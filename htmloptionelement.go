// Code generated DO NOT EDIT
// htmloptionelement.go
package dom

import "syscall/js"

type HTMLOptionElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetBaseURI() string
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	GetDefaultSelected() bool
	SetDefaultSelected(bool)
	GetDir() string
	SetDir(string)
	GetDisabled() bool
	SetDisabled(bool)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetFirstChild() Node
	GetForm() HTMLFormElement
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
	GetIndex() float64
	GetInnerText() string
	SetInnerText(string)
	InsertAdjacentElement(args ...interface{}) Element
	InsertAdjacentText(args ...interface{})
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLabel() string
	SetLabel(string)
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
	GetNamespaceURI() string
	GetNextSibling() Node
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
	GetSelected() bool
	SetSelected(bool)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetTagName() string
	GetText() string
	SetText(string)
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetValue() string
	SetValue(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLOptionElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func jsValueToHTMLOptionElement(val js.Value) HTMLOptionElement {
	return HTMLOptionElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLOptionElement() HTMLOptionElement { return HTMLOptionElement{Value: v} }
func (h HTMLOptionElement) GetDefaultSelected() bool {
	val := h.Get("defaultSelected")
	return val.Bool()
}
func (h HTMLOptionElement) SetDefaultSelected(val bool) {
	h.Set("defaultSelected", val)
}
func (h HTMLOptionElement) GetDisabled() bool {
	val := h.Get("disabled")
	return val.Bool()
}
func (h HTMLOptionElement) SetDisabled(val bool) {
	h.Set("disabled", val)
}
func (h HTMLOptionElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return jsValueToHTMLFormElement(val.JSValue())
}
func (h HTMLOptionElement) GetIndex() float64 {
	val := h.Get("index")
	return val.Float()
}
func (h HTMLOptionElement) GetLabel() string {
	val := h.Get("label")
	return val.String()
}
func (h HTMLOptionElement) SetLabel(val string) {
	h.Set("label", val)
}
func (h HTMLOptionElement) GetSelected() bool {
	val := h.Get("selected")
	return val.Bool()
}
func (h HTMLOptionElement) SetSelected(val bool) {
	h.Set("selected", val)
}
func (h HTMLOptionElement) GetText() string {
	val := h.Get("text")
	return val.String()
}
func (h HTMLOptionElement) SetText(val string) {
	h.Set("text", val)
}
func (h HTMLOptionElement) GetValue() string {
	val := h.Get("value")
	return val.String()
}
func (h HTMLOptionElement) SetValue(val string) {
	h.Set("value", val)
}
