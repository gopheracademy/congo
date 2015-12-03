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
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`default view.name`, mt.Name, 2, true, err)
		}
		tmp16 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
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
				if v, ok := val["href"]; ok {
					var tmp19 string
					if val, ok := v.(string); ok {
						tmp19 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp18.Href = tmp19
				}
				if v, ok := val["id"]; ok {
					var tmp20 int
					if f, ok := v.(float64); ok {
						tmp20 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp18.ID = tmp20
				}
				if v, ok := val["name"]; ok {
					var tmp21 string
					if val, ok := v.(string); ok {
						tmp21 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp21) < 2 {
							err = goa.InvalidLengthError(`[*].Name`, tmp21, 2, true, err)
						}
					}
					tmp18.Name = tmp21
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
	for i, tmp22 := range mt {
		if len(tmp22.Name) < 2 {
			err = goa.InvalidLengthError(`default view[*].name`, tmp22.Name, 2, true, err)
		}
		tmp23 := map[string]interface{}{
			"href": tmp22.Href,
			"id":   tmp22.ID,
			"name": tmp22.Name,
		}
		res[i] = tmp23
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

// An instance of an event or conference
// Identifier: application/vnd.congo.api.instance+json
type Instance struct {
	// API href of instance
	Href string
	// ID of Instance
	ID   int
	Name string
	// Series that this instance belongs to
	Series *Series
}

// Instance views
type InstanceViewEnum string

const (
	// Instance default view
	InstanceDefaultView InstanceViewEnum = "default"
	// Instance full view
	InstanceFullView InstanceViewEnum = "full"
	// Instance tiny view
	InstanceTinyView InstanceViewEnum = "tiny"
)

// LoadInstance loads raw data into an instance of Instance running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadInstance(raw interface{}) (*Instance, error) {
	var err error
	var res *Instance
	if val, ok := raw.(map[string]interface{}); ok {
		res = new(Instance)
		if v, ok := val["href"]; ok {
			var tmp24 string
			if val, ok := v.(string); ok {
				tmp24 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp24
		}
		if v, ok := val["id"]; ok {
			var tmp25 int
			if f, ok := v.(float64); ok {
				tmp25 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp25
		}
		if v, ok := val["name"]; ok {
			var tmp26 string
			if val, ok := v.(string); ok {
				tmp26 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp26) < 2 {
					err = goa.InvalidLengthError(`.Name`, tmp26, 2, true, err)
				}
			}
			res.Name = tmp26
		}
		if v, ok := val["series"]; ok {
			var tmp27 *Series
			if val, ok := v.(map[string]interface{}); ok {
				tmp27 = new(Series)
				if v, ok := val["account"]; ok {
					var tmp28 *Account
					if val, ok := v.(map[string]interface{}); ok {
						tmp28 = new(Account)
						if v, ok := val["href"]; ok {
							var tmp29 string
							if val, ok := v.(string); ok {
								tmp29 = val
							} else {
								err = goa.InvalidAttributeTypeError(`.Series.Account.Href`, v, "string", err)
							}
							tmp28.Href = tmp29
						}
						if v, ok := val["id"]; ok {
							var tmp30 int
							if f, ok := v.(float64); ok {
								tmp30 = int(f)
							} else {
								err = goa.InvalidAttributeTypeError(`.Series.Account.ID`, v, "int", err)
							}
							tmp28.ID = tmp30
						}
						if v, ok := val["name"]; ok {
							var tmp31 string
							if val, ok := v.(string); ok {
								tmp31 = val
							} else {
								err = goa.InvalidAttributeTypeError(`.Series.Account.Name`, v, "string", err)
							}
							if err == nil {
								if len(tmp31) < 2 {
									err = goa.InvalidLengthError(`.Series.Account.Name`, tmp31, 2, true, err)
								}
							}
							tmp28.Name = tmp31
						}
					} else {
						err = goa.InvalidAttributeTypeError(`.Series.Account`, v, "dictionary", err)
					}
					tmp27.Account = tmp28
				}
				if v, ok := val["href"]; ok {
					var tmp32 string
					if val, ok := v.(string); ok {
						tmp32 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Series.Href`, v, "string", err)
					}
					tmp27.Href = tmp32
				}
				if v, ok := val["id"]; ok {
					var tmp33 int
					if f, ok := v.(float64); ok {
						tmp33 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`.Series.ID`, v, "int", err)
					}
					tmp27.ID = tmp33
				}
				if v, ok := val["name"]; ok {
					var tmp34 string
					if val, ok := v.(string); ok {
						tmp34 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Series.Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp34) < 2 {
							err = goa.InvalidLengthError(`.Series.Name`, tmp34, 2, true, err)
						}
					}
					tmp27.Name = tmp34
				}
			} else {
				err = goa.InvalidAttributeTypeError(`.Series`, v, "dictionary", err)
			}
			res.Series = tmp27
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "dictionary", err)
	}
	return res, err
}

// Dump produces raw data from an instance of Instance running all the
// validations. See LoadInstance for the definition of raw data.
func (mt *Instance) Dump(view InstanceViewEnum) (map[string]interface{}, error) {
	var err error
	var res map[string]interface{}
	if view == InstanceDefaultView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`default view.name`, mt.Name, 2, true, err)
		}
		tmp35 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp35
	}
	if view == InstanceFullView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`full view.name`, mt.Name, 2, true, err)
		}
		tmp36 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		if mt.Series != nil {
			if len(mt.Series.Name) < 2 {
				err = goa.InvalidLengthError(`full view.Series.name`, mt.Series.Name, 2, true, err)
			}
			tmp37 := map[string]interface{}{
				"href": mt.Series.Href,
				"id":   mt.Series.ID,
				"name": mt.Series.Name,
			}
			tmp36["series"] = tmp37
		}
		res = tmp36
	}
	if view == InstanceTinyView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`tiny view.name`, mt.Name, 2, true, err)
		}
		tmp38 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp38
	}
	return res, err
}

// Validate validates the media type instance.
func (mt *Instance) Validate() (err error) {
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, 2, true, err)
	}
	if len(mt.Series.Account.Name) < 2 {
		err = goa.InvalidLengthError(`response.series.account.name`, mt.Series.Account.Name, 2, true, err)
	}
	if len(mt.Series.Name) < 2 {
		err = goa.InvalidLengthError(`response.series.name`, mt.Series.Name, 2, true, err)
	}
	return
}

// InstanceCollection media type
// Identifier: application/vnd.congo.api.instance+json; type=collection
type InstanceCollection []*Instance

// InstanceCollection views
type InstanceCollectionViewEnum string

const (
	// InstanceCollection default view
	InstanceCollectionDefaultView InstanceCollectionViewEnum = "default"
	// InstanceCollection tiny view
	InstanceCollectionTinyView InstanceCollectionViewEnum = "tiny"
)

// LoadInstanceCollection loads raw data into an instance of InstanceCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadInstanceCollection(raw interface{}) (InstanceCollection, error) {
	var err error
	var res InstanceCollection
	if val, ok := raw.([]interface{}); ok {
		res = make([]*Instance, len(val))
		for i, v := range val {
			var tmp39 *Instance
			if val, ok := v.(map[string]interface{}); ok {
				tmp39 = new(Instance)
				if v, ok := val["href"]; ok {
					var tmp40 string
					if val, ok := v.(string); ok {
						tmp40 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp39.Href = tmp40
				}
				if v, ok := val["id"]; ok {
					var tmp41 int
					if f, ok := v.(float64); ok {
						tmp41 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp39.ID = tmp41
				}
				if v, ok := val["name"]; ok {
					var tmp42 string
					if val, ok := v.(string); ok {
						tmp42 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp42) < 2 {
							err = goa.InvalidLengthError(`[*].Name`, tmp42, 2, true, err)
						}
					}
					tmp39.Name = tmp42
				}
				if v, ok := val["series"]; ok {
					var tmp43 *Series
					if val, ok := v.(map[string]interface{}); ok {
						tmp43 = new(Series)
						if v, ok := val["account"]; ok {
							var tmp44 *Account
							if val, ok := v.(map[string]interface{}); ok {
								tmp44 = new(Account)
								if v, ok := val["href"]; ok {
									var tmp45 string
									if val, ok := v.(string); ok {
										tmp45 = val
									} else {
										err = goa.InvalidAttributeTypeError(`[*].Series.Account.Href`, v, "string", err)
									}
									tmp44.Href = tmp45
								}
								if v, ok := val["id"]; ok {
									var tmp46 int
									if f, ok := v.(float64); ok {
										tmp46 = int(f)
									} else {
										err = goa.InvalidAttributeTypeError(`[*].Series.Account.ID`, v, "int", err)
									}
									tmp44.ID = tmp46
								}
								if v, ok := val["name"]; ok {
									var tmp47 string
									if val, ok := v.(string); ok {
										tmp47 = val
									} else {
										err = goa.InvalidAttributeTypeError(`[*].Series.Account.Name`, v, "string", err)
									}
									if err == nil {
										if len(tmp47) < 2 {
											err = goa.InvalidLengthError(`[*].Series.Account.Name`, tmp47, 2, true, err)
										}
									}
									tmp44.Name = tmp47
								}
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Series.Account`, v, "dictionary", err)
							}
							tmp43.Account = tmp44
						}
						if v, ok := val["href"]; ok {
							var tmp48 string
							if val, ok := v.(string); ok {
								tmp48 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Series.Href`, v, "string", err)
							}
							tmp43.Href = tmp48
						}
						if v, ok := val["id"]; ok {
							var tmp49 int
							if f, ok := v.(float64); ok {
								tmp49 = int(f)
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Series.ID`, v, "int", err)
							}
							tmp43.ID = tmp49
						}
						if v, ok := val["name"]; ok {
							var tmp50 string
							if val, ok := v.(string); ok {
								tmp50 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Series.Name`, v, "string", err)
							}
							if err == nil {
								if len(tmp50) < 2 {
									err = goa.InvalidLengthError(`[*].Series.Name`, tmp50, 2, true, err)
								}
							}
							tmp43.Name = tmp50
						}
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Series`, v, "dictionary", err)
					}
					tmp39.Series = tmp43
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "dictionary", err)
			}
			res[i] = tmp39
		}
	} else {
		err = goa.InvalidAttributeTypeError(``, raw, "array", err)
	}
	return res, err
}

