// Code generated DO NOT EDIT
// attr.go
package dom

import "syscall/js"

type AttrIFace interface {
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
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLastChild() Node
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetName() string
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOwnerDocument() Document
	GetOwnerElement() Element
	GetParentElement() Element
	GetParentNode() Node
	GetPrefix() string
	GetPreviousSibling() Node
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetSpecified() bool
	GetTextContent() string
	SetTextContent(string)
	GetValue() string
	SetValue(string)
}
type Attr struct {
	Value
	Node
	EventTarget
}

func jsValueToAttr(val js.Value) Attr { return Attr{Value: Value{Value: val}} }
func (v Value) AsAttr() Attr          { return Attr{Value: v} }
func (a Attr) GetLocalName() string {
	val := a.Get("localName")
	return val.String()
}
func (a Attr) GetName() string {
	val := a.Get("name")
	return val.String()
}
func (a Attr) GetNamespaceURI() string {
	val := a.Get("namespaceURI")
	return val.String()
}
func (a Attr) GetOwnerElement() Element {
	val := a.Get("ownerElement")
	return jsValueToElement(val.JSValue())
}
func (a Attr) GetPrefix() string {
	val := a.Get("prefix")
	return val.String()
}
func (a Attr) GetSpecified() bool {
	val := a.Get("specified")
	return val.Bool()
}
func (a Attr) GetValue() string {
	val := a.Get("value")
	return val.String()
}
func (a Attr) SetValue(val string) {
	a.Set("value", val)
}
