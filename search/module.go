package search

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"github.com/owezzy/soko-bora-mngt-system/customers/customerspb"
	"github.com/owezzy/soko-bora-mngt-system/internal/ddd"
	"github.com/owezzy/soko-bora-mngt-system/internal/di"
	"github.com/owezzy/soko-bora-mngt-system/internal/jetstream"
	"github.com/owezzy/soko-bora-mngt-system/internal/monolith"
	pg "github.com/owezzy/soko-bora-mngt-system/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/internal/registry"
	"github.com/owezzy/soko-bora-mngt-system/internal/tm"
	"github.com/owezzy/soko-bora-mngt-system/ordering/orderingpb"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/grpc"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/handlers"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/logging"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/rest"
	"github.com/owezzy/soko-bora-mngt-system/stores/storespb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := orderingpb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := customerspb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := storespb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return mono.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), c.Get("logger").(zerolog.Logger)), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return mono.DB(), nil
	})
	container.AddSingleton("conn", func(c di.Container) (any, error) {
		return grpc.Dial(ctx, mono.Config().Rpc.Address())
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*sql.DB)
		return db.Begin()
	})
	container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		inboxStore := pg.NewInboxStore("search.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	container.AddScoped("customers", func(c di.Container) (any, error) {
		return postgres.NewCustomerCacheRepository(
			"search.customers_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewCustomerRepository(c.Get("conn").(*grpc.ClientConn)),
		), nil
	})
	container.AddScoped("stores", func(c di.Container) (any, error) {
		return postgres.NewStoreCacheRepository(
			"search.stores_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewStoreRepository(c.Get("conn").(*grpc.ClientConn)),
		), nil
	})
	container.AddScoped("products", func(c di.Container) (any, error) {
		return postgres.NewProductCacheRepository(
			"search.products_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewProductRepository(c.Get("conn").(*grpc.ClientConn)),
		), nil
	})
	container.AddScoped("orders", func(c di.Container) (any, error) {
		return postgres.NewOrderRepository("search.orders", c.Get("tx").(*sql.Tx)), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.LogApplicationAccess(
			application.New(
				c.Get("orders").(application.OrderRepository),
			),
			c.Get("logger").(zerolog.Logger),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return logging.LogEventHandlerAccess[ddd.Event](
			handlers.NewIntegrationEventHandlers(
				c.Get("orders").(application.OrderRepository),
				c.Get("customers").(application.CustomerCacheRepository),
				c.Get("stores").(application.StoreCacheRepository),
				c.Get("products").(application.ProductCacheRepository),
			),
			"IntegrationEvents", c.Get("logger").(zerolog.Logger),
		), nil
	})

	// setup Driver adapters
	if err = grpc.RegisterServerTx(container, mono.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}

	return nil
}
