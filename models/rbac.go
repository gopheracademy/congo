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

package models

import "github.com/mikespook/gorbac"

const (
	ADMIN = "Admin"
	USER  = "User"

	ACCOUNTCREATE  = "account.create"
	ACCOUNTDELETE  = "account.delete"
	ACCOUNTLIST    = "account.list"
	ACCOUNTSHOW    = "account.show"
	ACCOUNTUPDATE  = "account.update"
	INSTANCECREATE = "instance.create"
	INSTANCEDELETE = "instance.delete"
	INSTANCELIST   = "instance.list"
	INSTANCESHOW   = "instance.show"
	INSTANCEUPDATE = "instance.update"
	SERIESCREATE   = "series.create"
	SERIESDELETE   = "series.delete"
	SERIESLIST     = "series.list"
	SERIESSHOW     = "series.show"
	SERIESUPDATE   = "series.update"
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
	Role() string
}

func Authorize(r Roler, perm string) bool {
	return RBAC.IsGranted(r.Role(), perm, nil)
}

// These are provided as a template.  Edit to suit as required by your applicaton
// Test roles in your controllers with models.RBAC.isGranted(ROLE, SOMEPERMISSION, nil)
func init() {
	RBAC = gorbac.New()
	RBAC.Add(USER, []string{
		ACCOUNTCREATE,
		ACCOUNTDELETE,
		ACCOUNTLIST,
		ACCOUNTSHOW,
		ACCOUNTUPDATE,
		INSTANCECREATE,
		INSTANCEDELETE,
		INSTANCELIST,
		INSTANCESHOW,
		INSTANCEUPDATE,
		SERIESCREATE,
		SERIESDELETE,
		SERIESLIST,
		SERIESSHOW,
		SERIESUPDATE,
		USERCREATE,
		USERDELETE,
		USERLIST,
		USERSHOW,
		USERUPDATE,
	}, nil)
	RBAC.Add(ADMIN, []string{
		ACCOUNTCREATE,
		ACCOUNTDELETE,
		ACCOUNTLIST,
		ACCOUNTSHOW,
		ACCOUNTUPDATE,
		INSTANCECREATE,
		INSTANCEDELETE,
		INSTANCELIST,
		INSTANCESHOW,
		INSTANCEUPDATE,
		SERIESCREATE,
		SERIESDELETE,
		SERIESLIST,
		SERIESSHOW,
		SERIESUPDATE,
		USERCREATE,
		USERDELETE,
		USERLIST,
		USERSHOW,
		USERUPDATE,
	}, []string{USER})
}
