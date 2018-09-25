// Code generated DO NOT EDIT
// characterdata.go
package dom

import "syscall/js"

type CharacterDataIFace interface {
	AddEventListener(args ...interface{})
	After(args ...interface{})
	AppendChild(args ...interface{}) Node
	AppendData(args ...interface{})
	GetBaseURI() string
	Before(args ...interface{})
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
	GetNextElementSibling() Element
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPreviousElementSibling() Element
	GetPreviousSibling() Node
	Remove(args ...interface{})
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	ReplaceData(args ...interface{})
	ReplaceWith(args ...interface{})
	SubstringData(args ...interface{}) string
	GetTextContent() string
	SetTextContent(string)
}
type CharacterData struct {
	Value
}

func JSValueToCharacterData(val js.Value) CharacterData {
	return CharacterData{Value: JSValueToValue(val)}
}
func (v Value) AsCharacterData() CharacterData { return CharacterData{Value: v} }
func NewCharacterData(args ...interface{}) CharacterData {
	return CharacterData{Value: JSValueToValue(js.Global().Get("CharacterData").New(args...))}
}
func (c CharacterData) AddEventListener(args ...interface{}) {
	c.Call("addEventListener", args...)
}
func (c CharacterData) After(args ...interface{}) {
	c.Call("after", args...)
}
func (c CharacterData) AppendChild(args ...interface{}) Node {
	val := c.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) AppendData(args ...interface{}) {
	c.Call("appendData", args...)
}
func (c CharacterData) GetBaseURI() string {
	val := c.Get("baseURI")
	return val.String()
}
func (c CharacterData) Before(args ...interface{}) {
	c.Call("before", args...)
}
func (c CharacterData) GetChildNodes() NodeList {
	val := c.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (c CharacterData) CloneNode(args ...interface{}) Node {
	val := c.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) CompareDocumentPosition(args ...interface{}) int {
	val := c.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (c CharacterData) Contains(args ...interface{}) bool {
	val := c.Call("contains", args...)
	return val.Bool()
}
func (c CharacterData) GetData() string {
	val := c.Get("data")
	return val.String()
}
func (c CharacterData) SetData(val string) {
	c.Set("data", val)
}
func (c CharacterData) DeleteData(args ...interface{}) {
	c.Call("deleteData", args...)
}
func (c CharacterData) DispatchEvent(args ...interface{}) bool {
	val := c.Call("dispatchEvent", args...)
	return val.Bool()
}
func (c CharacterData) GetFirstChild() Node {
	val := c.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) GetRootNode(args ...interface{}) Node {
	val := c.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) HasChildNodes(args ...interface{}) bool {
	val := c.Call("hasChildNodes", args...)
	return val.Bool()
}
func (c CharacterData) InsertBefore(args ...interface{}) Node {
	val := c.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) InsertData(args ...interface{}) {
	c.Call("insertData", args...)
}
func (c CharacterData) GetIsConnected() bool {
	val := c.Get("isConnected")
	return val.Bool()
}
func (c CharacterData) IsDefaultNamespace(args ...interface{}) bool {
	val := c.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (c CharacterData) IsEqualNode(args ...interface{}) bool {
	val := c.Call("isEqualNode", args...)
	return val.Bool()
}
func (c CharacterData) IsSameNode(args ...interface{}) bool {
	val := c.Call("isSameNode", args...)
	return val.Bool()
}
func (c CharacterData) GetLastChild() Node {
	val := c.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) GetLength() int {
	val := c.Get("length")
	return val.Int()
}
func (c CharacterData) LookupNamespaceURI(args ...interface{}) string {
	val := c.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (c CharacterData) LookupPrefix(args ...interface{}) string {
	val := c.Call("lookupPrefix", args...)
	return val.String()
}
func (c CharacterData) GetNextElementSibling() Element {
	val := c.Get("nextElementSibling")
	return JSValueToElement(val.JSValue())
}
func (c CharacterData) GetNextSibling() Node {
	val := c.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) GetNodeName() string {
	val := c.Get("nodeName")
	return val.String()
}
func (c CharacterData) GetNodeType() int {
	val := c.Get("nodeType")
	return val.Int()
}
func (c CharacterData) GetNodeValue() string {
	val := c.Get("nodeValue")
	return val.String()
}
func (c CharacterData) SetNodeValue(val string) {
	c.Set("nodeValue", val)
}
func (c CharacterData) Normalize(args ...interface{}) {
	c.Call("normalize", args...)
}
func (c CharacterData) GetOwnerDocument() Document {
	val := c.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (c CharacterData) GetParentElement() Element {
	val := c.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (c CharacterData) GetParentNode() Node {
	val := c.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) GetPreviousElementSibling() Element {
	val := c.Get("previousElementSibling")
	return JSValueToElement(val.JSValue())
}
func (c CharacterData) GetPreviousSibling() Node {
	val := c.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) Remove(args ...interface{}) {
	c.Call("remove", args...)
}
func (c CharacterData) RemoveChild(args ...interface{}) Node {
	val := c.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) RemoveEventListener(args ...interface{}) {
	c.Call("removeEventListener", args...)
}
func (c CharacterData) ReplaceChild(args ...interface{}) Node {
	val := c.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (c CharacterData) ReplaceData(args ...interface{}) {
	c.Call("replaceData", args...)
}
func (c CharacterData) ReplaceWith(args ...interface{}) {
	c.Call("replaceWith", args...)
}
func (c CharacterData) SubstringData(args ...interface{}) string {
	val := c.Call("substringData", args...)
	return val.String()
}
func (c CharacterData) GetTextContent() string {
	val := c.Get("textContent")
	return val.String()
}
func (c CharacterData) SetTextContent(val string) {
	c.Set("textContent", val)
}
