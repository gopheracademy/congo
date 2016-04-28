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

// CreatePresentationPayload is the presentation create action payload.
type CreatePresentationPayload struct {
	Abstract string  `json:"abstract" xml:"abstract"`
	Detail   *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

// CreatePresentationPath computes a request path to the create action of presentation.
func CreatePresentationPath(eventID string, speakerID string, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations", tenantID, eventID, speakerID)
}

// Record new presentation
func (c *Client) CreatePresentation(ctx context.Context, path string, payload *CreatePresentationPayload) (*http.Response, error) {
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

// DeletePresentationPath computes a request path to the delete action of presentation.
func DeletePresentationPath(eventID string, presentationID int, speakerID string, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID)
}

// DeletePresentation makes a request to the delete action endpoint of the presentation resource
func (c *Client) DeletePresentation(ctx context.Context, path string) (*http.Response, error) {
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

// ListPresentationPath computes a request path to the list action of presentation.
func ListPresentationPath(eventID string, speakerID string, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations", tenantID, eventID, speakerID)
}

// List all presentations
func (c *Client) ListPresentation(ctx context.Context, path string) (*http.Response, error) {
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

// ShowPresentationPath computes a request path to the show action of presentation.
func ShowPresentationPath(eventID string, presentationID int, speakerID string, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID)
}

// Retrieve presentation with given id
func (c *Client) ShowPresentation(ctx context.Context, path string) (*http.Response, error) {
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

// UpdatePresentationPayload is the presentation update action payload.
type UpdatePresentationPayload struct {
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail   *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

// UpdatePresentationPath computes a request path to the update action of presentation.
func UpdatePresentationPath(eventID string, presentationID int, speakerID string, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID)
}

// UpdatePresentation makes a request to the update action endpoint of the presentation resource
func (c *Client) UpdatePresentation(ctx context.Context, path string, payload *UpdatePresentationPayload) (*http.Response, error) {
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
