package http

import (
	"encoding/json"
	stdhttp "net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"

	"github.com/owezzy/soko-bora-mngt-system/iam/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/internal/auth"
)

type OAuthConfig struct {
	FrontendRedirectURL string
	GoogleEnabled       bool
}

type Server struct {
	app           application.App
	middleware    auth.HTTPMiddleware
	exchangeCodes *ExchangeCodeStore
	oauthConfig   OAuthConfig
}

func RegisterRoutes(mux *chi.Mux, app application.App, middleware auth.HTTPMiddleware, exchangeCodes *ExchangeCodeStore, oauthConfig OAuthConfig) {
	server := Server{app: app, middleware: middleware, exchangeCodes: exchangeCodes, oauthConfig: oauthConfig}
	mux.Get("/api/auth/google/login", server.beginOAuth)
	mux.Get("/api/auth/{provider}", server.beginOAuth)
	mux.Get("/api/auth/{provider}/callback", server.completeOAuth)
	mux.Post("/api/auth/bot/token", server.signInBot)
	mux.Post("/api/auth/exchange-code", server.exchangeCode)
	mux.Post("/api/auth/sign-in", server.signIn)
	mux.Post("/api/auth/sign-in-with-token", server.signInWithToken)
	mux.Post("/api/auth/logout", server.logout)
	mux.Get("/api/auth/session", middleware.RequireAuth(stdhttp.HandlerFunc(server.session)).ServeHTTP)
	mux.Get("/api/common/user", middleware.RequireAuth(stdhttp.HandlerFunc(server.currentUser)).ServeHTTP)
	mux.Patch("/api/common/user", middleware.RequireAuth(stdhttp.HandlerFunc(server.currentUser)).ServeHTTP)
}

func (s Server) signIn(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		auth.WriteJSON(w, stdhttp.StatusBadRequest, map[string]any{"error": "invalid request"})
		return
	}

	session, err := s.app.SignIn(r.Context(), credentials.Email, credentials.Password)
	if err != nil {
		auth.WriteJSON(w, stdhttp.StatusUnauthorized, map[string]any{"error": "invalid credentials"})
		return
	}

	auth.WriteJSON(w, stdhttp.StatusOK, session)
}

func (s Server) signInWithToken(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	var body struct {
		AccessToken string `json:"accessToken"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		auth.WriteJSON(w, stdhttp.StatusBadRequest, map[string]any{"error": "invalid request"})
		return
	}

	session, err := s.app.SignInWithToken(r.Context(), body.AccessToken)
	if err != nil {
		auth.WriteJSON(w, stdhttp.StatusUnauthorized, map[string]any{"error": "invalid token"})
		return
	}

	auth.WriteJSON(w, stdhttp.StatusOK, session)
}

func (s Server) signInBot(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	var body struct {
		ClientID     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		auth.WriteJSON(w, stdhttp.StatusBadRequest, map[string]any{"error": "invalid request"})
		return
	}

	session, err := s.app.SignInBot(r.Context(), body.ClientID, body.ClientSecret)
	if err != nil {
		auth.WriteJSON(w, stdhttp.StatusUnauthorized, map[string]any{"error": "invalid credentials"})
		return
	}

	auth.WriteJSON(w, stdhttp.StatusOK, session)
}

func (s Server) beginOAuth(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	provider := oauthProviderFromRequest(r)
	if !s.providerEnabled(provider) {
		auth.WriteJSON(w, stdhttp.StatusNotFound, map[string]any{"error": "provider not available"})
		return
	}

	s.storeRedirectURL(w, r)
	gothic.BeginAuthHandler(w, withProviderQuery(r, provider))
}

func (s Server) completeOAuth(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	provider := oauthProviderFromRequest(r)
	if !s.providerEnabled(provider) {
		auth.WriteJSON(w, stdhttp.StatusNotFound, map[string]any{"error": "provider not available"})
		return
	}

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		s.redirectToFrontend(w, r, "", "google_sign_in_failed")
		return
	}

	if strings.TrimSpace(user.UserID) == "" {
		s.redirectToFrontend(w, r, "", "google_sign_in_failed")
		return
	}

	session, err := s.app.SignInWithOAuth(r.Context(), application.OAuthPrincipal{
		Provider:     user.Provider,
		ProviderUser: user.UserID,
		Email:        user.Email,
		Name:         user.Name,
		Avatar:       user.AvatarURL,
	})
	if err != nil {
		s.redirectToFrontend(w, r, "", "google_sign_in_failed")
		return
	}

	code := s.exchangeCodes.Issue(session, 30*time.Second)
	s.redirectToFrontend(w, r, code, "")
}

func (s Server) exchangeCode(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	var body struct {
		Code string `json:"code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		auth.WriteJSON(w, stdhttp.StatusBadRequest, map[string]any{"error": "invalid request"})
		return
	}

	session, ok := s.exchangeCodes.Consume(body.Code)
	if !ok {
		auth.WriteJSON(w, stdhttp.StatusUnauthorized, map[string]any{"error": "invalid exchange code"})
		return
	}

	auth.WriteJSON(w, stdhttp.StatusOK, session)
}

