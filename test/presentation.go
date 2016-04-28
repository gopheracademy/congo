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

// ListPresentationOKAdmin test setup
func ListPresentationOKAdmin(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string) *app.PresentationAdminCollection {
	return ListPresentationOKAdminCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID)
}

// ListPresentationOKAdminCtx test setup
func ListPresentationOKAdminCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string) *app.PresentationAdminCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	listCtx, err := app.NewListPresentationContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.PresentationAdminCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.PresentationAdminCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ListPresentationOK test setup
func ListPresentationOK(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string) *app.PresentationCollection {
	return ListPresentationOKCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID)
}

// ListPresentationOKCtx test setup
func ListPresentationOKCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string) *app.PresentationCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	listCtx, err := app.NewListPresentationContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.PresentationCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.PresentationCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ListPresentationNotFound test setup
func ListPresentationNotFound(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string) {
	ListPresentationNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID)
}

// ListPresentationNotFoundCtx test setup
func ListPresentationNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	listCtx, err := app.NewListPresentationContext(goaCtx, service)
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

// ShowPresentationOK test setup
func ShowPresentationOK(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int) *app.Presentation {
	return ShowPresentationOKCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID, presentationID)
}

// ShowPresentationOKCtx test setup
func ShowPresentationOKCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int) *app.Presentation {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	showCtx, err := app.NewShowPresentationContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Presentation)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Presentation", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ShowPresentationOKAdmin test setup
func ShowPresentationOKAdmin(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int) *app.PresentationAdmin {
	return ShowPresentationOKAdminCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID, presentationID)
}

// ShowPresentationOKAdminCtx test setup
func ShowPresentationOKAdminCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int) *app.PresentationAdmin {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	showCtx, err := app.NewShowPresentationContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.PresentationAdmin)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.PresentationAdmin", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ShowPresentationNotFound test setup
func ShowPresentationNotFound(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int) {
	ShowPresentationNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID, presentationID)
}

// ShowPresentationNotFoundCtx test setup
func ShowPresentationNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	showCtx, err := app.NewShowPresentationContext(goaCtx, service)
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

// CreatePresentationCreated test setup
func CreatePresentationCreated(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, payload *app.CreatePresentationPayload) {
	CreatePresentationCreatedCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID, payload)
}

// CreatePresentationCreatedCtx test setup
func CreatePresentationCreatedCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, payload *app.CreatePresentationPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	createCtx, err := app.NewCreatePresentationContext(goaCtx, service)
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

// UpdatePresentationNoContent test setup
func UpdatePresentationNoContent(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int, payload *app.UpdatePresentationPayload) {
	UpdatePresentationNoContentCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID, presentationID, payload)
}

// UpdatePresentationNoContentCtx test setup
func UpdatePresentationNoContentCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int, payload *app.UpdatePresentationPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	updateCtx, err := app.NewUpdatePresentationContext(goaCtx, service)
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

// UpdatePresentationNotFound test setup
func UpdatePresentationNotFound(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int, payload *app.UpdatePresentationPayload) {
	UpdatePresentationNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID, presentationID, payload)
}

// UpdatePresentationNotFoundCtx test setup
func UpdatePresentationNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int, payload *app.UpdatePresentationPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	updateCtx, err := app.NewUpdatePresentationContext(goaCtx, service)
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

// DeletePresentationNoContent test setup
func DeletePresentationNoContent(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int) {
	DeletePresentationNoContentCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID, presentationID)
}

// DeletePresentationNoContentCtx test setup
func DeletePresentationNoContentCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	deleteCtx, err := app.NewDeletePresentationContext(goaCtx, service)
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

// DeletePresentationNotFound test setup
func DeletePresentationNotFound(t *testing.T, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int) {
	DeletePresentationNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID, presentationID)
}

// DeletePresentationNotFoundCtx test setup
func DeletePresentationNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.PresentationController, tenantID string, eventID string, speakerID string, presentationID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PresentationTest"), rw, req, nil)
	deleteCtx, err := app.NewDeletePresentationContext(goaCtx, service)
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
