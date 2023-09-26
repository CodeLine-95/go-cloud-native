package structs

import "reflect"

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
