package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// UserPayload defines the data structure used in the create user request body.
// It is also the base type for the user media type used to render users.
var UserPayload = Type("UserPayload", func() {
	Attribute("id", Integer)
	Attribute("first_name", func() {
		MinLength(2)
	})
	Attribute("last_name", func() {
		MinLength(2)
	})
	Attribute("email", func() {
		MinLength(2)
	})
	Attribute("password", func() {
		MinLength(8)
	})
	Attribute("validated", Boolean, func() {
	})

	Attribute("validation_code", String, func() {
		MinLength(8)
	})

	Attribute("role")
})

// TenantPayload defines the data structure used in the create tenant request body.
// It is also the base type for the tenant media type used to render tenants.
var TenantPayload = Type("TenantPayload", func() {
	Attribute("name", func() {
		MinLength(2)
	})
})

// SeriesPayload defines the data structure used in the create series request body.
// It is also the base type for the series media type used to render series.
var SeriesPayload = Type("SeriesPayload", func() {
	Attribute("name", func() {
		MinLength(2)
	})
})

// EventPayload defines the data structure used in the create event request body.
// It is also the base type for the event media type used to render events.
var EventPayload = Type("EventPayload", func() {
	Attribute("name", func() {
		MinLength(2)
	})
	Attribute("url", func() {
		MinLength(5)
	})
	Attribute("start_date", DateTime)
	Attribute("end_date", DateTime)

})

// SpeakerPayload defines the data structure used in the create speaker request body.
// It is also the base type for the speaker media type used to render speakers.
var SpeakerPayload = Type("SpeakerPayload", func() {
	Attribute("first_name", func() {
		MinLength(2)
	})
	Attribute("last_name", func() {
		MinLength(2)
	})
	Attribute("bio")
	Attribute("twitter")
	Attribute("linkedin")
	Attribute("github")
	Attribute("image_url", func() {
		MinLength(5)
	})

})

// PresentationPayload defines the data structure used in the create presentation request body.
// It is also the base type for the presentation media type used to render presentation.
var PresentationPayload = Type("PresentationPayload", func() {
	Attribute("abstract", func() {
		MinLength(10)
	})
	Attribute("name")
	Attribute("detail")
})

// AdminUserPayload defines the data structure used in the create user request body.
// It is also the base type for the user media type used to render users.
var AdminUserPayload = Type("AdminUserPayload", func() {
	Attribute("id", Integer)
	Attribute("tenant_id", Integer)
	Attribute("first_name", func() {
		MinLength(2)
	})
	Attribute("last_name", func() {
		MinLength(2)
	})
	Attribute("email", func() {
		MinLength(2)
	})
	Attribute("password", func() {
		MinLength(8)
	})
	Attribute("validated", Boolean, func() {
	})

	Attribute("validation_code", String, func() {
		MinLength(8)
	})

	Attribute("role")
})
