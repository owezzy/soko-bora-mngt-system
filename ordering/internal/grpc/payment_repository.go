package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/owezzy/soko-bora-mngt-system/ordering/internal/domain"
	"github.com/owezzy/soko-bora-mngt-system/payments/paymentspb"
)

type PaymentRepository struct {
	client paymentspb.PaymentsServiceClient
}

var _ domain.PaymentRepository = (*PaymentRepository)(nil)

func NewPaymentRepository(conn *grpc.ClientConn) PaymentRepository {
	return PaymentRepository{
		client: paymentspb.NewPaymentsServiceClient(conn),
	}
}

func (r PaymentRepository) Confirm(ctx context.Context, paymentID string) error {
	_, err := r.client.ConfirmPayment(ctx, &paymentspb.ConfirmPaymentRequest{Id: paymentID})
	return err
}