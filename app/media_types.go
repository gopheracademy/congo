//************************************************************************//
// congo: Application Media Types
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

// A tenant account
// Identifier: application/vnd.congo.api.account+json
type Account struct {
	// API href of account
	Href string
	// ID of account
	ID int
	// Name of account
	Name string
}

// Account views
type AccountViewEnum string

const (
	// Account default view
	AccountDefaultView AccountViewEnum = "default"
	// Account link view
	AccountLinkView AccountViewEnum = "link"
)

// LoadAccount loads raw data into an instance of Account running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadAccount(raw interface{}) (*Account, error) {
	var err error
	var res *Account
	if val, ok := raw.(map[string]interface{}); ok {
		res = new(Account)
		if v, ok := val["href"]; ok {
			var tmp11 string
			if val, ok := v.(string); ok {
				tmp11 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp11
		}
		if v, ok := val["id"]; ok {
			var tmp12 int
			if f, ok := v.(float64); ok {
				tmp12 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp12
		}
		if v, ok := val["name"]; ok {
			var tmp13 string
			if val, ok := v.(string); ok {
				tmp13 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp13) < 2 {
					err = goa.InvalidLengthError(`.Name`, tmp13, 2, true, err)
				}
			}
			res.Name = tmp13
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "dictionary", err)
	}
	return res, err
}

// Dump produces raw data from an instance of Account running all the
// validations. See LoadAccount for the definition of raw data.
func (mt *Account) Dump(view AccountViewEnum) (map[string]interface{}, error) {
	var err error
	var res map[string]interface{}
	if view == AccountDefaultView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`default view.name`, mt.Name, 2, true, err)
		}
		tmp14 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp14
	}
	if view == AccountLinkView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`link view.name`, mt.Name, 2, true, err)
		}
		tmp15 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp15
	}
	return res, err
}

// Validate validates the media type instance.
func (mt *Account) Validate() (err error) {
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, 2, true, err)
	}
	return
}

// AccountCollection media type
// Identifier: application/vnd.congo.api.account+json; type=collection
type AccountCollection []*Account

// LoadAccountCollection loads raw data into an instance of AccountCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadAccountCollection(raw interface{}) (AccountCollection, error) {
	var err error
	var res AccountCollection
	if val, ok := raw.([]interface{}); ok {
		res = make([]*Account, len(val))
		for i, v := range val {
			var tmp16 *Account
			if val, ok := v.(map[string]interface{}); ok {
				tmp16 = new(Account)
				if v, ok := val["href"]; ok {
					var tmp17 string
					if val, ok := v.(string); ok {
						tmp17 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp16.Href = tmp17
				}
				if v, ok := val["id"]; ok {
					var tmp18 int
					if f, ok := v.(float64); ok {
						tmp18 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp16.ID = tmp18
				}
				if v, ok := val["name"]; ok {
					var tmp19 string
					if val, ok := v.(string); ok {
						tmp19 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp19) < 2 {
							err = goa.InvalidLengthError(`[*].Name`, tmp19, 2, true, err)
						}
					}
					tmp16.Name = tmp19
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "dictionary", err)
			}
			res[i] = tmp16
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "array", err)
	}
	return res, err
}

// Dump produces raw data from an instance of AccountCollection running all the
// validations. See LoadAccountCollection for the definition of raw data.
func (mt AccountCollection) Dump() ([]map[string]interface{}, error) {
	var err error
	var res []map[string]interface{}
	res = make([]map[string]interface{}, len(mt))
	for i, tmp20 := range mt {
		if len(tmp20.Name) < 2 {
			err = goa.InvalidLengthError(`default view[*].name`, tmp20.Name, 2, true, err)
		}
		tmp21 := map[string]interface{}{
			"href": tmp20.Href,
			"id":   tmp20.ID,
			"name": tmp20.Name,
		}
		res[i] = tmp21
	}
	return res, err
}

// Validate validates the media type instance.
func (mt AccountCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, 2, true, err)
		}
	}
	return
}

// A recurring event or conference
// Identifier: application/vnd.congo.api.series+json
type Series struct {
	// Account that owns bottle
	Account *Account
	// API href of series
	Href string
	// ID of series
	ID   int
	Name string
}

// Series views
type SeriesViewEnum string

const (
	// Series default view
	SeriesDefaultView SeriesViewEnum = "default"
	// Series full view
	SeriesFullView SeriesViewEnum = "full"
	// Series tiny view
	SeriesTinyView SeriesViewEnum = "tiny"
)

