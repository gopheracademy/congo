package handlers

import "net/http"

func Accounts(ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ren.Render(w, http.StatusOK, "accounts", struct{}{}, "layout")
	})
}
