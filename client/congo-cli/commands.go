package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gopheracademy/congo/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

type (
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
