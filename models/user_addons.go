package models

import (
	"fmt"

	"github.com/kr/pretty"

	"gopkg.in/authboss.v0"
)

func (m *UserDB) Get(email string) (interface{}, error) {
	fmt.Println("GET", email)
	var obj User

	err := m.DB.Where("email = ?", email).First(&obj).Error
	fmt.Println(obj, err)
	if err != nil && err.Error() == "record not found" {
		fmt.Println("NOT FOUND")
		return nil, authboss.ErrUserNotFound
	}
	return &obj, err
}

func (m *UserDB) UserByOauth(id, provider string) (User, error) {
	var obj User

	err := m.DB.Where("oauth2_uid= ? and oauth2_provider= ?", id, provider).First(&obj).Error
	fmt.Println(obj, err)
	return obj, err
}
func (m *UserDB) ConfirmUser(confirmToken string) (interface{}, error) {

	var obj User

	err := m.DB.Where("confirm_token= ?", confirmToken).First(&obj).Error
	if err != nil && err.Error() == "record not found" {
		return nil, authboss.ErrUserNotFound
	}
	return &obj, err
}

func (m *UserDB) RecoverUser(recoverToken string) (interface{}, error) {

	var obj User

	err := m.DB.Where("recover_token= ?", recoverToken).First(&obj).Error
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
	err = m.DB.Model(&u).Updates(attr).Error
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
	err := m.DB.Create(&user).Error
	fmt.Println(err)
	return err
}
