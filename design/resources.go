package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("auth", func() {
	Security("password")
	DefaultMedia(Authorize)
	BasePath("/auth")

	Action("token", func() {
		Routing(
			POST("/token"),
		)
		Description("Obtain an access token")

		Response(Created, func() {
			Media(Authorize)
		})
	})
	Action("refresh", func() {
		Routing(
			POST("/refresh"),
		)
		Description("Obtain a refreshed access token")
		Response(Created, func() {
			Media(Authorize)
		})
	})
})

var _ = Resource("healthz", func() {
	NoSecurity()
	BasePath("/healthz")
	Action("status", func() {
		Routing(
			GET(""),
		)
		Description("Get Server Status")
	})
})

var _ = Resource("adminuser", func() {
	DefaultMedia(User)

	Security("jwt")
	BasePath("/admin/users")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all users")
		Response(OK, func() {
			Media(CollectionOf(User, func() {
				View("default")
				View("tenant")
			}))
		})
		Response(NotFound)
	})

	Action("show", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Retrieve user with given id")
		Params(func() {
			Param("userID", Integer)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new user")
		Payload(AdminUserPayload, func() {
			Required("first_name")
			Required("last_name")
			Required("email")
			Required("password")
			Required("role")
		})
		Response(Created, "^/admin/users/[0-9]+$")
	})

	Action("update", func() {
		Routing(
			PATCH("/:userID"),
		)
		Params(func() {
			Param("userID", Integer)
		})
		Payload(UserPayload)
		Response(NoContent)
		Response(NotFound)
	})
	Action("delete", func() {
		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", Integer, "UserID ")
		})
		Response(NoContent)
		Response(NotFound)
	})

})

var _ = Resource("validate", func() {
	BasePath("/validate")

	NoSecurity()
	Action("validate", func() {
		Routing(
			GET("/:userID"),
		)
		Description("validate user email")
		Params(func() {
			Param("v", String)
		})
		Response(OK)
		Response(NotFound)
	})

})

var _ = Resource("tenant", func() {
	DefaultMedia(Tenant)
	BasePath("/tenants")
	Security("jwt")
	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all tenants")
		Response(OK, func() {
			Media(CollectionOf(Tenant, func() {
				View("default")
			}))
		})
		Response(NotFound)
	})

	Action("show", func() {
		Routing(
			GET("/:tenantID"),
		)
		Description("Retrieve tenant with given id")
		Params(func() {
			Param("tenantID", Integer)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new tenant")
		Payload(TenantPayload, func() {
			Required("name")
		})
		Response(Created, "^/tenants/[0-9]+$")
	})

	Action("update", func() {
		Routing(
			PATCH("/:tenantID"),
		)
		Params(func() {
			Param("tenantID", Integer)
		})
		Payload(TenantPayload)
		Response(NoContent)
		Response(NotFound)
	})
	Action("delete", func() {
		Routing(
			DELETE("/:tenantID"),
		)
		Params(func() {
			Param("tenantID", Integer, "Tenant ID")
		})
		Response(NoContent)
		Response(NotFound)
	})

})
var _ = Resource("series", func() {
	DefaultMedia(Series)
	Parent("tenant")
	Security("jwt")
	BasePath("/series")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all series")
		Response(OK, func() {
			Media(CollectionOf(Series, func() {
				View("default")
			}))
		})
		Response(NotFound)
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
		Payload(TenantPayload, func() {
			Required("name")
		})
		Response(Created, "^/tenants/[0-9]/series/[0-9]+$")
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
var _ = Resource("event", func() {
	DefaultMedia(Event)
	Parent("tenant")
	BasePath("/events")
	Security("jwt")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all events")
		Response(OK, func() {
			Media(CollectionOf(Event, func() {
				View("default")
			}))
		})
		Response(NotFound)
	})

	Action("show", func() {
		NoSecurity()
		Routing(
			GET("/:eventID"),
		)
		Description("Retrieve event with given id")
		Params(func() {
			Param("eventID", Integer)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new event")
		Payload(TenantPayload, func() {
			Required("name")
		})
		Response(Created, "^/tenants/[0-9]/events/[0-9]+$")
	})

	Action("update", func() {
		Routing(
			PATCH("/:eventID"),
		)
		Params(func() {
			Param("eventID", Integer)
		})
		Payload(EventPayload)
		Response(NoContent)
		Response(NotFound)
	})
	Action("delete", func() {
		Routing(
			DELETE("/:eventID"),
		)
		Params(func() {
			Param("eventID", Integer, "Event ID")
		})
		Response(NoContent)
		Response(NotFound)
	})

})

var _ = Resource("user", func() {
	DefaultMedia(User)
	BasePath("/users")
	Parent("tenant")
	Security("jwt")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all users for a tenant")
		Response(OK, func() {
			Media(CollectionOf(User, func() {
				View("default")
			}))
		})
		Response(NotFound)
	})

	Action("show", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Retrieve user with given id")
		Params(func() {
			Param("userID", Integer)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new user")
		Payload(UserPayload, func() {
			Required("first_name")
			Required("last_name")
			Required("email")
			Required("password")
			Required("role")
		})
		Response(Created, "^/tenants/[0-9]+/users/[0-9]+$")
	})

	Action("update", func() {
		Routing(
			PATCH("/:userID"),
		)
		Params(func() {
			Param("userID", Integer)
		})
		Payload(UserPayload)
		Response(NoContent)
		Response(NotFound)
	})
	Action("delete", func() {
		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", Integer, "UserID ")
		})
		Response(NoContent)
		Response(NotFound)
	})

})
