// Code generated DO NOT EDIT
// imagebitmap.go
package dom

import "syscall/js"

type ImageBitmapIFace interface {
	Close(args ...interface{})
	GetHeight() int
	GetWidth() int
}
type ImageBitmap struct {
	Value
}

func JSValueToImageBitmap(val js.Value) ImageBitmap { return ImageBitmap{Value: Value{Value: val}} }
func (v Value) AsImageBitmap() ImageBitmap          { return ImageBitmap{Value: v} }
func (i ImageBitmap) Close(args ...interface{}) {
	i.Call("close", args...)
}
func (i ImageBitmap) GetHeight() int {
	val := i.Get("height")
	return val.Int()
}
func (i ImageBitmap) GetWidth() int {
	val := i.Get("width")
	return val.Int()
}
