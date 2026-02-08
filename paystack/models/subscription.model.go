package models

type CreateSubscriptionRequest struct {
	Customer string `json:"customer"`
	Plan     string `json:"plan"`
}

type Subscription struct {
	Customer         Customer `json:"customer"`
	Plan             Plan     `json:"plan"`
	Integration      int      `json:"integration"`
	Domain           string   `json:"domain"`
	Start            int      `json:"start"`
	Status           string   `json:"status"`
	Quantity         int      `json:"quantity"`
	Amount           int      `json:"amount"`
	SubscriptionCode string   `json:"subscription_code"`
	EmailToken       string   `json:"email_token"`
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
