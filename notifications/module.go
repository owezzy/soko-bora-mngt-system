package notifications

import (
	"context"
	"github.com/owezzy/soko-bora-mngt-system/customers/customerspb"
	"github.com/owezzy/soko-bora-mngt-system/internal/jetstream"
	"github.com/owezzy/soko-bora-mngt-system/internal/monolith"
	"github.com/owezzy/soko-bora-mngt-system/internal/registry"
	"github.com/owezzy/soko-bora-mngt-system/notifications/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/notifications/internal/grpc"
	"github.com/owezzy/soko-bora-mngt-system/notifications/internal/handlers"
	"github.com/owezzy/soko-bora-mngt-system/notifications/internal/logging"
	"github.com/owezzy/soko-bora-mngt-system/notifications/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/ordering/orderingpb"

	"github.com/owezzy/soko-bora-mngt-system/internal/am"
	"github.com/owezzy/soko-bora-mngt-system/internal/ddd"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger()))
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := postgres.NewCustomerCacheRepository("notifications.customers_cache", mono.DB(), grpc.NewCustomerRepository(conn))

	// setup application
	app := logging.LogApplicationAccess(
		application.New(customers),
		mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(app, customers),
		"IntegrationEvents", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}

	return nil
}
