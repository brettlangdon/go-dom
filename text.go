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
	GetLength() int
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

func JSValueToText(val js.Value) Text { return Text{Value: JSValueToValue(val)} }
func (v Value) AsText() Text          { return Text{Value: v} }
func NewText(args ...interface{}) Text {
	return Text{Value: JSValueToValue(js.Global().Get("Text").New(args...))}
}
func (t Text) GetAssignedSlot() HTMLSlotElement {
	val := t.Get("assignedSlot")
	return JSValueToHTMLSlotElement(val.JSValue())
}
func (t Text) SplitText(args ...interface{}) Text {
	val := t.Call("splitText", args...)
	return JSValueToText(val.JSValue())
}
func (t Text) GetWholeText() string {
	val := t.Get("wholeText")
	return val.String()
}
