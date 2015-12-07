package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

// SeriesModel defines the data structure used in the create series request body.
// It is also the base type for the series media type used to render series.
var SeriesModel = Type("SeriesModel", func() {
	Metadata("github.com/bketelsen/gorma", "Model")
	Metadata("github.com/bketelsen/gorma#belongsto", "Account")
	Attribute("name", func() {
		MinLength(2)
	})

})

// InstanceModel defines the data structure used in the create instance request body.
// It is also the base type for the instance media type used to render instances.
var InstanceModel = Type("InstanceModel", func() {
	Metadata("github.com/bketelsen/gorma", "Model")
	Metadata("github.com/bketelsen/gorma#belongsto", "Series")
	Attribute("name", func() {
		MinLength(2)
	})

})

// AccountModel defines the data structure used in the create account request body.
// It is also the base type for the account media type used to render accounts.
var AccountModel = Type("AccountModel", func() {
	Metadata("github.com/bketelsen/gorma", "Model")
	Metadata("github.com/bketelsen/gorma#hasmany", "User,Account")
	Attribute("name", func() {
		MinLength(2)
	})
})

// UserModel defines the data structure used in the create user request body.
// It is also the base type for the user media type used to render users.
var UserModel = Type("UserModel", func() {
	Metadata("github.com/bketelsen/gorma", "Model")
	Metadata("github.com/bketelsen/gorma#belongsto", "Account")
	Metadata("github.com/bketelsen/gorma#authboss", "All")
	Metadata("github.com/bketelsen/gorma#roler", "xxx")
	Attribute("id", Integer, func() {
	})
	Attribute("first_name", func() {
		MinLength(2)
	})
	Attribute("last_name", func() {
		MinLength(2)
	})
	Attribute("email", func() {
		MinLength(2)
	})
	Attribute("role", func() {
	})

})
