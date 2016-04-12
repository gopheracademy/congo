package client

import (
	"github.com/goadesign/goa"
	"github.com/spf13/cobra"
)

type (
	// Client is the congo service client.
	Client struct {
		*goa.Client
	}

	// ActionCommand represents a single action command as defined on the command line.
	// Each command is associated with a generated client method and contains the logic to
	// call the method passing in arguments computed from the command line.
	ActionCommand interface {
		// Run makes the HTTP request and returns the response.
		Run(c *Client) error
		// RegisterFlags defines the command flags.
		RegisterFlags(*cobra.Command)
	}
)

// New instantiates the client.
func New() *Client {
	return &Client{Client: goa.NewClient()}
}
