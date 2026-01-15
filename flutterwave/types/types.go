package types

type GenerateAccessTokenRequest struct {
	GrantType    string `json:"grant_type"` // client_credentials
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
type GenerateAccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"` // 600
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"` // Bearer
	NotBeforePolicy  int    `json:"not-before-policy"`
	Scope            string `json:"scope"` // profile email
}

// Response

type FlutterApiResponse[T any] struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
}

// generate customer object
type CustomerAddress struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}
type CustomerName struct {
	First  string `json:"first"`
	Middle string `json:"middle"`
	Last   string `json:"last"`
}
type CustomerPhone struct {
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}
type CreateCustomerObjectRequest struct {
	Address CustomerAddress `json:"address"`
	Name    CustomerName    `json:"name"`
	Phone   CustomerPhone   `json:"phone"`
	Email   string          `json:"email"`
}

type Customer struct {
	ID          string            `json:"id"`
	Address     CustomerAddress   `json:"address"`
	Email       string            `json:"email"`
	Name        CustomerName      `json:"name"`
	Phone       CustomerPhone     `json:"phone"`
	Meta        map[string]string `json:"meta"`
	CreatedDate string            `json:"created_datetime"`
}
type CreateCustomerObjectResponse = FlutterApiResponse[Customer]

type Card struct {
	EncryptedCardNumber  string `json:"encrypted_card_number"`
	EncryptedExpiryMonth string `json:"encrypted_expiry_month"`
	EncryptedExpiryYear  string `json:"encrypted_expiry_year"`
	EncryptedCvv         string `json:"encrypted_cvv"`
	Nonce                string `json:"nonce"`
}
type CreateCardObjectRequest struct {
	Type string `json:"type"` // card
	Card Card   `json:"card"`
}

type CardProcessed struct {
	ExpiryMonth int    `json:"expiry_month"`
	ExpiryYear  int    `json:"expiry_year"`
	First6      string `json:"first6"`
	Last4       string `json:"last4"`
	Network     string `json:"network"` // mastercard
}
type CardObject struct {
	Type        string            `json:"type"` // card
	Card        CardProcessed     `json:"card"`
	ID          string            `json:"id"`
	Meta        map[string]string `json:"meta"`
	CreatedDate string            `json:"created_datetime"`
}
type CreateCardObjectResponse = FlutterApiResponse[CardObject]

type CardChargeRequest struct {
	Reference       string            `json:"reference"`
	Currency        string            `json:"currency"`
	CustomerID      string            `json:"customer_id"`
	PaymentMethodID string            `json:"payment_method_id"`
	RedirectURL     string            `json:"redirect_url"`
	Amount          int               `json:"amount"`
	Meta            map[string]string `json:"meta"`
}

/**
{
        "id": "chg_VoUhmFMhmF",
        "status": "pending",
        "next_action": {
            "type": "authorize",
            "authorization": {
                "type": "pin"
            }
        },
        ...
    }
		**/
type CardChargeObject struct {
	Id            string `json:"id"`
	Status        string `json:"status"` // pending
	NextAction    struct {
		Type string `json:"type"` // authorize
		Authorization struct {
			Type string `json:"type"` // pin
		}
	} `json:"next_action"`
}
type CardChargeResponse = FlutterApiResponse[CardChargeObject]

type PinRequest struct {
	Authorization struct {
		Type string `json:"type"` // pin
		Pin  struct {
			Nonce        string `json:"nonce"`
			EncryptedPin string `json:"encrypted_pin"`
		} `json:"pin"`
	} `json:"authorization"`
}

type PinResponse struct {
	ID         string `json:"id"`
	Status     string `json:"status"` // pending
	NextAction struct {
		Type        string `json:"type"` // redirect_url
		RedirectURL struct {
			URL string `json:"url"` // redirect the customer to this link to complete payment
		} `json:"redirect_url"`
	} `json:"next_action"`
	PaymentMethodDetails any `json:"payment_method_details"`
}
type PaymentMethod struct {
	Type              string `json:"type"` // card
	Card              Card   `json:"card"`
	ID                string `json:"id"`
	CustomerID        string `json:"customer_id"`
	Meta              any    `json:"meta"`
	DeviceFingerprint any    `json:"device_fingerprint"`
	ClientIP          any    `json:"client_ip"`
	CreatedDate       string `json:"created_datetime"`
}
type ProcessorResponse struct {
	Type string `json:"type"` // approved
	Code string `json:"code"` // 00
}
type SuccessEventHookData struct {
	ID                string            `json:"id"`
	Amount            int               `json:"amount"`
	Currency          string            `json:"currency"`
	Customer          Customer          `json:"customer"`
	Description       string            `json:"description"`
	Meta              any               `json:"meta"`
	PaymentMethod     PaymentMethod     `json:"payment_method"`
	RedirectURL       string            `json:"redirect_url"`
	Reference         string            `json:"reference"`
	Status            string            `json:"status"`
	ProcessorResponse ProcessorResponse `json:"processor_response"`
	CreatedDatetime   string            `json:"created_datetime"`
}
type SuccessEventHook struct {
	WebhookID  string `json:"webhook_id"`
	Timestamp  int64  `json:"timestamp"`
	Type       string `json:"type"` //charge.completed
	Data       any    `json:"data"`
	Signature  string `json:"signature"`
	Secret     string `json:"secret"`
	IsLiveMode bool   `json:"is_live_mode"`
}
