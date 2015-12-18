//************************************************************************//
// congo JSON Hyper-schema
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
// --url=http://localhost
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package schema

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// MountController mounts the API JSON schema controller under "/schema.json".
func MountController(service goa.Service) {
	ctrl := service.NewController("Schema")
	service.Info("mount", "ctrl", "Schema", "action", "Show", "route", "GET /schema.json")
	h := ctrl.NewHTTPRouterHandle("Show", getSchema)
	service.HTTPHandler().(*httprouter.Router).Handle("GET", "/schema.json", h)
}

// getSchema is the httprouter handle that returns the API JSON hyper schema.
// func getSchema(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func getSchema(ctx *goa.Context) error {
	ctx.Header().Set("Content-Type", "application/schema+json")
	ctx.Header().Set("Cache-Control", "public, max-age=3600")
	return ctx.Respond(200, []byte(schema))
}

// Generated schema
const schema = `{"$schema":"http://json-schema.org/draft-04/hyper-schema","id":"http://localhost/schema","title":"Congo - Conference Management Made Easy","type":"object","properties":{"CreateUserPayload":{"$ref":"#/definitions/CreateUserPayload"},"UpdateUserPayload":{"$ref":"#/definitions/UpdateUserPayload"},"User":{"$ref":"#/definitions/User"},"user":{"$ref":"#/definitions/user"}},"definitions":{"CreateUserPayload":{"title":"CreateUserPayload","type":"object","properties":{"email":{"type":"string","minLength":2},"firstname":{"type":"string","minLength":2},"lastname":{"type":"string","minLength":2},"role":{"type":"string"}},"required":["email"]},"UpdateUserPayload":{"title":"UpdateUserPayload","type":"object","properties":{"email":{"type":"string","minLength":2},"firstname":{"type":"string","minLength":2},"lastname":{"type":"string","minLength":2},"role":{"type":"string"}}},"User":{"title":"Mediatype identifier: application/vnd.user+json","type":"object","properties":{"email":{"type":"string","description":"Email address of user","format":"email","minLength":2},"firstname":{"type":"string","description":"First name of user","minLength":2},"href":{"type":"string","description":"API href of user"},"id":{"type":"integer","description":"ID of user"},"lastname":{"type":"string","description":"Last name of user","minLength":2}},"description":"A user belonging to a tenant account","media":{"type":"application/vnd.user+json"}},"user":{"title":"user","type":"object","properties":{"email":{"type":"string","description":"Email address of user","format":"email","minLength":2},"firstname":{"type":"string","description":"First name of user","minLength":2},"href":{"type":"string","description":"API href of user"},"id":{"type":"integer","description":"ID of user"},"lastname":{"type":"string","description":"Last name of user","minLength":2}},"description":"A user belonging to a tenant account","media":{"type":"application/vnd.user+json"},"links":[{"title":"create","rel":"create","href":"/api/users","method":"POST","schema":{"description":"create payload","$ref":"#/definitions/CreateUserPayload"}},{"title":"delete","rel":"delete","href":"/api/users/{userID}","method":"DELETE"},{"title":"list","rel":"list","href":"/api/users","method":"GET"},{"title":"show","rel":"self","href":"/api/users/{userID}","method":"GET","targetSchema":{"$ref":"#/definitions/User"},"mediaType":"application/vnd.user+json"},{"title":"update","rel":"update","href":"/api/users/{userID}","method":"PATCH","schema":{"description":"update payload","$ref":"#/definitions/UpdateUserPayload"}}]}},"description":"Multi-tenant conference management application","links":[{"rel":"self","href":"http://localhost"},{"rel":"self","href":"/schema","method":"GET","targetSchema":{"$schema":"http://json-schema.org/draft-04/hyper-schema","additionalProperties":true}}]}`
