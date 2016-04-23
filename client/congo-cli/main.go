package main

import (
	"fmt"
	"github.com/gopheracademy/congo/client"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// PrettyPrint is true if the tool output should be formatted for human consumption.
var PrettyPrint bool

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "congo-cli",
		Short: `CLI client for the congo service (https://congodocs.gopheracademy.com)`,
	}
	c := client.New(nil)
	c.UserAgent = "congo-cli/1.0"
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "congo.gopheracademy.com", "API hostname")
	app.PersistentFlags().DurationVarP(&c.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")
	app.PersistentFlags().BoolVar(&PrettyPrint, "pp", false, "Pretty print response body")
	RegisterCommands(app, c)
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %s", err)
		os.Exit(-1)
	}
}

// RegisterCommands all the resource action subcommands to the application command line.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "bootstrap",
		Short: `Render single page app HTML`,
	}
	tmp1 := new(BootstrapUICommand)
	sub = &cobra.Command{
		Use:   `ui "/"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "create",
		Short: `create action`,
	}
	tmp2 := new(CreateAdminuserCommand)
	sub = &cobra.Command{
		Use:   `adminuser "/api/admin/users"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp3 := new(CreateEventCommand)
	sub = &cobra.Command{
		Use:   `event "/api/tenants/:tenantID/events"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp4 := new(CreatePresentationCommand)
	sub = &cobra.Command{
		Use:   `presentation "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp5 := new(CreateSeriesCommand)
	sub = &cobra.Command{
		Use:   `series "/api/tenants/:tenantID/series"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp6 := new(CreateSpeakerCommand)
	sub = &cobra.Command{
		Use:   `speaker "/api/tenants/:tenantID/events/:eventID/speakers"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp7 := new(CreateTenantCommand)
	sub = &cobra.Command{
		Use:   `tenant "/api/tenants"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp8 := new(CreateUserCommand)
	sub = &cobra.Command{
		Use:   `user "/api/tenants/:tenantID/users"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete action`,
	}
	tmp9 := new(DeleteAdminuserCommand)
	sub = &cobra.Command{
		Use:   `adminuser "/api/admin/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp10 := new(DeleteEventCommand)
	sub = &cobra.Command{
		Use:   `event "/api/tenants/:tenantID/events/:eventID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp11 := new(DeletePresentationCommand)
	sub = &cobra.Command{
		Use:   `presentation "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations/:presentationID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp12 := new(DeleteSeriesCommand)
	sub = &cobra.Command{
		Use:   `series "/api/tenants/:tenantID/series/:seriesID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp13 := new(DeleteSpeakerCommand)
	sub = &cobra.Command{
		Use:   `speaker "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp14 := new(DeleteTenantCommand)
	sub = &cobra.Command{
		Use:   `tenant "/api/tenants/:tenantID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp15 := new(DeleteUserCommand)
	sub = &cobra.Command{
		Use:   `user "/api/tenants/:tenantID/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp16 := new(ListAdminuserCommand)
	sub = &cobra.Command{
		Use:   `adminuser "/api/admin/users"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp16.Run(c, args) },
	}
	tmp16.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp17 := new(ListEventCommand)
	sub = &cobra.Command{
		Use:   `event "/api/tenants/:tenantID/events"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp17.Run(c, args) },
	}
	tmp17.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp18 := new(ListPresentationCommand)
	sub = &cobra.Command{
		Use:   `presentation "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp18.Run(c, args) },
	}
	tmp18.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp19 := new(ListSeriesCommand)
	sub = &cobra.Command{
		Use:   `series "/api/tenants/:tenantID/series"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp19.Run(c, args) },
	}
	tmp19.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp20 := new(ListSpeakerCommand)
	sub = &cobra.Command{
		Use:   `speaker "/api/tenants/:tenantID/events/:eventID/speakers"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp20.Run(c, args) },
	}
	tmp20.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp21 := new(ListTenantCommand)
	sub = &cobra.Command{
		Use:   `tenant "/api/tenants"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp21.Run(c, args) },
	}
	tmp21.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp22 := new(ListUserCommand)
	sub = &cobra.Command{
		Use:   `user "/api/tenants/:tenantID/users"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp22.Run(c, args) },
	}
	tmp22.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "refresh",
		Short: `Obtain a refreshed access token`,
	}
	tmp23 := new(RefreshAuthCommand)
	sub = &cobra.Command{
		Use:   `auth "/api/auth/refresh"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp23.Run(c, args) },
	}
	tmp23.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp24 := new(ShowAdminuserCommand)
	sub = &cobra.Command{
		Use:   `adminuser "/api/admin/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp24.Run(c, args) },
	}
	tmp24.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp25 := new(ShowEventCommand)
	sub = &cobra.Command{
		Use:   `event "/api/tenants/:tenantID/events/:eventID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp25.Run(c, args) },
	}
	tmp25.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp26 := new(ShowPresentationCommand)
	sub = &cobra.Command{
		Use:   `presentation "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations/:presentationID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp26.Run(c, args) },
	}
	tmp26.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp27 := new(ShowSeriesCommand)
	sub = &cobra.Command{
		Use:   `series "/api/tenants/:tenantID/series/:seriesID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp27.Run(c, args) },
	}
	tmp27.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp28 := new(ShowSpeakerCommand)
	sub = &cobra.Command{
		Use:   `speaker "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp28.Run(c, args) },
	}
	tmp28.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp29 := new(ShowTenantCommand)
	sub = &cobra.Command{
		Use:   `tenant "/api/tenants/:tenantID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp29.Run(c, args) },
	}
	tmp29.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp30 := new(ShowUserCommand)
	sub = &cobra.Command{
		Use:   `user "/api/tenants/:tenantID/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp30.Run(c, args) },
	}
	tmp30.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "status",
		Short: `Get Server Status`,
	}
	tmp31 := new(StatusHealthzCommand)
	sub = &cobra.Command{
		Use:   `healthz "/api/healthz"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp31.Run(c, args) },
	}
	tmp31.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "token",
		Short: `Obtain an access token`,
	}
	tmp32 := new(TokenAuthCommand)
	sub = &cobra.Command{
		Use:   `auth "/api/auth/token"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp32.Run(c, args) },
	}
	tmp32.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update action`,
	}
	tmp33 := new(UpdateAdminuserCommand)
	sub = &cobra.Command{
		Use:   `adminuser "/api/admin/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp33.Run(c, args) },
	}
	tmp33.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp34 := new(UpdateEventCommand)
	sub = &cobra.Command{
		Use:   `event "/api/tenants/:tenantID/events/:eventID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp34.Run(c, args) },
	}
	tmp34.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp35 := new(UpdatePresentationCommand)
	sub = &cobra.Command{
		Use:   `presentation "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations/:presentationID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp35.Run(c, args) },
	}
	tmp35.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp36 := new(UpdateSeriesCommand)
	sub = &cobra.Command{
		Use:   `series "/api/tenants/:tenantID/series/:seriesID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp36.Run(c, args) },
	}
	tmp36.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp37 := new(UpdateSpeakerCommand)
	sub = &cobra.Command{
		Use:   `speaker "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp37.Run(c, args) },
	}
	tmp37.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp38 := new(UpdateTenantCommand)
	sub = &cobra.Command{
		Use:   `tenant "/api/tenants/:tenantID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp38.Run(c, args) },
	}
	tmp38.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp39 := new(UpdateUserCommand)
	sub = &cobra.Command{
		Use:   `user "/api/tenants/:tenantID/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp39.Run(c, args) },
	}
	tmp39.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "validate",
		Short: `validate user email`,
	}
	tmp40 := new(ValidateValidateCommand)
	sub = &cobra.Command{
		Use:   `validate "/api/validate/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp40.Run(c, args) },
	}
	tmp40.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)

}
