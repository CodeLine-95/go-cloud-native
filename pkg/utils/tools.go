package utils

import (
	"crypto/md5"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// MD5 generate md5 by string
func MD5(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}

// BatchSearchToStringSlice 批量搜索对字符串的分割
func BatchSearchToStringSlice(text string) []string {
	var nilSlice []string
	if text == "" {
		return nilSlice
	}
	res := strings.Split(text, ",")
	if len(res) > 1 {
		return res
	}

	res = strings.Split(text, " ")
	if len(res) > 1 {
		return res
	}

	return []string{text}
}

// BatchSearchToInt64Slice 批量搜索对字符串的分割
func BatchSearchToInt64Slice(text string) []int64 {
	slice := BatchSearchToStringSlice(text)
	var tmp []int64
	for _, id := range slice {
		id = strings.Trim(id, " ")
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return []int64{}
		}
		tmp = append(tmp, intId)
	}
	return tmp
}

func SubStr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	return s[:l]
}

// Round 四舍五入 precision:需要的保留的小数位数
func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}
