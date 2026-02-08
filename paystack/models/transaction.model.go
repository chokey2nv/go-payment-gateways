package models

import "github.com/chokey2nv/go-payment-gateways/paystack/client"

type TransactionExport struct {
	Path      string `json:"path"`
	ExpiresAt string `json:"expiresAt"`
}
type TransactionTotals struct {
	TotalTransactions     int64 `json:"total_transactions"`
	TotalVolume           int64 `json:"total_volume"`
	TotalVolumeByCurrency []struct {
		Currency string `json:"currency"`
		Amount   int64  `json:"amount"`
	} `json:"total_volume_by_currency"`
	PendingTransfers           int64 `json:"pending_transfers"`
	PendingTransfersByCurrency []struct {
		Currency string `json:"currency"`
		Amount   int64  `json:"amount"`
	} `json:"pending_transfers_by_currency"`
}
type TransactionTimeline struct {
	StartTime int64 `json:"start_time"`
	TimeSpent int64 `json:"time_spent"`
	Attempts  int64 `json:"attempts"`
	Errors    int64 `json:"errors"`
	Success   bool  `json:"success"`
	Mobile    bool  `json:"mobile"`
	Input     []any `json:"input"`
	History   []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Time    int64  `json:"time"`
	}
}

type ChargeAuthorizationRequest struct {
	Email             string `json:"email"`
	Amount            string `json:"amount"`
	AuthorizationCode string `json:"authorization_code"`
}

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

// -------- List Transactions --------

type ListTransactionsRequest struct {
	PerPage int
	Page    int
}

type ListTransactionsResponse struct {
	Transactions []Transaction
	Meta         *client.Meta
}

type Authorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
	Channel           string `json:"channel"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	Reusable          bool   `json:"reusable"`
	Signature         string `json:"signature"`
	AccountName       string `json:"account_name"`
}

type TransactionLog struct {
	StartTime int64 `json:"start_time"`
	TimeSpent int64 `json:"time_spent"`
	Attempts  int64 `json:"attempts"`
	Errors    int64 `json:"errors"`
	Success   bool  `json:"success"`
	Mobile    bool  `json:"mobile"`
	Input     []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"input"`
	History []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"history"`
}
type Transaction struct {
	ID                 int64          `json:"id"`
	Status             string         `json:"status"`
	Reference          string         `json:"reference"`
	ReceiptNumber      string         `json:"receipt_number"`
	Amount             int64          `json:"amount"`
	Message            string         `json:"message"`
	GatewayResponse    string         `json:"gateway_response"`
	HelpdeskLink       string         `json:"helpdesk_link"`
	PaidAt             string         `json:"paid_at"`
	CreatedAt          string         `json:"created_at"`
	Channel            string         `json:"channel"`
	Currency           string         `json:"currency"`
	IPAddress          string         `json:"ip_address"`
	Metadata           string         `json:"metadata"`
	Log                TransactionLog `json:"log"`
	Fees               int64          `json:"fees"`
	FeesSplit          bool           `json:"fees_split"`
	Customer           Customer       `json:"customer"`
	Authorization      Authorization  `json:"authorization"`
	Plan               any            `json:"plan"`
	SubAccount         any            `json:"subaccount"`
	RequestedAmount    int64          `json:"requested_amount"`
	PosTransactionData any            `json:"pos_transaction_data"`
	Source             struct {
		Type       string `json:"type"`
		Source     string `json:"source"`
		Identifier string `json:"identifier"`
	} `json:"source"`
	FeeBreakdown any `json:"fee_breakdown"`
	Connect      any `json:"connect"`
}
