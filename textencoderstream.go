// Code generated DO NOT EDIT
// textencoderstream.go
package dom

import "syscall/js"

type TextEncoderStreamIFace interface {
	GetEncoding() string
	GetReadable() Value
	GetWritable() Value
}
type TextEncoderStream struct {
	Value
}

func jsValueToTextEncoderStream(val js.Value) TextEncoderStream {
	return TextEncoderStream{Value: Value{Value: val}}
}
func (v Value) AsTextEncoderStream() TextEncoderStream { return TextEncoderStream{Value: v} }
func (t TextEncoderStream) GetEncoding() string {
	val := t.Get("encoding")
	return val.String()
}
func (t TextEncoderStream) GetReadable() Value {
	val := t.Get("readable")
	return val
}
func (t TextEncoderStream) GetWritable() Value {
	val := t.Get("writable")
	return val
}
