package rest

import (
	"github.com/go-chi/chi/v5"

	"github.com/owezzy/soko-bora-mngt-system/internal/di"

	searchgrpc "github.com/owezzy/soko-bora-mngt-system/search/internal/grpc"
)

func RegisterConnect(container di.Container, mux *chi.Mux) error {
	path, handler := searchgrpc.NewConnectHandlerTx(container)
	mux.Mount(path, handler)
	return nil
}
