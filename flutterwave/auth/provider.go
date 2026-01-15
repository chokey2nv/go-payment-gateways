package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	flutterwave_types "github.com/chokey2nv/go-payment-gateways/flutterwave/types"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type Provider struct {
	clientID     string
	clientSecret string
	tokenURL     string

	mu     sync.Mutex
	token  *flutterwave_types.GenerateAccessTokenResponse
	expiry time.Time

	http *http.Client
}

func New(clientID, clientSecret string) *Provider {

	return &Provider{
		clientID:     clientID,
		clientSecret: clientSecret,
		tokenURL:     "https://idp.flutterwave.com/realms/flutterwave/protocol/openid-connect/token",
		http: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (p *Provider) GetToken(ctx context.Context) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.token != nil && time.Now().Before(p.expiry) {
		return p.token.AccessToken, nil
	}

	if err := p.fetch(ctx); err != nil {
		return "", err
	}

	return p.token.AccessToken, nil
}

func (p *Provider) fetch(ctx context.Context) error {
	form := url.Values{}
	form.Set("client_id", p.clientID)
	form.Set("client_secret", p.clientSecret)
	form.Set("grant_type", "client_credentials")

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		p.tokenURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := p.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf(
			"auth token request failed: status=%d body=%s",
			resp.StatusCode,
			string(raw),
		)
	}

	var token flutterwave_types.GenerateAccessTokenResponse
	if err := json.Unmarshal(raw, &token); err != nil {
		return err
	}

	ttl := time.Duration(token.ExpiresIn) * time.Second
	if ttl > 30*time.Second {
		ttl -= 30 * time.Second
	}

	p.token = &token
	p.expiry = time.Now().Add(ttl)

	return nil
}
