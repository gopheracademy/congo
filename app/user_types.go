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

import "github.com/raphael/goa"

// ProposalModel type
type ProposalModel struct {
	Abstract  *string
	Detail    *string
	Firstname *string
	Title     *string
	Withdrawn *bool
}

// Validate validates the type instance.
func (ut *ProposalModel) Validate() (err error) {
	if ut.Abstract != nil {
		if len(*ut.Abstract) < 50 {
			err = goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 50, true, err)
		}
	}
	if ut.Abstract != nil {
		if len(*ut.Abstract) > 500 {
			err = goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 500, false, err)
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) < 100 {
			err = goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 100, true, err)
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) > 2000 {
			err = goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 2000, false, err)
		}
	}
	if ut.Firstname != nil {
		if len(*ut.Firstname) < 2 {
			err = goa.InvalidLengthError(`response.firstname`, *ut.Firstname, len(*ut.Firstname), 2, true, err)
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) < 10 {
			err = goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 10, true, err)
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) > 200 {
			err = goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 200, false, err)
		}
	}
	return
}

// MarshalProposalModel validates and renders an instance of ProposalModel into a interface{}
func MarshalProposalModel(source *ProposalModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp30 := map[string]interface{}{
		"abstract":  source.Abstract,
		"detail":    source.Detail,
		"firstname": source.Firstname,
		"title":     source.Title,
		"withdrawn": source.Withdrawn,
	}
	target = tmp30
	return
}

// UnmarshalProposalModel unmarshals and validates a raw interface{} into an instance of ProposalModel
func UnmarshalProposalModel(source interface{}, inErr error) (target *ProposalModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ProposalModel)
		if v, ok := val["abstract"]; ok {
			var tmp31 string
			if val, ok := v.(string); ok {
				tmp31 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp31) < 50 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp31, len(tmp31), 50, true, err)
				}
				if len(tmp31) > 500 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp31, len(tmp31), 500, false, err)
				}
			}
			target.Abstract = &tmp31
		}
		if v, ok := val["detail"]; ok {
			var tmp32 string
			if val, ok := v.(string); ok {
				tmp32 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp32) < 100 {
					err = goa.InvalidLengthError(`load.Detail`, tmp32, len(tmp32), 100, true, err)
				}
				if len(tmp32) > 2000 {
					err = goa.InvalidLengthError(`load.Detail`, tmp32, len(tmp32), 2000, false, err)
				}
			}
			target.Detail = &tmp32
		}
		if v, ok := val["firstname"]; ok {
			var tmp33 string
			if val, ok := v.(string); ok {
				tmp33 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			if err == nil {
				if len(tmp33) < 2 {
					err = goa.InvalidLengthError(`load.Firstname`, tmp33, len(tmp33), 2, true, err)
				}
			}
			target.Firstname = &tmp33
		}
		if v, ok := val["title"]; ok {
			var tmp34 string
			if val, ok := v.(string); ok {
				tmp34 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp34) < 10 {
					err = goa.InvalidLengthError(`load.Title`, tmp34, len(tmp34), 10, true, err)
				}
				if len(tmp34) > 200 {
					err = goa.InvalidLengthError(`load.Title`, tmp34, len(tmp34), 200, false, err)
				}
			}
			target.Title = &tmp34
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp35 bool
			if val, ok := v.(bool); ok {
				tmp35 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = &tmp35
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// ReviewModel type
type ReviewModel struct {
	Comment *string
	Rating  *int
}

// Validate validates the type instance.
func (ut *ReviewModel) Validate() (err error) {
	if ut.Comment != nil {
		if len(*ut.Comment) < 10 {
			err = goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 10, true, err)
		}
	}
	if ut.Comment != nil {
		if len(*ut.Comment) > 200 {
			err = goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 200, false, err)
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *ut.Rating, 1, true, err)
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *ut.Rating, 5, false, err)
		}
	}
	return
}

