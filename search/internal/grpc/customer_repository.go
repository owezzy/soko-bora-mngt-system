package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/owezzy/soko-bora-mngt-system/customers/customerspb"
	"github.com/owezzy/soko-bora-mngt-system/internal/rpc"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/models"
)

type CustomerRepository struct {
	endpoint string
}

var _ application.CustomerRepository = (*CustomerRepository)(nil)

func NewCustomerRepository(endpoint string) CustomerRepository {
	return CustomerRepository{
		endpoint: endpoint,
	}
}

func (r CustomerRepository) Find(ctx context.Context, customerID string) (customer *models.Customer, err error) {
	var conn *grpc.ClientConn
	conn, err = r.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := customerspb.NewCustomersServiceClient(conn).GetCustomer(ctx, &customerspb.GetCustomerRequest{Id: customerID})
	if err != nil {
		return nil, err
	}

	return &models.Customer{
		ID:   resp.GetCustomer().GetId(),
		Name: resp.GetCustomer().GetName(),
	}, nil
}

func (r CustomerRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, r.endpoint)
}
