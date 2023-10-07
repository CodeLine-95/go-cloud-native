// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package logz

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"sync"
	"time"
	"unicode/utf8"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// For JSON-escaping; see JsonEncoder.safeAddString below.
const hex = "0123456789abcdef"

var (
	pool = buffer.NewPool()
)

var jsonPool = sync.Pool{New: func() interface{} {
	return &JsonEncoder{}
}}

func addFields(enc zapcore.ObjectEncoder, fields []zapcore.Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}

func getJSONEncoder() *JsonEncoder {
	return jsonPool.Get().(*JsonEncoder)
}

func putJSONEncoder(enc *JsonEncoder) {
	if enc.reflectBuf != nil {
		enc.reflectBuf.Free()
	}
	enc.EncoderConfig = nil
	enc.Buf = nil
	enc.spaced = false
	enc.openNamespaces = 0
	enc.reflectBuf = nil
	enc.reflectEnc = nil
	jsonPool.Put(enc)
}

type EncoderConfig struct {
	zapcore.EncoderConfig
	PrefixFns []PrefixFn
}

type JsonEncoder struct {
	*zapcore.EncoderConfig
	prefixes       *[]PrefixFn
	Buf            *buffer.Buffer
	spaced         bool // include spaces after colons and commas
	openNamespaces int

	// for encoding generic values by reflection
	reflectBuf *buffer.Buffer
	reflectEnc *json.Encoder
}

// NewJSONEncoder creates a fast, low-allocation JSON encoder. The encoder
// appropriately escapes all field keys and values.
//
// Note that the encoder doesn't deduplicate keys, so it's possible to produce
// a message like
//
//	{"foo":"bar","foo":"baz"}
//
// This is permitted by the JSON specification, but not encouraged. Many
// libraries will ignore duplicate key-value pairs (typically keeping the last
// pair) when unmarshaling, but users should attempt to avoid adding duplicate
// keys.
func NewJSONEncoder(cfg EncoderConfig) zapcore.Encoder {
	return newJSONEncoder(cfg, false)
}

func newJSONEncoder(cfg EncoderConfig, spaced bool) *JsonEncoder {
	return &JsonEncoder{
		EncoderConfig: &cfg.EncoderConfig,
		prefixes:      &cfg.PrefixFns,
		Buf:           pool.Get(),
		spaced:        spaced,
	}
}

func (enc *JsonEncoder) AddArray(key string, arr zapcore.ArrayMarshaler) error {
	enc.addKey(key)
	return enc.AppendArray(arr)
}

func (enc *JsonEncoder) AddObject(key string, obj zapcore.ObjectMarshaler) error {
	enc.addKey(key)
	return enc.AppendObject(obj)
}

func (enc *JsonEncoder) AddBinary(key string, val []byte) {
	enc.AddString(key, base64.StdEncoding.EncodeToString(val))
}

func (enc *JsonEncoder) AddByteString(key string, val []byte) {
	enc.addKey(key)
	enc.AppendByteString(val)
}

func (enc *JsonEncoder) AddBool(key string, val bool) {
	enc.addKey(key)
	enc.AppendBool(val)
}

func (enc *JsonEncoder) AddComplex128(key string, val complex128) {
	enc.addKey(key)
	enc.AppendComplex128(val)
}

func (enc *JsonEncoder) AddDuration(key string, val time.Duration) {
	enc.addKey(key)
	enc.AppendDuration(val)
}

func (enc *JsonEncoder) AddFloat64(key string, val float64) {
	enc.addKey(key)
	enc.AppendFloat64(val)
}

func (enc *JsonEncoder) AddInt64(key string, val int64) {
	enc.addKey(key)
	enc.AppendInt64(val)
}

func (enc *JsonEncoder) resetReflectBuf() {
	if enc.reflectBuf == nil {
		enc.reflectBuf = pool.Get()
		enc.reflectEnc = json.NewEncoder(enc.reflectBuf)
	} else {
		enc.reflectBuf.Reset()
	}
}

func (enc *JsonEncoder) AddReflected(key string, obj interface{}) error {
	enc.resetReflectBuf()
	err := enc.reflectEnc.Encode(obj)
	if err != nil {
		return err
	}
	enc.reflectBuf.TrimNewline()
	enc.addKey(key)
	_, err = enc.Buf.Write(enc.reflectBuf.Bytes())
	return err
}

