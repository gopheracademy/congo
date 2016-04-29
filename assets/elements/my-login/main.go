package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"code.palmstonegames.com/polymer"
	"github.com/go-humble/locstor"
	"github.com/gopherjs/gopherjs/js"
)

func init() {
	polymer.Register("my-login", &Login{})
}

type Login struct {
	*polymer.Proto
	DoLogin  chan *polymer.Event `polymer:"handler"`
	Username string              `polymer:"bind"`
	Password string              `polymer:"bind"`
	Message  string              `polymer:"bind"`
}

func (l *Login) Created() {
	l.Username = "user@example.com"
}

func (l *Login) Ready() {
	l.ListenToggleEvents()
}

func (l *Login) ListenToggleEvents() {
	go func() {
		for {
			select {
			case <-l.DoLogin:
				success := login(l.Username, l.Password)
				fmt.Println(success)
				if success {
					js.Global.Get("page").Call("redirect", "/")
				} else {
					l.Message = "Login Failed"
					l.Notify("message")
				}
			}

		}
	}()
}
func login(u, p string) bool {

	req, err := http.NewRequest("POST", "/api/auth/token", nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(u, p))
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("did not get acceptable status code: %v body: %q", resp.Status, string(body))
		return false
	}
	var a Authorize
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil {
		return false
	}

	if err := locstor.SetItem("user", a.User.String()); err != nil {
		fmt.Println("can't set user in local storage")
	}

	if err := locstor.SetItem("tenant", a.Tenant.String()); err != nil {
		fmt.Println("cant set tenant in local storage")
	}

	if err := locstor.SetItem("token", a.AccessToken); err != nil {
		fmt.Println("cant set token in local storage")
	}
	return true
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {}

type Authorize struct {
	// access token
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// Time to expiration in seconds
	ExpiresIn int `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// tenant
	Tenant Tenant `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// type of token
	TokenType string `json:"token_type,omitempty" xml:"token_type,omitempty"`
	// user
	User User `json:"user,omitempty" xml:"user,omitempty"`
}

type User struct {
	// Email address of user
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// First name of user
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// API href of record
	Href string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of record
	ID int `json:"id,omitempty" xml:"id,omitempty"`
	// Last name of user
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// User's role
	Role     string `json:"role,omitempty" xml:"role,omitempty"`
	TenantID int    `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

func (u User) String() string {

	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}

type Tenant struct {
	// ID
	ID int `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

func (t Tenant) String() string {

	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
