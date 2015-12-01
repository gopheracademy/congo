//************************************************************************//
// congo: Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"errors"
	"sync"

	"github.com/gopheracademy/congo/app"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

// A tenant account
// Identifier: application/vnd.congo.api.account+json
type Account struct {
	// Date of creation
	CreatedAt string `json:"created_at,omitempty"`
	// Email of account owner
	CreatedBy string `json:"created_by,omitempty"`
	// API href of account
	Href string `json:"href,omitempty"`
	// ID of account
	ID int `json:"id,omitempty"`
	// Name of account
	Name string `json:"name,omitempty"`
}

func AccountFromCreatePayload(ctx *app.CreateAccountContext) Account {
	payload := ctx.Payload
	m := Account{}
	copier.Copy(&m, payload)
	return m
}
func (m Account) ToApp() *app.Account {
	target := app.Account{}
	copier.Copy(&target, &m)
	return &target
}

type AccountStorage interface {
	List(ctx *app.ListAccountContext) []Account
	Get(ctx *app.ShowAccountContext) (Account, error)
	Add(ctx *app.CreateAccountContext) (Account, error)
	Update(ctx *app.UpdateAccountContext, m Account) (Account, error)
	Delete(ctx *app.DeleteAccountContext, id int) (bool, error)
}

type AccountDB struct {
	DB gorm.DB
}

func NewAccountDB(db gorm.DB) *AccountDB {
	return &AccountDB{DB: db}
}

func (m *AccountDB) List(ctx *app.ListAccountContext) []Account {

	var objs []Account
	m.DB.Find(&objs)
	return objs
}

func (m *AccountDB) Get(ctx *app.ShowAccountContext) (Account, error) {

	var obj Account

	err := m.DB.Find(&obj, ctx.AccountID).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}

func (m *AccountDB) Add(ctx *app.CreateAccountContext) (Account, error) {
	model := AccountFromCreatePayload(ctx)
	err := m.DB.Create(&model).Error
	return model, err
}
func (m *AccountDB) Update(ctx *app.UpdateAccountContext, model Account) (Account, error) {
	getCtx, err := app.NewShowAccountContext(ctx.Context)
	if err != nil {
		return Account{}, err
	}
	obj, err := m.Get(getCtx)
	if err != nil {
		return model, err
	}
	err = m.DB.Model(&obj).Updates(model).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}
func (m *AccountDB) Delete(ctx *app.DeleteAccountContext, id int) (bool, error) {
	var obj Account
	err := m.DB.Delete(&obj, id).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return false, err
	}
	return true, nil
}

type MockAccountStorage struct {
	AccountList map[int]Account
	nextID      int
	mut         sync.Mutex
}

func NewMockAccountStorage() *MockAccountStorage {
	ml := make(map[int]Account, 0)
	return &MockAccountStorage{AccountList: ml}
}

func (db *MockAccountStorage) List(ctx *app.ListAccountContext) []Account {
	var list []Account = make([]Account, 0)
	for _, v := range db.AccountList {
		list = append(list, v)
	}
	return list
}

func (db *MockAccountStorage) Get(ctx *app.ShowAccountContext) (Account, error) {

	var obj Account

	obj, ok := db.AccountList[ctx.AccountID]
	if ok {
		return obj, nil
	} else {
		return obj, errors.New("Account does not exist")
	}
}

func (db *MockAccountStorage) Add(ctx *app.CreateAccountContext) (Account, error) {
	u := AccountFromCreatePayload(ctx)
	db.mut.Lock()
	db.nextID = db.nextID + 1
	u.ID = db.nextID
	db.mut.Unlock()

	db.AccountList[u.ID] = u
	return u, nil
}

func (db *MockAccountStorage) Update(ctx *app.UpdateAccountContext, u Account) (Account, error) {
	id := u.ID
	_, ok := db.AccountList[id]
	if ok {
		db.AccountList[id] = u
		return db.AccountList[id], nil
	} else {
		return u, errors.New("Account does not exist")
	}
}

func (db *MockAccountStorage) Delete(ctx *app.DeleteAccountContext, id int) (bool, error) {
	_, ok := db.AccountList[id]
	if ok {
		delete(db.AccountList, id)
		return true, nil
	} else {
		return false, errors.New("Could not delete this user")
	}
}

