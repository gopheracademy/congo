package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/jwt"
	"github.com/gopheracademy/congo/models"
)

// AuthController manages authentication
type AuthController struct {
	*goa.Controller
	e    *Env
	tm   *jwt.TokenManager
	spec *jwt.Specification
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, e *Env, tm *jwt.TokenManager, spec *jwt.Specification) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController"),
		e:    e,
		tm:   tm,
		spec: spec,
	}
}

// Refresh runs the refresh action.
func (c *AuthController) Refresh(ctx *app.RefreshAuthContext) error {
	// TBD: implement
	return nil
}

// Token runs the token action.
func (c *AuthController) Token(ctx *app.TokenAuthContext) error {
	// authenticate
	u, p, _ := ctx.Request.BasicAuth()

	goa.LogInfo(ctx, "token auth", "user", u)
	login := models.Login{
		Email:    u,
		Password: p,
	}

	userdb := models.NewUserDB(c.e.DB)
	user, err := userdb.GetByLogin(ctx, login)
	if err != nil {
		goa.LogError(ctx, "find user by login", "error", err.Error())
		return err
	}
	tenantdb := models.NewTenantDB(c.e.DB)
	var tenant models.Tenant
	var e error
	if user.TenantID > 0 {
		tenant, e = tenantdb.Get(ctx, user.TenantID)
		if e != nil {
			goa.LogError(ctx, "find tenant by id", "error", e.Error())
			return err
		}
	}
	// create token
	claims := make(map[string]interface{})
	claims["sub"] = user.ID
	claims["role"] = user.Role

	t, err := c.tm.Create(claims)
	// return token
	a := &app.Authorize{}
	tt := "Bearer"
	a.TokenType = &tt
	a.AccessToken = &t
	ttl := c.spec.TTLMinutes * 60
	a.ExpiresIn = &ttl
	a.User = user.UserToUser()
	a.Tenant = tenant.TenantToTenant()
	return ctx.Created(a)
}
