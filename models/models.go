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

// AccountModel media type
// Identifier:
type AccountModel struct {
	gorm.Model
	Name string `json:"name,omitempty"`
}

func AccountModelFromCreatePayload(ctx *app.CreateAccountContext) AccountModel {
	payload := ctx.Payload
	m := AccountModel{}
	copier.Copy(&m, payload)

	return m
}

func AccountModelFromUpdatePayload(ctx *app.UpdateAccountContext) AccountModel {
	payload := ctx.Payload
	m := AccountModel{}
	copier.Copy(&m, payload)
	return m
}
func (m AccountModel) ToApp() *app.Account {
	target := app.Account{}
	copier.Copy(&target, &m)
	return &target
}

type AccountModelStorage interface {
	List(ctx *app.ListAccountContext) []AccountModel
	Get(ctx *app.ShowAccountContext) (AccountModel, error)
	Add(ctx *app.CreateAccountContext) (AccountModel, error)
	Update(ctx *app.UpdateAccountContext) error
	Delete(ctx *app.DeleteAccountContext) error
}

type AccountModelDB struct {
	DB gorm.DB
}

func NewAccountModelDB(db gorm.DB) *AccountModelDB {
	return &AccountModelDB{DB: db}
}

func (m *AccountModelDB) List(ctx *app.ListAccountContext) []AccountModel {

	var objs []AccountModel
	m.DB.Find(&objs)
	return objs
}

func (m *AccountModelDB) Get(ctx *app.ShowAccountContext) (AccountModel, error) {

	var obj AccountModel

	err := m.DB.Find(&obj, ctx.AccountID).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}

func (m *AccountModelDB) Add(ctx *app.CreateAccountContext) (AccountModel, error) {
	model := AccountModelFromCreatePayload(ctx)
	err := m.DB.Create(&model).Error
	return model, err
}
func (m *AccountModelDB) Update(ctx *app.UpdateAccountContext) error {
	getCtx, err := app.NewShowAccountContext(ctx.Context)
	if err != nil {
		return err
	}
	obj, err := m.Get(getCtx)
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Updates(AccountModelFromUpdatePayload(ctx)).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return err
}
func (m *AccountModelDB) Delete(ctx *app.DeleteAccountContext) error {
	var obj AccountModel
	err := m.DB.Delete(&obj, ctx.AccountID).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return err
	}
	return nil
}

type MockAccountModelStorage struct {
	AccountModelList map[uint]AccountModel
	nextID           uint
	mut              sync.Mutex
}

func NewMockAccountModelStorage() *MockAccountModelStorage {
	ml := make(map[uint]AccountModel, 0)
	return &MockAccountModelStorage{AccountModelList: ml}
}

func (db *MockAccountModelStorage) List(ctx *app.ListAccountContext) []AccountModel {
	var list []AccountModel = make([]AccountModel, 0)
	for _, v := range db.AccountModelList {
		list = append(list, v)
	}
	return list
}

func (db *MockAccountModelStorage) Get(ctx *app.ShowAccountContext) (AccountModel, error) {

	var obj AccountModel

	obj, ok := db.AccountModelList[uint(ctx.AccountID)]
	if ok {
		return obj, nil
	} else {
		return obj, errors.New("AccountModel does not exist")
	}
}

func (db *MockAccountModelStorage) Add(ctx *app.CreateAccountContext) (AccountModel, error) {
	u := AccountModelFromCreatePayload(ctx)
	db.mut.Lock()
	db.nextID = db.nextID + 1
	u.ID = db.nextID
	db.mut.Unlock()

	db.AccountModelList[u.ID] = u
	return u, nil
}

func (db *MockAccountModelStorage) Update(ctx *app.UpdateAccountContext) error {
	id := uint(ctx.AccountID)
	_, ok := db.AccountModelList[id]
	if ok {
		db.AccountModelList[id] = AccountModelFromUpdatePayload(ctx)
		return nil
	} else {
		return errors.New("AccountModel does not exist")
	}
}

func (db *MockAccountModelStorage) Delete(ctx *app.DeleteAccountContext) error {
	_, ok := db.AccountModelList[uint(ctx.AccountID)]
	if ok {
		delete(db.AccountModelList, uint(ctx.AccountID))
		return nil
	} else {
		return errors.New("Could not delete this user")
	}
}

