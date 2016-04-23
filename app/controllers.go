//************************************************************************//
// API "congo": Application Controllers
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

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder(goa.NewJSONEncoder, "application/json")
	service.Encoder(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder(goa.NewXMLEncoder, "application/xml")
	service.Decoder(goa.NewJSONDecoder, "application/json")
	service.Decoder(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder(goa.NewJSONEncoder, "*/*")
	service.Decoder(goa.NewJSONDecoder, "*/*")
}

// AdminuserController is the controller interface for the Adminuser actions.
type AdminuserController interface {
	goa.Muxer
	Create(*CreateAdminuserContext) error
	Delete(*DeleteAdminuserContext) error
	List(*ListAdminuserContext) error
	Show(*ShowAdminuserContext) error
	Update(*UpdateAdminuserContext) error
}

// MountAdminuserController "mounts" a Adminuser resource controller on the given service.
func MountAdminuserController(service *goa.Service, ctrl AdminuserController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/admin/users", cors.HandlePreflight(service.Context, handleAdminuserOrigin))
	service.Mux.Handle("OPTIONS", "/api/admin/users/:userID", cors.HandlePreflight(service.Context, handleAdminuserOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateAdminuserContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateAdminuserPayload)
		}
		return ctrl.Create(rctx)
	}
	h = handleAdminuserOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("POST", "/api/admin/users", ctrl.MuxHandler("Create", h, unmarshalCreateAdminuserPayload))
	service.LogInfo("mount", "ctrl", "Adminuser", "action", "Create", "route", "POST /api/admin/users", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteAdminuserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleAdminuserOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("DELETE", "/api/admin/users/:userID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Adminuser", "action", "Delete", "route", "DELETE /api/admin/users/:userID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListAdminuserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleAdminuserOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/admin/users", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Adminuser", "action", "List", "route", "GET /api/admin/users", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowAdminuserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleAdminuserOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/admin/users/:userID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Adminuser", "action", "Show", "route", "GET /api/admin/users/:userID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateAdminuserContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateAdminuserPayload)
		}
		return ctrl.Update(rctx)
	}
	h = handleAdminuserOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/admin/users/:userID", ctrl.MuxHandler("Update", h, unmarshalUpdateAdminuserPayload))
	service.LogInfo("mount", "ctrl", "Adminuser", "action", "Update", "route", "PATCH /api/admin/users/:userID", "security", "jwt")
}

// handleAdminuserOrigin applies the CORS response headers corresponding to the origin.
func handleAdminuserOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateAdminuserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateAdminuserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createAdminuserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateAdminuserPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateAdminuserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateAdminuserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Muxer
	Refresh(*RefreshAuthContext) error
	Token(*TokenAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service *goa.Service, ctrl AuthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/auth/refresh", cors.HandlePreflight(service.Context, handleAuthOrigin))
	service.Mux.Handle("OPTIONS", "/api/auth/token", cors.HandlePreflight(service.Context, handleAuthOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewRefreshAuthContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Refresh(rctx)
	}
	h = handleAuthOrigin(h)
	h = handleSecurity("password", h)
	service.Mux.Handle("POST", "/api/auth/refresh", ctrl.MuxHandler("Refresh", h, nil))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Refresh", "route", "POST /api/auth/refresh", "security", "password")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewTokenAuthContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Token(rctx)
	}
	h = handleAuthOrigin(h)
	h = handleSecurity("password", h)
	service.Mux.Handle("POST", "/api/auth/token", ctrl.MuxHandler("Token", h, nil))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Token", "route", "POST /api/auth/token", "security", "password")
}

// handleAuthOrigin applies the CORS response headers corresponding to the origin.
func handleAuthOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// EventController is the controller interface for the Event actions.
type EventController interface {
	goa.Muxer
	Create(*CreateEventContext) error
	Delete(*DeleteEventContext) error
	List(*ListEventContext) error
	Show(*ShowEventContext) error
	Update(*UpdateEventContext) error
}

