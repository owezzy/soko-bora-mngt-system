package application

import (
	"context"

	"github.com/owezzy/soko-bora-mngt-system/search/internal/models"
)

type StoreRepository interface {
	Find(ctx context.Context, storeID string) (*models.Store, error)
}

type StoreCacheRepository interface {
	Add(ctx context.Context, storeID, name string) error
	Rename(ctx context.Context, storeID, name string) error
	StoreRepository
}
