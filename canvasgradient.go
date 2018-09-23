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

func jsValueToCanvasGradient(val js.Value) CanvasGradient {
	return CanvasGradient{Value: Value{Value: val}}
}
func (v Value) AsCanvasGradient() CanvasGradient { return CanvasGradient{Value: v} }
func (c CanvasGradient) AddColorStop(args ...interface{}) {
	c.Call("addColorStop", args...)
}
