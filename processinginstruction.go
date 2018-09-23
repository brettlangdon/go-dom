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
	GetSheet() StyleSheet
	SubstringData(args ...interface{}) string
	GetTarget() string
	GetTextContent() string
	SetTextContent(string)
}
type ProcessingInstruction struct {
	Value
	CharacterData
	Node
	EventTarget
}

func jsValueToProcessingInstruction(val js.Value) ProcessingInstruction {
	return ProcessingInstruction{Value: Value{Value: val}}
}
func (v Value) AsProcessingInstruction() ProcessingInstruction { return ProcessingInstruction{Value: v} }
func (p ProcessingInstruction) GetSheet() StyleSheet {
	val := p.Get("sheet")
	return jsValueToStyleSheet(val.JSValue())
}
func (p ProcessingInstruction) GetTarget() string {
	val := p.Get("target")
	return val.String()
}
