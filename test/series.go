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

// CreateSeriesCreated test setup
func CreateSeriesCreated(t *testing.T, ctrl app.SeriesController, tenantID string, payload *app.CreateSeriesPayload) {
	CreateSeriesCreatedCtx(t, context.Background(), ctrl, tenantID, payload)
}

// CreateSeriesCreatedCtx test setup
func CreateSeriesCreatedCtx(t *testing.T, ctx context.Context, ctrl app.SeriesController, tenantID string, payload *app.CreateSeriesPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/tenants/%v/series", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SeriesTest"), rw, req, nil)
	createCtx, err := app.NewCreateSeriesContext(goaCtx, service)
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

// UpdateSeriesNoContent test setup
func UpdateSeriesNoContent(t *testing.T, ctrl app.SeriesController, tenantID string, seriesID int, payload *app.UpdateSeriesPayload) {
	UpdateSeriesNoContentCtx(t, context.Background(), ctrl, tenantID, seriesID, payload)
}

// UpdateSeriesNoContentCtx test setup
func UpdateSeriesNoContentCtx(t *testing.T, ctx context.Context, ctrl app.SeriesController, tenantID string, seriesID int, payload *app.UpdateSeriesPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/tenants/%v/series/%v", tenantID, seriesID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SeriesTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateSeriesContext(goaCtx, service)
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

// UpdateSeriesNotFound test setup
func UpdateSeriesNotFound(t *testing.T, ctrl app.SeriesController, tenantID string, seriesID int, payload *app.UpdateSeriesPayload) {
	UpdateSeriesNotFoundCtx(t, context.Background(), ctrl, tenantID, seriesID, payload)
}

// UpdateSeriesNotFoundCtx test setup
func UpdateSeriesNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.SeriesController, tenantID string, seriesID int, payload *app.UpdateSeriesPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/tenants/%v/series/%v", tenantID, seriesID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SeriesTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateSeriesContext(goaCtx, service)
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

// DeleteSeriesNoContent test setup
func DeleteSeriesNoContent(t *testing.T, ctrl app.SeriesController, tenantID string, seriesID int) {
	DeleteSeriesNoContentCtx(t, context.Background(), ctrl, tenantID, seriesID)
}

// DeleteSeriesNoContentCtx test setup
func DeleteSeriesNoContentCtx(t *testing.T, ctx context.Context, ctrl app.SeriesController, tenantID string, seriesID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/tenants/%v/series/%v", tenantID, seriesID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SeriesTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteSeriesContext(goaCtx, service)
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

// DeleteSeriesNotFound test setup
func DeleteSeriesNotFound(t *testing.T, ctrl app.SeriesController, tenantID string, seriesID int) {
	DeleteSeriesNotFoundCtx(t, context.Background(), ctrl, tenantID, seriesID)
}

// DeleteSeriesNotFoundCtx test setup
func DeleteSeriesNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.SeriesController, tenantID string, seriesID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/tenants/%v/series/%v", tenantID, seriesID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SeriesTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteSeriesContext(goaCtx, service)
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

// ListSeriesOK test setup
func ListSeriesOK(t *testing.T, ctrl app.SeriesController, tenantID string) *app.SeriesCollection {
	return ListSeriesOKCtx(t, context.Background(), ctrl, tenantID)
}

// ListSeriesOKCtx test setup
func ListSeriesOKCtx(t *testing.T, ctx context.Context, ctrl app.SeriesController, tenantID string) *app.SeriesCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/series", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SeriesTest"), rw, req, nil)
	listCtx, err := app.NewListSeriesContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.SeriesCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.SeriesCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ListSeriesNotFound test setup
func ListSeriesNotFound(t *testing.T, ctrl app.SeriesController, tenantID string) {
	ListSeriesNotFoundCtx(t, context.Background(), ctrl, tenantID)
}

// ListSeriesNotFoundCtx test setup
func ListSeriesNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.SeriesController, tenantID string) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/series", tenantID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SeriesTest"), rw, req, nil)
	listCtx, err := app.NewListSeriesContext(goaCtx, service)
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

// ShowSeriesOK test setup
func ShowSeriesOK(t *testing.T, ctrl app.SeriesController, tenantID string, seriesID int) *app.Series {
	return ShowSeriesOKCtx(t, context.Background(), ctrl, tenantID, seriesID)
}

// ShowSeriesOKCtx test setup
func ShowSeriesOKCtx(t *testing.T, ctx context.Context, ctrl app.SeriesController, tenantID string, seriesID int) *app.Series {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/series/%v", tenantID, seriesID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SeriesTest"), rw, req, nil)
	showCtx, err := app.NewShowSeriesContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Series)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Series", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ShowSeriesNotFound test setup
func ShowSeriesNotFound(t *testing.T, ctrl app.SeriesController, tenantID string, seriesID int) {
	ShowSeriesNotFoundCtx(t, context.Background(), ctrl, tenantID, seriesID)
}

// ShowSeriesNotFoundCtx test setup
func ShowSeriesNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.SeriesController, tenantID string, seriesID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/tenants/%v/series/%v", tenantID, seriesID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SeriesTest"), rw, req, nil)
	showCtx, err := app.NewShowSeriesContext(goaCtx, service)
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
