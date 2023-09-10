package ordering

import (
	"context"

	"github.com/owezzy/soko-bora-mngt-system/internal/monolith"
	"github.com/owezzy/soko-bora-mngt-system/ordering/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/ordering/internal/grpc"
	"github.com/owezzy/soko-bora-mngt-system/ordering/internal/logging"
	"github.com/owezzy/soko-bora-mngt-system/ordering/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/ordering/internal/rest"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	orders := postgres.NewOrderRepository("ordering.orders", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	invoices := grpc.NewInvoiceRepository(conn)
	shopping := grpc.NewShoppingListRepository(conn)
	notifications := grpc.NewNotificationRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, customers, payments, invoices, shopping, notifications)
	app = logging.NewApplication(app, mono.Logger())

	// setup Driver adapters
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
