package paystack

import "context"

// -------- Initialize Transaction --------

type Channel = string

// channels "card", "bank", "apple_pay", "ussd", "qr", "mobile_money", "bank_transfer", "eft", "payattitude"
const (
	ChannelCard         Channel = "card"
	ChannelBank         Channel = "bank"
	ChannelApplePay     Channel = "apple_pay"
	ChannelUssd         Channel = "ussd"
	ChannelQr           Channel = "qr"
	ChannelMobileMoney  Channel = "mobile_money"
	ChannelBankTransfer Channel = "bank_transfer"
	ChannelEft          Channel = "eft"
	ChannelPayattitude  Channel = "payattitude"
)

type InitializeTransactionRequest struct {
	Email  string `json:"email"`
	Amount int64  `json:"amount"`

	Channels     []Channel `json:"channels,omitempty"`
	Currency     string    `json:"currency,omitempty"`
	Reference    string    `json:"reference,omitempty"`    //Unique transaction reference. Only -, ., = and alphanumeric characters allowed.
	CallbackURL  string    `json:"callback_url,omitempty"` //Fully qualified url, e.g. https://example.com/ . Use this to override the callback url provided on the dashboard for this transaction
	Plan         string    `json:"plan,omitempty"`
	InvoiceLimit int       `json:"invoice_limit,omitempty"`
	Metadata     string    `json:"metadata,omitempty"`
	SplitCode    string    `json:"split_code,omitempty"`
	Bearer       string    `json:"bearer,omitempty"`
}

type InitializeTransactionResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

func (c *PayStackClient) InitializeTransaction(
	ctx context.Context,
	req InitializeTransactionRequest,
) (*InitializeTransactionResponse, error) {
	var res InitializeTransactionResponse
	err := c.do(ctx, "POST", "/transaction/initialize", req, &res)
	return &res, err
}

// -------- Verify Transaction --------

type VerifyTransactionResponse struct {
	ID        int64  `json:"id"`
	Status    string `json:"status"`
	Reference string `json:"reference"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	PaidAt    string `json:"paid_at"`
	Gateway   string `json:"gateway_response"`

	Customer struct {
		Email string `json:"email"`
	} `json:"customer"`

	Authorization struct {
		AuthorizationCode string `json:"authorization_code"`
		Reusable          bool   `json:"reusable"`
	} `json:"authorization"`
}

func (c *PayStackClient) VerifyTransaction(
	ctx context.Context,
	reference string,
) (*VerifyTransactionResponse, error) {
	var res VerifyTransactionResponse
	err := c.do(ctx, "GET", "/transaction/verify/"+reference, nil, &res)
	return &res, err
}

// -------- List Transactions --------

type ListTransactionsRequest struct {
	PerPage int
	Page    int
}

type Transaction struct {
	ID        int64  `json:"id"`
	Status    string `json:"status"`
	Reference string `json:"reference"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	CreatedAt string `json:"created_at"`
	PaidAt    string `json:"paid_at"`
}

type ListTransactionsResponse struct {
	Transactions []Transaction
	Meta         *Meta
}

func (c *PayStackClient) ListTransactions(
	ctx context.Context,
) (*ListTransactionsResponse, error) {
	var data []Transaction
	var meta Meta

	err := c.do(ctx, "GET", "/transaction", nil, &struct {
		Data []Transaction `json:"data"`
		Meta Meta          `json:"meta"`
	}{
		Data: data,
		Meta: meta,
	})

	return &ListTransactionsResponse{
		Transactions: data,
		Meta:         &meta,
	}, err
}
