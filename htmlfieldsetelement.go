// Code generated DO NOT EDIT
// htmlfieldsetelement.go
package dom

import "syscall/js"

type HTMLFieldSetElementIFace interface {
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
	GetDir() string
	SetDir(string)
	GetDisabled() bool
	SetDisabled(bool)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetElements() HTMLCollection
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
	WebkitMatchesSelector(args ...interface{}) bool
	GetWillValidate() bool
}
type HTMLFieldSetElement struct {
	Value
}

func JSValueToHTMLFieldSetElement(val js.Value) HTMLFieldSetElement {
	return HTMLFieldSetElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLFieldSetElement() HTMLFieldSetElement { return HTMLFieldSetElement{Value: v} }
func NewHTMLFieldSetElement(args ...interface{}) HTMLFieldSetElement {
	return HTMLFieldSetElement{Value: JSValueToValue(js.Global().Get("HTMLFieldSetElement").New(args...))}
}
func (h HTMLFieldSetElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLFieldSetElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLFieldSetElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLFieldSetElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLFieldSetElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLFieldSetElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLFieldSetElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLFieldSetElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLFieldSetElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLFieldSetElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLFieldSetElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLFieldSetElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLFieldSetElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLFieldSetElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLFieldSetElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLFieldSetElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLFieldSetElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLFieldSetElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLFieldSetElement) GetDisabled() bool {
	val := h.Get("disabled")
	return val.Bool()
}
func (h HTMLFieldSetElement) SetDisabled(val bool) {
	h.Set("disabled", val)
}
func (h HTMLFieldSetElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLFieldSetElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLFieldSetElement) GetElements() HTMLCollection {
	val := h.Get("elements")
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFieldSetElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLFieldSetElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLFieldSetElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLFieldSetElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLFieldSetElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFieldSetElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFieldSetElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFieldSetElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFieldSetElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFieldSetElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLFieldSetElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLFieldSetElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLFieldSetElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLFieldSetElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLFieldSetElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLFieldSetElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLFieldSetElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLFieldSetElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLFieldSetElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLFieldSetElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLFieldSetElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLFieldSetElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLFieldSetElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLFieldSetElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLFieldSetElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLFieldSetElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLFieldSetElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLFieldSetElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLFieldSetElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLFieldSetElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLFieldSetElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLFieldSetElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLFieldSetElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLFieldSetElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLFieldSetElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLFieldSetElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLFieldSetElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFieldSetElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLFieldSetElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFieldSetElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLFieldSetElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLFieldSetElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFieldSetElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFieldSetElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLFieldSetElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLFieldSetElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLFieldSetElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLFieldSetElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLFieldSetElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLFieldSetElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLFieldSetElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLFieldSetElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLFieldSetElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLFieldSetElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLFieldSetElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLFieldSetElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLFieldSetElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLFieldSetElement) GetValidationMessage() string {
	val := h.Get("validationMessage")
	return val.String()
}
func (h HTMLFieldSetElement) GetValidity() ValidityState {
	val := h.Get("validity")
	return JSValueToValidityState(val.JSValue())
}
func (h HTMLFieldSetElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLFieldSetElement) GetWillValidate() bool {
	val := h.Get("willValidate")
	return val.Bool()
}