// MarshalReviewModel validates and renders an instance of ReviewModel into a interface{}
func MarshalReviewModel(source *ReviewModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp36 := map[string]interface{}{
		"comment": source.Comment,
		"rating":  source.Rating,
	}
	target = tmp36
	return
}

// UnmarshalReviewModel unmarshals and validates a raw interface{} into an instance of ReviewModel
func UnmarshalReviewModel(source interface{}, inErr error) (target *ReviewModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ReviewModel)
		if v, ok := val["comment"]; ok {
			var tmp37 string
			if val, ok := v.(string); ok {
				tmp37 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp37) < 10 {
					err = goa.InvalidLengthError(`load.Comment`, tmp37, len(tmp37), 10, true, err)
				}
				if len(tmp37) > 200 {
					err = goa.InvalidLengthError(`load.Comment`, tmp37, len(tmp37), 200, false, err)
				}
			}
			target.Comment = &tmp37
		}
		if v, ok := val["rating"]; ok {
			var tmp38 int
			if f, ok := v.(float64); ok {
				tmp38 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp38 < 1 {
					err = goa.InvalidRangeError(`load.Rating`, tmp38, 1, true, err)
				}
				if tmp38 > 5 {
					err = goa.InvalidRangeError(`load.Rating`, tmp38, 5, false, err)
				}
			}
			target.Rating = &tmp38
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// UserModel type
type UserModel struct {
	Bio       *string
	City      *string
	Country   *string
	Email     *string
	Firstname *string
	Lastname  *string
	Role      *string
	State     *string
}

// Validate validates the type instance.
func (ut *UserModel) Validate() (err error) {
	if ut.Bio != nil {
		if len(*ut.Bio) > 500 {
			err = goa.InvalidLengthError(`response.bio`, *ut.Bio, len(*ut.Bio), 500, false, err)
		}
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}

// MarshalUserModel validates and renders an instance of UserModel into a interface{}
func MarshalUserModel(source *UserModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp39 := map[string]interface{}{
		"bio":       source.Bio,
		"city":      source.City,
		"country":   source.Country,
		"email":     source.Email,
		"firstname": source.Firstname,
		"lastname":  source.Lastname,
		"role":      source.Role,
		"state":     source.State,
	}
	target = tmp39
	return
}

// UnmarshalUserModel unmarshals and validates a raw interface{} into an instance of UserModel
func UnmarshalUserModel(source interface{}, inErr error) (target *UserModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UserModel)
		if v, ok := val["bio"]; ok {
			var tmp40 string
			if val, ok := v.(string); ok {
				tmp40 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp40) > 500 {
					err = goa.InvalidLengthError(`load.Bio`, tmp40, len(tmp40), 500, false, err)
				}
			}
			target.Bio = &tmp40
		}
		if v, ok := val["city"]; ok {
			var tmp41 string
			if val, ok := v.(string); ok {
				tmp41 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = &tmp41
		}
		if v, ok := val["country"]; ok {
			var tmp42 string
			if val, ok := v.(string); ok {
				tmp42 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = &tmp42
		}
		if v, ok := val["email"]; ok {
			var tmp43 string
			if val, ok := v.(string); ok {
				tmp43 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			if err == nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, tmp43); err2 != nil {
					err = goa.InvalidFormatError(`load.Email`, tmp43, goa.FormatEmail, err2, err)
				}
			}
			target.Email = &tmp43
		}
		if v, ok := val["firstname"]; ok {
			var tmp44 string
			if val, ok := v.(string); ok {
				tmp44 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			target.Firstname = &tmp44
		}
		if v, ok := val["lastname"]; ok {
			var tmp45 string
			if val, ok := v.(string); ok {
				tmp45 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Lastname`, v, "string", err)
			}
			target.Lastname = &tmp45
		}
		if v, ok := val["role"]; ok {
			var tmp46 string
			if val, ok := v.(string); ok {
				tmp46 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = &tmp46
		}
		if v, ok := val["state"]; ok {
			var tmp47 string
			if val, ok := v.(string); ok {
				tmp47 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = &tmp47
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
