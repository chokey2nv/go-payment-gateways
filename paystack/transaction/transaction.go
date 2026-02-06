package transaction

import (
	"context"
	"strconv"

	"github.com/chokey2nv/go-payment-gateways/paystack/client"
)

type TransactionService struct {
	client *client.PayStackClient
}

func (c *TransactionService) ExportTransaction(
	ctx context.Context,
) (*TransactionExport, error) {
	var res TransactionExport
	_, err := c.client.Do(ctx, "GET", "/transaction/export", nil, &res)
	return &res, err
}
func (c *TransactionService) TransactionTotals(
	ctx context.Context,
) (*TransactionTotals, error) {
	var res TransactionTotals
	_, err := c.client.Do(ctx, "GET", "/transaction/totals", nil, &res)
	return &res, err
}
func (c *TransactionService) ViewTransactionTimeline(
	ctx context.Context,
	idOrReference string,
) (*TransactionTimeline, error) {
	var res TransactionTimeline
	_, err := c.client.Do(ctx, "GET", "/timeline/"+idOrReference, nil, &res)
	return &res, err
}
func (c *TransactionService) InitializeTransaction(
	ctx context.Context,
	req InitializeTransactionRequest,
) (*InitializeTransactionResponse, error) {
	var res InitializeTransactionResponse
	_, err := c.client.Do(ctx, "POST", "/transaction/initialize", req, &res)
	return &res, err
}

// -------- Verify Transaction --------

func (c *TransactionService) VerifyTransaction(
	ctx context.Context,
	reference string,
) (*VerifyTransactionResponse, error) {
	var res VerifyTransactionResponse
	_, err := c.client.Do(ctx, "GET", "/transaction/verify/"+reference, nil, &res)
	return &res, err
}

func (c *TransactionService) ListTransactions(
	ctx context.Context,
) (*ListTransactionsResponse, error) {
	var data []Transaction
	var meta client.Meta

	_, err := c.client.Do(ctx, "GET", "/transaction", nil, &struct {
		Data []Transaction `json:"data"`
		Meta client.Meta   `json:"meta"`
	}{
		Data: data,
		Meta: meta,
	})

	return &ListTransactionsResponse{
		Transactions: data,
		Meta:         &meta,
	}, err
}

func (c *TransactionService) FetchTransaction(
	ctx context.Context,
	id int64,
) (*Transaction, error) {
	var res Transaction
	_, err := c.client.Do(ctx, "GET", "/transaction/"+strconv.Itoa(int(id)), nil, &res)
	return &res, err
}
func (c *TransactionService) ChargeAuthorization(
	ctx context.Context,
	req ChargeAuthorizationRequest,
) (*Transaction, error) {
	var res Transaction
	_, err := c.client.Do(ctx, "POST", "/transaction/charge_authorization", req, &res)
	return &res, err
}
