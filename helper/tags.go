package helper

import (
	"reflect"
	"strings"
)

// StructTags tag in struct
type StructTags struct {
	Tag    string     `json:"tag"`
	Fields []FieldTag `json:"fields"`
}

// FieldTag tag string in field
type FieldTag struct {
	Index     int    `json:"index"`
	Field     string `json:"field"`
	TagName   string `json:"tag_name"`
	TagOption string `json:"tag_option"`
}

// New tags in struct
func New(i any, tag string) *StructTags {
	s := i
	rv := reflect.ValueOf(s)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	st := rv.Type()

	tags := StructTags{Tag: tag}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		t := FieldTag{Index: i, Field: field.Name}
		if tag == "gorm" {
			t.TagName, t.TagOption = parseTagGORM(field)
		} else {
			t.TagName, t.TagOption = parseTagDefault(field, tag)
		}

		tags.Fields = append(tags.Fields, t)
	}
	return &tags
}

func parseTagDefault(field reflect.StructField, tagName string) (string, string) {
	if tag, ok := field.Tag.Lookup(tagName); ok {
		if tag == "-" {
			return "-", ""
		}
		idx := strings.SplitN(tag, ",", 2)
		if len(idx) > 1 {
			return idx[0], idx[1]
		}
		return idx[0], ""
	}
	return field.Name, ""
}

func parseTagGORM(field reflect.StructField) (string, string) {
	sep := ";"
	if tag, ok := field.Tag.Lookup("gorm"); ok {
		if tag == "-" {
			return "-", ""
		}
		names := strings.Split(tag, sep)
		for i := 0; i < len(names); i++ {
			j := i
			if len(names[j]) > 0 {
				for {
					if names[j][len(names[j])-1] == '\\' {
						i++
						names[j] = names[j][0:len(names[j])-1] + sep + names[i]
						names[i] = ""
					} else {
						break
					}
				}
			}
			values := strings.Split(names[j], ":")
			k := strings.TrimSpace(strings.ToUpper(values[0]))
			if k == "COLUMN" {
				if len(values) >= 2 {
					return values[1], tag
				}
			}
		}
	}
	return field.Name, ""
}
