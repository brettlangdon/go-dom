// Code generated DO NOT EDIT
// imagedata.go
package dom

import "syscall/js"

type ImageDataIFace interface {
	GetData() Uint8ClampedArray
	GetHeight() int
	GetWidth() int
}
type ImageData struct {
	Value
}

func JSValueToImageData(val js.Value) ImageData { return ImageData{Value: Value{Value: val}} }
func (v Value) AsImageData() ImageData          { return ImageData{Value: v} }
func (i ImageData) GetData() Uint8ClampedArray {
	val := i.Get("data")
	return JSValueToUint8ClampedArray(val.JSValue())
}
func (i ImageData) GetHeight() int {
	val := i.Get("height")
	return val.Int()
}
func (i ImageData) GetWidth() int {
	val := i.Get("width")
	return val.Int()
}
