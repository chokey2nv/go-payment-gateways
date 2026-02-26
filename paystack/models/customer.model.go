package models

type Customer struct {
	Authorizations []Authorization `json:"authorizations"`
	Subscriptions  []Subscription  `json:"subscriptions"`
	Integration    int             `json:"integration"`
	Id             int             `json:"id"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	Email          string          `json:"email"`
	CustomerCode   string          `json:"customer_code"`
	Phone          string          `json:"phone"`
	Metadata       struct {
		CustomFields []struct {
			DisplayName  string `json:"display_name"`
			VariableName string `json:"variable_name"`
			Value        string `json:"value"`
		} `json:"custom_fields"`
	} `json:"metadata"`
	RiskAction               string `json:"risk_action"`
	InternationalFormatPhone string `json:"international_format_phone"`
}
