package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TokenProvider interface {
	GetToken(ctx context.Context) (string, error)
}

type Client struct {
	baseURL string
	auth    TokenProvider
	http    *http.Client
}

func New(baseURL string, auth TokenProvider) *Client {
	return &Client{
		baseURL: baseURL,
		auth:    auth,
		http:    &http.Client{},
	}
}

type RequestOption func(*http.Request)

func WithTraceID(traceID string) RequestOption {
	return func(r *http.Request) {
		if traceID != "" {
			r.Header.Set("X-Trace-Id", traceID)
		}
	}
}
func WithIdempotency(idempotencyKey string) RequestOption {
	return func(r *http.Request) {
		if idempotencyKey != "" {
			r.Header.Set("X-Idempotency-Key", idempotencyKey)
		}
	}
}

func (c *Client) Do(
	ctx context.Context,
	method string,
	path string,
	body any,
	out any,
	opts ...RequestOption,
) error {
	var reader io.Reader

	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.baseURL+path,
		reader,
	)
	if err != nil {
		return err
	}

	token, err := c.auth.GetToken(ctx)

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	for _, opt := range opts {
		opt(req)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("api error (%d): %s", resp.StatusCode, string(data))
	}

	if out != nil {
		return json.Unmarshal(data, out)
	}

	return nil
}

func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
	return c.auth.GetToken(ctx)
}
