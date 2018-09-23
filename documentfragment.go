// Code generated DO NOT EDIT
// documentfragment.go
package dom

import "syscall/js"

type DocumentFragmentIFace interface {
	AddEventListener(args ...interface{})
	Append(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetBaseURI() string
	GetChildElementCount() float64
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
	Node
	EventTarget
}

func JSValueToDocumentFragment(val js.Value) DocumentFragment {
	return DocumentFragment{Value: Value{Value: val}}
}
func (v Value) AsDocumentFragment() DocumentFragment { return DocumentFragment{Value: v} }
func (d DocumentFragment) Append(args ...interface{}) {
	d.Call("append", args...)
}
func (d DocumentFragment) GetChildElementCount() float64 {
	val := d.Get("childElementCount")
	return val.Float()
}
func (d DocumentFragment) GetChildren() HTMLCollection {
	val := d.Get("children")
	return JSValueToHTMLCollection(val.JSValue())
}
func (d DocumentFragment) GetFirstElementChild() Element {
	val := d.Get("firstElementChild")
	return JSValueToElement(val.JSValue())
}
func (d DocumentFragment) GetElementById(args ...interface{}) Element {
	val := d.Call("getElementById", args...)
	return JSValueToElement(val.JSValue())
}
func (d DocumentFragment) GetLastElementChild() Element {
	val := d.Get("lastElementChild")
	return JSValueToElement(val.JSValue())
}
func (d DocumentFragment) Prepend(args ...interface{}) {
	d.Call("prepend", args...)
}
func (d DocumentFragment) QuerySelector(args ...interface{}) Element {
	val := d.Call("querySelector", args...)
	return JSValueToElement(val.JSValue())
}
func (d DocumentFragment) QuerySelectorAll(args ...interface{}) NodeList {
	val := d.Call("querySelectorAll", args...)
	return JSValueToNodeList(val.JSValue())
}
