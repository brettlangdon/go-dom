// Code generated DO NOT EDIT
// xmldocument.go
package dom

import "syscall/js"

type XMLDocumentIFace interface {
	AddEventListener(args ...interface{})
	AdoptNode(args ...interface{}) Node
	AppendChild(args ...interface{}) Node
	GetBaseURI() string
	GetCharacterSet() string
	GetCharset() string
	GetChildNodes() NodeList
	CloneNode(args ...interface{}) Node
	CompareDocumentPosition(args ...interface{}) int
	GetCompatMode() string
	Contains(args ...interface{}) bool
	GetContentType() string
	CreateAttribute(args ...interface{}) Attr
	CreateAttributeNS(args ...interface{}) Attr
	CreateCDATASection(args ...interface{}) CDATASection
	CreateComment(args ...interface{}) Comment
	CreateDocumentFragment(args ...interface{}) DocumentFragment
	CreateElement(args ...interface{}) Element
	CreateElementNS(args ...interface{}) Element
	CreateEvent(args ...interface{}) Event
	CreateNodeIterator(args ...interface{}) NodeIterator
	CreateProcessingInstruction(args ...interface{}) ProcessingInstruction
	CreateRange(args ...interface{}) Range
	CreateTextNode(args ...interface{}) Text
	CreateTreeWalker(args ...interface{}) TreeWalker
	DispatchEvent(args ...interface{}) bool
	GetDoctype() DocumentType
	GetDocumentElement() Element
	GetDocumentURI() string
	GetFirstChild() Node
	GetElementsByClassName(args ...interface{}) HTMLCollection
	GetElementsByTagName(args ...interface{}) HTMLCollection
	GetElementsByTagNameNS(args ...interface{}) HTMLCollection
	GetRootNode(args ...interface{}) Node
	HasChildNodes(args ...interface{}) bool
	GetImplementation() DOMImplementation
	ImportNode(args ...interface{}) Node
	GetInputEncoding() string
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
	GetOrigin() string
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPreviousSibling() Node
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetTextContent() string
	SetTextContent(string)
	GetURL() string
}
type XMLDocument struct {
	Value
}

