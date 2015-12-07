//************************************************************************//
// congo: Application User Types
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

// AccountModel type
type AccountModel struct {
	Name string
}

// MarshalAccountModel validates and renders an instance of AccountModel into a interface{}
func MarshalAccountModel(source *AccountModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp56 := map[string]interface{}{
		"name": source.Name,
	}
	target = tmp56
	return
}

// UnmarshalAccountModel unmarshals and validates a raw interface{} into an instance of AccountModel
func UnmarshalAccountModel(source interface{}, inErr error) (target *AccountModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(AccountModel)
		if v, ok := val["name"]; ok {
			var tmp57 string
			if val, ok := v.(string); ok {
				tmp57 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp57) < 2 {
					err = goa.InvalidLengthError(`load.Name`, tmp57, 2, true, err)
				}
			}
			target.Name = tmp57
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// InstanceModel type
type InstanceModel struct {
	Name string
}

// MarshalInstanceModel validates and renders an instance of InstanceModel into a interface{}
func MarshalInstanceModel(source *InstanceModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp58 := map[string]interface{}{
		"name": source.Name,
	}
	target = tmp58
	return
}

// UnmarshalInstanceModel unmarshals and validates a raw interface{} into an instance of InstanceModel
func UnmarshalInstanceModel(source interface{}, inErr error) (target *InstanceModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(InstanceModel)
		if v, ok := val["name"]; ok {
			var tmp59 string
			if val, ok := v.(string); ok {
				tmp59 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp59) < 2 {
					err = goa.InvalidLengthError(`load.Name`, tmp59, 2, true, err)
				}
			}
			target.Name = tmp59
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// SeriesModel type
type SeriesModel struct {
	Name string
}

// MarshalSeriesModel validates and renders an instance of SeriesModel into a interface{}
func MarshalSeriesModel(source *SeriesModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Name) < 2 {
		err = goa.InvalidLengthError(`.name`, source.Name, 2, true, err)
	}
	tmp60 := map[string]interface{}{
		"name": source.Name,
	}
	target = tmp60
	return
}

// UnmarshalSeriesModel unmarshals and validates a raw interface{} into an instance of SeriesModel
func UnmarshalSeriesModel(source interface{}, inErr error) (target *SeriesModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(SeriesModel)
		if v, ok := val["name"]; ok {
			var tmp61 string
			if val, ok := v.(string); ok {
				tmp61 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp61) < 2 {
					err = goa.InvalidLengthError(`load.Name`, tmp61, 2, true, err)
				}
			}
			target.Name = tmp61
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// UserModel type
type UserModel struct {
	Email     string
	FirstName string
	ID        int
	LastName  string
	Role      string
}

// MarshalUserModel validates and renders an instance of UserModel into a interface{}
func MarshalUserModel(source *UserModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Email) < 2 {
		err = goa.InvalidLengthError(`.email`, source.Email, 2, true, err)
	}
	if len(source.FirstName) < 2 {
		err = goa.InvalidLengthError(`.first_name`, source.FirstName, 2, true, err)
	}
	if len(source.LastName) < 2 {
		err = goa.InvalidLengthError(`.last_name`, source.LastName, 2, true, err)
	}
	tmp62 := map[string]interface{}{
		"email":      source.Email,
		"first_name": source.FirstName,
		"id":         source.ID,
		"last_name":  source.LastName,
		"role":       source.Role,
	}
	target = tmp62
	return
}

// UnmarshalUserModel unmarshals and validates a raw interface{} into an instance of UserModel
func UnmarshalUserModel(source interface{}, inErr error) (target *UserModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UserModel)
		if v, ok := val["email"]; ok {
			var tmp63 string
			if val, ok := v.(string); ok {
				tmp63 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp63) < 2 {
					err = goa.InvalidLengthError(`load.Email`, tmp63, 2, true, err)
				}
			}
			target.Email = tmp63
		}
		if v, ok := val["first_name"]; ok {
			var tmp64 string
			if val, ok := v.(string); ok {
				tmp64 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.FirstName`, v, "string", err)
			}
			if err == nil {
				if len(tmp64) < 2 {
					err = goa.InvalidLengthError(`load.FirstName`, tmp64, 2, true, err)
				}
			}
			target.FirstName = tmp64
		}
		if v, ok := val["id"]; ok {
			var tmp65 int
			if f, ok := v.(float64); ok {
				tmp65 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp65
		}
		if v, ok := val["last_name"]; ok {
			var tmp66 string
			if val, ok := v.(string); ok {
				tmp66 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.LastName`, v, "string", err)
			}
			if err == nil {
				if len(tmp66) < 2 {
					err = goa.InvalidLengthError(`load.LastName`, tmp66, 2, true, err)
				}
			}
			target.LastName = tmp66
		}
		if v, ok := val["role"]; ok {
			var tmp67 string
			if val, ok := v.(string); ok {
				tmp67 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = tmp67
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
