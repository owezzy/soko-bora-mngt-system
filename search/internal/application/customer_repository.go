package application

import (
	"context"

	"github.com/owezzy/soko-bora-mngt-system/search/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}

type CustomerCacheRepository interface {
	Add(ctx context.Context, customerID, name string) error
	CustomerRepository
}
