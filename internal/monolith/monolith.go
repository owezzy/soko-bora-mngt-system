package monolith

import (
	"context"
	"database/sql"
	"github.com/nats-io/nats.go"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/owezzy/soko-bora-mngt-system/internal/config"
	"github.com/owezzy/soko-bora-mngt-system/internal/waiter"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
	JS() nats.JetStreamContext
	Logger() zerolog.Logger
	Mux() *chi.Mux
	RPC() *grpc.Server
	Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, Monolith) error
}
