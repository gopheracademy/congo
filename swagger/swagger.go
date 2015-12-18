//************************************************************************//
// congo Swagger Spec
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package swagger

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// MountController mounts the swagger spec controller under "/swagger.json".
func MountController(service goa.Service) {
	ctrl := service.NewController("Swagger")
	service.Info("mount", "ctrl", "Swagger", "action", "Show", "route", "GET /swagger.json")
	h := ctrl.NewHTTPRouterHandle("Show", getSwagger)
	service.HTTPHandler().(*httprouter.Router).Handle("GET", "/swagger.json", h)
}

// getSwagger is the httprouter handle that returns the Swagger spec.
// func getSwagger(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func getSwagger(ctx *goa.Context) error {
	ctx.Header().Set("Content-Type", "application/swagger+json")
	ctx.Header().Set("Cache-Control", "public, max-age=3600")
	return ctx.Respond(200, []byte(spec))
}

// Generated spec
const spec = `{"swagger":"2.0","info":{"title":"Congo - Conference Management Made Easy","description":"Multi-tenant conference management application","contact":{"name":"Gopher Academy","email":"social@gopheracademy.com","url":"https://gopheracademy.com"},"license":{"name":"MIT","url":"https://github.com/gopheracademy/congo/blob/master/LICENSE"},"version":""},"host":"api.gopheracademy.com","basePath":"/api","schemes":["http"],"consumes":["application/json"],"produces":["application/json"],"paths":{"/users":{"get":{"description":"List all users in account","operationId":"user#list","consumes":["application/json"],"produces":["application/json"],"responses":{"200":{"description":""}},"schemes":["https"]},"post":{"description":"Record new user","operationId":"user#create","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/CreateUserPayload"}}],"responses":{"201":{"description":"Resource created","headers":{"Location":{"description":"href to created resource","type":"string","pattern":"^/accounts/[0-9]+/users/[0-9]+$"}}}},"schemes":["https"]}},"/users/{userID}":{"get":{"description":"Retrieve user with given id","operationId":"user#show","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"userID","in":"path","required":true,"type":"integer"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/User"}},"404":{"description":""}},"schemes":["https"]},"delete":{"operationId":"user#delete","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"userID","in":"path","description":"User ID","required":true,"type":"integer"}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]},"patch":{"operationId":"user#update","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"userID","in":"path","required":true,"type":"integer"},{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/UpdateUserPayload"}}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]}}},"definitions":{"CreateUserPayload":{"title":"CreateUserPayload","type":"object","properties":{"email":{"type":"string","minLength":2},"firstname":{"type":"string","minLength":2},"lastname":{"type":"string","minLength":2},"role":{"type":"string"}},"required":["email"]},"UpdateUserPayload":{"title":"UpdateUserPayload","type":"object","properties":{"email":{"type":"string","minLength":2},"firstname":{"type":"string","minLength":2},"lastname":{"type":"string","minLength":2},"role":{"type":"string"}}},"User":{"title":"Mediatype identifier: application/vnd.user+json","type":"object","properties":{"email":{"type":"string","description":"Email address of user","format":"email","minLength":2},"firstname":{"type":"string","description":"First name of user","minLength":2},"href":{"type":"string","description":"API href of user"},"id":{"type":"integer","description":"ID of user"},"lastname":{"type":"string","description":"Last name of user","minLength":2}},"description":"A user belonging to a tenant account"}},"externalDocs":{"description":"congo guide","url":"https://gopheracademy.com/congo/getting-started.html"}} `
