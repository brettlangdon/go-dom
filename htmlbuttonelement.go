// Code generated DO NOT EDIT
// htmlbuttonelement.go
package dom

import "syscall/js"

type HTMLButtonElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetAutofocus() bool
	SetAutofocus(bool)
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
	GetDisabled() bool
	SetDisabled(bool)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetFirstChild() Node
	GetForm() HTMLFormElement
	GetFormAction() string
	SetFormAction(string)
	GetFormEnctype() string
	SetFormEnctype(string)
	GetFormMethod() string
	SetFormMethod(string)
	GetFormNoValidate() bool
	SetFormNoValidate(bool)
	GetFormTarget() string
	SetFormTarget(string)
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
	SetType(string)
	GetValidationMessage() string
	GetValidity() ValidityState
	GetValue() string
	SetValue(string)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWillValidate() bool
}
type HTMLButtonElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLButtonElement(val js.Value) HTMLButtonElement {
	return HTMLButtonElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLButtonElement() HTMLButtonElement { return HTMLButtonElement{Value: v} }
func (h HTMLButtonElement) GetAutofocus() bool {
	val := h.Get("autofocus")
	return val.Bool()
}
func (h HTMLButtonElement) SetAutofocus(val bool) {
	h.Set("autofocus", val)
}
func (h HTMLButtonElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLButtonElement) GetDisabled() bool {
	val := h.Get("disabled")
	return val.Bool()
}
func (h HTMLButtonElement) SetDisabled(val bool) {
	h.Set("disabled", val)
}
func (h HTMLButtonElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLButtonElement) GetFormAction() string {
	val := h.Get("formAction")
	return val.String()
}
func (h HTMLButtonElement) SetFormAction(val string) {
	h.Set("formAction", val)
}
func (h HTMLButtonElement) GetFormEnctype() string {
	val := h.Get("formEnctype")
	return val.String()
}
func (h HTMLButtonElement) SetFormEnctype(val string) {
	h.Set("formEnctype", val)
}
func (h HTMLButtonElement) GetFormMethod() string {
	val := h.Get("formMethod")
	return val.String()
}
func (h HTMLButtonElement) SetFormMethod(val string) {
	h.Set("formMethod", val)
}
func (h HTMLButtonElement) GetFormNoValidate() bool {
	val := h.Get("formNoValidate")
	return val.Bool()
}
func (h HTMLButtonElement) SetFormNoValidate(val bool) {
	h.Set("formNoValidate", val)
}
func (h HTMLButtonElement) GetFormTarget() string {
	val := h.Get("formTarget")
	return val.String()
}
func (h HTMLButtonElement) SetFormTarget(val string) {
	h.Set("formTarget", val)
}
func (h HTMLButtonElement) GetLabels() NodeList {
	val := h.Get("labels")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLButtonElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLButtonElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLButtonElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLButtonElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLButtonElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLButtonElement) SetType(val string) {
	h.Set("type", val)
}
func (h HTMLButtonElement) GetValidationMessage() string {
	val := h.Get("validationMessage")
	return val.String()
}
func (h HTMLButtonElement) GetValidity() ValidityState {
	val := h.Get("validity")
	return JSValueToValidityState(val.JSValue())
}
func (h HTMLButtonElement) GetValue() string {
	val := h.Get("value")
	return val.String()
}
func (h HTMLButtonElement) SetValue(val string) {
	h.Set("value", val)
}
func (h HTMLButtonElement) GetWillValidate() bool {
	val := h.Get("willValidate")
	return val.Bool()
}
