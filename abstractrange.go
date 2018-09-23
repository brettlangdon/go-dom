// Code generated DO NOT EDIT
// abstractrange.go
package dom

import "syscall/js"

type AbstractRangeIFace interface {
	GetCollapsed() bool
	GetEndContainer() Node
	GetEndOffset() float64
	GetStartContainer() Node
	GetStartOffset() float64
}
type AbstractRange struct {
	Value
}

func jsValueToAbstractRange(val js.Value) AbstractRange {
	return AbstractRange{Value: Value{Value: val}}
}
func (v Value) AsAbstractRange() AbstractRange { return AbstractRange{Value: v} }
func (a AbstractRange) GetCollapsed() bool {
	val := a.Get("collapsed")
	return val.Bool()
}
func (a AbstractRange) GetEndContainer() Node {
	val := a.Get("endContainer")
	return jsValueToNode(val.JSValue())
}
func (a AbstractRange) GetEndOffset() float64 {
	val := a.Get("endOffset")
	return val.Float()
}
func (a AbstractRange) GetStartContainer() Node {
	val := a.Get("startContainer")
	return jsValueToNode(val.JSValue())
}
func (a AbstractRange) GetStartOffset() float64 {
	val := a.Get("startOffset")
	return val.Float()
}
