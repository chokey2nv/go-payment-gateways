package paystack

import (
	"context"
	"strconv"

	"github.com/chokey2nv/go-payment-gateways/paystack/client"
	"github.com/chokey2nv/go-payment-gateways/paystack/models"
)

type TransactionService struct {
	client *client.PayStackClient
}

func NewTransaction(client *client.PayStackClient) *TransactionService {
	return &TransactionService{
		client: client,
	}
}
func (c *TransactionService) ExportTransaction(
	ctx context.Context,
) (*models.TransactionExport, error) {
	res, _, err := client.Do[models.TransactionExport](
		ctx,
		c.client,
		"GET",
		"/transaction/export",
		nil,
	)
	return res, err
}
func (c *TransactionService) TransactionTotals(
	ctx context.Context,
) (*models.TransactionTotals, error) {
	res, _, err := client.Do[models.TransactionTotals](
		ctx,
		c.client,
		"GET",
		"/transaction/totals",
		nil,
	)
	return res, err
}
func (c *TransactionService) ViewTransactionTimeline(
	ctx context.Context,
	idOrReference string,
) (*models.TransactionTimeline, error) {
	res, _, err := client.Do[models.TransactionTimeline](
		ctx,
		c.client,
		"GET",
		"/timeline/"+idOrReference,
		nil,
	)
	return res, err
}
func (c *TransactionService) InitializeTransaction(
	ctx context.Context,
	req *models.InitializeTransactionRequest,
) (*models.InitializeTransactionResponse, error) {
	res, _, err := client.Do[models.InitializeTransactionResponse](
		ctx,
		c.client,
		"POST",
		"/transaction/initialize",
		req,
	)
	return res, err
}

// -------- Verify Transaction --------

func (c *TransactionService) VerifyTransaction(
	ctx context.Context,
	reference string,
) (*models.VerifyTransactionResponse, error) {
	var res models.VerifyTransactionResponse
	_, err := c.client.Do(ctx, "GET", "/transaction/verify/"+reference, nil, &res)
	return &res, err
}

func (c *TransactionService) ListTransactions(
	ctx context.Context,
) (*[]*models.Transaction, *client.Meta, error) {
	res, meta, err := client.Do[[]*models.Transaction](
		ctx,
		c.client,
		"GET",
		"/transaction",
		nil,
	)
	return res, meta, err
}

func (c *TransactionService) FetchTransaction(
	ctx context.Context,
	id int64,
) (*models.Transaction, error) {
	res, _, err := client.Do[models.Transaction](
		ctx,
		c.client,
		"GET",
		"/transaction/"+strconv.Itoa(int(id)),
		nil,
	)
	return res, err
}
func (c *TransactionService) ChargeAuthorization(
	ctx context.Context,
	req *models.ChargeAuthorizationRequest,
) (*models.Transaction, error) {
	res, _, err := client.Do[models.Transaction](
		ctx,
		c.client,
		"POST",
		"/transaction/charge_authorization",
		req,
	)
	return res, err
}
