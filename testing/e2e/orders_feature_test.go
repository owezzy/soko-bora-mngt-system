//go:build e2e

package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cucumber/godog"
	"github.com/go-openapi/strfmt"
	"github.com/stackus/errors"

	"github.com/owezzy/soko-bora-mngt-system/ordering/orderingclient"
	"github.com/owezzy/soko-bora-mngt-system/ordering/orderingclient/models"
	"github.com/owezzy/soko-bora-mngt-system/ordering/orderingclient/order"
)

type ordersFeature struct {
	client *orderingclient.OrderProcessing
	db     *sql.DB
}

func (c *ordersFeature) init(cfg featureConfig) (err error) {
	if cfg.useMonoDB {
		c.db, err = sql.Open("pgx", "postgres://mallbots_user:mallbots_pass@localhost:5432/mallbots?sslmode=disable")
	} else {
		c.db, err = sql.Open("pgx", "postgres://ordering_user:ordering_pass@localhost:5432/ordering?sslmode=disable&search_path=ordering,public")
	}
	if err != nil {
		return
	}
	c.client = orderingclient.New(cfg.transport, strfmt.Default)

	return
}

func (c *ordersFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^the checked out basket eventually becomes an order$`, c.theCheckedOutBasketEventuallyBecomesAnOrder)
	ctx.Step(`^(?:I )?(?:ensure |expect )?an order exists for the checked out basket$`, c.expectAnOrderExistsForTheCheckedOutBasket)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order status is "([^"]*)"$`, c.expectTheOrderStatusIs)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order belongs to the current customer$`, c.expectTheOrderBelongsToTheCurrentCustomer)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order references the authorized payment$`, c.expectTheOrderReferencesTheAuthorizedPayment)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order contains the basket items$`, c.expectTheOrderContainsTheBasketItems)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order is visible through the ordering API$`, c.expectTheOrderIsVisibleThroughTheOrderingAPI)
	}

func (c *ordersFeature) reset() {
	truncate := func(tableName string) {
		_, _ = c.db.Exec(fmt.Sprintf("TRUNCATE %s", tableName))
	}

	truncate("ordering.orders")
	truncate("ordering.events")
	truncate("ordering.snapshots")
	truncate("ordering.inbox")
	truncate("ordering.outbox")
	truncate("ordering.sagas")
	}

func (c *ordersFeature) theCheckedOutBasketEventuallyBecomesAnOrder(ctx context.Context) error {
	_, err := c.waitForOrder(ctx)
	return err
}

func (c *ordersFeature) expectAnOrderExistsForTheCheckedOutBasket(ctx context.Context) error {
	_, err := c.waitForOrder(ctx)
	return err
}

func (c *ordersFeature) expectTheOrderStatusIs(ctx context.Context, status string) error {
	ord, err := c.waitForOrderStatus(ctx, status)
	if err != nil {
		return err
	}

	if ord.Status != status {
		return errors.ErrBadRequest.Msgf("expected order status `%s`, got `%s`", status, ord.Status)
	}

	return nil
}

func (c *ordersFeature) expectTheOrderBelongsToTheCurrentCustomer(ctx context.Context) error {
	ord, err := c.waitForOrder(ctx)
	if err != nil {
		return err
	}
	customerID, err := lastCustomerID(ctx)
	if err != nil {
		return err
	}

	if ord.CustomerID != customerID {
		return errors.ErrBadRequest.Msgf("expected order customer `%s`, got `%s`", customerID, ord.CustomerID)
	}

	return nil
}

func (c *ordersFeature) expectTheOrderReferencesTheAuthorizedPayment(ctx context.Context) error {
	ord, err := c.waitForOrder(ctx)
	if err != nil {
		return err
	}
	paymentID, err := lastPaymentID(ctx)
	if err != nil {
		return err
	}

	if ord.PaymentID != paymentID {
		return errors.ErrBadRequest.Msgf("expected order payment `%s`, got `%s`", paymentID, ord.PaymentID)
	}

	return nil
}

func (c *ordersFeature) expectTheOrderContainsTheBasketItems(ctx context.Context) error {
	ord, err := c.waitForOrder(ctx)
	if err != nil {
		return err
	}
	basketItems, err := currentBasketItems(ctx)
	if err != nil {
		return err
	}

	if len(ord.Items) != len(basketItems) {
		return errors.ErrBadRequest.Msgf("expected `%d` order items, got `%d`", len(basketItems), len(ord.Items))
	}

	for i, basketItem := range basketItems {
		orderItem := ord.Items[i]
		if orderItem.ProductID != basketItem.ProductID || orderItem.StoreID != basketItem.StoreID || orderItem.Quantity != basketItem.Quantity {
			return errors.ErrBadRequest.Msgf("expected order item `%d` to match basket item", i)
		}
	}

	return nil
}

func (c *ordersFeature) expectTheOrderIsVisibleThroughTheOrderingAPI(ctx context.Context) error {
	_, err := c.waitForOrder(ctx)
	return err
}

func (c *ordersFeature) waitForOrder(ctx context.Context) (*models.OrderingpbOrder, error) {
	orderID, err := lastBasketID(ctx)
	if err != nil {
		return nil, err
	}

	deadline := time.Now().Add(5 * time.Second)
	var lastErr error
	for time.Now().Before(deadline) {
		resp, err := c.client.Order.GetOrder(order.NewGetOrderParams().WithID(orderID))
		ctx = setLastResponseAndError(ctx, resp, err)
		if err == nil && resp != nil && resp.Payload != nil && resp.Payload.Order != nil {
			return resp.Payload.Order, nil
		}
		lastErr = err
		time.Sleep(200 * time.Millisecond)
	}

	if lastErr != nil {
		return nil, lastErr
	}
	return nil, errors.ErrUnknown.Msgf("order `%s` was not created before timeout", orderID)
	}

func (c *ordersFeature) waitForOrderStatus(ctx context.Context, expected string) (*models.OrderingpbOrder, error) {
	deadline := time.Now().Add(5 * time.Second)
	var lastOrder *models.OrderingpbOrder
	for time.Now().Before(deadline) {
		ord, err := c.waitForOrder(ctx)
		if err == nil {
			lastOrder = ord
			if ord.Status == expected {
				return ord, nil
			}
		}
		time.Sleep(200 * time.Millisecond)
	}

	if lastOrder != nil {
		return nil, errors.ErrBadRequest.Msgf("expected order status `%s`, got `%s`", expected, lastOrder.Status)
	}

	return nil, errors.ErrUnknown.Msgf("order did not reach status `%s` before timeout", expected)
	}
