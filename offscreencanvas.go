// Code generated DO NOT EDIT
// offscreencanvas.go
package dom

import "syscall/js"

type OffscreenCanvasIFace interface {
	AddEventListener(args ...interface{})
	ConvertToBlob(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetContext(args ...interface{}) OffscreenRenderingContext
	GetHeight() float64
	SetHeight(float64)
	RemoveEventListener(args ...interface{})
	TransferToImageBitmap(args ...interface{}) ImageBitmap
	GetWidth() float64
	SetWidth(float64)
}
type OffscreenCanvas struct {
	Value
	EventTarget
}

func JSValueToOffscreenCanvas(val js.Value) OffscreenCanvas {
	return OffscreenCanvas{Value: Value{Value: val}}
}
func (v Value) AsOffscreenCanvas() OffscreenCanvas { return OffscreenCanvas{Value: v} }
func (o OffscreenCanvas) ConvertToBlob(args ...interface{}) {
	o.Call("convertToBlob", args...)
}
func (o OffscreenCanvas) GetContext(args ...interface{}) OffscreenRenderingContext {
	val := o.Call("getContext", args...)
	return JSValueToOffscreenRenderingContext(val.JSValue())
}
func (o OffscreenCanvas) GetHeight() float64 {
	val := o.Get("height")
	return val.Float()
}
func (o OffscreenCanvas) SetHeight(val float64) {
	o.Set("height", val)
}
func (o OffscreenCanvas) TransferToImageBitmap(args ...interface{}) ImageBitmap {
	val := o.Call("transferToImageBitmap", args...)
	return JSValueToImageBitmap(val.JSValue())
}
func (o OffscreenCanvas) GetWidth() float64 {
	val := o.Get("width")
	return val.Float()
}
func (o OffscreenCanvas) SetWidth(val float64) {
	o.Set("width", val)
}
