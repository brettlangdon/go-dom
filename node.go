// Code generated DO NOT EDIT
// node.go
package dom

import "syscall/js"

type NodeIFace interface {
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetBaseURI() string
	GetChildNodes() NodeList
	CloneNode(args ...interface{}) Node
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	DispatchEvent(args ...interface{}) bool
	GetFirstChild() Node
	GetRootNode(args ...interface{}) Node
	HasChildNodes(args ...interface{}) bool
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLastChild() Node
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPreviousSibling() Node
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetTextContent() string
	SetTextContent(string)
}
type Node struct {
	Value
}

func JSValueToNode(val js.Value) Node { return Node{Value: JSValueToValue(val)} }
func (v Value) AsNode() Node          { return Node{Value: v} }
func NewNode(args ...interface{}) Node {
	return Node{Value: JSValueToValue(js.Global().Get("Node").New(args...))}
}
func (n Node) AddEventListener(args ...interface{}) {
	n.Call("addEventListener", args...)
}
func (n Node) AppendChild(args ...interface{}) Node {
	val := n.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (n Node) GetBaseURI() string {
	val := n.Get("baseURI")
	return val.String()
}
func (n Node) GetChildNodes() NodeList {
	val := n.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (n Node) CloneNode(args ...interface{}) Node {
	val := n.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (n Node) CompareDocumentPosition(args ...interface{}) int {
	val := n.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (n Node) Contains(args ...interface{}) bool {
	val := n.Call("contains", args...)
	return val.Bool()
}
func (n Node) DispatchEvent(args ...interface{}) bool {
	val := n.Call("dispatchEvent", args...)
	return val.Bool()
}
func (n Node) GetFirstChild() Node {
	val := n.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (n Node) GetRootNode(args ...interface{}) Node {
	val := n.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (n Node) HasChildNodes(args ...interface{}) bool {
	val := n.Call("hasChildNodes", args...)
	return val.Bool()
}
func (n Node) InsertBefore(args ...interface{}) Node {
	val := n.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (n Node) GetIsConnected() bool {
	val := n.Get("isConnected")
	return val.Bool()
}
func (n Node) IsDefaultNamespace(args ...interface{}) bool {
	val := n.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (n Node) IsEqualNode(args ...interface{}) bool {
	val := n.Call("isEqualNode", args...)
	return val.Bool()
}
func (n Node) IsSameNode(args ...interface{}) bool {
	val := n.Call("isSameNode", args...)
	return val.Bool()
}
func (n Node) GetLastChild() Node {
	val := n.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (n Node) LookupNamespaceURI(args ...interface{}) string {
	val := n.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (n Node) LookupPrefix(args ...interface{}) string {
	val := n.Call("lookupPrefix", args...)
	return val.String()
}
func (n Node) GetNextSibling() Node {
	val := n.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (n Node) GetNodeName() string {
	val := n.Get("nodeName")
	return val.String()
}
func (n Node) GetNodeType() int {
	val := n.Get("nodeType")
	return val.Int()
}
func (n Node) GetNodeValue() string {
	val := n.Get("nodeValue")
	return val.String()
}
func (n Node) SetNodeValue(val string) {
	n.Set("nodeValue", val)
}
func (n Node) Normalize(args ...interface{}) {
	n.Call("normalize", args...)
}
func (n Node) GetOwnerDocument() Document {
	val := n.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (n Node) GetParentElement() Element {
	val := n.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (n Node) GetParentNode() Node {
	val := n.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (n Node) GetPreviousSibling() Node {
	val := n.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (n Node) RemoveChild(args ...interface{}) Node {
	val := n.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (n Node) RemoveEventListener(args ...interface{}) {
	n.Call("removeEventListener", args...)
}
func (n Node) ReplaceChild(args ...interface{}) Node {
	val := n.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (n Node) GetTextContent() string {
	val := n.Get("textContent")
	return val.String()
}
func (n Node) SetTextContent(val string) {
	n.Set("textContent", val)
}
