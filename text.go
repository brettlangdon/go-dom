// Code generated DO NOT EDIT
// text.go
package dom

import "syscall/js"

type TextIFace interface {
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AppendData(args ...interface{})
	GetAssignedSlot() HTMLSlotElement
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
type Text struct {
	Value
}

func JSValueToText(val js.Value) Text { return Text{Value: JSValueToValue(val)} }
func (v Value) AsText() Text          { return Text{Value: v} }
func NewText(args ...interface{}) Text {
	return Text{Value: JSValueToValue(js.Global().Get("Text").New(args...))}
}
func (t Text) AddEventListener(args ...interface{}) {
	t.Call("addEventListener", args...)
}
func (t Text) AppendChild(args ...interface{}) Node {
	val := t.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (t Text) AppendData(args ...interface{}) {
	t.Call("appendData", args...)
}
func (t Text) GetAssignedSlot() HTMLSlotElement {
	val := t.Get("assignedSlot")
	return JSValueToHTMLSlotElement(val.JSValue())
}
func (t Text) GetBaseURI() string {
	val := t.Get("baseURI")
	return val.String()
}
func (t Text) GetChildNodes() NodeList {
	val := t.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (t Text) CloneNode(args ...interface{}) Node {
	val := t.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (t Text) CompareDocumentPosition(args ...interface{}) int {
	val := t.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (t Text) Contains(args ...interface{}) bool {
	val := t.Call("contains", args...)
	return val.Bool()
}
func (t Text) GetData() string {
	val := t.Get("data")
	return val.String()
}
func (t Text) SetData(val string) {
	t.Set("data", val)
}
func (t Text) DeleteData(args ...interface{}) {
	t.Call("deleteData", args...)
}
func (t Text) DispatchEvent(args ...interface{}) bool {
	val := t.Call("dispatchEvent", args...)
	return val.Bool()
}
func (t Text) GetFirstChild() Node {
	val := t.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (t Text) GetRootNode(args ...interface{}) Node {
	val := t.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (t Text) HasChildNodes(args ...interface{}) bool {
	val := t.Call("hasChildNodes", args...)
	return val.Bool()
}
func (t Text) InsertBefore(args ...interface{}) Node {
	val := t.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (t Text) InsertData(args ...interface{}) {
	t.Call("insertData", args...)
}
func (t Text) GetIsConnected() bool {
	val := t.Get("isConnected")
	return val.Bool()
}
func (t Text) IsDefaultNamespace(args ...interface{}) bool {
	val := t.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (t Text) IsEqualNode(args ...interface{}) bool {
	val := t.Call("isEqualNode", args...)
	return val.Bool()
}
func (t Text) IsSameNode(args ...interface{}) bool {
	val := t.Call("isSameNode", args...)
	return val.Bool()
}
func (t Text) GetLastChild() Node {
	val := t.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (t Text) GetLength() int {
	val := t.Get("length")
	return val.Int()
}
func (t Text) LookupNamespaceURI(args ...interface{}) string {
	val := t.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (t Text) LookupPrefix(args ...interface{}) string {
	val := t.Call("lookupPrefix", args...)
	return val.String()
}
func (t Text) GetNextSibling() Node {
	val := t.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (t Text) GetNodeName() string {
	val := t.Get("nodeName")
	return val.String()
}
func (t Text) GetNodeType() int {
	val := t.Get("nodeType")
	return val.Int()
}
func (t Text) GetNodeValue() string {
	val := t.Get("nodeValue")
	return val.String()
}
func (t Text) SetNodeValue(val string) {
	t.Set("nodeValue", val)
}
func (t Text) Normalize(args ...interface{}) {
	t.Call("normalize", args...)
}
func (t Text) GetOwnerDocument() Document {
	val := t.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (t Text) GetParentElement() Element {
	val := t.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (t Text) GetParentNode() Node {
	val := t.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (t Text) GetPreviousSibling() Node {
	val := t.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (t Text) RemoveChild(args ...interface{}) Node {
	val := t.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (t Text) RemoveEventListener(args ...interface{}) {
	t.Call("removeEventListener", args...)
}
func (t Text) ReplaceChild(args ...interface{}) Node {
	val := t.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (t Text) ReplaceData(args ...interface{}) {
	t.Call("replaceData", args...)
}
func (t Text) SplitText(args ...interface{}) Text {
	val := t.Call("splitText", args...)
	return JSValueToText(val.JSValue())
}
func (t Text) SubstringData(args ...interface{}) string {
	val := t.Call("substringData", args...)
	return val.String()
}
func (t Text) GetTextContent() string {
	val := t.Get("textContent")
	return val.String()
}
func (t Text) SetTextContent(val string) {
	t.Set("textContent", val)
}
func (t Text) GetWholeText() string {
	val := t.Get("wholeText")
	return val.String()
}
