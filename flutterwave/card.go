package flutterwave

import (
	"context"

	"github.com/chokey2nv/go-payment-gateways/flutterwave/client"
	"github.com/chokey2nv/go-payment-gateways/flutterwave/types"
)

func (fw *Flutterwave) CreateCardObject(
	ctx context.Context,
	req *types.CreateCardObjectRequest,
	traceID string,
	idempotencyKey string,
) (*types.CreateCardObjectResponse, error) {
	var resp types.CreateCardObjectResponse

	
	err := fw.Client.Do(
		ctx,
		"POST",
		"/payment-methods",
		req,
		&resp,
		client.WithTraceID(traceID),
		client.WithIdempotency(idempotencyKey),
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
