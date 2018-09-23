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
	Document
	Node
	EventTarget
}

func JSValueToXMLDocument(val js.Value) XMLDocument { return XMLDocument{Value: Value{Value: val}} }
func (v Value) AsXMLDocument() XMLDocument          { return XMLDocument{Value: v} }