// LoadSeries loads raw data into an instance of Series running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadSeries(raw interface{}) (*Series, error) {
	var err error
	var res *Series
	if val, ok := raw.(map[string]interface{}); ok {
		res = new(Series)
		if v, ok := val["account"]; ok {
			var tmp22 *Account
			if val, ok := v.(map[string]interface{}); ok {
				tmp22 = new(Account)
				if v, ok := val["href"]; ok {
					var tmp23 string
					if val, ok := v.(string); ok {
						tmp23 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.Href`, v, "string", err)
					}
					tmp22.Href = tmp23
				}
				if v, ok := val["id"]; ok {
					var tmp24 int
					if f, ok := v.(float64); ok {
						tmp24 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.ID`, v, "int", err)
					}
					tmp22.ID = tmp24
				}
				if v, ok := val["name"]; ok {
					var tmp25 string
					if val, ok := v.(string); ok {
						tmp25 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp25) < 2 {
							err = goa.InvalidLengthError(`.Account.Name`, tmp25, 2, true, err)
						}
					}
					tmp22.Name = tmp25
				}
			} else {
				err = goa.InvalidAttributeTypeError(`.Account`, v, "dictionary", err)
			}
			res.Account = tmp22
		}
		if v, ok := val["href"]; ok {
			var tmp26 string
			if val, ok := v.(string); ok {
				tmp26 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp26
		}
		if v, ok := val["id"]; ok {
			var tmp27 int
			if f, ok := v.(float64); ok {
				tmp27 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp27
		}
		if v, ok := val["name"]; ok {
			var tmp28 string
			if val, ok := v.(string); ok {
				tmp28 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp28) < 2 {
					err = goa.InvalidLengthError(`.Name`, tmp28, 2, true, err)
				}
			}
			res.Name = tmp28
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "dictionary", err)
	}
	return res, err
}

// Dump produces raw data from an instance of Series running all the
// validations. See LoadSeries for the definition of raw data.
func (mt *Series) Dump(view SeriesViewEnum) (map[string]interface{}, error) {
	var err error
	var res map[string]interface{}
	if view == SeriesDefaultView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`default view.name`, mt.Name, 2, true, err)
		}
		tmp30 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp30
		if err == nil {
			links := make(map[string]interface{})
			if len(mt.Account.Name) < 2 {
				err = goa.InvalidLengthError(`link account.name`, mt.Account.Name, 2, true, err)
			}
			tmp29 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			links["account"] = tmp29
			res["links"] = links
		}
	}
	if view == SeriesFullView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`full view.name`, mt.Name, 2, true, err)
		}
		tmp32 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		if mt.Account != nil {
			if len(mt.Account.Name) < 2 {
				err = goa.InvalidLengthError(`full view.Account.name`, mt.Account.Name, 2, true, err)
			}
			tmp33 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			tmp32["account"] = tmp33
		}
		res = tmp32
		if err == nil {
			links := make(map[string]interface{})
			if len(mt.Account.Name) < 2 {
				err = goa.InvalidLengthError(`link account.name`, mt.Account.Name, 2, true, err)
			}
			tmp31 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			links["account"] = tmp31
			res["links"] = links
		}
	}
	if view == SeriesTinyView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`tiny view.name`, mt.Name, 2, true, err)
		}
		tmp35 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp35
		if err == nil {
			links := make(map[string]interface{})
			if len(mt.Account.Name) < 2 {
				err = goa.InvalidLengthError(`link account.name`, mt.Account.Name, 2, true, err)
			}
			tmp34 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			links["account"] = tmp34
			res["links"] = links
		}
	}
	return res, err
}

// Validate validates the media type instance.
func (mt *Series) Validate() (err error) {
	if len(mt.Account.Name) < 2 {
		err = goa.InvalidLengthError(`response.account.name`, mt.Account.Name, 2, true, err)
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, 2, true, err)
	}
	return
}

// SeriesCollection media type
// Identifier: application/vnd.congo.api.series+json; type=collection
type SeriesCollection []*Series

// SeriesCollection views
type SeriesCollectionViewEnum string

const (
	// SeriesCollection default view
	SeriesCollectionDefaultView SeriesCollectionViewEnum = "default"
	// SeriesCollection tiny view
	SeriesCollectionTinyView SeriesCollectionViewEnum = "tiny"
)

