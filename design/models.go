package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var sg = StorageGroup("XOR Storage", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the postgresrelational store")
		Model("User", func() {
			BuildsFrom(func() {
				Payload("user", "create")
				Payload("user", "update")
<<<<<<< HEAD
				Payload("adminuser", "create")
				Payload("adminuser", "update")
			})
			RendersTo(User)
			Description("This is the User model")
			BelongsTo("Member")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
			Field("first_name", gorma.String, func() {
				Alias("firstname")
			})

			Field("last_name", gorma.String, func() {
				Alias("lastname")
=======
			})
			RendersTo(User)
			Description("User Model")
			HasMany("Reviews", "Review")
			HasMany("Proposals", "Proposal")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the User Model PK field")
>>>>>>> d16d31f78e7efce748d2a2e8192efa72fe080e47
			})
			Field("href", gorma.String)
		})
		Model("MeasuredReciprocity", func() {
			BuildsFrom(func() {
				Payload("measured_reciprocity", "create")
				Payload("measured_reciprocity", "update")
			})
			RendersTo(MeasuredReciprocity)
			Description("This is the MR model")
			NoAutomaticIDFields()
			Field("OwningMemberID", gorma.Integer, func() {
				PrimaryKey()
				Description("This is part the ID PK field")
			})

<<<<<<< HEAD
			Field("href", gorma.String)
			Field("GrantMemberID", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the other part the ID PK field")
			})
		})
		Model("Signup", func() {
			BuildsFrom(func() {
				Payload("signup", "create")
				Payload("signup", "update")
				Payload("signup", "approve")
				Payload("signup", "reject")
			})
			RendersTo(Signup)
			Description("This is the Signup model")
			BelongsTo("Exchange")
			HasOne("Industry")
			HasOne("Product")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
			Field("href", gorma.String)
		})
		Model("Application", func() {
			BuildsFrom(func() {
				Payload("application", "create")
				Payload("application", "update")
			})
			Alias("apps")
			RendersTo(Application)
			Description("This is the Application model")
			BelongsTo("Member")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
			Field("href", gorma.String)
		})
		Model("GiveToGet", func() {
			BuildsFrom(func() {
				Payload("give_to_get", "create")
				Payload("give_to_get", "update")
			})
			RendersTo(GiveToGet)
			Description("This is the GiveToGet model")
			HasOne("Portfolio")
			HasOne("Datatype")
			HasOne("DatatypeField")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
			Field("href", gorma.String)
		})
		Model("Datatype", func() {
			BuildsFrom(func() {
				Payload("datatype", "create")
				Payload("datatype", "update")
			})
			RendersTo(Datatype)
			Description("This is the Datatype model")
			BelongsTo("Exchange")
			HasMany("DatatypeFields", "DatatypeField")
			Field("href", gorma.String)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})
		Model("DatatypeField", func() {
			BuildsFrom(func() {
				Payload("datatype_field", "create")
				Payload("datatype_field", "update")
			})
			RendersTo(DatatypeField)
			Description("This is the DatatypeField model")
			BelongsTo("Datatype")
			Field("href", gorma.String)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})

		Model("Industry", func() {
			BuildsFrom(func() {
				Payload("industry", "create")
				Payload("industry", "update")
			})
			RendersTo(Industry)
			Description("This is the Industry model")
			HasMany("Products", "Product")
			Field("href", gorma.String)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})
		Model("Member", func() {
			BuildsFrom(func() {
				Payload("member", "create")
				Payload("member", "update")
			})
			RendersTo(ExchangeMember)
			Description("This is the Member model")
			HasMany("Portfolios", "Portfolio")
			HasMany("Applications", "Application")
			HasMany("Users", "User")
			HasMany("Purchases", "Purchase")
			Field("href", gorma.String)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})
		Model("ProductPackage", func() {
			BuildsFrom(func() {
				Payload("product_package", "create")
				Payload("product_package", "update")
			})
			RendersTo(ProductPackage)
			Description("This is the ProductPackage model")
			BelongsTo("Product")
			Field("href", gorma.String)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})
		Model("Portfolio", func() {
			BuildsFrom(func() {
				Payload("portfolio", "create")
				Payload("portfolio", "update")
			})
			RendersTo(Portfolio)
			Description("This is the Portfolio model")
			BelongsTo("Member")
			HasMany("PortfolioPermissions", "PortfolioPermission")
			HasOne("Industry")
			ManyToMany("Product", "portfolio_products")
			Field("href", gorma.String)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})
		Model("PortfolioPermission", func() {
			BuildsFrom(func() {
				Payload("portfolio_permission", "create")
				Payload("portfolio_permission", "update")
			})
			RendersTo(PortfolioPermission)
			Description("This is the PortfolioPermissio model")
			BelongsTo("Portfolio")
			HasOne("Product")
			HasOne("Industry")
			HasOne("Datatype")
			HasOne("DatatypeField")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
			Field("href", gorma.String)
		})
		Model("Product", func() {
			BuildsFrom(func() {
				Payload("product", "create")
				Payload("product", "update")
			})
			RendersTo(Product)
			Description("This is the Product model")
			BelongsTo("Exchange")
			HasMany("ProductPackages", "ProductPackage")
			Field("href", gorma.String)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})
		Model("Purchase", func() {
			BuildsFrom(func() {
				Payload("purchase", "create")
				Payload("purchase", "update")
			})
			RendersTo(Purchase)
			Description("This is the Purchase model")
			BelongsTo("Member")
			HasOne("ProductPackage")
			Field("href", gorma.String)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})
		Model("Usage", func() {
			BuildsFrom(func() {
				Payload("usage", "create")
				Payload("usage", "update")
			})
			RendersTo(Usage)
			Description("This is the Usage model")
			BelongsTo("Purchase")
			Field("href", gorma.String)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
		})

		Model("Exchange", func() {
			BuildsFrom(func() {
				Payload("exchange", "create")
				Payload("exchange", "update")
			})
			RendersTo(Exchange)
			Description("This is the Exchange model")
			HasMany("Datatypes", "Datatype")
			HasMany("Products", "Product")
			ManyToMany("Industry", "exchange_industries")
			ManyToMany("Member", "exchange_members")
			ManyToMany("Portfolio", "exchange_memberportfolios")
			Field("href", gorma.String)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the ID PK field")
			})
=======
		Model("Proposal", func() {
			BuildsFrom(func() {
				Payload("proposal", "create")
				Payload("proposal", "update")
			})
			RendersTo(Proposal)
			Description("Proposal Model")
			BelongsTo("User")
			HasMany("Reviews", "Review")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the Payload Model PK field")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Review", func() {
			BuildsFrom(func() {
				Payload("review", "create")
				Payload("review", "update")
			})
			RendersTo(Review)
			Description("Review Model")
			BelongsTo("User")
			BelongsTo("Proposal")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the Review Model PK field")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
>>>>>>> d16d31f78e7efce748d2a2e8192efa72fe080e47
		})

	})
})
