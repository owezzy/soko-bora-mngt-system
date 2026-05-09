package grpc

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"

	"github.com/owezzy/soko-bora-mngt-system/customers/customerspb"
	"github.com/owezzy/soko-bora-mngt-system/stores/storespb"
)

func TestCustomerRepositoryFindReturnsNotFoundOnNilCustomer(t *testing.T) {
	t.Parallel()

	endpoint := startCustomersServer(t, customerspb.CustomersServiceServer(&customerTestServer{
		getCustomer: func(context.Context, *customerspb.GetCustomerRequest) (*customerspb.GetCustomerResponse, error) {
			return &customerspb.GetCustomerResponse{}, nil
		},
	}))

	repo := NewCustomerRepository(endpoint)

	customer, err := repo.Find(context.Background(), "missing-customer")
	if err == nil {
		t.Fatal("expected not found error")
	}
	if customer != nil {
		t.Fatalf("expected nil customer, got %#v", customer)
	}
}

func TestStoreRepositoryFindReturnsNotFoundOnNilStore(t *testing.T) {
	t.Parallel()

	endpoint := startStoresServer(t, storespb.StoresServiceServer(&storesTestServer{
		getStore: func(context.Context, *storespb.GetStoreRequest) (*storespb.GetStoreResponse, error) {
			return &storespb.GetStoreResponse{}, nil
		},
	}))

	repo := NewStoreRepository(endpoint)

	store, err := repo.Find(context.Background(), "missing-store")
	if err == nil {
		t.Fatal("expected not found error")
	}
	if store != nil {
		t.Fatalf("expected nil store, got %#v", store)
	}
}

func TestProductRepositoryFindReturnsNotFoundOnNilProduct(t *testing.T) {
	t.Parallel()

	endpoint := startStoresServer(t, storespb.StoresServiceServer(&storesTestServer{
		getProduct: func(context.Context, *storespb.GetProductRequest) (*storespb.GetProductResponse, error) {
			return &storespb.GetProductResponse{}, nil
		},
	}))

	repo := NewProductRepository(endpoint)

	product, err := repo.Find(context.Background(), "missing-product")
	if err == nil {
		t.Fatal("expected not found error")
	}
	if product != nil {
		t.Fatalf("expected nil product, got %#v", product)
	}
}

func TestProductRepositoryFindPreservesStoreID(t *testing.T) {
	t.Parallel()

	endpoint := startStoresServer(t, storespb.StoresServiceServer(&storesTestServer{
		getProduct: func(context.Context, *storespb.GetProductRequest) (*storespb.GetProductResponse, error) {
			return &storespb.GetProductResponse{Product: &storespb.Product{
				Id:      "product-1",
				StoreId: "store-1",
				Name:    "Bananas",
			}}, nil
		},
	}))

	repo := NewProductRepository(endpoint)

	product, err := repo.Find(context.Background(), "product-1")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if product == nil {
		t.Fatal("expected product")
	}
	if product.StoreID != "store-1" {
		t.Fatalf("expected store id to be preserved, got %q", product.StoreID)
	}
}

type customerTestServer struct {
	customerspb.UnimplementedCustomersServiceServer
	getCustomer func(context.Context, *customerspb.GetCustomerRequest) (*customerspb.GetCustomerResponse, error)
}

func (s *customerTestServer) GetCustomer(ctx context.Context, req *customerspb.GetCustomerRequest) (*customerspb.GetCustomerResponse, error) {
	return s.getCustomer(ctx, req)
}

type storesTestServer struct {
	storespb.UnimplementedStoresServiceServer
	getStore   func(context.Context, *storespb.GetStoreRequest) (*storespb.GetStoreResponse, error)
	getProduct func(context.Context, *storespb.GetProductRequest) (*storespb.GetProductResponse, error)
}

func (s *storesTestServer) GetStore(ctx context.Context, req *storespb.GetStoreRequest) (*storespb.GetStoreResponse, error) {
	if s.getStore == nil {
		return &storespb.GetStoreResponse{}, nil
	}
	return s.getStore(ctx, req)
}

func (s *storesTestServer) GetProduct(ctx context.Context, req *storespb.GetProductRequest) (*storespb.GetProductResponse, error) {
	if s.getProduct == nil {
		return &storespb.GetProductResponse{}, nil
	}
	return s.getProduct(ctx, req)
}

func startCustomersServer(t *testing.T, server customerspb.CustomersServiceServer) string {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	customerspb.RegisterCustomersServiceServer(grpcServer, server)

	go func() {
		_ = grpcServer.Serve(listener)
	}()

	t.Cleanup(func() {
		grpcServer.Stop()
		_ = listener.Close()
	})

	return listener.Addr().String()
}

func startStoresServer(t *testing.T, server storespb.StoresServiceServer) string {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	storespb.RegisterStoresServiceServer(grpcServer, server)

	go func() {
		_ = grpcServer.Serve(listener)
	}()

	t.Cleanup(func() {
		grpcServer.Stop()
		_ = listener.Close()
	})

	return listener.Addr().String()
}
