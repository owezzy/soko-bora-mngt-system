package application

import (
	"context"
	"fmt"

	"github.com/owezzy/soko-bora-mngt-system/notifications/internal/models"
)

type (
	OrderCreated struct {
		OrderID    string
		CustomerID string
	}

	OrderCanceled struct {
		OrderID    string
		CustomerID string
	}

	OrderReady struct {
		OrderID    string
		CustomerID string
	}

	OrderCompleted struct {
		OrderID    string
		CustomerID string
	}

	App interface {
		NotifyOrderCreated(ctx context.Context, notify OrderCreated) error
		NotifyOrderCanceled(ctx context.Context, notify OrderCanceled) error
		NotifyOrderReady(ctx context.Context, notify OrderReady) error
		NotifyOrderCompleted(ctx context.Context, notify OrderCompleted) error
	}

	Application struct {
		customers     CustomerRepository
		notifications NotificationRepository
	}
)

var _ App = (*Application)(nil)

func New(customers CustomerRepository, notifications NotificationRepository) *Application {
	return &Application{
		customers:     customers,
		notifications: notifications,
	}
}

func (a Application) NotifyOrderCreated(ctx context.Context, notify OrderCreated) error {
	return a.notify(ctx, notify.OrderID, notify.CustomerID, "created", "Order %s was created for %s")
}

func (a Application) NotifyOrderCanceled(ctx context.Context, notify OrderCanceled) error {
	return a.notify(ctx, notify.OrderID, notify.CustomerID, "canceled", "Order %s was canceled for %s")
}

func (a Application) NotifyOrderReady(ctx context.Context, notify OrderReady) error {
	return a.notify(ctx, notify.OrderID, notify.CustomerID, "ready", "Order %s is ready for pickup for %s")
}

func (a Application) NotifyOrderCompleted(ctx context.Context, notify OrderCompleted) error {
	return a.notify(ctx, notify.OrderID, notify.CustomerID, "completed", "Order %s was completed for %s")
}

func (a Application) notify(ctx context.Context, orderID, customerID, notificationType, messageFormat string) error {
	customer, err := a.customers.Find(ctx, customerID)
	if err != nil {
		return err
	}

	return a.notifications.Save(ctx, &models.Notification{
		ID:         fmt.Sprintf("%s:%s", notificationType, orderID),
		OrderID:    orderID,
		CustomerID: customerID,
		Type:       notificationType,
		SMSNumber:  customer.SmsNumber,
		Message:    fmt.Sprintf(messageFormat, orderID, customer.Name),
	})
}
