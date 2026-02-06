package client

import "net/http"

type Option func(*PayStackClient)

func WithHTTPClient(h *http.Client) Option {
	return func(c *PayStackClient) {
		c.http = h
	}
}
