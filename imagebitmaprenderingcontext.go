// Code generated DO NOT EDIT
// imagebitmaprenderingcontext.go
package dom

import "syscall/js"

type ImageBitmapRenderingContextIFace interface {
	GetCanvas() HTMLCanvasElement
	TransferFromImageBitmap(args ...interface{})
}
type ImageBitmapRenderingContext struct {
	Value
}

func JSValueToImageBitmapRenderingContext(val js.Value) ImageBitmapRenderingContext {
	return ImageBitmapRenderingContext{Value: JSValueToValue(val)}
}
func (v Value) AsImageBitmapRenderingContext() ImageBitmapRenderingContext {
	return ImageBitmapRenderingContext{Value: v}
}
func NewImageBitmapRenderingContext(args ...interface{}) ImageBitmapRenderingContext {
	return ImageBitmapRenderingContext{Value: JSValueToValue(js.Global().Get("ImageBitmapRenderingContext").New(args...))}
}
func (i ImageBitmapRenderingContext) GetCanvas() HTMLCanvasElement {
	val := i.Get("canvas")
	return JSValueToHTMLCanvasElement(val.JSValue())
}
func (i ImageBitmapRenderingContext) TransferFromImageBitmap(args ...interface{}) {
	i.Call("transferFromImageBitmap", args...)
}
