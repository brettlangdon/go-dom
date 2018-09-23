// Code generated DO NOT EDIT
// htmltextareaelement.go
package dom

import "syscall/js"

type HTMLTextAreaElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
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
	GetCols() float64
	SetCols(float64)
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	GetDefaultValue() string
	SetDefaultValue(string)
	GetDir() string
	SetDir(string)
	GetDirName() string
	SetDirName(string)
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
	GetLabels() NodeList
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
	GetMaxLength() float64
	SetMaxLength(float64)
	GetMinLength() float64
	SetMinLength(float64)
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
	GetPlaceholder() string
	SetPlaceholder(string)
	GetPrefix() string
	GetPreviousSibling() Node
	GetReadOnly() bool
	SetReadOnly(bool)
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	ReportValidity(args ...interface{}) bool
	GetRequired() bool
	SetRequired(bool)
	GetRows() float64
	SetRows(float64)
	Select(args ...interface{})
	GetSelectionDirection() string
	SetSelectionDirection(string)
	GetSelectionEnd() float64
	SetSelectionEnd(float64)
	GetSelectionStart() float64
	SetSelectionStart(float64)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	SetCustomValidity(args ...interface{})
	SetRangeText(args ...interface{})
	SetRangeTextWithArgs(args ...interface{})
	SetSelectionRange(args ...interface{})
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTextLength() float64
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
	GetWrap() string
	SetWrap(string)
}
type HTMLTextAreaElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLTextAreaElement(val js.Value) HTMLTextAreaElement {
	return HTMLTextAreaElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLTextAreaElement() HTMLTextAreaElement { return HTMLTextAreaElement{Value: v} }
func (h HTMLTextAreaElement) GetAutocomplete() string {
	val := h.Get("autocomplete")
	return val.String()
}
func (h HTMLTextAreaElement) SetAutocomplete(val string) {
	h.Set("autocomplete", val)
}
func (h HTMLTextAreaElement) GetAutofocus() bool {
	val := h.Get("autofocus")
	return val.Bool()
}
func (h HTMLTextAreaElement) SetAutofocus(val bool) {
	h.Set("autofocus", val)
}
func (h HTMLTextAreaElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) GetCols() float64 {
	val := h.Get("cols")
	return val.Float()
}
func (h HTMLTextAreaElement) SetCols(val float64) {
	h.Set("cols", val)
}
func (h HTMLTextAreaElement) GetDefaultValue() string {
	val := h.Get("defaultValue")
	return val.String()
}
func (h HTMLTextAreaElement) SetDefaultValue(val string) {
	h.Set("defaultValue", val)
}
func (h HTMLTextAreaElement) GetDirName() string {
	val := h.Get("dirName")
	return val.String()
}
func (h HTMLTextAreaElement) SetDirName(val string) {
	h.Set("dirName", val)
}
func (h HTMLTextAreaElement) GetDisabled() bool {
	val := h.Get("disabled")
	return val.Bool()
}
func (h HTMLTextAreaElement) SetDisabled(val bool) {
	h.Set("disabled", val)
}
func (h HTMLTextAreaElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLTextAreaElement) GetLabels() NodeList {
	val := h.Get("labels")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLTextAreaElement) GetMaxLength() float64 {
	val := h.Get("maxLength")
	return val.Float()
}
func (h HTMLTextAreaElement) SetMaxLength(val float64) {
	h.Set("maxLength", val)
}
func (h HTMLTextAreaElement) GetMinLength() float64 {
	val := h.Get("minLength")
	return val.Float()
}
func (h HTMLTextAreaElement) SetMinLength(val float64) {
	h.Set("minLength", val)
}
func (h HTMLTextAreaElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLTextAreaElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLTextAreaElement) GetPlaceholder() string {
	val := h.Get("placeholder")
	return val.String()
}
func (h HTMLTextAreaElement) SetPlaceholder(val string) {
	h.Set("placeholder", val)
}
func (h HTMLTextAreaElement) GetReadOnly() bool {
	val := h.Get("readOnly")
	return val.Bool()
}
func (h HTMLTextAreaElement) SetReadOnly(val bool) {
	h.Set("readOnly", val)
}
func (h HTMLTextAreaElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) GetRequired() bool {
	val := h.Get("required")
	return val.Bool()
}
func (h HTMLTextAreaElement) SetRequired(val bool) {
	h.Set("required", val)
}
func (h HTMLTextAreaElement) GetRows() float64 {
	val := h.Get("rows")
	return val.Float()
}
func (h HTMLTextAreaElement) SetRows(val float64) {
	h.Set("rows", val)
}
func (h HTMLTextAreaElement) Select(args ...interface{}) {
	h.Call("select", args...)
}
func (h HTMLTextAreaElement) GetSelectionDirection() string {
	val := h.Get("selectionDirection")
	return val.String()
}
func (h HTMLTextAreaElement) SetSelectionDirection(val string) {
	h.Set("selectionDirection", val)
}
func (h HTMLTextAreaElement) GetSelectionEnd() float64 {
	val := h.Get("selectionEnd")
	return val.Float()
}
func (h HTMLTextAreaElement) SetSelectionEnd(val float64) {
	h.Set("selectionEnd", val)
}
func (h HTMLTextAreaElement) GetSelectionStart() float64 {
	val := h.Get("selectionStart")
	return val.Float()
}
func (h HTMLTextAreaElement) SetSelectionStart(val float64) {
	h.Set("selectionStart", val)
}
func (h HTMLTextAreaElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLTextAreaElement) SetRangeText(args ...interface{}) {
	h.Call("setRangeText", args...)
}
func (h HTMLTextAreaElement) SetRangeTextWithArgs(args ...interface{}) {
	h.Call("setRangeTextWithArgs", args...)
}
func (h HTMLTextAreaElement) SetSelectionRange(args ...interface{}) {
	h.Call("setSelectionRange", args...)
}
func (h HTMLTextAreaElement) GetTextLength() float64 {
	val := h.Get("textLength")
	return val.Float()
}
func (h HTMLTextAreaElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLTextAreaElement) GetValidationMessage() string {
	val := h.Get("validationMessage")
	return val.String()
}
func (h HTMLTextAreaElement) GetValidity() ValidityState {
	val := h.Get("validity")
	return JSValueToValidityState(val.JSValue())
}
func (h HTMLTextAreaElement) GetValue() string {
	val := h.Get("value")
	return val.String()
}
func (h HTMLTextAreaElement) SetValue(val string) {
	h.Set("value", val)
}
func (h HTMLTextAreaElement) GetWillValidate() bool {
	val := h.Get("willValidate")
	return val.Bool()
}
func (h HTMLTextAreaElement) GetWrap() string {
	val := h.Get("wrap")
	return val.String()
}
func (h HTMLTextAreaElement) SetWrap(val string) {
	h.Set("wrap", val)
}
