package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gopheracademy/congo/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

type (
	// CreateProposalCommand is the command line data structure for the create action of proposal
	CreateProposalCommand struct {
		// Path is the HTTP request path.
		Path    string
		Payload string
	}
	// DeleteProposalCommand is the command line data structure for the delete action of proposal
	DeleteProposalCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// ListProposalCommand is the command line data structure for the list action of proposal
	ListProposalCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// ShowProposalCommand is the command line data structure for the show action of proposal
	ShowProposalCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// UpdateProposalCommand is the command line data structure for the update action of proposal
	UpdateProposalCommand struct {
		// Path is the HTTP request path.
		Path    string
		Payload string
	}
	// CreateReviewCommand is the command line data structure for the create action of review
	CreateReviewCommand struct {
		// Path is the HTTP request path.
		Path    string
		Payload string
	}
	// DeleteReviewCommand is the command line data structure for the delete action of review
	DeleteReviewCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// ListReviewCommand is the command line data structure for the list action of review
	ListReviewCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// ShowReviewCommand is the command line data structure for the show action of review
	ShowReviewCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// UpdateReviewCommand is the command line data structure for the update action of review
	UpdateReviewCommand struct {
		// Path is the HTTP request path.
		Path    string
		Payload string
	}
	// CreateUserCommand is the command line data structure for the create action of user
	CreateUserCommand struct {
		// Path is the HTTP request path.
		Path    string
		Payload string
	}
	// DeleteUserCommand is the command line data structure for the delete action of user
	DeleteUserCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// ListUserCommand is the command line data structure for the list action of user
	ListUserCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// ShowUserCommand is the command line data structure for the show action of user
	ShowUserCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// UpdateUserCommand is the command line data structure for the update action of user
	UpdateUserCommand struct {
		// Path is the HTTP request path.
		Path    string
		Payload string
	}
)

// Run makes the HTTP request corresponding to the CreateProposalCommand command.
func (cmd *CreateProposalCommand) Run(c *client.Client) (*http.Response, error) {
	var payload client.CreateProposalPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	return c.CreateProposal(cmd.Path, &payload)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateProposalCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID/proposals`).Required().StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}

// Run makes the HTTP request corresponding to the DeleteProposalCommand command.
func (cmd *DeleteProposalCommand) Run(c *client.Client) (*http.Response, error) {
	return c.DeleteProposal(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteProposalCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID/proposals/:proposalID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ListProposalCommand command.
func (cmd *ListProposalCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ListProposal(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListProposalCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID/proposals`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ShowProposalCommand command.
func (cmd *ShowProposalCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ShowProposal(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowProposalCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID/proposals/:proposalID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the UpdateProposalCommand command.
func (cmd *UpdateProposalCommand) Run(c *client.Client) (*http.Response, error) {
	var payload client.UpdateProposalPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	return c.UpdateProposal(cmd.Path, &payload)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateProposalCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID/proposals/:proposalID`).Required().StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}

// Run makes the HTTP request corresponding to the CreateReviewCommand command.
func (cmd *CreateReviewCommand) Run(c *client.Client) (*http.Response, error) {
	var payload client.CreateReviewPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	return c.CreateReview(cmd.Path, &payload)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateReviewCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID/proposals/:proposalID/review`).Required().StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}

// Run makes the HTTP request corresponding to the DeleteReviewCommand command.
func (cmd *DeleteReviewCommand) Run(c *client.Client) (*http.Response, error) {
	return c.DeleteReview(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteReviewCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID/proposals/:proposalID/review/:reviewID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ListReviewCommand command.
func (cmd *ListReviewCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ListReview(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListReviewCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID/proposals/:proposalID/review`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ShowReviewCommand command.
func (cmd *ShowReviewCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ShowReview(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowReviewCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID/proposals/:proposalID/review/:reviewID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the UpdateReviewCommand command.
func (cmd *UpdateReviewCommand) Run(c *client.Client) (*http.Response, error) {
	var payload client.UpdateReviewPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	return c.UpdateReview(cmd.Path, &payload)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateReviewCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID/proposals/:proposalID/review/:reviewID`).Required().StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}

// Run makes the HTTP request corresponding to the CreateUserCommand command.
func (cmd *CreateUserCommand) Run(c *client.Client) (*http.Response, error) {
	var payload client.CreateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	return c.CreateUser(cmd.Path, &payload)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateUserCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, default is "/api/users"`).Default("/api/users").StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}

// Run makes the HTTP request corresponding to the DeleteUserCommand command.
func (cmd *DeleteUserCommand) Run(c *client.Client) (*http.Response, error) {
	return c.DeleteUser(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteUserCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ListUserCommand command.
func (cmd *ListUserCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ListUser(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUserCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, default is "/api/users"`).Default("/api/users").StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ShowUserCommand command.
func (cmd *ShowUserCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ShowUser(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUserCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the UpdateUserCommand command.
func (cmd *UpdateUserCommand) Run(c *client.Client) (*http.Response, error) {
	var payload client.UpdateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	return c.UpdateUser(cmd.Path, &payload)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateUserCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/users/:userID`).Required().StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}