// MountEventController "mounts" a Event resource controller on the given service.
func MountEventController(service *goa.Service, ctrl EventController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID/events", cors.HandlePreflight(service.Context, handleEventOrigin))
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID/events/:eventID", cors.HandlePreflight(service.Context, handleEventOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateEventContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateEventPayload)
		}
		return ctrl.Create(rctx)
	}
	h = handleEventOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("POST", "/api/tenants/:tenantID/events", ctrl.MuxHandler("Create", h, unmarshalCreateEventPayload))
	service.LogInfo("mount", "ctrl", "Event", "action", "Create", "route", "POST /api/tenants/:tenantID/events", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteEventContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleEventOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("DELETE", "/api/tenants/:tenantID/events/:eventID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Event", "action", "Delete", "route", "DELETE /api/tenants/:tenantID/events/:eventID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListEventContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleEventOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID/events", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Event", "action", "List", "route", "GET /api/tenants/:tenantID/events", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowEventContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleEventOrigin(h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID/events/:eventID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Event", "action", "Show", "route", "GET /api/tenants/:tenantID/events/:eventID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateEventContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateEventPayload)
		}
		return ctrl.Update(rctx)
	}
	h = handleEventOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/tenants/:tenantID/events/:eventID", ctrl.MuxHandler("Update", h, unmarshalUpdateEventPayload))
	service.LogInfo("mount", "ctrl", "Event", "action", "Update", "route", "PATCH /api/tenants/:tenantID/events/:eventID", "security", "jwt")
}

// handleEventOrigin applies the CORS response headers corresponding to the origin.
func handleEventOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateEventPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateEventPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createEventPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateEventPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateEventPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateEventPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// HealthzController is the controller interface for the Healthz actions.
type HealthzController interface {
	goa.Muxer
	Status(*StatusHealthzContext) error
}

// MountHealthzController "mounts" a Healthz resource controller on the given service.
func MountHealthzController(service *goa.Service, ctrl HealthzController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/healthz", cors.HandlePreflight(service.Context, handleHealthzOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewStatusHealthzContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Status(rctx)
	}
	h = handleHealthzOrigin(h)
	service.Mux.Handle("GET", "/api/healthz", ctrl.MuxHandler("Status", h, nil))
	service.LogInfo("mount", "ctrl", "Healthz", "action", "Status", "route", "GET /api/healthz")
}

// handleHealthzOrigin applies the CORS response headers corresponding to the origin.
func handleHealthzOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// PresentationController is the controller interface for the Presentation actions.
type PresentationController interface {
	goa.Muxer
	Create(*CreatePresentationContext) error
	Delete(*DeletePresentationContext) error
	List(*ListPresentationContext) error
	Show(*ShowPresentationContext) error
	Update(*UpdatePresentationContext) error
}

// MountPresentationController "mounts" a Presentation resource controller on the given service.
func MountPresentationController(service *goa.Service, ctrl PresentationController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations", cors.HandlePreflight(service.Context, handlePresentationOrigin))
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations/:presentationID", cors.HandlePreflight(service.Context, handlePresentationOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreatePresentationContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreatePresentationPayload)
		}
		return ctrl.Create(rctx)
	}
	h = handlePresentationOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("POST", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations", ctrl.MuxHandler("Create", h, unmarshalCreatePresentationPayload))
	service.LogInfo("mount", "ctrl", "Presentation", "action", "Create", "route", "POST /api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeletePresentationContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handlePresentationOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("DELETE", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations/:presentationID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Presentation", "action", "Delete", "route", "DELETE /api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations/:presentationID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListPresentationContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handlePresentationOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Presentation", "action", "List", "route", "GET /api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowPresentationContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handlePresentationOrigin(h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations/:presentationID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Presentation", "action", "Show", "route", "GET /api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations/:presentationID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdatePresentationContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdatePresentationPayload)
		}
		return ctrl.Update(rctx)
	}
	h = handlePresentationOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations/:presentationID", ctrl.MuxHandler("Update", h, unmarshalUpdatePresentationPayload))
	service.LogInfo("mount", "ctrl", "Presentation", "action", "Update", "route", "PATCH /api/tenants/:tenantID/events/:eventID/speakers/:speakerID/presentations/:presentationID", "security", "jwt")
}

// handlePresentationOrigin applies the CORS response headers corresponding to the origin.
func handlePresentationOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreatePresentationPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreatePresentationPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createPresentationPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdatePresentationPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdatePresentationPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updatePresentationPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SeriesController is the controller interface for the Series actions.
type SeriesController interface {
	goa.Muxer
	Create(*CreateSeriesContext) error
	Delete(*DeleteSeriesContext) error
	List(*ListSeriesContext) error
	Show(*ShowSeriesContext) error
	Update(*UpdateSeriesContext) error
}

