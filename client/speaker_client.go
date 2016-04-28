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

// CreateSpeakerPayload is the speaker create action payload.
type CreateSpeakerPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	FirstName string  `json:"first_name" xml:"first_name"`
	Github    *string `json:"github,omitempty" xml:"github,omitempty"`
	ImageURL  *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	LastName  string  `json:"last_name" xml:"last_name"`
	Linkedin  *string `json:"linkedin,omitempty" xml:"linkedin,omitempty"`
	Twitter   *string `json:"twitter,omitempty" xml:"twitter,omitempty"`
}

// CreateSpeakerPath computes a request path to the create action of speaker.
func CreateSpeakerPath(eventID string, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers", tenantID, eventID)
}

// Record new speaker
func (c *Client) CreateSpeaker(ctx context.Context, path string, payload *CreateSpeakerPayload) (*http.Response, error) {
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

// DeleteSpeakerPath computes a request path to the delete action of speaker.
func DeleteSpeakerPath(eventID string, speakerID int, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID)
}

// DeleteSpeaker makes a request to the delete action endpoint of the speaker resource
func (c *Client) DeleteSpeaker(ctx context.Context, path string) (*http.Response, error) {
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

// ListSpeakerPath computes a request path to the list action of speaker.
func ListSpeakerPath(eventID string, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers", tenantID, eventID)
}

// List all speakers
func (c *Client) ListSpeaker(ctx context.Context, path string) (*http.Response, error) {
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

// ShowSpeakerPath computes a request path to the show action of speaker.
func ShowSpeakerPath(eventID string, speakerID int, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID)
}

// Retrieve speaker with given id
func (c *Client) ShowSpeaker(ctx context.Context, path string) (*http.Response, error) {
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

// UpdateSpeakerPayload is the speaker update action payload.
type UpdateSpeakerPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	Github    *string `json:"github,omitempty" xml:"github,omitempty"`
	ImageURL  *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	LastName  *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Linkedin  *string `json:"linkedin,omitempty" xml:"linkedin,omitempty"`
	Twitter   *string `json:"twitter,omitempty" xml:"twitter,omitempty"`
}

// UpdateSpeakerPath computes a request path to the update action of speaker.
func UpdateSpeakerPath(eventID string, speakerID int, tenantID string) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID)
}

// UpdateSpeaker makes a request to the update action endpoint of the speaker resource
func (c *Client) UpdateSpeaker(ctx context.Context, path string, payload *UpdateSpeakerPayload) (*http.Response, error) {
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
