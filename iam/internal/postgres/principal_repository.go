package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/owezzy/soko-bora-mngt-system/iam/internal/domain"
	"github.com/owezzy/soko-bora-mngt-system/internal/auth"
	sharedpostgres "github.com/owezzy/soko-bora-mngt-system/internal/postgres"
)

type PrincipalRepository struct {
	tableName string
	db        sharedpostgres.DB
}

func NewPrincipalRepository(tableName string, db sharedpostgres.DB) PrincipalRepository {
	return PrincipalRepository{tableName: tableName, db: db}
}

func (r PrincipalRepository) FindByEmail(ctx context.Context, email string) (domain.Principal, error) {
	const query = `SELECT id, name, email, password, avatar, status, roles, customer_id, provider, provider_user, kind FROM %s WHERE email = $1 LIMIT 1`
	return r.scanOne(ctx, fmt.Sprintf(query, r.tableName), email)
}

func (r PrincipalRepository) FindByID(ctx context.Context, id string) (domain.Principal, error) {
	const query = `SELECT id, name, email, password, avatar, status, roles, customer_id, provider, provider_user, kind FROM %s WHERE id = $1 LIMIT 1`
	return r.scanOne(ctx, fmt.Sprintf(query, r.tableName), id)
}

func (r PrincipalRepository) FindByProvider(ctx context.Context, provider, providerUser string) (domain.Principal, error) {
	const query = `SELECT id, name, email, password, avatar, status, roles, customer_id, provider, provider_user, kind FROM %s WHERE provider = $1 AND provider_user = $2 LIMIT 1`
	var principal domain.Principal
	var roles string
	var kind string
	err := r.db.QueryRowContext(ctx, fmt.Sprintf(query, r.tableName), provider, providerUser).Scan(&principal.ID, &principal.Name, &principal.Email, &principal.Password, &principal.Avatar, &principal.Status, &roles, &principal.CustomerID, &principal.Provider, &principal.ProviderUser, &kind)
	if err != nil {
		return domain.Principal{}, err
	}
	principal.Email = auth.NormalizeEmail(principal.Email)
	principal.Roles = parseRoles(roles)
	principal.Kind = auth.TokenKind(kind)
	return principal, nil
}

func (r PrincipalRepository) Save(ctx context.Context, principal domain.Principal) error {
	const query = `INSERT INTO %s (id, name, email, password, avatar, status, roles, customer_id, provider, provider_user, kind) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	_, err := r.db.ExecContext(ctx, fmt.Sprintf(query, r.tableName), principal.ID, principal.Name, principal.Email, principal.Password, principal.Avatar, principal.Status, strings.Join(roleStrings(principal.Roles), ","), principal.CustomerID, principal.Provider, principal.ProviderUser, string(principal.Kind))
	return err
}

func (r PrincipalRepository) Update(ctx context.Context, principal domain.Principal) error {
	const query = `UPDATE %s SET name = $2, email = $3, password = $4, avatar = $5, status = $6, roles = $7, customer_id = $8, provider = $9, provider_user = $10, kind = $11 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, fmt.Sprintf(query, r.tableName), principal.ID, principal.Name, principal.Email, principal.Password, principal.Avatar, principal.Status, strings.Join(roleStrings(principal.Roles), ","), principal.CustomerID, principal.Provider, principal.ProviderUser, string(principal.Kind))
	return err
}

func (r PrincipalRepository) scanOne(ctx context.Context, query string, arg string) (domain.Principal, error) {
	var principal domain.Principal
	var roles string
	var kind string
	err := r.db.QueryRowContext(ctx, query, arg).Scan(&principal.ID, &principal.Name, &principal.Email, &principal.Password, &principal.Avatar, &principal.Status, &roles, &principal.CustomerID, &principal.Provider, &principal.ProviderUser, &kind)
	if err != nil {
		return domain.Principal{}, err
	}
	principal.Email = auth.NormalizeEmail(principal.Email)
	principal.Roles = parseRoles(roles)
	principal.Kind = auth.TokenKind(kind)
	return principal, nil
}

func parseRoles(raw string) []auth.Role {
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	roles := make([]auth.Role, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			roles = append(roles, auth.Role(trimmed))
		}
	}
	return roles
}

func roleStrings(roles []auth.Role) []string {
	values := make([]string, 0, len(roles))
	for _, role := range roles {
		values = append(values, string(role))
	}
	return values
}