// MountSeriesController "mounts" a Series resource controller on the given service.
func MountSeriesController(service *goa.Service, ctrl SeriesController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID/series", cors.HandlePreflight(service.Context, handleSeriesOrigin))
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID/series/:seriesID", cors.HandlePreflight(service.Context, handleSeriesOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateSeriesContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateSeriesPayload)
		}
		return ctrl.Create(rctx)
	}
	h = handleSeriesOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("POST", "/api/tenants/:tenantID/series", ctrl.MuxHandler("Create", h, unmarshalCreateSeriesPayload))
	service.LogInfo("mount", "ctrl", "Series", "action", "Create", "route", "POST /api/tenants/:tenantID/series", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteSeriesContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSeriesOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("DELETE", "/api/tenants/:tenantID/series/:seriesID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Series", "action", "Delete", "route", "DELETE /api/tenants/:tenantID/series/:seriesID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListSeriesContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSeriesOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID/series", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Series", "action", "List", "route", "GET /api/tenants/:tenantID/series", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowSeriesContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSeriesOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID/series/:seriesID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Series", "action", "Show", "route", "GET /api/tenants/:tenantID/series/:seriesID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateSeriesContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateSeriesPayload)
		}
		return ctrl.Update(rctx)
	}
	h = handleSeriesOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/tenants/:tenantID/series/:seriesID", ctrl.MuxHandler("Update", h, unmarshalUpdateSeriesPayload))
	service.LogInfo("mount", "ctrl", "Series", "action", "Update", "route", "PATCH /api/tenants/:tenantID/series/:seriesID", "security", "jwt")
}

// handleSeriesOrigin applies the CORS response headers corresponding to the origin.
func handleSeriesOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateSeriesPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateSeriesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createSeriesPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateSeriesPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateSeriesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateSeriesPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SpeakerController is the controller interface for the Speaker actions.
type SpeakerController interface {
	goa.Muxer
	Create(*CreateSpeakerContext) error
	Delete(*DeleteSpeakerContext) error
	List(*ListSpeakerContext) error
	Show(*ShowSpeakerContext) error
	Update(*UpdateSpeakerContext) error
}

// MountSpeakerController "mounts" a Speaker resource controller on the given service.
func MountSpeakerController(service *goa.Service, ctrl SpeakerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID/events/:eventID/speakers", cors.HandlePreflight(service.Context, handleSpeakerOrigin))
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID", cors.HandlePreflight(service.Context, handleSpeakerOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateSpeakerContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateSpeakerPayload)
		}
		return ctrl.Create(rctx)
	}
	h = handleSpeakerOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("POST", "/api/tenants/:tenantID/events/:eventID/speakers", ctrl.MuxHandler("Create", h, unmarshalCreateSpeakerPayload))
	service.LogInfo("mount", "ctrl", "Speaker", "action", "Create", "route", "POST /api/tenants/:tenantID/events/:eventID/speakers", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteSpeakerContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSpeakerOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("DELETE", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Speaker", "action", "Delete", "route", "DELETE /api/tenants/:tenantID/events/:eventID/speakers/:speakerID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListSpeakerContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSpeakerOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID/events/:eventID/speakers", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Speaker", "action", "List", "route", "GET /api/tenants/:tenantID/events/:eventID/speakers", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowSpeakerContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSpeakerOrigin(h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Speaker", "action", "Show", "route", "GET /api/tenants/:tenantID/events/:eventID/speakers/:speakerID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateSpeakerContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateSpeakerPayload)
		}
		return ctrl.Update(rctx)
	}
	h = handleSpeakerOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/tenants/:tenantID/events/:eventID/speakers/:speakerID", ctrl.MuxHandler("Update", h, unmarshalUpdateSpeakerPayload))
	service.LogInfo("mount", "ctrl", "Speaker", "action", "Update", "route", "PATCH /api/tenants/:tenantID/events/:eventID/speakers/:speakerID", "security", "jwt")
}

// handleSpeakerOrigin applies the CORS response headers corresponding to the origin.
func handleSpeakerOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateSpeakerPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateSpeakerPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createSpeakerPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateSpeakerPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateSpeakerPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateSpeakerPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// TenantController is the controller interface for the Tenant actions.
type TenantController interface {
	goa.Muxer
	Create(*CreateTenantContext) error
	Delete(*DeleteTenantContext) error
	List(*ListTenantContext) error
	Show(*ShowTenantContext) error
	Update(*UpdateTenantContext) error
}

