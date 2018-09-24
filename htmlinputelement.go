// Code generated DO NOT EDIT
// htmlinputelement.go
package dom

import "syscall/js"

type HTMLInputElementIFace interface {
	GetAccept() string
	SetAccept(string)
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	GetAlign() string
	SetAlign(string)
	GetAlt() string
	SetAlt(string)
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
	GetChecked() bool
	SetChecked(bool)
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	GetDefaultChecked() bool
	SetDefaultChecked(bool)
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
	GetFiles() FileList
	SetFiles(FileList)
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
	GetHeight() int
	SetHeight(int)
	GetHidden() bool
	SetHidden(bool)
	GetId() string
	SetId(string)
	GetIndeterminate() bool
	SetIndeterminate(bool)
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
	GetList() HTMLElement
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
	GetMax() string
	SetMax(string)
	GetMaxLength() int
	SetMaxLength(int)
	GetMin() string
	SetMin(string)
	GetMinLength() int
	SetMinLength(int)
	GetMultiple() bool
	SetMultiple(bool)
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
	GetPattern() string
	SetPattern(string)
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
	Select(args ...interface{})
	GetSelectionDirection() string
	SetSelectionDirection(string)
	GetSelectionEnd() int
	SetSelectionEnd(int)
	GetSelectionStart() int
	SetSelectionStart(int)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	SetCustomValidity(args ...interface{})
	SetRangeText(args ...interface{})
	SetRangeTextWithArgs(args ...interface{})
	SetSelectionRange(args ...interface{})
	GetShadowRoot() ShadowRoot
	GetSize() int
	SetSize(int)
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetSrc() string
	SetSrc(string)
	GetStep() string
	SetStep(string)
	StepDown(args ...interface{})
	StepUp(args ...interface{})
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
	GetUseMap() string
	SetUseMap(string)
	GetValidationMessage() string
	GetValidity() ValidityState
	GetValue() string
	SetValue(string)
	GetValueAsDate() Value
	SetValueAsDate(Value)
	GetValueAsNumber() float64
	SetValueAsNumber(float64)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() int
	SetWidth(int)
	GetWillValidate() bool
}
type HTMLInputElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLInputElement(val js.Value) HTMLInputElement {
	return HTMLInputElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLInputElement() HTMLInputElement { return HTMLInputElement{Value: v} }
func NewHTMLInputElement(args ...interface{}) HTMLInputElement {
	return HTMLInputElement{Value: JSValueToValue(js.Global().Get("HTMLInputElement").New(args...))}
}
func (h HTMLInputElement) GetAccept() string {
	val := h.Get("accept")
	return val.String()
}
func (h HTMLInputElement) SetAccept(val string) {
	h.Set("accept", val)
}
func (h HTMLInputElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLInputElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLInputElement) GetAlt() string {
	val := h.Get("alt")
	return val.String()
}
func (h HTMLInputElement) SetAlt(val string) {
	h.Set("alt", val)
}
func (h HTMLInputElement) GetAutocomplete() string {
	val := h.Get("autocomplete")
	return val.String()
}
func (h HTMLInputElement) SetAutocomplete(val string) {
	h.Set("autocomplete", val)
}
func (h HTMLInputElement) GetAutofocus() bool {
	val := h.Get("autofocus")
	return val.Bool()
}
func (h HTMLInputElement) SetAutofocus(val bool) {
	h.Set("autofocus", val)
}
func (h HTMLInputElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLInputElement) GetChecked() bool {
	val := h.Get("checked")
	return val.Bool()
}
func (h HTMLInputElement) SetChecked(val bool) {
	h.Set("checked", val)
}
func (h HTMLInputElement) GetDefaultChecked() bool {
	val := h.Get("defaultChecked")
	return val.Bool()
}
func (h HTMLInputElement) SetDefaultChecked(val bool) {
	h.Set("defaultChecked", val)
}
func (h HTMLInputElement) GetDefaultValue() string {
	val := h.Get("defaultValue")
	return val.String()
}
func (h HTMLInputElement) SetDefaultValue(val string) {
	h.Set("defaultValue", val)
}
func (h HTMLInputElement) GetDirName() string {
	val := h.Get("dirName")
	return val.String()
}
func (h HTMLInputElement) SetDirName(val string) {
	h.Set("dirName", val)
}
func (h HTMLInputElement) GetDisabled() bool {
	val := h.Get("disabled")
	return val.Bool()
}
func (h HTMLInputElement) SetDisabled(val bool) {
	h.Set("disabled", val)
}
func (h HTMLInputElement) GetFiles() FileList {
	val := h.Get("files")
	return JSValueToFileList(val.JSValue())
}
func (h HTMLInputElement) SetFiles(val FileList) {
	h.Set("files", val)
}
func (h HTMLInputElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLInputElement) GetFormAction() string {
	val := h.Get("formAction")
	return val.String()
}
func (h HTMLInputElement) SetFormAction(val string) {
	h.Set("formAction", val)
}
func (h HTMLInputElement) GetFormEnctype() string {
	val := h.Get("formEnctype")
	return val.String()
}
func (h HTMLInputElement) SetFormEnctype(val string) {
	h.Set("formEnctype", val)
}
func (h HTMLInputElement) GetFormMethod() string {
	val := h.Get("formMethod")
	return val.String()
}
func (h HTMLInputElement) SetFormMethod(val string) {
	h.Set("formMethod", val)
}
func (h HTMLInputElement) GetFormNoValidate() bool {
	val := h.Get("formNoValidate")
	return val.Bool()
}
func (h HTMLInputElement) SetFormNoValidate(val bool) {
	h.Set("formNoValidate", val)
}
func (h HTMLInputElement) GetFormTarget() string {
	val := h.Get("formTarget")
	return val.String()
}
func (h HTMLInputElement) SetFormTarget(val string) {
	h.Set("formTarget", val)
}
func (h HTMLInputElement) GetHeight() int {
	val := h.Get("height")
	return val.Int()
}
func (h HTMLInputElement) SetHeight(val int) {
	h.Set("height", val)
}
func (h HTMLInputElement) GetIndeterminate() bool {
	val := h.Get("indeterminate")
	return val.Bool()
}
func (h HTMLInputElement) SetIndeterminate(val bool) {
	h.Set("indeterminate", val)
}
func (h HTMLInputElement) GetLabels() NodeList {
	val := h.Get("labels")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLInputElement) GetList() HTMLElement {
	val := h.Get("list")
	return JSValueToHTMLElement(val.JSValue())
}
func (h HTMLInputElement) GetMax() string {
	val := h.Get("max")
	return val.String()
}
func (h HTMLInputElement) SetMax(val string) {
	h.Set("max", val)
}
func (h HTMLInputElement) GetMaxLength() int {
	val := h.Get("maxLength")
	return val.Int()
}
func (h HTMLInputElement) SetMaxLength(val int) {
	h.Set("maxLength", val)
}
func (h HTMLInputElement) GetMin() string {
	val := h.Get("min")
	return val.String()
}
func (h HTMLInputElement) SetMin(val string) {
	h.Set("min", val)
}
func (h HTMLInputElement) GetMinLength() int {
	val := h.Get("minLength")
	return val.Int()
}
func (h HTMLInputElement) SetMinLength(val int) {
	h.Set("minLength", val)
}
func (h HTMLInputElement) GetMultiple() bool {
	val := h.Get("multiple")
	return val.Bool()
}
func (h HTMLInputElement) SetMultiple(val bool) {
	h.Set("multiple", val)
}
func (h HTMLInputElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLInputElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLInputElement) GetPattern() string {
	val := h.Get("pattern")
	return val.String()
}
func (h HTMLInputElement) SetPattern(val string) {
	h.Set("pattern", val)
}
func (h HTMLInputElement) GetPlaceholder() string {
	val := h.Get("placeholder")
	return val.String()
}
func (h HTMLInputElement) SetPlaceholder(val string) {
	h.Set("placeholder", val)
}
func (h HTMLInputElement) GetReadOnly() bool {
	val := h.Get("readOnly")
	return val.Bool()
}
func (h HTMLInputElement) SetReadOnly(val bool) {
	h.Set("readOnly", val)
}
func (h HTMLInputElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLInputElement) GetRequired() bool {
	val := h.Get("required")
	return val.Bool()
}
func (h HTMLInputElement) SetRequired(val bool) {
	h.Set("required", val)
}
func (h HTMLInputElement) Select(args ...interface{}) {
	h.Call("select", args...)
}
func (h HTMLInputElement) GetSelectionDirection() string {
	val := h.Get("selectionDirection")
	return val.String()
}
func (h HTMLInputElement) SetSelectionDirection(val string) {
	h.Set("selectionDirection", val)
}
func (h HTMLInputElement) GetSelectionEnd() int {
	val := h.Get("selectionEnd")
	return val.Int()
}
func (h HTMLInputElement) SetSelectionEnd(val int) {
	h.Set("selectionEnd", val)
}
func (h HTMLInputElement) GetSelectionStart() int {
	val := h.Get("selectionStart")
	return val.Int()
}
func (h HTMLInputElement) SetSelectionStart(val int) {
	h.Set("selectionStart", val)
}
func (h HTMLInputElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLInputElement) SetRangeText(args ...interface{}) {
	h.Call("setRangeText", args...)
}
func (h HTMLInputElement) SetRangeTextWithArgs(args ...interface{}) {
	h.Call("setRangeTextWithArgs", args...)
}
func (h HTMLInputElement) SetSelectionRange(args ...interface{}) {
	h.Call("setSelectionRange", args...)
}
func (h HTMLInputElement) GetSize() int {
	val := h.Get("size")
	return val.Int()
}
func (h HTMLInputElement) SetSize(val int) {
	h.Set("size", val)
}
func (h HTMLInputElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLInputElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLInputElement) GetStep() string {
	val := h.Get("step")
	return val.String()
}
func (h HTMLInputElement) SetStep(val string) {
	h.Set("step", val)
}
func (h HTMLInputElement) StepDown(args ...interface{}) {
	h.Call("stepDown", args...)
}
func (h HTMLInputElement) StepUp(args ...interface{}) {
	h.Call("stepUp", args...)
}
func (h HTMLInputElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLInputElement) SetType(val string) {
	h.Set("type", val)
}
func (h HTMLInputElement) GetUseMap() string {
	val := h.Get("useMap")
	return val.String()
}
func (h HTMLInputElement) SetUseMap(val string) {
	h.Set("useMap", val)
}
func (h HTMLInputElement) GetValidationMessage() string {
	val := h.Get("validationMessage")
	return val.String()
}
func (h HTMLInputElement) GetValidity() ValidityState {
	val := h.Get("validity")
	return JSValueToValidityState(val.JSValue())
}
func (h HTMLInputElement) GetValue() string {
	val := h.Get("value")
	return val.String()
}
func (h HTMLInputElement) SetValue(val string) {
	h.Set("value", val)
}
func (h HTMLInputElement) GetValueAsDate() Value {
	val := h.Get("valueAsDate")
	return val
}
func (h HTMLInputElement) SetValueAsDate(val Value) {
	h.Set("valueAsDate", val)
}
func (h HTMLInputElement) GetValueAsNumber() float64 {
	val := h.Get("valueAsNumber")
	return val.Float()
}
func (h HTMLInputElement) SetValueAsNumber(val float64) {
	h.Set("valueAsNumber", val)
}
func (h HTMLInputElement) GetWidth() int {
	val := h.Get("width")
	return val.Int()
}
func (h HTMLInputElement) SetWidth(val int) {
	h.Set("width", val)
}
func (h HTMLInputElement) GetWillValidate() bool {
	val := h.Get("willValidate")
	return val.Bool()
}
