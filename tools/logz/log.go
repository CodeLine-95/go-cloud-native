package logz

import (
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// defaultLogger 全局默认日志实例
var defaultLogger *Logger

type (
	// Logger 对zap.Logger进行了封装
	Logger struct {
		*zap.Logger
		atom          zap.AtomicLevel
		writer        writer
		level         level
		fields        []field
		prefixes      []PrefixFn // 日志记录前统一前缀
		encoderConfig *encoderConfig
	}

	// field 定义了一个私有的zap.Field类型别名，防止用户直接操作zap.Field
	// 用户可用过F(key, val)方法返回一个field
	field = zapcore.Field
)

func init() {
	SetDefaultLogger(New())
}

// New 构造logger实例
func New(ops ...Option) *Logger {
	// 构造Logger
	l := &Logger{
		atom:   zap.NewAtomicLevel(),
		Logger: new(zap.Logger),
		level:  InfoLevel, // 默认级别为info
	}

	fmtTimeMtx.Lock()
	fmtTime = time.Now().Format("2006-01-02 15:04:05")
	fmtTimeMtx.Unlock()
	go SyncTime()

	for _, op := range ops {
		op.apply(l)
	}

	zapOpts := []zap.Option{zap.AddStacktrace(ErrorLevel)}
	// 初始化fields
	if len(l.fields) != 0 {
		zapOpts = append(zapOpts, zap.Fields(l.fields...))
	}

	// 初始化level
	l.atom.SetLevel(l.level)

	// 初始化writer
	// 如果没有配置writer，默认设为stderr writer
	if l.writer == nil {
		l.writer = os.Stderr
	}

	// 初始化encoderConfig
	if l.encoderConfig == nil {
		cfg := zap.NewProductionEncoderConfig()
		l.encoderConfig = &cfg
	}

	// 直接在stderr write的时候加前缀有较大的性能损耗，因此不采用
	core := zapcore.NewCore(
		NewJSONEncoder(EncoderConfig{
			EncoderConfig: *l.encoderConfig,
			PrefixFns:     l.prefixes,
		}),
		l.writer,
		l.atom,
	)

	l.Logger.WithOptions()
	l.Logger = zap.New(core, zapOpts...)

	return l
}

// SetDefaultLogger 根据输入的Options初始化默认logger实例
func SetDefaultLogger(l *Logger) {
	defaultLogger = l
}

// DefaultLogger 返回默认logger实例
func DefaultLogger() *Logger {
	return defaultLogger
}

// Handler 返回atom httpHandler
func Handler() http.Handler {
	return defaultLogger.atom
}

// Handler 返回defaultLogger的atom httpHandler
func (l *Logger) Handler() http.Handler {
	return l.atom
}

func write(level zapcore.Level, msg string, fields ...field) {
	if ce := defaultLogger.Check(level, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (l *Logger) write(level zapcore.Level, msg string, fields ...field) {
	if ce := l.Check(level, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Debug 使用defaultLogger实例打印Debug级别日志
func Debug(msg string, fields ...field) {
	write(DebugLevel, msg, fields...)
}

// Debug 打印Debug级别日志
func (l *Logger) Debug(msg string, fields ...field) {
	l.write(DebugLevel, msg, fields...)
}

// Info 使用defaultLogger实例打印Info级别日志
func Info(msg string, fields ...field) {
	write(InfoLevel, msg, fields...)
}

// Info 打印Info级别日志
func (l *Logger) Info(msg string, fields ...field) {
	l.write(InfoLevel, msg, fields...)
}

// Warn 使用defaultLogger实例打印Warn级别日志
func Warn(msg string, fields ...field) {
	write(WarnLevel, msg, fields...)
}

// Warn 打印Warn级别日志
func (l *Logger) Warn(msg string, fields ...field) {
	l.write(WarnLevel, msg, fields...)
}

// Error 使用defaultLogger实例打印Error级别日志。注：此级别下会打印调用栈，请根据需求使用
func Error(msg string, fields ...field) {
	write(ErrorLevel, msg, fields...)
}

// Error 打印Error级别日志。注：此级别下会打印调用栈，请根据需求使用
func (l *Logger) Error(msg string, fields ...field) {
	l.write(ErrorLevel, msg, fields...)
}

// Panic 使用defaultLogger实例打印Panic级别日志。注：此级别下会打印调用栈，同时会Panic，请根据需求使用
func Panic(msg string, fields ...field) {
	write(PanicLevel, msg, fields...)
}

// Panic 打印Panic级别日志。注：此级别下会打印调用栈，同时会Panic，请根据需求使用
func (l *Logger) Panic(msg string, fields ...field) {
	l.write(PanicLevel, msg, fields...)
}

// Fatal 使用defaultLogger实例打印Fatal级别日志。注：此级别下会打印调用栈，同时会调用os.Exit(1)退出，请根据需求使用
func Fatal(msg string, fields ...field) {
	write(FatalLevel, msg, fields...)
}

// Fatal 打印Fatal级别日志。注：此级别下会打印调用栈，同时会调用os.Exit(1)退出，请根据需求使用
func (l *Logger) Fatal(msg string, fields ...field) {
	l.write(FatalLevel, msg, fields...)
}
