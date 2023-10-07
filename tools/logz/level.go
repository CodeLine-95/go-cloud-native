package logz

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// DebugLevel 打印Debug级别日志
	DebugLevel level = zap.DebugLevel
	// InfoLevel 打印Info级别日志
	InfoLevel level = zap.InfoLevel
	// WarnLevel 打印Warn级别日志
	WarnLevel level = zap.WarnLevel
	// ErrorLevel 打印Error级别日志和调用栈
	ErrorLevel level = zap.ErrorLevel
	// PanicLevel 打印Panic级别日志和调用栈并panic退出
	PanicLevel level = zap.PanicLevel
	// FatalLevel 打印Fatal级别日志和调用栈并调用os.Exit(1)退出
	FatalLevel level = zap.FatalLevel
)

// Lvl 定义了一个私有的zap.Level类型别名
type level = zapcore.Level

var LevelM = map[level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "NOTICE",
	WarnLevel:  "WARN",
	ErrorLevel: "FATAL",
	PanicLevel: "FATAL",
	FatalLevel: "FATAL",
}
