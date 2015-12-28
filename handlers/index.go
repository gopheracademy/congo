package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/context"

	"github.com/gopheracademy/congo/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/jinzhu/gorm"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func Index(db gorm.DB, ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := getPageData(db, r)
		ren.Render(w, http.StatusOK, "index", data, "layout")
	})
}
func Login(db gorm.DB, ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := getPageData(db, r)
		ren.Render(w, http.StatusOK, "login", data, "layout")
	})
}
func Logout(db gorm.DB, ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := map[string]string{}
		if encoded, err := s.Encode("cookie-name", value); err == nil {
			cookie := &http.Cookie{
				Name:  "cookie-name",
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		}
		session, _ := gothic.Store.Get(r, gothic.SessionName)
		session.Options.MaxAge = -1
		err := session.Save(r, w)
		if err != nil {
			fmt.Println("error saving session", err.Error())
		}
		data := PageData{}
		ren.Render(w, http.StatusOK, "index", data, "layout")
	})
}
func Profile(db gorm.DB, ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := getPageData(db, r)
		fmt.Println(data)
		ren.Render(w, http.StatusOK, "profile", data, "layout")
	})
}
func Landing(db gorm.DB, ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// print our state string to the console. Ideally, you should verify
		// that it's the same string as the one you set in `setState`
		fmt.Println("State: ", gothic.GetState(r))

		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		udb := models.NewUserDB(db)
		prov, _ := provider(r)
		var cuser models.User
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

		value := map[string]string{
			"userid": strconv.Itoa(cuser.ID),
		}
		if encoded, err := s.Encode("cookie-name", value); err == nil {
			cookie := &http.Cookie{
				Name:  "cookie-name",
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		}
		data := getPageData(db, r)
		data.GUser = user
		data.User = cuser
		data.CurrentUserName = cuser.Firstname
		data.UserID = strconv.Itoa(cuser.ID)
		ren.Render(w, http.StatusOK, "profile", data, "layout")
	})
}

type PageData struct {
	Loggedin        bool
	User            models.User
	UserID          string
	CurrentUserName string
	FlashSuccess    string
	FlashError      string
	GUser           goth.User
}

var s *securecookie.SecureCookie

func init() {
	// Hash keys should be at least 32 bytes long
	var hashKey = []byte("very-secret time in a gopher's existence where he likes to be S3cure")
	s = securecookie.New(hashKey, nil)

}

func getPageData(db gorm.DB, r *http.Request) PageData {
	data := PageData{}
	session, err := gothic.Store.Get(r, gothic.SessionName)
	if err != nil {
		fmt.Println("Sesson Error: ", err)
	}
	if !session.IsNew {
		data.Loggedin = true

	}

	if cookie, err := r.Cookie("cookie-name"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("cookie-name", cookie.Value, &value); err == nil {
			fmt.Println(value)
			userid, ok := value["userid"]
			if ok {
				fmt.Println("Found USER ID: ", userid)
				data.UserID = userid
				id, err := strconv.Atoi(userid)
				if err == nil {
					udb := models.NewUserDB(db)
					var cuser models.User
					cuser, err = udb.One(context.Background(), id)
					data.User = cuser
					data.CurrentUserName = cuser.Firstname

				}
			}
		}
	}
	return data
}

func provider(r *http.Request) (string, error) {

	vars := mux.Vars(r)
	provider := vars["provider"]
	if provider == "" {
		return "", errors.New("No Provider specified")
	}
	return provider, nil
}
