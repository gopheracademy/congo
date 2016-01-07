//************************************************************************//
// API "congo": Application Media Types
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

import "github.com/raphael/goa"

// Token authorization response
// Identifier: application/vnd.authorize+json
type Authorize struct {
	// access token
	AccessToken string
	// Time to expiration in seconds
	ExpiresIn int
	// type of token
	TokenType string
}

// LoadAuthorize loads raw data into an instance of Authorize
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadAuthorize(raw interface{}) (res *Authorize, err error) {
	res, err = UnmarshalAuthorize(raw, err)
	return
}

// Dump produces raw data from an instance of Authorize running all the
// validations. See LoadAuthorize for the definition of raw data.
func (mt *Authorize) Dump() (res map[string]interface{}, err error) {
	res, err = MarshalAuthorize(mt, err)
	return
}

// MarshalAuthorize validates and renders an instance of Authorize into a interface{}
// using view "default".
func MarshalAuthorize(source *Authorize, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp7 := map[string]interface{}{
		"access_token": source.AccessToken,
		"expires_in":   source.ExpiresIn,
		"token_type":   source.TokenType,
	}
	target = tmp7
	return
}

// UnmarshalAuthorize unmarshals and validates a raw interface{} into an instance of Authorize
func UnmarshalAuthorize(source interface{}, inErr error) (target *Authorize, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Authorize)
		if v, ok := val["access_token"]; ok {
			var tmp8 string
			if val, ok := v.(string); ok {
				tmp8 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.AccessToken`, v, "string", err)
			}
			target.AccessToken = tmp8
		}
		if v, ok := val["expires_in"]; ok {
			var tmp9 int
			if f, ok := v.(float64); ok {
				tmp9 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ExpiresIn`, v, "int", err)
			}
			target.ExpiresIn = tmp9
		}
		if v, ok := val["token_type"]; ok {
			var tmp10 string
			if val, ok := v.(string); ok {
				tmp10 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.TokenType`, v, "string", err)
			}
			target.TokenType = tmp10
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// Login media type
// Identifier: application/vnd.login+json
type Login struct {
	// UUID of requesting application
	Application string
	// email
	Email string
	// password
	Password string
}

// LoadLogin loads raw data into an instance of Login
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadLogin(raw interface{}) (res *Login, err error) {
	res, err = UnmarshalLogin(raw, err)
	return
}

// Dump produces raw data from an instance of Login running all the
// validations. See LoadLogin for the definition of raw data.
func (mt *Login) Dump() (res map[string]interface{}, err error) {
	res, err = MarshalLogin(mt, err)
	return
}

// MarshalLogin validates and renders an instance of Login into a interface{}
// using view "default".
func MarshalLogin(source *Login, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp11 := map[string]interface{}{
		"application": source.Application,
		"email":       source.Email,
		"password":    source.Password,
	}
	target = tmp11
	return
}

// UnmarshalLogin unmarshals and validates a raw interface{} into an instance of Login
func UnmarshalLogin(source interface{}, inErr error) (target *Login, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Login)
		if v, ok := val["application"]; ok {
			var tmp12 string
			if val, ok := v.(string); ok {
				tmp12 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Application`, v, "string", err)
			}
			target.Application = tmp12
		}
		if v, ok := val["email"]; ok {
			var tmp13 string
			if val, ok := v.(string); ok {
				tmp13 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			target.Email = tmp13
		}
		if v, ok := val["password"]; ok {
			var tmp14 string
			if val, ok := v.(string); ok {
				tmp14 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Password`, v, "string", err)
			}
			target.Password = tmp14
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// A user belonging to a tenant account
// Identifier: application/vnd.user+json
type User struct {
	// Biography of user
	Bio string
	// City of residence
	City string
	// Country of residence
	Country string
	// Email address of user
	Email string
	// First name of user
	Firstname string
	// API href of user
	Href string
	// ID of user
	ID int
	// Last name of user
	Lastname string
	// Role of user
	Role string
	// State of residence
	State string
}

// User views
type UserViewEnum string

const (
	// User default view
	UserDefaultView UserViewEnum = "default"
	// User link view
	UserLinkView UserViewEnum = "link"
)

// LoadUser loads raw data into an instance of User
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUser(raw interface{}) (res *User, err error) {
	res, err = UnmarshalUser(raw, err)
	return
}

// Dump produces raw data from an instance of User running all the
// validations. See LoadUser for the definition of raw data.
func (mt *User) Dump(view UserViewEnum) (res map[string]interface{}, err error) {
	if view == UserDefaultView {
		res, err = MarshalUser(mt, err)
	}
	if view == UserLinkView {
		res, err = MarshalUserLink(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *User) Validate() (err error) {
	if len(mt.Bio) > 500 {
		err = goa.InvalidLengthError(`response.bio`, mt.Bio, len(mt.Bio), 500, false, err)
	}
	if mt.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2, err)
		}
	}
	if mt.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}

// MarshalUser validates and renders an instance of User into a interface{}
// using view "default".
func MarshalUser(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Bio) > 500 {
		err = goa.InvalidLengthError(`.bio`, source.Bio, len(source.Bio), 500, false, err)
	}
	if source.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, source.Email); err2 != nil {
			err = goa.InvalidFormatError(`.email`, source.Email, goa.FormatEmail, err2, err)
		}
	}
	if source.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, source.Email); err2 != nil {
			err = goa.InvalidFormatError(`.email`, source.Email, goa.FormatEmail, err2, err)
		}
	}
	tmp15 := map[string]interface{}{
		"bio":       source.Bio,
		"city":      source.City,
		"country":   source.Country,
		"email":     source.Email,
		"firstname": source.Firstname,
		"href":      source.Href,
		"id":        source.ID,
		"lastname":  source.Lastname,
		"role":      source.Role,
		"state":     source.State,
	}
	target = tmp15
	return
}

// MarshalUserLink validates and renders an instance of User into a interface{}
// using view "link".
func MarshalUserLink(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if source.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, source.Email); err2 != nil {
			err = goa.InvalidFormatError(`.email`, source.Email, goa.FormatEmail, err2, err)
		}
	}
	if source.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, source.Email); err2 != nil {
			err = goa.InvalidFormatError(`.email`, source.Email, goa.FormatEmail, err2, err)
		}
	}
	tmp16 := map[string]interface{}{
		"email": source.Email,
		"href":  source.Href,
		"id":    source.ID,
	}
	target = tmp16
	return
}

// UnmarshalUser unmarshals and validates a raw interface{} into an instance of User
func UnmarshalUser(source interface{}, inErr error) (target *User, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(User)
		if v, ok := val["bio"]; ok {
			var tmp17 string
			if val, ok := v.(string); ok {
				tmp17 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp17) > 500 {
					err = goa.InvalidLengthError(`load.Bio`, tmp17, len(tmp17), 500, false, err)
				}
			}
			target.Bio = tmp17
		}
		if v, ok := val["city"]; ok {
			var tmp18 string
			if val, ok := v.(string); ok {
				tmp18 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = tmp18
		}
		if v, ok := val["country"]; ok {
			var tmp19 string
			if val, ok := v.(string); ok {
				tmp19 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = tmp19
		}
		if v, ok := val["email"]; ok {
			var tmp20 string
			if val, ok := v.(string); ok {
				tmp20 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			if err == nil {
				if tmp20 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp20); err2 != nil {
						err = goa.InvalidFormatError(`load.Email`, tmp20, goa.FormatEmail, err2, err)
					}
				}
				if tmp20 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp20); err2 != nil {
						err = goa.InvalidFormatError(`load.Email`, tmp20, goa.FormatEmail, err2, err)
					}
				}
			}
			target.Email = tmp20
		}
		if v, ok := val["firstname"]; ok {
			var tmp21 string
			if val, ok := v.(string); ok {
				tmp21 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			target.Firstname = tmp21
		}
		if v, ok := val["href"]; ok {
			var tmp22 string
			if val, ok := v.(string); ok {
				tmp22 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp22
		}
		if v, ok := val["id"]; ok {
			var tmp23 int
			if f, ok := v.(float64); ok {
				tmp23 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp23
		}
		if v, ok := val["lastname"]; ok {
			var tmp24 string
			if val, ok := v.(string); ok {
				tmp24 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Lastname`, v, "string", err)
			}
			target.Lastname = tmp24
		}
		if v, ok := val["role"]; ok {
			var tmp25 string
			if val, ok := v.(string); ok {
				tmp25 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = tmp25
		}
		if v, ok := val["state"]; ok {
			var tmp26 string
			if val, ok := v.(string); ok {
				tmp26 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = tmp26
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// UserCollection media type
// Identifier: application/vnd.user+json; type=collection
type UserCollection []*User

// LoadUserCollection loads raw data into an instance of UserCollection
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUserCollection(raw interface{}) (res UserCollection, err error) {
	res, err = UnmarshalUserCollection(raw, err)
	return
}

// Dump produces raw data from an instance of UserCollection running all the
// validations. See LoadUserCollection for the definition of raw data.
func (mt UserCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp27 := range mt {
		var tmp28 map[string]interface{}
		tmp28, err = MarshalUser(tmp27, err)
		res[i] = tmp28
	}
	return
}

// Validate validates the media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Bio) > 500 {
			err = goa.InvalidLengthError(`response[*].bio`, e.Bio, len(e.Bio), 500, false, err)
		}
		if e.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, e.Email, goa.FormatEmail, err2, err)
			}
		}
		if e.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, e.Email, goa.FormatEmail, err2, err)
			}
		}
	}
	return
}

// MarshalUserCollection validates and renders an instance of UserCollection into a interface{}
// using view "default".
func MarshalUserCollection(source UserCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalUser(res, err)
	}
	return
}

// UnmarshalUserCollection unmarshals and validates a raw interface{} into an instance of UserCollection
func UnmarshalUserCollection(source interface{}, inErr error) (target UserCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*User, len(val))
		for tmp29, v := range val {
			target[tmp29], err = UnmarshalUser(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}
