package main

import (
	"html/template"

	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
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

// UIController implements the ui resource.
type UIController struct {
	*goa.Controller
}

// NewUIController creates a ui controller.
func NewUIController(service *goa.Service) *UIController {
	return &UIController{Controller: service.NewController("UIController")}
}

// Bootstrap runs the bootstrap action.
func (c *UIController) Bootstrap(ctx *app.BootstrapUIContext) error {
	return ctx.OK([]byte(indexT))
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
		<link href="/assets/css/bootstrap.min.css" rel="stylesheet">

	</head>
	<body>
		<div id="app-container" data-auth='blue'></div>
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
		<link href="/assets/css/bootstrap.min.css" rel="stylesheet">

		<!-- Custom styles for this template -->
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
						<h3>To Proceed, Please Sign In </h3>
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
