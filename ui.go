package main

import (
	"bytes"
	"encoding/json"
	"text/template"
	"time"

	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"
	"github.com/raphael/goa"
	"github.com/raphael/goa-middleware/jwt"
)

// Template used to render bootstrap page.
var indexTmpl *template.Template

// Initialize template
func init() {
	tmpl, err := template.New("index").Parse(indexT)
	if err != nil {
		panic("index template: " + err.Error())
	}
	indexTmpl = tmpl
}

// UIController implements the UI bootstrap action.
type UIController struct {
	goa.Controller
	db *gorm.DB
}

// NewUIController initializes a new UI controller.
func NewUIController(service goa.Service, db *gorm.DB) *UIController {
	return &UIController{
		Controller: service.NewController("ui"),
		db:         db,
	}

}

type okCtx interface {
	OK([]byte) error
}

// RenderBootstrap writes the bootstrap HTML for the given user using the context OK function.
func RenderBootstrap(ctx okCtx, auth *app.Authorize) error {
	var js string
	if auth != nil {
		bs, err := json.Marshal(auth)
		if err != nil {
			return err
		}
		js = string(bs)
	}
	var out bytes.Buffer
	err := indexTmpl.Execute(&out, js)
	if err != nil {
		return err
	}
	return ctx.OK(out.Bytes())
}

// Bootstrap renders the index.html.
func (c *UIController) Bootstrap(ctx *app.BootstrapUiContext) error {
	var auth *app.Authorize
	token, err := jwt.GetToken(ctx.Context, jwtSpec)
	if err == nil {
		userdb := models.NewUserDB(*c.db)
		u, err := userdb.One(ctx, token.Claims["sub"].(int))
		if err != nil {
			auth = &app.Authorize{
				AccessToken: u.Oauth2Token,
				ExpiresIn:   time.Now().Second() - u.Oauth2Expiry.Second(),
				TokenType:   "",
			}
		}
	}
	return RenderBootstrap(ctx, auth)
}

// indexT is the template used to render the index page
const indexT = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8"/>
    <link rel="stylesheet" type="text/css" href="/assets/bootstrap.min.css">
  </head>
  <body>
    <div id="app-container"></div>
    <script src="/assets/app.js"></script>
    <script type="text/babel">
      {{if .}}var AuthedMaster = React.createClass({
          render() {
            return <Master auth={{.}} />;
          }
        });
{{end}}
      ReactDOM.render(
        <Router history={hashHistory}>
          <Route path="/" component={{if .}}{AuthedMaster}{{else}}{Login}{{end}}>
            <Route path="/profile" component={Profile} />
          </Route>
        </Router>,
        document.getElementById("app-container")
      );
    </script>
  </body>
</html>
`
