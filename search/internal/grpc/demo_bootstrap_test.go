package grpc

import (
	"context"
	"testing"

	"github.com/owezzy/soko-bora-mngt-system/internal/demo"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/models"
	"github.com/owezzy/soko-bora-mngt-system/search/searchpb"
)

func TestGetDemoBootstrapReturnsDeterministicSeedContext(t *testing.T) {
	t.Parallel()

	svc := server{app: stubApplication{bootstrap: demo.Spec()}}

	resp, err := svc.GetDemoBootstrap(context.Background(), &searchpb.GetDemoBootstrapRequest{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if resp.GetCustomer().GetId() != demo.Spec().Customer.ID {
		t.Fatalf("expected demo customer id %q, got %q", demo.Spec().Customer.ID, resp.GetCustomer().GetId())
	}
	if len(resp.GetStores()) != len(demo.Spec().Stores) {
		t.Fatalf("expected %d stores, got %d", len(demo.Spec().Stores), len(resp.GetStores()))
	}

	productCount := 0
	for _, store := range demo.Spec().Stores {
		productCount += len(store.Products)
	}
	if len(resp.GetProducts()) != productCount {
		t.Fatalf("expected %d demo products, got %d", productCount, len(resp.GetProducts()))
	}

	firstStore := demo.Spec().Stores[0]
	firstProduct := firstStore.Products[0]

	if resp.GetStores()[0].GetId() != firstStore.ID || resp.GetStores()[0].GetName() != firstStore.Name {
		t.Fatal("expected first demo store to match bootstrap spec")
	}
	if resp.GetProducts()[0].GetId() != firstProduct.ID || resp.GetProducts()[0].GetStoreId() != firstStore.ID || resp.GetProducts()[0].GetStoreName() != firstStore.Name {
		t.Fatal("expected first demo product to carry store context from bootstrap spec")
	}
}

type stubApplication struct {
	bootstrap application.DemoBootstrap
}

func (s stubApplication) SearchOrders(context.Context, application.SearchOrders) ([]*models.Order, error) {
	panic("unused")
}

func (s stubApplication) GetOrder(context.Context, application.GetOrder) (*models.Order, error) {
	panic("unused")
}

func (s stubApplication) GetDemoBootstrap(context.Context) (application.DemoBootstrap, error) {
	return s.bootstrap, nil
}
