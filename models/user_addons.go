package models

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

func (m *UserDB) GetByLogin(ctx context.Context, l Login) (User, error) {
	var u User
	fmt.Println(l.Password)
	err := m.DB.Where("email= ?", l.Email).First(&u).Error
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
