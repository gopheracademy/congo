package design

import (
	"github.com/bketelsen/gorma"
	. "github.com/bketelsen/gorma/dsl"
)

var _ = StorageGroup("CongoStorageGroup", func() {
	Description("This is the global storage group")
	RelationalStore("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		RelationalModel("User", UserPayload, func() {
			Description("User Model")
			HasMany("Reviews", "Review")
			HasMany("Proposals", "Proposal")
			RelationalField("ID", gorma.PKInteger, func() {
				Description("This is the User Model PK field")
			})
			RelationalField("CreatedAt", gorma.Timestamp, func() {})
			RelationalField("UpdatedAt", gorma.Timestamp, func() {})
			RelationalField("DeletedAt", gorma.NullableTimestamp, func() {})
		})

		RelationalModel("Proposal", ProposalPayload, func() {
			Description("Proposal Model")
			BelongsTo("User")
			HasMany("Reviews", "Review")
			RelationalField("ID", gorma.PKInteger, func() {
				Description("This is the Payload Model PK field")
			})
			RelationalField("CreatedAt", gorma.Timestamp, func() {})
			RelationalField("UpdatedAt", gorma.Timestamp, func() {})
			RelationalField("DeletedAt", gorma.NullableTimestamp, func() {})
		})

		RelationalModel("Review", ReviewPayload, func() {
			Description("Review Model")
			BelongsTo("User")
			BelongsTo("Proposal")
			RelationalField("ID", gorma.PKInteger, func() {
				Description("This is the Review Model PK field")
			})
			RelationalField("CreatedAt", gorma.Timestamp, func() {})
			RelationalField("UpdatedAt", gorma.Timestamp, func() {})
			RelationalField("DeletedAt", gorma.NullableTimestamp, func() {})
		})
	})
})
