package queries

import (
	"context"

	"github.com/owezzy/soko-bora-mngt-system/stores/internal/domain"
)

type GetStores struct{}

type GetStoresHandler struct {
	mall domain.MallRepository
}

func NewGetStoresHandler(mall domain.MallRepository) GetStoresHandler {
	return GetStoresHandler{mall: mall}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*domain.MallStore, error) {
	return h.mall.All(ctx)
}
