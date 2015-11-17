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