// SeriesModel media type
// Identifier:
type SeriesModel struct {
	gorm.Model
	Name string `json:"name,omitempty"`
}

func SeriesModelFromCreatePayload(ctx *app.CreateSeriesContext) SeriesModel {
	payload := ctx.Payload
	m := SeriesModel{}
	copier.Copy(&m, payload)

	return m
}

func SeriesModelFromUpdatePayload(ctx *app.UpdateSeriesContext) SeriesModel {
	payload := ctx.Payload
	m := SeriesModel{}
	copier.Copy(&m, payload)
	return m
}
func (m SeriesModel) ToApp() *app.Series {
	target := app.Series{}
	copier.Copy(&target, &m)
	return &target
}

type SeriesModelStorage interface {
	List(ctx *app.ListSeriesContext) []SeriesModel
	Get(ctx *app.ShowSeriesContext) (SeriesModel, error)
	Add(ctx *app.CreateSeriesContext) (SeriesModel, error)
	Update(ctx *app.UpdateSeriesContext) error
	Delete(ctx *app.DeleteSeriesContext) error
}

type SeriesModelDB struct {
	DB gorm.DB
}

func NewSeriesModelDB(db gorm.DB) *SeriesModelDB {
	return &SeriesModelDB{DB: db}
}

func (m *SeriesModelDB) List(ctx *app.ListSeriesContext) []SeriesModel {

	var objs []SeriesModel
	m.DB.Find(&objs)
	return objs
}

func (m *SeriesModelDB) Get(ctx *app.ShowSeriesContext) (SeriesModel, error) {

	var obj SeriesModel

	err := m.DB.Find(&obj, ctx.SeriesID).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}

func (m *SeriesModelDB) Add(ctx *app.CreateSeriesContext) (SeriesModel, error) {
	model := SeriesModelFromCreatePayload(ctx)
	err := m.DB.Create(&model).Error
	return model, err
}
func (m *SeriesModelDB) Update(ctx *app.UpdateSeriesContext) error {
	getCtx, err := app.NewShowSeriesContext(ctx.Context)
	if err != nil {
		return err
	}
	obj, err := m.Get(getCtx)
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Updates(SeriesModelFromUpdatePayload(ctx)).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return err
}
func (m *SeriesModelDB) Delete(ctx *app.DeleteSeriesContext) error {
	var obj SeriesModel
	err := m.DB.Delete(&obj, ctx.SeriesID).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return err
	}
	return nil
}

type MockSeriesModelStorage struct {
	SeriesModelList map[uint]SeriesModel
	nextID          uint
	mut             sync.Mutex
}

func NewMockSeriesModelStorage() *MockSeriesModelStorage {
	ml := make(map[uint]SeriesModel, 0)
	return &MockSeriesModelStorage{SeriesModelList: ml}
}

func (db *MockSeriesModelStorage) List(ctx *app.ListSeriesContext) []SeriesModel {
	var list []SeriesModel = make([]SeriesModel, 0)
	for _, v := range db.SeriesModelList {
		list = append(list, v)
	}
	return list
}

func (db *MockSeriesModelStorage) Get(ctx *app.ShowSeriesContext) (SeriesModel, error) {

	var obj SeriesModel

	obj, ok := db.SeriesModelList[uint(ctx.SeriesID)]
	if ok {
		return obj, nil
	} else {
		return obj, errors.New("SeriesModel does not exist")
	}
}

func (db *MockSeriesModelStorage) Add(ctx *app.CreateSeriesContext) (SeriesModel, error) {
	u := SeriesModelFromCreatePayload(ctx)
	db.mut.Lock()
	db.nextID = db.nextID + 1
	u.ID = db.nextID
	db.mut.Unlock()

	db.SeriesModelList[u.ID] = u
	return u, nil
}

func (db *MockSeriesModelStorage) Update(ctx *app.UpdateSeriesContext) error {
	id := uint(ctx.SeriesID)
	_, ok := db.SeriesModelList[id]
	if ok {
		db.SeriesModelList[id] = SeriesModelFromUpdatePayload(ctx)
		return nil
	} else {
		return errors.New("SeriesModel does not exist")
	}
}

func (db *MockSeriesModelStorage) Delete(ctx *app.DeleteSeriesContext) error {
	_, ok := db.SeriesModelList[uint(ctx.SeriesID)]
	if ok {
		delete(db.SeriesModelList, uint(ctx.SeriesID))
		return nil
	} else {
		return errors.New("Could not delete this user")
	}
}

