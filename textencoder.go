// Code generated DO NOT EDIT
// textencoder.go
package dom

import "syscall/js"

type TextEncoderIFace interface {
	Encode(args ...interface{}) Uint8Array
	GetEncoding() string
}
type TextEncoder struct {
	Value
}

func JSValueToTextEncoder(val js.Value) TextEncoder { return TextEncoder{Value: JSValueToValue(val)} }
func (v Value) AsTextEncoder() TextEncoder          { return TextEncoder{Value: v} }
func NewTextEncoder(args ...interface{}) TextEncoder {
	return TextEncoder{Value: JSValueToValue(js.Global().Get("TextEncoder").New(args...))}
}
func (t TextEncoder) Encode(args ...interface{}) Uint8Array {
	val := t.Call("encode", args...)
	return JSValueToUint8Array(val.JSValue())
}
func (t TextEncoder) GetEncoding() string {
	val := t.Get("encoding")
	return val.String()
}
