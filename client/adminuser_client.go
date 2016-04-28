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

// CreateAdminuserPayload is the adminuser create action payload.
type CreateAdminuserPayload struct {
	Email          string  `json:"email" xml:"email"`
	FirstName      string  `json:"first_name" xml:"first_name"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       string  `json:"last_name" xml:"last_name"`
	Password       string  `json:"password" xml:"password"`
	Role           string  `json:"role" xml:"role"`
	TenantID       *int    `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// CreateAdminuserPath computes a request path to the create action of adminuser.
func CreateAdminuserPath() string {
	return fmt.Sprintf("/api/admin/users")
}

// Record new user
func (c *Client) CreateAdminuser(ctx context.Context, path string, payload *CreateAdminuserPayload) (*http.Response, error) {
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

// DeleteAdminuserPath computes a request path to the delete action of adminuser.
func DeleteAdminuserPath(userID int) string {
	return fmt.Sprintf("/api/admin/users/%v", userID)
}

// DeleteAdminuser makes a request to the delete action endpoint of the adminuser resource
func (c *Client) DeleteAdminuser(ctx context.Context, path string) (*http.Response, error) {
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

// ListAdminuserPath computes a request path to the list action of adminuser.
func ListAdminuserPath() string {
	return fmt.Sprintf("/api/admin/users")
}

// List all users
func (c *Client) ListAdminuser(ctx context.Context, path string) (*http.Response, error) {
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

// ShowAdminuserPath computes a request path to the show action of adminuser.
func ShowAdminuserPath(userID int) string {
	return fmt.Sprintf("/api/admin/users/%v", userID)
}

// Retrieve user with given id
func (c *Client) ShowAdminuser(ctx context.Context, path string) (*http.Response, error) {
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

// UpdateAdminuserPayload is the adminuser update action payload.
type UpdateAdminuserPayload struct {
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Role           *string `json:"role,omitempty" xml:"role,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// UpdateAdminuserPath computes a request path to the update action of adminuser.
func UpdateAdminuserPath(userID int) string {
	return fmt.Sprintf("/api/admin/users/%v", userID)
}

// UpdateAdminuser makes a request to the update action endpoint of the adminuser resource
func (c *Client) UpdateAdminuser(ctx context.Context, path string, payload *UpdateAdminuserPayload) (*http.Response, error) {
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
