// Code generated DO NOT EDIT
// nodelist.go
package dom

import "syscall/js"

type NodeListIFace interface {
	Item(args ...interface{}) Node
	GetLength() int
}
type NodeList struct {
	Value
}

func JSValueToNodeList(val js.Value) NodeList { return NodeList{Value: JSValueToValue(val)} }
func (v Value) AsNodeList() NodeList          { return NodeList{Value: v} }
func NewNodeList(args ...interface{}) NodeList {
	return NodeList{Value: JSValueToValue(js.Global().Get("NodeList").New(args...))}
}
func (n NodeList) Item(args ...interface{}) Node {
	val := n.Call("item", args...)
	return JSValueToNode(val.JSValue())
}
func (n NodeList) GetLength() int {
	val := n.Get("length")
	return val.Int()
}
