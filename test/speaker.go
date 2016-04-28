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

// DeleteSpeakerNoContent test setup
func DeleteSpeakerNoContent(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) {
	DeleteSpeakerNoContentCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID)
}

// DeleteSpeakerNoContentCtx test setup
func DeleteSpeakerNoContentCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteSpeakerContext(goaCtx, service)
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

// DeleteSpeakerNotFound test setup
func DeleteSpeakerNotFound(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) {
	DeleteSpeakerNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID)
}

// DeleteSpeakerNotFoundCtx test setup
func DeleteSpeakerNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteSpeakerContext(goaCtx, service)
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

// ListSpeakerOKLink test setup
func ListSpeakerOKLink(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string) *app.SpeakerLinkCollection {
	return ListSpeakerOKLinkCtx(t, context.Background(), ctrl, tenantID, eventID)
}

// ListSpeakerOKLinkCtx test setup
func ListSpeakerOKLinkCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string) *app.SpeakerLinkCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	listCtx, err := app.NewListSpeakerContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.SpeakerLinkCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.SpeakerLinkCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ListSpeakerOK test setup
func ListSpeakerOK(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string) *app.SpeakerCollection {
	return ListSpeakerOKCtx(t, context.Background(), ctrl, tenantID, eventID)
}

// ListSpeakerOKCtx test setup
func ListSpeakerOKCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string) *app.SpeakerCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	listCtx, err := app.NewListSpeakerContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.SpeakerCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.SpeakerCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ListSpeakerOKAdmin test setup
func ListSpeakerOKAdmin(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string) *app.SpeakerAdminCollection {
	return ListSpeakerOKAdminCtx(t, context.Background(), ctrl, tenantID, eventID)
}

// ListSpeakerOKAdminCtx test setup
func ListSpeakerOKAdminCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string) *app.SpeakerAdminCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	listCtx, err := app.NewListSpeakerContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.SpeakerAdminCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.SpeakerAdminCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ListSpeakerNotFound test setup
func ListSpeakerNotFound(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string) {
	ListSpeakerNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID)
}

// ListSpeakerNotFoundCtx test setup
func ListSpeakerNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	listCtx, err := app.NewListSpeakerContext(goaCtx, service)
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

// ShowSpeakerOKLink test setup
func ShowSpeakerOKLink(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) *app.SpeakerLink {
	return ShowSpeakerOKLinkCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID)
}

// ShowSpeakerOKLinkCtx test setup
func ShowSpeakerOKLinkCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) *app.SpeakerLink {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	showCtx, err := app.NewShowSpeakerContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.SpeakerLink)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.SpeakerLink", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ShowSpeakerOK test setup
func ShowSpeakerOK(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) *app.Speaker {
	return ShowSpeakerOKCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID)
}

// ShowSpeakerOKCtx test setup
func ShowSpeakerOKCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) *app.Speaker {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	showCtx, err := app.NewShowSpeakerContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Speaker)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Speaker", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ShowSpeakerOKAdmin test setup
func ShowSpeakerOKAdmin(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) *app.SpeakerAdmin {
	return ShowSpeakerOKAdminCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID)
}

// ShowSpeakerOKAdminCtx test setup
func ShowSpeakerOKAdminCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) *app.SpeakerAdmin {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	showCtx, err := app.NewShowSpeakerContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.SpeakerAdmin)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.SpeakerAdmin", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ShowSpeakerNotFound test setup
func ShowSpeakerNotFound(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) {
	ShowSpeakerNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID)
}

// ShowSpeakerNotFoundCtx test setup
func ShowSpeakerNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	showCtx, err := app.NewShowSpeakerContext(goaCtx, service)
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

// CreateSpeakerCreated test setup
func CreateSpeakerCreated(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string, payload *app.CreateSpeakerPayload) {
	CreateSpeakerCreatedCtx(t, context.Background(), ctrl, tenantID, eventID, payload)
}

// CreateSpeakerCreatedCtx test setup
func CreateSpeakerCreatedCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string, payload *app.CreateSpeakerPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/tenants/%v/events/%v/speakers", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	createCtx, err := app.NewCreateSpeakerContext(goaCtx, service)
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

// UpdateSpeakerNotFound test setup
func UpdateSpeakerNotFound(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int, payload *app.UpdateSpeakerPayload) {
	UpdateSpeakerNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID, payload)
}

// UpdateSpeakerNotFoundCtx test setup
func UpdateSpeakerNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int, payload *app.UpdateSpeakerPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateSpeakerContext(goaCtx, service)
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

// UpdateSpeakerNoContent test setup
func UpdateSpeakerNoContent(t *testing.T, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int, payload *app.UpdateSpeakerPayload) {
	UpdateSpeakerNoContentCtx(t, context.Background(), ctrl, tenantID, eventID, speakerID, payload)
}

// UpdateSpeakerNoContentCtx test setup
func UpdateSpeakerNoContentCtx(t *testing.T, ctx context.Context, ctrl app.SpeakerController, tenantID string, eventID string, speakerID int, payload *app.UpdateSpeakerPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpeakerTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateSpeakerContext(goaCtx, service)
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
