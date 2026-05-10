//go:build e2e

package e2e

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/cucumber/godog"
	"github.com/go-openapi/strfmt"
	"github.com/stackus/errors"

	"github.com/owezzy/soko-bora-mngt-system/ordering/orderingclient"
	orderingmodels "github.com/owezzy/soko-bora-mngt-system/ordering/orderingclient/models"
	"github.com/owezzy/soko-bora-mngt-system/ordering/orderingclient/order"
	"github.com/owezzy/soko-bora-mngt-system/search/searchpb"
	"github.com/owezzy/soko-bora-mngt-system/search/searchpb/searchpbconnect"
)

type shoppingListIDKey struct{}

const orderPropagationTimeout = 30 * time.Second

type ordersFeature struct {
	client       *orderingclient.OrderProcessing
	searchClient searchpbconnect.SearchServiceClient
	httpClient   *http.Client
	orderingDB   *sql.DB
	depotDB      *sql.DB
	searchDB     *sql.DB
}

func (c *ordersFeature) init(cfg featureConfig) (err error) {
	if cfg.useMonoDB {
		c.orderingDB, err = sql.Open("pgx", "postgres://mallbots_user:mallbots_pass@localhost:5432/mallbots?sslmode=disable")
		if err != nil {
			return err
		}
		c.depotDB, err = sql.Open("pgx", "postgres://mallbots_user:mallbots_pass@localhost:5432/mallbots?sslmode=disable")
		if err != nil {
			return err
		}
		c.searchDB, err = sql.Open("pgx", "postgres://mallbots_user:mallbots_pass@localhost:5432/mallbots?sslmode=disable")
		if err != nil {
			return err
		}
	} else {
		c.orderingDB, err = sql.Open("pgx", "postgres://ordering_user:ordering_pass@localhost:5432/ordering?sslmode=disable&search_path=ordering,public")
		if err != nil {
			return err
		}
		c.depotDB, err = sql.Open("pgx", "postgres://depot_user:depot_pass@localhost:5432/depot?sslmode=disable&search_path=depot,public")
		if err != nil {
			return err
		}
		c.searchDB, err = sql.Open("pgx", "postgres://search_user:search_pass@localhost:5432/search?sslmode=disable&search_path=search,public")
		if err != nil {
			return err
		}
	}

	c.client = orderingclient.New(cfg.transport, strfmt.Default)
	c.httpClient = &http.Client{Timeout: 5 * time.Second}
	c.searchClient = searchpbconnect.NewSearchServiceClient(c.httpClient, "http://localhost:8080")

	return nil
}

