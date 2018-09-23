// Code generated DO NOT EDIT
// htmlselectelement.go
package dom

import "syscall/js"

type HTMLSelectElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	Add(args ...interface{})
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetAutocomplete() string
	SetAutocomplete(string)
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
	Item(args ...interface{}) Element
	GetLabels() NodeList
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	GetLength() float64
	SetLength(float64)
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
	GetMultiple() bool
	SetMultiple(bool)
	GetName() string
	SetName(string)
	NamedItem(args ...interface{}) HTMLOptionElement
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOptions() HTMLOptionsCollection
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPrefix() string
	GetPreviousSibling() Node
	Remove(args ...interface{})
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	RemoveWithArgs(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	ReportValidity(args ...interface{}) bool
	GetRequired() bool
	SetRequired(bool)
	GetSelectedIndex() float64
	SetSelectedIndex(float64)
	GetSelectedOptions() HTMLCollection
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	SetCustomValidity(args ...interface{})
	GetShadowRoot() ShadowRoot
	GetSize() float64
	SetSize(float64)
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
type HTMLSelectElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func jsValueToHTMLSelectElement(val js.Value) HTMLSelectElement {
	return HTMLSelectElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLSelectElement() HTMLSelectElement { return HTMLSelectElement{Value: v} }
func (h HTMLSelectElement) Add(args ...interface{}) {
	h.Call("add", args...)
}
func (h HTMLSelectElement) GetAutocomplete() string {
	val := h.Get("autocomplete")
	return val.String()
}
func (h HTMLSelectElement) SetAutocomplete(val string) {
	h.Set("autocomplete", val)
}
func (h HTMLSelectElement) GetAutofocus() bool {
	val := h.Get("autofocus")
	return val.Bool()
}
func (h HTMLSelectElement) SetAutofocus(val bool) {
	h.Set("autofocus", val)
}
func (h HTMLSelectElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLSelectElement) GetDisabled() bool {
	val := h.Get("disabled")
	return val.Bool()
}
func (h HTMLSelectElement) SetDisabled(val bool) {
	h.Set("disabled", val)
}
func (h HTMLSelectElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return jsValueToHTMLFormElement(val.JSValue())
}
func (h HTMLSelectElement) Item(args ...interface{}) Element {
	val := h.Call("item", args...)
	return jsValueToElement(val.JSValue())
}
func (h HTMLSelectElement) GetLabels() NodeList {
	val := h.Get("labels")
	return jsValueToNodeList(val.JSValue())
}
func (h HTMLSelectElement) GetLength() float64 {
	val := h.Get("length")
	return val.Float()
}
func (h HTMLSelectElement) SetLength(val float64) {
	h.Set("length", val)
}
func (h HTMLSelectElement) GetMultiple() bool {
	val := h.Get("multiple")
	return val.Bool()
}
func (h HTMLSelectElement) SetMultiple(val bool) {
	h.Set("multiple", val)
}
func (h HTMLSelectElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLSelectElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLSelectElement) NamedItem(args ...interface{}) HTMLOptionElement {
	val := h.Call("namedItem", args...)
	return jsValueToHTMLOptionElement(val.JSValue())
}
func (h HTMLSelectElement) GetOptions() HTMLOptionsCollection {
	val := h.Get("options")
	return jsValueToHTMLOptionsCollection(val.JSValue())
}
func (h HTMLSelectElement) Remove(args ...interface{}) {
	h.Call("remove", args...)
}
func (h HTMLSelectElement) RemoveWithArgs(args ...interface{}) {
	h.Call("removeWithArgs", args...)
}
func (h HTMLSelectElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLSelectElement) GetRequired() bool {
	val := h.Get("required")
	return val.Bool()
}
func (h HTMLSelectElement) SetRequired(val bool) {
	h.Set("required", val)
}
func (h HTMLSelectElement) GetSelectedIndex() float64 {
	val := h.Get("selectedIndex")
	return val.Float()
}
func (h HTMLSelectElement) SetSelectedIndex(val float64) {
	h.Set("selectedIndex", val)
}
func (h HTMLSelectElement) GetSelectedOptions() HTMLCollection {
	val := h.Get("selectedOptions")
	return jsValueToHTMLCollection(val.JSValue())
}
func (h HTMLSelectElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLSelectElement) GetSize() float64 {
	val := h.Get("size")
	return val.Float()
}
func (h HTMLSelectElement) SetSize(val float64) {
	h.Set("size", val)
}
func (h HTMLSelectElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLSelectElement) GetValidationMessage() string {
	val := h.Get("validationMessage")
	return val.String()
}
func (h HTMLSelectElement) GetValidity() ValidityState {
	val := h.Get("validity")
	return jsValueToValidityState(val.JSValue())
}
func (h HTMLSelectElement) GetValue() string {
	val := h.Get("value")
	return val.String()
}
func (h HTMLSelectElement) SetValue(val string) {
	h.Set("value", val)
}
func (h HTMLSelectElement) GetWillValidate() bool {
	val := h.Get("willValidate")
	return val.Bool()
}
