// Code generated DO NOT EDIT
// nodeiterator.go
package dom

import "syscall/js"

type NodeIteratorIFace interface {
	Detach(args ...interface{})
	GetFilter() NodeFilter
	NextNode(args ...interface{}) Node
	GetPointerBeforeReferenceNode() bool
	PreviousNode(args ...interface{}) Node
	GetReferenceNode() Node
	GetRoot() Node
	GetWhatToShow() float64
}
type NodeIterator struct {
	Value
}

func jsValueToNodeIterator(val js.Value) NodeIterator { return NodeIterator{Value: Value{Value: val}} }
func (v Value) AsNodeIterator() NodeIterator          { return NodeIterator{Value: v} }
func (n NodeIterator) Detach(args ...interface{}) {
	n.Call("detach", args...)
}
func (n NodeIterator) GetFilter() NodeFilter {
	val := n.Get("filter")
	return jsValueToNodeFilter(val.JSValue())
}
func (n NodeIterator) NextNode(args ...interface{}) Node {
	val := n.Call("nextNode", args...)
	return jsValueToNode(val.JSValue())
}
func (n NodeIterator) GetPointerBeforeReferenceNode() bool {
	val := n.Get("pointerBeforeReferenceNode")
	return val.Bool()
}
func (n NodeIterator) PreviousNode(args ...interface{}) Node {
	val := n.Call("previousNode", args...)
	return jsValueToNode(val.JSValue())
}
func (n NodeIterator) GetReferenceNode() Node {
	val := n.Get("referenceNode")
	return jsValueToNode(val.JSValue())
}
func (n NodeIterator) GetRoot() Node {
	val := n.Get("root")
	return jsValueToNode(val.JSValue())
}
func (n NodeIterator) GetWhatToShow() float64 {
	val := n.Get("whatToShow")
	return val.Float()
}
