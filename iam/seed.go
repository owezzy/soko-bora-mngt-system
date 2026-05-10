package iam

import (
	"context"

	"github.com/owezzy/soko-bora-mngt-system/iam/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/iam/internal/domain"
	"github.com/owezzy/soko-bora-mngt-system/internal/auth"
	"github.com/owezzy/soko-bora-mngt-system/internal/demo"
)

func seedDemoData(ctx context.Context, app application.App) error {
	bootstrap := demo.Spec()
	for _, spec := range []demo.PrincipalSpec{bootstrap.Auth, bootstrap.BotAuth} {
		roles := make([]auth.Role, 0, len(spec.Roles))
		for _, role := range spec.Roles {
			roles = append(roles, auth.Role(role))
		}

		if err := app.SeedPrincipal(ctx, domain.NewPrincipal(
			spec.ID,
			spec.Name,
			spec.Email,
			spec.Password,
			roles,
			spec.CustomerID,
			spec.Avatar,
			spec.Status,
			auth.TokenKind(spec.Kind),
		)); err != nil {
			return err
		}
	}

	return nil
}
