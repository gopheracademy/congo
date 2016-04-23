//************************************************************************//
// API "congo": Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// AdminuserHref returns the resource href.
func AdminuserHref(userID interface{}) string {
	return fmt.Sprintf("/api/admin/users/%v", userID)
}

// EventHref returns the resource href.
func EventHref(tenantID, eventID interface{}) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v", tenantID, eventID)
}

// PresentationHref returns the resource href.
func PresentationHref(tenantID, eventID, speakerID, presentationID interface{}) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v/presentations/%v", tenantID, eventID, speakerID, presentationID)
}

// SeriesHref returns the resource href.
func SeriesHref(tenantID, seriesID interface{}) string {
	return fmt.Sprintf("/api/tenants/%v/series/%v", tenantID, seriesID)
}

// SpeakerHref returns the resource href.
func SpeakerHref(tenantID, eventID, speakerID interface{}) string {
	return fmt.Sprintf("/api/tenants/%v/events/%v/speakers/%v", tenantID, eventID, speakerID)
}

// TenantHref returns the resource href.
func TenantHref(tenantID interface{}) string {
	return fmt.Sprintf("/api/tenants/%v", tenantID)
}

// UserHref returns the resource href.
func UserHref(tenantID, userID interface{}) string {
	return fmt.Sprintf("/api/tenants/%v/users/%v", tenantID, userID)
}
