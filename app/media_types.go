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
			var tmp5 string
			if val, ok := v.(string); ok {
				tmp5 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.CreatedAt`, v, "string", err)
			}
			if err == nil {
				if tmp5 != "" {
					if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp5); err2 != nil {
						err = goa.InvalidFormatError(`.CreatedAt`, tmp5, goa.FormatDateTime, err2, err)
					}
				}
			}
			res.CreatedAt = tmp5
		}
		if v, ok := val["created_by"]; ok {
			var tmp6 string
			if val, ok := v.(string); ok {
				tmp6 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.CreatedBy`, v, "string", err)
			}
			if err == nil {
				if tmp6 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp6); err2 != nil {
						err = goa.InvalidFormatError(`.CreatedBy`, tmp6, goa.FormatEmail, err2, err)
					}
				}
			}
			res.CreatedBy = tmp6
		}
		if v, ok := val["href"]; ok {
			var tmp7 string
			if val, ok := v.(string); ok {
				tmp7 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp7
		}
		if v, ok := val["id"]; ok {
			var tmp8 int
			if f, ok := v.(float64); ok {
				tmp8 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp8
		}
		if v, ok := val["name"]; ok {
			var tmp9 string
			if val, ok := v.(string); ok {
				tmp9 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Name`, v, "string", err)
			}
			res.Name = tmp9
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
		tmp10 := map[string]interface{}{
			"created_at": mt.CreatedAt,
			"created_by": mt.CreatedBy,
			"href":       mt.Href,
			"id":         mt.ID,
			"name":       mt.Name,
		}
		res = tmp10
	}
	if view == AccountLinkView {
		tmp11 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp11
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
			var tmp12 *Account
			if val, ok := v.(map[string]interface{}); ok {
				tmp12 = new(Account)
				if v, ok := val["created_at"]; ok {
					var tmp13 string
					if val, ok := v.(string); ok {
						tmp13 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.CreatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp13 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp13); err2 != nil {
								err = goa.InvalidFormatError(`.Account.CreatedAt`, tmp13, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp12.CreatedAt = tmp13
				}
				if v, ok := val["created_by"]; ok {
					var tmp14 string
					if val, ok := v.(string); ok {
						tmp14 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.CreatedBy`, v, "string", err)
					}
					if err == nil {
						if tmp14 != "" {
							if err2 := goa.ValidateFormat(goa.FormatEmail, tmp14); err2 != nil {
								err = goa.InvalidFormatError(`.Account.CreatedBy`, tmp14, goa.FormatEmail, err2, err)
							}
						}
					}
					tmp12.CreatedBy = tmp14
				}
				if v, ok := val["href"]; ok {
					var tmp15 string
					if val, ok := v.(string); ok {
						tmp15 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.Href`, v, "string", err)
					}
					tmp12.Href = tmp15
				}
				if v, ok := val["id"]; ok {
					var tmp16 int
					if f, ok := v.(float64); ok {
						tmp16 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.ID`, v, "int", err)
					}
					tmp12.ID = tmp16
				}
				if v, ok := val["name"]; ok {
					var tmp17 string
					if val, ok := v.(string); ok {
						tmp17 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.Name`, v, "string", err)
					}
					tmp12.Name = tmp17
				}
			} else {
				err = goa.InvalidAttributeTypeError(`.Account`, v, "map[string]interface{}", err)
			}
			res.Account = tmp12
		}
		if v, ok := val["created_at"]; ok {
			var tmp18 string
			if val, ok := v.(string); ok {
				tmp18 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.CreatedAt`, v, "string", err)
			}
			if err == nil {
				if tmp18 != "" {
					if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp18); err2 != nil {
						err = goa.InvalidFormatError(`.CreatedAt`, tmp18, goa.FormatDateTime, err2, err)
					}
				}
			}
			res.CreatedAt = tmp18
		}
		if v, ok := val["href"]; ok {
			var tmp19 string
			if val, ok := v.(string); ok {
				tmp19 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp19
		}
		if v, ok := val["id"]; ok {
			var tmp20 int
			if f, ok := v.(float64); ok {
				tmp20 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp20
		}
		if v, ok := val["name"]; ok {
			var tmp21 string
			if val, ok := v.(string); ok {
				tmp21 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp21) < 2 {
					err = goa.InvalidLengthError(`.Name`, tmp21, 2, true, err)
				}
			}
			res.Name = tmp21
		}
		if v, ok := val["updated_at"]; ok {
			var tmp22 string
			if val, ok := v.(string); ok {
				tmp22 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.UpdatedAt`, v, "string", err)
			}
			if err == nil {
				if tmp22 != "" {
					if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp22); err2 != nil {
						err = goa.InvalidFormatError(`.UpdatedAt`, tmp22, goa.FormatDateTime, err2, err)
					}
				}
			}
			res.UpdatedAt = tmp22
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
		tmp24 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp24
		if err == nil {
			links := make(map[string]interface{})
			tmp23 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			links["account"] = tmp23
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
		tmp26 := map[string]interface{}{
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
			tmp27 := map[string]interface{}{
				"created_at": mt.Account.CreatedAt,
				"created_by": mt.Account.CreatedBy,
				"href":       mt.Account.Href,
				"id":         mt.Account.ID,
				"name":       mt.Account.Name,
			}
			tmp26["account"] = tmp27
		}
		res = tmp26
		if err == nil {
			links := make(map[string]interface{})
			tmp25 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			links["account"] = tmp25
			res["links"] = links
		}
	}
	if view == SeriesTinyView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`tiny view.name`, mt.Name, 2, true, err)
		}
		tmp29 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp29
		if err == nil {
			links := make(map[string]interface{})
			tmp28 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			links["account"] = tmp28
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
			var tmp30 *Series
			if val, ok := v.(map[string]interface{}); ok {
				tmp30 = new(Series)
				if v, ok := val["account"]; ok {
					var tmp31 *Account
					if val, ok := v.(map[string]interface{}); ok {
						tmp31 = new(Account)
						if v, ok := val["created_at"]; ok {
							var tmp32 string
							if val, ok := v.(string); ok {
								tmp32 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.CreatedAt`, v, "string", err)
							}
							if err == nil {
								if tmp32 != "" {
									if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp32); err2 != nil {
										err = goa.InvalidFormatError(`[*].Account.CreatedAt`, tmp32, goa.FormatDateTime, err2, err)
									}
								}
							}
							tmp31.CreatedAt = tmp32
						}
						if v, ok := val["created_by"]; ok {
							var tmp33 string
							if val, ok := v.(string); ok {
								tmp33 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.CreatedBy`, v, "string", err)
							}
							if err == nil {
								if tmp33 != "" {
									if err2 := goa.ValidateFormat(goa.FormatEmail, tmp33); err2 != nil {
										err = goa.InvalidFormatError(`[*].Account.CreatedBy`, tmp33, goa.FormatEmail, err2, err)
									}
								}
							}
							tmp31.CreatedBy = tmp33
						}
						if v, ok := val["href"]; ok {
							var tmp34 string
							if val, ok := v.(string); ok {
								tmp34 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.Href`, v, "string", err)
							}
							tmp31.Href = tmp34
						}
						if v, ok := val["id"]; ok {
							var tmp35 int
							if f, ok := v.(float64); ok {
								tmp35 = int(f)
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.ID`, v, "int", err)
							}
							tmp31.ID = tmp35
						}
						if v, ok := val["name"]; ok {
							var tmp36 string
							if val, ok := v.(string); ok {
								tmp36 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.Name`, v, "string", err)
							}
							tmp31.Name = tmp36
						}
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Account`, v, "map[string]interface{}", err)
					}
					tmp30.Account = tmp31
				}
				if v, ok := val["created_at"]; ok {
					var tmp37 string
					if val, ok := v.(string); ok {
						tmp37 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].CreatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp37 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp37); err2 != nil {
								err = goa.InvalidFormatError(`[*].CreatedAt`, tmp37, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp30.CreatedAt = tmp37
				}
				if v, ok := val["href"]; ok {
					var tmp38 string
					if val, ok := v.(string); ok {
						tmp38 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp30.Href = tmp38
				}
				if v, ok := val["id"]; ok {
					var tmp39 int
					if f, ok := v.(float64); ok {
						tmp39 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp30.ID = tmp39
				}
				if v, ok := val["name"]; ok {
					var tmp40 string
					if val, ok := v.(string); ok {
						tmp40 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp40) < 2 {
							err = goa.InvalidLengthError(`[*].Name`, tmp40, 2, true, err)
						}
					}
					tmp30.Name = tmp40
				}
				if v, ok := val["updated_at"]; ok {
					var tmp41 string
					if val, ok := v.(string); ok {
						tmp41 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].UpdatedAt`, v, "string", err)
					}
					if err == nil {
						if tmp41 != "" {
							if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp41); err2 != nil {
								err = goa.InvalidFormatError(`[*].UpdatedAt`, tmp41, goa.FormatDateTime, err2, err)
							}
						}
					}
					tmp30.UpdatedAt = tmp41
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "map[string]interface{}", err)
			}
			res[i] = tmp30
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
		for i, tmp42 := range mt {
			if len(tmp42.Name) < 2 {
				err = goa.InvalidLengthError(`default view[*].name`, tmp42.Name, 2, true, err)
			}
			tmp44 := map[string]interface{}{
				"href": tmp42.Href,
				"id":   tmp42.ID,
				"name": tmp42.Name,
			}
			res[i] = tmp44
			if err == nil {
				links := make(map[string]interface{})
				tmp43 := map[string]interface{}{
					"href": tmp42.Account.Href,
					"id":   tmp42.Account.ID,
					"name": tmp42.Account.Name,
				}
				links["account"] = tmp43
				res[i]["links"] = links
			}
		}
	}
	if view == SeriesCollectionTinyView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp45 := range mt {
			if len(tmp45.Name) < 2 {
				err = goa.InvalidLengthError(`tiny view[*].name`, tmp45.Name, 2, true, err)
			}
			tmp47 := map[string]interface{}{
				"href": tmp45.Href,
				"id":   tmp45.ID,
				"name": tmp45.Name,
			}
			res[i] = tmp47
			if err == nil {
				links := make(map[string]interface{})
				tmp46 := map[string]interface{}{
					"href": tmp45.Account.Href,
					"id":   tmp45.Account.ID,
					"name": tmp45.Account.Name,
				}
				links["account"] = tmp46
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
