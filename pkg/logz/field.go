package logz

import (
	"go.uber.org/zap"
)

// F 返回一个指定key的field
func F(key string, value interface{}) field {
	return zap.Any(key, value)
}

// E 是F("error", err)的语法糖，用来构造错误err的filed，其key为"error"
func E(err error) field {
	return zap.Error(err)
}
