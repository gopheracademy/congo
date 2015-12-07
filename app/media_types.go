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
// Identifier: application/vnd.account+json
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
func LoadAccount(raw interface{}) (res *Account, err error) {
	res, err = UnmarshalAccount(raw, err)
	return
}

// Dump produces raw data from an instance of Account running all the
// validations. See LoadAccount for the definition of raw data.
func (mt *Account) Dump(view AccountViewEnum) (res map[string]interface{}, err error) {
	if view == AccountDefaultView {
		res, err = MarshalAccount(mt, err)
	}
	if view == AccountLinkView {
		res, err = MarshalAccountLink(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *Account) Validate() (err error) {
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, 2, true, err)
	}
	return
}

// MarshalAccount validates and renders an instance of Account into a interface{}
// using view "default".
func MarshalAccount(source *Account, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp17 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	target = tmp17
	return
}

// MarshalAccountLink validates and renders an instance of Account into a interface{}
// using view "link".
func MarshalAccountLink(source *Account, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp18 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	target = tmp18
	return
}

// UnmarshalAccount unmarshals and validates a raw interface{} into an instance of Account
func UnmarshalAccount(source interface{}, inErr error) (target *Account, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Account)
		if v, ok := val["href"]; ok {
			var tmp19 string
			if val, ok := v.(string); ok {
				tmp19 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp19
		}
		if v, ok := val["id"]; ok {
			var tmp20 int
			if f, ok := v.(float64); ok {
				tmp20 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp20
		}
		if v, ok := val["name"]; ok {
			var tmp21 string
			if val, ok := v.(string); ok {
				tmp21 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp21) < 2 {
					err = goa.InvalidLengthError(`load.Name`, tmp21, 2, true, err)
				}
			}
			target.Name = tmp21
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// AccountCollection media type
// Identifier: application/vnd.account+json; type=collection
type AccountCollection []*Account

// LoadAccountCollection loads raw data into an instance of AccountCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadAccountCollection(raw interface{}) (res AccountCollection, err error) {
	res, err = UnmarshalAccountCollection(raw, err)
	return
}

// Dump produces raw data from an instance of AccountCollection running all the
// validations. See LoadAccountCollection for the definition of raw data.
func (mt AccountCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp22 := range mt {
		var tmp23 map[string]interface{}
		tmp23, err = MarshalAccount(tmp22, err)
		res[i] = tmp23
	}
	return
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

// MarshalAccountCollection validates and renders an instance of AccountCollection into a interface{}
// using view "default".
func MarshalAccountCollection(source AccountCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalAccount(res, err)
	}
	return
}

// UnmarshalAccountCollection unmarshals and validates a raw interface{} into an instance of AccountCollection
func UnmarshalAccountCollection(source interface{}, inErr error) (target AccountCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Account, len(val))
		for i, v := range val {
			target[i], err = UnmarshalAccount(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}

// An instance of an event or conference
// Identifier: application/vnd.instance+json
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
func LoadInstance(raw interface{}) (res *Instance, err error) {
	res, err = UnmarshalInstance(raw, err)
	return
}

// Dump produces raw data from an instance of Instance running all the
// validations. See LoadInstance for the definition of raw data.
func (mt *Instance) Dump(view InstanceViewEnum) (res map[string]interface{}, err error) {
	if view == InstanceDefaultView {
		res, err = MarshalInstance(mt, err)
	}
	if view == InstanceFullView {
		res, err = MarshalInstanceFull(mt, err)
	}
	if view == InstanceTinyView {
		res, err = MarshalInstanceTiny(mt, err)
	}
	return
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

// MarshalInstance validates and renders an instance of Instance into a interface{}
// using view "default".
func MarshalInstance(source *Instance, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp24 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	target = tmp24
	return
}

// MarshalInstanceFull validates and renders an instance of Instance into a interface{}
// using view "full".
func MarshalInstanceFull(source *Instance, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp25 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	if source.Series != nil {
		tmp25["series"], err = MarshalSeries(source.Series, err)
	}
	target = tmp25
	return
}

// MarshalInstanceTiny validates and renders an instance of Instance into a interface{}
// using view "tiny".
func MarshalInstanceTiny(source *Instance, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp26 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	target = tmp26
	return
}

// UnmarshalInstance unmarshals and validates a raw interface{} into an instance of Instance
func UnmarshalInstance(source interface{}, inErr error) (target *Instance, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Instance)
		if v, ok := val["href"]; ok {
			var tmp27 string
			if val, ok := v.(string); ok {
				tmp27 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp27
		}
		if v, ok := val["id"]; ok {
			var tmp28 int
			if f, ok := v.(float64); ok {
				tmp28 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp28
		}
		if v, ok := val["name"]; ok {
			var tmp29 string
			if val, ok := v.(string); ok {
				tmp29 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp29) < 2 {
					err = goa.InvalidLengthError(`load.Name`, tmp29, 2, true, err)
				}
			}
			target.Name = tmp29
		}
		if v, ok := val["series"]; ok {
			var tmp30 *Series
			tmp30, err = UnmarshalSeries(v, err)
			target.Series = tmp30
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// InstanceCollection media type
// Identifier: application/vnd.instance+json; type=collection
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
func LoadInstanceCollection(raw interface{}) (res InstanceCollection, err error) {
	res, err = UnmarshalInstanceCollection(raw, err)
	return
}

// Dump produces raw data from an instance of InstanceCollection running all the
// validations. See LoadInstanceCollection for the definition of raw data.
func (mt InstanceCollection) Dump(view InstanceCollectionViewEnum) (res []map[string]interface{}, err error) {
	if view == InstanceCollectionDefaultView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp31 := range mt {
			var tmp32 map[string]interface{}
			tmp32, err = MarshalInstance(tmp31, err)
			res[i] = tmp32
		}
	}
	if view == InstanceCollectionTinyView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp33 := range mt {
			var tmp34 map[string]interface{}
			tmp34, err = MarshalInstanceTiny(tmp33, err)
			res[i] = tmp34
		}
	}
	return
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

// MarshalInstanceCollection validates and renders an instance of InstanceCollection into a interface{}
// using view "default".
func MarshalInstanceCollection(source InstanceCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalInstance(res, err)
	}
	return
}

// MarshalInstanceCollectionTiny validates and renders an instance of InstanceCollection into a interface{}
// using view "tiny".
func MarshalInstanceCollectionTiny(source InstanceCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalInstanceTiny(res, err)
	}
	return
}

// UnmarshalInstanceCollection unmarshals and validates a raw interface{} into an instance of InstanceCollection
func UnmarshalInstanceCollection(source interface{}, inErr error) (target InstanceCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Instance, len(val))
		for i, v := range val {
			target[i], err = UnmarshalInstance(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}

// A recurring event or conference
// Identifier: application/vnd.series+json
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
func LoadSeries(raw interface{}) (res *Series, err error) {
	res, err = UnmarshalSeries(raw, err)
	return
}

// Dump produces raw data from an instance of Series running all the
// validations. See LoadSeries for the definition of raw data.
func (mt *Series) Dump(view SeriesViewEnum) (res map[string]interface{}, err error) {
	if view == SeriesDefaultView {
		res, err = MarshalSeries(mt, err)
	}
	if view == SeriesFullView {
		res, err = MarshalSeriesFull(mt, err)
	}
	if view == SeriesLinkView {
		res, err = MarshalSeriesLink(mt, err)
	}
	if view == SeriesTinyView {
		res, err = MarshalSeriesTiny(mt, err)
	}
	return
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

// MarshalSeries validates and renders an instance of Series into a interface{}
// using view "default".
func MarshalSeries(source *Series, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp35 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	target = tmp35
	return
}

// MarshalSeriesFull validates and renders an instance of Series into a interface{}
// using view "full".
func MarshalSeriesFull(source *Series, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp36 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	if source.Account != nil {
		tmp36["account"], err = MarshalAccount(source.Account, err)
	}
	target = tmp36
	return
}

// MarshalSeriesLink validates and renders an instance of Series into a interface{}
// using view "link".
func MarshalSeriesLink(source *Series, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp37 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	target = tmp37
	return
}

// MarshalSeriesTiny validates and renders an instance of Series into a interface{}
// using view "tiny".
func MarshalSeriesTiny(source *Series, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp38 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	target = tmp38
	return
}

// UnmarshalSeries unmarshals and validates a raw interface{} into an instance of Series
func UnmarshalSeries(source interface{}, inErr error) (target *Series, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Series)
		if v, ok := val["account"]; ok {
			var tmp39 *Account
			tmp39, err = UnmarshalAccount(v, err)
			target.Account = tmp39
		}
		if v, ok := val["href"]; ok {
			var tmp40 string
			if val, ok := v.(string); ok {
				tmp40 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp40
		}
		if v, ok := val["id"]; ok {
			var tmp41 int
			if f, ok := v.(float64); ok {
				tmp41 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp41
		}
		if v, ok := val["name"]; ok {
			var tmp42 string
			if val, ok := v.(string); ok {
				tmp42 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp42) < 2 {
					err = goa.InvalidLengthError(`load.Name`, tmp42, 2, true, err)
				}
			}
			target.Name = tmp42
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// SeriesCollection media type
// Identifier: application/vnd.series+json; type=collection
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
func LoadSeriesCollection(raw interface{}) (res SeriesCollection, err error) {
	res, err = UnmarshalSeriesCollection(raw, err)
	return
}

// Dump produces raw data from an instance of SeriesCollection running all the
// validations. See LoadSeriesCollection for the definition of raw data.
func (mt SeriesCollection) Dump(view SeriesCollectionViewEnum) (res []map[string]interface{}, err error) {
	if view == SeriesCollectionDefaultView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp43 := range mt {
			var tmp44 map[string]interface{}
			tmp44, err = MarshalSeries(tmp43, err)
			res[i] = tmp44
		}
	}
	if view == SeriesCollectionTinyView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp45 := range mt {
			var tmp46 map[string]interface{}
			tmp46, err = MarshalSeriesTiny(tmp45, err)
			res[i] = tmp46
		}
	}
	return
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

// MarshalSeriesCollection validates and renders an instance of SeriesCollection into a interface{}
// using view "default".
func MarshalSeriesCollection(source SeriesCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalSeries(res, err)
	}
	return
}

// MarshalSeriesCollectionTiny validates and renders an instance of SeriesCollection into a interface{}
// using view "tiny".
func MarshalSeriesCollectionTiny(source SeriesCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalSeriesTiny(res, err)
	}
	return
}

// UnmarshalSeriesCollection unmarshals and validates a raw interface{} into an instance of SeriesCollection
func UnmarshalSeriesCollection(source interface{}, inErr error) (target SeriesCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Series, len(val))
		for i, v := range val {
			target[i], err = UnmarshalSeries(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}

// A user belonging to a tenant account
// Identifier: application/vnd.user+json
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
	if len(mt.Email) < 2 {
		err = goa.InvalidLengthError(`response.email`, mt.Email, 2, true, err)
	}
	if mt.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2, err)
		}
	}
	if len(mt.FirstName) < 2 {
		err = goa.InvalidLengthError(`response.first_name`, mt.FirstName, 2, true, err)
	}
	if len(mt.LastName) < 2 {
		err = goa.InvalidLengthError(`response.last_name`, mt.LastName, 2, true, err)
	}
	return
}

