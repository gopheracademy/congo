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

// ListAdminuserOK test setup
func ListAdminuserOK(t *testing.T, ctrl app.AdminuserController) *app.UserCollection {
	return ListAdminuserOKCtx(t, context.Background(), ctrl)
}

// ListAdminuserOKCtx test setup
func ListAdminuserOKCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController) *app.UserCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/admin/users"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	listCtx, err := app.NewListAdminuserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.UserCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.UserCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ListAdminuserOKTenant test setup
func ListAdminuserOKTenant(t *testing.T, ctrl app.AdminuserController) *app.UserTenantCollection {
	return ListAdminuserOKTenantCtx(t, context.Background(), ctrl)
}

// ListAdminuserOKTenantCtx test setup
func ListAdminuserOKTenantCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController) *app.UserTenantCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/admin/users"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	listCtx, err := app.NewListAdminuserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.UserTenantCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.UserTenantCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ListAdminuserNotFound test setup
func ListAdminuserNotFound(t *testing.T, ctrl app.AdminuserController) {
	ListAdminuserNotFoundCtx(t, context.Background(), ctrl)
}

// ListAdminuserNotFoundCtx test setup
func ListAdminuserNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/admin/users"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	listCtx, err := app.NewListAdminuserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// ShowAdminuserOK test setup
func ShowAdminuserOK(t *testing.T, ctrl app.AdminuserController, userID int) *app.User {
	return ShowAdminuserOKCtx(t, context.Background(), ctrl, userID)
}

// ShowAdminuserOKCtx test setup
func ShowAdminuserOKCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController, userID int) *app.User {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/admin/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	showCtx, err := app.NewShowAdminuserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.User)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.User", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowAdminuserOKTenant test setup
func ShowAdminuserOKTenant(t *testing.T, ctrl app.AdminuserController, userID int) *app.UserTenant {
	return ShowAdminuserOKTenantCtx(t, context.Background(), ctrl, userID)
}

// ShowAdminuserOKTenantCtx test setup
func ShowAdminuserOKTenantCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController, userID int) *app.UserTenant {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/admin/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	showCtx, err := app.NewShowAdminuserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.UserTenant)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.UserTenant", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowAdminuserOKLink test setup
func ShowAdminuserOKLink(t *testing.T, ctrl app.AdminuserController, userID int) *app.UserLink {
	return ShowAdminuserOKLinkCtx(t, context.Background(), ctrl, userID)
}

// ShowAdminuserOKLinkCtx test setup
func ShowAdminuserOKLinkCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController, userID int) *app.UserLink {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/admin/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	showCtx, err := app.NewShowAdminuserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.UserLink)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.UserLink", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowAdminuserNotFound test setup
func ShowAdminuserNotFound(t *testing.T, ctrl app.AdminuserController, userID int) {
	ShowAdminuserNotFoundCtx(t, context.Background(), ctrl, userID)
}

// ShowAdminuserNotFoundCtx test setup
func ShowAdminuserNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController, userID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/admin/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	showCtx, err := app.NewShowAdminuserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// CreateAdminuserCreated test setup
func CreateAdminuserCreated(t *testing.T, ctrl app.AdminuserController, payload *app.CreateAdminuserPayload) {
	CreateAdminuserCreatedCtx(t, context.Background(), ctrl, payload)
}

// CreateAdminuserCreatedCtx test setup
func CreateAdminuserCreatedCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController, payload *app.CreateAdminuserPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/admin/users"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	createCtx, err := app.NewCreateAdminuserContext(goaCtx, service)
	createCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Create(createCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}

}

// UpdateAdminuserNoContent test setup
func UpdateAdminuserNoContent(t *testing.T, ctrl app.AdminuserController, userID int, payload *app.UpdateAdminuserPayload) {
	UpdateAdminuserNoContentCtx(t, context.Background(), ctrl, userID, payload)
}

// UpdateAdminuserNoContentCtx test setup
func UpdateAdminuserNoContentCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController, userID int, payload *app.UpdateAdminuserPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/admin/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateAdminuserContext(goaCtx, service)
	updateCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

}

// UpdateAdminuserNotFound test setup
func UpdateAdminuserNotFound(t *testing.T, ctrl app.AdminuserController, userID int, payload *app.UpdateAdminuserPayload) {
	UpdateAdminuserNotFoundCtx(t, context.Background(), ctrl, userID, payload)
}

// UpdateAdminuserNotFoundCtx test setup
func UpdateAdminuserNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController, userID int, payload *app.UpdateAdminuserPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/admin/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateAdminuserContext(goaCtx, service)
	updateCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// DeleteAdminuserNoContent test setup
func DeleteAdminuserNoContent(t *testing.T, ctrl app.AdminuserController, userID int) {
	DeleteAdminuserNoContentCtx(t, context.Background(), ctrl, userID)
}

// DeleteAdminuserNoContentCtx test setup
func DeleteAdminuserNoContentCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController, userID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/admin/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteAdminuserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Delete(deleteCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

}

// DeleteAdminuserNotFound test setup
func DeleteAdminuserNotFound(t *testing.T, ctrl app.AdminuserController, userID int) {
	DeleteAdminuserNotFoundCtx(t, context.Background(), ctrl, userID)
}

// DeleteAdminuserNotFoundCtx test setup
func DeleteAdminuserNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.AdminuserController, userID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/admin/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AdminuserTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteAdminuserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Delete(deleteCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}
