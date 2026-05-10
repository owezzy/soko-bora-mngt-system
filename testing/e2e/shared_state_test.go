//go:build e2e

package e2e

import (
	"context"
	"math"

	basketmodels "github.com/owezzy/soko-bora-mngt-system/baskets/basketsclient/models"
	"github.com/stackus/errors"
)

type basketItemsKey struct{}
type basketTotalKey struct{}
type seededDemoScenarioKey struct{}

type basketItemSnapshot struct {
	ProductID string
	StoreID   string
	Quantity  int32
}

func setSeededDemoScenario(ctx context.Context, seeded bool) context.Context {
	return context.WithValue(ctx, seededDemoScenarioKey{}, seeded)
}

func isSeededDemoScenario(ctx context.Context) bool {
	v := ctx.Value(seededDemoScenarioKey{})
	if v == nil {
		return false
	}

	seeded, ok := v.(bool)
	return ok && seeded
}

func setCurrentBasketItems(ctx context.Context, items []*basketmodels.BasketspbItem) context.Context {
	cloned := make([]basketItemSnapshot, len(items))
	var total float64
	for i, item := range items {
		cloned[i] = basketItemSnapshot{
			ProductID: item.ProductID,
			StoreID:   item.StoreID,
			Quantity:  item.Quantity,
		}
		total += item.ProductPrice * float64(item.Quantity)
	}

	ctx = context.WithValue(ctx, basketItemsKey{}, cloned)
	return context.WithValue(ctx, basketTotalKey{}, total)
}

func currentBasketItems(ctx context.Context) ([]basketItemSnapshot, error) {
	v := ctx.Value(basketItemsKey{})
	if v == nil {
		return nil, errors.ErrNotFound.Msg("no basket items are available")
	}

	items, ok := v.([]basketItemSnapshot)
	if !ok {
		return nil, errors.ErrInternal.Msg("basket items are in an unexpected format")
	}

	return items, nil
}

func currentBasketTotal(ctx context.Context) (float64, error) {
	v := ctx.Value(basketTotalKey{})
	if v == nil {
		return 0, errors.ErrNotFound.Msg("no basket snapshot is available to calculate total")
	}

	total, ok := v.(float64)
	if !ok {
		return 0, errors.ErrInternal.Msg("basket total is in an unexpected format")
	}

	return total, nil
}

func nearlyEqualFloat64(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}
