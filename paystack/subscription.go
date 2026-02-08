package paystack

import (
	"context"

	"github.com/chokey2nv/go-payment-gateways/paystack/client"
	"github.com/chokey2nv/go-payment-gateways/paystack/models"
)

type SubscriptionService struct {
	client *client.PayStackClient
}

func NewSubscriptionService(client *client.PayStackClient) *SubscriptionService {
	return &SubscriptionService{
		client: client,
	}
}

// update plan - will return nil when successful
func (c *SubscriptionService) UpdateSubscription(ctx context.Context, idOrCode string, req models.Subscription) error {
	var res models.Subscription
	_, err := c.client.Do(ctx, "PUT", "/subscription/"+idOrCode, req, &res)
	return err
}
func (c *SubscriptionService) CreateSubscription(ctx context.Context, req models.CreateSubscriptionRequest) (*models.Subscription, error) {
	var res models.Subscription
	_, err := c.client.Do(ctx, "POST", "/subscription", req, &res)
	return &res, err
}

// query parameter -> status, perPage, page, interval, amount
func (c *SubscriptionService) ListSubscription(ctx context.Context) (*[]models.Subscription, error) {
	var res []models.Subscription
	_, err := c.client.Do(ctx, "GET", "/subscription", nil, &res)
	return &res, err
}
func (c *SubscriptionService) FetchSubscription(ctx context.Context, idOrCode string) (*models.Subscription, error) {
	var res models.Subscription
	_, err := c.client.Do(ctx, "GET", "/subscription/"+idOrCode, nil, &res)
	return &res, err
}