func (enc *JsonEncoder) OpenNamespace(key string) {
	enc.addKey(key)
	enc.Buf.AppendByte('{')
	enc.openNamespaces++
}

func (enc *JsonEncoder) AddString(key, val string) {
	enc.addKey(key)
	enc.AppendString(val)
}

func (enc *JsonEncoder) AddTime(key string, val time.Time) {
	enc.addKey(key)
	enc.AppendTime(val)
}

func (enc *JsonEncoder) AddUint64(key string, val uint64) {
	enc.addKey(key)
	enc.AppendUint64(val)
}

func (enc *JsonEncoder) AppendArray(arr zapcore.ArrayMarshaler) error {
	enc.addElementSeparator()
	enc.Buf.AppendByte('[')
	err := arr.MarshalLogArray(enc)
	enc.Buf.AppendByte(']')
	return err
}

func (enc *JsonEncoder) AppendObject(obj zapcore.ObjectMarshaler) error {
	enc.addElementSeparator()
	enc.Buf.AppendByte('{')
	err := obj.MarshalLogObject(enc)
	enc.Buf.AppendByte('}')
	return err
}

func (enc *JsonEncoder) AppendBool(val bool) {
	enc.addElementSeparator()
	enc.Buf.AppendBool(val)
}

func (enc *JsonEncoder) AppendByteString(val []byte) {
	enc.addElementSeparator()
	enc.Buf.AppendByte('"')
	enc.safeAddByteString(val)
	enc.Buf.AppendByte('"')
}

func (enc *JsonEncoder) AppendComplex128(val complex128) {
	enc.addElementSeparator()
	// Cast to a platform-independent, fixed-size type.
	r, i := float64(real(val)), float64(imag(val))
	enc.Buf.AppendByte('"')
	// Because we're always in a quoted string, we can use strconv without
	// special-casing NaN and +/-Inf.
	enc.Buf.AppendFloat(r, 64)
	enc.Buf.AppendByte('+')
	enc.Buf.AppendFloat(i, 64)
	enc.Buf.AppendByte('i')
	enc.Buf.AppendByte('"')
}

func (enc *JsonEncoder) AppendDuration(val time.Duration) {
	cur := enc.Buf.Len()
	enc.EncodeDuration(val, enc)
	if cur == enc.Buf.Len() {
		// User-supplied EncodeDuration is a no-op. Fall back to nanoseconds to keep
		// JSON valid.
		enc.AppendInt64(int64(val))
	}
}

func (enc *JsonEncoder) AppendInt64(val int64) {
	enc.addElementSeparator()
	enc.Buf.AppendInt(val)
}

func (enc *JsonEncoder) AppendReflected(val interface{}) error {
	enc.resetReflectBuf()
	err := enc.reflectEnc.Encode(val)
	if err != nil {
		return err
	}
	enc.reflectBuf.TrimNewline()
	enc.addElementSeparator()
	_, err = enc.Buf.Write(enc.reflectBuf.Bytes())
	return err
}

func (enc *JsonEncoder) AppendString(val string) {
	enc.addElementSeparator()
	enc.Buf.AppendByte('"')
	enc.safeAddString(val)
	enc.Buf.AppendByte('"')
}

func (enc *JsonEncoder) AppendTime(val time.Time) {
	cur := enc.Buf.Len()
	enc.EncodeTime(val, enc)
	if cur == enc.Buf.Len() {
		// User-supplied EncodeTime is a no-op. Fall back to nanos since epoch to keep
		// output JSON valid.
		enc.AppendInt64(val.UnixNano())
	}
}

func (enc *JsonEncoder) AppendUint64(val uint64) {
	enc.addElementSeparator()
	enc.Buf.AppendUint(val)
}

