// Code generated DO NOT EDIT
// blob.go
package dom

import "syscall/js"

type BlobIFace interface {
	GetSize() float64
	Slice(args ...interface{}) Blob
	GetType() string
}
type Blob struct {
	Value
}

func jsValueToBlob(val js.Value) Blob { return Blob{Value: Value{Value: val}} }
func (v Value) AsBlob() Blob          { return Blob{Value: v} }
func (b Blob) GetSize() float64 {
	val := b.Get("size")
	return val.Float()
}
func (b Blob) Slice(args ...interface{}) Blob {
	val := b.Call("slice", args...)
	return jsValueToBlob(val.JSValue())
}
func (b Blob) GetType() string {
	val := b.Get("type")
	return val.String()
}
