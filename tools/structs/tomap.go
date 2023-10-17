package structs

import (
	"reflect"
)

const jsonFilter = "child_node"

func ToMap(item any) map[string]any {
	if isNil(item) {
		return nil
	}
	s := newStruct(item, "json", false)
	return s.tomap()
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
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
			fields = append(fields, tagValue)
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
