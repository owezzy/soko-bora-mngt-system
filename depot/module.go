package depot

import (
	"context"

	"github.com/owezzy/soko-bora-mngt-system/depot/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/depot/internal/grpc"
	"github.com/owezzy/soko-bora-mngt-system/depot/internal/logging"
	"github.com/owezzy/soko-bora-mngt-system/depot/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/depot/internal/rest"
	"github.com/owezzy/soko-bora-mngt-system/internal/monolith"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	shoppingLists := postgres.NewShoppingListRepository("depot.shopping_lists", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	var app application.App
	app = application.New(shoppingLists, stores, products, orders)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup Driver adapters
	if err := grpc.Register(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
