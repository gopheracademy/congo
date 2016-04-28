package client

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
)

// RefreshAuthPath computes a request path to the refresh action of auth.
func RefreshAuthPath() string {
	return fmt.Sprintf("/api/auth/refresh")
}

// Obtain a refreshed access token
func (c *Client) RefreshAuth(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	c.SignerPassword.Sign(ctx, req)
	return c.Client.Do(ctx, req)
}

// TokenAuthPath computes a request path to the token action of auth.
func TokenAuthPath() string {
	return fmt.Sprintf("/api/auth/token")
}

// Obtain an access token
func (c *Client) TokenAuth(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	c.SignerPassword.Sign(ctx, req)
	return c.Client.Do(ctx, req)
}
