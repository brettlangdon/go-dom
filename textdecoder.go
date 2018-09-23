// Code generated DO NOT EDIT
// textdecoder.go
package dom

import "syscall/js"

type TextDecoderIFace interface {
	Decode(args ...interface{}) string
	GetEncoding() string
	GetFatal() bool
	GetIgnoreBOM() bool
}
type TextDecoder struct {
	Value
}

func jsValueToTextDecoder(val js.Value) TextDecoder { return TextDecoder{Value: Value{Value: val}} }
func (v Value) AsTextDecoder() TextDecoder          { return TextDecoder{Value: v} }
func (t TextDecoder) Decode(args ...interface{}) string {
	val := t.Call("decode", args...)
	return val.String()
}
func (t TextDecoder) GetEncoding() string {
	val := t.Get("encoding")
	return val.String()
}
func (t TextDecoder) GetFatal() bool {
	val := t.Get("fatal")
	return val.Bool()
}
func (t TextDecoder) GetIgnoreBOM() bool {
	val := t.Get("ignoreBOM")
	return val.Bool()
}
