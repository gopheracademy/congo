package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"

	"code.palmstonegames.com/polymer"
)

func init() {
	polymer.Register("my-login", &Login{})
}

type Login struct {
	*polymer.Proto
	DoLogin  chan *polymer.Event `polymer:"handler"`
	Username string              `polymer:"bind"`
	Password string              `polymer:"bind"`
	Hideform bool                `polymer:"bind"`
}

func (l *Login) Created() {
	l.Username = "user@example.com"
	l.Hideform = false
}

func (l *Login) Ready() {
	l.Hideform = false
	l.ListenToggleEvents()
}

func (l *Login) ListenToggleEvents() {
	go func() {
		for {
			select {
			case e := <-l.DoLogin:
				fmt.Println("Something Happened!", e)
				fmt.Println(l.Hideform)
				fmt.Println(l.Username, l.Password)
				l.Hideform = true
				l.Notify("hideform")
				fmt.Println(l.Hideform)
				js.Global.Get("page").Call("redirect", "/")
			}

		}
	}()
}

func main() {}
