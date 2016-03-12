package main

import (
	"encoding/json"
	"fmt"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/client"
	"github.com/spf13/cobra"
)

type (
	// CallbackAuthCommand is the command line data structure for the callback action of auth
	CallbackAuthCommand struct {
	}

	// OauthAuthCommand is the command line data structure for the oauth action of auth
	OauthAuthCommand struct {
	}

	// RefreshAuthCommand is the command line data structure for the refresh action of auth
	RefreshAuthCommand struct {
		Payload string
	}

	// TokenAuthCommand is the command line data structure for the token action of auth
	TokenAuthCommand struct {
		Payload string
	}

	// CreateProposalCommand is the command line data structure for the create action of proposal
	CreateProposalCommand struct {
		Payload string
	}

	// DeleteProposalCommand is the command line data structure for the delete action of proposal
	DeleteProposalCommand struct {
	}

	// ListProposalCommand is the command line data structure for the list action of proposal
	ListProposalCommand struct {
	}

	// ShowProposalCommand is the command line data structure for the show action of proposal
	ShowProposalCommand struct {
	}

	// UpdateProposalCommand is the command line data structure for the update action of proposal
	UpdateProposalCommand struct {
		Payload string
	}

	// CreateReviewCommand is the command line data structure for the create action of review
	CreateReviewCommand struct {
		Payload string
	}

	// DeleteReviewCommand is the command line data structure for the delete action of review
	DeleteReviewCommand struct {
	}

	// ListReviewCommand is the command line data structure for the list action of review
	ListReviewCommand struct {
	}

	// ShowReviewCommand is the command line data structure for the show action of review
	ShowReviewCommand struct {
	}

	// UpdateReviewCommand is the command line data structure for the update action of review
	UpdateReviewCommand struct {
		Payload string
	}

	// BootstrapUICommand is the command line data structure for the bootstrap action of ui
	BootstrapUICommand struct {
	}

	// CreateUserCommand is the command line data structure for the create action of user
	CreateUserCommand struct {
		Payload string
	}

	// DeleteUserCommand is the command line data structure for the delete action of user
	DeleteUserCommand struct {
	}

	// ListUserCommand is the command line data structure for the list action of user
	ListUserCommand struct {
	}

	// ShowUserCommand is the command line data structure for the show action of user
	ShowUserCommand struct {
	}

	// UpdateUserCommand is the command line data structure for the update action of user
	UpdateUserCommand struct {
		Payload string
	}
)

// Run makes the HTTP request corresponding to the CallbackAuthCommand command.
func (cmd *CallbackAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.CallbackAuth(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CallbackAuthCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the OauthAuthCommand command.
func (cmd *OauthAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.OauthAuth(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *OauthAuthCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the RefreshAuthCommand command.
func (cmd *RefreshAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/auth/refresh"
	}
	var payload app.RefreshAuthPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.RefreshAuth(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RefreshAuthCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the TokenAuthCommand command.
func (cmd *TokenAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/auth/token"
	}
	var payload app.TokenAuthPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.TokenAuth(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *TokenAuthCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the CreateProposalCommand command.
func (cmd *CreateProposalCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.CreateProposalPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.CreateProposal(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateProposalCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the DeleteProposalCommand command.
func (cmd *DeleteProposalCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.DeleteProposal(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteProposalCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ListProposalCommand command.
func (cmd *ListProposalCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ListProposal(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListProposalCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ShowProposalCommand command.
func (cmd *ShowProposalCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ShowProposal(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowProposalCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the UpdateProposalCommand command.
func (cmd *UpdateProposalCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.UpdateProposalPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.UpdateProposal(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateProposalCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the CreateReviewCommand command.
func (cmd *CreateReviewCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.CreateReviewPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.CreateReview(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateReviewCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the DeleteReviewCommand command.
func (cmd *DeleteReviewCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.DeleteReview(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteReviewCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ListReviewCommand command.
func (cmd *ListReviewCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ListReview(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListReviewCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ShowReviewCommand command.
func (cmd *ShowReviewCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ShowReview(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowReviewCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the UpdateReviewCommand command.
func (cmd *UpdateReviewCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.UpdateReviewPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.UpdateReview(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateReviewCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the BootstrapUICommand command.
func (cmd *BootstrapUICommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/"
	}
	resp, err := c.BootstrapUI(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *BootstrapUICommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the CreateUserCommand command.
func (cmd *CreateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/users"
	}
	var payload app.CreateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.CreateUser(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateUserCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the DeleteUserCommand command.
func (cmd *DeleteUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.DeleteUser(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteUserCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ListUserCommand command.
func (cmd *ListUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/users"
	}
	resp, err := c.ListUser(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUserCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ShowUserCommand command.
func (cmd *ShowUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.ShowUser(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUserCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the UpdateUserCommand command.
func (cmd *UpdateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.UpdateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.UpdateUser(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateUserCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}
