package application

import (
	"context"
	"crypto/subtle"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/owezzy/soko-bora-mngt-system/iam/internal/domain"
	"github.com/owezzy/soko-bora-mngt-system/internal/auth"
)

type App interface {
	SignIn(ctx context.Context, email, password string) (AuthSession, error)
	SignInBot(ctx context.Context, clientID, clientSecret string) (AuthSession, error)
	SignInWithToken(ctx context.Context, token string) (AuthSession, error)
	SignInWithOAuth(ctx context.Context, principal OAuthPrincipal) (AuthSession, error)
	CurrentUser(ctx context.Context) (map[string]any, error)
	SeedPrincipal(ctx context.Context, principal domain.Principal) error
}

type Options struct {
	AdminEmails       []string
	BotClientID       string
	BotClientSecret   string
	BotPrincipalEmail string
}

type OAuthPrincipal struct {
	Provider     string
	ProviderUser string
	Email        string
	Name         string
	Avatar       string
}

type AuthSession struct {
	User        map[string]any `json:"user"`
	AccessToken string         `json:"accessToken"`
	TokenType   string         `json:"tokenType"`
}

type Application struct {
	principals domain.PrincipalRepository
	tokens     *auth.TokenManager
	now        func() time.Time
	newID      func() string
	adminEmails map[string]struct{}
	botClientID string
	botClientSecret string
	botPrincipalEmail string
}


func New(principals domain.PrincipalRepository, tokens *auth.TokenManager, options Options) *Application {
	adminEmailSet := make(map[string]struct{}, len(options.AdminEmails))
	for _, email := range options.AdminEmails {
		normalized := auth.NormalizeEmail(email)
		if normalized != "" {
			adminEmailSet[normalized] = struct{}{}
		}
	}

	return &Application{
		principals: principals,
		tokens:     tokens,
		now:        time.Now,
		newID:      uuid.NewString,
		adminEmails: adminEmailSet,
		botClientID: options.BotClientID,
		botClientSecret: options.BotClientSecret,
		botPrincipalEmail: auth.NormalizeEmail(options.BotPrincipalEmail),
	}
}

func (a *Application) SignIn(ctx context.Context, email, password string) (AuthSession, error) {
	principal, err := a.principals.FindByEmail(ctx, auth.NormalizeEmail(email))
	if err != nil {
		return AuthSession{}, err
	}

	if err := principal.VerifyPassword(password); err != nil {
		return AuthSession{}, err
	}

	return a.sessionForPrincipal(principal)
}

func (a *Application) SignInBot(ctx context.Context, clientID, clientSecret string) (AuthSession, error) {
	if subtle.ConstantTimeCompare([]byte(clientID), []byte(a.botClientID)) != 1 || subtle.ConstantTimeCompare([]byte(clientSecret), []byte(a.botClientSecret)) != 1 {
		return AuthSession{}, domain.ErrInvalidCredentials
	}

	principal, err := a.principals.FindByEmail(ctx, a.botPrincipalEmail)
	if err != nil {
		return AuthSession{}, err
	}

	return a.sessionForPrincipal(principal)
}

func (a *Application) SignInWithToken(ctx context.Context, token string) (AuthSession, error) {
	claims, err := a.tokens.Parse(token)
	if err != nil {
		return AuthSession{}, err
	}

	principal, err := a.principals.FindByID(ctx, claims.RegisteredClaims.Subject)
	if err != nil {
		return AuthSession{}, err
	}

	return a.sessionForPrincipal(principal)
}

func (a *Application) SignInWithOAuth(ctx context.Context, oauthPrincipal OAuthPrincipal) (AuthSession, error) {
	principal, err := a.upsertOAuthPrincipal(ctx, oauthPrincipal)
	if err != nil {
		return AuthSession{}, err
	}

	return a.sessionForPrincipal(principal)
}

func (a *Application) CurrentUser(ctx context.Context) (map[string]any, error) {
	claims, ok := auth.ClaimsFromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}

	principal, err := a.principals.FindByID(ctx, claims.RegisteredClaims.Subject)
	if err != nil {
		return nil, err
	}

	return principal.UserResponse(), nil
}

