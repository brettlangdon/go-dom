// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Node struct {
	Value
}

func NewNode(v js.Value) *Node {
	val := Value{Value: v}
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return val.ToNode()
}
func (v Value) ToNode() *Node { return &Node{Value: v} }
func (n *Node) AddEventListener(t string, listener *Callback) Value {
	val := Value{Value: n.Call("addEventListener", ToJSValue(t), ToJSValue(listener))}
	return val
}
func (n *Node) AppendChild(aChild *Element) *Element {
	val := Value{Value: n.Call("appendChild", ToJSValue(aChild))}
	return NewElement(val.JSValue())
}
func (n *Node) GetBaseURI() string {
	val := Value{Value: n.Get("baseURI")}
	return val.String()
}
func (n *Node) GetFirstChild() *Element {
	val := Value{Value: n.Get("firstChild")}
	return NewElement(val.JSValue())
}
func (n *Node) GetLastChild() *Element {
	val := Value{Value: n.Get("lastChild")}
	return NewElement(val.JSValue())
}
func (n *Node) GetNextSibling() *Element {
	val := Value{Value: n.Get("nextSibling")}
	return NewElement(val.JSValue())
}
func (n *Node) GetPreviousSibling() *Element {
	val := Value{Value: n.Get("previousSibling")}
	return NewElement(val.JSValue())
}
func (n *Node) GetParentElement() *Element {
	val := Value{Value: n.Get("parentElement")}
	return NewElement(val.JSValue())
}
func (n *Node) GetRootElement() *Element {
	val := Value{Value: n.Get("rootElement")}
	return NewElement(val.JSValue())
}
func (n *Node) GetPrefix() string {
	val := Value{Value: n.Get("prefix")}
	return val.String()
}
func (n *Node) GetNodeName() string {
	val := Value{Value: n.Get("nodeName")}
	return val.String()
}
func (n *Node) GetTextContent() string {
	val := Value{Value: n.Get("textContent")}
	return val.String()
}
func (n *Node) SetTextContent(v string) {
	n.Set("textContent", v)
}
