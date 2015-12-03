package handlers

import "net/http"

func Series(ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ren.Render(w, http.StatusOK, "series", struct{}{}, "layout")
	})
}
