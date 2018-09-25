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
}

func JSValueToHTMLOutputElement(val js.Value) HTMLOutputElement {
	return HTMLOutputElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLOutputElement() HTMLOutputElement { return HTMLOutputElement{Value: v} }
func NewHTMLOutputElement(args ...interface{}) HTMLOutputElement {
	return HTMLOutputElement{Value: JSValueToValue(js.Global().Get("HTMLOutputElement").New(args...))}
}
func (h HTMLOutputElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLOutputElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLOutputElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLOutputElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLOutputElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLOutputElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLOutputElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLOutputElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLOutputElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLOutputElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLOutputElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLOutputElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLOutputElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLOutputElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLOutputElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLOutputElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLOutputElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLOutputElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLOutputElement) GetDefaultValue() string {
	val := h.Get("defaultValue")
	return val.String()
}
func (h HTMLOutputElement) SetDefaultValue(val string) {
	h.Set("defaultValue", val)
}
func (h HTMLOutputElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLOutputElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLOutputElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLOutputElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLOutputElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLOutputElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLOutputElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLOutputElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLOutputElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLOutputElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOutputElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOutputElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLOutputElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLOutputElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLOutputElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLOutputElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLOutputElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLOutputElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLOutputElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLOutputElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLOutputElement) GetHtmlFor() DOMTokenList {
	val := h.Get("htmlFor")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLOutputElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLOutputElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLOutputElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLOutputElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLOutputElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLOutputElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLOutputElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLOutputElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLOutputElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLOutputElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLOutputElement) GetLabels() NodeList {
	val := h.Get("labels")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLOutputElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLOutputElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLOutputElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLOutputElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLOutputElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLOutputElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLOutputElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLOutputElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLOutputElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLOutputElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLOutputElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLOutputElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLOutputElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLOutputElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLOutputElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLOutputElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLOutputElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLOutputElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLOutputElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLOutputElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOutputElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLOutputElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOutputElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLOutputElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLOutputElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLOutputElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOutputElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOutputElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLOutputElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLOutputElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLOutputElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLOutputElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLOutputElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLOutputElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLOutputElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLOutputElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLOutputElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLOutputElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLOutputElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLOutputElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLOutputElement) SetTranslate(val bool) {
	h.Set("translate", val)
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
func (h HTMLOutputElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLOutputElement) GetWillValidate() bool {
	val := h.Get("willValidate")
	return val.Bool()
}
