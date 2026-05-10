package config

import (
	"os"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/stackus/dotenv"

	"github.com/owezzy/soko-bora-mngt-system/internal/rpc"
	"github.com/owezzy/soko-bora-mngt-system/internal/web"
)

type (
	PGConfig struct {
		Conn string `required:"true"`
	}

	NatsConfig struct {
		URL    string `required:"true"`
		Stream string `default:"mallbots"`
	}

	OtelConfig struct {
		ServiceName      string `envconfig:"SERVICE_NAME" default:"mallbots"`
		ExporterEndpoint string `envconfig:"EXPORTER_OTLP_ENDPOINT" default:"http://collector:4317"`
	}

	AuthConfig struct {
		JWTIssuer           string        `envconfig:"AUTH_JWT_ISSUER" default:"mallbots"`
		JWTAudience         string        `envconfig:"AUTH_JWT_AUDIENCE" default:"soko-bora-web-app"`
		JWTSecret           string        `envconfig:"AUTH_JWT_SECRET" default:"mallbots-demo-secret-change-me"`
		AccessTokenTTL      time.Duration `envconfig:"AUTH_ACCESS_TOKEN_TTL" default:"24h"`
		SessionSecret       string        `envconfig:"AUTH_SESSION_SECRET" default:"mallbots-demo-session-secret-change-me-please-override"`
		FrontendRedirectURL string        `envconfig:"AUTH_FRONTEND_REDIRECT_URL" default:"http://localhost:4200/sign-in"`
		GoogleClientID      string        `envconfig:"AUTH_GOOGLE_CLIENT_ID"`
		GoogleClientSecret  string        `envconfig:"AUTH_GOOGLE_CLIENT_SECRET"`
		GoogleCallbackURL   string        `envconfig:"AUTH_GOOGLE_CALLBACK_URL" default:"http://localhost:8080/api/auth/google/callback"`
		AdminEmails         string        `envconfig:"AUTH_ADMIN_EMAILS" default:"hughes.brian@company.com"`
		BotClientID         string        `envconfig:"AUTH_BOT_CLIENT_ID" default:"mallbots-bot"`
		BotClientSecret     string        `envconfig:"AUTH_BOT_CLIENT_SECRET" default:"mallbots-bot-secret-change-me"`
		BotPrincipalEmail   string        `envconfig:"AUTH_BOT_PRINCIPAL_EMAIL" default:"bot@mallbots.internal"`
	}
	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		PG              PGConfig
		Nats            NatsConfig
		Rpc             rpc.RpcConfig
		Web             web.WebConfig
		Otel            OtelConfig
		Auth            AuthConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func (c AuthConfig) AdminEmailList() []string {
	if c.AdminEmails == "" {
		return nil
	}

	parts := strings.Split(c.AdminEmails, ",")
	adminEmails := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(strings.ToLower(part))
		if trimmed != "" {
			adminEmails = append(adminEmails, trimmed)
		}
	}

	return adminEmails
}

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
