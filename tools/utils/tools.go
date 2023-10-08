package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// MD5 generate md5 by string
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Get16BitStrMd5(str string) string {
	return MD5(str)[8:24]
}

func Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func Base64Decode(input string) string {
	ret, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return ""
	}
	return string(ret)
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

func NewMkdir(path string) string {
	floderName := time.Now().Format(time.DateOnly)
	floderPath := filepath.Join(path, floderName)
	err := os.MkdirAll(floderPath, os.ModePerm)
	if err != nil {
		return ""
	}
	return floderPath
}

// RemoveHtmlTag 去除代码中的html标签
func RemoveHtmlTag(rawHtml string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		return "", err
	}
	htmlString := doc.Text()
	return htmlString, nil
}

// NumToAsc 暂时仅支持大小写字母
func NumToAsc(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	var result string
	for _, v := range nums {
		result += string(rune(v))

	}
	return result
}

func RandStringRunes(n int) string {
	rand.NewSource(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// DistinctI64 -
// 引入泛型前，为了类型安全暂时包一下 RemoveDuplicateElement
func DistinctI64(v []int64) []int64 {
	res, _ := RemoveDuplicateElement(v)
	return res.([]int64)
}

// DistinctStr -
// 引入泛型前，为了类型安全暂时包一下 RemoveDuplicateElement
func DistinctStr(v []string) []string {
	res, _ := RemoveDuplicateElement(v)
	return res.([]string)
}

// RemoveDuplicateElement slice去重
func RemoveDuplicateElement(ori any) (any, error) {
	temp := map[any]struct{}{}

	switch sType := ori.(type) {
	case []string:
		result := make([]string, 0, len(ori.([]string)))

		for _, item := range sType {
			if _, ok := temp[item]; !ok {
				temp[item] = struct{}{}
				result = append(result, item)
			}
		}

		return result, nil
	case []int64:
		result := make([]int64, 0, len(ori.([]int64)))

		for _, item := range sType {
			if _, ok := temp[item]; !ok {
				temp[item] = struct{}{}
				result = append(result, item)
			}
		}
		return result, nil
	default:
		err := fmt.Errorf("unknown type: %T", sType)
		return nil, err
	}
}

func InArray(src string, arr []string) bool {
	for _, v := range arr {
		if src == v {
			return true
		}
	}
	return false
}

func InArrayInt(d int, arr []int) bool {
	for _, v := range arr {
		if d == v {
			return true
		}
	}
	return false
}