// LoadSeriesCollection loads raw data into an instance of SeriesCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadSeriesCollection(raw interface{}) (SeriesCollection, error) {
	var err error
	var res SeriesCollection
	if val, ok := raw.([]interface{}); ok {
		res = make([]*Series, len(val))
		for i, v := range val {
			var tmp36 *Series
			if val, ok := v.(map[string]interface{}); ok {
				tmp36 = new(Series)
				if v, ok := val["account"]; ok {
					var tmp37 *Account
					if val, ok := v.(map[string]interface{}); ok {
						tmp37 = new(Account)
						if v, ok := val["href"]; ok {
							var tmp38 string
							if val, ok := v.(string); ok {
								tmp38 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.Href`, v, "string", err)
							}
							tmp37.Href = tmp38
						}
						if v, ok := val["id"]; ok {
							var tmp39 int
							if f, ok := v.(float64); ok {
								tmp39 = int(f)
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.ID`, v, "int", err)
							}
							tmp37.ID = tmp39
						}
						if v, ok := val["name"]; ok {
							var tmp40 string
							if val, ok := v.(string); ok {
								tmp40 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.Name`, v, "string", err)
							}
							if err == nil {
								if len(tmp40) < 2 {
									err = goa.InvalidLengthError(`[*].Account.Name`, tmp40, 2, true, err)
								}
							}
							tmp37.Name = tmp40
						}
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Account`, v, "dictionary", err)
					}
					tmp36.Account = tmp37
				}
				if v, ok := val["href"]; ok {
					var tmp41 string
					if val, ok := v.(string); ok {
						tmp41 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp36.Href = tmp41
				}
				if v, ok := val["id"]; ok {
					var tmp42 int
					if f, ok := v.(float64); ok {
						tmp42 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp36.ID = tmp42
				}
				if v, ok := val["name"]; ok {
					var tmp43 string
					if val, ok := v.(string); ok {
						tmp43 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp43) < 2 {
							err = goa.InvalidLengthError(`[*].Name`, tmp43, 2, true, err)
						}
					}
					tmp36.Name = tmp43
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "dictionary", err)
			}
			res[i] = tmp36
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "array", err)
	}
	return res, err
}

// Dump produces raw data from an instance of SeriesCollection running all the
// validations. See LoadSeriesCollection for the definition of raw data.
func (mt SeriesCollection) Dump(view SeriesCollectionViewEnum) ([]map[string]interface{}, error) {
	var err error
	var res []map[string]interface{}
	if view == SeriesCollectionDefaultView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp44 := range mt {
			if len(tmp44.Name) < 2 {
				err = goa.InvalidLengthError(`default view[*].name`, tmp44.Name, 2, true, err)
			}
			tmp46 := map[string]interface{}{
				"href": tmp44.Href,
				"id":   tmp44.ID,
				"name": tmp44.Name,
			}
			res[i] = tmp46
			if err == nil {
				links := make(map[string]interface{})
				if len(tmp44.Account.Name) < 2 {
					err = goa.InvalidLengthError(`link account.name`, tmp44.Account.Name, 2, true, err)
				}
				tmp45 := map[string]interface{}{
					"href": tmp44.Account.Href,
					"id":   tmp44.Account.ID,
					"name": tmp44.Account.Name,
				}
				links["account"] = tmp45
				res[i]["links"] = links
			}
		}
	}
	if view == SeriesCollectionTinyView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp47 := range mt {
			if len(tmp47.Name) < 2 {
				err = goa.InvalidLengthError(`tiny view[*].name`, tmp47.Name, 2, true, err)
			}
			tmp49 := map[string]interface{}{
				"href": tmp47.Href,
				"id":   tmp47.ID,
				"name": tmp47.Name,
			}
			res[i] = tmp49
			if err == nil {
				links := make(map[string]interface{})
				if len(tmp47.Account.Name) < 2 {
					err = goa.InvalidLengthError(`link account.name`, tmp47.Account.Name, 2, true, err)
				}
				tmp48 := map[string]interface{}{
					"href": tmp47.Account.Href,
					"id":   tmp47.Account.ID,
					"name": tmp47.Account.Name,
				}
				links["account"] = tmp48
				res[i]["links"] = links
			}
		}
	}
	return res, err
}

// Validate validates the media type instance.
func (mt SeriesCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Account.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].account.name`, e.Account.Name, 2, true, err)
		}
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, 2, true, err)
		}
	}
	return
}

// A user belonging to a tenant account
// Identifier: application/vnd.congo.api.user+json
type User struct {
	// Email address of user
	Email string
	// First name of user
	FirstName string
	// API href of user
	Href string
	// ID of user
	ID int
	// Last name of user
	LastName string
}

// User views
type UserViewEnum string

const (
	// User default view
	UserDefaultView UserViewEnum = "default"
	// User link view
	UserLinkView UserViewEnum = "link"
)

// LoadUser loads raw data into an instance of User running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUser(raw interface{}) (*User, error) {
	var err error
	var res *User
	if val, ok := raw.(map[string]interface{}); ok {
		res = new(User)
		if v, ok := val["email"]; ok {
			var tmp50 string
			if val, ok := v.(string); ok {
				tmp50 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Email`, v, "string", err)
			}
			if err == nil {
				if tmp50 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp50); err2 != nil {
						err = goa.InvalidFormatError(`.Email`, tmp50, goa.FormatEmail, err2, err)
					}
				}
			}
			res.Email = tmp50
		}
		if v, ok := val["first_name"]; ok {
			var tmp51 string
			if val, ok := v.(string); ok {
				tmp51 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.FirstName`, v, "string", err)
			}
			res.FirstName = tmp51
		}
		if v, ok := val["href"]; ok {
			var tmp52 string
			if val, ok := v.(string); ok {
				tmp52 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp52
		}
		if v, ok := val["id"]; ok {
			var tmp53 int
			if f, ok := v.(float64); ok {
				tmp53 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp53
		}
		if v, ok := val["last_name"]; ok {
			var tmp54 string
			if val, ok := v.(string); ok {
				tmp54 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.LastName`, v, "string", err)
			}
			res.LastName = tmp54
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "dictionary", err)
	}
	return res, err
}

