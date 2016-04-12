package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Authorize is the authorize resource media type.
var Authorize = MediaType("application/vnd.authorize+json", func() {
	Description("Token authorization response")
	Attributes(func() {
		Attribute("access_token", String, "access token")
		Attribute("expires_in", Integer, "Time to expiration in seconds")
		Attribute("token_type", String, "type of token")
		Attribute("user", User, "user")
		Attribute("tenant", Tenant, "tenant")

	})
	View("default", func() {
		Attribute("access_token")
		Attribute("expires_in")
		Attribute("token_type")
		Attribute("user")
		Attribute("tenant")
	})
})

// Login is the Login resource media type.
var Login = MediaType("application/vnd.login+json", func() {
	Description("")
	Attributes(func() {
		Attribute("email", String, "email")
		Attribute("password", String, "password")
	})
	View("default", func() {
		Attribute("email")
		Attribute("password")
	})
})

// Tenant is the tenant resource media type.
// GopherAcademy
var Tenant = MediaType("application/vnd.tenant+json", func() {
	Description("Tenant account in Congo")
	Attributes(func() {
		Attribute("name", String, "name")
	})
	View("default", func() {
		Attribute("name")
	})
})

// Series is the series resource media type.
// GopherCon
var Series = MediaType("application/vnd.series+json", func() {
	Description("A series is a recurring event, e.g. GopherCon")
	Attributes(func() {
		Attribute("name", String, "name")
	})
	View("default", func() {
		Attribute("name")
	})
})

// Event is the event resource media type.
// GopherCon 2016
var Event = MediaType("application/vnd.event+json", func() {
	Description("An event is a specific instance of a conference, e.g. GopherCon 2016")
	Attributes(func() {
		Attribute("name", String, "name")
		Attribute("location", String, "location")
		Attribute("url", String, "event URL")
		Attribute("start_date", DateTime, "start_date")
		Attribute("end_date", DateTime, "end_date")
	})
	View("default", func() {
		Attribute("name")
		Attribute("location")
		Attribute("url")
		Attribute("start_date")
		Attribute("end_date")
	})
})

// User is the user resource media type.
var User = MediaType("application/vnd.user+json", func() {
	Description("A user belonging to a member")
	Reference(UserPayload)
	Attributes(func() {
		Attribute("id", Integer, "ID of record")
		Attribute("href", String, "API href of record")
		Attribute("first_name", String, "First name of user")
		Attribute("last_name", String, "Last name of user")
		Attribute("active", Boolean, "Activation status")
		Attribute("validated", Boolean, "Validation status")
		Attribute("role", String, "User's role")
		Attribute("tenant_id", Integer)
		Attribute("email", String, "Email address of user", func() {
			Format("email")
		})
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("email")
		Attribute("role")
		Attribute("tenant_id")
	})

	View("tenant", func() {
		Attribute("id")
		Attribute("href")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("email")
		Attribute("active")
		Attribute("validated")
		Attribute("role")
		Attribute("tenant_id")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
		Attribute("email")
	})
})
