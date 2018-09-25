// Code generated DO NOT EDIT
// documenttype.go
package dom

import "syscall/js"

type DocumentTypeIFace interface {
	AddEventListener(args ...interface{})
	After(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetBaseURI() string
	Before(args ...interface{})
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
	GetName() string
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
	GetPublicId() string
	Remove(args ...interface{})
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	ReplaceWith(args ...interface{})
	GetSystemId() string
	GetTextContent() string
	SetTextContent(string)
}
type DocumentType struct {
	Value
}

func JSValueToDocumentType(val js.Value) DocumentType { return DocumentType{Value: JSValueToValue(val)} }
func (v Value) AsDocumentType() DocumentType          { return DocumentType{Value: v} }
func NewDocumentType(args ...interface{}) DocumentType {
	return DocumentType{Value: JSValueToValue(js.Global().Get("DocumentType").New(args...))}
}
func (d DocumentType) AddEventListener(args ...interface{}) {
	d.Call("addEventListener", args...)
}
func (d DocumentType) After(args ...interface{}) {
	d.Call("after", args...)
}
func (d DocumentType) AppendChild(args ...interface{}) Node {
	val := d.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) GetBaseURI() string {
	val := d.Get("baseURI")
	return val.String()
}
func (d DocumentType) Before(args ...interface{}) {
	d.Call("before", args...)
}
func (d DocumentType) GetChildNodes() NodeList {
	val := d.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (d DocumentType) CloneNode(args ...interface{}) Node {
	val := d.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) CompareDocumentPosition(args ...interface{}) int {
	val := d.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (d DocumentType) Contains(args ...interface{}) bool {
	val := d.Call("contains", args...)
	return val.Bool()
}
func (d DocumentType) DispatchEvent(args ...interface{}) bool {
	val := d.Call("dispatchEvent", args...)
	return val.Bool()
}
func (d DocumentType) GetFirstChild() Node {
	val := d.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) GetRootNode(args ...interface{}) Node {
	val := d.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) HasChildNodes(args ...interface{}) bool {
	val := d.Call("hasChildNodes", args...)
	return val.Bool()
}
func (d DocumentType) InsertBefore(args ...interface{}) Node {
	val := d.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) GetIsConnected() bool {
	val := d.Get("isConnected")
	return val.Bool()
}
func (d DocumentType) IsDefaultNamespace(args ...interface{}) bool {
	val := d.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (d DocumentType) IsEqualNode(args ...interface{}) bool {
	val := d.Call("isEqualNode", args...)
	return val.Bool()
}
func (d DocumentType) IsSameNode(args ...interface{}) bool {
	val := d.Call("isSameNode", args...)
	return val.Bool()
}
func (d DocumentType) GetLastChild() Node {
	val := d.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) LookupNamespaceURI(args ...interface{}) string {
	val := d.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (d DocumentType) LookupPrefix(args ...interface{}) string {
	val := d.Call("lookupPrefix", args...)
	return val.String()
}
func (d DocumentType) GetName() string {
	val := d.Get("name")
	return val.String()
}
func (d DocumentType) GetNextSibling() Node {
	val := d.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) GetNodeName() string {
	val := d.Get("nodeName")
	return val.String()
}
func (d DocumentType) GetNodeType() int {
	val := d.Get("nodeType")
	return val.Int()
}
func (d DocumentType) GetNodeValue() string {
	val := d.Get("nodeValue")
	return val.String()
}
func (d DocumentType) SetNodeValue(val string) {
	d.Set("nodeValue", val)
}
func (d DocumentType) Normalize(args ...interface{}) {
	d.Call("normalize", args...)
}
func (d DocumentType) GetOwnerDocument() Document {
	val := d.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (d DocumentType) GetParentElement() Element {
	val := d.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (d DocumentType) GetParentNode() Node {
	val := d.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) GetPreviousSibling() Node {
	val := d.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) GetPublicId() string {
	val := d.Get("publicId")
	return val.String()
}
func (d DocumentType) Remove(args ...interface{}) {
	d.Call("remove", args...)
}
func (d DocumentType) RemoveChild(args ...interface{}) Node {
	val := d.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) RemoveEventListener(args ...interface{}) {
	d.Call("removeEventListener", args...)
}
func (d DocumentType) ReplaceChild(args ...interface{}) Node {
	val := d.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentType) ReplaceWith(args ...interface{}) {
	d.Call("replaceWith", args...)
}
func (d DocumentType) GetSystemId() string {
	val := d.Get("systemId")
	return val.String()
}
func (d DocumentType) GetTextContent() string {
	val := d.Get("textContent")
	return val.String()
}
func (d DocumentType) SetTextContent(val string) {
	d.Set("textContent", val)
}
