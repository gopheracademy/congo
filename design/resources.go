package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

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
