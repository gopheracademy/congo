package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/gopheracademy/congo/app"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ValidateValidateOK test setup
func ValidateValidateOK(t *testing.T, ctrl app.ValidateController, userID string) {
	ValidateValidateOKCtx(t, context.Background(), ctrl, userID)
}

// ValidateValidateOKCtx test setup
func ValidateValidateOKCtx(t *testing.T, ctx context.Context, ctrl app.ValidateController, userID string) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/validate/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ValidateTest"), rw, req, nil)
	validateCtx, err := app.NewValidateValidateContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Validate(validateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

}

// ValidateValidateNotFound test setup
func ValidateValidateNotFound(t *testing.T, ctrl app.ValidateController, userID string) {
	ValidateValidateNotFoundCtx(t, context.Background(), ctrl, userID)
}

// ValidateValidateNotFoundCtx test setup
func ValidateValidateNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.ValidateController, userID string) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/validate/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ValidateTest"), rw, req, nil)
	validateCtx, err := app.NewValidateValidateContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Validate(validateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}
