package plan

import (
	"context"

	"github.com/chokey2nv/go-payment-gateways/paystack/client"
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
func (c *PlanService) UpdatePlan(ctx context.Context, idOrCode string, req CreatePlanRequest) error {
	var res Plan
	_, err := c.client.Do(ctx, "PUT", "/plan/"+idOrCode, req, &res)
	return err
}
func (c *PlanService) CreatePlan(ctx context.Context, req CreatePlanRequest) (*Plan, error) {
	var res Plan
	_, err := c.client.Do(ctx, "POST", "/plan", req, &res)
	return &res, err
}

// query parameter -> status, perPage, page, interval, amount
func (c *PlanService) ListPlan(ctx context.Context) (*[]Plan, error) {
	var res []Plan
	_, err := c.client.Do(ctx, "GET", "/plan", nil, &res)
	return &res, err
}
func (c *PlanService) FetchPlan(ctx context.Context, idOrCode string) (*Plan, error) {
	var res Plan
	_, err := c.client.Do(ctx, "GET", "/plan/"+idOrCode, nil, &res)
	return &res, err
}
