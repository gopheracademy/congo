package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aarondl/tpl"
	"github.com/codegangsta/negroni"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/handlers"
	"github.com/gopheracademy/congo/models"
	"github.com/gopheracademy/congo/swagger"
	"github.com/gopheracademy/congo/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/nosurf"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/raphael/goa"
	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/confirm"
	_ "gopkg.in/authboss.v0/lock"
	_ "gopkg.in/authboss.v0/recover"
	_ "gopkg.in/authboss.v0/register"
	//	_ "gopkg.in/authboss.v0/remember"
)

var (
	Ab        = authboss.New()
	templates = tpl.Must(tpl.Load("views", "views/partials", "layout.html.tpl", funcs))
	schemaDec = schema.NewDecoder()
)

func LayoutData(w http.ResponseWriter, r *http.Request) authboss.HTMLData {
	currentUserName := ""
	userInter, err := Ab.CurrentUser(w, r)
	if userInter != nil && err == nil {
		currentUserName = userInter.(*models.User).Firstname
	}

	return authboss.HTMLData{
		"loggedin":               userInter != nil,
		"username":               "",
		"firstname":              "",
		"lastname":               "",
		authboss.FlashSuccessKey: Ab.FlashSuccess(w, r),
		authboss.FlashErrorKey:   Ab.FlashError(w, r),
		"current_user_name":      currentUserName,
	}
}

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
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	// Mount "user" controller
	c3 := NewUserController(service, models.NewUserDB(db))
	app.MountUserController(service, c3)

	// Mount Swagger spec provider controller
	swagger.MountController(service)

	goarouter := service.HTTPHandler().(*httprouter.Router)

	renderer := handlers.NewRenderRenderer("templates", []string{".html"}, handlers.Funcs, dev)
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`NpEPi8pEjKVjLGJ6kYCS+VTCzi6BUuDzU0wrwXyf5uDPArtlofn2AG6aTMiPmN3C909rsEWMNqJqhIVPGP3Exg==`)
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(`Ab.Ywmmt8UCwUuhd9qvfNA9UCuN1cVcKJN1ofbiky6xCyyBj20whe40rJa3Su0WOWLWcPpO1taqJdsEI/65+JA==`)
	cookieStore = securecookie.New(cookieStoreKey, nil)
	sessionStore = sessions.NewCookieStore(sessionStoreKey)

	Ab.MountPath = "/auth"
	Ab.LogWriter = os.Stdout
	Ab.Mailer = authboss.LogMailer(os.Stdout)
	Ab.Storer = models.NewUserDB(db)
	Ab.ViewsPath = "ab_views"
	Ab.LayoutDataMaker = LayoutData
	b, err := ioutil.ReadFile(filepath.Join("views", "layout.html.tpl"))
	if err != nil {
		panic(err)
	}
	Ab.Layout = template.Must(template.New("layout").Funcs(funcs).Parse(string(b)))
	Ab.RootURL = `http://localhost:8080`
	Ab.CookieStoreMaker = NewCookieStorer
	Ab.SessionStoreMaker = NewSessionStorer
	Ab.XSRFName = "csrf_token"
	Ab.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return nosurf.Token(r)
	}
	if err := Ab.Init(); err != nil {
		// Handle error, don't let program continue to run
		log.Fatalln(err)
	}

	// Make sure to put authboss's router somewhere
	r := mux.NewRouter()
	// Make sure to put authboss's router somewhere
	r.Handle("/", handlers.Index(renderer, LayoutData)).Methods("GET")
	r.Handle("/users", handlers.Users(renderer, LayoutData)).Methods("GET")

	r.PathPrefix("/api").Handler(goarouter)
	r.PathPrefix("/auth").Handler(Ab.NewRouter())

	n := negroni.Classic()
	n.UseHandler(r)
	hostStr := fmt.Sprintf(":%d", cfg.Port)
	n.Run(hostStr)
}

after := func(ctx *authboss.Context) error () {

	pretty.Print(ctx)
	return nil
}

var nextUserID uint
var funcs = template.FuncMap{
	"formatDate": func(date time.Time) string {
		return date.Format("2006/01/02 03:04pm")
	},
	"yield": func() string { return "" },
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

const sessionCookieName = "Ab.blog"

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

func connectDB() (gorm.DB, error) {
	var db gorm.DB
	constr := fmt.Sprintf("user=%s host=%s port=%d dbname=%s password=%s sslmode=disable",
		"postgres",
		"127.0.0.1",
		5433,
		"postgres",
		"postgres")

	var err error
	db, err = gorm.Open("postgres", constr)
	if err == nil {
		db.AutoMigrate(&models.User{})
	}
	db.LogMode(true)
	return db, err
}
