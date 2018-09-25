// Code generated DO NOT EDIT
// range.go
package dom

import "syscall/js"

type RangeIFace interface {
	CloneContents(args ...interface{}) DocumentFragment
	CloneRange(args ...interface{}) Range
	Collapse(args ...interface{})
	GetCollapsed() bool
	GetCommonAncestorContainer() Node
	CompareBoundaryPoints(args ...interface{}) int
	ComparePoint(args ...interface{}) int
	CreateContextualFragment(args ...interface{}) DocumentFragment
	DeleteContents(args ...interface{})
	Detach(args ...interface{})
	GetEndContainer() Node
	GetEndOffset() int
	ExtractContents(args ...interface{}) DocumentFragment
	InsertNode(args ...interface{})
	IntersectsNode(args ...interface{}) bool
	IsPointInRange(args ...interface{}) bool
	SelectNode(args ...interface{})
	SelectNodeContents(args ...interface{})
	SetEnd(args ...interface{})
	SetEndAfter(args ...interface{})
	SetEndBefore(args ...interface{})
	SetStart(args ...interface{})
	SetStartAfter(args ...interface{})
	SetStartBefore(args ...interface{})
	GetStartContainer() Node
	GetStartOffset() int
	SurroundContents(args ...interface{})
}
type Range struct {
	Value
}

func JSValueToRange(val js.Value) Range { return Range{Value: JSValueToValue(val)} }
func (v Value) AsRange() Range          { return Range{Value: v} }
func NewRange(args ...interface{}) Range {
	return Range{Value: JSValueToValue(js.Global().Get("Range").New(args...))}
}
func (r Range) CloneContents(args ...interface{}) DocumentFragment {
	val := r.Call("cloneContents", args...)
	return JSValueToDocumentFragment(val.JSValue())
}
func (r Range) CloneRange(args ...interface{}) Range {
	val := r.Call("cloneRange", args...)
	return JSValueToRange(val.JSValue())
}
func (r Range) Collapse(args ...interface{}) {
	r.Call("collapse", args...)
}
func (r Range) GetCollapsed() bool {
	val := r.Get("collapsed")
	return val.Bool()
}
func (r Range) GetCommonAncestorContainer() Node {
	val := r.Get("commonAncestorContainer")
	return JSValueToNode(val.JSValue())
}
func (r Range) CompareBoundaryPoints(args ...interface{}) int {
	val := r.Call("compareBoundaryPoints", args...)
	return val.Int()
}
func (r Range) ComparePoint(args ...interface{}) int {
	val := r.Call("comparePoint", args...)
	return val.Int()
}
func (r Range) CreateContextualFragment(args ...interface{}) DocumentFragment {
	val := r.Call("createContextualFragment", args...)
	return JSValueToDocumentFragment(val.JSValue())
}
func (r Range) DeleteContents(args ...interface{}) {
	r.Call("deleteContents", args...)
}
func (r Range) Detach(args ...interface{}) {
	r.Call("detach", args...)
}
func (r Range) GetEndContainer() Node {
	val := r.Get("endContainer")
	return JSValueToNode(val.JSValue())
}
func (r Range) GetEndOffset() int {
	val := r.Get("endOffset")
	return val.Int()
}
func (r Range) ExtractContents(args ...interface{}) DocumentFragment {
	val := r.Call("extractContents", args...)
	return JSValueToDocumentFragment(val.JSValue())
}
func (r Range) InsertNode(args ...interface{}) {
	r.Call("insertNode", args...)
}
func (r Range) IntersectsNode(args ...interface{}) bool {
	val := r.Call("intersectsNode", args...)
	return val.Bool()
}
func (r Range) IsPointInRange(args ...interface{}) bool {
	val := r.Call("isPointInRange", args...)
	return val.Bool()
}
func (r Range) SelectNode(args ...interface{}) {
	r.Call("selectNode", args...)
}
func (r Range) SelectNodeContents(args ...interface{}) {
	r.Call("selectNodeContents", args...)
}
func (r Range) SetEnd(args ...interface{}) {
	r.Call("setEnd", args...)
}
func (r Range) SetEndAfter(args ...interface{}) {
	r.Call("setEndAfter", args...)
}
func (r Range) SetEndBefore(args ...interface{}) {
	r.Call("setEndBefore", args...)
}
func (r Range) SetStart(args ...interface{}) {
	r.Call("setStart", args...)
}
func (r Range) SetStartAfter(args ...interface{}) {
	r.Call("setStartAfter", args...)
}
func (r Range) SetStartBefore(args ...interface{}) {
	r.Call("setStartBefore", args...)
}
func (r Range) GetStartContainer() Node {
	val := r.Get("startContainer")
	return JSValueToNode(val.JSValue())
}
func (r Range) GetStartOffset() int {
	val := r.Get("startOffset")
	return val.Int()
}
func (r Range) SurroundContents(args ...interface{}) {
	r.Call("surroundContents", args...)
}
