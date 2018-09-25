// Code generated DO NOT EDIT
// shadowroot.go
package dom

import "syscall/js"

type ShadowRootIFace interface {
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
	GetHost() Element
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLastChild() Node
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetMode() ShadowRootMode
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
type ShadowRoot struct {
	Value
}

func JSValueToShadowRoot(val js.Value) ShadowRoot { return ShadowRoot{Value: JSValueToValue(val)} }
func (v Value) AsShadowRoot() ShadowRoot          { return ShadowRoot{Value: v} }
func NewShadowRoot(args ...interface{}) ShadowRoot {
	return ShadowRoot{Value: JSValueToValue(js.Global().Get("ShadowRoot").New(args...))}
}
func (s ShadowRoot) AddEventListener(args ...interface{}) {
	s.Call("addEventListener", args...)
}
func (s ShadowRoot) AppendChild(args ...interface{}) Node {
	val := s.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) GetBaseURI() string {
	val := s.Get("baseURI")
	return val.String()
}
func (s ShadowRoot) GetChildNodes() NodeList {
	val := s.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (s ShadowRoot) CloneNode(args ...interface{}) Node {
	val := s.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) CompareDocumentPosition(args ...interface{}) int {
	val := s.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (s ShadowRoot) Contains(args ...interface{}) bool {
	val := s.Call("contains", args...)
	return val.Bool()
}
func (s ShadowRoot) DispatchEvent(args ...interface{}) bool {
	val := s.Call("dispatchEvent", args...)
	return val.Bool()
}
func (s ShadowRoot) GetFirstChild() Node {
	val := s.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) GetRootNode(args ...interface{}) Node {
	val := s.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) HasChildNodes(args ...interface{}) bool {
	val := s.Call("hasChildNodes", args...)
	return val.Bool()
}
func (s ShadowRoot) GetHost() Element {
	val := s.Get("host")
	return JSValueToElement(val.JSValue())
}
func (s ShadowRoot) InsertBefore(args ...interface{}) Node {
	val := s.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) GetIsConnected() bool {
	val := s.Get("isConnected")
	return val.Bool()
}
func (s ShadowRoot) IsDefaultNamespace(args ...interface{}) bool {
	val := s.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (s ShadowRoot) IsEqualNode(args ...interface{}) bool {
	val := s.Call("isEqualNode", args...)
	return val.Bool()
}
func (s ShadowRoot) IsSameNode(args ...interface{}) bool {
	val := s.Call("isSameNode", args...)
	return val.Bool()
}
func (s ShadowRoot) GetLastChild() Node {
	val := s.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) LookupNamespaceURI(args ...interface{}) string {
	val := s.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (s ShadowRoot) LookupPrefix(args ...interface{}) string {
	val := s.Call("lookupPrefix", args...)
	return val.String()
}
func (s ShadowRoot) GetMode() ShadowRootMode {
	val := s.Get("mode")
	return JSValueToShadowRootMode(val.JSValue())
}
func (s ShadowRoot) GetNextSibling() Node {
	val := s.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) GetNodeName() string {
	val := s.Get("nodeName")
	return val.String()
}
func (s ShadowRoot) GetNodeType() int {
	val := s.Get("nodeType")
	return val.Int()
}
func (s ShadowRoot) GetNodeValue() string {
	val := s.Get("nodeValue")
	return val.String()
}
func (s ShadowRoot) SetNodeValue(val string) {
	s.Set("nodeValue", val)
}
func (s ShadowRoot) Normalize(args ...interface{}) {
	s.Call("normalize", args...)
}
func (s ShadowRoot) GetOwnerDocument() Document {
	val := s.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (s ShadowRoot) GetParentElement() Element {
	val := s.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (s ShadowRoot) GetParentNode() Node {
	val := s.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) GetPreviousSibling() Node {
	val := s.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) RemoveChild(args ...interface{}) Node {
	val := s.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) RemoveEventListener(args ...interface{}) {
	s.Call("removeEventListener", args...)
}
func (s ShadowRoot) ReplaceChild(args ...interface{}) Node {
	val := s.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (s ShadowRoot) GetTextContent() string {
	val := s.Get("textContent")
	return val.String()
}
func (s ShadowRoot) SetTextContent(val string) {
	s.Set("textContent", val)
}
