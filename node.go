// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Node struct {
	js.Value
}

func (n *Node) JSValue() js.Value { return n.Value }
func (n *Node) AddEventListener(t string, listener js.Callback) {
	n.Call("addEventListener", ToJSValue(t), ToJSValue(listener))
}
func (n *Node) AppendChild(aChild *Element) *Element {
	val := n.Call("appendChild", ToJSValue(aChild))
	return &Element{Value: val}
}
func (n *Node) GetBaseURI() string {
	val := n.Get("baseURI")
	return val.String()
}
func (n *Node) GetFirstChild() *Element {
	val := n.Get("firstChild")
	return &Element{Value: val}
}
func (n *Node) GetLastChild() *Element {
	val := n.Get("lastChild")
	return &Element{Value: val}
}
func (n *Node) GetNextSibling() *Element {
	val := n.Get("nextSibling")
	return &Element{Value: val}
}
func (n *Node) GetPreviousSibling() *Element {
	val := n.Get("previousSibling")
	return &Element{Value: val}
}
func (n *Node) GetParentElement() *Element {
	val := n.Get("parentElement")
	return &Element{Value: val}
}
func (n *Node) GetRootElement() *Element {
	val := n.Get("rootElement")
	return &Element{Value: val}
}
func (n *Node) GetPrefix() string {
	val := n.Get("prefix")
	return val.String()
}
func (n *Node) GetNodeName() string {
	val := n.Get("nodeName")
	return val.String()
}
func (n *Node) GetTextContent() string {
	val := n.Get("textContent")
	return val.String()
}
func (n *Node) SetTextContent(v string) {
	n.Set("textContent", v)
}
