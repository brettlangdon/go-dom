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
}

func JSValueToStaticRange(val js.Value) StaticRange { return StaticRange{Value: JSValueToValue(val)} }
func (v Value) AsStaticRange() StaticRange          { return StaticRange{Value: v} }
func NewStaticRange(args ...interface{}) StaticRange {
	return StaticRange{Value: JSValueToValue(js.Global().Get("StaticRange").New(args...))}
}
func (s StaticRange) GetCollapsed() bool {
	val := s.Get("collapsed")
	return val.Bool()
}
func (s StaticRange) GetEndContainer() Node {
	val := s.Get("endContainer")
	return JSValueToNode(val.JSValue())
}
func (s StaticRange) GetEndOffset() int {
	val := s.Get("endOffset")
	return val.Int()
}
func (s StaticRange) GetStartContainer() Node {
	val := s.Get("startContainer")
	return JSValueToNode(val.JSValue())
}
func (s StaticRange) GetStartOffset() int {
	val := s.Get("startOffset")
	return val.Int()
}
