//************************************************************************//
// API "congo": Application User Types
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

// Validate validates the type instance.
func (ut *ProposalModel) Validate() (err error) {
	if len(ut.Abstract) < 50 {
		err = goa.InvalidLengthError(`response.abstract`, ut.Abstract, len(ut.Abstract), 50, true, err)
	}
	if len(ut.Abstract) > 500 {
		err = goa.InvalidLengthError(`response.abstract`, ut.Abstract, len(ut.Abstract), 500, false, err)
	}
	if len(ut.Detail) < 100 {
		err = goa.InvalidLengthError(`response.detail`, ut.Detail, len(ut.Detail), 100, true, err)
	}
	if len(ut.Detail) > 2000 {
		err = goa.InvalidLengthError(`response.detail`, ut.Detail, len(ut.Detail), 2000, false, err)
	}
	if len(ut.Firstname) < 2 {
		err = goa.InvalidLengthError(`response.firstname`, ut.Firstname, len(ut.Firstname), 2, true, err)
	}
	if len(ut.Title) < 10 {
		err = goa.InvalidLengthError(`response.title`, ut.Title, len(ut.Title), 10, true, err)
	}
	if len(ut.Title) > 200 {
		err = goa.InvalidLengthError(`response.title`, ut.Title, len(ut.Title), 200, false, err)
	}
	return
}

// MarshalProposalModel validates and renders an instance of ProposalModel into a interface{}
func MarshalProposalModel(source *ProposalModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Abstract) < 50 {
		err = goa.InvalidLengthError(`.abstract`, source.Abstract, len(source.Abstract), 50, true, err)
	}
	if len(source.Abstract) > 500 {
		err = goa.InvalidLengthError(`.abstract`, source.Abstract, len(source.Abstract), 500, false, err)
	}
	if len(source.Detail) < 100 {
		err = goa.InvalidLengthError(`.detail`, source.Detail, len(source.Detail), 100, true, err)
	}
	if len(source.Detail) > 2000 {
		err = goa.InvalidLengthError(`.detail`, source.Detail, len(source.Detail), 2000, false, err)
	}
	if len(source.Firstname) < 2 {
		err = goa.InvalidLengthError(`.firstname`, source.Firstname, len(source.Firstname), 2, true, err)
	}
	if len(source.Title) < 10 {
		err = goa.InvalidLengthError(`.title`, source.Title, len(source.Title), 10, true, err)
	}
	if len(source.Title) > 200 {
		err = goa.InvalidLengthError(`.title`, source.Title, len(source.Title), 200, false, err)
	}
	tmp79 := map[string]interface{}{
		"abstract":  source.Abstract,
		"detail":    source.Detail,
		"firstname": source.Firstname,
		"title":     source.Title,
		"withdrawn": source.Withdrawn,
	}
	target = tmp79
	return
}

