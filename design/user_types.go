package design

import (
	. "github.com/raphael/goa/design/dsl"
)

// SeriesModel defines the data structure used in the create series request body.
// It is also the base type for the series media type used to render series.
var SeriesModel = Type("SeriesModel", func() {
	Attribute("name", func() {
		MinLength(2)
	})

})

// AccountModel defines the data structure used in the create account request body.
// It is also the base type for the account media type used to render accounts.
var AccountModel = Type("AccountModel", func() {
	Metadata("github.com/bketelsen/gorma#onetomany", "UserModel")
	Attribute("name", func() {
		MinLength(2)
	})
})

// UserModel defines the data structure used in the create user request body.
// It is also the base type for the user media type used to render users.
var UserModel = Type("UserModel", func() {
	Metadata("github.com/bketelsen/gorma#belongsto", "AccountID")
	Attribute("first_name", func() {
		MinLength(2)
	})
	Attribute("last_name", func() {
		MinLength(2)
	})
	Attribute("email", func() {
		MinLength(2)
	})

})
