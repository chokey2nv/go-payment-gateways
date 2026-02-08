package paystack

import "github.com/chokey2nv/go-payment-gateways/paystack/client"

type Paystack struct {
	Client       *client.PayStackClient
	Transaction  *TransactionService
	Subscription *SubscriptionService
	Plan         *PlanService
	Customer     *CustomerService
}

func New(apiKey string) *Paystack {
	client := client.New(apiKey)
	return &Paystack{
		Transaction:  NewTransaction(client),
		Subscription: NewSubscriptionService(client),
		Plan:         NewPlanService(client),
		Customer:     NewCustomerService(client),
	}
}
