//go:build e2e

package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cucumber/godog"
	"github.com/go-openapi/strfmt"
	"github.com/stackus/errors"

	"github.com/owezzy/soko-bora-mngt-system/payments/paymentsclient"
	"github.com/owezzy/soko-bora-mngt-system/payments/paymentsclient/invoice"
	"github.com/owezzy/soko-bora-mngt-system/payments/paymentsclient/models"
	"github.com/owezzy/soko-bora-mngt-system/payments/paymentsclient/payment"
)

type paymentIDKey struct{}

type paymentsFeature struct {
	client *paymentsclient.Payments
	db     *sql.DB
}

func (c *paymentsFeature) init(cfg featureConfig) (err error) {
	if cfg.useMonoDB {
		c.db, err = sql.Open("pgx", "postgres://mallbots_user:mallbots_pass@localhost:5432/mallbots?sslmode=disable")
	} else {
		c.db, err = sql.Open("pgx", "postgres://payments_user:payments_pass@localhost:5432/payments?sslmode=disable&search_path=payments,public")
	}
	if err != nil {
		return
	}
	c.client = paymentsclient.New(cfg.transport, strfmt.Default)

	return
}

func (c *paymentsFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I authorize payment for the basket total$`, c.iAuthorizePaymentForTheBasketTotal)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the payment (?:was|is) authorized$`, c.expectThePaymentWasAuthorized)
	ctx.Step(`^(?:I )?(?:ensure |expect )?a payment record exists for the basket total$`, c.expectAPaymentRecordExistsForTheBasketTotal)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the authorized payment belongs to the current customer$`, c.expectTheAuthorizedPaymentBelongsToTheCurrentCustomer)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the authorized payment amount is "([^"]*)"$`, c.expectTheAuthorizedPaymentAmountIs)
	ctx.Step(`^(?:I )?(?:ensure |expect )?an invoice exists for the checked out basket$`, c.expectAnInvoiceExistsForTheCheckedOutBasket)
	ctx.Step(`^I pay the invoice for the checked out basket$`, c.iPayTheInvoiceForTheCheckedOutBasket)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the invoice status is "([^"]*)"$`, c.expectTheInvoiceStatusIs)
	}

func (c *paymentsFeature) reset() {
	truncate := func(tableName string) {
		_, _ = c.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", tableName))
	}

	truncate("payments.payments")
	truncate("payments.invoices")
	truncate("payments.inbox")
	truncate("payments.outbox")
}

func (c *paymentsFeature) iAuthorizePaymentForTheBasketTotal(ctx context.Context) (context.Context, error) {
	customerID, err := lastCustomerID(ctx)
	if err != nil {
		return ctx, err
	}

	total, err := currentBasketTotal(ctx)
	if err != nil {
		return ctx, err
	}

	resp, err := c.client.Payment.AuthorizePayment(payment.NewAuthorizePaymentParams().WithBody(&models.PaymentspbAuthorizePaymentRequest{
		CustomerID: customerID,
		Amount:     total,
	}))
	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx, nil
	}

	return context.WithValue(ctx, paymentIDKey{}, resp.Payload.ID), nil
}

func (c *paymentsFeature) expectThePaymentWasAuthorized(ctx context.Context) error {
	if err := lastResponseWas(ctx, &payment.AuthorizePaymentOK{}); err != nil {
		return err
	}

	return nil
}

func (c *paymentsFeature) expectAPaymentRecordExistsForTheBasketTotal(ctx context.Context) error {
	customerID, err := lastCustomerID(ctx)
	if err != nil {
		return err
	}

	total, err := currentBasketTotal(ctx)
	if err != nil {
		return err
	}

	var paymentID string
	row := c.db.QueryRow("SELECT id FROM payments.payments WHERE customer_id = $1 AND amount = $2", customerID, total)
	if err = row.Scan(&paymentID); err != nil {
		if err == sql.ErrNoRows {
			return errors.ErrNotFound.Msgf("payment for customer `%s` and amount `%0.2f` does not exist", customerID, total)
		}
		return err
	}

	return nil
}

func (c *paymentsFeature) expectTheAuthorizedPaymentBelongsToTheCurrentCustomer(ctx context.Context) error {
	paymentID, err := lastPaymentID(ctx)
	if err != nil {
		return err
	}
	customerID, err := lastCustomerID(ctx)
	if err != nil {
		return err
	}

	var actualCustomerID string
	row := c.db.QueryRow("SELECT customer_id FROM payments.payments WHERE id = $1", paymentID)
	if err = row.Scan(&actualCustomerID); err != nil {
		return err
	}

	if actualCustomerID != customerID {
		return errors.ErrBadRequest.Msgf("expected payment customer `%s`, got `%s`", customerID, actualCustomerID)
	}

	return nil
}

func (c *paymentsFeature) expectTheAuthorizedPaymentAmountIs(ctx context.Context, amount float64) error {
	paymentID, err := lastPaymentID(ctx)
	if err != nil {
		return err
	}

	var actualAmount float64
	row := c.db.QueryRow("SELECT amount FROM payments.payments WHERE id = $1", paymentID)
	if err = row.Scan(&actualAmount); err != nil {
		return err
	}

	if !nearlyEqualFloat64(actualAmount, amount) {
		return errors.ErrBadRequest.Msgf("expected payment amount `%0.2f`, got `%0.2f`", amount, actualAmount)
	}

	return nil
}

func (c *paymentsFeature) iPayTheInvoiceForTheCheckedOutBasket(ctx context.Context) (context.Context, error) {
	invoiceID, err := c.waitForInvoiceID(ctx)
	if err != nil {
		return ctx, err
	}

	resp, err := c.client.Invoice.PayInvoice(invoice.NewPayInvoiceParams().WithID(invoiceID))
	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx, nil
	}

	return ctx, nil
}

func (c *paymentsFeature) expectAnInvoiceExistsForTheCheckedOutBasket(ctx context.Context) error {
	_, err := c.waitForInvoiceID(ctx)
	return err
}

func (c *paymentsFeature) expectTheInvoiceStatusIs(ctx context.Context, expected string) error {
	invoiceID, err := c.waitForInvoiceID(ctx)
	if err != nil {
		return err
	}

	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		var status string
		row := c.db.QueryRow("SELECT status FROM payments.invoices WHERE id = $1", invoiceID)
		if err = row.Scan(&status); err == nil && status == expected {
			return nil
		}
		time.Sleep(200 * time.Millisecond)
	}

	return errors.ErrUnknown.Msgf("invoice `%s` did not reach status `%s` before timeout", invoiceID, expected)
}

func (c *paymentsFeature) waitForInvoiceID(ctx context.Context) (string, error) {
	orderID, err := lastBasketID(ctx)
	if err != nil {
		return "", err
	}

	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		var invoiceID string
		row := c.db.QueryRow("SELECT id FROM payments.invoices WHERE order_id = $1", orderID)
		if err = row.Scan(&invoiceID); err == nil {
			return invoiceID, nil
		}
		time.Sleep(200 * time.Millisecond)
	}

	return "", errors.ErrUnknown.Msgf("invoice for order `%s` was not created before timeout", orderID)
}

func lastPaymentID(ctx context.Context) (string, error) {
	v := ctx.Value(paymentIDKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no payment ID to work with")
	}
	return v.(string), nil
}
