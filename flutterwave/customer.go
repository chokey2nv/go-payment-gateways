package flutterwave

import (
	"context"

	"github.com/chokey2nv/go-payment-gateways/flutterwave/types"
)

func (fw *Flutterwave) CreateCustomerObject(
	ctx context.Context,
	req *types.CreateCustomerObjectRequest,
	traceID string,
) (*types.CreateCustomerObjectResponse, error) {
	var resp types.CreateCustomerObjectResponse

	err := fw.Client.Do(
		ctx,
		"POST",
		"/customers",
		req,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
