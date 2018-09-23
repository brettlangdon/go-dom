// Code generated DO NOT EDIT
// htmlformelement.go
package dom

import "syscall/js"

type HTMLFormElementIFace interface {
	GetAcceptCharset() string
	SetAcceptCharset(string)
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	GetAction() string
	SetAction(string)
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetAutocomplete() string
	SetAutocomplete(string)
	GetBaseURI() string
	CheckValidity(args ...interface{}) bool
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetElements() HTMLFormControlsCollection
	GetEncoding() string
	SetEncoding(string)
	GetEnctype() string
	SetEnctype(string)
	GetFirstChild() Node
	GetAttribute(args ...interface{}) string
	GetAttributeNS(args ...interface{}) string
	GetAttributeNames(args ...interface{})
	GetAttributeNode(args ...interface{}) Attr
	GetAttributeNodeNS(args ...interface{}) Attr
	GetElementsByClassName(args ...interface{}) HTMLCollection
	GetElementsByTagName(args ...interface{}) HTMLCollection
	GetElementsByTagNameNS(args ...interface{}) HTMLCollection
	GetRootNode(args ...interface{}) Node
	HasAttribute(args ...interface{}) bool
	HasAttributeNS(args ...interface{}) bool
	HasAttributes(args ...interface{}) bool
	HasChildNodes(args ...interface{}) bool
	GetHidden() bool
	SetHidden(bool)
	GetId() string
	SetId(string)
	GetInnerText() string
	SetInnerText(string)
	InsertAdjacentElement(args ...interface{}) Element
	InsertAdjacentText(args ...interface{})
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	GetLength() int
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
	GetMethod() string
	SetMethod(string)
	GetName() string
	SetName(string)
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNoValidate() bool
	SetNoValidate(bool)
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPrefix() string
	GetPreviousSibling() Node
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	ReportValidity(args ...interface{}) bool
	Reset(args ...interface{})
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	Submit(args ...interface{})
	GetTagName() string
	GetTarget() string
	SetTarget(string)
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLFormElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLFormElement(val js.Value) HTMLFormElement {
	return HTMLFormElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLFormElement() HTMLFormElement { return HTMLFormElement{Value: v} }
func (h HTMLFormElement) GetAcceptCharset() string {
	val := h.Get("acceptCharset")
	return val.String()
}
func (h HTMLFormElement) SetAcceptCharset(val string) {
	h.Set("acceptCharset", val)
}
func (h HTMLFormElement) GetAction() string {
	val := h.Get("action")
	return val.String()
}
func (h HTMLFormElement) SetAction(val string) {
	h.Set("action", val)
}
func (h HTMLFormElement) GetAutocomplete() string {
	val := h.Get("autocomplete")
	return val.String()
}
func (h HTMLFormElement) SetAutocomplete(val string) {
	h.Set("autocomplete", val)
}
func (h HTMLFormElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLFormElement) GetElements() HTMLFormControlsCollection {
	val := h.Get("elements")
	return JSValueToHTMLFormControlsCollection(val.JSValue())
}
func (h HTMLFormElement) GetEncoding() string {
	val := h.Get("encoding")
	return val.String()
}
func (h HTMLFormElement) SetEncoding(val string) {
	h.Set("encoding", val)
}
func (h HTMLFormElement) GetEnctype() string {
	val := h.Get("enctype")
	return val.String()
}
func (h HTMLFormElement) SetEnctype(val string) {
	h.Set("enctype", val)
}
func (h HTMLFormElement) GetLength() int {
	val := h.Get("length")
	return val.Int()
}
func (h HTMLFormElement) GetMethod() string {
	val := h.Get("method")
	return val.String()
}
func (h HTMLFormElement) SetMethod(val string) {
	h.Set("method", val)
}
func (h HTMLFormElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLFormElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLFormElement) GetNoValidate() bool {
	val := h.Get("noValidate")
	return val.Bool()
}
func (h HTMLFormElement) SetNoValidate(val bool) {
	h.Set("noValidate", val)
}
func (h HTMLFormElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLFormElement) Reset(args ...interface{}) {
	h.Call("reset", args...)
}
func (h HTMLFormElement) Submit(args ...interface{}) {
	h.Call("submit", args...)
}
func (h HTMLFormElement) GetTarget() string {
	val := h.Get("target")
	return val.String()
}
func (h HTMLFormElement) SetTarget(val string) {
	h.Set("target", val)
}
