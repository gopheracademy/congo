package main

import (
	"net/http"

	"github.com/gopheracademy/congo/models"

	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// ErrBasicAuthFailed means it wasn't able to authenticate you with your login/password.
var ErrBasicAuthFailed = goa.NewErrorClass("basic_auth_failed", 401)

// BasicAuth creates a username/password handler for your `goa.BasicAuthSecurity` method.
//
// Example:
//    app.MyBasicAuthSecurity.Use(BasicAuth(env))
//
func BasicAuth(e *Env) goa.BasicAuthSecurityConfigFunc {
	return func(scheme *goa.BasicAuthSecurity) goa.Middleware {

		middleware, _ := goa.NewMiddleware(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			u, p, ok := r.BasicAuth()
			login := models.Login{
				Email:    u,
				Password: p,
			}
			userdb := models.NewUserDB(e.DB)
			_, err := userdb.GetByLogin(ctx, login)
			if !ok || err != nil {
				return ErrBasicAuthFailed("Authentication failed")
			}
			return nil
		})
		return middleware
	}
}