// UserModel media type
// Identifier:
type UserModel struct {
	gorm.Model
	AccountID uint
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func UserModelFromCreatePayload(ctx *app.CreateUserContext) UserModel {
	payload := ctx.Payload
	m := UserModel{}
	copier.Copy(&m, payload)
	m.AccountID = ctx.AccountID
	return m
}

func UserModelFromUpdatePayload(ctx *app.UpdateUserContext) UserModel {
	payload := ctx.Payload
	m := UserModel{}
	copier.Copy(&m, payload)
	return m
}
func (m UserModel) ToApp() *app.User {
	target := app.User{}
	copier.Copy(&target, &m)
	return &target
}

type UserModelStorage interface {
	List(ctx *app.ListUserContext) []UserModel
	Get(ctx *app.ShowUserContext) (UserModel, error)
	Add(ctx *app.CreateUserContext) (UserModel, error)
	Update(ctx *app.UpdateUserContext) error
	Delete(ctx *app.DeleteUserContext) error
}

type UserModelDB struct {
	DB gorm.DB
}

func UserModelFilter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
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
func NewUserModelDB(db gorm.DB) *UserModelDB {
	return &UserModelDB{DB: db}
}

func (m *UserModelDB) List(ctx *app.ListUserContext) []UserModel {

	var objs []UserModel
	m.DB.Scopes(UserModelFilter(ctx.AccountID, &m.DB)).Find(&objs)
	return objs
}

func (m *UserModelDB) Get(ctx *app.ShowUserContext) (UserModel, error) {

	var obj UserModel

	err := m.DB.Find(&obj, ctx.UserID).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}

func (m *UserModelDB) Add(ctx *app.CreateUserContext) (UserModel, error) {
	model := UserModelFromCreatePayload(ctx)
	err := m.DB.Create(&model).Error
	return model, err
}
func (m *UserModelDB) Update(ctx *app.UpdateUserContext) error {
	getCtx, err := app.NewShowUserContext(ctx.Context)
	if err != nil {
		return err
	}
	obj, err := m.Get(getCtx)
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Updates(UserModelFromUpdatePayload(ctx)).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return err
}
func (m *UserModelDB) Delete(ctx *app.DeleteUserContext) error {
	var obj UserModel
	err := m.DB.Delete(&obj, ctx.UserID).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return err
	}
	return nil
}

type MockUserModelStorage struct {
	UserModelList map[uint]UserModel
	nextID        uint
	mut           sync.Mutex
}

func NewMockUserModelStorage() *MockUserModelStorage {
	ml := make(map[uint]UserModel, 0)
	return &MockUserModelStorage{UserModelList: ml}
}

func (db *MockUserModelStorage) List(ctx *app.ListUserContext) []UserModel {
	var list []UserModel = make([]UserModel, 0)
	for _, v := range db.UserModelList {
		list = append(list, v)
	}
	return list
}

func (db *MockUserModelStorage) Get(ctx *app.ShowUserContext) (UserModel, error) {

	var obj UserModel

	obj, ok := db.UserModelList[uint(ctx.UserID)]
	if ok {
		return obj, nil
	} else {
		return obj, errors.New("UserModel does not exist")
	}
}

func (db *MockUserModelStorage) Add(ctx *app.CreateUserContext) (UserModel, error) {
	u := UserModelFromCreatePayload(ctx)
	db.mut.Lock()
	db.nextID = db.nextID + 1
	u.ID = db.nextID
	db.mut.Unlock()

	db.UserModelList[u.ID] = u
	return u, nil
}

func (db *MockUserModelStorage) Update(ctx *app.UpdateUserContext) error {
	id := uint(ctx.UserID)
	_, ok := db.UserModelList[id]
	if ok {
		db.UserModelList[id] = UserModelFromUpdatePayload(ctx)
		return nil
	} else {
		return errors.New("UserModel does not exist")
	}
}

func (db *MockUserModelStorage) Delete(ctx *app.DeleteUserContext) error {
	_, ok := db.UserModelList[uint(ctx.UserID)]
	if ok {
		delete(db.UserModelList, uint(ctx.UserID))
		return nil
	} else {
		return errors.New("Could not delete this user")
	}
}
