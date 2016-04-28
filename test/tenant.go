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

// ListTenantOK test setup
func ListTenantOK(t *testing.T, ctrl app.TenantController) *app.TenantCollection {
	return ListTenantOKCtx(t, context.Background(), ctrl)
}

// ListTenantOKCtx test setup
func ListTenantOKCtx(t *testing.T, ctx context.Context, ctrl app.TenantController) *app.TenantCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TenantTest"), rw, req, nil)
	listCtx, err := app.NewListTenantContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.TenantCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.TenantCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ListTenantNotFound test setup
func ListTenantNotFound(t *testing.T, ctrl app.TenantController) {
	ListTenantNotFoundCtx(t, context.Background(), ctrl)
}

// ListTenantNotFoundCtx test setup
func ListTenantNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.TenantController) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TenantTest"), rw, req, nil)
	listCtx, err := app.NewListTenantContext(goaCtx, service)
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

// ShowTenantOK test setup
func ShowTenantOK(t *testing.T, ctrl app.TenantController, tenantID int) *app.Tenant {
	return ShowTenantOKCtx(t, context.Background(), ctrl, tenantID)
}

// ShowTenantOKCtx test setup
func ShowTenantOKCtx(t *testing.T, ctx context.Context, ctrl app.TenantController, tenantID int) *app.Tenant {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TenantTest"), rw, req, nil)
	showCtx, err := app.NewShowTenantContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Tenant)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Tenant", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ShowTenantNotFound test setup
func ShowTenantNotFound(t *testing.T, ctrl app.TenantController, tenantID int) {
	ShowTenantNotFoundCtx(t, context.Background(), ctrl, tenantID)
}

// ShowTenantNotFoundCtx test setup
func ShowTenantNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.TenantController, tenantID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TenantTest"), rw, req, nil)
	showCtx, err := app.NewShowTenantContext(goaCtx, service)
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

// CreateTenantCreated test setup
func CreateTenantCreated(t *testing.T, ctrl app.TenantController, payload *app.CreateTenantPayload) {
	CreateTenantCreatedCtx(t, context.Background(), ctrl, payload)
}

// CreateTenantCreatedCtx test setup
func CreateTenantCreatedCtx(t *testing.T, ctx context.Context, ctrl app.TenantController, payload *app.CreateTenantPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/tenants"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TenantTest"), rw, req, nil)
	createCtx, err := app.NewCreateTenantContext(goaCtx, service)
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

// UpdateTenantNoContent test setup
func UpdateTenantNoContent(t *testing.T, ctrl app.TenantController, tenantID int, payload *app.UpdateTenantPayload) {
	UpdateTenantNoContentCtx(t, context.Background(), ctrl, tenantID, payload)
}

// UpdateTenantNoContentCtx test setup
func UpdateTenantNoContentCtx(t *testing.T, ctx context.Context, ctrl app.TenantController, tenantID int, payload *app.UpdateTenantPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/tenants/%v", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TenantTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateTenantContext(goaCtx, service)
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

// UpdateTenantNotFound test setup
func UpdateTenantNotFound(t *testing.T, ctrl app.TenantController, tenantID int, payload *app.UpdateTenantPayload) {
	UpdateTenantNotFoundCtx(t, context.Background(), ctrl, tenantID, payload)
}

// UpdateTenantNotFoundCtx test setup
func UpdateTenantNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.TenantController, tenantID int, payload *app.UpdateTenantPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/tenants/%v", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TenantTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateTenantContext(goaCtx, service)
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

// DeleteTenantNoContent test setup
func DeleteTenantNoContent(t *testing.T, ctrl app.TenantController, tenantID int) {
	DeleteTenantNoContentCtx(t, context.Background(), ctrl, tenantID)
}

// DeleteTenantNoContentCtx test setup
func DeleteTenantNoContentCtx(t *testing.T, ctx context.Context, ctrl app.TenantController, tenantID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/tenants/%v", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TenantTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteTenantContext(goaCtx, service)
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

// DeleteTenantNotFound test setup
func DeleteTenantNotFound(t *testing.T, ctrl app.TenantController, tenantID int) {
	DeleteTenantNotFoundCtx(t, context.Background(), ctrl, tenantID)
}

// DeleteTenantNotFoundCtx test setup
func DeleteTenantNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.TenantController, tenantID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/tenants/%v", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TenantTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteTenantContext(goaCtx, service)
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