func (enc *JsonEncoder) AddComplex64(k string, v complex64) { enc.AddComplex128(k, complex128(v)) }
func (enc *JsonEncoder) AddFloat32(k string, v float32)     { enc.AddFloat64(k, float64(v)) }
func (enc *JsonEncoder) AddInt(k string, v int)             { enc.AddInt64(k, int64(v)) }
func (enc *JsonEncoder) AddInt32(k string, v int32)         { enc.AddInt64(k, int64(v)) }
func (enc *JsonEncoder) AddInt16(k string, v int16)         { enc.AddInt64(k, int64(v)) }
func (enc *JsonEncoder) AddInt8(k string, v int8)           { enc.AddInt64(k, int64(v)) }
func (enc *JsonEncoder) AddUint(k string, v uint)           { enc.AddUint64(k, uint64(v)) }
func (enc *JsonEncoder) AddUint32(k string, v uint32)       { enc.AddUint64(k, uint64(v)) }
func (enc *JsonEncoder) AddUint16(k string, v uint16)       { enc.AddUint64(k, uint64(v)) }
func (enc *JsonEncoder) AddUint8(k string, v uint8)         { enc.AddUint64(k, uint64(v)) }
func (enc *JsonEncoder) AddUintptr(k string, v uintptr)     { enc.AddUint64(k, uint64(v)) }
func (enc *JsonEncoder) AppendComplex64(v complex64)        { enc.AppendComplex128(complex128(v)) }
func (enc *JsonEncoder) AppendFloat64(v float64)            { enc.appendFloat(v, 64) }
func (enc *JsonEncoder) AppendFloat32(v float32)            { enc.appendFloat(float64(v), 32) }
func (enc *JsonEncoder) AppendInt(v int)                    { enc.AppendInt64(int64(v)) }
func (enc *JsonEncoder) AppendInt32(v int32)                { enc.AppendInt64(int64(v)) }
func (enc *JsonEncoder) AppendInt16(v int16)                { enc.AppendInt64(int64(v)) }
func (enc *JsonEncoder) AppendInt8(v int8)                  { enc.AppendInt64(int64(v)) }
func (enc *JsonEncoder) AppendUint(v uint)                  { enc.AppendUint64(uint64(v)) }
func (enc *JsonEncoder) AppendUint32(v uint32)              { enc.AppendUint64(uint64(v)) }
func (enc *JsonEncoder) AppendUint16(v uint16)              { enc.AppendUint64(uint64(v)) }
func (enc *JsonEncoder) AppendUint8(v uint8)                { enc.AppendUint64(uint64(v)) }
func (enc *JsonEncoder) AppendUintptr(v uintptr)            { enc.AppendUint64(uint64(v)) }

func (enc *JsonEncoder) Clone() zapcore.Encoder {
	clone := enc.clone()
	clone.Buf.Write(enc.Buf.Bytes())
	return clone
}

func (enc *JsonEncoder) clone() *JsonEncoder {
	clone := getJSONEncoder()
	clone.EncoderConfig = enc.EncoderConfig
	clone.spaced = enc.spaced
	clone.openNamespaces = enc.openNamespaces
	clone.Buf = pool.Get()
	clone.prefixes = enc.prefixes
	return clone
}

func (enc *JsonEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	final := enc.clone()

	for _, p := range *enc.prefixes {
		p(final, ent)
	}
	final.Buf.AppendByte('{')

	if final.LevelKey != "" {
		final.addKey(final.LevelKey)
		cur := final.Buf.Len()
		final.EncodeLevel(ent.Level, final)
		if cur == final.Buf.Len() {
			// User-supplied EncodeLevel was a no-op. Fall back to strings to keep
			// output JSON valid.
			final.AppendString(ent.Level.String())
		}
	}
	if final.TimeKey != "" {
		final.AddTime(final.TimeKey, ent.Time)
	}
	if ent.LoggerName != "" && final.NameKey != "" {
		final.addKey(final.NameKey)
		cur := final.Buf.Len()
		nameEncoder := final.EncodeName

		// if no name encoder provided, fall back to FullNameEncoder for backwards
		// compatibility
		if nameEncoder == nil {
			nameEncoder = zapcore.FullNameEncoder
		}

		nameEncoder(ent.LoggerName, final)
		if cur == final.Buf.Len() {
			// User-supplied EncodeName was a no-op. Fall back to strings to
			// keep output JSON valid.
			final.AppendString(ent.LoggerName)
		}
	}
	if ent.Caller.Defined && final.CallerKey != "" {
		final.addKey(final.CallerKey)
		cur := final.Buf.Len()
		final.EncodeCaller(ent.Caller, final)
		if cur == final.Buf.Len() {
			// User-supplied EncodeCaller was a no-op. Fall back to strings to
			// keep output JSON valid.
			final.AppendString(ent.Caller.String())
		}
	}
	if final.MessageKey != "" {
		final.addKey(enc.MessageKey)
		final.AppendString(ent.Message)
	}
	if enc.Buf.Len() > 0 {
		final.addElementSeparator()
		final.Buf.Write(enc.Buf.Bytes())
	}
	addFields(final, fields)
	final.closeOpenNamespaces()
	if ent.Stack != "" && final.StacktraceKey != "" {
		final.AddString(final.StacktraceKey, ent.Stack)
	}
	final.Buf.AppendByte('}')
	if final.LineEnding != "" {
		final.Buf.AppendString(final.LineEnding)
	} else {
		final.Buf.AppendString(zapcore.DefaultLineEnding)
	}

	ret := final.Buf
	putJSONEncoder(final)
	return ret, nil
}

