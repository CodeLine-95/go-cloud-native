package structs

import (
	"reflect"
	"strings"
)

// tagOptions contains a slice of tag options
type tagOptions []string

// Has returns true if the given option is available in tagOptions
func (t tagOptions) Has(opt string) bool {
	for _, tagOpt := range t {
		if tagOpt == opt {
			return true
		}
	}

	return false
}

// parseTag splits a struct field's tag into its name and a list of options
// which comes after a name. A tag is in the form of: "name,option1,option2".
// The name can be neglectected.
func parseTag(tag string) (string, tagOptions) {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"

	res := strings.Split(tag, ",")
	return res[0], res[1:]
}

// ToTags 解析 struct 种的 tag 标签
func ToTags(c any, tagName string) (fields []string) {
	tagFields := []reflect.StructField{}

	ct := reflect.TypeOf(c)
	for i := 0; i < ct.NumField(); i++ {
		tagFields = append(tagFields, typeFields(ct.Field(i))...)
	}

	for _, val := range tagFields {
		tagValue := val.Tag.Get(tagName)
		if tagValue != "" && tagValue != jsonFilter {
			fields = append(fields, strings.Split(tagValue, ",")[0])
		}
	}
	return fields
}

// typeFields 通过递归获取全部的 tag
func typeFields(field reflect.StructField) []reflect.StructField {
	fieldsMap := []reflect.StructField{}
	switch field.Type.Kind() {
	case reflect.Struct:
		for j := 0; j < field.Type.NumField(); j++ {
			fieldsMap = append(fieldsMap, typeFields(field.Type.Field(j))...)
		}
	default:
		fieldsMap = append(fieldsMap, field)
	}
	return fieldsMap
}
