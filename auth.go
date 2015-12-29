package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"

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