func (enc *JsonEncoder) truncate() {
	enc.Buf.Reset()
}

func (enc *JsonEncoder) closeOpenNamespaces() {
	for i := 0; i < enc.openNamespaces; i++ {
		enc.Buf.AppendByte('}')
	}
}

func (enc *JsonEncoder) addKey(key string) {
	enc.addElementSeparator()
	enc.Buf.AppendByte('"')
	enc.safeAddString(key)
	enc.Buf.AppendByte('"')
	enc.Buf.AppendByte(':')
	if enc.spaced {
		enc.Buf.AppendByte(' ')
	}
}

func (enc *JsonEncoder) addElementSeparator() {
	last := enc.Buf.Len() - 1
	if last < 0 {
		return
	}
	switch enc.Buf.Bytes()[last] {
	case '{', '[', ':', ',', ' ':
		return
	default:
		enc.Buf.AppendByte(',')
		if enc.spaced {
			enc.Buf.AppendByte(' ')
		}
	}
}

func (enc *JsonEncoder) appendFloat(val float64, bitSize int) {
	enc.addElementSeparator()
	switch {
	case math.IsNaN(val):
		enc.Buf.AppendString(`"NaN"`)
	case math.IsInf(val, 1):
		enc.Buf.AppendString(`"+Inf"`)
	case math.IsInf(val, -1):
		enc.Buf.AppendString(`"-Inf"`)
	default:
		enc.Buf.AppendFloat(val, bitSize)
	}
}

// safeAddString JSON-escapes a string and appends it to the internal buffer.
// Unlike the standard library's encoder, it doesn't attempt to protect the
// user from browser vulnerabilities or JSONP-related problems.
func (enc *JsonEncoder) safeAddString(s string) {
	for i := 0; i < len(s); {
		if enc.tryAddRuneSelf(s[i]) {
			i++
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if enc.tryAddRuneError(r, size) {
			i++
			continue
		}
		enc.Buf.AppendString(s[i : i+size])
		i += size
	}
}

// safeAddByteString is no-alloc equivalent of safeAddString(string(s)) for s []byte.
func (enc *JsonEncoder) safeAddByteString(s []byte) {
	for i := 0; i < len(s); {
		if enc.tryAddRuneSelf(s[i]) {
			i++
			continue
		}
		r, size := utf8.DecodeRune(s[i:])
		if enc.tryAddRuneError(r, size) {
			i++
			continue
		}
		enc.Buf.Write(s[i : i+size])
		i += size
	}
}

// tryAddRuneSelf appends b if it is valid UTF-8 character represented in a single byte.
func (enc *JsonEncoder) tryAddRuneSelf(b byte) bool {
	if b >= utf8.RuneSelf {
		return false
	}
	if 0x20 <= b && b != '\\' && b != '"' {
		enc.Buf.AppendByte(b)
		return true
	}
	switch b {
	case '\\', '"':
		enc.Buf.AppendByte('\\')
		enc.Buf.AppendByte(b)
	case '\n':
		enc.Buf.AppendByte('\\')
		enc.Buf.AppendByte('n')
	case '\r':
		enc.Buf.AppendByte('\\')
		enc.Buf.AppendByte('r')
	case '\t':
		enc.Buf.AppendByte('\\')
		enc.Buf.AppendByte('t')
	default:
		// Encode bytes < 0x20, except for the escape sequences above.
		enc.Buf.AppendString(`\u00`)
		enc.Buf.AppendByte(hex[b>>4])
		enc.Buf.AppendByte(hex[b&0xF])
	}
	return true
}

func (enc *JsonEncoder) tryAddRuneError(r rune, size int) bool {
	if r == utf8.RuneError && size == 1 {
		enc.Buf.AppendString(`\ufffd`)
		return true
	}
	return false
}
