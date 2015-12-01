//************************************************************************//
// congo: Application User Types
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

// SeriesPayload type
type SeriesPayload struct {
	Name string
}

// UserPayload type
type UserPayload struct {
	Email     string
	FirstName string
	LastName  string
}
