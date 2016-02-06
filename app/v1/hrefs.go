//************************************************************************//
// API "congo" version v1: Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v1

import "fmt"

// ProposalHref returns the resource href.
func ProposalHref(version, userID, proposalID interface{}) string {
	return fmt.Sprintf("/%v/api/users/%v/proposals/%v", version, userID, proposalID)
}

// ReviewHref returns the resource href.
func ReviewHref(version, userID, proposalID, reviewID interface{}) string {
	return fmt.Sprintf("/%v/api/users/%v/proposals/%v/review/%v", version, userID, proposalID, reviewID)
}

// UserHref returns the resource href.
func UserHref(version, userID interface{}) string {
	return fmt.Sprintf("/%v/api/users/%v", version, userID)
}
