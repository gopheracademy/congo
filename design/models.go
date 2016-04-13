package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var sg = StorageGroup("Congo Storage", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the postgres relational store")
		Model("User", func() {
			BuildsFrom(func() {
				Payload("user", "create")
				Payload("user", "update")
				Payload("adminuser", "create")
				Payload("adminuser", "update")
			})
			RendersTo(User)
			Description("This is the User model")
			BelongsTo("Tenant")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
			Field("first_name", gorma.String, func() {
				Alias("firstname")
			})

			Field("last_name", gorma.String, func() {
				Alias("lastname")
			})
			Field("href", gorma.String)
		})

		Model("Tenant", func() {
			BuildsFrom(func() {
				Payload("tenant", "create")
				Payload("tenant", "update")
			})
			RendersTo(Tenant)
			Description("This is the Tenant model")
			HasMany("users", "User")
			HasMany("events", "Event")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})
		Model("Series", func() {
			BuildsFrom(func() {
				Payload("series", "create")
				Payload("series", "update")
			})
			RendersTo(Series)
			BelongsTo("Tenant")
			Description("This is the Series model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})
		Model("Event", func() {
			BuildsFrom(func() {
				Payload("event", "create")
				Payload("event", "update")
			})
			RendersTo(Event)
			BelongsTo("Tenant")
			Description("This is the Event model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})

	})
})
