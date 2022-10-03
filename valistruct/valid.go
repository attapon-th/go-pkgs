// Package validstruct helper for validate struct
// with https://github.com/go-playground/validator v10
package validstruct

import (
	"fmt"
	"reflect"

	"github.com/attapon-th/go-pkgs/helper"
	"github.com/go-playground/validator/v10"
)

// ValidStruct - object validate struct
type ValidStruct struct {
	*validator.Validate

	tag string

	ValidErrs []FieldError

	Err error
}

// New default
func New(filedNameByTag string) *ValidStruct {
	v := &ValidStruct{
		Validate: validator.New(),
		tag:      "json",
	}
	v.RegisterTagNameFunc(v.getTag)
	return v
}

func (v ValidStruct) getTag(fld reflect.StructField) string {
	tag := v.tag
	if tag == "" {
		tag = "json"
	}
	s, _ := helper.ParseField(fld, tag)
	if s == "-" {
		return ""
	}
	return s
}

func (v ValidStruct) Error() string {
	if v.Err != nil {
		return v.Error()
	}
	s := ""
	for _, v := range v.ValidErrs {
		s += fmt.Sprintf("%s;\n", v.Error())
	}
	return s
}

// Struct vaclidate struct
//
//	@receiver v
//	@param st any(struct only)
//	@return error
func (v *ValidStruct) Struct(st any) error {
	v.ValidErrs = []FieldError{}
	t := reflect.TypeOf(st)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		v.Err = ErrInValidStruct
		return v.Err
	}
	if err := v.Validate.Struct(st); err != nil {
		if er, ok := err.(*validator.InvalidValidationError); ok {
			v.Err = er
			return v.Err
		}
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, er := range errs {
				v.ValidErrs = append(v.ValidErrs, fromValidatorFieldError(er))
			}
		}
	}
	return nil
}

// IsValid is sucess valid
//
//	@return bool
func (v *ValidStruct) IsValid(st any) bool {
	if v.Err != nil || len(v.ValidErrs) > 0 {
		return false
	}
	return true
}
