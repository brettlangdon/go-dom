// Code generated DO NOT EDIT
// documentfragment.go
package dom

import "syscall/js"

type DocumentFragmentIFace interface {
	AddEventListener(args ...interface{})
	Append(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetBaseURI() string
	GetChildElementCount() int
	GetChildNodes() NodeList
	GetChildren() HTMLCollection
	CloneNode(args ...interface{}) Node
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	DispatchEvent(args ...interface{}) bool
	GetFirstChild() Node
	GetFirstElementChild() Element
	GetElementById(args ...interface{}) Element
	GetRootNode(args ...interface{}) Node
	HasChildNodes(args ...interface{}) bool
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLastChild() Node
	GetLastElementChild() Element
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
	Prepend(args ...interface{})
	GetPreviousSibling() Node
	QuerySelector(args ...interface{}) Element
	QuerySelectorAll(args ...interface{}) NodeList
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetTextContent() string
	SetTextContent(string)
}
type DocumentFragment struct {
	Value
}

func JSValueToDocumentFragment(val js.Value) DocumentFragment {
	return DocumentFragment{Value: JSValueToValue(val)}
}
func (v Value) AsDocumentFragment() DocumentFragment { return DocumentFragment{Value: v} }
func NewDocumentFragment(args ...interface{}) DocumentFragment {
	return DocumentFragment{Value: JSValueToValue(js.Global().Get("DocumentFragment").New(args...))}
}
func (d DocumentFragment) AddEventListener(args ...interface{}) {
	d.Call("addEventListener", args...)
}
func (d DocumentFragment) Append(args ...interface{}) {
	d.Call("append", args...)
}
func (d DocumentFragment) AppendChild(args ...interface{}) Node {
	val := d.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) GetBaseURI() string {
	val := d.Get("baseURI")
	return val.String()
}
func (d DocumentFragment) GetChildElementCount() int {
	val := d.Get("childElementCount")
	return val.Int()
}
func (d DocumentFragment) GetChildNodes() NodeList {
	val := d.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (d DocumentFragment) GetChildren() HTMLCollection {
	val := d.Get("children")
	return JSValueToHTMLCollection(val.JSValue())
}
func (d DocumentFragment) CloneNode(args ...interface{}) Node {
	val := d.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) CompareDocumentPosition(args ...interface{}) int {
	val := d.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (d DocumentFragment) Contains(args ...interface{}) bool {
	val := d.Call("contains", args...)
	return val.Bool()
}
func (d DocumentFragment) DispatchEvent(args ...interface{}) bool {
	val := d.Call("dispatchEvent", args...)
	return val.Bool()
}
func (d DocumentFragment) GetFirstChild() Node {
	val := d.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) GetFirstElementChild() Element {
	val := d.Get("firstElementChild")
	return JSValueToElement(val.JSValue())
}
func (d DocumentFragment) GetElementById(args ...interface{}) Element {
	val := d.Call("getElementById", args...)
	return JSValueToElement(val.JSValue())
}
func (d DocumentFragment) GetRootNode(args ...interface{}) Node {
	val := d.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) HasChildNodes(args ...interface{}) bool {
	val := d.Call("hasChildNodes", args...)
	return val.Bool()
}
func (d DocumentFragment) InsertBefore(args ...interface{}) Node {
	val := d.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) GetIsConnected() bool {
	val := d.Get("isConnected")
	return val.Bool()
}
func (d DocumentFragment) IsDefaultNamespace(args ...interface{}) bool {
	val := d.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (d DocumentFragment) IsEqualNode(args ...interface{}) bool {
	val := d.Call("isEqualNode", args...)
	return val.Bool()
}
func (d DocumentFragment) IsSameNode(args ...interface{}) bool {
	val := d.Call("isSameNode", args...)
	return val.Bool()
}
func (d DocumentFragment) GetLastChild() Node {
	val := d.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) GetLastElementChild() Element {
	val := d.Get("lastElementChild")
	return JSValueToElement(val.JSValue())
}
func (d DocumentFragment) LookupNamespaceURI(args ...interface{}) string {
	val := d.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (d DocumentFragment) LookupPrefix(args ...interface{}) string {
	val := d.Call("lookupPrefix", args...)
	return val.String()
}
func (d DocumentFragment) GetNextSibling() Node {
	val := d.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) GetNodeName() string {
	val := d.Get("nodeName")
	return val.String()
}
func (d DocumentFragment) GetNodeType() int {
	val := d.Get("nodeType")
	return val.Int()
}
func (d DocumentFragment) GetNodeValue() string {
	val := d.Get("nodeValue")
	return val.String()
}
func (d DocumentFragment) SetNodeValue(val string) {
	d.Set("nodeValue", val)
}
func (d DocumentFragment) Normalize(args ...interface{}) {
	d.Call("normalize", args...)
}
func (d DocumentFragment) GetOwnerDocument() Document {
	val := d.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (d DocumentFragment) GetParentElement() Element {
	val := d.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (d DocumentFragment) GetParentNode() Node {
	val := d.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) Prepend(args ...interface{}) {
	d.Call("prepend", args...)
}
func (d DocumentFragment) GetPreviousSibling() Node {
	val := d.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) QuerySelector(args ...interface{}) Element {
	val := d.Call("querySelector", args...)
	return JSValueToElement(val.JSValue())
}
func (d DocumentFragment) QuerySelectorAll(args ...interface{}) NodeList {
	val := d.Call("querySelectorAll", args...)
	return JSValueToNodeList(val.JSValue())
}
func (d DocumentFragment) RemoveChild(args ...interface{}) Node {
	val := d.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) RemoveEventListener(args ...interface{}) {
	d.Call("removeEventListener", args...)
}
func (d DocumentFragment) ReplaceChild(args ...interface{}) Node {
	val := d.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (d DocumentFragment) GetTextContent() string {
	val := d.Get("textContent")
	return val.String()
}
func (d DocumentFragment) SetTextContent(val string) {
	d.Set("textContent", val)
}
