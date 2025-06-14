package handlers

import (
	"context"
	"database/sql"

	"github.com/owezzy/soko-bora-mngt-system/internal/am"
	"github.com/owezzy/soko-bora-mngt-system/internal/ddd"
	"github.com/owezzy/soko-bora-mngt-system/internal/di"
	"github.com/owezzy/soko-bora-mngt-system/internal/registry"
)

func RegisterCommandHandlersTx(container di.Container) error {
	cmdMsgHandlers := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) (err error) {
		ctx = container.Scoped(ctx)
		defer func(tx *sql.Tx) {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p)
			} else if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}(di.Get(ctx, "tx").(*sql.Tx))

		cmdMsgHandlers := am.RawMessageHandlerWithMiddleware(
			am.NewCommandMessageHandler(
				di.Get(ctx, "registry").(registry.Registry),
				di.Get(ctx, "replyStream").(am.ReplyStream),
				di.Get(ctx, "commandHandlers").(ddd.CommandHandler[ddd.Command]),
			).(am.RawMessageHandler),
			di.Get(ctx, "inboxMiddleware").(am.RawMessageHandlerMiddleware),
		)

		return cmdMsgHandlers.HandleMessage(ctx, msg)
	})

	subscriber := container.Get("stream").(am.RawMessageStream)

	return RegisterCommandHandlers(subscriber, cmdMsgHandlers)
}