// A recurring event or conference
// Identifier: application/vnd.congo.api.series+json
type Series struct {
	// Account that owns bottle
	Account *Account `json:"account,omitempty"`
	// Date of creation
	CreatedAt string `json:"created_at,omitempty"`
	// API href of series
	Href string `json:"href,omitempty"`
	// ID of series
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// Date of last update
	UpdatedAt string `json:"updated_at,omitempty"`
}

func SeriesFromCreatePayload(ctx *app.CreateSeriesContext) Series {
	payload := ctx.Payload
	m := Series{}
	copier.Copy(&m, payload)
	return m
}
func (m Series) ToApp() *app.Series {
	target := app.Series{}
	copier.Copy(&target, &m)
	return &target
}

type SeriesStorage interface {
	List(ctx *app.ListSeriesContext) []Series
	Get(ctx *app.ShowSeriesContext) (Series, error)
	Add(ctx *app.CreateSeriesContext) (Series, error)
	Update(ctx *app.UpdateSeriesContext, m Series) (Series, error)
	Delete(ctx *app.DeleteSeriesContext, id int) (bool, error)
}

type SeriesDB struct {
	DB gorm.DB
}

func NewSeriesDB(db gorm.DB) *SeriesDB {
	return &SeriesDB{DB: db}
}

func (m *SeriesDB) List(ctx *app.ListSeriesContext) []Series {

	var objs []Series
	m.DB.Find(&objs)
	return objs
}

func (m *SeriesDB) Get(ctx *app.ShowSeriesContext) (Series, error) {

	var obj Series

	err := m.DB.Find(&obj, ctx.SeriesID).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}

func (m *SeriesDB) Add(ctx *app.CreateSeriesContext) (Series, error) {
	model := SeriesFromCreatePayload(ctx)
	err := m.DB.Create(&model).Error
	return model, err
}
func (m *SeriesDB) Update(ctx *app.UpdateSeriesContext, model Series) (Series, error) {
	getCtx, err := app.NewShowSeriesContext(ctx.Context)
	if err != nil {
		return Series{}, err
	}
	obj, err := m.Get(getCtx)
	if err != nil {
		return model, err
	}
	err = m.DB.Model(&obj).Updates(model).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}
func (m *SeriesDB) Delete(ctx *app.DeleteSeriesContext, id int) (bool, error) {
	var obj Series
	err := m.DB.Delete(&obj, id).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return false, err
	}
	return true, nil
}

type MockSeriesStorage struct {
	SeriesList map[int]Series
	nextID     int
	mut        sync.Mutex
}

func NewMockSeriesStorage() *MockSeriesStorage {
	ml := make(map[int]Series, 0)
	return &MockSeriesStorage{SeriesList: ml}
}

func (db *MockSeriesStorage) List(ctx *app.ListSeriesContext) []Series {
	var list []Series = make([]Series, 0)
	for _, v := range db.SeriesList {
		list = append(list, v)
	}
	return list
}

func (db *MockSeriesStorage) Get(ctx *app.ShowSeriesContext) (Series, error) {

	var obj Series

	obj, ok := db.SeriesList[ctx.SeriesID]
	if ok {
		return obj, nil
	} else {
		return obj, errors.New("Series does not exist")
	}
}

func (db *MockSeriesStorage) Add(ctx *app.CreateSeriesContext) (Series, error) {
	u := SeriesFromCreatePayload(ctx)
	db.mut.Lock()
	db.nextID = db.nextID + 1
	u.ID = db.nextID
	db.mut.Unlock()

	db.SeriesList[u.ID] = u
	return u, nil
}

func (db *MockSeriesStorage) Update(ctx *app.UpdateSeriesContext, u Series) (Series, error) {
	id := u.ID
	_, ok := db.SeriesList[id]
	if ok {
		db.SeriesList[id] = u
		return db.SeriesList[id], nil
	} else {
		return u, errors.New("Series does not exist")
	}
}

