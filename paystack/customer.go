package paystack

import (
	"context"

	"github.com/chokey2nv/go-payment-gateways/paystack/client"
	"github.com/chokey2nv/go-payment-gateways/paystack/models"
)

type CustomerService struct {
	client *client.PayStackClient
}

func NewCustomerService(client *client.PayStackClient) *CustomerService {
	return &CustomerService{
		client: client,
	}
}

func (c *CustomerService) CreateCustomer(ctx context.Context, req *models.Customer) (*models.Customer, error) {
	res, _, err := client.Do[models.Customer](
		ctx,
		c.client,
		"POST",
		"/customer",
		req,
	)
	return res, err
}

// query parameter -> status, perPage, page, interval, amount
func (c *CustomerService) ListCustomer(ctx context.Context) (*[]*models.Customer, *client.Meta, error) {
	res, meta, err := client.Do[[]*models.Customer](
		ctx,
		c.client,
		"GET",
		"/customer",
		nil,
	)
	return res, meta, err
}
func (c *CustomerService) FetchCustomer(ctx context.Context, idOrCode string) (*models.Customer, error) {
	res, _, err := client.Do[models.Customer](
		ctx,
		c.client,
		"GET",
		"/customer/"+idOrCode,
		nil,
	)
	return res, err
}

// update plan - will return nil when successful
func (c *CustomerService) UpdateCustomer(ctx context.Context, idOrCode string, req *models.Customer) (*models.Customer, error) {
	res, _, err := client.Do[models.Customer](
		ctx,
		c.client,
		"PUT",
		"/customer/"+idOrCode,
		req,
	)
	return res, err
}
