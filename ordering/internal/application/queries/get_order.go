package queries

import (
	"context"

	"github.com/stackus/errors"

	"github.com/owezzy/soko-bora-mngt-system/ordering/internal/domain"
)

type GetOrder struct {
	ID string
}

type GetOrderHandler struct {
	repo domain.OrderRepository
}

func NewGetOrderHandler(repo domain.OrderRepository) GetOrderHandler {
	return GetOrderHandler{repo: repo}
}

func (h GetOrderHandler) GetOrder(ctx context.Context, query GetOrder) (*domain.Order, error) {
	order, err := h.repo.Load(ctx, query.ID)
	if err != nil {
		return nil, errors.Wrap(err, "get order query")
	}

	if order.Version() == 0 {
		return nil, errors.Wrapf(errors.ErrNotFound, "order `%s` does not exist", query.ID)
	}

	return order, nil
}
