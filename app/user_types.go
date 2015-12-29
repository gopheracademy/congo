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

import (
	"github.com/raphael/goa"
)

// ProposalModel type
type ProposalModel struct {
	Abstract  string
	Detail    string
	Firstname string
	Title     string
	Withdrawn bool
}

// MarshalProposalModel validates and renders an instance of ProposalModel into a interface{}
func MarshalProposalModel(source *ProposalModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Abstract) < 50 {
		err = goa.InvalidLengthError(`.abstract`, source.Abstract, 50, true, err)
	}
	if len(source.Abstract) > 500 {
		err = goa.InvalidLengthError(`.abstract`, source.Abstract, 500, false, err)
	}
	if len(source.Detail) < 100 {
		err = goa.InvalidLengthError(`.detail`, source.Detail, 100, true, err)
	}
	if len(source.Detail) > 2000 {
		err = goa.InvalidLengthError(`.detail`, source.Detail, 2000, false, err)
	}
	if len(source.Firstname) < 2 {
		err = goa.InvalidLengthError(`.firstname`, source.Firstname, 2, true, err)
	}
	if len(source.Title) < 10 {
		err = goa.InvalidLengthError(`.title`, source.Title, 10, true, err)
	}
	if len(source.Title) > 200 {
		err = goa.InvalidLengthError(`.title`, source.Title, 200, false, err)
	}
	tmp81 := map[string]interface{}{
		"abstract":  source.Abstract,
		"detail":    source.Detail,
		"firstname": source.Firstname,
		"title":     source.Title,
		"withdrawn": source.Withdrawn,
	}
	target = tmp81
	return
}

// UnmarshalProposalModel unmarshals and validates a raw interface{} into an instance of ProposalModel
func UnmarshalProposalModel(source interface{}, inErr error) (target *ProposalModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ProposalModel)
		if v, ok := val["abstract"]; ok {
			var tmp82 string
			if val, ok := v.(string); ok {
				tmp82 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp82) < 50 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp82, 50, true, err)
				}
				if len(tmp82) > 500 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp82, 500, false, err)
				}
			}
			target.Abstract = tmp82
		}
		if v, ok := val["detail"]; ok {
			var tmp83 string
			if val, ok := v.(string); ok {
				tmp83 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp83) < 100 {
					err = goa.InvalidLengthError(`load.Detail`, tmp83, 100, true, err)
				}
				if len(tmp83) > 2000 {
					err = goa.InvalidLengthError(`load.Detail`, tmp83, 2000, false, err)
				}
			}
			target.Detail = tmp83
		}
		if v, ok := val["firstname"]; ok {
			var tmp84 string
			if val, ok := v.(string); ok {
				tmp84 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			if err == nil {
				if len(tmp84) < 2 {
					err = goa.InvalidLengthError(`load.Firstname`, tmp84, 2, true, err)
				}
			}
			target.Firstname = tmp84
		}
		if v, ok := val["title"]; ok {
			var tmp85 string
			if val, ok := v.(string); ok {
				tmp85 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp85) < 10 {
					err = goa.InvalidLengthError(`load.Title`, tmp85, 10, true, err)
				}
				if len(tmp85) > 200 {
					err = goa.InvalidLengthError(`load.Title`, tmp85, 200, false, err)
				}
			}
			target.Title = tmp85
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp86 bool
			if val, ok := v.(bool); ok {
				tmp86 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = tmp86
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// ReviewModel type
type ReviewModel struct {
	Comment string
	Rating  int
}

// MarshalReviewModel validates and renders an instance of ReviewModel into a interface{}
func MarshalReviewModel(source *ReviewModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Comment) < 10 {
		err = goa.InvalidLengthError(`.comment`, source.Comment, 10, true, err)
	}
	if len(source.Comment) > 200 {
		err = goa.InvalidLengthError(`.comment`, source.Comment, 200, false, err)
	}
	if source.Rating < 1 {
		err = goa.InvalidRangeError(`.rating`, source.Rating, 1, true, err)
	}
	if source.Rating > 5 {
		err = goa.InvalidRangeError(`.rating`, source.Rating, 5, false, err)
	}
	tmp87 := map[string]interface{}{
		"comment": source.Comment,
		"rating":  source.Rating,
	}
	target = tmp87
	return
}

// UnmarshalReviewModel unmarshals and validates a raw interface{} into an instance of ReviewModel
func UnmarshalReviewModel(source interface{}, inErr error) (target *ReviewModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ReviewModel)
		if v, ok := val["comment"]; ok {
			var tmp88 string
			if val, ok := v.(string); ok {
				tmp88 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp88) < 10 {
					err = goa.InvalidLengthError(`load.Comment`, tmp88, 10, true, err)
				}
				if len(tmp88) > 200 {
					err = goa.InvalidLengthError(`load.Comment`, tmp88, 200, false, err)
				}
			}
			target.Comment = tmp88
		}
		if v, ok := val["rating"]; ok {
			var tmp89 int
			if f, ok := v.(float64); ok {
				tmp89 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp89 < 1 {
					err = goa.InvalidRangeError(`load.Rating`, tmp89, 1, true, err)
				}
				if tmp89 > 5 {
					err = goa.InvalidRangeError(`load.Rating`, tmp89, 5, false, err)
				}
			}
			target.Rating = tmp89
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// UserModel type
type UserModel struct {
	Bio       string
	City      string
	Country   string
	Email     string
	Firstname string
	Lastname  string
	Password  string
	Role      string
	State     string
}

// MarshalUserModel validates and renders an instance of UserModel into a interface{}
func MarshalUserModel(source *UserModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Bio) > 500 {
		err = goa.InvalidLengthError(`.bio`, source.Bio, 500, false, err)
	}
	if len(source.Email) < 2 {
		err = goa.InvalidLengthError(`.email`, source.Email, 2, true, err)
	}
	if len(source.Password) < 8 {
		err = goa.InvalidLengthError(`.password`, source.Password, 8, true, err)
	}
	tmp90 := map[string]interface{}{
		"bio":       source.Bio,
		"city":      source.City,
		"country":   source.Country,
		"email":     source.Email,
		"firstname": source.Firstname,
		"lastname":  source.Lastname,
		"password":  source.Password,
		"role":      source.Role,
		"state":     source.State,
	}
	target = tmp90
	return
}

