package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gopheracademy/congo/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

// PrettyPrint is true if the tool output should be formatted for human consumption.
var PrettyPrint bool

func main() {
	// Create command line parser
	app := kingpin.New("congo-cli", "CLI client for the congo service (https://gopheracademy.com/congo/getting-started.html)")
	c := client.New()
	c.UserAgent = "congo-cli/1.0"
	app.Flag("scheme", "Set the requests scheme").Short('s').Default("http").StringVar(&c.Scheme)
	app.Flag("host", "API hostname").Short('h').Default("api.gopheracademy.com").StringVar(&c.Host)
	app.Flag("timeout", "Set the request timeout, defaults to 20s").Short('t').Default("20s").DurationVar(&c.Timeout)
	app.Flag("dump", "Dump HTTP request and response.").BoolVar(&c.Dump)
	app.Flag("pp", "Pretty print response body").BoolVar(&PrettyPrint)
	commands := RegisterCommands(app)
	// Make "client-cli <action> [<resource>] --help" equivalent to
	// "client-cli help <action> [<resource>]"
	if os.Args[len(os.Args)-1] == "--help" {
		args := append([]string{os.Args[0], "help"}, os.Args[1:len(os.Args)-1]...)
		os.Args = args
	}
	cmdName, err := app.Parse(os.Args[1:])
	if err != nil {
		kingpin.Fatalf(err.Error())
	}
	cmd, ok := commands[cmdName]
	if !ok {
		kingpin.Fatalf("unknown command %s", cmdName)
	}
	resp, err := cmd.Run(c)
	if err != nil {
		kingpin.Fatalf("request failed: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		kingpin.Fatalf("failed to read body: %s", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Let user know if something went wrong
		var sbody string
		if len(body) > 0 {
			sbody = ": " + string(body)
		}
		fmt.Printf("error: %d%s", resp.StatusCode, sbody)
	} else if !c.Dump && len(body) > 0 {
		var out string
		if PrettyPrint {
			var jbody interface{}
			err = json.Unmarshal(body, &jbody)
			if err != nil {
				out = string(body)
			} else {
				var b []byte
				b, err = json.MarshalIndent(jbody, "", "    ")
				if err == nil {
					out = string(b)
				} else {
					out = string(body)
				}
			}
		} else {
			out = string(body)
		}
		fmt.Print(out)
	}

	// Figure out exit code
	exitStatus := 0
	switch {
	case resp.StatusCode == 401:
		exitStatus = 1
	case resp.StatusCode == 403:
		exitStatus = 3
	case resp.StatusCode == 404:
		exitStatus = 4
	case resp.StatusCode > 399 && resp.StatusCode < 500:
		exitStatus = 2
	case resp.StatusCode > 499:
		exitStatus = 5
	}
	os.Exit(exitStatus)
}

// RegisterCommands all the resource action subcommands to the application command line.
func RegisterCommands(app *kingpin.Application) map[string]client.ActionCommand {
	res := make(map[string]client.ActionCommand)
	var command, sub *kingpin.CmdClause
	command = app.Command("create", "create action")
	tmp1 := new(CreateAccountCommand)
	sub = command.Command("account", "Create new account")
	tmp1.RegisterFlags(sub)
	res["create account"] = tmp1
	tmp2 := new(CreateInstanceCommand)
	sub = command.Command("instance", "Record new instance")
	tmp2.RegisterFlags(sub)
	res["create instance"] = tmp2
	tmp3 := new(CreateSeriesCommand)
	sub = command.Command("series", "Record new series")
	tmp3.RegisterFlags(sub)
	res["create series"] = tmp3
	tmp4 := new(CreateUserCommand)
	sub = command.Command("user", "Record new user")
	tmp4.RegisterFlags(sub)
	res["create user"] = tmp4
	command = app.Command("delete", "delete action")
	tmp5 := new(DeleteAccountCommand)
	sub = command.Command("account", "")
	tmp5.RegisterFlags(sub)
	res["delete account"] = tmp5
	tmp6 := new(DeleteInstanceCommand)
	sub = command.Command("instance", "")
	tmp6.RegisterFlags(sub)
	res["delete instance"] = tmp6
	tmp7 := new(DeleteSeriesCommand)
	sub = command.Command("series", "")
	tmp7.RegisterFlags(sub)
	res["delete series"] = tmp7
	tmp8 := new(DeleteUserCommand)
	sub = command.Command("user", "")
	tmp8.RegisterFlags(sub)
	res["delete user"] = tmp8
	command = app.Command("list", "list action")
	tmp9 := new(ListAccountCommand)
	sub = command.Command("account", "List all accounts")
	tmp9.RegisterFlags(sub)
	res["list account"] = tmp9
	tmp10 := new(ListInstanceCommand)
	sub = command.Command("instance", "List all instances of a series")
	tmp10.RegisterFlags(sub)
	res["list instance"] = tmp10
	tmp11 := new(ListSeriesCommand)
	sub = command.Command("series", "List all series in account")
	tmp11.RegisterFlags(sub)
	res["list series"] = tmp11
	tmp12 := new(ListUserCommand)
	sub = command.Command("user", "List all users in account")
	tmp12.RegisterFlags(sub)
	res["list user"] = tmp12
	command = app.Command("show", "show action")
	tmp13 := new(ShowAccountCommand)
	sub = command.Command("account", "Retrieve account with given id")
	tmp13.RegisterFlags(sub)
	res["show account"] = tmp13
	tmp14 := new(ShowInstanceCommand)
	sub = command.Command("instance", "Retrieve instance with given id")
	tmp14.RegisterFlags(sub)
	res["show instance"] = tmp14
	tmp15 := new(ShowSeriesCommand)
	sub = command.Command("series", "Retrieve series with given id")
	tmp15.RegisterFlags(sub)
	res["show series"] = tmp15
	tmp16 := new(ShowUserCommand)
	sub = command.Command("user", "Retrieve user with given id")
	tmp16.RegisterFlags(sub)
	res["show user"] = tmp16
	command = app.Command("update", "update action")
	tmp17 := new(UpdateAccountCommand)
	sub = command.Command("account", "Change account name")
	tmp17.RegisterFlags(sub)
	res["update account"] = tmp17
	tmp18 := new(UpdateInstanceCommand)
	sub = command.Command("instance", "")
	tmp18.RegisterFlags(sub)
	res["update instance"] = tmp18
	tmp19 := new(UpdateSeriesCommand)
	sub = command.Command("series", "")
	tmp19.RegisterFlags(sub)
	res["update series"] = tmp19
	tmp20 := new(UpdateUserCommand)
	sub = command.Command("user", "")
	tmp20.RegisterFlags(sub)
	res["update user"] = tmp20

	return res
}
