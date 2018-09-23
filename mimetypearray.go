// Code generated DO NOT EDIT
// mimetypearray.go
package dom

import "syscall/js"

type MimeTypeArrayIFace interface {
	Item(args ...interface{}) MimeType
	GetLength() float64
	NamedItem(args ...interface{}) MimeType
}
type MimeTypeArray struct {
	Value
}

func JSValueToMimeTypeArray(val js.Value) MimeTypeArray {
	return MimeTypeArray{Value: Value{Value: val}}
}
func (v Value) AsMimeTypeArray() MimeTypeArray { return MimeTypeArray{Value: v} }
func (m MimeTypeArray) Item(args ...interface{}) MimeType {
	val := m.Call("item", args...)
	return JSValueToMimeType(val.JSValue())
}
func (m MimeTypeArray) GetLength() float64 {
	val := m.Get("length")
	return val.Float()
}
func (m MimeTypeArray) NamedItem(args ...interface{}) MimeType {
	val := m.Call("namedItem", args...)
	return JSValueToMimeType(val.JSValue())
}
