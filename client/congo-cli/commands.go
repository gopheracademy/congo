package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alecthomas/kingpin"
	"github.com/gopheracademy/congo/client"
)

type (
	// CreateAccountCommand is the command line data structure for the create action of account
	CreateAccountCommand struct {
		// Path is the HTTP request path.
		Path    string
		Payload string
	}
	// DeleteAccountCommand is the command line data structure for the delete action of account
	DeleteAccountCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// ListAccountCommand is the command line data structure for the list action of account
	ListAccountCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// ShowAccountCommand is the command line data structure for the show action of account
	ShowAccountCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// UpdateAccountCommand is the command line data structure for the update action of account
	UpdateAccountCommand struct {
		// Path is the HTTP request path.
		Path    string
		Payload string
	}
	// CreateSeriesCommand is the command line data structure for the create action of series
	CreateSeriesCommand struct {
		// Path is the HTTP request path.
		Path    string
		Payload string
	}
	// DeleteSeriesCommand is the command line data structure for the delete action of series
	DeleteSeriesCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// ListSeriesCommand is the command line data structure for the list action of series
	ListSeriesCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// ShowSeriesCommand is the command line data structure for the show action of series
	ShowSeriesCommand struct {
		// Path is the HTTP request path.
		Path string
	}
	// UpdateSeriesCommand is the command line data structure for the update action of series
	UpdateSeriesCommand struct {
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

// Run makes the HTTP request corresponding to the CreateAccountCommand command.
func (cmd *CreateAccountCommand) Run(c *client.Client) (*http.Response, error) {
	var payload client.CreateAccountPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	return c.CreateAccount(cmd.Path, &payload)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateAccountCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, default is "/api/accounts"`).Default("/api/accounts").StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}

// Run makes the HTTP request corresponding to the DeleteAccountCommand command.
func (cmd *DeleteAccountCommand) Run(c *client.Client) (*http.Response, error) {
	return c.DeleteAccount(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteAccountCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ListAccountCommand command.
func (cmd *ListAccountCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ListAccount(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListAccountCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, default is "/api/accounts"`).Default("/api/accounts").StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ShowAccountCommand command.
func (cmd *ShowAccountCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ShowAccount(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowAccountCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the UpdateAccountCommand command.
func (cmd *UpdateAccountCommand) Run(c *client.Client) (*http.Response, error) {
	var payload client.UpdateAccountPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	return c.UpdateAccount(cmd.Path, &payload)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateAccountCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID`).Required().StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}

// Run makes the HTTP request corresponding to the CreateSeriesCommand command.
func (cmd *CreateSeriesCommand) Run(c *client.Client) (*http.Response, error) {
	var payload client.CreateSeriesPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	return c.CreateSeries(cmd.Path, &payload)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateSeriesCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID/series`).Required().StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}

// Run makes the HTTP request corresponding to the DeleteSeriesCommand command.
func (cmd *DeleteSeriesCommand) Run(c *client.Client) (*http.Response, error) {
	return c.DeleteSeries(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteSeriesCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID/series/:seriesID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ListSeriesCommand command.
func (cmd *ListSeriesCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ListSeries(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListSeriesCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID/series`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ShowSeriesCommand command.
func (cmd *ShowSeriesCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ShowSeries(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowSeriesCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID/series/:seriesID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the UpdateSeriesCommand command.
func (cmd *UpdateSeriesCommand) Run(c *client.Client) (*http.Response, error) {
	var payload client.UpdateSeriesPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	return c.UpdateSeries(cmd.Path, &payload)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateSeriesCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID/series/:seriesID`).Required().StringVar(&cmd.Path)
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
	cc.Arg("path", `Request path, format is /api/accounts/:accountID/users`).Required().StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}

// Run makes the HTTP request corresponding to the DeleteUserCommand command.
func (cmd *DeleteUserCommand) Run(c *client.Client) (*http.Response, error) {
	return c.DeleteUser(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteUserCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID/users/:userID`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ListUserCommand command.
func (cmd *ListUserCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ListUser(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUserCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID/users`).Required().StringVar(&cmd.Path)
}

// Run makes the HTTP request corresponding to the ShowUserCommand command.
func (cmd *ShowUserCommand) Run(c *client.Client) (*http.Response, error) {
	return c.ShowUser(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUserCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /api/accounts/:accountID/users/:userID`).Required().StringVar(&cmd.Path)
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
	cc.Arg("path", `Request path, format is /api/accounts/:accountID/users/:userID`).Required().StringVar(&cmd.Path)
	cc.Flag("payload", "Request JSON body").StringVar(&cmd.Payload)
}
