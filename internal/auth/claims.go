package auth

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Role string

const (
	RoleCustomer Role = "customer"
	RoleAdmin    Role = "admin"
	RoleBot      Role = "bot"
)

type TokenKind string

const (
	TokenKindUser TokenKind = "user"
	TokenKindBot  TokenKind = "bot"
)

type Claims struct {
	jwt.RegisteredClaims
	Kind       TokenKind `json:"kind"`
	Email      string    `json:"email,omitempty"`
	Name       string    `json:"name,omitempty"`
	Roles      []Role    `json:"roles"`
	CustomerID string    `json:"customer_id,omitempty"`
}

func (c Claims) HasRole(role Role) bool {
	for _, candidate := range c.Roles {
		if candidate == role {
			return true
		}
	}

	return false
}

func (c Claims) ExpiresIn(now time.Time) time.Duration {
	if c.ExpiresAt == nil {
		return 0
	}

	return c.ExpiresAt.Time.Sub(now)
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
