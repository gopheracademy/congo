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
	c.SignerJWT.Sign(ctx, req)
	return c.Client.Do(ctx, req)
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
	c.SignerJWT.Sign(ctx, req)
	return c.Client.Do(ctx, req)
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
	c.SignerJWT.Sign(ctx, req)
	return c.Client.Do(ctx, req)
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
	c.SignerJWT.Sign(ctx, req)
	return c.Client.Do(ctx, req)
}
