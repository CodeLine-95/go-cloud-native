package xlog

import (
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"go.uber.org/zap/zapcore"
	"strings"
	"time"
)

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	PanicLevel = "panic"
	FatalLevel = "fatal"
)

var levelMap = map[string]zapcore.Level{
	DebugLevel: zapcore.DebugLevel,
	InfoLevel:  zapcore.InfoLevel,
	WarnLevel:  zapcore.WarnLevel,
	ErrorLevel: zapcore.ErrorLevel,
	PanicLevel: zapcore.PanicLevel,
	FatalLevel: zapcore.FatalLevel,
}

type LogContext interface {
	Formatter() (string, []zapcore.Field)
}

type log struct {
	errLogger     *logz.Logger
	defaultLogger *logz.Logger
}

var logger *log

func InitLog(dir string, level string, serviceName string) {
	logger = &log{}
	if dir == "" {
		logger.defaultLogger = logz.DefaultLogger()
		logger.errLogger = logz.DefaultLogger()
		return
	}
	l, ok := levelMap[strings.ToLower(level)]
	if !ok {
		l = zapcore.InfoLevel
	}
	logFile := dir + "/" + serviceName + "/" + time.Now().Format(time.DateOnly)
	defaultLogger := logz.New(logz.Writer(logz.NewFileWriter(logFile+".log")), logz.Level(l))
	logz.SetDefaultLogger(defaultLogger)
	errLogger := logz.New(logz.Writer(logz.NewFileWriter(logFile+".wf")), logz.Level(levelMap[WarnLevel]))
	logger.defaultLogger = defaultLogger
	logger.errLogger = errLogger
	logz.Info("init logger succ")
}

func Debug(ctx LogContext) {
	msg, fields := ctx.Formatter()
	logger.defaultLogger.Debug(msg, fields...)
}

func Info(ctx LogContext) {
	msg, fields := ctx.Formatter()
	logger.defaultLogger.Info(msg, fields...)
}

func Warn(ctx LogContext) {
	msg, fields := ctx.Formatter()
	logger.errLogger.Warn(msg, fields...)
}

func Error(ctx LogContext) {
	msg, fields := ctx.Formatter()
	logger.errLogger.Error(msg, fields...)
}

func Panic(ctx LogContext) {
	msg, fields := ctx.Formatter()
	logger.errLogger.Panic(msg, fields...)
}

func Fatal(ctx LogContext) {
	msg, fields := ctx.Formatter()
	logger.errLogger.Fatal(msg, fields...)
}