// UnmarshalUserModel unmarshals and validates a raw interface{} into an instance of UserModel
func UnmarshalUserModel(source interface{}, inErr error) (target *UserModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UserModel)
		if v, ok := val["bio"]; ok {
			var tmp91 string
			if val, ok := v.(string); ok {
				tmp91 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp91) > 500 {
					err = goa.InvalidLengthError(`load.Bio`, tmp91, 500, false, err)
				}
			}
			target.Bio = tmp91
		}
		if v, ok := val["city"]; ok {
			var tmp92 string
			if val, ok := v.(string); ok {
				tmp92 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = tmp92
		}
		if v, ok := val["country"]; ok {
			var tmp93 string
			if val, ok := v.(string); ok {
				tmp93 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = tmp93
		}
		if v, ok := val["email"]; ok {
			var tmp94 string
			if val, ok := v.(string); ok {
				tmp94 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp94) < 2 {
					err = goa.InvalidLengthError(`load.Email`, tmp94, 2, true, err)
				}
			}
			target.Email = tmp94
		}
		if v, ok := val["firstname"]; ok {
			var tmp95 string
			if val, ok := v.(string); ok {
				tmp95 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			target.Firstname = tmp95
		}
		if v, ok := val["lastname"]; ok {
			var tmp96 string
			if val, ok := v.(string); ok {
				tmp96 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Lastname`, v, "string", err)
			}
			target.Lastname = tmp96
		}
		if v, ok := val["password"]; ok {
			var tmp97 string
			if val, ok := v.(string); ok {
				tmp97 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Password`, v, "string", err)
			}
			if err == nil {
				if len(tmp97) < 8 {
					err = goa.InvalidLengthError(`load.Password`, tmp97, 8, true, err)
				}
			}
			target.Password = tmp97
		}
		if v, ok := val["role"]; ok {
			var tmp98 string
			if val, ok := v.(string); ok {
				tmp98 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = tmp98
		}
		if v, ok := val["state"]; ok {
			var tmp99 string
			if val, ok := v.(string); ok {
				tmp99 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = tmp99
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
