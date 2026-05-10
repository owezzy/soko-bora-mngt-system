package domain

import (
	"context"
	"errors"
	"strings"

	"github.com/owezzy/soko-bora-mngt-system/internal/auth"
)

type Principal struct {
	ID           string
	Name         string
	Email        string
	Password     string
	Avatar       string
	Status       string
	Roles        []auth.Role
	CustomerID   string
	Provider     string
	ProviderUser string
	Kind         auth.TokenKind
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func NewPrincipal(id, name, email, password string, roles []auth.Role, customerID, avatar, status string, kind auth.TokenKind) Principal {
	return Principal{
		ID:         id,
		Name:       name,
		Email:      auth.NormalizeEmail(email),
		Password:   password,
		Roles:      append([]auth.Role(nil), roles...),
		CustomerID: customerID,
		Avatar:     avatar,
		Status:     status,
		Kind:       kind,
	}
}

func (p Principal) VerifyPassword(password string) error {
	if strings.TrimSpace(password) == "" || p.Password != password {
		return ErrInvalidCredentials
	}

	return nil
}

func (p Principal) UserResponse() map[string]any {
	return map[string]any{
		"id":     p.ID,
		"name":   p.Name,
		"email":  p.Email,
		"avatar": p.Avatar,
		"status": p.Status,
	}
}

type PrincipalRepository interface {
	FindByEmail(ctx context.Context, email string) (Principal, error)
	FindByID(ctx context.Context, id string) (Principal, error)
	FindByProvider(ctx context.Context, provider, providerUser string) (Principal, error)
	Save(ctx context.Context, principal Principal) error
	Update(ctx context.Context, principal Principal) error
}
