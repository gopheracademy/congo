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

// app.AccountModel storage type
// Identifier:
type Account struct {
	gorm.Model
	Name string `json:"name,omitempty"`
}

func AccountFromCreatePayload(ctx *app.CreateAccountContext) Account {
	payload := ctx.Payload
	m := Account{}
	copier.Copy(&m, payload)

	return m
}

func AccountFromUpdatePayload(ctx *app.UpdateAccountContext) Account {
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
	Update(ctx *app.UpdateAccountContext) error
	Delete(ctx *app.DeleteAccountContext) error
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
func (m *AccountDB) Update(ctx *app.UpdateAccountContext) error {
	getCtx, err := app.NewShowAccountContext(ctx.Context)
	if err != nil {
		return err
	}
	obj, err := m.Get(getCtx)
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Updates(AccountFromUpdatePayload(ctx)).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return err
}
func (m *AccountDB) Delete(ctx *app.DeleteAccountContext) error {
	var obj Account
	err := m.DB.Delete(&obj, ctx.AccountID).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return err
	}
	return nil
}

type MockAccountStorage struct {
	AccountList map[uint]Account
	nextID      uint
	mut         sync.Mutex
}

func NewMockAccountStorage() *MockAccountStorage {
	ml := make(map[uint]Account, 0)
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

	obj, ok := db.AccountList[uint(ctx.AccountID)]
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

func (db *MockAccountStorage) Update(ctx *app.UpdateAccountContext) error {
	id := uint(ctx.AccountID)
	_, ok := db.AccountList[id]
	if ok {
		db.AccountList[id] = AccountFromUpdatePayload(ctx)
		return nil
	} else {
		return errors.New("Account does not exist")
	}
}

func (db *MockAccountStorage) Delete(ctx *app.DeleteAccountContext) error {
	_, ok := db.AccountList[uint(ctx.AccountID)]
	if ok {
		delete(db.AccountList, uint(ctx.AccountID))
		return nil
	} else {
		return errors.New("Could not delete this user")
	}
}

// app.SeriesModel storage type
// Identifier:
type Series struct {
	gorm.Model
	Name string `json:"name,omitempty"`
}

func SeriesFromCreatePayload(ctx *app.CreateSeriesContext) Series {
	payload := ctx.Payload
	m := Series{}
	copier.Copy(&m, payload)

	return m
}

func SeriesFromUpdatePayload(ctx *app.UpdateSeriesContext) Series {
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
	Update(ctx *app.UpdateSeriesContext) error
	Delete(ctx *app.DeleteSeriesContext) error
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
func (m *SeriesDB) Update(ctx *app.UpdateSeriesContext) error {
	getCtx, err := app.NewShowSeriesContext(ctx.Context)
	if err != nil {
		return err
	}
	obj, err := m.Get(getCtx)
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Updates(SeriesFromUpdatePayload(ctx)).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return err
}
func (m *SeriesDB) Delete(ctx *app.DeleteSeriesContext) error {
	var obj Series
	err := m.DB.Delete(&obj, ctx.SeriesID).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return err
	}
	return nil
}

type MockSeriesStorage struct {
	SeriesList map[uint]Series
	nextID     uint
	mut        sync.Mutex
}

func NewMockSeriesStorage() *MockSeriesStorage {
	ml := make(map[uint]Series, 0)
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

	obj, ok := db.SeriesList[uint(ctx.SeriesID)]
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

func (db *MockSeriesStorage) Update(ctx *app.UpdateSeriesContext) error {
	id := uint(ctx.SeriesID)
	_, ok := db.SeriesList[id]
	if ok {
		db.SeriesList[id] = SeriesFromUpdatePayload(ctx)
		return nil
	} else {
		return errors.New("Series does not exist")
	}
}

func (db *MockSeriesStorage) Delete(ctx *app.DeleteSeriesContext) error {
	_, ok := db.SeriesList[uint(ctx.SeriesID)]
	if ok {
		delete(db.SeriesList, uint(ctx.SeriesID))
		return nil
	} else {
		return errors.New("Could not delete this user")
	}
}

// app.UserModel storage type
// Identifier:
type User struct {
	gorm.Model
	AccountID uint
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func UserFromCreatePayload(ctx *app.CreateUserContext) User {
	payload := ctx.Payload
	m := User{}
	copier.Copy(&m, payload)
	m.AccountID = uint(ctx.AccountID)
	return m
}

func UserFromUpdatePayload(ctx *app.UpdateUserContext) User {
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
	Update(ctx *app.UpdateUserContext) error
	Delete(ctx *app.DeleteUserContext) error
}

type UserDB struct {
	DB gorm.DB
}

func UserFilter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("account_id = ?", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}
func NewUserDB(db gorm.DB) *UserDB {
	return &UserDB{DB: db}
}

func (m *UserDB) List(ctx *app.ListUserContext) []User {

	var objs []User
	m.DB.Scopes(UserFilter(ctx.AccountID, &m.DB)).Find(&objs)
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
func (m *UserDB) Update(ctx *app.UpdateUserContext) error {
	getCtx, err := app.NewShowUserContext(ctx.Context)
	if err != nil {
		return err
	}
	obj, err := m.Get(getCtx)
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Updates(UserFromUpdatePayload(ctx)).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return err
}
func (m *UserDB) Delete(ctx *app.DeleteUserContext) error {
	var obj User
	err := m.DB.Delete(&obj, ctx.UserID).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return err
	}
	return nil
}

type MockUserStorage struct {
	UserList map[uint]User
	nextID   uint
	mut      sync.Mutex
}

func NewMockUserStorage() *MockUserStorage {
	ml := make(map[uint]User, 0)
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

	obj, ok := db.UserList[uint(ctx.UserID)]
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

func (db *MockUserStorage) Update(ctx *app.UpdateUserContext) error {
	id := uint(ctx.UserID)
	_, ok := db.UserList[id]
	if ok {
		db.UserList[id] = UserFromUpdatePayload(ctx)
		return nil
	} else {
		return errors.New("User does not exist")
	}
}

func (db *MockUserStorage) Delete(ctx *app.DeleteUserContext) error {
	_, ok := db.UserList[uint(ctx.UserID)]
	if ok {
		delete(db.UserList, uint(ctx.UserID))
		return nil
	} else {
		return errors.New("Could not delete this user")
	}
}