// Dump produces raw data from an instance of User running all the
// validations. See LoadUser for the definition of raw data.
func (mt *User) Dump(view UserViewEnum) (map[string]interface{}, error) {
	var err error
	var res map[string]interface{}
	if view == UserDefaultView {
		if mt.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
				err = goa.InvalidFormatError(`default view.email`, mt.Email, goa.FormatEmail, err2, err)
			}
		}
		tmp55 := map[string]interface{}{
			"email":      mt.Email,
			"first_name": mt.FirstName,
			"href":       mt.Href,
			"id":         mt.ID,
			"last_name":  mt.LastName,
		}
		res = tmp55
	}
	if view == UserLinkView {
		if mt.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
				err = goa.InvalidFormatError(`link view.email`, mt.Email, goa.FormatEmail, err2, err)
			}
		}
		tmp56 := map[string]interface{}{
			"email": mt.Email,
			"href":  mt.Href,
			"id":    mt.ID,
		}
		res = tmp56
	}
	return res, err
}

// Validate validates the media type instance.
func (mt *User) Validate() (err error) {
	if mt.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}

// UserCollection media type
// Identifier: application/vnd.congo.api.user+json; type=collection
type UserCollection []*User

// LoadUserCollection loads raw data into an instance of UserCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUserCollection(raw interface{}) (UserCollection, error) {
	var err error
	var res UserCollection
	if val, ok := raw.([]interface{}); ok {
		res = make([]*User, len(val))
		for i, v := range val {
			var tmp57 *User
			if val, ok := v.(map[string]interface{}); ok {
				tmp57 = new(User)
				if v, ok := val["email"]; ok {
					var tmp58 string
					if val, ok := v.(string); ok {
						tmp58 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Email`, v, "string", err)
					}
					if err == nil {
						if tmp58 != "" {
							if err2 := goa.ValidateFormat(goa.FormatEmail, tmp58); err2 != nil {
								err = goa.InvalidFormatError(`[*].Email`, tmp58, goa.FormatEmail, err2, err)
							}
						}
					}
					tmp57.Email = tmp58
				}
				if v, ok := val["first_name"]; ok {
					var tmp59 string
					if val, ok := v.(string); ok {
						tmp59 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].FirstName`, v, "string", err)
					}
					tmp57.FirstName = tmp59
				}
				if v, ok := val["href"]; ok {
					var tmp60 string
					if val, ok := v.(string); ok {
						tmp60 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp57.Href = tmp60
				}
				if v, ok := val["id"]; ok {
					var tmp61 int
					if f, ok := v.(float64); ok {
						tmp61 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp57.ID = tmp61
				}
				if v, ok := val["last_name"]; ok {
					var tmp62 string
					if val, ok := v.(string); ok {
						tmp62 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].LastName`, v, "string", err)
					}
					tmp57.LastName = tmp62
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "dictionary", err)
			}
			res[i] = tmp57
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "array", err)
	}
	return res, err
}

// Dump produces raw data from an instance of UserCollection running all the
// validations. See LoadUserCollection for the definition of raw data.
func (mt UserCollection) Dump() ([]map[string]interface{}, error) {
	var err error
	var res []map[string]interface{}
	res = make([]map[string]interface{}, len(mt))
	for i, tmp63 := range mt {
		if tmp63.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, tmp63.Email); err2 != nil {
				err = goa.InvalidFormatError(`default view[*].email`, tmp63.Email, goa.FormatEmail, err2, err)
			}
		}
		tmp64 := map[string]interface{}{
			"email":      tmp63.Email,
			"first_name": tmp63.FirstName,
			"href":       tmp63.Href,
			"id":         tmp63.ID,
			"last_name":  tmp63.LastName,
		}
		res[i] = tmp64
	}
	return res, err
}

// Validate validates the media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, e.Email, goa.FormatEmail, err2, err)
			}
		}
	}
	return
}
