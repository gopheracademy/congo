package models

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"

	"github.com/kr/pretty"

	"gopkg.in/authboss.v0"
)

func (m *UserDB) Get(email string) (interface{}, error) {
	fmt.Println("GET", email)
	var obj User

	err := m.db.Where("email = ?", email).First(&obj).Error
	fmt.Println(obj, err)
	if err != nil && err.Error() == "record not found" {
		fmt.Println("NOT FOUND")
		return nil, authboss.ErrUserNotFound
	}
	return &obj, err
}

func (m *UserDB) UserByOauth(id, provider string) (User, error) {
	var obj User

	err := m.db.Where("oauth2_uid= ? and oauth2_provider= ?", id, provider).First(&obj).Error
	fmt.Println(obj, err)
	return obj, err
}
func (m *UserDB) ConfirmUser(confirmToken string) (interface{}, error) {

	var obj User

	err := m.db.Where("confirm_token= ?", confirmToken).First(&obj).Error
	if err != nil && err.Error() == "record not found" {
		return nil, authboss.ErrUserNotFound
	}
	return &obj, err
}

func (m *UserDB) RecoverUser(recoverToken string) (interface{}, error) {

	var obj User

	err := m.db.Where("recover_token= ?", recoverToken).First(&obj).Error
	if err != nil && err.Error() == "record not found" {
		return nil, authboss.ErrUserNotFound
	}

	return &obj, err
}
func (m *UserDB) Put(email string, attr authboss.Attributes) error {
	obj, err := m.Get(email)
	u := obj.(*User)
	if err == authboss.ErrUserNotFound {

		return err
	}
	err = m.db.Model(&u).Updates(attr).Error
	fmt.Println(err)
	return err
}

func (m *UserDB) Create(email string, attr authboss.Attributes) error {
	fmt.Println("CREATE", email)
	var user User
	if err := attr.Bind(&user, true); err != nil {
		fmt.Println(err)
		return err
	}
	pretty.Print(user)
	err := m.db.Create(&user).Error
	fmt.Println(err)
	return err
}

func (m *UserDB) GetByLogin(ctx context.Context, l Login) (User, error) {
	var u User
	fmt.Println(l.Password)
	err := m.db.Where("email= ?", l.Email).First(&u).Error
	fmt.Println("ERR:", err)
	fmt.Println(u.Password)
	if e := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password)); e != nil {
		fmt.Println("Password Mismatch")
		fmt.Println(e.Error())
		return u, errors.New("Invalid Login Attempt")
	}
	return u, err
}
func (u *User) BeforeCreate() (err error) {
	var pw []byte
	fmt.Println("Fixing Password")
	fmt.Println(u.Password)
	fmt.Println(u)
	pw, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {

		return err
	}

	u.Password = string(pw)

	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	return
}

type Login struct {
	Email    string
	Password string
}
