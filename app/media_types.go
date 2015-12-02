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
			if err == nil {
				if len(tmp15) < 2 {
					err = goa.InvalidLengthError(`.Name`, tmp15, 2, true, err)
				}
			}
			res.Name = tmp15
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
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`default view.name`, mt.Name, 2, true, err)
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
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`link view.name`, mt.Name, 2, true, err)
		}
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
			var tmp18 *Account
			if val, ok := v.(map[string]interface{}); ok {
				tmp18 = new(Account)
				if v, ok := val["created_at"]; ok {
					var tmp19 string
					if val, ok := v.(string); ok {
						tmp19 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].CreatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp19 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp19); err2 != nil {
								err = goa.InvalidFormatError(`[*].CreatedAt`, tmp19, goa.FormatDateTime, err2, err)
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
						err = goa.InvalidAttributeTypeError(`[*].CreatedBy`, v, "string", err)
					}
					if err == nil {
						if tmp20 != "" {
							if err2 := goa.ValidateFormat(goa.FormatEmail, tmp20); err2 != nil {
								err = goa.InvalidFormatError(`[*].CreatedBy`, tmp20, goa.FormatEmail, err2, err)
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
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp18.Href = tmp21
				}
				if v, ok := val["id"]; ok {
					var tmp22 int
					if f, ok := v.(float64); ok {
						tmp22 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp18.ID = tmp22
				}
				if v, ok := val["name"]; ok {
					var tmp23 string
					if val, ok := v.(string); ok {
						tmp23 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp23) < 2 {
							err = goa.InvalidLengthError(`[*].Name`, tmp23, 2, true, err)
						}
					}
					tmp18.Name = tmp23
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "dictionary", err)
			}
			res[i] = tmp18
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
	for i, tmp24 := range mt {
		if tmp24.CreatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp24.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`default view[*].created_at`, tmp24.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if tmp24.CreatedBy != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, tmp24.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`default view[*].created_by`, tmp24.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
		if len(tmp24.Name) < 2 {
			err = goa.InvalidLengthError(`default view[*].name`, tmp24.Name, 2, true, err)
		}
		tmp25 := map[string]interface{}{
			"created_at": tmp24.CreatedAt,
			"created_by": tmp24.CreatedBy,
			"href":       tmp24.Href,
			"id":         tmp24.ID,
			"name":       tmp24.Name,
		}
		res[i] = tmp25
	}
	return res, err
}

// Validate validates the media type instance.
func (mt AccountCollection) Validate() (err error) {
	for _, e := range mt {
		if e.CreatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, e.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].created_at`, e.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if e.CreatedBy != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, e.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response[*].created_by`, e.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
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
			var tmp26 *Account
			if val, ok := v.(map[string]interface{}); ok {
				tmp26 = new(Account)
				if v, ok := val["created_at"]; ok {
					var tmp27 string
					if val, ok := v.(string); ok {
						tmp27 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.CreatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp27 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp27); err2 != nil {
								err = goa.InvalidFormatError(`.Account.CreatedAt`, tmp27, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp26.CreatedAt = tmp27
				}
				if v, ok := val["created_by"]; ok {
					var tmp28 string
					if val, ok := v.(string); ok {
						tmp28 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.CreatedBy`, v, "string", err)
					}
					if err == nil {
						if tmp28 != "" {
							if err2 := goa.ValidateFormat(goa.FormatEmail, tmp28); err2 != nil {
								err = goa.InvalidFormatError(`.Account.CreatedBy`, tmp28, goa.FormatEmail, err2, err)
							}
						}
					}
					tmp26.CreatedBy = tmp28
				}
				if v, ok := val["href"]; ok {
					var tmp29 string
					if val, ok := v.(string); ok {
						tmp29 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.Href`, v, "string", err)
					}
					tmp26.Href = tmp29
				}
				if v, ok := val["id"]; ok {
					var tmp30 int
					if f, ok := v.(float64); ok {
						tmp30 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.ID`, v, "int", err)
					}
					tmp26.ID = tmp30
				}
				if v, ok := val["name"]; ok {
					var tmp31 string
					if val, ok := v.(string); ok {
						tmp31 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp31) < 2 {
							err = goa.InvalidLengthError(`.Account.Name`, tmp31, 2, true, err)
						}
					}
					tmp26.Name = tmp31
				}
			} else {
				err = goa.InvalidAttributeTypeError(`.Account`, v, "dictionary", err)
			}
			res.Account = tmp26
		}
		if v, ok := val["created_at"]; ok {
			var tmp32 string
			if val, ok := v.(string); ok {
				tmp32 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.CreatedAt`, v, "string", err)
			}
			if err == nil {
				if tmp32 != "" {
					if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp32); err2 != nil {
						err = goa.InvalidFormatError(`.CreatedAt`, tmp32, goa.FormatDateTime, err2, err)
					}
				}
			}
			res.CreatedAt = tmp32
		}
		if v, ok := val["href"]; ok {
			var tmp33 string
			if val, ok := v.(string); ok {
				tmp33 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp33
		}
		if v, ok := val["id"]; ok {
			var tmp34 int
			if f, ok := v.(float64); ok {
				tmp34 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp34
		}
		if v, ok := val["name"]; ok {
			var tmp35 string
			if val, ok := v.(string); ok {
				tmp35 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp35) < 2 {
					err = goa.InvalidLengthError(`.Name`, tmp35, 2, true, err)
				}
			}
			res.Name = tmp35
		}
		if v, ok := val["updated_at"]; ok {
			var tmp36 string
			if val, ok := v.(string); ok {
				tmp36 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.UpdatedAt`, v, "string", err)
			}
			if err == nil {
				if tmp36 != "" {
					if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp36); err2 != nil {
						err = goa.InvalidFormatError(`.UpdatedAt`, tmp36, goa.FormatDateTime, err2, err)
					}
				}
			}
			res.UpdatedAt = tmp36
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
		tmp38 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp38
		if err == nil {
			links := make(map[string]interface{})
			if len(mt.Account.Name) < 2 {
				err = goa.InvalidLengthError(`link account.name`, mt.Account.Name, 2, true, err)
			}
			tmp37 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			links["account"] = tmp37
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
		tmp40 := map[string]interface{}{
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
			if len(mt.Account.Name) < 2 {
				err = goa.InvalidLengthError(`full view.Account.name`, mt.Account.Name, 2, true, err)
			}
			tmp41 := map[string]interface{}{
				"created_at": mt.Account.CreatedAt,
				"created_by": mt.Account.CreatedBy,
				"href":       mt.Account.Href,
				"id":         mt.Account.ID,
				"name":       mt.Account.Name,
			}
			tmp40["account"] = tmp41
		}
		res = tmp40
		if err == nil {
			links := make(map[string]interface{})
			if len(mt.Account.Name) < 2 {
				err = goa.InvalidLengthError(`link account.name`, mt.Account.Name, 2, true, err)
			}
			tmp39 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			links["account"] = tmp39
			res["links"] = links
		}
	}
	if view == SeriesTinyView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`tiny view.name`, mt.Name, 2, true, err)
		}
		tmp43 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp43
		if err == nil {
			links := make(map[string]interface{})
			if len(mt.Account.Name) < 2 {
				err = goa.InvalidLengthError(`link account.name`, mt.Account.Name, 2, true, err)
			}
			tmp42 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			links["account"] = tmp42
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
	if len(mt.Account.Name) < 2 {
		err = goa.InvalidLengthError(`response.account.name`, mt.Account.Name, 2, true, err)
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
			var tmp44 *Series
			if val, ok := v.(map[string]interface{}); ok {
				tmp44 = new(Series)
				if v, ok := val["account"]; ok {
					var tmp45 *Account
					if val, ok := v.(map[string]interface{}); ok {
						tmp45 = new(Account)
						if v, ok := val["created_at"]; ok {
							var tmp46 string
							if val, ok := v.(string); ok {
								tmp46 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.CreatedAt`, v, "string", err)
							}
							if err == nil {
								if tmp46 != "" {
									if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp46); err2 != nil {
										err = goa.InvalidFormatError(`[*].Account.CreatedAt`, tmp46, goa.FormatDateTime, err2, err)
									}
								}
							}
							tmp45.CreatedAt = tmp46
						}
						if v, ok := val["created_by"]; ok {
							var tmp47 string
							if val, ok := v.(string); ok {
								tmp47 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.CreatedBy`, v, "string", err)
							}
							if err == nil {
								if tmp47 != "" {
									if err2 := goa.ValidateFormat(goa.FormatEmail, tmp47); err2 != nil {
										err = goa.InvalidFormatError(`[*].Account.CreatedBy`, tmp47, goa.FormatEmail, err2, err)
									}
								}
							}
							tmp45.CreatedBy = tmp47
						}
						if v, ok := val["href"]; ok {
							var tmp48 string
							if val, ok := v.(string); ok {
								tmp48 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.Href`, v, "string", err)
							}
							tmp45.Href = tmp48
						}
						if v, ok := val["id"]; ok {
							var tmp49 int
							if f, ok := v.(float64); ok {
								tmp49 = int(f)
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.ID`, v, "int", err)
							}
							tmp45.ID = tmp49
						}
						if v, ok := val["name"]; ok {
							var tmp50 string
							if val, ok := v.(string); ok {
								tmp50 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.Name`, v, "string", err)
							}
							if err == nil {
								if len(tmp50) < 2 {
									err = goa.InvalidLengthError(`[*].Account.Name`, tmp50, 2, true, err)
								}
							}
							tmp45.Name = tmp50
						}
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Account`, v, "dictionary", err)
					}
					tmp44.Account = tmp45
				}
				if v, ok := val["created_at"]; ok {
					var tmp51 string
					if val, ok := v.(string); ok {
						tmp51 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].CreatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp51 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp51); err2 != nil {
								err = goa.InvalidFormatError(`[*].CreatedAt`, tmp51, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp44.CreatedAt = tmp51
				}
				if v, ok := val["href"]; ok {
					var tmp52 string
					if val, ok := v.(string); ok {
						tmp52 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp44.Href = tmp52
				}
				if v, ok := val["id"]; ok {
					var tmp53 int
					if f, ok := v.(float64); ok {
						tmp53 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp44.ID = tmp53
				}
				if v, ok := val["name"]; ok {
					var tmp54 string
					if val, ok := v.(string); ok {
						tmp54 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp54) < 2 {
							err = goa.InvalidLengthError(`[*].Name`, tmp54, 2, true, err)
						}
					}
					tmp44.Name = tmp54
				}
				if v, ok := val["updated_at"]; ok {
					var tmp55 string
					if val, ok := v.(string); ok {
						tmp55 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].UpdatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp55 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp55); err2 != nil {
								err = goa.InvalidFormatError(`[*].UpdatedAt`, tmp55, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp44.UpdatedAt = tmp55
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "dictionary", err)
			}
			res[i] = tmp44
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
		for i, tmp56 := range mt {
			if len(tmp56.Name) < 2 {
				err = goa.InvalidLengthError(`default view[*].name`, tmp56.Name, 2, true, err)
			}
			tmp58 := map[string]interface{}{
				"href": tmp56.Href,
				"id":   tmp56.ID,
				"name": tmp56.Name,
			}
			res[i] = tmp58
			if err == nil {
				links := make(map[string]interface{})
				if len(tmp56.Account.Name) < 2 {
					err = goa.InvalidLengthError(`link account.name`, tmp56.Account.Name, 2, true, err)
				}
				tmp57 := map[string]interface{}{
					"href": tmp56.Account.Href,
					"id":   tmp56.Account.ID,
					"name": tmp56.Account.Name,
				}
				links["account"] = tmp57
				res[i]["links"] = links
			}
		}
	}
	if view == SeriesCollectionTinyView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp59 := range mt {
			if len(tmp59.Name) < 2 {
				err = goa.InvalidLengthError(`tiny view[*].name`, tmp59.Name, 2, true, err)
			}
			tmp61 := map[string]interface{}{
				"href": tmp59.Href,
				"id":   tmp59.ID,
				"name": tmp59.Name,
			}
			res[i] = tmp61
			if err == nil {
				links := make(map[string]interface{})
				if len(tmp59.Account.Name) < 2 {
					err = goa.InvalidLengthError(`link account.name`, tmp59.Account.Name, 2, true, err)
				}
				tmp60 := map[string]interface{}{
					"href": tmp59.Account.Href,
					"id":   tmp59.Account.ID,
					"name": tmp59.Account.Name,
				}
				links["account"] = tmp60
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
		if len(e.Account.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].account.name`, e.Account.Name, 2, true, err)
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
// Identifier: application/vnd.congo.api.user+json
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
		if v, ok := val["created_at"]; ok {
			var tmp62 string
			if val, ok := v.(string); ok {
				tmp62 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.CreatedAt`, v, "string", err)
			}
			if err == nil {
				if tmp62 != "" {
					if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp62); err2 != nil {
						err = goa.InvalidFormatError(`.CreatedAt`, tmp62, goa.FormatDateTime, err2, err)
					}
				}
			}
			res.CreatedAt = tmp62
		}
		if v, ok := val["email"]; ok {
			var tmp63 string
			if val, ok := v.(string); ok {
				tmp63 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Email`, v, "string", err)
			}
			if err == nil {
				if tmp63 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp63); err2 != nil {
						err = goa.InvalidFormatError(`.Email`, tmp63, goa.FormatEmail, err2, err)
					}
				}
			}
			res.Email = tmp63
		}
		if v, ok := val["first_name"]; ok {
			var tmp64 string
			if val, ok := v.(string); ok {
				tmp64 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.FirstName`, v, "string", err)
			}
			res.FirstName = tmp64
		}
		if v, ok := val["href"]; ok {
			var tmp65 string
			if val, ok := v.(string); ok {
				tmp65 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp65
		}
		if v, ok := val["id"]; ok {
			var tmp66 int
			if f, ok := v.(float64); ok {
				tmp66 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp66
		}
		if v, ok := val["last_name"]; ok {
			var tmp67 string
			if val, ok := v.(string); ok {
				tmp67 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.LastName`, v, "string", err)
			}
			res.LastName = tmp67
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
		tmp68 := map[string]interface{}{
			"created_at": mt.CreatedAt,
			"email":      mt.Email,
			"first_name": mt.FirstName,
			"href":       mt.Href,
			"id":         mt.ID,
			"last_name":  mt.LastName,
		}
		res = tmp68
	}
	if view == UserLinkView {
		if mt.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
				err = goa.InvalidFormatError(`link view.email`, mt.Email, goa.FormatEmail, err2, err)
			}
		}
		tmp69 := map[string]interface{}{
			"email": mt.Email,
			"href":  mt.Href,
			"id":    mt.ID,
		}
		res = tmp69
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
			var tmp70 *User
			if val, ok := v.(map[string]interface{}); ok {
				tmp70 = new(User)
				if v, ok := val["created_at"]; ok {
					var tmp71 string
					if val, ok := v.(string); ok {
						tmp71 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].CreatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp71 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp71); err2 != nil {
								err = goa.InvalidFormatError(`[*].CreatedAt`, tmp71, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp70.CreatedAt = tmp71
				}
				if v, ok := val["email"]; ok {
					var tmp72 string
					if val, ok := v.(string); ok {
						tmp72 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Email`, v, "string", err)
					}
					if err == nil {
						if tmp72 != "" {
							if err2 := goa.ValidateFormat(goa.FormatEmail, tmp72); err2 != nil {
								err = goa.InvalidFormatError(`[*].Email`, tmp72, goa.FormatEmail, err2, err)
							}
						}
					}
					tmp70.Email = tmp72
				}
				if v, ok := val["first_name"]; ok {
					var tmp73 string
					if val, ok := v.(string); ok {
						tmp73 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].FirstName`, v, "string", err)
					}
					tmp70.FirstName = tmp73
				}
				if v, ok := val["href"]; ok {
					var tmp74 string
					if val, ok := v.(string); ok {
						tmp74 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp70.Href = tmp74
				}
				if v, ok := val["id"]; ok {
					var tmp75 int
					if f, ok := v.(float64); ok {
						tmp75 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp70.ID = tmp75
				}
				if v, ok := val["last_name"]; ok {
					var tmp76 string
					if val, ok := v.(string); ok {
						tmp76 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].LastName`, v, "string", err)
					}
					tmp70.LastName = tmp76
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "dictionary", err)
			}
			res[i] = tmp70
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
	for i, tmp77 := range mt {
		if tmp77.CreatedAt != "" {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp77.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`default view[*].created_at`, tmp77.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if tmp77.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, tmp77.Email); err2 != nil {
				err = goa.InvalidFormatError(`default view[*].email`, tmp77.Email, goa.FormatEmail, err2, err)
			}
		}
		tmp78 := map[string]interface{}{
			"created_at": tmp77.CreatedAt,
			"email":      tmp77.Email,
			"first_name": tmp77.FirstName,
			"href":       tmp77.Href,
			"id":         tmp77.ID,
			"last_name":  tmp77.LastName,
		}
		res[i] = tmp78
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
