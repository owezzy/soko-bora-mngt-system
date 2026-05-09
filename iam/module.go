package iam

import (
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"

	"github.com/owezzy/soko-bora-mngt-system/iam/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/iam/internal/constants"
	"github.com/owezzy/soko-bora-mngt-system/iam/internal/domain"
	iahttp "github.com/owezzy/soko-bora-mngt-system/iam/internal/http"
	"github.com/owezzy/soko-bora-mngt-system/iam/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/internal/auth"
	"github.com/owezzy/soko-bora-mngt-system/internal/config"
	"github.com/owezzy/soko-bora-mngt-system/internal/di"
	"github.com/owezzy/soko-bora-mngt-system/internal/system"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, svc system.Service) error {
	container := di.New()
	configureOAuth(svc.Config().Auth)
	tokens := auth.NewTokenManager(svc.Config().Auth)
	exchangeCodes := iahttp.NewExchangeCodeStore()
	container.AddSingleton(constants.PrincipalsRepoKey, func(c di.Container) (any, error) {
		return postgres.NewPrincipalRepository(constants.PrincipalsTableName, svc.DB()), nil
	})
	container.AddSingleton(constants.ApplicationKey, func(c di.Container) (any, error) {
		return application.New(c.Get(constants.PrincipalsRepoKey).(domain.PrincipalRepository), tokens, application.Options{
			AdminEmails:       svc.Config().Auth.AdminEmailList(),
			BotClientID:       svc.Config().Auth.BotClientID,
			BotClientSecret:   svc.Config().Auth.BotClientSecret,
			BotPrincipalEmail: svc.Config().Auth.BotPrincipalEmail,
		}), nil
	})

	app := container.Get(constants.ApplicationKey).(*application.Application)
	if err := seedDemoData(ctx, app); err != nil {
		return err
	}

	iahttp.RegisterRoutes(svc.Mux(), app, auth.NewHTTPMiddleware(tokens), exchangeCodes, iahttp.OAuthConfig{
		FrontendRedirectURL: svc.Config().Auth.FrontendRedirectURL,
		GoogleEnabled:       googleEnabled(svc.Config().Auth),
	})
	return nil
}

func configureOAuth(cfg config.AuthConfig) {
	store := sessions.NewCookieStore([]byte(cfg.SessionSecret))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   strings.HasPrefix(cfg.GoogleCallbackURL, "https://"),
	}
	gothic.Store = store

	if !googleEnabled(cfg) {
		goth.UseProviders()
		return
	}

	goth.UseProviders(
		google.New(
			cfg.GoogleClientID,
			cfg.GoogleClientSecret,
			cfg.GoogleCallbackURL,
		),
	)
}

func googleEnabled(cfg config.AuthConfig) bool {
	return cfg.GoogleClientID != "" && cfg.GoogleClientSecret != "" && cfg.GoogleCallbackURL != ""
}