// MarshalUser validates and renders an instance of User into a interface{}
// using view "default".
func MarshalUser(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Email) < 2 {
		err = goa.InvalidLengthError(`.email`, source.Email, 2, true, err)
	}
	if source.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, source.Email); err2 != nil {
			err = goa.InvalidFormatError(`.email`, source.Email, goa.FormatEmail, err2, err)
		}
	}
	if len(source.FirstName) < 2 {
		err = goa.InvalidLengthError(`.first_name`, source.FirstName, 2, true, err)
	}
	if len(source.LastName) < 2 {
		err = goa.InvalidLengthError(`.last_name`, source.LastName, 2, true, err)
	}
	tmp47 := map[string]interface{}{
		"email":      source.Email,
		"first_name": source.FirstName,
		"href":       source.Href,
		"id":         source.ID,
		"last_name":  source.LastName,
	}
	target = tmp47
	return
}

// MarshalUserLink validates and renders an instance of User into a interface{}
// using view "link".
func MarshalUserLink(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Email) < 2 {
		err = goa.InvalidLengthError(`.email`, source.Email, 2, true, err)
	}
	if source.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, source.Email); err2 != nil {
			err = goa.InvalidFormatError(`.email`, source.Email, goa.FormatEmail, err2, err)
		}
	}
	tmp48 := map[string]interface{}{
		"email": source.Email,
		"href":  source.Href,
		"id":    source.ID,
	}
	target = tmp48
	return
}

