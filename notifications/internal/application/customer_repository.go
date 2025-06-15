package application

import (
	"context"

	"github.com/owezzy/soko-bora-mngt-system/notifications/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
