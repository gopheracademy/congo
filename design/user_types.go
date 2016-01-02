package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

// UserModel defines the data structure used in the create user request body.
// It is also the base type for the user media type used to render users.
var UserModel = Type("UserModel", func() {
	Metadata("github.com/bketelsen/gorma", "Model")
	Metadata("github.com/bketelsen/gorma#authboss", "All")
	Metadata("github.com/bketelsen/gorma#roler", "true")
	Metadata("github.com/bketelsen/gorma#hasmany", "Proposal,Review")
	Attribute("firstname", func() {
	})
	Attribute("lastname", func() {
	})
	Attribute("city", func() {
	})
	Attribute("state", func() {
	})
	Attribute("country", func() {
	})
	Attribute("email", func() {
		Format("email")
	})
	Attribute("bio", func() {
		MaxLength(500)
	})
	Attribute("role", func() {
	})

})

// ProposalModel defines the data structure used in the create proposal request body.
// It is also the base type for the proposal media type used to render users.
var ProposalModel = Type("ProposalModel", func() {
	Metadata("github.com/bketelsen/gorma", "Model")
	Metadata("github.com/bketelsen/gorma#belongsto", "User")
	Metadata("github.com/bketelsen/gorma#hasmany", "Review")
	Attribute("firstname", func() {
		MinLength(2)
	})
	Attribute("title", func() {
		MinLength(10)
		MaxLength(200)
	})
	Attribute("abstract", func() {
		MinLength(50)
		MaxLength(500)
	})
	Attribute("detail", func() {
		MinLength(100)
		MaxLength(2000)
	})
	Attribute("withdrawn", Boolean)
})

// ReviewModel defines the data structure used to create a review request body
// It is also the base type for the review media type used to render reviews
var ReviewModel = Type("ReviewModel", func() {
	Metadata("github.com/bketelsen/gorma", "Model")
	Metadata("github.com/bketelsen/gorma#belongsto", "Proposal,User")
	Attribute("comment", func() {
		MinLength(10)
		MaxLength(200)
	})
	Attribute("rating", Integer, func() {
		Minimum(1)
		Maximum(5)
	})
})
