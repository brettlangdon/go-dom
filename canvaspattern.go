// Code generated DO NOT EDIT
// canvaspattern.go
package dom

import "syscall/js"

type CanvasPatternIFace interface {
	SetTransform(args ...interface{})
}
type CanvasPattern struct {
	Value
}

func JSValueToCanvasPattern(val js.Value) CanvasPattern {
	return CanvasPattern{Value: JSValueToValue(val)}
}
func (v Value) AsCanvasPattern() CanvasPattern { return CanvasPattern{Value: v} }
func NewCanvasPattern(args ...interface{}) CanvasPattern {
	return CanvasPattern{Value: JSValueToValue(js.Global().Get("CanvasPattern").New(args...))}
}
func (c CanvasPattern) SetTransform(args ...interface{}) {
	c.Call("setTransform", args...)
}
