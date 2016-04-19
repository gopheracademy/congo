package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// CreateAdminuserCommand is the command line data structure for the create action of adminuser
	CreateAdminuserCommand struct {
		Payload string
	}
	// DeleteAdminuserCommand is the command line data structure for the delete action of adminuser
	DeleteAdminuserCommand struct {
	}
	// ListAdminuserCommand is the command line data structure for the list action of adminuser
	ListAdminuserCommand struct {
	}
	// ShowAdminuserCommand is the command line data structure for the show action of adminuser
	ShowAdminuserCommand struct {
	}
	// UpdateAdminuserCommand is the command line data structure for the update action of adminuser
	UpdateAdminuserCommand struct {
		Payload string
	}
	// RefreshAuthCommand is the command line data structure for the refresh action of auth
	RefreshAuthCommand struct {
	}
	// TokenAuthCommand is the command line data structure for the token action of auth
	TokenAuthCommand struct {
	}
	// CreateEventCommand is the command line data structure for the create action of event
	CreateEventCommand struct {
		Payload string
	}
	// DeleteEventCommand is the command line data structure for the delete action of event
	DeleteEventCommand struct {
	}
	// ListEventCommand is the command line data structure for the list action of event
	ListEventCommand struct {
	}
	// ShowEventCommand is the command line data structure for the show action of event
	ShowEventCommand struct {
	}
	// UpdateEventCommand is the command line data structure for the update action of event
	UpdateEventCommand struct {
		Payload string
	}
	// StatusHealthzCommand is the command line data structure for the status action of healthz
	StatusHealthzCommand struct {
	}
	// CreateSeriesCommand is the command line data structure for the create action of series
	CreateSeriesCommand struct {
		Payload string
	}
	// DeleteSeriesCommand is the command line data structure for the delete action of series
	DeleteSeriesCommand struct {
	}
	// ListSeriesCommand is the command line data structure for the list action of series
	ListSeriesCommand struct {
	}
	// ShowSeriesCommand is the command line data structure for the show action of series
	ShowSeriesCommand struct {
	}
	// UpdateSeriesCommand is the command line data structure for the update action of series
	UpdateSeriesCommand struct {
		Payload string
	}
	// CreateTenantCommand is the command line data structure for the create action of tenant
	CreateTenantCommand struct {
		Payload string
	}
	// DeleteTenantCommand is the command line data structure for the delete action of tenant
	DeleteTenantCommand struct {
	}
	// ListTenantCommand is the command line data structure for the list action of tenant
	ListTenantCommand struct {
	}
	// ShowTenantCommand is the command line data structure for the show action of tenant
	ShowTenantCommand struct {
	}
	// UpdateTenantCommand is the command line data structure for the update action of tenant
	UpdateTenantCommand struct {
		Payload string
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
	// ValidateValidateCommand is the command line data structure for the validate action of validate
	ValidateValidateCommand struct {
		V string
	}
)

// Run makes the HTTP request corresponding to the CreateAdminuserCommand command.
func (cmd *CreateAdminuserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/admin/users"
	}
	var payload app.CreateAdminuserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateAdminuser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateAdminuserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the DeleteAdminuserCommand command.
func (cmd *DeleteAdminuserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteAdminuser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteAdminuserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ListAdminuserCommand command.
func (cmd *ListAdminuserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/admin/users"
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListAdminuser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListAdminuserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ShowAdminuserCommand command.
func (cmd *ShowAdminuserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowAdminuser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowAdminuserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the UpdateAdminuserCommand command.
func (cmd *UpdateAdminuserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.UpdateAdminuserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateAdminuser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateAdminuserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the RefreshAuthCommand command.
func (cmd *RefreshAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/auth/refresh"
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RefreshAuth(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RefreshAuthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerPassword.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the TokenAuthCommand command.
func (cmd *TokenAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/auth/token"
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.TokenAuth(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *TokenAuthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerPassword.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the CreateEventCommand command.
func (cmd *CreateEventCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.CreateEventPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateEvent(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateEventCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the DeleteEventCommand command.
func (cmd *DeleteEventCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteEvent(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteEventCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ListEventCommand command.
func (cmd *ListEventCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListEvent(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListEventCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ShowEventCommand command.
func (cmd *ShowEventCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowEvent(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowEventCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UpdateEventCommand command.
func (cmd *UpdateEventCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.UpdateEventPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateEvent(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateEventCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the StatusHealthzCommand command.
func (cmd *StatusHealthzCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/healthz"
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.StatusHealthz(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *StatusHealthzCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the CreateSeriesCommand command.
func (cmd *CreateSeriesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.CreateSeriesPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateSeries(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateSeriesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the DeleteSeriesCommand command.
func (cmd *DeleteSeriesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteSeries(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteSeriesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ListSeriesCommand command.
func (cmd *ListSeriesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListSeries(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListSeriesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ShowSeriesCommand command.
func (cmd *ShowSeriesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowSeries(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowSeriesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the UpdateSeriesCommand command.
func (cmd *UpdateSeriesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.UpdateSeriesPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateSeries(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateSeriesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the CreateTenantCommand command.
func (cmd *CreateTenantCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/tenants"
	}
	var payload app.CreateTenantPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateTenant(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateTenantCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the DeleteTenantCommand command.
func (cmd *DeleteTenantCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteTenant(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteTenantCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ListTenantCommand command.
func (cmd *ListTenantCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/tenants"
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListTenant(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListTenantCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowTenantCommand command.
func (cmd *ShowTenantCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowTenant(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowTenantCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UpdateTenantCommand command.
func (cmd *UpdateTenantCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.UpdateTenantPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateTenant(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateTenantCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the CreateUserCommand command.
func (cmd *CreateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	var payload app.CreateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the DeleteUserCommand command.
func (cmd *DeleteUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ListUserCommand command.
func (cmd *ListUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ShowUserCommand command.
func (cmd *ShowUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	c.SignerJWT.RegisterFlags(cc)
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
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ValidateValidateCommand command.
func (cmd *ValidateValidateCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ValidateValidate(ctx, path, cmd.V)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ValidateValidateCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var tmp30 string
	cc.Flags().StringVar(&cmd.V, "v", tmp30, ``)
}
