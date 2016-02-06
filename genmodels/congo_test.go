package genmodels

import (
	"fmt"
	"os"
	"testing"
	"time"

	"gopkg.in/inconshreveable/log15.v2"

	"golang.org/x/net/context"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db gorm.DB
var logger log15.Logger
var ctx *goa.Context

func TestMain(m *testing.M) {
	//c := dockertest.Launch("registry.xordataexchange.com/xor/postgres-dev")
	//fmt.Println(c)
	var err error
	//port, err := strconv.Atoi(strings.Split(c.Port(5432), ":")[1])
	//host := strings.Split(c.Port(5432), ":")[0]
	url := fmt.Sprintf("dbname=xorapidb user=docker password=docker sslmode=disable port=%d host=%s", 32779, "192.168.100.5")
	fmt.Println(url)
	time.Sleep(10)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.DropTable(&Proposal{}, &Review{}, &User{})
	db.AutoMigrate(&Proposal{}, &Review{}, &User{})
	logger = log15.New("tests", "congo")
	gctx := context.Background()
	ctx = goa.NewContext(gctx, goa.New("test"), nil, nil, nil)
	ctx.Logger = logger
	setup()
	//defer c.Close()
	os.Exit(m.Run())
}

func TestOneUser(t *testing.T) {
	db.LogMode(true)
	bdb := NewUserDB(db, logger)
	btl := bdb.OneUser(ctx, 1)
	if *btl.ID != 1 {
		t.Error("Expected UserID to be 1")
	}
	if *btl.City != "Tampa" {
		t.Error("Expected tampa, got squat")
	}
}

/*
func TestOneUserTiny(t *testing.T) {
	db.LogMode(true)
	bdb := NewBottleDB(db, logger)
	btl := bdb.OneBottleTiny(*ctx, 1)
	if btl.Name != "Red Horse" {
		t.Error("Expected name to be set")
	}
}
func TestGetBottle(t *testing.T) {
	db.LogMode(true)
	bdb := NewBottleDB(db, logger)
	btl := bdb.Get(*ctx, 1)
	if btl.ID != 1 {
		t.Error("Expected Bottle")
	}

}
func TestBottleToBottle(t *testing.T) {
	db.LogMode(true)
	bdb := NewBottleDB(db, logger)
	btl := bdb.Get(*ctx, 1)
	if btl.ID != 1 {
		t.Error("Expected Bottle")
	}
	appbottle := btl.BottleToBottle()
	if appbottle.ID != btl.ID {
		t.Error("Expected bottle id to transfer")
	}
	if appbottle.Vintage != *btl.Vintage {
		t.Error("Expected vintager to transfer")
	}
	if *appbottle.Rating != *btl.Rating {
		t.Error("Expected rating to transfer")
	}
}
*/
func TestOneProposal(t *testing.T) {
	db.LogMode(true)
	adb := NewProposalDB(db, logger)
	act := adb.OneProposal(ctx, 1)
	fmt.Println(act.ID)
}

func TestGetProposal(t *testing.T) {
	db.LogMode(true)
	adb := NewProposalDB(db, logger)
	act := adb.Get(ctx, 1)
	if act.ID != 1 {
		t.Error("Expected Proposal")
	}
}
func setup() error {
	adb := NewUserDB(db, logger)
	bio := "awesome"
	city := "Tampa"
	country := "US of A"
	email := "bketelsen@gmail.com"
	usr, err := adb.Add(*ctx, User{
		Bio:     &bio,
		City:    &city,
		Country: &country,
		Email:   &email,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(usr.ID)

	bdb := NewProposalDB(db, logger)

	var abstract string
	var detail string
	abstract = "abstract"
	detail = "Detail"
	prop, err := bdb.Add(*ctx, Proposal{
		Abstract: &abstract,
		Detail:   &detail,
		UserID:   usr.ID,
	})
	fmt.Println(prop.ID, prop.UserID)
	return err
}
