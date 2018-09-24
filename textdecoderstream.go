// Code generated DO NOT EDIT
// textdecoderstream.go
package dom

import "syscall/js"

type TextDecoderStreamIFace interface {
	GetEncoding() string
	GetFatal() bool
	GetIgnoreBOM() bool
	GetReadable() Value
	GetWritable() Value
}
type TextDecoderStream struct {
	Value
}

func JSValueToTextDecoderStream(val js.Value) TextDecoderStream {
	return TextDecoderStream{Value: JSValueToValue(val)}
}
func (v Value) AsTextDecoderStream() TextDecoderStream { return TextDecoderStream{Value: v} }
func NewTextDecoderStream(args ...interface{}) TextDecoderStream {
	return TextDecoderStream{Value: JSValueToValue(js.Global().Get("TextDecoderStream").New(args...))}
}
func (t TextDecoderStream) GetEncoding() string {
	val := t.Get("encoding")
	return val.String()
}
func (t TextDecoderStream) GetFatal() bool {
	val := t.Get("fatal")
	return val.Bool()
}
func (t TextDecoderStream) GetIgnoreBOM() bool {
	val := t.Get("ignoreBOM")
	return val.Bool()
}
func (t TextDecoderStream) GetReadable() Value {
	val := t.Get("readable")
	return val
}
func (t TextDecoderStream) GetWritable() Value {
	val := t.Get("writable")
	return val
}
