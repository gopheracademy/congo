package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gopheracademy/congo/app"
	"io"
	"net/http"
	"net/url"
)

// Create a new review
func (c *Client) CreateReview(path string, payload *app.CreateReviewPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// DeleteReview makes a request to the delete action endpoint of the review resource
func (c *Client) DeleteReview(path string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// List all reviews for a proposal
func (c *Client) ListReview(path string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// Retrieve review with given id
func (c *Client) ShowReview(path string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// UpdateReview makes a request to the update action endpoint of the review resource
func (c *Client) UpdateReview(path string, payload *app.UpdateReviewPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}
