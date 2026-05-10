package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/owezzy/soko-bora-mngt-system/internal/config"
)

type Principal struct {
	Subject    string
	Email      string
	Name       string
	Roles      []Role
	CustomerID string
	Kind       TokenKind
}

type TokenManager struct {
	issuer    string
	audience  string
	secretKey []byte
	ttl       time.Duration
}

func NewTokenManager(cfg config.AuthConfig) *TokenManager {
	return &TokenManager{
		issuer:    cfg.JWTIssuer,
		audience:  cfg.JWTAudience,
		secretKey: []byte(cfg.JWTSecret),
		ttl:       cfg.AccessTokenTTL,
	}
}

func (m *TokenManager) Issue(principal Principal, now time.Time) (string, Claims, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			Audience:  jwt.ClaimStrings{m.audience},
			Subject:   principal.Subject,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.ttl)),
		},
		Kind:       principal.Kind,
		Email:      principal.Email,
		Name:       principal.Name,
		Roles:      append([]Role(nil), principal.Roles...),
		CustomerID: principal.CustomerID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", Claims{}, err
	}

	return signed, claims, nil
}

func (m *TokenManager) Parse(raw string) (Claims, error) {
	token, err := jwt.ParseWithClaims(raw, &Claims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Method.Alg())
		}

		return m.secretKey, nil
	}, jwt.WithAudience(m.audience), jwt.WithIssuer(m.issuer))
	if err != nil {
		return Claims{}, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return Claims{}, fmt.Errorf("invalid token claims")
	}

	return *claims, nil
}
