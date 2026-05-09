package grpc

import (
	"context"

	"github.com/stackus/errors"
	"google.golang.org/grpc"

	"github.com/owezzy/soko-bora-mngt-system/internal/rpc"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/models"
	"github.com/owezzy/soko-bora-mngt-system/stores/storespb"
)

type ProductRepository struct {
	endpoint string
}

var _ application.ProductRepository = (*ProductRepository)(nil)

func NewProductRepository(endpoint string) ProductRepository {
	return ProductRepository{
		endpoint: endpoint,
	}
}

func (r ProductRepository) Find(ctx context.Context, productID string) (product *models.Product, err error) {
	var conn *grpc.ClientConn
	conn, err = r.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := storespb.NewStoresServiceClient(conn).GetProduct(ctx, &storespb.GetProductRequest{Id: productID})
	if err != nil {
		return nil, err
	}
	if resp.GetProduct() == nil {
		return nil, errors.ErrNotFound.Msgf("product with id: `%s` does not exist", productID)
	}

	return r.productToDomain(resp.GetProduct()), nil
}

func (r ProductRepository) productToDomain(product *storespb.Product) *models.Product {
	return &models.Product{
		ID:      product.GetId(),
		StoreID: product.GetStoreId(),
		Name:    product.GetName(),
	}
}

func (r ProductRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, r.endpoint)
}
