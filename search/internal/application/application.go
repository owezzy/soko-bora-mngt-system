package application

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/stackus/errors"

	"github.com/owezzy/soko-bora-mngt-system/internal/demo"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/models"
)

const DefaultSearchLimit = 25

type (
	Filters struct {
		CustomerID string
		After      time.Time
		Before     time.Time
		StoreIDs   []string
		ProductIDs []string
		MinTotal   float64
		MaxTotal   float64
		Status     string
	}
	SearchOrders struct {
		Filters Filters
		Next    string
		Limit   int
	}

	GetOrder struct {
		OrderID string
	}

	DemoBootstrap = demo.BootstrapSpec

	Application interface {
		SearchOrders(ctx context.Context, search SearchOrders) ([]*models.Order, error)
		GetOrder(ctx context.Context, get GetOrder) (*models.Order, error)
		GetDemoBootstrap(ctx context.Context) (DemoBootstrap, error)
	}

	app struct {
		orders OrderRepository
	}
)

var _ Application = (*app)(nil)

func New(orders OrderRepository) *app {
	return &app{
		orders: orders,
	}
}

func (a app) SearchOrders(ctx context.Context, search SearchOrders) ([]*models.Order, error) {
	if search.Limit < 0 {
		return nil, errors.ErrInvalidArgument.Msg("search limit cannot be negative")
	}
	if search.Limit == 0 {
		search.Limit = DefaultSearchLimit
	}
	if search.Next != "" {
		next, err := strconv.Atoi(search.Next)
		if err != nil || next < 0 {
			return nil, errors.ErrInvalidArgument.Msg("search next cursor must be a non-negative offset")
		}
	}

	search.Filters.CustomerID = strings.TrimSpace(search.Filters.CustomerID)
	search.Filters.Status = strings.TrimSpace(search.Filters.Status)

	orders, err := a.orders.Search(ctx, search)
	if err != nil {
		return nil, errors.Wrap(err, "search orders query")
	}

	return orders, nil
}

func (a app) GetOrder(ctx context.Context, get GetOrder) (*models.Order, error) {
	get.OrderID = strings.TrimSpace(get.OrderID)
	if get.OrderID == "" {
		return nil, errors.ErrInvalidArgument.Msg("order id is required")
	}

	order, err := a.orders.Get(ctx, get.OrderID)
	if err != nil {
		return nil, errors.Wrap(err, "get order query")
	}

	return order, nil
}

func (a app) GetDemoBootstrap(ctx context.Context) (DemoBootstrap, error) {
	return demo.Spec(), nil
}
