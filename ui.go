package main

import (
	"bytes"
	"encoding/json"
	"text/template"
	"time"

	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/gopheracademy/congo/models/user"
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
func RenderBootstrap(ctx okCtx, userID int, admin bool, auth *app.Authorize) error {
	if auth == nil {
		return ctx.OK([]byte(loginT))
	}
	data := struct {
		UserID int
		Admin  bool
		Auth   *app.Authorize
	}{
		UserID: userID,
		Admin:  admin,
		Auth:   auth,
	}
	bs, err := json.Marshal(data)
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
	var userID int
	var admin bool
	var auth *app.Authorize
	token, err := jwt.GetToken(ctx.Context, jwtSpec)
	if err == nil {
		userdb := user.NewUserDB(*c.db)
		userID = token.Claims["sub"].(int)
		u, err := userdb.One(ctx, userID)
		if err == nil {
			if u.Role == models.ADMIN {
				admin = true
			}
			auth = &app.Authorize{
				AccessToken: token.Raw,
				ExpiresIn:   time.Now().Second() - u.Oauth2Expiry.Second(),
				TokenType:   "Bearer",
			}
		}
	}
	return RenderBootstrap(ctx, userID, admin, auth)
}

// indexT is the template used to render the index page
const indexT = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta name="description" content="">
		<meta name="author" content="">

		 <title>Congo - CFP</title>

    <!-- Bootstrap core CSS -->
		<link href="/assets/css/bootstrap.css" rel="stylesheet">

    <!-- Custom styles for this template -->
		<link href="/assets/css/main.css" rel="stylesheet">
		<link href="/assets/css/font-awesome.min.css" rel="stylesheet">
	</head>
	<body>
		<div id="app-container" data-auth='{{.}}'></div>
		<script src="/assets/app.js"></script>
	</body>
</html>
`

// loginT is the template used to render the login page
const loginT = `
<!DOCTYPE html>
<html>
	<head>
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta name="description" content="">
		<meta name="author" content="">

		<title>Congo - CFP</title>

		<!-- Bootstrap core CSS -->
		<link href="/assets/css/bootstrap.css" rel="stylesheet">

		<!-- Custom styles for this template -->
		<link href="/assets/css/main.css" rel="stylesheet">
		<link href="/assets/css/font-awesome.min.css" rel="stylesheet">
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/0.14.4/react.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/0.14.4/react-dom.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.23/browser.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/require.js/2.1.16/require.min.js"></script>
	</head>
	<body>
		<div id="app-container"></div>
		<script type="text/babel">
		var Login = React.createClass({
			render() {
				return (
					<div className="container">
						<h2>Welcome to Congo Conference Management</h2>
						<h3>To Proceed, Please Sign In With</h3>
						<a className="btn btn-twitter" href="/api/auth/twitter"><i className="fa fa-twitter"></i> | Twitter </a>
						<a className="btn btn-github" href="/api/auth/github"><i className="fa fa-github"></i> | GitHub </a>
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
