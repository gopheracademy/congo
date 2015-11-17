package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

// This is the api definition used by goa to generate the api
var _ = API("congo", func() {
	Title("Congo - Conference Management Made Easy")
	Description("Multi-tenant conference management application")
	Contact(func() {
		Name("Gopher Academy")
		Email("social@gopheracademy.com")
		URL("https://gopheracademy.com")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/gopheracademy/congo/blob/master/LICENSE")
	})
	Docs(func() {
		Description("congo guide")
		URL("https://gopheracademy.com/congo/getting-started.html")
	})
	Host("api.gopheracademy.com")
	Schemes("http")
	BasePath("/congo")

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})

// Account is the account resource media type.
var Account = MediaType("application/vnd.congo.api.account", func() {
	Description("A tenant account")
	Attributes(func() {
		Attribute("id", Integer, "ID of account")
		Attribute("href", String, "API href of account")
		Attribute("name", String, "Name of account")
		Attribute("created_at", String, "Date of creation", func() {
			Format("date-time")
		})
		Attribute("created_by", String, "Email of account owner", func() {
			Format("email")
		})
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("created_at")
		Attribute("created_by")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})
})

// Series represents a recurring conference (GopherCon)
var Series = MediaType("application/vnd.congo.api.series", func() {
	Description("A recurring event or conference")
	Reference(SeriesPayload)
	Attributes(func() {
		Attribute("id", Integer, "ID of series")
		Attribute("href", String, "API href of series")
		Attribute("account", Account, "Account that owns bottle")
		Attribute("created_at", String, "Date of creation", func() {
			Format("date-time")
		})
		Attribute("updated_at", String, "Date of last update", func() {
			Format("date-time")
		})
		// Attributes below inherit from the base type
		Attribute("name")

		Links(func() {
			Link("account")
		})

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
			Attribute("created_at")
			Attribute("updated_at")
			Attribute("links")
		})
	})
})

var _ = Resource("account", func() {

	DefaultMedia(Account)
	BasePath("/accounts")

	Action("show", func() {
		Routing(
			GET("/:accountID"),
		)
		Description("Retrieve account with given id")
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new account")
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(Created, "/accounts/[0-9]+")
	})

	Action("update", func() {
		Routing(
			PUT("/:accountID"),
		)
		Description("Change account name")
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(NoContent)
		Response(NotFound)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:accountID"),
		)
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
})

var _ = Resource("series", func() {

	DefaultMedia(Series)
	BasePath("series")
	Parent("account")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all series in account")
		Response(OK, func() {
			Media(CollectionOf(Series, func() {
				View("default")
				View("tiny")
			}))
		})
	})

	Action("show", func() {
		Routing(
			GET("/:seriesID"),
		)
		Description("Retrieve series with given id")
		Params(func() {
			Param("seriesID", Integer)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new series")
		Payload(SeriesPayload, func() {
			Required("name")
		})
		Response(Created, "^/accounts/[0-9]+/series/[0-9]+$")
	})

	Action("update", func() {
		Routing(
			PATCH("/:seriesID"),
		)
		Params(func() {
			Param("seriesID", Integer)
		})
		Payload(SeriesPayload)
		Response(NoContent)
		Response(NotFound)
	})
})

// SeriesPayload defines the data structure used in the create series request body.
// It is also the base type for the series media type used to render series.
var SeriesPayload = Type("SeriesPayload", func() {
	Attribute("name", func() {
		MinLength(2)
	})

})
