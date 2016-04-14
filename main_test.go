package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	jg "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	mjwt "github.com/goadesign/goa/middleware/security/jwt"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/jwt"
	"github.com/gopheracademy/congo/swagger"
	"github.com/kelseyhightower/envconfig"
	"github.com/kr/pretty"
	_ "github.com/lib/pq"
)

var reader io.Reader

func TestMain(m *testing.M) {

	err := envconfig.Process("congotest", &settings)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	env, err := envFromSettings()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	Migrate(env)

	storageList := makeStorage(env)
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	spec := &jwt.Specification{
		AllowParam:       false,
		AuthOptions:      false,
		TTLMinutes:       60,
		Issuer:           "congo.gopheracademy.com",
		KeySigningMethod: jwt.RSA256,
		SigningKeyFunc:   privateKey,
		ValidationFunc:   pubKey,
	}

	tm := jwt.NewTokenManager(spec)

	pubkey, err := readPubKey()
	if err != nil {
		panic(err)
	}
	pk, err := jg.ParseRSAPublicKeyFromPEM(pubkey)
	if err != nil {
		panic(err)
	}
	app.ConfigureJWTSecurity(service, mjwt.New(pk, nil))
	app.ConfigurePasswordSecurity(service, BasicAuth(env))
	// Mount "adminuser" controller
	c := NewAdminuserController(service)
	app.MountAdminuserController(service, c)
	// Mount "auth" controller
	c2 := NewAuthController(service, env, tm, spec)
	app.MountAuthController(service, c2)
	// Mount "event" controller
	c3 := NewEventController(service, storageList, env)
	app.MountEventController(service, c3)
	// Mount "healthz" controller
	c4 := NewHealthzController(service)
	app.MountHealthzController(service, c4)
	// Mount "series" controller
	c5 := NewSeriesController(service, storageList, env)
	app.MountSeriesController(service, c5)
	// Mount "tenant" controller
	c6 := NewTenantController(service, storageList, env)
	app.MountTenantController(service, c6)
	// Mount "user" controller
	c7 := NewUserController(service, storageList, env)
	app.MountUserController(service, c7)
	// Mount "validate" controller
	c8 := NewValidateController(service)
	app.MountValidateController(service, c8)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	go service.ListenAndServe(":18081")

	status := m.Run()
	service.CancelAll()
	os.Exit(status)
}

func TestTenantsNoAuth(t *testing.T) {
	tenantsUrl := "http://localhost:18081/api/tenants"

	request, err := http.NewRequest("GET", tenantsUrl, reader) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode != 401 {
		t.Errorf("expected: %d got %d", 401, res.StatusCode) //Uh-oh this means our test failed
	}
}

func TestTenantsWithAuth(t *testing.T) {
	token, err := getToken()
	if err != nil {
		t.Error(err)
	}
	tenantsUrl := "http://localhost:18081/api/tenants"

	request, err := http.NewRequest("GET", tenantsUrl, reader) //Create request with JSON body
	request.Header.Add("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode != 200 {
		t.Errorf("expected: %d got %d", 200, res.StatusCode) //Uh-oh this means our test failed
	}

	var a app.TenantCollection
	err = json.NewDecoder(res.Body).Decode(&a)
	if err != nil {
		t.Error(err)
	}

	if len(a) < 1 {
		t.Errorf("expected a tenant, got none")
	}
	pretty.Println(a)

}

func TestAuth(t *testing.T) {

	token, err := getToken()

	if err != nil {
		t.Error(err)
	}
	if len(token) < 1 {
		t.Error("Expected a token, got %s", token)
	}
	fmt.Println(token)

}

func getToken() (string, error) {

	authurl := "http://localhost:18081/api/auth/token"

	request, err := http.NewRequest("POST", authurl, reader) //Create request with JSON body
	request.SetBasicAuth("bketelsen@gmail.com", "GopherCon")
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		return "", err
	}

	if res.StatusCode != 201 {
		return "", err
	}

	defer res.Body.Close()
	var a app.Authorize
	err = json.NewDecoder(res.Body).Decode(&a)
	return *a.AccessToken, err

}
