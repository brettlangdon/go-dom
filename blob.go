// Code generated DO NOT EDIT
// blob.go
package dom

import "syscall/js"

type BlobIFace interface {
	GetSize() int
	Slice(args ...interface{}) Blob
	GetType() string
}
type Blob struct {
	Value
}

func JSValueToBlob(val js.Value) Blob { return Blob{Value: Value{Value: val}} }
func (v Value) AsBlob() Blob          { return Blob{Value: v} }
func (b Blob) GetSize() int {
	val := b.Get("size")
	return val.Int()
}
func (b Blob) Slice(args ...interface{}) Blob {
	val := b.Call("slice", args...)
	return JSValueToBlob(val.JSValue())
}
func (b Blob) GetType() string {
	val := b.Get("type")
	return val.String()
}
