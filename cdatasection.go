// Code generated DO NOT EDIT
// cdatasection.go
package dom

import "syscall/js"

type CDATASectionIFace interface {
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
	SplitText(args ...interface{}) Text
	SubstringData(args ...interface{}) string
	GetTextContent() string
	SetTextContent(string)
	GetWholeText() string
}
type CDATASection struct {
	Value
}

func JSValueToCDATASection(val js.Value) CDATASection { return CDATASection{Value: JSValueToValue(val)} }
func (v Value) AsCDATASection() CDATASection          { return CDATASection{Value: v} }
func NewCDATASection(args ...interface{}) CDATASection {
	return CDATASection{Value: JSValueToValue(js.Global().Get("CDATASection").New(args...))}
}
func (c CDATASection) AddEventListener(args ...interface{}) {
	c.Call("addEventListener", args...)
}
func (c CDATASection) AppendChild(args ...interface{}) Node {
	val := c.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) AppendData(args ...interface{}) {
	c.Call("appendData", args...)
}
func (c CDATASection) GetBaseURI() string {
	val := c.Get("baseURI")
	return val.String()
}
func (c CDATASection) GetChildNodes() NodeList {
	val := c.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (c CDATASection) CloneNode(args ...interface{}) Node {
	val := c.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) CompareDocumentPosition(args ...interface{}) int {
	val := c.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (c CDATASection) Contains(args ...interface{}) bool {
	val := c.Call("contains", args...)
	return val.Bool()
}
func (c CDATASection) GetData() string {
	val := c.Get("data")
	return val.String()
}
func (c CDATASection) SetData(val string) {
	c.Set("data", val)
}
func (c CDATASection) DeleteData(args ...interface{}) {
	c.Call("deleteData", args...)
}
func (c CDATASection) DispatchEvent(args ...interface{}) bool {
	val := c.Call("dispatchEvent", args...)
	return val.Bool()
}
func (c CDATASection) GetFirstChild() Node {
	val := c.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) GetRootNode(args ...interface{}) Node {
	val := c.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) HasChildNodes(args ...interface{}) bool {
	val := c.Call("hasChildNodes", args...)
	return val.Bool()
}
func (c CDATASection) InsertBefore(args ...interface{}) Node {
	val := c.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) InsertData(args ...interface{}) {
	c.Call("insertData", args...)
}
func (c CDATASection) GetIsConnected() bool {
	val := c.Get("isConnected")
	return val.Bool()
}
func (c CDATASection) IsDefaultNamespace(args ...interface{}) bool {
	val := c.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (c CDATASection) IsEqualNode(args ...interface{}) bool {
	val := c.Call("isEqualNode", args...)
	return val.Bool()
}
func (c CDATASection) IsSameNode(args ...interface{}) bool {
	val := c.Call("isSameNode", args...)
	return val.Bool()
}
func (c CDATASection) GetLastChild() Node {
	val := c.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) GetLength() int {
	val := c.Get("length")
	return val.Int()
}
func (c CDATASection) LookupNamespaceURI(args ...interface{}) string {
	val := c.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (c CDATASection) LookupPrefix(args ...interface{}) string {
	val := c.Call("lookupPrefix", args...)
	return val.String()
}
func (c CDATASection) GetNextSibling() Node {
	val := c.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) GetNodeName() string {
	val := c.Get("nodeName")
	return val.String()
}
func (c CDATASection) GetNodeType() int {
	val := c.Get("nodeType")
	return val.Int()
}
func (c CDATASection) GetNodeValue() string {
	val := c.Get("nodeValue")
	return val.String()
}
func (c CDATASection) SetNodeValue(val string) {
	c.Set("nodeValue", val)
}
func (c CDATASection) Normalize(args ...interface{}) {
	c.Call("normalize", args...)
}
func (c CDATASection) GetOwnerDocument() Document {
	val := c.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (c CDATASection) GetParentElement() Element {
	val := c.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (c CDATASection) GetParentNode() Node {
	val := c.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) GetPreviousSibling() Node {
	val := c.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) RemoveChild(args ...interface{}) Node {
	val := c.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) RemoveEventListener(args ...interface{}) {
	c.Call("removeEventListener", args...)
}
func (c CDATASection) ReplaceChild(args ...interface{}) Node {
	val := c.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (c CDATASection) ReplaceData(args ...interface{}) {
	c.Call("replaceData", args...)
}
func (c CDATASection) SplitText(args ...interface{}) Text {
	val := c.Call("splitText", args...)
	return JSValueToText(val.JSValue())
}
func (c CDATASection) SubstringData(args ...interface{}) string {
	val := c.Call("substringData", args...)
	return val.String()
}
func (c CDATASection) GetTextContent() string {
	val := c.Get("textContent")
	return val.String()
}
func (c CDATASection) SetTextContent(val string) {
	c.Set("textContent", val)
}
func (c CDATASection) GetWholeText() string {
	val := c.Get("wholeText")
	return val.String()
}
