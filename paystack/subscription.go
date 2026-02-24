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
func (c *SubscriptionService) UpdateSubscription(ctx context.Context, idOrCode string, req *models.Subscription) error {
	_, _, err := client.Do[models.Subscription](
		ctx,
		c.client,
		"PUT",
		"/subscription/"+idOrCode,
		req,
	)
	return err
}
func (c *SubscriptionService) CreateSubscription(ctx context.Context, req *models.CreateSubscriptionRequest) (*models.Subscription, error) {
	res, _, err := client.Do[models.Subscription](
		ctx,
		c.client,
		"POST",
		"/subscription",
		req,
	)
	return res, err
}

// query parameter -> status, perPage, page, interval, amount
func (c *SubscriptionService) ListSubscription(ctx context.Context) (*[]*models.Subscription, *client.Meta, error) {
	res, meta, err := client.Do[[]*models.Subscription](
		ctx,
		c.client,
		"GET",
		"/subscription",
		nil,
	)
	return res, meta, err
}
func (c *SubscriptionService) FetchSubscription(ctx context.Context, idOrCode string) (*models.Subscription, error) {
	res, _, err := client.Do[models.Subscription](
		ctx,
		c.client,
		"GET",
		"/subscription/"+idOrCode,
		nil,
	)
	return res, err
}
