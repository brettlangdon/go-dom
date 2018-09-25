// Code generated DO NOT EDIT
// comment.go
package dom

import "syscall/js"

type CommentIFace interface {
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AppendData(args ...interface{})
	GetBaseURI() string
	GetChildNodes() NodeList
	CloneNode(args ...interface{}) Node
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	GetData() string
	SetData(string)
	DeleteData(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetFirstChild() Node
	GetRootNode(args ...interface{}) Node
	HasChildNodes(args ...interface{}) bool
	InsertBefore(args ...interface{}) Node
	InsertData(args ...interface{})
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLastChild() Node
	GetLength() int
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
	ReplaceData(args ...interface{})
	SubstringData(args ...interface{}) string
	GetTextContent() string
	SetTextContent(string)
}
type Comment struct {
	Value
}

func JSValueToComment(val js.Value) Comment { return Comment{Value: JSValueToValue(val)} }
func (v Value) AsComment() Comment          { return Comment{Value: v} }
func NewComment(args ...interface{}) Comment {
	return Comment{Value: JSValueToValue(js.Global().Get("Comment").New(args...))}
}
func (c Comment) AddEventListener(args ...interface{}) {
	c.Call("addEventListener", args...)
}
func (c Comment) AppendChild(args ...interface{}) Node {
	val := c.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (c Comment) AppendData(args ...interface{}) {
	c.Call("appendData", args...)
}
func (c Comment) GetBaseURI() string {
	val := c.Get("baseURI")
	return val.String()
}
func (c Comment) GetChildNodes() NodeList {
	val := c.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (c Comment) CloneNode(args ...interface{}) Node {
	val := c.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (c Comment) CompareDocumentPosition(args ...interface{}) int {
	val := c.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (c Comment) Contains(args ...interface{}) bool {
	val := c.Call("contains", args...)
	return val.Bool()
}
func (c Comment) GetData() string {
	val := c.Get("data")
	return val.String()
}
func (c Comment) SetData(val string) {
	c.Set("data", val)
}
func (c Comment) DeleteData(args ...interface{}) {
	c.Call("deleteData", args...)
}
func (c Comment) DispatchEvent(args ...interface{}) bool {
	val := c.Call("dispatchEvent", args...)
	return val.Bool()
}
func (c Comment) GetFirstChild() Node {
	val := c.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (c Comment) GetRootNode(args ...interface{}) Node {
	val := c.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (c Comment) HasChildNodes(args ...interface{}) bool {
	val := c.Call("hasChildNodes", args...)
	return val.Bool()
}
func (c Comment) InsertBefore(args ...interface{}) Node {
	val := c.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (c Comment) InsertData(args ...interface{}) {
	c.Call("insertData", args...)
}
func (c Comment) GetIsConnected() bool {
	val := c.Get("isConnected")
	return val.Bool()
}
func (c Comment) IsDefaultNamespace(args ...interface{}) bool {
	val := c.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (c Comment) IsEqualNode(args ...interface{}) bool {
	val := c.Call("isEqualNode", args...)
	return val.Bool()
}
func (c Comment) IsSameNode(args ...interface{}) bool {
	val := c.Call("isSameNode", args...)
	return val.Bool()
}
func (c Comment) GetLastChild() Node {
	val := c.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (c Comment) GetLength() int {
	val := c.Get("length")
	return val.Int()
}
func (c Comment) LookupNamespaceURI(args ...interface{}) string {
	val := c.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (c Comment) LookupPrefix(args ...interface{}) string {
	val := c.Call("lookupPrefix", args...)
	return val.String()
}
func (c Comment) GetNextSibling() Node {
	val := c.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (c Comment) GetNodeName() string {
	val := c.Get("nodeName")
	return val.String()
}
func (c Comment) GetNodeType() int {
	val := c.Get("nodeType")
	return val.Int()
}
func (c Comment) GetNodeValue() string {
	val := c.Get("nodeValue")
	return val.String()
}
func (c Comment) SetNodeValue(val string) {
	c.Set("nodeValue", val)
}
func (c Comment) Normalize(args ...interface{}) {
	c.Call("normalize", args...)
}
func (c Comment) GetOwnerDocument() Document {
	val := c.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (c Comment) GetParentElement() Element {
	val := c.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (c Comment) GetParentNode() Node {
	val := c.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (c Comment) GetPreviousSibling() Node {
	val := c.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (c Comment) RemoveChild(args ...interface{}) Node {
	val := c.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (c Comment) RemoveEventListener(args ...interface{}) {
	c.Call("removeEventListener", args...)
}
func (c Comment) ReplaceChild(args ...interface{}) Node {
	val := c.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (c Comment) ReplaceData(args ...interface{}) {
	c.Call("replaceData", args...)
}
func (c Comment) SubstringData(args ...interface{}) string {
	val := c.Call("substringData", args...)
	return val.String()
}
func (c Comment) GetTextContent() string {
	val := c.Get("textContent")
	return val.String()
}
func (c Comment) SetTextContent(val string) {
	c.Set("textContent", val)
}