func JSValueToXMLDocument(val js.Value) XMLDocument { return XMLDocument{Value: JSValueToValue(val)} }
func (v Value) AsXMLDocument() XMLDocument          { return XMLDocument{Value: v} }
func NewXMLDocument(args ...interface{}) XMLDocument {
	return XMLDocument{Value: JSValueToValue(js.Global().Get("XMLDocument").New(args...))}
}
func (x XMLDocument) AddEventListener(args ...interface{}) {
	x.Call("addEventListener", args...)
}
func (x XMLDocument) AdoptNode(args ...interface{}) Node {
	val := x.Call("adoptNode", args...)
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) AppendChild(args ...interface{}) Node {
	val := x.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) GetBaseURI() string {
	val := x.Get("baseURI")
	return val.String()
}
func (x XMLDocument) GetCharacterSet() string {
	val := x.Get("characterSet")
	return val.String()
}
func (x XMLDocument) GetCharset() string {
	val := x.Get("charset")
	return val.String()
}
func (x XMLDocument) GetChildNodes() NodeList {
	val := x.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (x XMLDocument) CloneNode(args ...interface{}) Node {
	val := x.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) CompareDocumentPosition(args ...interface{}) int {
	val := x.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (x XMLDocument) GetCompatMode() string {
	val := x.Get("compatMode")
	return val.String()
}
func (x XMLDocument) Contains(args ...interface{}) bool {
	val := x.Call("contains", args...)
	return val.Bool()
}
func (x XMLDocument) GetContentType() string {
	val := x.Get("contentType")
	return val.String()
}
func (x XMLDocument) CreateAttribute(args ...interface{}) Attr {
	val := x.Call("createAttribute", args...)
	return JSValueToAttr(val.JSValue())
}
func (x XMLDocument) CreateAttributeNS(args ...interface{}) Attr {
	val := x.Call("createAttributeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (x XMLDocument) CreateCDATASection(args ...interface{}) CDATASection {
	val := x.Call("createCDATASection", args...)
	return JSValueToCDATASection(val.JSValue())
}
func (x XMLDocument) CreateComment(args ...interface{}) Comment {
	val := x.Call("createComment", args...)
	return JSValueToComment(val.JSValue())
}
func (x XMLDocument) CreateDocumentFragment(args ...interface{}) DocumentFragment {
	val := x.Call("createDocumentFragment", args...)
	return JSValueToDocumentFragment(val.JSValue())
}
func (x XMLDocument) CreateElement(args ...interface{}) Element {
	val := x.Call("createElement", args...)
	return JSValueToElement(val.JSValue())
}
func (x XMLDocument) CreateElementNS(args ...interface{}) Element {
	val := x.Call("createElementNS", args...)
	return JSValueToElement(val.JSValue())
}
func (x XMLDocument) CreateEvent(args ...interface{}) Event {
	val := x.Call("createEvent", args...)
	return JSValueToEvent(val.JSValue())
}
func (x XMLDocument) CreateNodeIterator(args ...interface{}) NodeIterator {
	val := x.Call("createNodeIterator", args...)
	return JSValueToNodeIterator(val.JSValue())
}
func (x XMLDocument) CreateProcessingInstruction(args ...interface{}) ProcessingInstruction {
	val := x.Call("createProcessingInstruction", args...)
	return JSValueToProcessingInstruction(val.JSValue())
}
func (x XMLDocument) CreateRange(args ...interface{}) Range {
	val := x.Call("createRange", args...)
	return JSValueToRange(val.JSValue())
}
func (x XMLDocument) CreateTextNode(args ...interface{}) Text {
	val := x.Call("createTextNode", args...)
	return JSValueToText(val.JSValue())
}
func (x XMLDocument) CreateTreeWalker(args ...interface{}) TreeWalker {
	val := x.Call("createTreeWalker", args...)
	return JSValueToTreeWalker(val.JSValue())
}
func (x XMLDocument) DispatchEvent(args ...interface{}) bool {
	val := x.Call("dispatchEvent", args...)
	return val.Bool()
}
func (x XMLDocument) GetDoctype() DocumentType {
	val := x.Get("doctype")
	return JSValueToDocumentType(val.JSValue())
}
func (x XMLDocument) GetDocumentElement() Element {
	val := x.Get("documentElement")
	return JSValueToElement(val.JSValue())
}
func (x XMLDocument) GetDocumentURI() string {
	val := x.Get("documentURI")
	return val.String()
}
func (x XMLDocument) GetFirstChild() Node {
	val := x.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := x.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (x XMLDocument) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := x.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (x XMLDocument) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := x.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (x XMLDocument) GetRootNode(args ...interface{}) Node {
	val := x.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) HasChildNodes(args ...interface{}) bool {
	val := x.Call("hasChildNodes", args...)
	return val.Bool()
}
func (x XMLDocument) GetImplementation() DOMImplementation {
	val := x.Get("implementation")
	return JSValueToDOMImplementation(val.JSValue())
}
func (x XMLDocument) ImportNode(args ...interface{}) Node {
	val := x.Call("importNode", args...)
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) GetInputEncoding() string {
	val := x.Get("inputEncoding")
	return val.String()
}
func (x XMLDocument) InsertBefore(args ...interface{}) Node {
	val := x.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) GetIsConnected() bool {
	val := x.Get("isConnected")
	return val.Bool()
}
func (x XMLDocument) IsDefaultNamespace(args ...interface{}) bool {
	val := x.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (x XMLDocument) IsEqualNode(args ...interface{}) bool {
	val := x.Call("isEqualNode", args...)
	return val.Bool()
}
func (x XMLDocument) IsSameNode(args ...interface{}) bool {
	val := x.Call("isSameNode", args...)
	return val.Bool()
}
func (x XMLDocument) GetLastChild() Node {
	val := x.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) LookupNamespaceURI(args ...interface{}) string {
	val := x.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (x XMLDocument) LookupPrefix(args ...interface{}) string {
	val := x.Call("lookupPrefix", args...)
	return val.String()
}
func (x XMLDocument) GetNextSibling() Node {
	val := x.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) GetNodeName() string {
	val := x.Get("nodeName")
	return val.String()
}
func (x XMLDocument) GetNodeType() int {
	val := x.Get("nodeType")
	return val.Int()
}
func (x XMLDocument) GetNodeValue() string {
	val := x.Get("nodeValue")
	return val.String()
}
func (x XMLDocument) SetNodeValue(val string) {
	x.Set("nodeValue", val)
}
func (x XMLDocument) Normalize(args ...interface{}) {
	x.Call("normalize", args...)
}
func (x XMLDocument) GetOrigin() string {
	val := x.Get("origin")
	return val.String()
}
func (x XMLDocument) GetOwnerDocument() Document {
	val := x.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (x XMLDocument) GetParentElement() Element {
	val := x.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (x XMLDocument) GetParentNode() Node {
	val := x.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) GetPreviousSibling() Node {
	val := x.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) RemoveChild(args ...interface{}) Node {
	val := x.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) RemoveEventListener(args ...interface{}) {
	x.Call("removeEventListener", args...)
}
func (x XMLDocument) ReplaceChild(args ...interface{}) Node {
	val := x.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (x XMLDocument) GetTextContent() string {
	val := x.Get("textContent")
	return val.String()
}
func (x XMLDocument) SetTextContent(val string) {
	x.Set("textContent", val)
}
func (x XMLDocument) GetURL() string {
	val := x.Get("URL")
	return val.String()
}
