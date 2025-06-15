package search

import (
	"context"
	"database/sql"

	"github.com/owezzy/soko-bora-mngt-system/customers/customerspb"
	"github.com/owezzy/soko-bora-mngt-system/internal/am"
	"github.com/owezzy/soko-bora-mngt-system/internal/amotel"
	"github.com/owezzy/soko-bora-mngt-system/internal/amprom"
	"github.com/owezzy/soko-bora-mngt-system/internal/di"
	"github.com/owezzy/soko-bora-mngt-system/internal/jetstream"
	pg "github.com/owezzy/soko-bora-mngt-system/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/internal/postgresotel"
	"github.com/owezzy/soko-bora-mngt-system/internal/registry"
	"github.com/owezzy/soko-bora-mngt-system/internal/system"
	"github.com/owezzy/soko-bora-mngt-system/internal/tm"
	"github.com/owezzy/soko-bora-mngt-system/ordering/orderingpb"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/constants"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/grpc"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/handlers"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/rest"
	"github.com/owezzy/soko-bora-mngt-system/stores/storespb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
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
	stream := jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return svc.DB().Begin()
	})
	container.AddSingleton(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			stream,
			amotel.OtelMessageContextExtractor(),
			amprom.ReceivedMessagesCounter(constants.ServiceName),
		), nil
	})
	container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		tx := postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx))
		return pg.NewInboxStore(constants.InboxTableName, tx), nil
	})
	container.AddScoped(constants.CustomersRepoKey, func(c di.Container) (any, error) {
		return postgres.NewCustomerCacheRepository(
			constants.CustomersCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewCustomerRepository(svc.Config().Rpc.Service(constants.CustomersServiceName)),
		), nil
	})
	container.AddScoped(constants.StoresRepoKey, func(c di.Container) (any, error) {
		return postgres.NewStoreCacheRepository(
			constants.StoresCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewStoreRepository(svc.Config().Rpc.Service(constants.StoresServiceName)),
		), nil
	})
	container.AddScoped(constants.ProductsRepoKey, func(c di.Container) (any, error) {
		return postgres.NewProductCacheRepository(
			constants.ProductsCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewProductRepository(svc.Config().Rpc.Service(constants.StoresServiceName)),
		), nil
	})
	container.AddScoped(constants.OrdersRepoKey, func(c di.Container) (any, error) {
		return postgres.NewOrderRepository(
			constants.OrdersTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
		), nil
	})

	// setup application
	container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		return application.New(
			c.Get(constants.OrdersRepoKey).(application.OrderRepository),
		), nil
	})
	container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.OrdersRepoKey).(application.OrderRepository),
			c.Get(constants.CustomersRepoKey).(application.CustomerCacheRepository),
			c.Get(constants.StoresRepoKey).(application.StoreCacheRepository),
			c.Get(constants.ProductsRepoKey).(application.ProductCacheRepository),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})

	// setup Driver adapters
	if err = grpc.RegisterServerTx(container, svc.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, svc.Mux(), svc.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(svc.Mux()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}

	return nil
}
