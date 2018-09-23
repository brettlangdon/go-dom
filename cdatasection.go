// Code generated DO NOT EDIT
// cdatasection.go
package dom

import "syscall/js"

type CDATASectionIFace interface {
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
	GetLength() float64
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
type CDATASection struct {
	Value
	Text
	CharacterData
	Node
	EventTarget
}

func JSValueToCDATASection(val js.Value) CDATASection { return CDATASection{Value: Value{Value: val}} }
func (v Value) AsCDATASection() CDATASection          { return CDATASection{Value: v} }
