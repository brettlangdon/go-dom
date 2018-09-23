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

func jsValueToTextEncoder(val js.Value) TextEncoder { return TextEncoder{Value: Value{Value: val}} }
func (v Value) AsTextEncoder() TextEncoder          { return TextEncoder{Value: v} }
func (t TextEncoder) Encode(args ...interface{}) Uint8Array {
	val := t.Call("encode", args...)
	return jsValueToUint8Array(val.JSValue())
}
func (t TextEncoder) GetEncoding() string {
	val := t.Get("encoding")
	return val.String()
}
