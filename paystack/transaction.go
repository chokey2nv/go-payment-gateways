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
	var res models.TransactionExport
	_, err := c.client.Do(ctx, "GET", "/transaction/export", nil, &res)
	return &res, err
}
func (c *TransactionService) TransactionTotals(
	ctx context.Context,
) (*models.TransactionTotals, error) {
	var res models.TransactionTotals
	_, err := c.client.Do(ctx, "GET", "/transaction/totals", nil, &res)
	return &res, err
}
func (c *TransactionService) ViewTransactionTimeline(
	ctx context.Context,
	idOrReference string,
) (*models.TransactionTimeline, error) {
	var res models.TransactionTimeline
	_, err := c.client.Do(ctx, "GET", "/timeline/"+idOrReference, nil, &res)
	return &res, err
}
func (c *TransactionService) InitializeTransaction(
	ctx context.Context,
	req models.InitializeTransactionRequest,
) (*models.InitializeTransactionResponse, error) {
	var res models.InitializeTransactionResponse
	_, err := c.client.Do(ctx, "POST", "/transaction/initialize", req, &res)
	return &res, err
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
) (*models.ListTransactionsResponse, error) {
	var data []models.Transaction
	var meta client.Meta

	_, err := c.client.Do(ctx, "GET", "/transaction", nil, &struct {
		Data []models.Transaction `json:"data"`
		Meta client.Meta          `json:"meta"`
	}{
		Data: data,
		Meta: meta,
	})

	return &models.ListTransactionsResponse{
		Transactions: data,
		Meta:         &meta,
	}, err
}

func (c *TransactionService) FetchTransaction(
	ctx context.Context,
	id int64,
) (*models.Transaction, error) {
	var res models.Transaction
	_, err := c.client.Do(ctx, "GET", "/transaction/"+strconv.Itoa(int(id)), nil, &res)
	return &res, err
}
func (c *TransactionService) ChargeAuthorization(
	ctx context.Context,
	req models.ChargeAuthorizationRequest,
) (*models.Transaction, error) {
	var res models.Transaction
	_, err := c.client.Do(ctx, "POST", "/transaction/charge_authorization", req, &res)
	return &res, err
}
