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

// ListEventOK test setup
func ListEventOK(t *testing.T, ctrl app.EventController, tenantID string) *app.EventCollection {
	return ListEventOKCtx(t, context.Background(), ctrl, tenantID)
}

// ListEventOKCtx test setup
func ListEventOKCtx(t *testing.T, ctx context.Context, ctrl app.EventController, tenantID string) *app.EventCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EventTest"), rw, req, nil)
	listCtx, err := app.NewListEventContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.EventCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.EventCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ListEventNotFound test setup
func ListEventNotFound(t *testing.T, ctrl app.EventController, tenantID string) {
	ListEventNotFoundCtx(t, context.Background(), ctrl, tenantID)
}

// ListEventNotFoundCtx test setup
func ListEventNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.EventController, tenantID string) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EventTest"), rw, req, nil)
	listCtx, err := app.NewListEventContext(goaCtx, service)
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

// ShowEventOK test setup
func ShowEventOK(t *testing.T, ctrl app.EventController, tenantID string, eventID int) *app.Event {
	return ShowEventOKCtx(t, context.Background(), ctrl, tenantID, eventID)
}

// ShowEventOKCtx test setup
func ShowEventOKCtx(t *testing.T, ctx context.Context, ctrl app.EventController, tenantID string, eventID int) *app.Event {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EventTest"), rw, req, nil)
	showCtx, err := app.NewShowEventContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Event)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Event", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ShowEventNotFound test setup
func ShowEventNotFound(t *testing.T, ctrl app.EventController, tenantID string, eventID int) {
	ShowEventNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID)
}

// ShowEventNotFoundCtx test setup
func ShowEventNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.EventController, tenantID string, eventID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/events/%v", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EventTest"), rw, req, nil)
	showCtx, err := app.NewShowEventContext(goaCtx, service)
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

// CreateEventCreated test setup
func CreateEventCreated(t *testing.T, ctrl app.EventController, tenantID string, payload *app.CreateEventPayload) {
	CreateEventCreatedCtx(t, context.Background(), ctrl, tenantID, payload)
}

// CreateEventCreatedCtx test setup
func CreateEventCreatedCtx(t *testing.T, ctx context.Context, ctrl app.EventController, tenantID string, payload *app.CreateEventPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/tenants/%v/events", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EventTest"), rw, req, nil)
	createCtx, err := app.NewCreateEventContext(goaCtx, service)
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

// UpdateEventNoContent test setup
func UpdateEventNoContent(t *testing.T, ctrl app.EventController, tenantID string, eventID int, payload *app.UpdateEventPayload) {
	UpdateEventNoContentCtx(t, context.Background(), ctrl, tenantID, eventID, payload)
}

// UpdateEventNoContentCtx test setup
func UpdateEventNoContentCtx(t *testing.T, ctx context.Context, ctrl app.EventController, tenantID string, eventID int, payload *app.UpdateEventPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/tenants/%v/events/%v", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EventTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateEventContext(goaCtx, service)
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

// UpdateEventNotFound test setup
func UpdateEventNotFound(t *testing.T, ctrl app.EventController, tenantID string, eventID int, payload *app.UpdateEventPayload) {
	UpdateEventNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID, payload)
}

// UpdateEventNotFoundCtx test setup
func UpdateEventNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.EventController, tenantID string, eventID int, payload *app.UpdateEventPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/tenants/%v/events/%v", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EventTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateEventContext(goaCtx, service)
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

// DeleteEventNoContent test setup
func DeleteEventNoContent(t *testing.T, ctrl app.EventController, tenantID string, eventID int) {
	DeleteEventNoContentCtx(t, context.Background(), ctrl, tenantID, eventID)
}

// DeleteEventNoContentCtx test setup
func DeleteEventNoContentCtx(t *testing.T, ctx context.Context, ctrl app.EventController, tenantID string, eventID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/tenants/%v/events/%v", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EventTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteEventContext(goaCtx, service)
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

// DeleteEventNotFound test setup
func DeleteEventNotFound(t *testing.T, ctrl app.EventController, tenantID string, eventID int) {
	DeleteEventNotFoundCtx(t, context.Background(), ctrl, tenantID, eventID)
}

// DeleteEventNotFoundCtx test setup
func DeleteEventNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.EventController, tenantID string, eventID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/tenants/%v/events/%v", tenantID, eventID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EventTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteEventContext(goaCtx, service)
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
