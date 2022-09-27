package valistruct

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

var cacheValidate *validator.Validate

func init() {
	cacheValidate = NewValidate()
}

func NewValidate() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(GetJsonTag)
	return v
}

func GetValidate() *validator.Validate {
	return cacheValidate
}

func GetJsonTag(fld reflect.StructField) string {
	s, _ := helper.New(fld, "json")
	if name == "-" {
		return
	}
	return name
}
