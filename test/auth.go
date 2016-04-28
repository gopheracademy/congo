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

// TokenAuthCreated test setup
func TokenAuthCreated(t *testing.T, ctrl app.AuthController) *app.Authorize {
	return TokenAuthCreatedCtx(t, context.Background(), ctrl)
}

// TokenAuthCreatedCtx test setup
func TokenAuthCreatedCtx(t *testing.T, ctx context.Context, ctrl app.AuthController) *app.Authorize {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/auth/token"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AuthTest"), rw, req, nil)
	tokenCtx, err := app.NewTokenAuthContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Token(tokenCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Authorize)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Authorize", resp)
	}

	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// RefreshAuthCreated test setup
func RefreshAuthCreated(t *testing.T, ctrl app.AuthController) *app.Authorize {
	return RefreshAuthCreatedCtx(t, context.Background(), ctrl)
}

// RefreshAuthCreatedCtx test setup
func RefreshAuthCreatedCtx(t *testing.T, ctx context.Context, ctrl app.AuthController) *app.Authorize {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/auth/refresh"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AuthTest"), rw, req, nil)
	refreshCtx, err := app.NewRefreshAuthContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Refresh(refreshCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Authorize)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Authorize", resp)
	}

	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}
