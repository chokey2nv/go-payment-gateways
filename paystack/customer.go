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

func (c *CustomerService) CreateCustomer(ctx context.Context, req models.Customer) (*models.Customer, error) {
	var res models.Customer
	_, err := c.client.Do(ctx, "POST", "/customer", req, &res)
	return &res, err
}

// query parameter -> status, perPage, page, interval, amount
func (c *CustomerService) ListCustomer(ctx context.Context) (*[]models.Customer, error) {
	var res []models.Customer
	_, err := c.client.Do(ctx, "GET", "/customer", nil, &res)
	return &res, err
}
func (c *CustomerService) FetchCustomer(ctx context.Context, idOrCode string) (*models.Customer, error) {
	var res models.Customer
	_, err := c.client.Do(ctx, "GET", "/customer/"+idOrCode, nil, &res)
	return &res, err
}

// update plan - will return nil when successful
func (c *CustomerService) UpdateCustomer(ctx context.Context, idOrCode string, req models.Customer) error {
	var res models.Customer
	_, err := c.client.Do(ctx, "PUT", "/customer/"+idOrCode, req, &res)
	return err
}