// UnmarshalProposalModel unmarshals and validates a raw interface{} into an instance of ProposalModel
func UnmarshalProposalModel(source interface{}, inErr error) (target *ProposalModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ProposalModel)
		if v, ok := val["abstract"]; ok {
			var tmp80 string
			if val, ok := v.(string); ok {
				tmp80 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp80) < 50 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp80, len(tmp80), 50, true, err)
				}
				if len(tmp80) > 500 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp80, len(tmp80), 500, false, err)
				}
			}
			target.Abstract = tmp80
		}
		if v, ok := val["detail"]; ok {
			var tmp81 string
			if val, ok := v.(string); ok {
				tmp81 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp81) < 100 {
					err = goa.InvalidLengthError(`load.Detail`, tmp81, len(tmp81), 100, true, err)
				}
				if len(tmp81) > 2000 {
					err = goa.InvalidLengthError(`load.Detail`, tmp81, len(tmp81), 2000, false, err)
				}
			}
			target.Detail = tmp81
		}
		if v, ok := val["firstname"]; ok {
			var tmp82 string
			if val, ok := v.(string); ok {
				tmp82 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			if err == nil {
				if len(tmp82) < 2 {
					err = goa.InvalidLengthError(`load.Firstname`, tmp82, len(tmp82), 2, true, err)
				}
			}
			target.Firstname = tmp82
		}
		if v, ok := val["title"]; ok {
			var tmp83 string
			if val, ok := v.(string); ok {
				tmp83 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp83) < 10 {
					err = goa.InvalidLengthError(`load.Title`, tmp83, len(tmp83), 10, true, err)
				}
				if len(tmp83) > 200 {
					err = goa.InvalidLengthError(`load.Title`, tmp83, len(tmp83), 200, false, err)
				}
			}
			target.Title = tmp83
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp84 bool
			if val, ok := v.(bool); ok {
				tmp84 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = tmp84
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

// Validate validates the type instance.
func (ut *ReviewModel) Validate() (err error) {
	if len(ut.Comment) < 10 {
		err = goa.InvalidLengthError(`response.comment`, ut.Comment, len(ut.Comment), 10, true, err)
	}
	if len(ut.Comment) > 200 {
		err = goa.InvalidLengthError(`response.comment`, ut.Comment, len(ut.Comment), 200, false, err)
	}
	if ut.Rating < 1 {
		err = goa.InvalidRangeError(`response.rating`, ut.Rating, 1, true, err)
	}
	if ut.Rating > 5 {
		err = goa.InvalidRangeError(`response.rating`, ut.Rating, 5, false, err)
	}
	return
}

// MarshalReviewModel validates and renders an instance of ReviewModel into a interface{}
func MarshalReviewModel(source *ReviewModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Comment) < 10 {
		err = goa.InvalidLengthError(`.comment`, source.Comment, len(source.Comment), 10, true, err)
	}
	if len(source.Comment) > 200 {
		err = goa.InvalidLengthError(`.comment`, source.Comment, len(source.Comment), 200, false, err)
	}
	if source.Rating < 1 {
		err = goa.InvalidRangeError(`.rating`, source.Rating, 1, true, err)
	}
	if source.Rating > 5 {
		err = goa.InvalidRangeError(`.rating`, source.Rating, 5, false, err)
	}
	tmp85 := map[string]interface{}{
		"comment": source.Comment,
		"rating":  source.Rating,
	}
	target = tmp85
	return
}

// UnmarshalReviewModel unmarshals and validates a raw interface{} into an instance of ReviewModel
func UnmarshalReviewModel(source interface{}, inErr error) (target *ReviewModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ReviewModel)
		if v, ok := val["comment"]; ok {
			var tmp86 string
			if val, ok := v.(string); ok {
				tmp86 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp86) < 10 {
					err = goa.InvalidLengthError(`load.Comment`, tmp86, len(tmp86), 10, true, err)
				}
				if len(tmp86) > 200 {
					err = goa.InvalidLengthError(`load.Comment`, tmp86, len(tmp86), 200, false, err)
				}
			}
			target.Comment = tmp86
		}
		if v, ok := val["rating"]; ok {
			var tmp87 int
			if f, ok := v.(float64); ok {
				tmp87 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp87 < 1 {
					err = goa.InvalidRangeError(`load.Rating`, tmp87, 1, true, err)
				}
				if tmp87 > 5 {
					err = goa.InvalidRangeError(`load.Rating`, tmp87, 5, false, err)
				}
			}
			target.Rating = tmp87
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

// Validate validates the type instance.
func (ut *UserModel) Validate() (err error) {
	if len(ut.Bio) > 500 {
		err = goa.InvalidLengthError(`response.bio`, ut.Bio, len(ut.Bio), 500, false, err)
	}
	if ut.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, ut.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}

// MarshalUserModel validates and renders an instance of UserModel into a interface{}
func MarshalUserModel(source *UserModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Bio) > 500 {
		err = goa.InvalidLengthError(`.bio`, source.Bio, len(source.Bio), 500, false, err)
	}
	if source.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, source.Email); err2 != nil {
			err = goa.InvalidFormatError(`.email`, source.Email, goa.FormatEmail, err2, err)
		}
	}
	tmp88 := map[string]interface{}{
		"bio":       source.Bio,
		"city":      source.City,
		"country":   source.Country,
		"email":     source.Email,
		"firstname": source.Firstname,
		"lastname":  source.Lastname,
		"role":      source.Role,
		"state":     source.State,
	}
	target = tmp88
	return
}

// UnmarshalUserModel unmarshals and validates a raw interface{} into an instance of UserModel
func UnmarshalUserModel(source interface{}, inErr error) (target *UserModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UserModel)
		if v, ok := val["bio"]; ok {
			var tmp89 string
			if val, ok := v.(string); ok {
				tmp89 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp89) > 500 {
					err = goa.InvalidLengthError(`load.Bio`, tmp89, len(tmp89), 500, false, err)
				}
			}
			target.Bio = tmp89
		}
		if v, ok := val["city"]; ok {
			var tmp90 string
			if val, ok := v.(string); ok {
				tmp90 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = tmp90
		}
		if v, ok := val["country"]; ok {
			var tmp91 string
			if val, ok := v.(string); ok {
				tmp91 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = tmp91
		}
		if v, ok := val["email"]; ok {
			var tmp92 string
			if val, ok := v.(string); ok {
				tmp92 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			if err == nil {
				if tmp92 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp92); err2 != nil {
						err = goa.InvalidFormatError(`load.Email`, tmp92, goa.FormatEmail, err2, err)
					}
				}
			}
			target.Email = tmp92
		}
		if v, ok := val["firstname"]; ok {
			var tmp93 string
			if val, ok := v.(string); ok {
				tmp93 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			target.Firstname = tmp93
		}
		if v, ok := val["lastname"]; ok {
			var tmp94 string
			if val, ok := v.(string); ok {
				tmp94 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Lastname`, v, "string", err)
			}
			target.Lastname = tmp94
		}
		if v, ok := val["role"]; ok {
			var tmp95 string
			if val, ok := v.(string); ok {
				tmp95 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = tmp95
		}
		if v, ok := val["state"]; ok {
			var tmp96 string
			if val, ok := v.(string); ok {
				tmp96 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = tmp96
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
