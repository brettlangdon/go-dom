// Code generated DO NOT EDIT
// shadowroot.go
package dom

import "syscall/js"

type ShadowRootIFace interface {
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
	GetHost() Element
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLastChild() Node
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetMode() ShadowRootMode
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
	GetTextContent() string
	SetTextContent(string)
}
type ShadowRoot struct {
	Value
	DocumentFragment
	Node
	EventTarget
}

func JSValueToShadowRoot(val js.Value) ShadowRoot { return ShadowRoot{Value: JSValueToValue(val)} }
func (v Value) AsShadowRoot() ShadowRoot          { return ShadowRoot{Value: v} }
func NewShadowRoot(args ...interface{}) ShadowRoot {
	return ShadowRoot{Value: JSValueToValue(js.Global().Get("ShadowRoot").New(args...))}
}
func (s ShadowRoot) GetHost() Element {
	val := s.Get("host")
	return JSValueToElement(val.JSValue())
}
func (s ShadowRoot) GetMode() ShadowRootMode {
	val := s.Get("mode")
	return JSValueToShadowRootMode(val.JSValue())
}
