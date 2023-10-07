package logz

import (
	"os"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type writer = zapcore.WriteSyncer

type fileWriter struct {
	*lumberjack.Logger
}

func (fileWriter) Sync() error { return nil }

const (
	defaultFileMaxSiz     = 500 // 默认最大文件大小 MB
	defaultFileMaxAge     = 1   // 默认文件分隔周期 days
	defaultFileMaxBackups = 100 // 默认最大旧日志备份数
)

// NewFileWriter 返回写文件的writer
func NewFileWriter(path string) writer {
	return &fileWriter{
		&lumberjack.Logger{
			Filename:   path,
			MaxSize:    defaultFileMaxSiz,
			MaxAge:     defaultFileMaxAge,
			MaxBackups: defaultFileMaxBackups,
			LocalTime:  true,
			Compress:   true,
		},
	}
}

// NewStdErrWriter 返回写标准错误输出的writer
func NewStdErrWriter() writer {
	return zapcore.Lock(os.Stderr)
}
