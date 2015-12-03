package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var _ = Resource("account", func() {

	DefaultMedia(Account)
	BasePath("/accounts")
	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all accounts")
		Response(OK, func() {
			Media(CollectionOf(Account, func() {
				View("default")
			}))
		})
	})
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
		Payload(AccountModel, func() {
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
		Payload(AccountModel, func() {
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
var _ = Resource("user", func() {

	DefaultMedia(User)
	BasePath("users")
	Parent("account")
	Metadata("resource", "123")
	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all users in account")
		Response(OK, func() {
			Media(CollectionOf(User, func() {
				View("default")
			}))
		})
	})

	Action("show", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Retrieve user with given id")
		Params(func() {
			Param("userID", Integer)
		})
		Metadata("action", "123")
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new user")
		Payload(UserModel, func() {
			Required("first_name")
			Required("last_name")
			Required("email")
		})
		Response(Created, "^/accounts/[0-9]+/users/[0-9]+$")
	})

	Action("update", func() {
		Routing(
			PATCH("/:userID"),
		)
		Params(func() {
			Param("userID", Integer)
		})
		Payload(UserModel)
		Response(NoContent)
		Response(NotFound)
	})
	Action("delete", func() {
		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", Integer, "User ID")
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
		Payload(SeriesModel, func() {
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
		Payload(SeriesModel)
		Response(NoContent)
		Response(NotFound)
	})
	Action("delete", func() {
		Routing(
			DELETE("/:seriesID"),
		)
		Params(func() {
			Param("seriesID", Integer, "Series ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
})
