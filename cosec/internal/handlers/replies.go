package handlers

import (
	"context"

	"github.com/owezzy/soko-bora-mngt-system/cosec/internal/models"
	"github.com/owezzy/soko-bora-mngt-system/internal/am"
	"github.com/owezzy/soko-bora-mngt-system/internal/sec"
)

func RegisterReplyHandlers(subscriber am.ReplySubscriber, orchestrator sec.Orchestrator[*models.CreateOrderData]) error {
	replyMsgHandler := am.MessageHandlerFunc[am.IncomingReplyMessage](func(ctx context.Context, replyMsg am.IncomingReplyMessage) error {
		return orchestrator.HandleReply(ctx, replyMsg)
	})
	return subscriber.Subscribe(orchestrator.ReplyTopic(), replyMsgHandler, am.GroupName("cosec-replies"))
}
