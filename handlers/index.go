package handlers

import (
	"net/http"

	"gopkg.in/authboss.v0"
)

func Index(ren Renderer, ld func(w http.ResponseWriter, r *http.Request) authboss.HTMLData) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		data := ld(w, r)
		ren.Render(w, http.StatusOK, "index", data, "layout")
	})
}
