// Code generated DO NOT EDIT
// abstractrange.go
package dom

import "syscall/js"

type AbstractRangeIFace interface {
	GetCollapsed() bool
	GetEndContainer() Node
	GetEndOffset() int
	GetStartContainer() Node
	GetStartOffset() int
}
type AbstractRange struct {
	Value
}

func JSValueToAbstractRange(val js.Value) AbstractRange {
	return AbstractRange{Value: JSValueToValue(val)}
}
func (v Value) AsAbstractRange() AbstractRange { return AbstractRange{Value: v} }
func NewAbstractRange(args ...interface{}) AbstractRange {
	return AbstractRange{Value: JSValueToValue(js.Global().Get("AbstractRange").New(args...))}
}
func (a AbstractRange) GetCollapsed() bool {
	val := a.Get("collapsed")
	return val.Bool()
}
func (a AbstractRange) GetEndContainer() Node {
	val := a.Get("endContainer")
	return JSValueToNode(val.JSValue())
}
func (a AbstractRange) GetEndOffset() int {
	val := a.Get("endOffset")
	return val.Int()
}
func (a AbstractRange) GetStartContainer() Node {
	val := a.Get("startContainer")
	return JSValueToNode(val.JSValue())
}
func (a AbstractRange) GetStartOffset() int {
	val := a.Get("startOffset")
	return val.Int()
}