// MountTenantController "mounts" a Tenant resource controller on the given service.
func MountTenantController(service *goa.Service, ctrl TenantController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/tenants", cors.HandlePreflight(service.Context, handleTenantOrigin))
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID", cors.HandlePreflight(service.Context, handleTenantOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateTenantContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateTenantPayload)
		}
		return ctrl.Create(rctx)
	}
	h = handleTenantOrigin(h)
	service.Mux.Handle("POST", "/api/tenants", ctrl.MuxHandler("Create", h, unmarshalCreateTenantPayload))
	service.LogInfo("mount", "ctrl", "Tenant", "action", "Create", "route", "POST /api/tenants")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteTenantContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleTenantOrigin(h)
	service.Mux.Handle("DELETE", "/api/tenants/:tenantID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Tenant", "action", "Delete", "route", "DELETE /api/tenants/:tenantID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListTenantContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleTenantOrigin(h)
	service.Mux.Handle("GET", "/api/tenants", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Tenant", "action", "List", "route", "GET /api/tenants")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowTenantContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleTenantOrigin(h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Tenant", "action", "Show", "route", "GET /api/tenants/:tenantID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateTenantContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateTenantPayload)
		}
		return ctrl.Update(rctx)
	}
	h = handleTenantOrigin(h)
	service.Mux.Handle("PATCH", "/api/tenants/:tenantID", ctrl.MuxHandler("Update", h, unmarshalUpdateTenantPayload))
	service.LogInfo("mount", "ctrl", "Tenant", "action", "Update", "route", "PATCH /api/tenants/:tenantID")
}

// handleTenantOrigin applies the CORS response headers corresponding to the origin.
func handleTenantOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateTenantPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateTenantPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createTenantPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateTenantPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateTenantPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateTenantPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// UIController is the controller interface for the UI actions.
type UIController interface {
	goa.Muxer
	Bootstrap(*BootstrapUIContext) error
}

// MountUIController "mounts" a UI resource controller on the given service.
func MountUIController(service *goa.Service, ctrl UIController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/", cors.HandlePreflight(service.Context, handleUIOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewBootstrapUIContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Bootstrap(rctx)
	}
	h = handleUIOrigin(h)
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("Bootstrap", h, nil))
	service.LogInfo("mount", "ctrl", "UI", "action", "Bootstrap", "route", "GET /")
}

// handleUIOrigin applies the CORS response headers corresponding to the origin.
func handleUIOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Create(*CreateUserContext) error
	Delete(*DeleteUserContext) error
	List(*ListUserContext) error
	Show(*ShowUserContext) error
	Update(*UpdateUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID/users", cors.HandlePreflight(service.Context, handleUserOrigin))
	service.Mux.Handle("OPTIONS", "/api/tenants/:tenantID/users/:userID", cors.HandlePreflight(service.Context, handleUserOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateUserContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateUserPayload)
		}
		return ctrl.Create(rctx)
	}
	h = handleUserOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("POST", "/api/tenants/:tenantID/users", ctrl.MuxHandler("Create", h, unmarshalCreateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Create", "route", "POST /api/tenants/:tenantID/users", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleUserOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("DELETE", "/api/tenants/:tenantID/users/:userID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Delete", "route", "DELETE /api/tenants/:tenantID/users/:userID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleUserOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID/users", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "List", "route", "GET /api/tenants/:tenantID/users", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleUserOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID/users/:userID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Show", "route", "GET /api/tenants/:tenantID/users/:userID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateUserContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUserPayload)
		}
		return ctrl.Update(rctx)
	}
	h = handleUserOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/tenants/:tenantID/users/:userID", ctrl.MuxHandler("Update", h, unmarshalUpdateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Update", "route", "PATCH /api/tenants/:tenantID/users/:userID", "security", "jwt")
}

// handleUserOrigin applies the CORS response headers corresponding to the origin.
func handleUserOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// ValidateController is the controller interface for the Validate actions.
type ValidateController interface {
	goa.Muxer
	Validate(*ValidateValidateContext) error
}

// MountValidateController "mounts" a Validate resource controller on the given service.
func MountValidateController(service *goa.Service, ctrl ValidateController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/validate/:userID", cors.HandlePreflight(service.Context, handleValidateOrigin))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewValidateValidateContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Validate(rctx)
	}
	h = handleValidateOrigin(h)
	service.Mux.Handle("GET", "/api/validate/:userID", ctrl.MuxHandler("Validate", h, nil))
	service.LogInfo("mount", "ctrl", "Validate", "action", "Validate", "route", "GET /api/validate/:userID")
}

// handleValidateOrigin applies the CORS response headers corresponding to the origin.
func handleValidateOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:5000") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
