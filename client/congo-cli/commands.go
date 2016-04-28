package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
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
		// UserID
		UserID int
	}
	// ListAdminuserCommand is the command line data structure for the list action of adminuser
	ListAdminuserCommand struct {
	}
	// ShowAdminuserCommand is the command line data structure for the show action of adminuser
	ShowAdminuserCommand struct {
		UserID int
	}
	// UpdateAdminuserCommand is the command line data structure for the update action of adminuser
	UpdateAdminuserCommand struct {
		Payload string
		UserID  int
	}
	// RefreshAuthCommand is the command line data structure for the refresh action of auth
	RefreshAuthCommand struct {
	}
	// TokenAuthCommand is the command line data structure for the token action of auth
	TokenAuthCommand struct {
	}
	// CreateEventCommand is the command line data structure for the create action of event
	CreateEventCommand struct {
		Payload  string
		TenantID int
	}
	// DeleteEventCommand is the command line data structure for the delete action of event
	DeleteEventCommand struct {
		// Event ID
		EventID  int
		TenantID int
	}
	// ListEventCommand is the command line data structure for the list action of event
	ListEventCommand struct {
		TenantID int
	}
	// ShowEventCommand is the command line data structure for the show action of event
	ShowEventCommand struct {
		EventID  int
		TenantID int
	}
	// UpdateEventCommand is the command line data structure for the update action of event
	UpdateEventCommand struct {
		Payload  string
		EventID  int
		TenantID int
	}
	// StatusHealthzCommand is the command line data structure for the status action of healthz
	StatusHealthzCommand struct {
	}
	// CreatePresentationCommand is the command line data structure for the create action of presentation
	CreatePresentationCommand struct {
		Payload   string
		EventID   int
		SpeakerID int
		TenantID  int
	}
	// DeletePresentationCommand is the command line data structure for the delete action of presentation
	DeletePresentationCommand struct {
		EventID int
		// Presentation ID
		PresentationID int
		SpeakerID      int
		TenantID       int
	}
	// ListPresentationCommand is the command line data structure for the list action of presentation
	ListPresentationCommand struct {
		EventID   int
		SpeakerID int
		TenantID  int
	}
	// ShowPresentationCommand is the command line data structure for the show action of presentation
	ShowPresentationCommand struct {
		EventID        int
		PresentationID int
		SpeakerID      int
		TenantID       int
	}
	// UpdatePresentationCommand is the command line data structure for the update action of presentation
	UpdatePresentationCommand struct {
		Payload        string
		EventID        int
		PresentationID int
		SpeakerID      int
		TenantID       int
	}
	// CreateSeriesCommand is the command line data structure for the create action of series
	CreateSeriesCommand struct {
		Payload  string
		TenantID int
	}
	// DeleteSeriesCommand is the command line data structure for the delete action of series
	DeleteSeriesCommand struct {
		// Series ID
		SeriesID int
		TenantID int
	}
	// ListSeriesCommand is the command line data structure for the list action of series
	ListSeriesCommand struct {
		TenantID int
	}
	// ShowSeriesCommand is the command line data structure for the show action of series
	ShowSeriesCommand struct {
		SeriesID int
		TenantID int
	}
	// UpdateSeriesCommand is the command line data structure for the update action of series
	UpdateSeriesCommand struct {
		Payload  string
		SeriesID int
		TenantID int
	}
	// CreateSpeakerCommand is the command line data structure for the create action of speaker
	CreateSpeakerCommand struct {
		Payload  string
		EventID  int
		TenantID int
	}
	// DeleteSpeakerCommand is the command line data structure for the delete action of speaker
	DeleteSpeakerCommand struct {
		EventID int
		// Speaker ID
		SpeakerID int
		TenantID  int
	}
	// ListSpeakerCommand is the command line data structure for the list action of speaker
	ListSpeakerCommand struct {
		EventID  int
		TenantID int
	}
	// ShowSpeakerCommand is the command line data structure for the show action of speaker
	ShowSpeakerCommand struct {
		EventID   int
		SpeakerID int
		TenantID  int
	}
	// UpdateSpeakerCommand is the command line data structure for the update action of speaker
	UpdateSpeakerCommand struct {
		Payload   string
		EventID   int
		SpeakerID int
		TenantID  int
	}
	// CreateTenantCommand is the command line data structure for the create action of tenant
	CreateTenantCommand struct {
		Payload string
	}
	// DeleteTenantCommand is the command line data structure for the delete action of tenant
	DeleteTenantCommand struct {
		// Tenant ID
		TenantID int
	}
	// ListTenantCommand is the command line data structure for the list action of tenant
	ListTenantCommand struct {
	}
	// ShowTenantCommand is the command line data structure for the show action of tenant
	ShowTenantCommand struct {
		TenantID int
	}
	// UpdateTenantCommand is the command line data structure for the update action of tenant
	UpdateTenantCommand struct {
		Payload  string
		TenantID int
	}
	// CreateUserCommand is the command line data structure for the create action of user
	CreateUserCommand struct {
		Payload  string
		TenantID int
	}
	// DeleteUserCommand is the command line data structure for the delete action of user
	DeleteUserCommand struct {
		TenantID int
		// UserID
		UserID int
	}
	// ListUserCommand is the command line data structure for the list action of user
	ListUserCommand struct {
		TenantID int
	}
	// ShowUserCommand is the command line data structure for the show action of user
	ShowUserCommand struct {
		TenantID int
		UserID   int
	}
	// UpdateUserCommand is the command line data structure for the update action of user
	UpdateUserCommand struct {
		Payload  string
		TenantID int
		UserID   int
	}
	// ValidateValidateCommand is the command line data structure for the validate action of validate
	ValidateValidateCommand struct {
		UserID string
		V      string
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
	var payload client.CreateAdminuserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
		path = fmt.Sprintf("/api/admin/users/%v", cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, `UserID `)
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
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
		path = fmt.Sprintf("/api/admin/users/%v", cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the UpdateAdminuserCommand command.
func (cmd *UpdateAdminuserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/admin/users/%v", cmd.UserID)
	}
	var payload client.UpdateAdminuserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
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
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
		path = fmt.Sprintf("/api/tenants/%v/events", cmd.TenantID)
	}
	var payload client.CreateEventPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the DeleteEventCommand command.
func (cmd *DeleteEventCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v", cmd.EventID, cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, `Event ID`)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ListEventCommand command.
func (cmd *ListEventCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events", cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ShowEventCommand command.
func (cmd *ShowEventCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v", cmd.EventID, cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
}

// Run makes the HTTP request corresponding to the UpdateEventCommand command.
func (cmd *UpdateEventCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v", cmd.EventID, cmd.TenantID)
	}
	var payload client.UpdateEventPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
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
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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

// Run makes the HTTP request corresponding to the CreatePresentationCommand command.
func (cmd *CreatePresentationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations", cmd.EventID, cmd.SpeakerID, cmd.TenantID)
	}
	var payload client.CreatePresentationPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreatePresentation(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreatePresentationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var speakerID int
	cc.Flags().IntVar(&cmd.SpeakerID, "speakerID", speakerID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the DeletePresentationCommand command.
func (cmd *DeletePresentationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", cmd.EventID, cmd.PresentationID, cmd.SpeakerID, cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeletePresentation(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeletePresentationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var presentationID int
	cc.Flags().IntVar(&cmd.PresentationID, "presentationID", presentationID, `Presentation ID`)
	var speakerID int
	cc.Flags().IntVar(&cmd.SpeakerID, "speakerID", speakerID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ListPresentationCommand command.
func (cmd *ListPresentationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations", cmd.EventID, cmd.SpeakerID, cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListPresentation(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListPresentationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var speakerID int
	cc.Flags().IntVar(&cmd.SpeakerID, "speakerID", speakerID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ShowPresentationCommand command.
func (cmd *ShowPresentationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", cmd.EventID, cmd.PresentationID, cmd.SpeakerID, cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowPresentation(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowPresentationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var presentationID int
	cc.Flags().IntVar(&cmd.PresentationID, "presentationID", presentationID, ``)
	var speakerID int
	cc.Flags().IntVar(&cmd.SpeakerID, "speakerID", speakerID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
}

// Run makes the HTTP request corresponding to the UpdatePresentationCommand command.
func (cmd *UpdatePresentationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", cmd.EventID, cmd.PresentationID, cmd.SpeakerID, cmd.TenantID)
	}
	var payload client.UpdatePresentationPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdatePresentation(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdatePresentationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var presentationID int
	cc.Flags().IntVar(&cmd.PresentationID, "presentationID", presentationID, ``)
	var speakerID int
	cc.Flags().IntVar(&cmd.SpeakerID, "speakerID", speakerID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the CreateSeriesCommand command.
func (cmd *CreateSeriesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/series", cmd.TenantID)
	}
	var payload client.CreateSeriesPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the DeleteSeriesCommand command.
func (cmd *DeleteSeriesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/series/%v", cmd.SeriesID, cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var seriesID int
	cc.Flags().IntVar(&cmd.SeriesID, "seriesID", seriesID, `Series ID`)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ListSeriesCommand command.
func (cmd *ListSeriesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/series", cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ShowSeriesCommand command.
func (cmd *ShowSeriesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/series/%v", cmd.SeriesID, cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var seriesID int
	cc.Flags().IntVar(&cmd.SeriesID, "seriesID", seriesID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the UpdateSeriesCommand command.
func (cmd *UpdateSeriesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/series/%v", cmd.SeriesID, cmd.TenantID)
	}
	var payload client.UpdateSeriesPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var seriesID int
	cc.Flags().IntVar(&cmd.SeriesID, "seriesID", seriesID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the CreateSpeakerCommand command.
func (cmd *CreateSpeakerCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v/speakers", cmd.EventID, cmd.TenantID)
	}
	var payload client.CreateSpeakerPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateSpeaker(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateSpeakerCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
}

// Run makes the HTTP request corresponding to the DeleteSpeakerCommand command.
func (cmd *DeleteSpeakerCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", cmd.EventID, cmd.SpeakerID, cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteSpeaker(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteSpeakerCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var speakerID int
	cc.Flags().IntVar(&cmd.SpeakerID, "speakerID", speakerID, `Speaker ID`)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
}

// Run makes the HTTP request corresponding to the ListSpeakerCommand command.
func (cmd *ListSpeakerCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v/speakers", cmd.EventID, cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListSpeaker(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListSpeakerCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
}

// Run makes the HTTP request corresponding to the ShowSpeakerCommand command.
func (cmd *ShowSpeakerCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", cmd.EventID, cmd.SpeakerID, cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowSpeaker(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowSpeakerCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var speakerID int
	cc.Flags().IntVar(&cmd.SpeakerID, "speakerID", speakerID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
}

// Run makes the HTTP request corresponding to the UpdateSpeakerCommand command.
func (cmd *UpdateSpeakerCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", cmd.EventID, cmd.SpeakerID, cmd.TenantID)
	}
	var payload client.UpdateSpeakerPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateSpeaker(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateSpeakerCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, ``)
	var speakerID int
	cc.Flags().IntVar(&cmd.SpeakerID, "speakerID", speakerID, ``)
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
}

// Run makes the HTTP request corresponding to the CreateTenantCommand command.
func (cmd *CreateTenantCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/tenants"
	}
	var payload client.CreateTenantPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
		path = fmt.Sprintf("/api/tenants/%v", cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, `Tenant ID`)
}

// Run makes the HTTP request corresponding to the ListTenantCommand command.
func (cmd *ListTenantCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/tenants"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
		path = fmt.Sprintf("/api/tenants/%v", cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
}

// Run makes the HTTP request corresponding to the UpdateTenantCommand command.
func (cmd *UpdateTenantCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v", cmd.TenantID)
	}
	var payload client.UpdateTenantPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
}

// Run makes the HTTP request corresponding to the CreateUserCommand command.
func (cmd *CreateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/users", cmd.TenantID)
	}
	var payload client.CreateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the DeleteUserCommand command.
func (cmd *DeleteUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/users/%v", cmd.TenantID, cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, `UserID `)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ListUserCommand command.
func (cmd *ListUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/users", cmd.TenantID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ShowUserCommand command.
func (cmd *ShowUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/users/%v", cmd.TenantID, cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the UpdateUserCommand command.
func (cmd *UpdateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/tenants/%v/users/%v", cmd.TenantID, cmd.UserID)
	}
	var payload client.UpdateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var tenantID int
	cc.Flags().IntVar(&cmd.TenantID, "tenantID", tenantID, ``)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
	c.SignerJWT.RegisterFlags(cc)
}

// Run makes the HTTP request corresponding to the ValidateValidateCommand command.
func (cmd *ValidateValidateCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/validate/%v", cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
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
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "userID", userID, ``)
	var v string
	cc.Flags().StringVar(&cmd.V, "v", v, ``)
}
