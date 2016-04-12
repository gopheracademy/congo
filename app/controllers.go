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
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/admin/users/:userID", ctrl.MuxHandler("Update", h, unmarshalUpdateAdminuserPayload))
	service.LogInfo("mount", "ctrl", "Adminuser", "action", "Update", "route", "PATCH /api/admin/users/:userID", "security", "jwt")
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

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewRefreshAuthContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Refresh(rctx)
	}
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
	h = handleSecurity("password", h)
	service.Mux.Handle("POST", "/api/auth/token", ctrl.MuxHandler("Token", h, nil))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Token", "route", "POST /api/auth/token", "security", "password")
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
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/tenants/:tenantID/events/:eventID", ctrl.MuxHandler("Update", h, unmarshalUpdateEventPayload))
	service.LogInfo("mount", "ctrl", "Event", "action", "Update", "route", "PATCH /api/tenants/:tenantID/events/:eventID", "security", "jwt")
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

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewStatusHealthzContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Status(rctx)
	}
	service.Mux.Handle("GET", "/api/healthz", ctrl.MuxHandler("Status", h, nil))
	service.LogInfo("mount", "ctrl", "Healthz", "action", "Status", "route", "GET /api/healthz")
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
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/tenants/:tenantID/series/:seriesID", ctrl.MuxHandler("Update", h, unmarshalUpdateSeriesPayload))
	service.LogInfo("mount", "ctrl", "Series", "action", "Update", "route", "PATCH /api/tenants/:tenantID/series/:seriesID", "security", "jwt")
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
	h = handleSecurity("jwt", h)
	service.Mux.Handle("POST", "/api/tenants", ctrl.MuxHandler("Create", h, unmarshalCreateTenantPayload))
	service.LogInfo("mount", "ctrl", "Tenant", "action", "Create", "route", "POST /api/tenants", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteTenantContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("jwt", h)
	service.Mux.Handle("DELETE", "/api/tenants/:tenantID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Tenant", "action", "Delete", "route", "DELETE /api/tenants/:tenantID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListTenantContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/tenants", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Tenant", "action", "List", "route", "GET /api/tenants", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowTenantContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/api/tenants/:tenantID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Tenant", "action", "Show", "route", "GET /api/tenants/:tenantID", "security", "jwt")

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
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/tenants/:tenantID", ctrl.MuxHandler("Update", h, unmarshalUpdateTenantPayload))
	service.LogInfo("mount", "ctrl", "Tenant", "action", "Update", "route", "PATCH /api/tenants/:tenantID", "security", "jwt")
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
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PATCH", "/api/tenants/:tenantID/users/:userID", ctrl.MuxHandler("Update", h, unmarshalUpdateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Update", "route", "PATCH /api/tenants/:tenantID/users/:userID", "security", "jwt")
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

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewValidateValidateContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Validate(rctx)
	}
	service.Mux.Handle("GET", "/api/validate/:userID", ctrl.MuxHandler("Validate", h, nil))
	service.LogInfo("mount", "ctrl", "Validate", "action", "Validate", "route", "GET /api/validate/:userID")
}
