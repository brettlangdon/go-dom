// Code generated DO NOT EDIT
// imagebitmap.go
package dom

import "syscall/js"

type ImageBitmapIFace interface {
	Close(args ...interface{})
	GetHeight() float64
	GetWidth() float64
}
type ImageBitmap struct {
	Value
}

func jsValueToImageBitmap(val js.Value) ImageBitmap { return ImageBitmap{Value: Value{Value: val}} }
func (v Value) AsImageBitmap() ImageBitmap          { return ImageBitmap{Value: v} }
func (i ImageBitmap) Close(args ...interface{}) {
	i.Call("close", args...)
}
func (i ImageBitmap) GetHeight() float64 {
	val := i.Get("height")
	return val.Float()
}
func (i ImageBitmap) GetWidth() float64 {
	val := i.Get("width")
	return val.Float()
}
