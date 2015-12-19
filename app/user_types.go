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
	tmp64 := map[string]interface{}{
		"abstract":  source.Abstract,
		"detail":    source.Detail,
		"firstname": source.Firstname,
		"title":     source.Title,
		"withdrawn": source.Withdrawn,
	}
	target = tmp64
	return
}

// UnmarshalProposalModel unmarshals and validates a raw interface{} into an instance of ProposalModel
func UnmarshalProposalModel(source interface{}, inErr error) (target *ProposalModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ProposalModel)
		if v, ok := val["abstract"]; ok {
			var tmp65 string
			if val, ok := v.(string); ok {
				tmp65 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp65) < 50 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp65, 50, true, err)
				}
				if len(tmp65) > 500 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp65, 500, false, err)
				}
			}
			target.Abstract = tmp65
		}
		if v, ok := val["detail"]; ok {
			var tmp66 string
			if val, ok := v.(string); ok {
				tmp66 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp66) < 100 {
					err = goa.InvalidLengthError(`load.Detail`, tmp66, 100, true, err)
				}
				if len(tmp66) > 2000 {
					err = goa.InvalidLengthError(`load.Detail`, tmp66, 2000, false, err)
				}
			}
			target.Detail = tmp66
		}
		if v, ok := val["firstname"]; ok {
			var tmp67 string
			if val, ok := v.(string); ok {
				tmp67 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			if err == nil {
				if len(tmp67) < 2 {
					err = goa.InvalidLengthError(`load.Firstname`, tmp67, 2, true, err)
				}
			}
			target.Firstname = tmp67
		}
		if v, ok := val["title"]; ok {
			var tmp68 string
			if val, ok := v.(string); ok {
				tmp68 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp68) < 10 {
					err = goa.InvalidLengthError(`load.Title`, tmp68, 10, true, err)
				}
				if len(tmp68) > 200 {
					err = goa.InvalidLengthError(`load.Title`, tmp68, 200, false, err)
				}
			}
			target.Title = tmp68
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp69 bool
			if val, ok := v.(bool); ok {
				tmp69 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = tmp69
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
	tmp70 := map[string]interface{}{
		"comment": source.Comment,
		"rating":  source.Rating,
	}
	target = tmp70
	return
}

// UnmarshalReviewModel unmarshals and validates a raw interface{} into an instance of ReviewModel
func UnmarshalReviewModel(source interface{}, inErr error) (target *ReviewModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ReviewModel)
		if v, ok := val["comment"]; ok {
			var tmp71 string
			if val, ok := v.(string); ok {
				tmp71 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp71) < 10 {
					err = goa.InvalidLengthError(`load.Comment`, tmp71, 10, true, err)
				}
				if len(tmp71) > 200 {
					err = goa.InvalidLengthError(`load.Comment`, tmp71, 200, false, err)
				}
			}
			target.Comment = tmp71
		}
		if v, ok := val["rating"]; ok {
			var tmp72 int
			if f, ok := v.(float64); ok {
				tmp72 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp72 < 1 {
					err = goa.InvalidRangeError(`load.Rating`, tmp72, 1, true, err)
				}
				if tmp72 > 5 {
					err = goa.InvalidRangeError(`load.Rating`, tmp72, 5, false, err)
				}
			}
			target.Rating = tmp72
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
	Role      string
	State     string
}

// MarshalUserModel validates and renders an instance of UserModel into a interface{}
func MarshalUserModel(source *UserModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Bio) > 500 {
		err = goa.InvalidLengthError(`.bio`, source.Bio, 500, false, err)
	}
	tmp73 := map[string]interface{}{
		"bio":       source.Bio,
		"city":      source.City,
		"country":   source.Country,
		"email":     source.Email,
		"firstname": source.Firstname,
		"lastname":  source.Lastname,
		"role":      source.Role,
		"state":     source.State,
	}
	target = tmp73
	return
}

// UnmarshalUserModel unmarshals and validates a raw interface{} into an instance of UserModel
func UnmarshalUserModel(source interface{}, inErr error) (target *UserModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UserModel)
		if v, ok := val["bio"]; ok {
			var tmp74 string
			if val, ok := v.(string); ok {
				tmp74 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp74) > 500 {
					err = goa.InvalidLengthError(`load.Bio`, tmp74, 500, false, err)
				}
			}
			target.Bio = tmp74
		}
		if v, ok := val["city"]; ok {
			var tmp75 string
			if val, ok := v.(string); ok {
				tmp75 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = tmp75
		}
		if v, ok := val["country"]; ok {
			var tmp76 string
			if val, ok := v.(string); ok {
				tmp76 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = tmp76
		}
		if v, ok := val["email"]; ok {
			var tmp77 string
			if val, ok := v.(string); ok {
				tmp77 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			target.Email = tmp77
		}
		if v, ok := val["firstname"]; ok {
			var tmp78 string
			if val, ok := v.(string); ok {
				tmp78 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			target.Firstname = tmp78
		}
		if v, ok := val["lastname"]; ok {
			var tmp79 string
			if val, ok := v.(string); ok {
				tmp79 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Lastname`, v, "string", err)
			}
			target.Lastname = tmp79
		}
		if v, ok := val["role"]; ok {
			var tmp80 string
			if val, ok := v.(string); ok {
				tmp80 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = tmp80
		}
		if v, ok := val["state"]; ok {
			var tmp81 string
			if val, ok := v.(string); ok {
				tmp81 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = tmp81
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
