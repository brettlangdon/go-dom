// Code generated DO NOT EDIT
// staticrange.go
package dom

import "syscall/js"

type StaticRangeIFace interface {
	GetCollapsed() bool
	GetEndContainer() Node
	GetEndOffset() float64
	GetStartContainer() Node
	GetStartOffset() float64
}
type StaticRange struct {
	Value
	AbstractRange
}

func JSValueToStaticRange(val js.Value) StaticRange { return StaticRange{Value: Value{Value: val}} }
func (v Value) AsStaticRange() StaticRange          { return StaticRange{Value: v} }
