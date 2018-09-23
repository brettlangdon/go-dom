// Code generated DO NOT EDIT
// characterdata.go
package dom

import "syscall/js"

type CharacterDataIFace interface {
	AddEventListener(args ...interface{})
	After(args ...interface{})
	AppendChild(args ...interface{}) Node
	AppendData(args ...interface{})
	GetBaseURI() string
	Before(args ...interface{})
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
	GetNextElementSibling() Element
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPreviousElementSibling() Element
	GetPreviousSibling() Node
	Remove(args ...interface{})
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	ReplaceData(args ...interface{})
	ReplaceWith(args ...interface{})
	SubstringData(args ...interface{}) string
	GetTextContent() string
	SetTextContent(string)
}
type CharacterData struct {
	Value
	Node
	EventTarget
}

func JSValueToCharacterData(val js.Value) CharacterData {
	return CharacterData{Value: Value{Value: val}}
}
func (v Value) AsCharacterData() CharacterData { return CharacterData{Value: v} }
func (c CharacterData) After(args ...interface{}) {
	c.Call("after", args...)
}
func (c CharacterData) AppendData(args ...interface{}) {
	c.Call("appendData", args...)
}
func (c CharacterData) Before(args ...interface{}) {
	c.Call("before", args...)
}
func (c CharacterData) GetData() string {
	val := c.Get("data")
	return val.String()
}
func (c CharacterData) SetData(val string) {
	c.Set("data", val)
}
func (c CharacterData) DeleteData(args ...interface{}) {
	c.Call("deleteData", args...)
}
func (c CharacterData) InsertData(args ...interface{}) {
	c.Call("insertData", args...)
}
func (c CharacterData) GetLength() float64 {
	val := c.Get("length")
	return val.Float()
}
func (c CharacterData) GetNextElementSibling() Element {
	val := c.Get("nextElementSibling")
	return JSValueToElement(val.JSValue())
}
func (c CharacterData) GetPreviousElementSibling() Element {
	val := c.Get("previousElementSibling")
	return JSValueToElement(val.JSValue())
}
func (c CharacterData) Remove(args ...interface{}) {
	c.Call("remove", args...)
}
func (c CharacterData) ReplaceData(args ...interface{}) {
	c.Call("replaceData", args...)
}
func (c CharacterData) ReplaceWith(args ...interface{}) {
	c.Call("replaceWith", args...)
}
func (c CharacterData) SubstringData(args ...interface{}) string {
	val := c.Call("substringData", args...)
	return val.String()
}
