package logz

import (
	"go.uber.org/zap/zapcore"
)

type PrefixFn func(encoder *JsonEncoder, ent zapcore.Entry)
