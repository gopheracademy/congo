//************************************************************************//
// congo JSON Hyper-schema
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=/home/bketelsen/src/github.com/bketelsen/congo
// --design=github.com/bketelsen/congo/design
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
	service.Info("mount", "ctrl", "Schema", "action", "Show", "route", "GET /schema.json")
	h := goa.NewHTTPRouterHandle(service, "Schema", "Show", getSchema)
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
const schema = `{"$schema":"http://json-schema.org/draft-04/hyper-schema","id":"http://localhost/schema","title":"Congo - Conference Management Made Easy","type":"object","properties":{"Account":{"$ref":"#/definitions/Account"},"CreateAccountPayload":{"$ref":"#/definitions/CreateAccountPayload"},"CreateSeriesPayload":{"$ref":"#/definitions/CreateSeriesPayload"},"Series":{"$ref":"#/definitions/Series"},"SeriesCollection":{"$ref":"#/definitions/SeriesCollection"},"UpdateAccountPayload":{"$ref":"#/definitions/UpdateAccountPayload"},"UpdateSeriesPayload":{"$ref":"#/definitions/UpdateSeriesPayload"},"account":{"$ref":"#/definitions/account"},"series":{"$ref":"#/definitions/series"}},"definitions":{"Account":{"title":"Mediatype identifier: application/vnd.congo.api.account","type":"object","properties":{"created_at":{"type":"string","description":"Date of creation","format":"date-time"},"created_by":{"type":"string","description":"Email of account owner","format":"email"},"href":{"type":"string","description":"API href of account"},"id":{"type":"integer","description":"ID of account"},"name":{"type":"string","description":"Name of account"}},"description":"A tenant account","media":{"type":"application/vnd.congo.api.account"}},"CreateAccountPayload":{"title":"CreateAccountPayload","type":"object","properties":{"name":{"type":"string","description":"Name of account"}},"required":["name"]},"CreateSeriesPayload":{"title":"CreateSeriesPayload","type":"object","properties":{"name":{"type":"string","minLength":2}},"required":["name"]},"Series":{"title":"Mediatype identifier: application/vnd.congo.api.series","type":"object","properties":{"account":{"description":"Account that owns bottle","$ref":"#/definitions/Account"},"created_at":{"type":"string","description":"Date of creation","format":"date-time"},"href":{"type":"string","description":"API href of series"},"id":{"type":"integer","description":"ID of series"},"name":{"type":"string","minLength":2},"updated_at":{"type":"string","description":"Date of last update","format":"date-time"}},"description":"A recurring event or conference","media":{"type":"application/vnd.congo.api.series"},"links":[{"title":"account","description":"Account that owns bottle","rel":"account","href":"/congo/accounts/{accountID}","method":"GET","targetSchema":{"$ref":"#/definitions/Account"},"mediaType":"application/vnd.congo.api.account"}]},"SeriesCollection":{"title":"Mediatype identifier: application/vnd.congo.api.series; type=collection","type":"array","items":{"$ref":"#/definitions/Series"},"media":{"type":"application/vnd.congo.api.series; type=collection"}},"UpdateAccountPayload":{"title":"UpdateAccountPayload","type":"object","properties":{"name":{"type":"string","description":"Name of account"}},"required":["name"]},"UpdateSeriesPayload":{"title":"UpdateSeriesPayload","type":"object","properties":{"name":{"type":"string","minLength":2}}},"account":{"title":"account","type":"object","properties":{"created_at":{"type":"string","description":"Date of creation","format":"date-time"},"created_by":{"type":"string","description":"Email of account owner","format":"email"},"href":{"type":"string","description":"API href of account"},"id":{"type":"integer","description":"ID of account"},"name":{"type":"string","description":"Name of account"}},"description":"A tenant account","media":{"type":"application/vnd.congo.api.account"},"links":[{"title":"show","rel":"self","href":"/congo/accounts/{accountID}","method":"GET","targetSchema":{"$ref":"#/definitions/Account"},"mediaType":"application/vnd.congo.api.account"},{"title":"create","rel":"create","href":"/congo/accounts","method":"POST","schema":{"description":"create payload","$ref":"#/definitions/CreateAccountPayload"}},{"title":"update","rel":"update","href":"/congo/accounts/{accountID}","method":"PUT","schema":{"description":"update payload","$ref":"#/definitions/UpdateAccountPayload"}},{"title":"delete","rel":"delete","href":"/congo/accounts/{accountID}","method":"DELETE"}]},"series":{"title":"series","type":"object","properties":{"account":{"description":"Account that owns bottle","$ref":"#/definitions/Account"},"created_at":{"type":"string","description":"Date of creation","format":"date-time"},"href":{"type":"string","description":"API href of series"},"id":{"type":"integer","description":"ID of series"},"name":{"type":"string","minLength":2},"updated_at":{"type":"string","description":"Date of last update","format":"date-time"}},"description":"A recurring event or conference","media":{"type":"application/vnd.congo.api.series"},"links":[{"title":"account","description":"Account that owns bottle","rel":"account","href":"/congo/accounts/{accountID}","method":"GET","targetSchema":{"$ref":"#/definitions/Account"},"mediaType":"application/vnd.congo.api.account"},{"title":"list","rel":"list","href":"/congo/accounts/{accountID}/series","method":"GET","targetSchema":{"$ref":"#/definitions/SeriesCollection"},"mediaType":"application/vnd.congo.api.series; type=collection"},{"title":"show","rel":"self","href":"/congo/accounts/{accountID}/series/{seriesID}","method":"GET","targetSchema":{"$ref":"#/definitions/Series"},"mediaType":"application/vnd.congo.api.series"},{"title":"create","rel":"create","href":"/congo/accounts/{accountID}/series","method":"POST","schema":{"description":"create payload","$ref":"#/definitions/CreateSeriesPayload"}},{"title":"update","rel":"update","href":"/congo/accounts/{accountID}/series/{seriesID}","method":"PATCH","schema":{"description":"update payload","$ref":"#/definitions/UpdateSeriesPayload"}}]}},"description":"Multi-tenant conference management application","links":[{"rel":"self","href":"http://localhost"},{"rel":"self","href":"/schema","method":"GET","targetSchema":{"$schema":"http://json-schema.org/draft-04/hyper-schema","additionalProperties":true}}]}`
