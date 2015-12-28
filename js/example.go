//************************************************************************//
// congo JavaScript Client Example
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package js

import (
	"net/http"

	"github.com/raphael/goa"
)

// MountController mounts the JavaScript example controller under "/js".
func MountController(service goa.Service) {
	// Serve static files under js
	service.ServeFiles("/api/js/*filepath", http.Dir("/home/bketelsen/src/github.com/gopheracademy/congo/js"))
	service.Info("mount", "ctrl", "JS", "action", "ServeFiles", "route", "GET /js/*")
}
