package design

import (
	. "github.com/raphael/goa/design/dsl"
)

// SeriesPayload defines the data structure used in the create series request body.
// It is also the base type for the series media type used to render series.
var SeriesPayload = Type("SeriesPayload", func() {
	Attribute("name", func() {
		MinLength(2)
	})

})

// UserPayload defines the data structure used in the create user request body.
// It is also the base type for the user media type used to render users.
var UserPayload = Type("UserPayload", func() {
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
