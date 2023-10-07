package logz

import (
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var messages = fakeMessages(1000)

func init() {
	hst, _ := os.Hostname()
	SetDefaultLogger(New(
		Fields(
			F("SERVICE", "logz"),
			F("GROUP", "42"),
		),
		Level(DebugLevel),
		Prefix("li_web_api", hst),
	))
}

func fakeMessages(n int) []string {
	messages := make([]string, n)
	for i := range messages {
		messages[i] = fmt.Sprintf("Test logging, but use a somewhat realistic message length. (#%v)", i)
	}
	return messages
}

func getMessage(iter int) string {
	return messages[iter%1000]
}

func TestDebug(t *testing.T) {
	Debug("用户lee的相关信息", F("Name", "lee"), F("Age", 42), F("Country", "CN"))
	// TODO 完善测试用例
}

func TestFileWriter(t *testing.T) {
	// 手动构造logger实例
	logger := New(
		Writer(NewFileWriter("test.log")),
		Level(DebugLevel),
	)
	logger.Info("用户lee的相关信息", F("Name", "lee"), F("Age", 42), F("Country", "CN"))

	// 测试SetDefaultLogger方法，将logger实例设置为默认实例
	SetDefaultLogger(logger)
	Error("用户lee的相关信息", F("Name", "lee"), F("Age", 42), F("Country", "CN"))
	// TODO 完善测试用例
}

func BenchmarkLogzInfo(b *testing.B) {
	b.Logf("Logging at a disabled level with some accumulated context.")
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)
	zapLogger := zap.New(zapcore.NewCore(
		enc,
		zapcore.Lock(os.Stderr),
		zap.DebugLevel,
	))

	b.Run("zap", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				zapLogger.Info(getMessage(0))
			}
		})
	})
	b.Run("logz", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Info(getMessage(0))
			}
		})
	})
}
