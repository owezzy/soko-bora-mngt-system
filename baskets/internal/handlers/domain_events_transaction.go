package handlers

import (
	"context"
	"github.com/owezzy/soko-bora-mngt-system/baskets/internal/constants"
	"github.com/owezzy/soko-bora-mngt-system/internal/ddd"
	"github.com/owezzy/soko-bora-mngt-system/internal/di"
)

func RegisterDomainEventHandlersTx(container di.Container) {
	handlers := ddd.EventHandlerFunc[ddd.Event](func(ctx context.Context, event ddd.Event) error {
		domainHandlers := di.Get(ctx, constants.DomainEventHandlersKey).(ddd.EventHandler[ddd.Event])

		return domainHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get(constants.DomainDispatcherKey).(*ddd.EventDispatcher[ddd.Event])

	RegisterDomainEventHandlers(subscriber, handlers)
}
