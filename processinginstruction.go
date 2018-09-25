// Code generated DO NOT EDIT
// processinginstruction.go
package dom

import "syscall/js"

type ProcessingInstructionIFace interface {
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
	GetSheet() StyleSheet
	SubstringData(args ...interface{}) string
	GetTarget() string
	GetTextContent() string
	SetTextContent(string)
}
type ProcessingInstruction struct {
	Value
}

func JSValueToProcessingInstruction(val js.Value) ProcessingInstruction {
	return ProcessingInstruction{Value: JSValueToValue(val)}
}
func (v Value) AsProcessingInstruction() ProcessingInstruction { return ProcessingInstruction{Value: v} }
func NewProcessingInstruction(args ...interface{}) ProcessingInstruction {
	return ProcessingInstruction{Value: JSValueToValue(js.Global().Get("ProcessingInstruction").New(args...))}
}
func (p ProcessingInstruction) AddEventListener(args ...interface{}) {
	p.Call("addEventListener", args...)
}
func (p ProcessingInstruction) AppendChild(args ...interface{}) Node {
	val := p.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) AppendData(args ...interface{}) {
	p.Call("appendData", args...)
}
func (p ProcessingInstruction) GetBaseURI() string {
	val := p.Get("baseURI")
	return val.String()
}
func (p ProcessingInstruction) GetChildNodes() NodeList {
	val := p.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (p ProcessingInstruction) CloneNode(args ...interface{}) Node {
	val := p.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) CompareDocumentPosition(args ...interface{}) int {
	val := p.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (p ProcessingInstruction) Contains(args ...interface{}) bool {
	val := p.Call("contains", args...)
	return val.Bool()
}
func (p ProcessingInstruction) GetData() string {
	val := p.Get("data")
	return val.String()
}
func (p ProcessingInstruction) SetData(val string) {
	p.Set("data", val)
}
func (p ProcessingInstruction) DeleteData(args ...interface{}) {
	p.Call("deleteData", args...)
}
func (p ProcessingInstruction) DispatchEvent(args ...interface{}) bool {
	val := p.Call("dispatchEvent", args...)
	return val.Bool()
}
func (p ProcessingInstruction) GetFirstChild() Node {
	val := p.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) GetRootNode(args ...interface{}) Node {
	val := p.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) HasChildNodes(args ...interface{}) bool {
	val := p.Call("hasChildNodes", args...)
	return val.Bool()
}
func (p ProcessingInstruction) InsertBefore(args ...interface{}) Node {
	val := p.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) InsertData(args ...interface{}) {
	p.Call("insertData", args...)
}
func (p ProcessingInstruction) GetIsConnected() bool {
	val := p.Get("isConnected")
	return val.Bool()
}
func (p ProcessingInstruction) IsDefaultNamespace(args ...interface{}) bool {
	val := p.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (p ProcessingInstruction) IsEqualNode(args ...interface{}) bool {
	val := p.Call("isEqualNode", args...)
	return val.Bool()
}
func (p ProcessingInstruction) IsSameNode(args ...interface{}) bool {
	val := p.Call("isSameNode", args...)
	return val.Bool()
}
func (p ProcessingInstruction) GetLastChild() Node {
	val := p.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) GetLength() int {
	val := p.Get("length")
	return val.Int()
}
func (p ProcessingInstruction) LookupNamespaceURI(args ...interface{}) string {
	val := p.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (p ProcessingInstruction) LookupPrefix(args ...interface{}) string {
	val := p.Call("lookupPrefix", args...)
	return val.String()
}
func (p ProcessingInstruction) GetNextSibling() Node {
	val := p.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) GetNodeName() string {
	val := p.Get("nodeName")
	return val.String()
}
func (p ProcessingInstruction) GetNodeType() int {
	val := p.Get("nodeType")
	return val.Int()
}
func (p ProcessingInstruction) GetNodeValue() string {
	val := p.Get("nodeValue")
	return val.String()
}
func (p ProcessingInstruction) SetNodeValue(val string) {
	p.Set("nodeValue", val)
}
func (p ProcessingInstruction) Normalize(args ...interface{}) {
	p.Call("normalize", args...)
}
func (p ProcessingInstruction) GetOwnerDocument() Document {
	val := p.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (p ProcessingInstruction) GetParentElement() Element {
	val := p.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (p ProcessingInstruction) GetParentNode() Node {
	val := p.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) GetPreviousSibling() Node {
	val := p.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) RemoveChild(args ...interface{}) Node {
	val := p.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) RemoveEventListener(args ...interface{}) {
	p.Call("removeEventListener", args...)
}
func (p ProcessingInstruction) ReplaceChild(args ...interface{}) Node {
	val := p.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (p ProcessingInstruction) ReplaceData(args ...interface{}) {
	p.Call("replaceData", args...)
}
func (p ProcessingInstruction) GetSheet() StyleSheet {
	val := p.Get("sheet")
	return JSValueToStyleSheet(val.JSValue())
}
func (p ProcessingInstruction) SubstringData(args ...interface{}) string {
	val := p.Call("substringData", args...)
	return val.String()
}
func (p ProcessingInstruction) GetTarget() string {
	val := p.Get("target")
	return val.String()
}
func (p ProcessingInstruction) GetTextContent() string {
	val := p.Get("textContent")
	return val.String()
}
func (p ProcessingInstruction) SetTextContent(val string) {
	p.Set("textContent", val)
}
