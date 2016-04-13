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
		Short: `CLI client for the congo service (https://congo.com)`,
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
		Use:   "create",
		Short: `create action`,
	}
	tmp1 := new(CreateAdminuserCommand)
	sub = &cobra.Command{
		Use:   `adminuser "/api/admin/users"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp2 := new(CreateEventCommand)
	sub = &cobra.Command{
		Use:   `event "/api/tenants/:tenantID/events"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp3 := new(CreateSeriesCommand)
	sub = &cobra.Command{
		Use:   `series "/api/tenants/:tenantID/series"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp4 := new(CreateTenantCommand)
	sub = &cobra.Command{
		Use:   `tenant "/api/tenants"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp5 := new(CreateUserCommand)
	sub = &cobra.Command{
		Use:   `user "/api/tenants/:tenantID/users"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete action`,
	}
	tmp6 := new(DeleteAdminuserCommand)
	sub = &cobra.Command{
		Use:   `adminuser "/api/admin/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp7 := new(DeleteEventCommand)
	sub = &cobra.Command{
		Use:   `event "/api/tenants/:tenantID/events/:eventID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp8 := new(DeleteSeriesCommand)
	sub = &cobra.Command{
		Use:   `series "/api/tenants/:tenantID/series/:seriesID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp9 := new(DeleteTenantCommand)
	sub = &cobra.Command{
		Use:   `tenant "/api/tenants/:tenantID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp10 := new(DeleteUserCommand)
	sub = &cobra.Command{
		Use:   `user "/api/tenants/:tenantID/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp11 := new(ListAdminuserCommand)
	sub = &cobra.Command{
		Use:   `adminuser "/api/admin/users"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp12 := new(ListEventCommand)
	sub = &cobra.Command{
		Use:   `event "/api/tenants/:tenantID/events"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp13 := new(ListSeriesCommand)
	sub = &cobra.Command{
		Use:   `series "/api/tenants/:tenantID/series"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp14 := new(ListTenantCommand)
	sub = &cobra.Command{
		Use:   `tenant "/api/tenants"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp15 := new(ListUserCommand)
	sub = &cobra.Command{
		Use:   `user "/api/tenants/:tenantID/users"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "refresh",
		Short: `Obtain a refreshed access token`,
	}
	tmp16 := new(RefreshAuthCommand)
	sub = &cobra.Command{
		Use:   `auth "/api/auth/refresh"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp16.Run(c, args) },
	}
	tmp16.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp17 := new(ShowAdminuserCommand)
	sub = &cobra.Command{
		Use:   `adminuser "/api/admin/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp17.Run(c, args) },
	}
	tmp17.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp18 := new(ShowEventCommand)
	sub = &cobra.Command{
		Use:   `event "/api/tenants/:tenantID/events/:eventID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp18.Run(c, args) },
	}
	tmp18.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp19 := new(ShowSeriesCommand)
	sub = &cobra.Command{
		Use:   `series "/api/tenants/:tenantID/series/:seriesID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp19.Run(c, args) },
	}
	tmp19.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp20 := new(ShowTenantCommand)
	sub = &cobra.Command{
		Use:   `tenant "/api/tenants/:tenantID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp20.Run(c, args) },
	}
	tmp20.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp21 := new(ShowUserCommand)
	sub = &cobra.Command{
		Use:   `user "/api/tenants/:tenantID/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp21.Run(c, args) },
	}
	tmp21.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "status",
		Short: `Get Server Status`,
	}
	tmp22 := new(StatusHealthzCommand)
	sub = &cobra.Command{
		Use:   `healthz "/api/healthz"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp22.Run(c, args) },
	}
	tmp22.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "token",
		Short: `Obtain an access token`,
	}
	tmp23 := new(TokenAuthCommand)
	sub = &cobra.Command{
		Use:   `auth "/api/auth/token"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp23.Run(c, args) },
	}
	tmp23.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update action`,
	}
	tmp24 := new(UpdateAdminuserCommand)
	sub = &cobra.Command{
		Use:   `adminuser "/api/admin/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp24.Run(c, args) },
	}
	tmp24.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp25 := new(UpdateEventCommand)
	sub = &cobra.Command{
		Use:   `event "/api/tenants/:tenantID/events/:eventID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp25.Run(c, args) },
	}
	tmp25.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp26 := new(UpdateSeriesCommand)
	sub = &cobra.Command{
		Use:   `series "/api/tenants/:tenantID/series/:seriesID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp26.Run(c, args) },
	}
	tmp26.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp27 := new(UpdateTenantCommand)
	sub = &cobra.Command{
		Use:   `tenant "/api/tenants/:tenantID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp27.Run(c, args) },
	}
	tmp27.RegisterFlags(sub)
	command.AddCommand(sub)
	tmp28 := new(UpdateUserCommand)
	sub = &cobra.Command{
		Use:   `user "/api/tenants/:tenantID/users/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp28.Run(c, args) },
	}
	tmp28.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "validate",
		Short: `validate user email`,
	}
	tmp29 := new(ValidateValidateCommand)
	sub = &cobra.Command{
		Use:   `validate "/api/validate/:userID"`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp29.Run(c, args) },
	}
	tmp29.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)

}
