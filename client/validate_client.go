package client

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
)

// ValidateValidatePath computes a request path to the validate action of validate.
func ValidateValidatePath(userID string) string {
	return fmt.Sprintf("/api/validate/%v", userID)
}

// validate user email
func (c *Client) ValidateValidate(ctx context.Context, path string, v string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("v", v)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}
