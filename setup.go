package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"

	jg "github.com/dgrijalva/jwt-go"
)

func pubKey(*jg.Token) (interface{}, error) {
	return ioutil.ReadFile("keys/gc.rsa.pub")
}
func readPubKey() ([]byte, error) {
	return ioutil.ReadFile("keys/gc.rsa.pub")
}

func privateKey() (interface{}, error) {
	return ioutil.ReadFile("keys/gc.rsa")
}

func envFromSettings() (*Env, error) {
	var e Env
	constr := fmt.Sprintf("user=%s host=%s port=%d dbname=%s password=%s sslmode=disable",
		settings.DatabaseUsername,
		settings.DatabaseHost,
		settings.DatabasePort,
		settings.DatabaseName,
		settings.DatabasePassword)

	db, err := gorm.Open("postgres", constr)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(settings.MaxIdleConns)
	db.DB().SetMaxOpenConns(settings.MaxOpenConns)
	db.LogMode(settings.Debug)
	e.DB = *db
	return &e, nil
}

// Migrate creates seed records and updates the database schema
func Migrate(e *Env) {
	//	autoMigrate(e, &AccessToken{}, &RefreshToken{})
	autoMigrate(e, &models.Tenant{})
	autoMigrate(e, &models.User{})
	autoMigrate(e, &models.Series{})
	autoMigrate(e, &models.Event{})
	createTenants(e.DB)
	createUsers(e.DB)
	createSeries(e.DB)
	createEvents(e.DB)
}

func autoMigrate(e *Env, values ...interface{}) {
	for _, value := range values {
		e.DB.AutoMigrate(value)
	}
}

// Storage is an interface for database access
type Storage interface{}

func makeStorage(e *Env) map[string]Storage {

	storage := make(map[string]Storage, 0)
	userStorage := models.NewUserDB(e.DB)
	storage["USERSTORAGE"] = userStorage

	tenantStorage := models.NewTenantDB(e.DB)
	storage["TENANTSTORAGE"] = tenantStorage

	eventStorage := models.NewEventDB(e.DB)
	storage["EVENTSTORAGE"] = eventStorage

	seriesStorage := models.NewSeriesDB(e.DB)
	storage["SERIESSTORAGE"] = seriesStorage
	return storage
}
