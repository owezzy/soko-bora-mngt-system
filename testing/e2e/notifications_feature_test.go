//go:build e2e

package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cucumber/godog"
	"github.com/stackus/errors"
)

type notificationsFeature struct {
	db *sql.DB
}

func (c *notificationsFeature) init(cfg featureConfig) (err error) {
	if cfg.useMonoDB {
		c.db, err = sql.Open("pgx", "postgres://mallbots_user:mallbots_pass@localhost:5432/mallbots?sslmode=disable")
	} else {
		c.db, err = sql.Open("pgx", "postgres://notifications_user:notifications_pass@localhost:5432/notifications?sslmode=disable&search_path=notifications,public")
	}
	return err
}

func (c *notificationsFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^(?:I )?(?:ensure |expect )?a "([^"]*)" notification exists for the order$`, c.expectANotificationExistsForTheOrder)
	}

func (c *notificationsFeature) reset() {
	truncate := func(tableName string) {
		_, _ = c.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", tableName))
	}

	truncate("notifications.notifications")
	truncate("notifications.inbox")
	truncate("notifications.outbox")
	truncate("notifications.customers_cache")
}

func (c *notificationsFeature) expectANotificationExistsForTheOrder(ctx context.Context, notificationType string) error {
	orderID, err := lastBasketID(ctx)
	if err != nil {
		return err
	}

	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		var count int
		row := c.db.QueryRow("SELECT COUNT(*) FROM notifications.notifications WHERE order_id = $1 AND type = $2", orderID, notificationType)
		if err = row.Scan(&count); err == nil && count > 0 {
			return nil
		}
		time.Sleep(200 * time.Millisecond)
	}

	return errors.ErrNotFound.Msgf("notification `%s` for order `%s` was not recorded", notificationType, orderID)
}
