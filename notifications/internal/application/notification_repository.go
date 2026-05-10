package application

import (
	"context"

	"github.com/owezzy/soko-bora-mngt-system/notifications/internal/models"
)

type NotificationRepository interface {
	Save(ctx context.Context, notification *models.Notification) error
}