func (db *MockSeriesStorage) Delete(ctx *app.DeleteSeriesContext, id int) (bool, error) {
	_, ok := db.SeriesList[id]
	if ok {
		delete(db.SeriesList, id)
		return true, nil
	} else {
		return false, errors.New("Could not delete this user")
	}
}

// A user belonging to a tenant account
// Identifier: application/vnd.congo.api.user+json
type User struct {
	// Date of creation
	CreatedAt string `json:"created_at,omitempty"`
	// Email address of user
	Email string `json:"email,omitempty"`
	// First name of user
	FirstName string `json:"first_name,omitempty"`
	// API href of user
	Href string `json:"href,omitempty"`
	// ID of user
	ID int `json:"id,omitempty"`
	// Last name of user
	LastName string `json:"last_name,omitempty"`
}

func UserFromCreatePayload(ctx *app.CreateUserContext) User {
	payload := ctx.Payload
	m := User{}
	copier.Copy(&m, payload)
	return m
}
func (m User) ToApp() *app.User {
	target := app.User{}
	copier.Copy(&target, &m)
	return &target
}

type UserStorage interface {
	List(ctx *app.ListUserContext) []User
	Get(ctx *app.ShowUserContext) (User, error)
	Add(ctx *app.CreateUserContext) (User, error)
	Update(ctx *app.UpdateUserContext, m User) (User, error)
	Delete(ctx *app.DeleteUserContext, id int) (bool, error)
}

type UserDB struct {
	DB gorm.DB
}

func NewUserDB(db gorm.DB) *UserDB {
	return &UserDB{DB: db}
}

func (m *UserDB) List(ctx *app.ListUserContext) []User {

	var objs []User
	m.DB.Find(&objs)
	return objs
}

func (m *UserDB) Get(ctx *app.ShowUserContext) (User, error) {

	var obj User

	err := m.DB.Find(&obj, ctx.UserID).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}

func (m *UserDB) Add(ctx *app.CreateUserContext) (User, error) {
	model := UserFromCreatePayload(ctx)
	err := m.DB.Create(&model).Error
	return model, err
}
func (m *UserDB) Update(ctx *app.UpdateUserContext, model User) (User, error) {
	getCtx, err := app.NewShowUserContext(ctx.Context)
	if err != nil {
		return User{}, err
	}
	obj, err := m.Get(getCtx)
	if err != nil {
		return model, err
	}
	err = m.DB.Model(&obj).Updates(model).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}
func (m *UserDB) Delete(ctx *app.DeleteUserContext, id int) (bool, error) {
	var obj User
	err := m.DB.Delete(&obj, id).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return false, err
	}
	return true, nil
}

type MockUserStorage struct {
	UserList map[int]User
	nextID   int
	mut      sync.Mutex
}

func NewMockUserStorage() *MockUserStorage {
	ml := make(map[int]User, 0)
	return &MockUserStorage{UserList: ml}
}

func (db *MockUserStorage) List(ctx *app.ListUserContext) []User {
	var list []User = make([]User, 0)
	for _, v := range db.UserList {
		list = append(list, v)
	}
	return list
}

func (db *MockUserStorage) Get(ctx *app.ShowUserContext) (User, error) {

	var obj User

	obj, ok := db.UserList[ctx.UserID]
	if ok {
		return obj, nil
	} else {
		return obj, errors.New("User does not exist")
	}
}

func (db *MockUserStorage) Add(ctx *app.CreateUserContext) (User, error) {
	u := UserFromCreatePayload(ctx)
	db.mut.Lock()
	db.nextID = db.nextID + 1
	u.ID = db.nextID
	db.mut.Unlock()

	db.UserList[u.ID] = u
	return u, nil
}

func (db *MockUserStorage) Update(ctx *app.UpdateUserContext, u User) (User, error) {
	id := u.ID
	_, ok := db.UserList[id]
	if ok {
		db.UserList[id] = u
		return db.UserList[id], nil
	} else {
		return u, errors.New("User does not exist")
	}
}

func (db *MockUserStorage) Delete(ctx *app.DeleteUserContext, id int) (bool, error) {
	_, ok := db.UserList[id]
	if ok {
		delete(db.UserList, id)
		return true, nil
	} else {
		return false, errors.New("Could not delete this user")
	}
}
