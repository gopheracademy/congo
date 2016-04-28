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

// CreateTenantPayload is the tenant create action payload.
type CreateTenantPayload struct {
	Name string `json:"name" xml:"name"`
}

// CreateTenantPath computes a request path to the create action of tenant.
func CreateTenantPath() string {
	return fmt.Sprintf("/api/tenants")
}

// Record new tenant
func (c *Client) CreateTenant(ctx context.Context, path string, payload *CreateTenantPayload) (*http.Response, error) {
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
	return c.Client.Do(ctx, req)
}

// DeleteTenantPath computes a request path to the delete action of tenant.
func DeleteTenantPath(tenantID int) string {
	return fmt.Sprintf("/api/tenants/%v", tenantID)
}

// DeleteTenant makes a request to the delete action endpoint of the tenant resource
func (c *Client) DeleteTenant(ctx context.Context, path string) (*http.Response, error) {
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
	return c.Client.Do(ctx, req)
}

// ListTenantPath computes a request path to the list action of tenant.
func ListTenantPath() string {
	return fmt.Sprintf("/api/tenants")
}

// List all tenants
func (c *Client) ListTenant(ctx context.Context, path string) (*http.Response, error) {
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
	return c.Client.Do(ctx, req)
}

// ShowTenantPath computes a request path to the show action of tenant.
func ShowTenantPath(tenantID int) string {
	return fmt.Sprintf("/api/tenants/%v", tenantID)
}

// Retrieve tenant with given id
func (c *Client) ShowTenant(ctx context.Context, path string) (*http.Response, error) {
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
	return c.Client.Do(ctx, req)
}

// UpdateTenantPayload is the tenant update action payload.
type UpdateTenantPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// UpdateTenantPath computes a request path to the update action of tenant.
func UpdateTenantPath(tenantID int) string {
	return fmt.Sprintf("/api/tenants/%v", tenantID)
}

// UpdateTenant makes a request to the update action endpoint of the tenant resource
func (c *Client) UpdateTenant(ctx context.Context, path string, payload *UpdateTenantPayload) (*http.Response, error) {
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
	return c.Client.Do(ctx, req)
}
