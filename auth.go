package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"
	"github.com/markbates/goth/gothic"

	djwt "github.com/dgrijalva/jwt-go"
	"github.com/raphael/goa"
	"github.com/raphael/goa-middleware/jwt"
)

// AuthController implements the auth resource.
type AuthController struct {
	goa.Controller
	db *gorm.DB

	tm   *jwt.TokenManager
	spec *jwt.Specification
}

// NewAuthController creates a auth controller.
func NewAuthController(service goa.Service, db *gorm.DB, tm *jwt.TokenManager, spec *jwt.Specification) app.AuthController {
	return &AuthController{Controller: service.NewController("AuthController"),
		tm:   tm,
		db:   db,
		spec: spec,
	}
}

// Callback implements the endpoint called by the service during the oauth flow.
func (c *AuthController) Callback(ctx *app.CallbackAuthContext) error {

	user, err := gothic.CompleteUserAuth(ctx, ctx.Request())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(user)

	udb := models.NewUserDB(*c.db)
	prov, _ := provider(ctx.Request())
	var cuser models.User
	var admin bool
	cuser, err = udb.UserByOauth(user.UserID, prov)
	if err != nil {
		if err.Error() == "record not found" { // todo: get a real error here?
			cuser = models.User{
				// OAuth2
				Oauth2Uid:      user.UserID,
				Oauth2Provider: prov,
				Oauth2Token:    user.AccessToken,

				Email: user.Email,
				Role:  models.USER,
			}
			cuser, err = udb.Add(context.Background(), cuser)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(cuser.ID)
		}
	} else {
		fmt.Println("updating existing user")
		cuser.Oauth2Token = user.AccessToken
		err := udb.Update(context.Background(), cuser)
		if err != nil {
			fmt.Println("Update Error: ", err)
		}

	}
	if cuser.Role == models.ADMIN {
		admin = true
	}
	claims := make(map[string]interface{})
	claims["sub"] = strconv.Itoa(cuser.ID)
	claims["provider"] = prov
	claims["first_name"] = cuser.Firstname
	t, err := c.tm.Create(claims)
	auth := app.Authorize{}
	auth.AccessToken = t
	auth.ExpiresIn = 60 // TBD extract from auth response raw data?

	return RenderBootstrap(ctx, cuser.ID, admin, &auth)
}

func (c *AuthController) Oauth(ctx *app.OauthAuthContext) error {
	url, err := gothic.GetAuthURL(ctx, ctx.Request())
	fmt.Println(url)
	if err != nil {
		ctx.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(ctx, err)
		return err
	}

	http.Redirect(ctx, ctx.Request(), url, http.StatusTemporaryRedirect)
	return nil
}

// Token runs the token action.
func (c *AuthController) Token(ctx *app.TokenAuthContext) error {

	// authenticate
	login := models.Login{
		Email:    ctx.Payload.Email,
		Password: ctx.Payload.Password,
	}
	fmt.Println(login)
	userdb := models.NewUserDB(*c.db)
	user, err := userdb.GetByLogin(ctx, login)
	if err != nil {
		//Logger(m.e, ctx).Error(err)
		return ctx.Respond(401, []byte("Unauthorized"))
	}
	// create token
	claims := make(map[string]interface{})
	claims["sub"] = user.ID
	claims["role"] = user.Role

	t, err := c.tm.Create(claims)
	if err != nil {
		return err
	}
	// return token
	a := &app.Authorize{}
	a.TokenType = "Bearer"
	a.AccessToken = t
	a.ExpiresIn = c.spec.TTLMinutes * 60
	return ctx.Created(a)
}

// Refresh runs the refresh action.
func (c *AuthController) Refresh(ctx *app.RefreshAuthContext) error {
	// validate existing token

	var found bool
	var token string
	header := ctx.Request().Header.Get(c.spec.TokenHeader)

	if header != "" {
		parts := strings.Split(header, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			// This is an error
		}
		token = parts[1]
		found = true
	}
	if !found && c.spec.AllowParam {
		token = ctx.Request().URL.Query().Get(c.spec.TokenParam)

	}

	if token == "" {
		err := ctx.Respond(http.StatusUnauthorized, []byte(http.StatusText(http.StatusUnauthorized)))
		return err
	}
	parsed, err := djwt.Parse(token, keyFuncWrapper(c.spec.ValidationFunc))
	if err != nil {
		msg := fmt.Sprintf("Error parsing token: %s", err.Error())
		err = ctx.Respond(http.StatusUnauthorized, []byte(msg))
		return err
	}
	if parsed.Valid {
		ctx.SetValue(jwt.JWTKey, parsed)
	} else {
		msg := "Invalid Token"
		err = ctx.Respond(http.StatusUnauthorized, []byte(msg))
		return err
	}
	// authenticate
	login := models.Login{
		Email:    ctx.Payload.Email,
		Password: ctx.Payload.Password,
	}

	userdb := models.NewUserDB(*c.db)
	user, err := userdb.GetByLogin(ctx, login)
	if err != nil {
		//Logger(m.e, ctx).Error(err)
		return ctx.Respond(401, []byte("Unauthorized"))
	}
	// create token
	claims := make(map[string]interface{})
	claims["sub"] = user.ID
	claims["role"] = user.Role

	t, err := c.tm.Create(claims)
	// return token
	a := &app.Authorize{}
	a.TokenType = "Bearer"
	a.AccessToken = t
	a.ExpiresIn = c.spec.TTLMinutes * 60
	return ctx.Created(a)
}

func keyFuncWrapper(k jwt.ValidationKeyfunc) djwt.Keyfunc {
	return func(tok *djwt.Token) (interface{}, error) {
		return k(tok)
	}
}
