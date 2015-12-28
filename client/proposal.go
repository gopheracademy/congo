package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// CreateProposalPayload is the data structure used to initialize the proposal create request body.
type CreateProposalPayload struct {
	Abstract  string `json:"abstract,omitempty"`
	Detail    string `json:"detail,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Title     string `json:"title"`
	Withdrawn bool   `json:"withdrawn,omitempty"`
}

// Create a new proposal
func (c *Client) CreateProposal(path string, payload *CreateProposalPayload) (*http.Response, error) {
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

// DeleteProposal makes a request to the delete action endpoint of the proposal resource
func (c *Client) DeleteProposal(path string) (*http.Response, error) {
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

// List all proposals for a user
func (c *Client) ListProposal(path string) (*http.Response, error) {
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

// Retrieve proposal with given id
func (c *Client) ShowProposal(path string) (*http.Response, error) {
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

// UpdateProposalPayload is the data structure used to initialize the proposal update request body.
type UpdateProposalPayload struct {
	Abstract  string `json:"abstract,omitempty"`
	Detail    string `json:"detail,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Title     string `json:"title,omitempty"`
	Withdrawn bool   `json:"withdrawn,omitempty"`
}

// UpdateProposal makes a request to the update action endpoint of the proposal resource
func (c *Client) UpdateProposal(path string, payload *UpdateProposalPayload) (*http.Response, error) {
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
