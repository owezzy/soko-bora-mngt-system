package postgres

import (
	"context"
	"fmt"

	"github.com/owezzy/soko-bora-mngt-system/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/notifications/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/notifications/internal/models"
)

type NotificationRepository struct {
	tableName string
	db        postgres.DB
}

var _ application.NotificationRepository = (*NotificationRepository)(nil)

func NewNotificationRepository(tableName string, db postgres.DB) NotificationRepository {
	return NotificationRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r NotificationRepository) Save(ctx context.Context, notification *models.Notification) error {
	const query = `INSERT INTO %s (id, order_id, customer_id, type, sms_number, message)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (id) DO UPDATE SET
	  order_id = EXCLUDED.order_id,
	  customer_id = EXCLUDED.customer_id,
	  type = EXCLUDED.type,
	  sms_number = EXCLUDED.sms_number,
	  message = EXCLUDED.message`

	_, err := r.db.ExecContext(ctx, r.table(query),
		notification.ID,
		notification.OrderID,
		notification.CustomerID,
		notification.Type,
		notification.SMSNumber,
		notification.Message,
	)

	return err
}

func (r NotificationRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