func (c *ordersFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^the checked out basket eventually becomes an order$`, c.theCheckedOutBasketEventuallyBecomesAnOrder)
	ctx.Step(`^(?:I )?(?:ensure |expect )?an order exists for the checked out basket$`, c.expectAnOrderExistsForTheCheckedOutBasket)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order status is "([^"]*)"$`, c.expectTheOrderStatusIs)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order belongs to the current customer$`, c.expectTheOrderBelongsToTheCurrentCustomer)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order references the authorized payment$`, c.expectTheOrderReferencesTheAuthorizedPayment)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order contains the basket items$`, c.expectTheOrderContainsTheBasketItems)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order is visible through the ordering API$`, c.expectTheOrderIsVisibleThroughTheOrderingAPI)
	ctx.Step(`^I assign the shopping list to bot "([^"]*)"$`, c.iAssignTheShoppingListToBot)
	ctx.Step(`^I complete the shopping list$`, c.iCompleteTheShoppingList)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the shopping list status is "([^"]*)"$`, c.expectTheShoppingListStatusIs)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the search projection status is "([^"]*)"$`, c.expectTheSearchProjectionStatusIs)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the depot shopping list exists for the order$`, c.expectTheDepotShoppingListExistsForTheOrder)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the order has an invoice id$`, c.expectTheOrderHasAnInvoiceID)
	}

func (c *ordersFeature) reset() {
	truncate := func(tableName string) {
		_, _ = c.orderingDB.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", tableName))
	}

	truncate("ordering.orders")
	truncate("ordering.events")
	truncate("ordering.snapshots")
	truncate("ordering.inbox")
	truncate("ordering.outbox")
	truncate("ordering.sagas")
	truncate("depot.shopping_lists")
	truncate("depot.inbox")
	truncate("depot.outbox")
	truncate("search.orders")
	truncate("search.inbox")
	truncate("search.outbox")
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

	orderItems := make(map[string]int32, len(ord.Items))
	for _, orderItem := range ord.Items {
		orderItems[fmt.Sprintf("%s|%s", orderItem.ProductID, orderItem.StoreID)] = orderItem.Quantity
	}

	for _, basketItem := range basketItems {
		key := fmt.Sprintf("%s|%s", basketItem.ProductID, basketItem.StoreID)
		quantity, ok := orderItems[key]
		if !ok {
			return errors.ErrBadRequest.Msgf("expected order item `%s` to exist", key)
		}
		if quantity != basketItem.Quantity {
			return errors.ErrBadRequest.Msgf("expected order item `%s` quantity `%d`, got `%d`", key, basketItem.Quantity, quantity)
		}
	}

	return nil
}

func (c *ordersFeature) expectTheOrderIsVisibleThroughTheOrderingAPI(ctx context.Context) error {
	_, err := c.waitForOrder(ctx)
	return err
}

func (c *ordersFeature) iAssignTheShoppingListToBot(ctx context.Context, botID string) error {
	shoppingListID, err := c.waitForShoppingList(ctx)
	if err != nil {
		return err
	}

	body, err := json.Marshal(map[string]string{"botId": botID})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("http://localhost:8080/api/depot/shopping/%s/assign", shoppingListID), bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode/100 != 2 {
		return errors.ErrUnknown.Msgf("assign shopping list failed with status %s", resp.Status)
	}

	return nil
}

func (c *ordersFeature) iCompleteTheShoppingList(ctx context.Context) error {
	shoppingListID, err := c.waitForShoppingList(ctx)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("http://localhost:8080/api/depot/shopping/%s/complete", shoppingListID), nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode/100 != 2 {
		return errors.ErrUnknown.Msgf("complete shopping list failed with status %s", resp.Status)
	}

	return nil
}

func (c *ordersFeature) expectTheShoppingListStatusIs(ctx context.Context, expected string) error {
	orderID, err := lastBasketID(ctx)
	if err != nil {
		return err
	}

	deadline := time.Now().Add(orderPropagationTimeout)
	for time.Now().Before(deadline) {
		var status string
		row := c.depotDB.QueryRow("SELECT status FROM depot.shopping_lists WHERE order_id = $1", orderID)
		if err = row.Scan(&status); err == nil && status == expected {
			return nil
		}
		time.Sleep(200 * time.Millisecond)
	}

	return errors.ErrUnknown.Msgf("shopping list for order `%s` did not reach status `%s` before timeout", orderID, expected)
	}

func (c *ordersFeature) expectTheSearchProjectionStatusIs(ctx context.Context, expected string) error {
	orderID, err := lastBasketID(ctx)
	if err != nil {
		return err
	}

	deadline := time.Now().Add(orderPropagationTimeout)
	for time.Now().Before(deadline) {
		resp, callErr := c.searchClient.GetOrder(ctx, connect.NewRequest(&searchpb.GetOrderRequest{Id: orderID}))
		if callErr == nil && resp.Msg != nil && resp.Msg.Order != nil && resp.Msg.Order.Status == expected {
			return nil
		}
		time.Sleep(200 * time.Millisecond)
	}

	return errors.ErrUnknown.Msgf("search projection for order `%s` did not reach status `%s` before timeout", orderID, expected)
	}

func (c *ordersFeature) expectTheDepotShoppingListExistsForTheOrder(ctx context.Context) error {
	_, err := c.waitForShoppingList(ctx)
	return err
}

func (c *ordersFeature) expectTheOrderHasAnInvoiceID(ctx context.Context) error {
	orderID, err := lastBasketID(ctx)
	if err != nil {
		return err
	}

	deadline := time.Now().Add(orderPropagationTimeout)
	for time.Now().Before(deadline) {
		var invoiceID string
		row := c.orderingDB.QueryRow("SELECT id FROM payments.invoices WHERE order_id = $1", orderID)
		if err = row.Scan(&invoiceID); err == nil && invoiceID != "" {
			return nil
		}
		time.Sleep(200 * time.Millisecond)
	}

	return errors.ErrNotFound.Msg("no invoice id recorded for order")
}

func (c *ordersFeature) waitForOrder(ctx context.Context) (*orderingmodels.OrderingpbOrder, error) {
	orderID, err := lastBasketID(ctx)
	if err != nil {
		return nil, err
	}

	deadline := time.Now().Add(orderPropagationTimeout)
	var lastErr error
	for time.Now().Before(deadline) {
		resp, err := c.client.Order.GetOrder(order.NewGetOrderParams().WithID(orderID))
		ctx = setLastResponseAndError(ctx, resp, err)
		if err == nil && resp != nil && resp.Payload != nil && resp.Payload.Order != nil {
			ord := resp.Payload.Order
			if ord.ID == orderID && ord.CustomerID != "" && ord.PaymentID != "" && len(ord.Items) > 0 && ord.Status != "" {
				return ord, nil
			}
		}
		lastErr = err
		time.Sleep(200 * time.Millisecond)
	}

	if lastErr != nil {
		return nil, lastErr
	}
	return nil, errors.ErrUnknown.Msgf("order `%s` was not created before timeout", orderID)
}

func (c *ordersFeature) waitForOrderStatus(ctx context.Context, expected string) (*orderingmodels.OrderingpbOrder, error) {
	deadline := time.Now().Add(orderPropagationTimeout)
	var lastOrder *orderingmodels.OrderingpbOrder
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

func (c *ordersFeature) waitForShoppingList(ctx context.Context) (string, error) {
	orderID, err := lastBasketID(ctx)
	if err != nil {
		return "", err
	}

	deadline := time.Now().Add(orderPropagationTimeout)
	for time.Now().Before(deadline) {
		var shoppingListID string
		row := c.depotDB.QueryRow("SELECT id FROM depot.shopping_lists WHERE order_id = $1", orderID)
		if err = row.Scan(&shoppingListID); err == nil {
			return shoppingListID, nil
		}
		time.Sleep(200 * time.Millisecond)
	}

	return "", errors.ErrUnknown.Msgf("shopping list for order `%s` was not created before timeout", orderID)
}
