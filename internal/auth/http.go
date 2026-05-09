package auth

import (
	"encoding/json"
	"net/http"
	"strings"
)

type HTTPMiddleware struct {
	tokens *TokenManager
}

func NewHTTPMiddleware(tokens *TokenManager) HTTPMiddleware {
	return HTTPMiddleware{tokens: tokens}
}

func (m HTTPMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := m.claimsFromRequest(r)
		if err != nil {
			WriteJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
			return
		}

		next.ServeHTTP(w, r.WithContext(WithClaims(r.Context(), claims)))
	})
}

func (m HTTPMiddleware) RequireRoles(roles ...Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return m.RequireAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := ClaimsFromContext(r.Context())
			if !ok {
				WriteJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
				return
			}

			for _, role := range roles {
				if !claims.HasRole(role) {
					WriteJSON(w, http.StatusForbidden, map[string]any{"error": "forbidden"})
					return
				}
			}

			next.ServeHTTP(w, r)
		}))
	}
}

func (m HTTPMiddleware) claimsFromRequest(r *http.Request) (Claims, error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return Claims{}, http.ErrNoCookie
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authorization, bearerPrefix) {
		return Claims{}, http.ErrNoCookie
	}

	return m.tokens.Parse(strings.TrimSpace(strings.TrimPrefix(authorization, bearerPrefix)))
}

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
