package flutterwave

import (
	"context"
	"net/http"

	"github.com/chokey2nv/go-payment-gateways/flutterwave/client"
	"github.com/chokey2nv/go-payment-gateways/flutterwave/types"
)

func (fw *Flutterwave) CreateCardCharge(
	ctx context.Context,
	req *types.CardChargeRequest,
	traceID string,
	idempotencyKey string,
) (*types.CardChargeResponse, error) {
	var resp types.CardChargeResponse

	err := fw.Client.Do(
		ctx,
		"POST",
		"/charges",
		req,
		&resp,
		client.WithTraceID(traceID),
		client.WithIdempotency(idempotencyKey),
		func(r *http.Request) {
			r.Header.Set("X-Scenario-Key", "scenario:auth_3ds&issuer:approved")
		},
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
