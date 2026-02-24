package models

type PlanSubscription struct {
	Customer         int    `json:"customer,omitempty"`
	Plan             int    `json:"plan,omitempty"`
	Integration      int    `json:"integration,omitempty"`
	Domain           string `json:"domain,omitempty"`
	Start            int    `json:"start,omitempty"`
	Status           string `json:"status,omitempty"`
	Quantity         int    `json:"quantity,omitempty"`
	Amount           int    `json:"amount,omitempty"`
	SubscriptionCode string `json:"subscription_code,omitempty"`
	EmailToken       string `json:"email_token,omitempty"`
	Authorization    struct {
		AuthorizationCode string `json:"authorization_code,omitempty"`
		Bin               string `json:"bin,omitempty"`
		Last4             string `json:"last4,omitempty"`
		ExpMonth          int    `json:"exp_month,omitempty"`
		ExpYear           int    `json:"exp_year,omitempty"`
		Channel           string `json:"channel,omitempty"`
		CardType          string `json:"card_type,omitempty"`
		Bank              string `json:"bank,omitempty"`
		CountryCode       string `json:"country_code,omitempty"`
		Brand             string `json:"brand,omitempty"`
		Reusable          bool   `json:"reusable,omitempty"`
		Signature         string `json:"signature,omitempty"`
		AccountName       string `json:"account_name,omitempty"`
	} `json:"authorization,omitempty"`
	EasyCronId      string `json:"easy_cron_id,omitempty"`    // used for recurring subscription
	CronExpression  string `json:"cron_expression,omitempty"` // used for recurring subscription eg "0 0 * * 0"
	NextPaymentDate string `json:"next_payment_date,omitempty"`
	OpenInvoice     bool   `json:"open_invoice,omitempty"`
	Id              int    `json:"id,omitempty"`
	CreatedAt       string `json:"createdAt,omitempty"`
	UpdatedAt       string `json:"updatedAt,omitempty"`
}

type PlanInterval = string

const (
	Daily      PlanInterval = "daily"
	Weekly     PlanInterval = "weekly"
	Quarterly  PlanInterval = "quarterly"
	Biannually PlanInterval = "biannually"
	Annually   PlanInterval = "annually"
)

type Plan struct {
	Subscriptions     []PlanSubscription `json:"subscriptions,omitempty"` // omit nil
	Amount            int                `json:"amount,omitempty"`
	Domain            string             `json:"domain,omitempty"`
	Interval          PlanInterval       `json:"interval,omitempty"`
	Integration       int                `json:"integration,omitempty"`
	InvoiceLimit      int                `json:"invoice_limit,omitempty"`
	IsArchived        bool               `json:"is_archived,omitempty"`
	Migrate           bool               `json:"migrate,omitempty"`
	Name              string             `json:"name,omitempty"`
	PlanCode          string             `json:"plan_code,omitempty"`
	SendInvoices      bool               `json:"send_invoices,omitempty"`
	SendSms           bool               `json:"send_sms,omitempty"`
	HostedPage        bool               `json:"hosted_page,omitempty"`
	HostedPageUrl     string             `json:"hosted_page_url,omitempty"`
	HostedPageSummary string             `json:"hosted_page_summary,omitempty"`
	Currency          string             `json:"currency,omitempty"`
	ID                int                `json:"id,omitempty"`
	CreatedAt         string             `json:"createdAt,omitempty"`
	UpdatedAt         string             `json:"updatedAt,omitempty"`
}

type CreatePlanRequest struct {
	Name     string       `json:"name"`
	Amount   int          `json:"amount"` // in kobo
	Interval PlanInterval `json:"interval"`
	Currency string       `json:"currency,omitempty"`
}
