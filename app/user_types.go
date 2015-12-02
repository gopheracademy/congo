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

// AccountModel type
type AccountModel struct {
	Name string
}

// SeriesModel type
type SeriesModel struct {
	Name string
}

// UserModel type
type UserModel struct {
	Email     string
	FirstName string
	LastName  string
}
