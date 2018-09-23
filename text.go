// Code generated DO NOT EDIT
// text.go
package dom

import "syscall/js"

type TextIFace interface {
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AppendData(args ...interface{})
	GetAssignedSlot() HTMLSlotElement
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
type Text struct {
	Value
	CharacterData
	Node
	EventTarget
}

func jsValueToText(val js.Value) Text { return Text{Value: Value{Value: val}} }
func (v Value) AsText() Text          { return Text{Value: v} }
func (t Text) GetAssignedSlot() HTMLSlotElement {
	val := t.Get("assignedSlot")
	return jsValueToHTMLSlotElement(val.JSValue())
}
func (t Text) SplitText(args ...interface{}) Text {
	val := t.Call("splitText", args...)
	return jsValueToText(val.JSValue())
}
func (t Text) GetWholeText() string {
	val := t.Get("wholeText")
	return val.String()
}