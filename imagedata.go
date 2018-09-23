// Code generated DO NOT EDIT
// imagedata.go
package dom

import "syscall/js"

type ImageDataIFace interface {
	GetData() Uint8ClampedArray
	GetHeight() float64
	GetWidth() float64
}
type ImageData struct {
	Value
}

func jsValueToImageData(val js.Value) ImageData { return ImageData{Value: Value{Value: val}} }
func (v Value) AsImageData() ImageData          { return ImageData{Value: v} }
func (i ImageData) GetData() Uint8ClampedArray {
	val := i.Get("data")
	return jsValueToUint8ClampedArray(val.JSValue())
}
func (i ImageData) GetHeight() float64 {
	val := i.Get("height")
	return val.Float()
}
func (i ImageData) GetWidth() float64 {
	val := i.Get("width")
	return val.Float()
}
