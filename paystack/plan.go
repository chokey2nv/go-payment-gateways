package paystack

import (
	"context"

	"github.com/chokey2nv/go-payment-gateways/paystack/client"
	"github.com/chokey2nv/go-payment-gateways/paystack/models"
)

type PlanService struct {
	client *client.PayStackClient
}

func NewPlanService(client *client.PayStackClient) *PlanService {
	return &PlanService{
		client: client,
	}
}

// update plan - will return nil when successful
func (c *PlanService) UpdatePlan(ctx context.Context, idOrCode string, req models.Plan) (*models.Plan, error) {
	res, _, err := client.Do[models.Plan](
		ctx,
		c.client,
		"PUT",
		"/plan/"+idOrCode,
		req,
	)
	return res, err
}
func (c *PlanService) CreatePlan(ctx context.Context, req models.CreatePlanRequest) (*models.Plan, error) {
	res, _, err := client.Do[models.Plan](
		ctx,
		c.client,
		"POST",
		"/plan",
		req,
	)
	return res, err
}

// query parameter -> status, perPage, page, interval, amount
func (c *PlanService) ListPlan(ctx context.Context) (*[]models.Plan, *client.Meta, error) {
	res, meta, err := client.Do[[]models.Plan](
		ctx,
		c.client,
		"GET",
		"/plan",
		nil,
	)
	return res, meta, err
}
func (c *PlanService) FetchPlan(ctx context.Context, idOrCode string) (*models.Plan, error) {
	res, _, err := client.Do[models.Plan](
		ctx,
		c.client,
		"GET",
		"/plan/"+idOrCode,
		nil,
	)
	return res, err
}
