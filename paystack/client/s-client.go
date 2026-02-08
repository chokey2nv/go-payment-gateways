package client

import "net/http"

type Client struct {
	httpClient *http.Client
	baseURL    string
	headers    map[string]string
}
