package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
)

// CreateSeriesPayload is the series create action payload.
type CreateSeriesPayload struct {
	Name string `json:"name" xml:"name"`
}

// CreateSeriesPath computes a request path to the create action of series.
func CreateSeriesPath(tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/series", tenantID)
}

// Record new series
func (c *Client) CreateSeries(ctx context.Context, path string, payload *CreateSeriesPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
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
	c.SignerJWT.Sign(ctx, req)
	return c.Client.Do(ctx, req)
}

// DeleteSeriesPath computes a request path to the delete action of series.
func DeleteSeriesPath(seriesID int, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/series/%v", tenantID, seriesID)
}

// DeleteSeries makes a request to the delete action endpoint of the series resource
func (c *Client) DeleteSeries(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	c.SignerJWT.Sign(ctx, req)
	return c.Client.Do(ctx, req)
}

// ListSeriesPath computes a request path to the list action of series.
func ListSeriesPath(tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/series", tenantID)
}

// List all series
func (c *Client) ListSeries(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	c.SignerJWT.Sign(ctx, req)
	return c.Client.Do(ctx, req)
}

// ShowSeriesPath computes a request path to the show action of series.
func ShowSeriesPath(seriesID int, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/series/%v", tenantID, seriesID)
}

// Retrieve series with given id
func (c *Client) ShowSeries(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	c.SignerJWT.Sign(ctx, req)
	return c.Client.Do(ctx, req)
}

// UpdateSeriesPayload is the series update action payload.
type UpdateSeriesPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// UpdateSeriesPath computes a request path to the update action of series.
func UpdateSeriesPath(seriesID int, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/series/%v", tenantID, seriesID)
}

// UpdateSeries makes a request to the update action endpoint of the series resource
func (c *Client) UpdateSeries(ctx context.Context, path string, payload *UpdateSeriesPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	c.SignerJWT.Sign(ctx, req)
	return c.Client.Do(ctx, req)
}
