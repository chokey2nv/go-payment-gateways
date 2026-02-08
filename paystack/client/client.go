package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	utils "github.com/chokey2nv/go-payment-gateways/utils"
)

type PayStackClient struct {
	secretKey string
	http      *http.Client
	baseURL   string
}

func New(secretKey string, opts ...Option) *PayStackClient {
	c := &PayStackClient{
		secretKey: secretKey,
		baseURL:   "https://api.paystack.co",
		http: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *PayStackClient) Do(
	ctx context.Context,
	method string,
	path string,
	body any,
	out any,
) (*Meta, error) {
	reqBody, err := encodeJSON(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.baseURL+path,
		reqBody,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.secretKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var envelope responseEnvelope
	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		return nil, err
	}

	if !envelope.Status {
		return nil, errors.New(envelope.Message)
	}

	utils.ErrorLog("*********************\n")
	utils.Errorf("Method: %s\n Path: %s\n", method, path)
	utils.ErrorLog(envelope.Data)
	utils.ErrorLog("*********************\n")
	// var logData  map[string]interface{}
	// if envelope.Data != nil {
	// 	logData = envelope.Data.(map[string]interface{})

	// }
	if out != nil {
		raw, _ := json.Marshal(envelope.Data)
		return envelope.Meta, json.Unmarshal(raw, out)
	}

	return envelope.Meta, nil
}
func Do[T any](
	ctx context.Context,
	client *PayStackClient,
	method string,
	path string,
	body any,
) (*T, *Meta, error) {
	reqBody, err := encodeJSON(body)
	if err != nil {
		return nil, nil, utils.Errorf(err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		client.baseURL+path,
		reqBody,
	)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Authorization", "Bearer "+client.secretKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.http.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var envelope responseEnvelope
	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		return nil, nil, utils.Errorf(err)
	}

	if !envelope.Status {
		return nil, nil, errors.New(envelope.Message)
	}

	utils.ErrorLog("*********************\n")
	utils.Errorf("Method: %s\n Path: %s\n", method, path)
	utils.ErrorLog(envelope.Data)
	utils.ErrorLog("*********************\n")
	// var logData  map[string]interface{}
	// if envelope.Data != nil {
	// 	logData = envelope.Data.(map[string]interface{})

	// }
	var out T
	raw, _ := json.Marshal(envelope.Data)
	err = json.Unmarshal(raw, &out)
	if err != nil {
		return nil, nil, utils.Errorf(err)
	}
	return &out, envelope.Meta, nil

}
