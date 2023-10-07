package logz

import (
	"log"

	"go.uber.org/zap/zapcore"
)

type encoderConfig = zapcore.EncoderConfig

// An Option configures a Logger.
type Option interface {
	apply(*Logger)
}

// optionFunc wraps a func, so it satisfies the Option interface.
type optionFunc func(*Logger)

func (f optionFunc) apply(l *Logger) {
	f(l)
}

// Writer 为实例设置默认writer
func Writer(w writer) Option {
	return optionFunc(func(l *Logger) {
		l.writer = w
	})
}

// Level 为实例设置默认日志级别
func Level(lvl level) Option {
	return optionFunc(func(l *Logger) {
		l.level = lvl
	})
}

// LevelStr 为实例设置默认日志级别，传入的参数是string
func LevelStr(lvlStr string) Option {
	return optionFunc(func(l *Logger) {
		var lvl level
		if err := lvl.UnmarshalText([]byte(lvlStr)); err != nil {
			log.Printf("invalid level string: %s, use InfoLevel instead!\n", lvlStr)
			lvl = InfoLevel
		}
		l.level = lvl
	})
}

// Fields 为实例添加染色字段
func Fields(fields ...field) Option {
	return optionFunc(func(l *Logger) {
		l.fields = fields
	})
}

// Prefix 设置输出前缀
func Prefix(p ...string) Option {
	return optionFunc(func(l *Logger) {
		// 添加 [时间] 前缀
		l.prefixes = append(l.prefixes, func(encoder *JsonEncoder, ent zapcore.Entry) {
			encoder.Buf.AppendByte('[')
			encoder.Buf.AppendString(FmtTime())
			encoder.Buf.AppendString("] ")
		})
		// 添加 [服务名] [Hostname] 前缀
		for i := range p {
			v := p[i]
			l.prefixes = append(l.prefixes, func(encoder *JsonEncoder, ent zapcore.Entry) {
				encoder.Buf.AppendByte('[')
				encoder.Buf.AppendString(v)
				encoder.Buf.AppendString("] ")
			})
		}
		// 添加 [日志级别] 前缀
		l.prefixes = append(l.prefixes, func(encoder *JsonEncoder, ent zapcore.Entry) {
			encoder.Buf.AppendByte('[')
			encoder.Buf.AppendString(LevelM[ent.Level])
			encoder.Buf.AppendString("] ")
		})
		// 添加 [logID] 前缀
		l.prefixes = append(l.prefixes, func(encoder *JsonEncoder, ent zapcore.Entry) {
			encoder.Buf.AppendString("[] ")
		})
	})
}
