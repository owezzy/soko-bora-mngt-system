//go:build e2e

package e2e

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cucumber/messages-go/v16"
	"github.com/cucumber/godog"
	"github.com/go-openapi/strfmt"
	"github.com/stackus/errors"

	"github.com/owezzy/soko-bora-mngt-system/baskets/basketsclient"
	"github.com/owezzy/soko-bora-mngt-system/baskets/basketsclient/basket"
	"github.com/owezzy/soko-bora-mngt-system/baskets/basketsclient/item"
	"github.com/owezzy/soko-bora-mngt-system/baskets/basketsclient/models"
	"github.com/owezzy/soko-bora-mngt-system/internal/demo"
)

type basketIDKey struct{}

type basketsFeature struct {
	client *basketsclient.ShoppingBaskets
	db     *sql.DB
}

func (c *basketsFeature) init(cfg featureConfig) (err error) {
	if cfg.useMonoDB {
		c.db, err = sql.Open("pgx", "postgres://mallbots_user:mallbots_pass@localhost:5432/mallbots?sslmode=disable")
	} else {
		c.db, err = sql.Open("pgx", "postgres://baskets_user:baskets_pass@localhost:5432/baskets?sslmode=disable&search_path=baskets,public")
	}
	if err != nil {
		return
	}
	c.client = basketsclient.New(cfg.transport, strfmt.Default)

	return
}

func (c *basketsFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I start a new basket$`, c.iStartANewBasket)
	ctx.Step(`^I start a new basket for the seeded demo customer$`, c.iStartANewBasketForTheSeededDemoCustomer)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the basket (?:was|is) started$`, c.expectTheBasketWasStarted)

	ctx.Step(`^I add the items$`, c.iAddTheItems)
	ctx.Step(`^I add the seeded demo items$`, c.iAddTheSeededDemoItems)
	ctx.Step(`^I check out the basket$`, c.iCheckOutTheBasket)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the basket (?:was|is) checked out$`, c.expectTheBasketWasCheckedOut)
	ctx.Step(`^I fetch the basket snapshot$`, c.iFetchTheBasketSnapshot)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the items (?:were|are) added$`, c.expectTheItemsWereAdded)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the current basket has (\d+) item(?:s)?$`, c.expectTheCurrentBasketHasItems)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the current basket total is "([^"]*)"$`, c.expectTheCurrentBasketTotalIs)
}

func (c *basketsFeature) reset() {
	truncate := func(tableName string) {
		_, _ = c.db.Exec(fmt.Sprintf("TRUNCATE %s", tableName))
	}

	truncate("baskets.events")
	truncate("baskets.snapshots")
	truncate("baskets.inbox")
	truncate("baskets.outbox")
	truncate("baskets.products_cache")
	truncate("baskets.stores_cache")
}

func (c *basketsFeature) iStartANewBasket(ctx context.Context) (context.Context, error) {
	customerID, err := lastCustomerID(ctx)
	if err != nil {
		return ctx, err
	}
	resp, err := c.client.Basket.StartBasket(basket.NewStartBasketParams().WithBody(&models.BasketspbStartBasketRequest{
		CustomerID: customerID,
	}))

	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx, nil
	}
	return context.WithValue(ctx, basketIDKey{}, resp.Payload.ID), nil
}

func (c *basketsFeature) iStartANewBasketForTheSeededDemoCustomer(ctx context.Context) (context.Context, error) {
	ctx = context.WithValue(ctx, customerIDKey{}, demo.Spec().Customer.ID)
	return c.iStartANewBasket(ctx)
}

func (c *basketsFeature) expectTheBasketWasStarted(ctx context.Context) error {
	if err := lastResponseWas(ctx, &basket.StartBasketOK{}); err != nil {
		return err
	}

	return nil
}

