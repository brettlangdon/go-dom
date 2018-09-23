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

func jsValueToCanvasPattern(val js.Value) CanvasPattern {
	return CanvasPattern{Value: Value{Value: val}}
}
func (v Value) AsCanvasPattern() CanvasPattern { return CanvasPattern{Value: v} }
func (c CanvasPattern) SetTransform(args ...interface{}) {
	c.Call("setTransform", args...)
}
