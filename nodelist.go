// Code generated DO NOT EDIT
// nodelist.go
package dom

import "syscall/js"

type NodeListIFace interface {
	Item(args ...interface{}) Node
	GetLength() float64
}
type NodeList struct {
	Value
}

func JSValueToNodeList(val js.Value) NodeList { return NodeList{Value: Value{Value: val}} }
func (v Value) AsNodeList() NodeList          { return NodeList{Value: v} }
func (n NodeList) Item(args ...interface{}) Node {
	val := n.Call("item", args...)
	return JSValueToNode(val.JSValue())
}
func (n NodeList) GetLength() float64 {
	val := n.Get("length")
	return val.Float()
}
