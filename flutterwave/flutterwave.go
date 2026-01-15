package flutterwave

import (
	"github.com/chokey2nv/go-payment-gateways/flutterwave/auth"
	"github.com/chokey2nv/go-payment-gateways/flutterwave/client"
)

type Flutterwave struct {
	Client *client.Client
}

func New(clientID, clientSecret string) *Flutterwave {
	authProvider := auth.New(clientID, clientSecret)

	return &Flutterwave{
		Client: client.New(
			"https://developersandbox-api.flutterwave.com",
			authProvider,
		),
	}
}
