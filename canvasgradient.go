// Code generated DO NOT EDIT
// canvasgradient.go
package dom

import "syscall/js"

type CanvasGradientIFace interface {
	AddColorStop(args ...interface{})
}
type CanvasGradient struct {
	Value
}

func JSValueToCanvasGradient(val js.Value) CanvasGradient {
	return CanvasGradient{Value: JSValueToValue(val)}
}
func (v Value) AsCanvasGradient() CanvasGradient { return CanvasGradient{Value: v} }
func NewCanvasGradient(args ...interface{}) CanvasGradient {
	return CanvasGradient{Value: JSValueToValue(js.Global().Get("CanvasGradient").New(args...))}
}
func (c CanvasGradient) AddColorStop(args ...interface{}) {
	c.Call("addColorStop", args...)
}
