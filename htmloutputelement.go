// Code generated DO NOT EDIT
// htmloutputelement.go
package dom

import "syscall/js"

type HTMLOutputElementIFace interface {
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
	GetDefaultValue() string
	SetDefaultValue(string)
	GetDir() string
	SetDir(string)
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
	GetHtmlFor() DOMTokenList
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
	GetLabels() NodeList
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
	GetName() string
	SetName(string)
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
	ReportValidity(args ...interface{}) bool
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	SetCustomValidity(args ...interface{})
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetType() string
	GetValidationMessage() string
	GetValidity() ValidityState
	GetValue() string
	SetValue(string)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWillValidate() bool
}
type HTMLOutputElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLOutputElement(val js.Value) HTMLOutputElement {
	return HTMLOutputElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLOutputElement() HTMLOutputElement { return HTMLOutputElement{Value: v} }
func (h HTMLOutputElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLOutputElement) GetDefaultValue() string {
	val := h.Get("defaultValue")
	return val.String()
}
func (h HTMLOutputElement) SetDefaultValue(val string) {
	h.Set("defaultValue", val)
}
func (h HTMLOutputElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLOutputElement) GetHtmlFor() DOMTokenList {
	val := h.Get("htmlFor")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLOutputElement) GetLabels() NodeList {
	val := h.Get("labels")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLOutputElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLOutputElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLOutputElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLOutputElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLOutputElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLOutputElement) GetValidationMessage() string {
	val := h.Get("validationMessage")
	return val.String()
}
func (h HTMLOutputElement) GetValidity() ValidityState {
	val := h.Get("validity")
	return JSValueToValidityState(val.JSValue())
}
func (h HTMLOutputElement) GetValue() string {
	val := h.Get("value")
	return val.String()
}
func (h HTMLOutputElement) SetValue(val string) {
	h.Set("value", val)
}
func (h HTMLOutputElement) GetWillValidate() bool {
	val := h.Get("willValidate")
	return val.Bool()
}