// Dump produces raw data from an instance of InstanceCollection running all the
// validations. See LoadInstanceCollection for the definition of raw data.
func (mt InstanceCollection) Dump(view InstanceCollectionViewEnum) ([]map[string]interface{}, error) {
	var err error
	var res []map[string]interface{}
	if view == InstanceCollectionDefaultView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp51 := range mt {
			if len(tmp51.Name) < 2 {
				err = goa.InvalidLengthError(`default view[*].name`, tmp51.Name, 2, true, err)
			}
			tmp52 := map[string]interface{}{
				"href": tmp51.Href,
				"id":   tmp51.ID,
				"name": tmp51.Name,
			}
			res[i] = tmp52
		}
	}
	if view == InstanceCollectionTinyView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp53 := range mt {
			if len(tmp53.Name) < 2 {
				err = goa.InvalidLengthError(`tiny view[*].name`, tmp53.Name, 2, true, err)
			}
			tmp54 := map[string]interface{}{
				"href": tmp53.Href,
				"id":   tmp53.ID,
				"name": tmp53.Name,
			}
			res[i] = tmp54
		}
	}
	return res, err
}

// Validate validates the media type instance.
func (mt InstanceCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, 2, true, err)
		}
		if len(e.Series.Account.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].series.account.name`, e.Series.Account.Name, 2, true, err)
		}
		if len(e.Series.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].series.name`, e.Series.Name, 2, true, err)
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
	// Series link view
	SeriesLinkView SeriesViewEnum = "link"
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
			var tmp55 *Account
			if val, ok := v.(map[string]interface{}); ok {
				tmp55 = new(Account)
				if v, ok := val["href"]; ok {
					var tmp56 string
					if val, ok := v.(string); ok {
						tmp56 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.Href`, v, "string", err)
					}
					tmp55.Href = tmp56
				}
				if v, ok := val["id"]; ok {
					var tmp57 int
					if f, ok := v.(float64); ok {
						tmp57 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.ID`, v, "int", err)
					}
					tmp55.ID = tmp57
				}
				if v, ok := val["name"]; ok {
					var tmp58 string
					if val, ok := v.(string); ok {
						tmp58 = val
					} else {
						err = goa.InvalidAttributeTypeError(`.Account.Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp58) < 2 {
							err = goa.InvalidLengthError(`.Account.Name`, tmp58, 2, true, err)
						}
					}
					tmp55.Name = tmp58
				}
			} else {
				err = goa.InvalidAttributeTypeError(`.Account`, v, "dictionary", err)
			}
			res.Account = tmp55
		}
		if v, ok := val["href"]; ok {
			var tmp59 string
			if val, ok := v.(string); ok {
				tmp59 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp59
		}
		if v, ok := val["id"]; ok {
			var tmp60 int
			if f, ok := v.(float64); ok {
				tmp60 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp60
		}
		if v, ok := val["name"]; ok {
			var tmp61 string
			if val, ok := v.(string); ok {
				tmp61 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp61) < 2 {
					err = goa.InvalidLengthError(`.Name`, tmp61, 2, true, err)
				}
			}
			res.Name = tmp61
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
		tmp62 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp62
	}
	if view == SeriesFullView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`full view.name`, mt.Name, 2, true, err)
		}
		tmp63 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		if mt.Account != nil {
			if len(mt.Account.Name) < 2 {
				err = goa.InvalidLengthError(`full view.Account.name`, mt.Account.Name, 2, true, err)
			}
			tmp64 := map[string]interface{}{
				"href": mt.Account.Href,
				"id":   mt.Account.ID,
				"name": mt.Account.Name,
			}
			tmp63["account"] = tmp64
		}
		res = tmp63
	}
	if view == SeriesLinkView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`link view.name`, mt.Name, 2, true, err)
		}
		tmp65 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp65
	}
	if view == SeriesTinyView {
		if len(mt.Name) < 2 {
			err = goa.InvalidLengthError(`tiny view.name`, mt.Name, 2, true, err)
		}
		tmp66 := map[string]interface{}{
			"href": mt.Href,
			"id":   mt.ID,
			"name": mt.Name,
		}
		res = tmp66
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
			var tmp67 *Series
			if val, ok := v.(map[string]interface{}); ok {
				tmp67 = new(Series)
				if v, ok := val["account"]; ok {
					var tmp68 *Account
					if val, ok := v.(map[string]interface{}); ok {
						tmp68 = new(Account)
						if v, ok := val["href"]; ok {
							var tmp69 string
							if val, ok := v.(string); ok {
								tmp69 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.Href`, v, "string", err)
							}
							tmp68.Href = tmp69
						}
						if v, ok := val["id"]; ok {
							var tmp70 int
							if f, ok := v.(float64); ok {
								tmp70 = int(f)
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.ID`, v, "int", err)
							}
							tmp68.ID = tmp70
						}
						if v, ok := val["name"]; ok {
							var tmp71 string
							if val, ok := v.(string); ok {
								tmp71 = val
							} else {
								err = goa.InvalidAttributeTypeError(`[*].Account.Name`, v, "string", err)
							}
							if err == nil {
								if len(tmp71) < 2 {
									err = goa.InvalidLengthError(`[*].Account.Name`, tmp71, 2, true, err)
								}
							}
							tmp68.Name = tmp71
						}
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Account`, v, "dictionary", err)
					}
					tmp67.Account = tmp68
				}
				if v, ok := val["href"]; ok {
					var tmp72 string
					if val, ok := v.(string); ok {
						tmp72 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp67.Href = tmp72
				}
				if v, ok := val["id"]; ok {
					var tmp73 int
					if f, ok := v.(float64); ok {
						tmp73 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp67.ID = tmp73
				}
				if v, ok := val["name"]; ok {
					var tmp74 string
					if val, ok := v.(string); ok {
						tmp74 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Name`, v, "string", err)
					}
					if err == nil {
						if len(tmp74) < 2 {
							err = goa.InvalidLengthError(`[*].Name`, tmp74, 2, true, err)
						}
					}
					tmp67.Name = tmp74
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "dictionary", err)
			}
			res[i] = tmp67
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
		for i, tmp75 := range mt {
			if len(tmp75.Name) < 2 {
				err = goa.InvalidLengthError(`default view[*].name`, tmp75.Name, 2, true, err)
			}
			tmp76 := map[string]interface{}{
				"href": tmp75.Href,
				"id":   tmp75.ID,
				"name": tmp75.Name,
			}
			res[i] = tmp76
		}
	}
	if view == SeriesCollectionTinyView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp77 := range mt {
			if len(tmp77.Name) < 2 {
				err = goa.InvalidLengthError(`tiny view[*].name`, tmp77.Name, 2, true, err)
			}
			tmp78 := map[string]interface{}{
				"href": tmp77.Href,
				"id":   tmp77.ID,
				"name": tmp77.Name,
			}
			res[i] = tmp78
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
			var tmp79 string
			if val, ok := v.(string); ok {
				tmp79 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Email`, v, "string", err)
			}
			if err == nil {
				if tmp79 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp79); err2 != nil {
						err = goa.InvalidFormatError(`.Email`, tmp79, goa.FormatEmail, err2, err)
					}
				}
			}
			res.Email = tmp79
		}
		if v, ok := val["first_name"]; ok {
			var tmp80 string
			if val, ok := v.(string); ok {
				tmp80 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.FirstName`, v, "string", err)
			}
			res.FirstName = tmp80
		}
		if v, ok := val["href"]; ok {
			var tmp81 string
			if val, ok := v.(string); ok {
				tmp81 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.Href`, v, "string", err)
			}
			res.Href = tmp81
		}
		if v, ok := val["id"]; ok {
			var tmp82 int
			if f, ok := v.(float64); ok {
				tmp82 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`.ID`, v, "int", err)
			}
			res.ID = tmp82
		}
		if v, ok := val["last_name"]; ok {
			var tmp83 string
			if val, ok := v.(string); ok {
				tmp83 = val
			} else {
				err = goa.InvalidAttributeTypeError(`.LastName`, v, "string", err)
			}
			res.LastName = tmp83
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
		tmp84 := map[string]interface{}{
			"email":      mt.Email,
			"first_name": mt.FirstName,
			"href":       mt.Href,
			"id":         mt.ID,
			"last_name":  mt.LastName,
		}
		res = tmp84
	}
	if view == UserLinkView {
		if mt.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
				err = goa.InvalidFormatError(`link view.email`, mt.Email, goa.FormatEmail, err2, err)
			}
		}
		tmp85 := map[string]interface{}{
			"email": mt.Email,
			"href":  mt.Href,
			"id":    mt.ID,
		}
		res = tmp85
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
			var tmp86 *User
			if val, ok := v.(map[string]interface{}); ok {
				tmp86 = new(User)
				if v, ok := val["email"]; ok {
					var tmp87 string
					if val, ok := v.(string); ok {
						tmp87 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Email`, v, "string", err)
					}
					if err == nil {
						if tmp87 != "" {
							if err2 := goa.ValidateFormat(goa.FormatEmail, tmp87); err2 != nil {
								err = goa.InvalidFormatError(`[*].Email`, tmp87, goa.FormatEmail, err2, err)
							}
						}
					}
					tmp86.Email = tmp87
				}
				if v, ok := val["first_name"]; ok {
					var tmp88 string
					if val, ok := v.(string); ok {
						tmp88 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].FirstName`, v, "string", err)
					}
					tmp86.FirstName = tmp88
				}
				if v, ok := val["href"]; ok {
					var tmp89 string
					if val, ok := v.(string); ok {
						tmp89 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].Href`, v, "string", err)
					}
					tmp86.Href = tmp89
				}
				if v, ok := val["id"]; ok {
					var tmp90 int
					if f, ok := v.(float64); ok {
						tmp90 = int(f)
					} else {
						err = goa.InvalidAttributeTypeError(`[*].ID`, v, "int", err)
					}
					tmp86.ID = tmp90
				}
				if v, ok := val["last_name"]; ok {
					var tmp91 string
					if val, ok := v.(string); ok {
						tmp91 = val
					} else {
						err = goa.InvalidAttributeTypeError(`[*].LastName`, v, "string", err)
					}
					tmp86.LastName = tmp91
				}
			} else {
				err = goa.InvalidAttributeTypeError(`[*]`, v, "dictionary", err)
			}
			res[i] = tmp86
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
	for i, tmp92 := range mt {
		if tmp92.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, tmp92.Email); err2 != nil {
				err = goa.InvalidFormatError(`default view[*].email`, tmp92.Email, goa.FormatEmail, err2, err)
			}
		}
		tmp93 := map[string]interface{}{
			"email":      tmp92.Email,
			"first_name": tmp92.FirstName,
			"href":       tmp92.Href,
			"id":         tmp92.ID,
			"last_name":  tmp92.LastName,
		}
		res[i] = tmp93
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
