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
	if auth == nil {
		return ctx.OK([]byte(loginT))
	}
	bs, err := json.Marshal(auth)
	if err != nil {
		return err
	}
	js := string(bs)
	var out bytes.Buffer
	err = indexTmpl.Execute(&out, js)
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
				AccessToken: token.Raw,
				ExpiresIn:   time.Now().Second() - u.Oauth2Expiry.Second(),
				TokenType:   "Bearer",
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
		<div id="app-container" data-jwt='{{.}}'></div>
		<script src="/assets/app.js"></script>
	</body>
</html>
`

// loginT is the template used to render the login page
const loginT = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8"/>
		<link rel="stylesheet" type="text/css" href="/assets/bootstrap.min.css">
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/0.14.4/react.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/0.14.4/react-dom.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.23/browser.min.js"></script>
	</head>
	<body>
		<div id="app-container"></div>
		<script type="text/babel">
		var Login = React.createClass({
			render: function() {
				return (
					<div>
						<h3>Sign In With</h3>
						<a className="btn btn-twitter" href="/api/auth/twitter"><i className="icon-twitter"></i> | Twitter </a>
						<a className="btn btn-github" href="/api/auth/github"><i className="icon-github"></i> | GitHub </a>
					</div>
				);
			}
		});
		ReactDOM.render(
			<Login />,
			document.getElementById("app-container")
		);
		</script>
	</body>
</html>
`
