// Code generated DO NOT EDIT
// offscreencanvas.go
package dom

import "syscall/js"

type OffscreenCanvasIFace interface {
	AddEventListener(args ...interface{})
	ConvertToBlob(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetContext(args ...interface{}) OffscreenRenderingContext
	GetHeight() int
	SetHeight(int)
	RemoveEventListener(args ...interface{})
	TransferToImageBitmap(args ...interface{}) ImageBitmap
	GetWidth() int
	SetWidth(int)
}
type OffscreenCanvas struct {
	Value
	EventTarget
}

func JSValueToOffscreenCanvas(val js.Value) OffscreenCanvas {
	return OffscreenCanvas{Value: JSValueToValue(val)}
}
func (v Value) AsOffscreenCanvas() OffscreenCanvas { return OffscreenCanvas{Value: v} }
func NewOffscreenCanvas(args ...interface{}) OffscreenCanvas {
	return OffscreenCanvas{Value: JSValueToValue(js.Global().Get("OffscreenCanvas").New(args...))}
}
func (o OffscreenCanvas) ConvertToBlob(args ...interface{}) {
	o.Call("convertToBlob", args...)
}
func (o OffscreenCanvas) GetContext(args ...interface{}) OffscreenRenderingContext {
	val := o.Call("getContext", args...)
	return JSValueToOffscreenRenderingContext(val.JSValue())
}
func (o OffscreenCanvas) GetHeight() int {
	val := o.Get("height")
	return val.Int()
}
func (o OffscreenCanvas) SetHeight(val int) {
	o.Set("height", val)
}
func (o OffscreenCanvas) TransferToImageBitmap(args ...interface{}) ImageBitmap {
	val := o.Call("transferToImageBitmap", args...)
	return JSValueToImageBitmap(val.JSValue())
}
func (o OffscreenCanvas) GetWidth() int {
	val := o.Get("width")
	return val.Int()
}
func (o OffscreenCanvas) SetWidth(val int) {
	o.Set("width", val)
}
