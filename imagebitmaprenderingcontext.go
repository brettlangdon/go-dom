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
	return ImageBitmapRenderingContext{Value: Value{Value: val}}
}
func (v Value) AsImageBitmapRenderingContext() ImageBitmapRenderingContext {
	return ImageBitmapRenderingContext{Value: v}
}
func (i ImageBitmapRenderingContext) GetCanvas() HTMLCanvasElement {
	val := i.Get("canvas")
	return JSValueToHTMLCanvasElement(val.JSValue())
}
func (i ImageBitmapRenderingContext) TransferFromImageBitmap(args ...interface{}) {
	i.Call("transferFromImageBitmap", args...)
}
