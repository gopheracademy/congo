package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
	"time"
)

// CreateEventPayload is the event create action payload.
type CreateEventPayload struct {
	EndDate   *time.Time `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Name      string     `json:"name" xml:"name"`
	StartDate *time.Time `json:"start_date,omitempty" xml:"start_date,omitempty"`
	URL       *string    `json:"url,omitempty" xml:"url,omitempty"`
}

// CreateEventPath computes a request path to the create action of event.
func CreateEventPath(tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events", tenantID)
}

// Record new event
func (c *Client) CreateEvent(ctx context.Context, path string, payload *CreateEventPayload) (*http.Response, error) {
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

// DeleteEventPath computes a request path to the delete action of event.
func DeleteEventPath(eventID int, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v", tenantID, eventID)
}

// DeleteEvent makes a request to the delete action endpoint of the event resource
func (c *Client) DeleteEvent(ctx context.Context, path string) (*http.Response, error) {
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

// ListEventPath computes a request path to the list action of event.
func ListEventPath(tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events", tenantID)
}

// List all events
func (c *Client) ListEvent(ctx context.Context, path string) (*http.Response, error) {
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

// ShowEventPath computes a request path to the show action of event.
func ShowEventPath(eventID int, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v", tenantID, eventID)
}

// Retrieve event with given id
func (c *Client) ShowEvent(ctx context.Context, path string) (*http.Response, error) {
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

// UpdateEventPayload is the event update action payload.
type UpdateEventPayload struct {
	EndDate   *time.Time `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Name      *string    `json:"name,omitempty" xml:"name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty" xml:"start_date,omitempty"`
	URL       *string    `json:"url,omitempty" xml:"url,omitempty"`
}

// UpdateEventPath computes a request path to the update action of event.
func UpdateEventPath(eventID int, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v", tenantID, eventID)
}

// UpdateEvent makes a request to the update action endpoint of the event resource
func (c *Client) UpdateEvent(ctx context.Context, path string, payload *UpdateEventPayload) (*http.Response, error) {
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