func (c *basketsFeature) iAddTheItems(ctx context.Context, table *godog.Table) (context.Context, error) {
	type Item struct {
		Name     string
		Quantity int
	}

	basketID, err := lastBasketID(ctx)
	if err != nil {
		return ctx, err
	}

	items, err := assist.CreateSlice(new(Item), table)
	if err != nil {
		return ctx, err
	}
	for _, i := range items.([]*Item) {
		productID := getProductID(ctx, i.Name)
		resp, err := c.client.Item.AddItem(item.NewAddItemParams().WithID(basketID).WithBody(&models.BasketServiceAddItemBody{
			ProductID: productID,
			Quantity:  int32(i.Quantity),
		}))
		ctx = setLastResponseAndError(ctx, resp, err)
		if err != nil {
			break
		}
	}

	return ctx, nil
}

func (c *basketsFeature) iAddTheSeededDemoItems(ctx context.Context) (context.Context, error) {
	ctx = addProduct(ctx, demo.Spec().Stores[0].Products[0].ID, demo.Spec().Stores[0].Products[0].Name)
	ctx = addProduct(ctx, demo.Spec().Stores[1].Products[0].ID, demo.Spec().Stores[1].Products[0].Name)

	table := &godog.Table{Rows: []*messages.PickleTableRow{
		{Cells: []*messages.PickleTableCell{{Value: "Name"}, {Value: "Quantity"}}},
		{Cells: []*messages.PickleTableCell{{Value: demo.Spec().Stores[0].Products[0].Name}, {Value: "2"}}},
		{Cells: []*messages.PickleTableCell{{Value: demo.Spec().Stores[1].Products[0].Name}, {Value: "1"}}},
	}}

	return c.iAddTheItems(ctx, table)
}

func (c *basketsFeature) expectTheItemsWereAdded(ctx context.Context) error {
	if err := lastResponseWas(ctx, &item.AddItemOK{}); err != nil {
		return err
	}

	return nil
}

func (c *basketsFeature) iCheckOutTheBasket(ctx context.Context) (context.Context, error) {
	basketID, err := lastBasketID(ctx)
	if err != nil {
		return ctx, err
	}
	paymentID, err := lastPaymentID(ctx)
	if err != nil {
		return ctx, err
	}

	resp, err := c.client.Basket.CheckoutBasket(basket.NewCheckoutBasketParams().WithID(basketID).WithBody(&models.BasketServiceCheckoutBasketBody{
		PaymentID: paymentID,
	}))
	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx, nil
	}

	return ctx, nil
}

func (c *basketsFeature) expectTheBasketWasCheckedOut(ctx context.Context) error {
	if err := lastResponseWas(ctx, &basket.CheckoutBasketOK{}); err != nil {
		return err
	}

	return nil
}

func (c *basketsFeature) iFetchTheBasketSnapshot(ctx context.Context) (context.Context, error) {
	basketID, err := lastBasketID(ctx)
	if err != nil {
		return ctx, err
	}

	resp, err := c.client.Basket.GetBasket(basket.NewGetBasketParams().WithID(basketID))
	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx, nil
	}

	if resp.Payload == nil || resp.Payload.Basket == nil {
		return ctx, errors.ErrNotFound.Msg("basket snapshot was empty")
	}

	return setCurrentBasketItems(ctx, resp.Payload.Basket.Items), nil
}

func (c *basketsFeature) expectTheCurrentBasketHasItems(ctx context.Context, expected int) error {
	items, err := currentBasketItems(ctx)
	if err != nil {
		return err
	}

	if len(items) != expected {
		return errors.ErrBadRequest.Msgf("expected `%d` basket items, got `%d`", expected, len(items))
	}

	return nil
}

func (c *basketsFeature) expectTheCurrentBasketTotalIs(ctx context.Context, expected float64) error {
	total, err := currentBasketTotal(ctx)
	if err != nil {
		return err
	}

	if !nearlyEqualFloat64(total, expected) {
		return errors.ErrBadRequest.Msgf("expected basket total `%0.2f`, got `%0.2f`", expected, total)
	}

	return nil
}

func lastBasketID(ctx context.Context) (string, error) {
	v := ctx.Value(basketIDKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no basket ID to work with")
	}
	return v.(string), nil
}
