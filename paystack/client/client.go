package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const baseURL = "https://api.paystack.co"

type PayStackClient struct {
	secretKey string
	http      *http.Client
}

func New(secretKey string, opts ...Option) *PayStackClient {
	c := &PayStackClient{
		secretKey: secretKey,
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
		baseURL+path,
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

	if out != nil {
		raw, _ := json.Marshal(envelope.Data)
		return envelope.Meta, json.Unmarshal(raw, out)
	}

	return envelope.Meta, nil
}
