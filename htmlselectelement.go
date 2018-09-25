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
	GetLength() int
	SetLength(int)
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
	GetSelectedIndex() int
	SetSelectedIndex(int)
	GetSelectedOptions() HTMLCollection
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	SetCustomValidity(args ...interface{})
	GetShadowRoot() ShadowRoot
	GetSize() int
	SetSize(int)
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
}

func JSValueToHTMLSelectElement(val js.Value) HTMLSelectElement {
	return HTMLSelectElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLSelectElement() HTMLSelectElement { return HTMLSelectElement{Value: v} }
func NewHTMLSelectElement(args ...interface{}) HTMLSelectElement {
	return HTMLSelectElement{Value: JSValueToValue(js.Global().Get("HTMLSelectElement").New(args...))}
}
func (h HTMLSelectElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLSelectElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLSelectElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLSelectElement) Add(args ...interface{}) {
	h.Call("add", args...)
}
func (h HTMLSelectElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLSelectElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSelectElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLSelectElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLSelectElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLSelectElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
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
func (h HTMLSelectElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLSelectElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLSelectElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLSelectElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLSelectElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLSelectElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLSelectElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLSelectElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSelectElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLSelectElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLSelectElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLSelectElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLSelectElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLSelectElement) GetDisabled() bool {
	val := h.Get("disabled")
	return val.Bool()
}
func (h HTMLSelectElement) SetDisabled(val bool) {
	h.Set("disabled", val)
}
func (h HTMLSelectElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLSelectElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLSelectElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLSelectElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLSelectElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLSelectElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLSelectElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLSelectElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLSelectElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLSelectElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLSelectElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLSelectElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLSelectElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLSelectElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSelectElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLSelectElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLSelectElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLSelectElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLSelectElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLSelectElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLSelectElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLSelectElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLSelectElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLSelectElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLSelectElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLSelectElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLSelectElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSelectElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLSelectElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLSelectElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLSelectElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLSelectElement) Item(args ...interface{}) Element {
	val := h.Call("item", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLSelectElement) GetLabels() NodeList {
	val := h.Get("labels")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLSelectElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLSelectElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLSelectElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLSelectElement) GetLength() int {
	val := h.Get("length")
	return val.Int()
}
func (h HTMLSelectElement) SetLength(val int) {
	h.Set("length", val)
}
func (h HTMLSelectElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLSelectElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLSelectElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLSelectElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
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
	return JSValueToHTMLOptionElement(val.JSValue())
}
func (h HTMLSelectElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLSelectElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLSelectElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLSelectElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLSelectElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLSelectElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLSelectElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLSelectElement) GetOptions() HTMLOptionsCollection {
	val := h.Get("options")
	return JSValueToHTMLOptionsCollection(val.JSValue())
}
func (h HTMLSelectElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLSelectElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLSelectElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLSelectElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLSelectElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLSelectElement) Remove(args ...interface{}) {
	h.Call("remove", args...)
}
func (h HTMLSelectElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLSelectElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLSelectElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLSelectElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSelectElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLSelectElement) RemoveWithArgs(args ...interface{}) {
	h.Call("removeWithArgs", args...)
}
func (h HTMLSelectElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
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
func (h HTMLSelectElement) GetSelectedIndex() int {
	val := h.Get("selectedIndex")
	return val.Int()
}
func (h HTMLSelectElement) SetSelectedIndex(val int) {
	h.Set("selectedIndex", val)
}
func (h HTMLSelectElement) GetSelectedOptions() HTMLCollection {
	val := h.Get("selectedOptions")
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLSelectElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLSelectElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLSelectElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLSelectElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLSelectElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLSelectElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLSelectElement) GetSize() int {
	val := h.Get("size")
	return val.Int()
}
func (h HTMLSelectElement) SetSize(val int) {
	h.Set("size", val)
}
func (h HTMLSelectElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLSelectElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLSelectElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLSelectElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLSelectElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLSelectElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLSelectElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLSelectElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLSelectElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLSelectElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLSelectElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLSelectElement) SetTranslate(val bool) {
	h.Set("translate", val)
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
	return JSValueToValidityState(val.JSValue())
}
func (h HTMLSelectElement) GetValue() string {
	val := h.Get("value")
	return val.String()
}
func (h HTMLSelectElement) SetValue(val string) {
	h.Set("value", val)
}
func (h HTMLSelectElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLSelectElement) GetWillValidate() bool {
	val := h.Get("willValidate")
	return val.Bool()
}
