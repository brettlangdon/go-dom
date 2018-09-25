// Code generated DO NOT EDIT
// attr.go
package dom

import "syscall/js"

type AttrIFace interface {
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
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetName() string
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOwnerDocument() Document
	GetOwnerElement() Element
	GetParentElement() Element
	GetParentNode() Node
	GetPrefix() string
	GetPreviousSibling() Node
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetSpecified() bool
	GetTextContent() string
	SetTextContent(string)
	GetValue() string
	SetValue(string)
}
type Attr struct {
	Value
}

func JSValueToAttr(val js.Value) Attr { return Attr{Value: JSValueToValue(val)} }
func (v Value) AsAttr() Attr          { return Attr{Value: v} }
func NewAttr(args ...interface{}) Attr {
	return Attr{Value: JSValueToValue(js.Global().Get("Attr").New(args...))}
}
func (a Attr) AddEventListener(args ...interface{}) {
	a.Call("addEventListener", args...)
}
func (a Attr) AppendChild(args ...interface{}) Node {
	val := a.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (a Attr) GetBaseURI() string {
	val := a.Get("baseURI")
	return val.String()
}
func (a Attr) GetChildNodes() NodeList {
	val := a.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (a Attr) CloneNode(args ...interface{}) Node {
	val := a.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (a Attr) CompareDocumentPosition(args ...interface{}) int {
	val := a.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (a Attr) Contains(args ...interface{}) bool {
	val := a.Call("contains", args...)
	return val.Bool()
}
func (a Attr) DispatchEvent(args ...interface{}) bool {
	val := a.Call("dispatchEvent", args...)
	return val.Bool()
}
func (a Attr) GetFirstChild() Node {
	val := a.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (a Attr) GetRootNode(args ...interface{}) Node {
	val := a.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (a Attr) HasChildNodes(args ...interface{}) bool {
	val := a.Call("hasChildNodes", args...)
	return val.Bool()
}
func (a Attr) InsertBefore(args ...interface{}) Node {
	val := a.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (a Attr) GetIsConnected() bool {
	val := a.Get("isConnected")
	return val.Bool()
}
func (a Attr) IsDefaultNamespace(args ...interface{}) bool {
	val := a.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (a Attr) IsEqualNode(args ...interface{}) bool {
	val := a.Call("isEqualNode", args...)
	return val.Bool()
}
func (a Attr) IsSameNode(args ...interface{}) bool {
	val := a.Call("isSameNode", args...)
	return val.Bool()
}
func (a Attr) GetLastChild() Node {
	val := a.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (a Attr) GetLocalName() string {
	val := a.Get("localName")
	return val.String()
}
func (a Attr) LookupNamespaceURI(args ...interface{}) string {
	val := a.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (a Attr) LookupPrefix(args ...interface{}) string {
	val := a.Call("lookupPrefix", args...)
	return val.String()
}
func (a Attr) GetName() string {
	val := a.Get("name")
	return val.String()
}
func (a Attr) GetNamespaceURI() string {
	val := a.Get("namespaceURI")
	return val.String()
}
func (a Attr) GetNextSibling() Node {
	val := a.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (a Attr) GetNodeName() string {
	val := a.Get("nodeName")
	return val.String()
}
func (a Attr) GetNodeType() int {
	val := a.Get("nodeType")
	return val.Int()
}
func (a Attr) GetNodeValue() string {
	val := a.Get("nodeValue")
	return val.String()
}
func (a Attr) SetNodeValue(val string) {
	a.Set("nodeValue", val)
}
func (a Attr) Normalize(args ...interface{}) {
	a.Call("normalize", args...)
}
func (a Attr) GetOwnerDocument() Document {
	val := a.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (a Attr) GetOwnerElement() Element {
	val := a.Get("ownerElement")
	return JSValueToElement(val.JSValue())
}
func (a Attr) GetParentElement() Element {
	val := a.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (a Attr) GetParentNode() Node {
	val := a.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (a Attr) GetPrefix() string {
	val := a.Get("prefix")
	return val.String()
}
func (a Attr) GetPreviousSibling() Node {
	val := a.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (a Attr) RemoveChild(args ...interface{}) Node {
	val := a.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (a Attr) RemoveEventListener(args ...interface{}) {
	a.Call("removeEventListener", args...)
}
func (a Attr) ReplaceChild(args ...interface{}) Node {
	val := a.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (a Attr) GetSpecified() bool {
	val := a.Get("specified")
	return val.Bool()
}
func (a Attr) GetTextContent() string {
	val := a.Get("textContent")
	return val.String()
}
func (a Attr) SetTextContent(val string) {
	a.Set("textContent", val)
}
func (a Attr) GetValue() string {
	val := a.Get("value")
	return val.String()
}
func (a Attr) SetValue(val string) {
	a.Set("value", val)
}
