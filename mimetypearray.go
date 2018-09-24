// Code generated DO NOT EDIT
// mimetypearray.go
package dom

import "syscall/js"

type MimeTypeArrayIFace interface {
	Item(args ...interface{}) MimeType
	GetLength() int
	NamedItem(args ...interface{}) MimeType
}
type MimeTypeArray struct {
	Value
}

func JSValueToMimeTypeArray(val js.Value) MimeTypeArray {
	return MimeTypeArray{Value: JSValueToValue(val)}
}
func (v Value) AsMimeTypeArray() MimeTypeArray { return MimeTypeArray{Value: v} }
func NewMimeTypeArray(args ...interface{}) MimeTypeArray {
	return MimeTypeArray{Value: JSValueToValue(js.Global().Get("MimeTypeArray").New(args...))}
}
func (m MimeTypeArray) Item(args ...interface{}) MimeType {
	val := m.Call("item", args...)
	return JSValueToMimeType(val.JSValue())
}
func (m MimeTypeArray) GetLength() int {
	val := m.Get("length")
	return val.Int()
}
func (m MimeTypeArray) NamedItem(args ...interface{}) MimeType {
	val := m.Call("namedItem", args...)
	return JSValueToMimeType(val.JSValue())
}
