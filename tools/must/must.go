package must

import (
	"fmt"
	"strconv"
)

func StringToInt(str string) (ret int) {
	if str == "" {
		return
	}

	ret, _ = strconv.Atoi(str)
	return
}

func MustString(any any) (resp string) {
	if any == nil {
		return ""
	}

	switch val := any.(type) {
	case string:
		resp = val
	case []byte:
		resp = string(val)
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64:
		resp = fmt.Sprintf("%d", val)
	case float32:
		resp = strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		resp = strconv.FormatFloat(val, 'f', -1, 64)
	default:
		resp = fmt.Sprintf("%v", val)
	}
	return
}

func MustInt(any any) (resp int) {
	if any == nil {
		return 0
	}
	switch val := any.(type) {
	case string:
		resp, _ = strconv.Atoi(val)
	case int, int16, int32, int64, float32, float64:
		resp, _ = fmt.Printf("%d", val)
	case bool:
		if val {
			resp = 1
		} else {
			resp = 0
		}
	default:
		resp, _ = fmt.Printf("%d", val)
	}
	return
}

func MustInt32(any any) (resp int32) {
	if any == nil {
		return 0
	}
	switch val := any.(type) {
	case string:
		respInt64, _ := strconv.ParseInt(val, 10, 64)
		resp = int32(respInt64)
	case int, int16, int32, int64, float32, float64:
		respInt, _ := fmt.Printf("%d", val)
		resp = int32(respInt)
	case bool:
		if val {
			resp = 1
		} else {
			resp = 0
		}
	default:
		respInt, _ := fmt.Printf("%d", val)
		resp = int32(respInt)
	}
	return
}

func MustInt64(any any) (resp int64) {
	if any == nil {
		return 0
	}
	switch val := any.(type) {
	case string:
		resp, _ = strconv.ParseInt(val, 10, 64)
	case int, int16, int32, int64, float32, float64:
		respInt, _ := fmt.Printf("%d", val)
		resp = int64(respInt)
	case bool:
		if val {
			resp = 1
		} else {
			resp = 0
		}
	default:
		respInt, _ := fmt.Printf("%d", val)
		resp = int64(respInt)
	}
	return
}

func MustFloat32(any any) (resp float32) {
	if any == nil {
		return 0
	}
	switch val := any.(type) {
	case string:
		respFloat64, _ := strconv.ParseFloat(val, 32)
		resp = float32(respFloat64)
	case int, int16, int32, int64, float32, float64:
		respInt, _ := fmt.Printf("%d", val)
		resp = float32(respInt)
	case bool:
		if val {
			resp = 1
		} else {
			resp = 0
		}
	default:
		respInt, _ := fmt.Printf("%d", val)
		resp = float32(respInt)
	}
	return
}

func MustFloat64(any any) (resp float64) {
	if any == nil {
		return 0
	}
	switch val := any.(type) {
	case string:
		resp, _ = strconv.ParseFloat(val, 32)
	case int, int16, int32, int64, float32, float64:
		respInt, _ := fmt.Printf("%d", val)
		resp = float64(respInt)
	case bool:
		if val {
			resp = 1
		} else {
			resp = 0
		}
	default:
		respInt, _ := fmt.Printf("%d", val)
		resp = float64(respInt)
	}
	return
}

func MustBool(any any) bool {
	if any == nil {
		return false
	}
	resp, ok := any.(bool)
	if !ok {
		str := fmt.Sprintf("%v", any)
		if str == "true" || str == "True" {
			return true
		}
		return false
	}
	return resp
}
