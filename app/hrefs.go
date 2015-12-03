//************************************************************************//
// congo: Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// AccountHref returns the resource href.
func AccountHref(accountID interface{}) string {
	return fmt.Sprintf("/api/accounts/%v", accountID)
}

// InstanceHref returns the resource href.
func InstanceHref(accountID, seriesID, instanceID interface{}) string {
	return fmt.Sprintf("/api/accounts/%v/series/%v/instances/%v", accountID, seriesID, instanceID)
}

// SeriesHref returns the resource href.
func SeriesHref(accountID, seriesID interface{}) string {
	return fmt.Sprintf("/api/accounts/%v/series/%v", accountID, seriesID)
}

// UserHref returns the resource href.
func UserHref(accountID, userID interface{}) string {
	return fmt.Sprintf("/api/accounts/%v/users/%v", accountID, userID)
}
