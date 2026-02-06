package plan

type Subscription struct {
	Customer         int    `json:"customer"`
	Plan             int    `json:"plan"`
	Integration      int    `json:"integration"`
	Domain           string `json:"domain"`
	Start            int    `json:"start"`
	Status           string `json:"status"`
	Quantity         int    `json:"quantity"`
	Amount           int    `json:"amount"`
	SubscriptionCode string `json:"subscription_code"`
	EmailToken       string `json:"email_token"`
	Authorization    struct {
		AuthorizationCode string `json:"authorization_code"`
		Bin               string `json:"bin"`
		Last4             string `json:"last4"`
		ExpMonth          int    `json:"exp_month"`
		ExpYear           int    `json:"exp_year"`
		Channel           string `json:"channel"`
		CardType          string `json:"card_type"`
		Bank              string `json:"bank"`
		CountryCode       string `json:"country_code"`
		Brand             string `json:"brand"`
		Reusable          bool   `json:"reusable"`
		Signature         string `json:"signature"`
		AccountName       string `json:"account_name"`
	} `json:"authorization"`
	EasyCronId      string `json:"easy_cron_id"`    // used for recurring subscription
	CronExpression  string `json:"cron_expression"` // used for recurring subscription eg "0 0 * * 0"
	NextPaymentDate string `json:"next_payment_date"`
	OpenInvoice     bool   `json:"open_invoice"`
	Id              int    `json:"id"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

type Plan struct {
	Subscriptions     []Subscription `json:"subscriptions"`
	Name              string         `json:"name"`
	Amount            string         `json:"amount"`
	Interval          PlanInterval   `json:"interval"`
	Integration       int            `json:"integration"`
	Domain            string         `json:"domain"`
	PlanCode          string         `json:"plan_code"`
	SendInvoices      bool           `json:"send_invoices"`
	SendSms           bool           `json:"send_sms"`
	HostedPage        bool           `json:"hosted_page"`
	HostedPageUrl     string         `json:"hosted_page_url"`
	HostedPageSummary string         `json:"hosted_page_summary"`
	Currency          string         `json:"currency"`
	ID                int            `json:"id"`
	CreatedAt         string         `json:"createdAt"`
	UpdatedAt         string         `json:"updatedAt"`
}

type PlanInterval = string

const (
	Daily      PlanInterval = "daily"
	Weekly     PlanInterval = "weekly"
	Quarterly  PlanInterval = "quarterly"
	Biannually PlanInterval = "biannually"
	Annually   PlanInterval = "annually"
)

type CreatePlanRequest struct {
	Name     string       `json:"name"`
	Interval PlanInterval `json:"interval"`
	Amount   string       `json:"amount"`
}
