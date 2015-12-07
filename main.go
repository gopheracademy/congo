package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/davecgh/go-spew/spew"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/handlers"
	"github.com/gopheracademy/congo/models"
	"github.com/gopheracademy/congo/swagger"
	"github.com/gopheracademy/congo/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/nosurf"
	"github.com/kelseyhightower/envconfig"
	"github.com/raphael/goa"
	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/confirm"
	_ "gopkg.in/authboss.v0/lock"
	_ "gopkg.in/authboss.v0/recover"
	_ "gopkg.in/authboss.v0/register"
	_ "gopkg.in/authboss.v0/remember"
)

func main() {

	cfg := util.Config{}
	if err := envconfig.Process("congo", &cfg); err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}

	env, err := cfg.Env()
	if err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}
	dev := env == util.EnvDev
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(goa.RequestID())
	service.Use(goa.LogRequest())
	service.Use(goa.Recover())

	// Mount "account" controller
	c := NewAccountController(service, models.NewMockAccountStorage())
	app.MountAccountController(service, c)
	// Mount "series" controller
	c2 := NewSeriesController(service, models.NewMockSeriesStorage())
	app.MountSeriesController(service, c2)
	// Mount "user" controller
	c3 := NewUserController(service, models.NewMockUserStorage())
	app.MountUserController(service, c3)

	// Mount "user" controller
	c4 := NewInstanceController(service, models.NewMockInstanceStorage())
	app.MountInstanceController(service, c4)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	goarouter := service.HTTPHandler().(*httprouter.Router)

	renderer := handlers.NewRenderRenderer("templates", []string{".html"}, handlers.Funcs, dev)
	database := NewMemStorer()
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`NpEPi8pEjKVjLGJ6kYCS+VTCzi6BUuDzU0wrwXyf5uDPArtlofn2AG6aTMiPmN3C909rsEWMNqJqhIVPGP3Exg==`)
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(`AbfYwmmt8UCwUuhd9qvfNA9UCuN1cVcKJN1ofbiky6xCyyBj20whe40rJa3Su0WOWLWcPpO1taqJdsEI/65+JA==`)
	cookieStore = securecookie.New(cookieStoreKey, nil)
	sessionStore = sessions.NewCookieStore(sessionStoreKey)
	ab := authboss.New() // Usually store this globally
	ab.MountPath = "/auth"
	ab.LogWriter = os.Stdout
	ab.Storer = database
	ab.CookieStoreMaker = NewCookieStorer
	ab.SessionStoreMaker = NewSessionStorer
	ab.XSRFName = "csrf_token"
	ab.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return nosurf.Token(r)
	}
	if err := ab.Init(); err != nil {
		// Handle error, don't let program continue to run
		log.Fatalln(err)
	}

	// Make sure to put authboss's router somewhere
	r := mux.NewRouter()
	// Make sure to put authboss's router somewhere
	r.Handle("/", handlers.Index(renderer)).Methods("GET")
	r.Handle("/accounts/{accountID}/users", handlers.Users(renderer)).Methods("GET")
	r.Handle("/accounts/{accountID}/series/{seriesID}/instances", handlers.Instances(renderer)).Methods("GET")
	r.Handle("/accounts/{accountID}/series", handlers.Series(renderer)).Methods("GET")
	r.Handle("/accounts", handlers.Accounts(renderer)).Methods("GET")

	r.PathPrefix("/api").Handler(goarouter)
	r.PathPrefix("/auth").Handler(ab.NewRouter())

	n := negroni.Classic()
	n.UseHandler(r)
	hostStr := fmt.Sprintf(":%d", cfg.Port)
	n.Run(hostStr)
}

var nextUserID uint

type MemStorer struct {
	Users  map[string]models.User
	Tokens map[string][]string
}

func NewMemStorer() *MemStorer {
	return &MemStorer{
		Users:  map[string]models.User{},
		Tokens: make(map[string][]string),
	}
}

