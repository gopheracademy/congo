package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

// Account is the account resource media type.
var Account = MediaType("application/vnd.account+json", func() {
	Description("A tenant account")
	Reference(AccountModel)
	Attributes(func() {
		Attribute("id", Integer, "ID of account")
		Attribute("href", String, "API href of account")
		Attribute("name", String, "Name of account")
	})
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})
	View("link", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})
})

// User is the user resource media type.
var User = MediaType("application/vnd.user+json", func() {
	Description("A user belonging to a tenant account")
	Metadata("mediatype", "123")
	Reference(UserModel)
	Attributes(func() {
		Metadata("test", "123")
		Attribute("id", Integer, "ID of user")
		Attribute("href", String, "API href of user")
		Attribute("first_name", String, "First name of user")
		Attribute("last_name", String, "Last name of user")
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
	})
	View("link", func() {
		Attribute("id")
		Attribute("href")
		Attribute("email")
	})
})

// Series represents a recurring conference (GopherCon)
var Series = MediaType("application/vnd.series+json", func() {
	Description("A recurring event or conference")
	Reference(SeriesModel)
	Attributes(func() {
		Attribute("id", Integer, "ID of series")
		Attribute("href", String, "API href of series")
		Attribute("account", Account, "Account that owns bottle")
		// Attributes below inherit from the base type
		Attribute("name")

		View("default", func() {
			Attribute("id")
			Attribute("href")
			Attribute("name")
			Attribute("links")
		})
		View("tiny", func() {
			Attribute("id")
			Attribute("href")
			Attribute("name")
			Attribute("links")
		})
		View("full", func() {
			Attribute("id")
			Attribute("href")
			Attribute("name")
			Attribute("account", func() {
				View("full")
			})
			Attribute("links")
		})
		View("link", func() {
			Attribute("id")
			Attribute("href")
			Attribute("name")
		})
	})
})

// Instance represents a specific occurence of a series (GopherCon 2016)
var Instance = MediaType("application/vnd.instance+json", func() {
	Description("An instance of an event or conference")
	Reference(InstanceModel)
	Attributes(func() {
		Attribute("id", Integer, "ID of Instance")
		Attribute("href", String, "API href of instance")
		Attribute("series", Series, "Series that this instance belongs to")
		// Attributes below inherit from the base type
		Attribute("name")

		View("default", func() {
			Attribute("id")
			Attribute("href")
			Attribute("name")
			Attribute("links")
		})
		View("tiny", func() {
			Attribute("id")
			Attribute("href")
			Attribute("name")
			Attribute("links")
		})
		View("full", func() {
			Attribute("id")
			Attribute("href")
			Attribute("name")
			Attribute("series", func() {
				View("full")
			})
			Attribute("links")
		})
	})
})
