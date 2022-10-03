package validstruct

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FieldError field error
type FieldError struct {
	Namespace   string       `json:"-"`
	StructField string       `json:"-"`
	Field       string       `json:"field"`
	Tag         string       `json:"valid"`
	ActualTag   string       `json:"detail"`
	Param       string       `json:"param"`
	Value       interface{}  `json:"-"`
	Kind        reflect.Kind `json:"-"`
	Type        reflect.Type `json:"-"`
	Err         string       `json:"massage"`
}

func (f FieldError) Error() string {
	return f.Err
}

func fromValidatorFieldError(er validator.FieldError) FieldError {
	var v FieldError
	v.Namespace = er.Namespace()
	v.Field = er.Field()
	v.StructField = er.StructField()
	v.Tag = er.Tag()
	v.ActualTag = er.ActualTag()
	v.Param = er.Param()
	v.Value = er.Value()
	v.Type = er.Type()
	v.Kind = er.Kind()
	// ut.New(locales.Translator)
	// th.RegisterDefaultTranslations(v)
	// log.Fatalf("%s", v.Namespace)
	s := strings.ReplaceAll(er.Error(), v.Namespace, v.Field)
	v.Err = s
	return v
}