func (s MemStorer) Create(key string, attr authboss.Attributes) error {
	var user models.User
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	user.ID = nextUserID
	nextUserID++

	s.Users[key] = user
	fmt.Println("Create")
	spew.Dump(s.Users)
	return nil
}

func (s MemStorer) Put(key string, attr authboss.Attributes) error {
	return s.Create(key, attr)
}

func (s MemStorer) Get(key string) (result interface{}, err error) {
	user, ok := s.Users[key]
	if !ok {
		return nil, authboss.ErrUserNotFound
	}

	return &user, nil
}

func (s MemStorer) PutOAuth(uid, provider string, attr authboss.Attributes) error {
	return s.Create(uid+provider, attr)
}

func (s MemStorer) GetOAuth(uid, provider string) (result interface{}, err error) {
	user, ok := s.Users[uid+provider]
	if !ok {
		return nil, authboss.ErrUserNotFound
	}

	return &user, nil
}

func (s MemStorer) AddToken(key, token string) error {
	s.Tokens[key] = append(s.Tokens[key], token)
	fmt.Println("AddToken")
	spew.Dump(s.Tokens)
	return nil
}

func (s MemStorer) DelTokens(key string) error {
	delete(s.Tokens, key)
	fmt.Println("DelTokens")
	spew.Dump(s.Tokens)
	return nil
}

func (s MemStorer) UseToken(givenKey, token string) error {
	toks, ok := s.Tokens[givenKey]
	if !ok {
		return authboss.ErrTokenNotFound
	}

	for i, tok := range toks {
		if tok == token {
			toks[i], toks[len(toks)-1] = toks[len(toks)-1], toks[i]
			s.Tokens[givenKey] = toks[:len(toks)-1]
			return nil
		}
	}

	return authboss.ErrTokenNotFound
}

func (s MemStorer) ConfirmUser(tok string) (result interface{}, err error) {
	fmt.Println("==============", tok)

	for _, u := range s.Users {
		if u.ConfirmToken == tok {
			return &u, nil
		}
	}

	return nil, authboss.ErrUserNotFound
}

func (s MemStorer) RecoverUser(rec string) (result interface{}, err error) {
	for _, u := range s.Users {
		if u.RecoverToken == rec {
			return &u, nil
		}
	}

	return nil, authboss.ErrUserNotFound
}

var cookieStore *securecookie.SecureCookie

type CookieStorer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewCookieStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &CookieStorer{w, r}
}

func (s CookieStorer) Get(key string) (string, bool) {
	cookie, err := s.r.Cookie(key)
	if err != nil {
		return "", false
	}

	var value string
	err = cookieStore.Decode(key, cookie.Value, &value)
	if err != nil {
		return "", false
	}

	return value, true
}

func (s CookieStorer) Put(key, value string) {
	encoded, err := cookieStore.Encode(key, value)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Expires: time.Now().UTC().AddDate(1, 0, 0),
		Name:    key,
		Value:   encoded,
		Path:    "/",
	}
	http.SetCookie(s.w, cookie)
}

func (s CookieStorer) Del(key string) {
	cookie := &http.Cookie{
		MaxAge: -1,
		Name:   key,
		Path:   "/",
	}
	http.SetCookie(s.w, cookie)
}

const sessionCookieName = "ab_blog"

var sessionStore *sessions.CookieStore

type SessionStorer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewSessionStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &SessionStorer{w, r}
}

func (s SessionStorer) Get(key string) (string, bool) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	strInf, ok := session.Values[key]
	if !ok {
		return "", false
	}

	str, ok := strInf.(string)
	if !ok {
		return "", false
	}

	return str, true
}

func (s SessionStorer) Put(key, value string) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.Values[key] = value
	session.Save(s.r, s.w)
}

func (s SessionStorer) Del(key string) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	delete(session.Values, key)
	session.Save(s.r, s.w)
}
