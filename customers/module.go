package customers

import (
	"context"

	"github.com/owezzy/soko-bora-mngt-system/customers/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/customers/internal/grpc"
	"github.com/owezzy/soko-bora-mngt-system/customers/internal/logging"
	"github.com/owezzy/soko-bora-mngt-system/customers/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/customers/internal/rest"
	"github.com/owezzy/soko-bora-mngt-system/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	var app application.App
	app = application.New(customers)
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
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
