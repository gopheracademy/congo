//************************************************************************//
// congo: Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=/home/bketelsen/src/github.com/bketelsen/congo
// --design=github.com/bketelsen/congo/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/raphael/goa"

// A tenant account
// Identifier: application/vnd.congo.api.account
type Account struct {
	// Date of creation
	CreatedAt string
	// Email of account owner
	CreatedBy string
	// API href of account
	Href string
	// ID of account
	ID int
	// Name of account
	Name string
}

// object views
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
		if v, ok := val["created_at"]; ok {
			var tmp11 string
			if val, ok := v.(string); ok {
				tmp11 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.CreatedAt`, v, "string", err)
			}
			if err == nil {
				if tmp11 != "" {
					if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp11); err2 != nil {
						err = goa.InvalidFormatError(`.CreatedAt`, tmp11, goa.FormatDateTime, err2, err)
					}
				}
			}
			res.CreatedAt = tmp11
		}
		if v, ok := val["created_by"]; ok {
			var tmp12 string
			if val, ok := v.(string); ok {
				tmp12 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.CreatedBy`, v, "string", err)
			}
			if err == nil {
				if tmp12 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp12); err2 != nil {
						err = goa.InvalidFormatError(`.CreatedBy`, tmp12, goa.FormatEmail, err2, err)
					}
				}
			}
			res.CreatedBy = tmp12
		}
		if v, ok := val["href"]; ok {
			var tmp13 string
			if val, ok := v.(string); ok {
				tmp13 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp13
		}
		if v, ok := val["id"]; ok {
			var tmp14 int
			if f, ok := v.(float64); ok {
				tmp14 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp14
		}
		if v, ok := val["name"]; ok {
			var tmp15 string
			if val, ok := v.(string); ok {
				tmp15 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Name`, v, "string", err)
			}
			res.Name = tmp15
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "map[string]interface{}", err)
	}
	return res, err
}

// Dump produces raw data from an instance of Account running all the
// validations. See LoadAccount for the definition of raw data.
func (mt *Account) Dump(view AccountViewEnum) (map[string]interface{}, error) {
	var err error
	var res map[string]interface{}
	if view == AccountDefaultView {
		if mt.CreatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, mt.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`default view.created_at`, mt.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if mt.CreatedBy != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, mt.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`default view.created_by`, mt.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
		tmp16 := map[string]interface{}{
			"created_at": mt.CreatedAt,
			"created_by": mt.CreatedBy,
			"href":       mt.Href,
			"id":         mt.ID,
			"name":       mt.Name,
		}
		res = tmp16
	}
	if view == AccountLinkView {
		tmp17 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp17
	}
	return res, err
}

// Validate validates the media type instance.
func (mt *Account) Validate() (err error) {
	if mt.CreatedAt != "" {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, mt.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if mt.CreatedBy != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, mt.CreatedBy); err2 != nil {
			err = goa.InvalidFormatError(`response.created_by`, mt.CreatedBy, goa.FormatEmail, err2, err)
		}
	}
	return
}

// A recurring event or conference
// Identifier: application/vnd.congo.api.series
type Series struct {
	// Account that owns bottle
	Account *Account
	// Date of creation
	CreatedAt string
	// API href of series
	Href string
	// ID of series
	ID   int
	Name string
	// Date of last update
	UpdatedAt string
}

// object views
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
			var tmp18 *Account
			if val, ok := v.(map[string]interface{}); ok {
				tmp18 = new(Account)
				if v, ok := val["created_at"]; ok {
					var tmp19 string
					if val, ok := v.(string); ok {
						tmp19 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.CreatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp19 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp19); err2 != nil {
								err = goa.InvalidFormatError(`.Account.CreatedAt`, tmp19, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp18.CreatedAt = tmp19
				}
				if v, ok := val["created_by"]; ok {
					var tmp20 string
					if val, ok := v.(string); ok {
						tmp20 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.CreatedBy`, v, "string", err)
					}
					if err == nil {
						if tmp20 != "" {
							if err2 := goa.ValidateFormat(goa.FormatEmail, tmp20); err2 != nil {
								err = goa.InvalidFormatError(`.Account.CreatedBy`, tmp20, goa.FormatEmail, err2, err)
							}
						}
					}
					tmp18.CreatedBy = tmp20
				}
				if v, ok := val["href"]; ok {
					var tmp21 string
					if val, ok := v.(string); ok {
						tmp21 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.Href`, v, "string", err)
					}
					tmp18.Href = tmp21
				}
				if v, ok := val["id"]; ok {
					var tmp22 int
					if f, ok := v.(float64); ok {
						tmp22 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.ID`, v, "int", err)
					}
					tmp18.ID = tmp22
				}
				if v, ok := val["name"]; ok {
					var tmp23 string
					if val, ok := v.(string); ok {
						tmp23 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.Name`, v, "string", err)
					}
					tmp18.Name = tmp23
				}
			} else {
				err = goa.InvalidAttributeTypeError(`.Account`, v, "map[string]interface{}", err)
			}
			res.Account = tmp18
		}
		if v, ok := val["created_at"]; ok {
			var tmp24 string
			if val, ok := v.(string); ok {
				tmp24 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.CreatedAt`, v, "string", err)
			}
			if err == nil {
				if tmp24 != "" {
					if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp24); err2 != nil {
						err = goa.InvalidFormatError(`.CreatedAt`, tmp24, goa.FormatDateTime, err2, err)
					}
				}
			}
			res.CreatedAt = tmp24
		}
		if v, ok := val["href"]; ok {
			var tmp25 string
			if val, ok := v.(string); ok {
				tmp25 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp25
		}
		if v, ok := val["id"]; ok {
			var tmp26 int
			if f, ok := v.(float64); ok {
				tmp26 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp26
		}
		if v, ok := val["name"]; ok {
			var tmp27 string
			if val, ok := v.(string); ok {
				tmp27 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp27) < 2 {
					err = goa.InvalidLengthError(`.Name`, tmp27, 2, true, err)
				}
			}
			res.Name = tmp27
		}
		if v, ok := val["updated_at"]; ok {
			var tmp28 string
			if val, ok := v.(string); ok {
				tmp28 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.UpdatedAt`, v, "string", err)
			}
			if err == nil {
				if tmp28 != "" {
					if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp28); err2 != nil {
						err = goa.InvalidFormatError(`.UpdatedAt`, tmp28, goa.FormatDateTime, err2, err)
					}
				}
			}
			res.UpdatedAt = tmp28
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "map[string]interface{}", err)
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
		if mt.CreatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, mt.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`full view.created_at`, mt.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`full view.name`, mt.Name, 2, true, err)
		}
		if mt.UpdatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, mt.UpdatedAt); err2 != nil {
				err = goa.InvalidFormatError(`full view.updated_at`, mt.UpdatedAt, goa.FormatDateTime, err2, err)
			}
		}
		tmp32 := map[string]interface{}{
			"created_at": mt.CreatedAt,
			"href":       mt.Href,
			"id":         mt.ID,
			"name":       mt.Name,
			"updated_at": mt.UpdatedAt,
		}
		if mt.Account != nil {
			if mt.Account.CreatedAt != "" {
				if err2 := goa.ValidateFormat(goa.FormatDateTime, mt.Account.CreatedAt); err2 != nil {
					err = goa.InvalidFormatError(`full view.Account.created_at`, mt.Account.CreatedAt, goa.FormatDateTime, err2, err)
				}
			}
			if mt.Account.CreatedBy != "" {
				if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Account.CreatedBy); err2 != nil {
					err = goa.InvalidFormatError(`full view.Account.created_by`, mt.Account.CreatedBy, goa.FormatEmail, err2, err)
				}
			}
			tmp33 := map[string]interface{}{
				"created_at": mt.Account.CreatedAt,
				"created_by": mt.Account.CreatedBy,
				"href":       mt.Account.Href,
				"id":         mt.Account.ID,
				"name":       mt.Account.Name,
			}
			tmp32["account"] = tmp33
		}
		res = tmp32
		if err == nil {
			links := make(map[string]interface{})
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
	if mt.Account.CreatedAt != "" {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, mt.Account.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.account.created_at`, mt.Account.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if mt.Account.CreatedBy != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Account.CreatedBy); err2 != nil {
			err = goa.InvalidFormatError(`response.account.created_by`, mt.Account.CreatedBy, goa.FormatEmail, err2, err)
		}
	}
	if mt.CreatedAt != "" {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, mt.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, 2, true, err)
	}
	if mt.UpdatedAt != "" {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, mt.UpdatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.updated_at`, mt.UpdatedAt, goa.FormatDateTime, err2, err)
		}
	}
	return
}

// SeriesCollection media type
// Identifier: application/vnd.congo.api.series; type=collection
type SeriesCollection []*Series

// array views
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
						if v, ok := val["created_at"]; ok {
							var tmp38 string
							if val, ok := v.(string); ok {
								tmp38 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.CreatedAt`, v, "string", err)
							}
							if err == nil {
								if tmp38 != "" {
									if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp38); err2 != nil {
										err = goa.InvalidFormatError(`[*].Account.CreatedAt`, tmp38, goa.FormatDateTime, err2, err)
									}
								}
							}
							tmp37.CreatedAt = tmp38
						}
						if v, ok := val["created_by"]; ok {
							var tmp39 string
							if val, ok := v.(string); ok {
								tmp39 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.CreatedBy`, v, "string", err)
							}
							if err == nil {
								if tmp39 != "" {
									if err2 := goa.ValidateFormat(goa.FormatEmail, tmp39); err2 != nil {
										err = goa.InvalidFormatError(`[*].Account.CreatedBy`, tmp39, goa.FormatEmail, err2, err)
									}
								}
							}
							tmp37.CreatedBy = tmp39
						}
						if v, ok := val["href"]; ok {
							var tmp40 string
							if val, ok := v.(string); ok {
								tmp40 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.Href`, v, "string", err)
							}
							tmp37.Href = tmp40
						}
						if v, ok := val["id"]; ok {
							var tmp41 int
							if f, ok := v.(float64); ok {
								tmp41 = int(f)
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.ID`, v, "int", err)
							}
							tmp37.ID = tmp41
						}
						if v, ok := val["name"]; ok {
							var tmp42 string
							if val, ok := v.(string); ok {
								tmp42 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.Name`, v, "string", err)
							}
							tmp37.Name = tmp42
						}
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Account`, v, "map[string]interface{}", err)
					}
					tmp36.Account = tmp37
				}
				if v, ok := val["created_at"]; ok {
					var tmp43 string
					if val, ok := v.(string); ok {
						tmp43 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].CreatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp43 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp43); err2 != nil {
								err = goa.InvalidFormatError(`[*].CreatedAt`, tmp43, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp36.CreatedAt = tmp43
				}
				if v, ok := val["href"]; ok {
					var tmp44 string
					if val, ok := v.(string); ok {
						tmp44 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp36.Href = tmp44
				}
				if v, ok := val["id"]; ok {
					var tmp45 int
					if f, ok := v.(float64); ok {
						tmp45 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp36.ID = tmp45
				}
				if v, ok := val["name"]; ok {
					var tmp46 string
					if val, ok := v.(string); ok {
						tmp46 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp46) < 2 {
							err = goa.InvalidLengthError(`[*].Name`, tmp46, 2, true, err)
						}
					}
					tmp36.Name = tmp46
				}
				if v, ok := val["updated_at"]; ok {
					var tmp47 string
					if val, ok := v.(string); ok {
						tmp47 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].UpdatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp47 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp47); err2 != nil {
								err = goa.InvalidFormatError(`[*].UpdatedAt`, tmp47, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp36.UpdatedAt = tmp47
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "map[string]interface{}", err)
			}
			res[i] = tmp36
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "[]interface{}", err)
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
		for i, tmp48 := range mt {
			if len(tmp48.Name) < 2 {
				err = goa.InvalidLengthError(`default view[*].name`, tmp48.Name, 2, true, err)
			}
			tmp50 := map[string]interface{}{
				"href": tmp48.Href,
				"id":   tmp48.ID,
				"name": tmp48.Name,
			}
			res[i] = tmp50
			if err == nil {
				links := make(map[string]interface{})
				tmp49 := map[string]interface{}{
					"href": tmp48.Account.Href,
					"id":   tmp48.Account.ID,
					"name": tmp48.Account.Name,
				}
				links["account"] = tmp49
				res[i]["links"] = links
			}
		}
	}
	if view == SeriesCollectionTinyView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp51 := range mt {
			if len(tmp51.Name) < 2 {
				err = goa.InvalidLengthError(`tiny view[*].name`, tmp51.Name, 2, true, err)
			}
			tmp53 := map[string]interface{}{
				"href": tmp51.Href,
				"id":   tmp51.ID,
				"name": tmp51.Name,
			}
			res[i] = tmp53
			if err == nil {
				links := make(map[string]interface{})
				tmp52 := map[string]interface{}{
					"href": tmp51.Account.Href,
					"id":   tmp51.Account.ID,
					"name": tmp51.Account.Name,
				}
				links["account"] = tmp52
				res[i]["links"] = links
			}
		}
	}
	return res, err
}

// Validate validates the media type instance.
func (mt SeriesCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Account.CreatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, e.Account.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].account.created_at`, e.Account.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if e.Account.CreatedBy != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, e.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response[*].account.created_by`, e.Account.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
		if e.CreatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, e.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].created_at`, e.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, 2, true, err)
		}
		if e.UpdatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, e.UpdatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].updated_at`, e.UpdatedAt, goa.FormatDateTime, err2, err)
			}
		}
	}
	return
}

// A user belonging to a tenant account
// Identifier: application/vnd.congo.api.user
type User struct {
	// Date of creation
	CreatedAt string
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

// object views
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
		if v, ok := val["created_at"]; ok {
			var tmp54 string
			if val, ok := v.(string); ok {
				tmp54 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.CreatedAt`, v, "string", err)
			}
			if err == nil {
				if tmp54 != "" {
					if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp54); err2 != nil {
						err = goa.InvalidFormatError(`.CreatedAt`, tmp54, goa.FormatDateTime, err2, err)
					}
				}
			}
			res.CreatedAt = tmp54
		}
		if v, ok := val["email"]; ok {
			var tmp55 string
			if val, ok := v.(string); ok {
				tmp55 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Email`, v, "string", err)
			}
			if err == nil {
				if tmp55 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp55); err2 != nil {
						err = goa.InvalidFormatError(`.Email`, tmp55, goa.FormatEmail, err2, err)
					}
				}
			}
			res.Email = tmp55
		}
		if v, ok := val["first_name"]; ok {
			var tmp56 string
			if val, ok := v.(string); ok {
				tmp56 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.FirstName`, v, "string", err)
			}
			res.FirstName = tmp56
		}
		if v, ok := val["href"]; ok {
			var tmp57 string
			if val, ok := v.(string); ok {
				tmp57 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp57
		}
		if v, ok := val["id"]; ok {
			var tmp58 int
			if f, ok := v.(float64); ok {
				tmp58 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp58
		}
		if v, ok := val["last_name"]; ok {
			var tmp59 string
			if val, ok := v.(string); ok {
				tmp59 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.LastName`, v, "string", err)
			}
			res.LastName = tmp59
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "map[string]interface{}", err)
	}
	return res, err
}

// Dump produces raw data from an instance of User running all the
// validations. See LoadUser for the definition of raw data.
func (mt *User) Dump(view UserViewEnum) (map[string]interface{}, error) {
	var err error
	var res map[string]interface{}
	if view == UserDefaultView {
		if mt.CreatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, mt.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`default view.created_at`, mt.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if mt.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
				err = goa.InvalidFormatError(`default view.email`, mt.Email, goa.FormatEmail, err2, err)
			}
		}
		tmp60 := map[string]interface{}{
			"created_at": mt.CreatedAt,
			"email":      mt.Email,
			"first_name": mt.FirstName,
			"href":       mt.Href,
			"id":         mt.ID,
			"last_name":  mt.LastName,
		}
		res = tmp60
	}
	if view == UserLinkView {
		if mt.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
				err = goa.InvalidFormatError(`link view.email`, mt.Email, goa.FormatEmail, err2, err)
			}
		}
		tmp61 := map[string]interface{}{
			"email": mt.Email,
			"href":  mt.Href,
			"id":    mt.ID,
		}
		res = tmp61
	}
	return res, err
}

// Validate validates the media type instance.
func (mt *User) Validate() (err error) {
	if mt.CreatedAt != "" {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, mt.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if mt.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}

// UserCollection media type
// Identifier: application/vnd.congo.api.user; type=collection
type UserCollection []*User

// array views
type UserCollectionViewEnum string

const (
	// UserCollection default view
	UserCollectionDefaultView UserCollectionViewEnum = "default"
)

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
			var tmp62 *User
			if val, ok := v.(map[string]interface{}); ok {
				tmp62 = new(User)
				if v, ok := val["created_at"]; ok {
					var tmp63 string
					if val, ok := v.(string); ok {
						tmp63 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].CreatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp63 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp63); err2 != nil {
								err = goa.InvalidFormatError(`[*].CreatedAt`, tmp63, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp62.CreatedAt = tmp63
				}
				if v, ok := val["email"]; ok {
					var tmp64 string
					if val, ok := v.(string); ok {
						tmp64 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Email`, v, "string", err)
					}
					if err == nil {
						if tmp64 != "" {
							if err2 := goa.ValidateFormat(goa.FormatEmail, tmp64); err2 != nil {
								err = goa.InvalidFormatError(`[*].Email`, tmp64, goa.FormatEmail, err2, err)
							}
						}
					}
					tmp62.Email = tmp64
				}
				if v, ok := val["first_name"]; ok {
					var tmp65 string
					if val, ok := v.(string); ok {
						tmp65 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].FirstName`, v, "string", err)
					}
					tmp62.FirstName = tmp65
				}
				if v, ok := val["href"]; ok {
					var tmp66 string
					if val, ok := v.(string); ok {
						tmp66 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp62.Href = tmp66
				}
				if v, ok := val["id"]; ok {
					var tmp67 int
					if f, ok := v.(float64); ok {
						tmp67 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp62.ID = tmp67
				}
				if v, ok := val["last_name"]; ok {
					var tmp68 string
					if val, ok := v.(string); ok {
						tmp68 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].LastName`, v, "string", err)
					}
					tmp62.LastName = tmp68
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "map[string]interface{}", err)
			}
			res[i] = tmp62
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "[]interface{}", err)
	}
	return res, err
}

// Dump produces raw data from an instance of UserCollection running all the
// validations. See LoadUserCollection for the definition of raw data.
func (mt UserCollection) Dump() ([]map[string]interface{}, error) {
	var err error
	var res []map[string]interface{}
	res = make([]map[string]interface{}, len(mt))
	for i, tmp69 := range mt {
		if tmp69.CreatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp69.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`default view[*].created_at`, tmp69.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if tmp69.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, tmp69.Email); err2 != nil {
				err = goa.InvalidFormatError(`default view[*].email`, tmp69.Email, goa.FormatEmail, err2, err)
			}
		}
		tmp70 := map[string]interface{}{
			"created_at": tmp69.CreatedAt,
			"email":      tmp69.Email,
			"first_name": tmp69.FirstName,
			"href":       tmp69.Href,
			"id":         tmp69.ID,
			"last_name":  tmp69.LastName,
		}
		res[i] = tmp70
	}
	return res, err
}

// Validate validates the media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.CreatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, e.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].created_at`, e.CreatedAt, goa.FormatDateTime, err2, err)
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
