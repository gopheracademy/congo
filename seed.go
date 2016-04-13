package main

import (
	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"
)

func createTenants(db gorm.DB) error {
	var err error
	var count int
	err = db.Model(&models.Tenant{}).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		var name string
		name = "GopherAcademy, LLC"

		tenant := models.Tenant{}
		tenant.Name = name

		err = db.Create(&tenant).Error
	}
	return err

}

func createSeries(db gorm.DB) error {
	var err error
	var count int
	err = db.Model(&models.Series{}).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		var name string
		name = "GopherCon"

		series := models.Series{}
		series.Name = name

		err = db.Create(&series).Error
	}
	return err

}

func createEvents(db gorm.DB) error {
	var err error
	var count int
	err = db.Model(&models.Event{}).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		var name string
		name = "GopherCon 2016"

		event := models.Event{}
		event.Name = name
		event.TenantID = 1

		err = db.Create(&event).Error
	}
	return err

}
func createUsers(db gorm.DB) error {
	var err error
	var count int
	err = db.Model(&models.User{}).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		var firstname string
		var lastname string
		var password string
		var email string
		var validated bool
		firstname = "Brian"
		lastname = "Ketelsen"
		password = "GopherCon"
		email = "bketelsen@gmail.com"
		validated = true

		user := models.User{}
		user.FirstName = firstname
		user.LastName = lastname
		user.Password = password
		user.Email = email
		user.Validated = &validated
		user.TenantID = 1

		err = db.Create(&user).Error
	}
	return err

}
