// Code generated DO NOT EDIT
// staticrange.go
package dom

import "syscall/js"

type StaticRangeIFace interface {
	GetCollapsed() bool
	GetEndContainer() Node
	GetEndOffset() int
	GetStartContainer() Node
	GetStartOffset() int
}
type StaticRange struct {
	Value
	AbstractRange
}

func JSValueToStaticRange(val js.Value) StaticRange { return StaticRange{Value: JSValueToValue(val)} }
func (v Value) AsStaticRange() StaticRange          { return StaticRange{Value: v} }
func NewStaticRange(args ...interface{}) StaticRange {
	return StaticRange{Value: JSValueToValue(js.Global().Get("StaticRange").New(args...))}
}