func (a *Application) SeedPrincipal(ctx context.Context, principal domain.Principal) error {
	_, err := a.principals.FindByID(ctx, principal.ID)
	if err == nil {
		return a.principals.Update(ctx, principal)
	}

	return a.principals.Save(ctx, principal)
}

func (a *Application) upsertOAuthPrincipal(ctx context.Context, oauthPrincipal OAuthPrincipal) (domain.Principal, error) {
	provider := oauthPrincipal.Provider
	providerUser := oauthPrincipal.ProviderUser
	email := auth.NormalizeEmail(oauthPrincipal.Email)

	if provider != "" && providerUser != "" {
		principal, err := a.principals.FindByProvider(ctx, provider, providerUser)
		switch {
		case err == nil:
			return a.refreshOAuthPrincipal(ctx, principal, oauthPrincipal)
		case !errors.Is(err, sql.ErrNoRows):
			return domain.Principal{}, err
		}
	}

	if email != "" {
		principal, err := a.principals.FindByEmail(ctx, email)
		switch {
		case err == nil:
			principal.Provider = provider
			principal.ProviderUser = providerUser
			return a.refreshOAuthPrincipal(ctx, principal, oauthPrincipal)
		case !errors.Is(err, sql.ErrNoRows):
			return domain.Principal{}, err
		}
	}

	principal := domain.NewPrincipal(
		a.newID(),
		a.displayName(oauthPrincipal),
		email,
		"",
		a.rolesForEmail(email),
		"",
		oauthPrincipal.Avatar,
		"online",
		auth.TokenKindUser,
	)
	principal.Provider = provider
	principal.ProviderUser = providerUser

	if err := a.principals.Save(ctx, principal); err != nil {
		return domain.Principal{}, err
	}

	return principal, nil
}

func (a *Application) refreshOAuthPrincipal(ctx context.Context, principal domain.Principal, oauthPrincipal OAuthPrincipal) (domain.Principal, error) {
	updated := false
	name := a.displayName(oauthPrincipal)
	if name != "" && principal.Name != name {
		principal.Name = name
		updated = true
	}

	avatar := oauthPrincipal.Avatar
	if avatar != "" && principal.Avatar != avatar {
		principal.Avatar = avatar
		updated = true
	}

	email := auth.NormalizeEmail(oauthPrincipal.Email)
	if email != "" && principal.Email != email {
		principal.Email = email
		updated = true
	}

	if oauthPrincipal.Provider != "" && principal.Provider != oauthPrincipal.Provider {
		principal.Provider = oauthPrincipal.Provider
		updated = true
	}

	if oauthPrincipal.ProviderUser != "" && principal.ProviderUser != oauthPrincipal.ProviderUser {
		principal.ProviderUser = oauthPrincipal.ProviderUser
		updated = true
	}

	if principal.Status == "" {
		principal.Status = "online"
		updated = true
	}

	if !updated {
		return principal, nil
	}

	if err := a.principals.Update(ctx, principal); err != nil {
		return domain.Principal{}, err
	}

	return principal, nil
}

func (a *Application) displayName(oauthPrincipal OAuthPrincipal) string {
	if oauthPrincipal.Name != "" {
		return oauthPrincipal.Name
	}

	return auth.NormalizeEmail(oauthPrincipal.Email)
}

func (a *Application) rolesForEmail(email string) []auth.Role {
	roles := []auth.Role{auth.RoleCustomer}
	if _, ok := a.adminEmails[auth.NormalizeEmail(email)]; ok {
		roles = append(roles, auth.RoleAdmin)
	}

	return roles
}

func (a *Application) sessionForPrincipal(principal domain.Principal) (AuthSession, error) {
	token, _, err := a.tokens.Issue(auth.Principal{
		Subject:    principal.ID,
		Email:      principal.Email,
		Name:       principal.Name,
		Roles:      principal.Roles,
		CustomerID: principal.CustomerID,
		Kind:       principal.Kind,
	}, a.now())
	if err != nil {
		return AuthSession{}, err
	}

	return AuthSession{
		User:        principal.UserResponse(),
		AccessToken: token,
		TokenType:   "bearer",
	}, nil
}