// UnmarshalUser unmarshals and validates a raw interface{} into an instance of User
func UnmarshalUser(source interface{}, inErr error) (target *User, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(User)
		if v, ok := val["email"]; ok {
			var tmp49 string
			if val, ok := v.(string); ok {
				tmp49 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp49) < 2 {
					err = goa.InvalidLengthError(`load.Email`, tmp49, 2, true, err)
				}
				if tmp49 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp49); err2 != nil {
						err = goa.InvalidFormatError(`load.Email`, tmp49, goa.FormatEmail, err2, err)
					}
				}
			}
			target.Email = tmp49
		}
		if v, ok := val["first_name"]; ok {
			var tmp50 string
			if val, ok := v.(string); ok {
				tmp50 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.FirstName`, v, "string", err)
			}
			if err == nil {
				if len(tmp50) < 2 {
					err = goa.InvalidLengthError(`load.FirstName`, tmp50, 2, true, err)
				}
			}
			target.FirstName = tmp50
		}
		if v, ok := val["href"]; ok {
			var tmp51 string
			if val, ok := v.(string); ok {
				tmp51 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp51
		}
		if v, ok := val["id"]; ok {
			var tmp52 int
			if f, ok := v.(float64); ok {
				tmp52 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp52
		}
		if v, ok := val["last_name"]; ok {
			var tmp53 string
			if val, ok := v.(string); ok {
				tmp53 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.LastName`, v, "string", err)
			}
			if err == nil {
				if len(tmp53) < 2 {
					err = goa.InvalidLengthError(`load.LastName`, tmp53, 2, true, err)
				}
			}
			target.LastName = tmp53
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// UserCollection media type
// Identifier: application/vnd.user+json; type=collection
type UserCollection []*User

// LoadUserCollection loads raw data into an instance of UserCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
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
	for i, tmp54 := range mt {
		var tmp55 map[string]interface{}
		tmp55, err = MarshalUser(tmp54, err)
		res[i] = tmp55
	}
	return
}

// Validate validates the media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Email) < 2 {
			err = goa.InvalidLengthError(`response[*].email`, e.Email, 2, true, err)
		}
		if e.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, e.Email, goa.FormatEmail, err2, err)
			}
		}
		if len(e.FirstName) < 2 {
			err = goa.InvalidLengthError(`response[*].first_name`, e.FirstName, 2, true, err)
		}
		if len(e.LastName) < 2 {
			err = goa.InvalidLengthError(`response[*].last_name`, e.LastName, 2, true, err)
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
		for i, v := range val {
			target[i], err = UnmarshalUser(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}