func (s Server) session(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	user, err := s.app.CurrentUser(r.Context())
	if err != nil {
		auth.WriteJSON(w, stdhttp.StatusUnauthorized, map[string]any{"error": "unauthorized"})
		return
	}

	auth.WriteJSON(w, stdhttp.StatusOK, map[string]any{"user": user})
}

func (s Server) currentUser(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	user, err := s.app.CurrentUser(r.Context())
	if err != nil {
		auth.WriteJSON(w, stdhttp.StatusUnauthorized, map[string]any{"error": "unauthorized"})
		return
	}

	auth.WriteJSON(w, stdhttp.StatusOK, user)
}

func (s Server) logout(w stdhttp.ResponseWriter, _ *stdhttp.Request) {
	auth.WriteJSON(w, stdhttp.StatusOK, true)
}

func (s Server) providerEnabled(provider string) bool {
	return provider == "google" && s.oauthConfig.GoogleEnabled
}

func oauthProviderFromRequest(r *stdhttp.Request) string {
	provider := chi.URLParam(r, "provider")
	if provider != "" {
		return provider
	}

	if strings.HasSuffix(r.URL.Path, "/google/login") {
		return "google"
	}

	return ""
}

func (s Server) redirectToFrontend(w stdhttp.ResponseWriter, r *stdhttp.Request, code, authError string) {
	target, err := url.Parse(s.oauthConfig.FrontendRedirectURL)
	if err != nil || target.String() == "" {
		auth.WriteJSON(w, stdhttp.StatusInternalServerError, map[string]any{"error": "invalid frontend redirect"})
		return
	}

	redirectURL := s.consumeRedirectURL(w, r)
	query := target.Query()
	if code != "" {
		query.Set("auth_code", code)
	}
	if authError != "" {
		query.Set("auth_error", authError)
	}
	if redirectURL != "" {
		query.Set("redirectURL", redirectURL)
	}
	target.RawQuery = query.Encode()
	stdhttp.Redirect(w, r, target.String(), stdhttp.StatusSeeOther)
}

func (s Server) storeRedirectURL(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	redirectURL := strings.TrimSpace(r.URL.Query().Get("redirectURL"))
	if redirectURL == "" {
		return
	}

	stdhttp.SetCookie(w, &stdhttp.Cookie{
		Name:     "auth_redirect_url",
		Value:    url.QueryEscape(redirectURL),
		Path:     "/",
		HttpOnly: true,
		SameSite: stdhttp.SameSiteLaxMode,
		MaxAge:   300,
		Secure:   strings.HasPrefix(s.oauthConfig.FrontendRedirectURL, "https://"),
	})
}

func (s Server) consumeRedirectURL(w stdhttp.ResponseWriter, r *stdhttp.Request) string {
	cookie, err := r.Cookie("auth_redirect_url")
	if err != nil {
		return ""
	}

	stdhttp.SetCookie(w, &stdhttp.Cookie{
		Name:     "auth_redirect_url",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: stdhttp.SameSiteLaxMode,
		MaxAge:   -1,
		Secure:   strings.HasPrefix(s.oauthConfig.FrontendRedirectURL, "https://"),
	})

	redirectURL, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return ""
	}

	return redirectURL
}

func withProviderQuery(r *stdhttp.Request, provider string) *stdhttp.Request {
	query := r.URL.Query()
	query.Set("provider", provider)
	clone := r.Clone(r.Context())
	clone.URL.RawQuery = query.Encode()
	return clone
}
