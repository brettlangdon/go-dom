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
	GetCols() int
	SetCols(int)
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
	GetMaxLength() int
	SetMaxLength(int)
	GetMinLength() int
	SetMinLength(int)
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
	GetRows() int
	SetRows(int)
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
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTextLength() int
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
}

func JSValueToHTMLTextAreaElement(val js.Value) HTMLTextAreaElement {
	return HTMLTextAreaElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLTextAreaElement() HTMLTextAreaElement { return HTMLTextAreaElement{Value: v} }
func NewHTMLTextAreaElement(args ...interface{}) HTMLTextAreaElement {
	return HTMLTextAreaElement{Value: JSValueToValue(js.Global().Get("HTMLTextAreaElement").New(args...))}
}
func (h HTMLTextAreaElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLTextAreaElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLTextAreaElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLTextAreaElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLTextAreaElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTextAreaElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTextAreaElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLTextAreaElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLTextAreaElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
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
func (h HTMLTextAreaElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLTextAreaElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLTextAreaElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLTextAreaElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLTextAreaElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLTextAreaElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLTextAreaElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTextAreaElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTextAreaElement) GetCols() int {
	val := h.Get("cols")
	return val.Int()
}
func (h HTMLTextAreaElement) SetCols(val int) {
	h.Set("cols", val)
}
func (h HTMLTextAreaElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLTextAreaElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) GetDefaultValue() string {
	val := h.Get("defaultValue")
	return val.String()
}
func (h HTMLTextAreaElement) SetDefaultValue(val string) {
	h.Set("defaultValue", val)
}
func (h HTMLTextAreaElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLTextAreaElement) SetDir(val string) {
	h.Set("dir", val)
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
func (h HTMLTextAreaElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLTextAreaElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLTextAreaElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTextAreaElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLTextAreaElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLTextAreaElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLTextAreaElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLTextAreaElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTextAreaElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTextAreaElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTextAreaElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTextAreaElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTextAreaElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTextAreaElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLTextAreaElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLTextAreaElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLTextAreaElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLTextAreaElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLTextAreaElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLTextAreaElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTextAreaElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLTextAreaElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTextAreaElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLTextAreaElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) GetLabels() NodeList {
	val := h.Get("labels")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLTextAreaElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLTextAreaElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLTextAreaElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTextAreaElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLTextAreaElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLTextAreaElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLTextAreaElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) GetMaxLength() int {
	val := h.Get("maxLength")
	return val.Int()
}
func (h HTMLTextAreaElement) SetMaxLength(val int) {
	h.Set("maxLength", val)
}
func (h HTMLTextAreaElement) GetMinLength() int {
	val := h.Get("minLength")
	return val.Int()
}
func (h HTMLTextAreaElement) SetMinLength(val int) {
	h.Set("minLength", val)
}
func (h HTMLTextAreaElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLTextAreaElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLTextAreaElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLTextAreaElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTextAreaElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLTextAreaElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLTextAreaElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLTextAreaElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLTextAreaElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLTextAreaElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLTextAreaElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLTextAreaElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTextAreaElement) GetPlaceholder() string {
	val := h.Get("placeholder")
	return val.String()
}
func (h HTMLTextAreaElement) SetPlaceholder(val string) {
	h.Set("placeholder", val)
}
func (h HTMLTextAreaElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLTextAreaElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTextAreaElement) GetReadOnly() bool {
	val := h.Get("readOnly")
	return val.Bool()
}
func (h HTMLTextAreaElement) SetReadOnly(val bool) {
	h.Set("readOnly", val)
}
func (h HTMLTextAreaElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLTextAreaElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLTextAreaElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTextAreaElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTextAreaElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLTextAreaElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
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
func (h HTMLTextAreaElement) GetRows() int {
	val := h.Get("rows")
	return val.Int()
}
func (h HTMLTextAreaElement) SetRows(val int) {
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
func (h HTMLTextAreaElement) GetSelectionEnd() int {
	val := h.Get("selectionEnd")
	return val.Int()
}
func (h HTMLTextAreaElement) SetSelectionEnd(val int) {
	h.Set("selectionEnd", val)
}
func (h HTMLTextAreaElement) GetSelectionStart() int {
	val := h.Get("selectionStart")
	return val.Int()
}
func (h HTMLTextAreaElement) SetSelectionStart(val int) {
	h.Set("selectionStart", val)
}
func (h HTMLTextAreaElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLTextAreaElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLTextAreaElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTextAreaElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
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
func (h HTMLTextAreaElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTextAreaElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLTextAreaElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLTextAreaElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLTextAreaElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLTextAreaElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLTextAreaElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLTextAreaElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLTextAreaElement) GetTextLength() int {
	val := h.Get("textLength")
	return val.Int()
}
func (h HTMLTextAreaElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLTextAreaElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLTextAreaElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLTextAreaElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLTextAreaElement) SetTranslate(val bool) {
	h.Set("translate", val)
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
func (h HTMLTextAreaElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
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
