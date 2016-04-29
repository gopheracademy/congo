package main

import (
	"encoding/json"
	"fmt"
)

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
