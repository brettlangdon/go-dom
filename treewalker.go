// Code generated DO NOT EDIT
// treewalker.go
package dom

import "syscall/js"

type TreeWalkerIFace interface {
	GetCurrentNode() Node
	SetCurrentNode(Node)
	GetFilter() NodeFilter
	FirstChild(args ...interface{}) Node
	LastChild(args ...interface{}) Node
	NextNode(args ...interface{}) Node
	NextSibling(args ...interface{}) Node
	ParentNode(args ...interface{}) Node
	PreviousNode(args ...interface{}) Node
	PreviousSibling(args ...interface{}) Node
	GetRoot() Node
	GetWhatToShow() int
}
type TreeWalker struct {
	Value
}

func JSValueToTreeWalker(val js.Value) TreeWalker { return TreeWalker{Value: Value{Value: val}} }
func (v Value) AsTreeWalker() TreeWalker          { return TreeWalker{Value: v} }
func (t TreeWalker) GetCurrentNode() Node {
	val := t.Get("currentNode")
	return JSValueToNode(val.JSValue())
}
func (t TreeWalker) SetCurrentNode(val Node) {
	t.Set("currentNode", val)
}
func (t TreeWalker) GetFilter() NodeFilter {
	val := t.Get("filter")
	return JSValueToNodeFilter(val.JSValue())
}
func (t TreeWalker) FirstChild(args ...interface{}) Node {
	val := t.Call("firstChild", args...)
	return JSValueToNode(val.JSValue())
}
func (t TreeWalker) LastChild(args ...interface{}) Node {
	val := t.Call("lastChild", args...)
	return JSValueToNode(val.JSValue())
}
func (t TreeWalker) NextNode(args ...interface{}) Node {
	val := t.Call("nextNode", args...)
	return JSValueToNode(val.JSValue())
}
func (t TreeWalker) NextSibling(args ...interface{}) Node {
	val := t.Call("nextSibling", args...)
	return JSValueToNode(val.JSValue())
}
func (t TreeWalker) ParentNode(args ...interface{}) Node {
	val := t.Call("parentNode", args...)
	return JSValueToNode(val.JSValue())
}
func (t TreeWalker) PreviousNode(args ...interface{}) Node {
	val := t.Call("previousNode", args...)
	return JSValueToNode(val.JSValue())
}
func (t TreeWalker) PreviousSibling(args ...interface{}) Node {
	val := t.Call("previousSibling", args...)
	return JSValueToNode(val.JSValue())
}
func (t TreeWalker) GetRoot() Node {
	val := t.Get("root")
	return JSValueToNode(val.JSValue())
}
func (t TreeWalker) GetWhatToShow() int {
	val := t.Get("whatToShow")
	return val.Int()
}
