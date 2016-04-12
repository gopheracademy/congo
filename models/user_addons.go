package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	sendgrid "github.com/sendgrid/sendgrid-go"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

func (m *UserDB) GetByLogin(ctx context.Context, l Login) (User, error) {
	db := m.DB().(*gorm.DB)
	var u User
	err := db.Where("email= ?", l.Email).First(&u)
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password)); err != nil {
		return u, errors.New("Invalid Login Attempt")
	}
	return u, err.Error
}

//TODO: Gorm changed something here, what is it?
func (u *User) BeforeCreate() (err error) {
	var pw []byte
	pw, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {

		return err
	}
	u.Password = string(pw)
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	return
}
func (u *User) AfterCreate() error {
	if *u.Validated {
		return nil
	}
	go u.ValidationEmail()
	return nil
}

func (u *User) ValidationEmail() error {
	sg := sendgrid.NewSendGridClientWithApiKey("")
	message := sendgrid.NewMail()
	message.AddTo(u.Email)
	message.AddToName(u.FirstName + " " + u.LastName)
	message.SetSubject("Validate your account")
	msg := fmt.Sprintf("<html><body>Welcome.<br/><br/>Please click <a href='<validationurl>validaate_email?u=%d&validation=%s'>here</a> to validate your account.</body></html>", u.ID, *u.ValidationCode)
	message.SetHTML(msg)
	message.SetFrom(supportEmail)
	r := sg.Send(message)

	return r
}

type Login struct {
	Email    string
	Password string
}
