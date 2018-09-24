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
	Node
	EventTarget
}

func JSValueToDocumentType(val js.Value) DocumentType { return DocumentType{Value: JSValueToValue(val)} }
func (v Value) AsDocumentType() DocumentType          { return DocumentType{Value: v} }
func NewDocumentType(args ...interface{}) DocumentType {
	return DocumentType{Value: JSValueToValue(js.Global().Get("DocumentType").New(args...))}
}
func (d DocumentType) After(args ...interface{}) {
	d.Call("after", args...)
}
func (d DocumentType) Before(args ...interface{}) {
	d.Call("before", args...)
}
func (d DocumentType) GetName() string {
	val := d.Get("name")
	return val.String()
}
func (d DocumentType) GetPublicId() string {
	val := d.Get("publicId")
	return val.String()
}
func (d DocumentType) Remove(args ...interface{}) {
	d.Call("remove", args...)
}
func (d DocumentType) ReplaceWith(args ...interface{}) {
	d.Call("replaceWith", args...)
}
func (d DocumentType) GetSystemId() string {
	val := d.Get("systemId")
	return val.String()
}
