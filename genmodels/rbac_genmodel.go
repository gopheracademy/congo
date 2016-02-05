//************************************************************************//
// congo: RBAC
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package genmodels

import "github.com/mikespook/gorbac"

const (
	ADMIN = "Admin"
	USER  = "User"

	AUTHCALLBACK   = "auth.callback"
	AUTHOAUTH      = "auth.oauth"
	AUTHREFRESH    = "auth.refresh"
	AUTHTOKEN      = "auth.token"
	PROPOSALCREATE = "proposal.create"
	PROPOSALDELETE = "proposal.delete"
	PROPOSALLIST   = "proposal.list"
	PROPOSALSHOW   = "proposal.show"
	PROPOSALUPDATE = "proposal.update"
	REVIEWCREATE   = "review.create"
	REVIEWDELETE   = "review.delete"
	REVIEWLIST     = "review.list"
	REVIEWSHOW     = "review.show"
	REVIEWUPDATE   = "review.update"
	USERCREATE     = "user.create"
	USERDELETE     = "user.delete"
	USERLIST       = "user.list"
	USERSHOW       = "user.show"
	USERUPDATE     = "user.update"
)

var RBAC *gorbac.RBAC

// Roler is an interface that provides a Role function
// which returns a string value representing the role to which a user
// belongs.
type Roler interface {
	GetRole() string
}

func Authorize(r Roler, perm string) bool {
	return RBAC.IsGranted(r.GetRole(), perm, nil)
}

// These are provided as a template.  Edit to suit as required by your applicaton
// Test roles in your controllers with models.RBAC.isGranted(ROLE, SOMEPERMISSION, nil)
func init() {
	RBAC = gorbac.New()
	RBAC.Add(USER, []string{
		AUTHCALLBACK,
		AUTHOAUTH,
		AUTHREFRESH,
		AUTHTOKEN,
		PROPOSALCREATE,
		PROPOSALDELETE,
		PROPOSALLIST,
		PROPOSALSHOW,
		PROPOSALUPDATE,
		REVIEWCREATE,
		REVIEWDELETE,
		REVIEWLIST,
		REVIEWSHOW,
		REVIEWUPDATE,
		USERCREATE,
		USERDELETE,
		USERLIST,
		USERSHOW,
		USERUPDATE,
	}, nil)
	RBAC.Add(ADMIN, []string{
		AUTHCALLBACK,
		AUTHOAUTH,
		AUTHREFRESH,
		AUTHTOKEN,
		PROPOSALCREATE,
		PROPOSALDELETE,
		PROPOSALLIST,
		PROPOSALSHOW,
		PROPOSALUPDATE,
		REVIEWCREATE,
		REVIEWDELETE,
		REVIEWLIST,
		REVIEWSHOW,
		REVIEWUPDATE,
		USERCREATE,
		USERDELETE,
		USERLIST,
		USERSHOW,
		USERUPDATE,
	}, []string{USER})
}
