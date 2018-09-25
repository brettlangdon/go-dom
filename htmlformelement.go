// Code generated DO NOT EDIT
// htmlformelement.go
package dom

import "syscall/js"

type HTMLFormElementIFace interface {
	GetAcceptCharset() string
	SetAcceptCharset(string)
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	GetAction() string
	SetAction(string)
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetAutocomplete() string
	SetAutocomplete(string)
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
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetElements() HTMLFormControlsCollection
	GetEncoding() string
	SetEncoding(string)
	GetEnctype() string
	SetEnctype(string)
	GetFirstChild() Node
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
	GetLength() int
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
	GetMethod() string
	SetMethod(string)
	GetName() string
	SetName(string)
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNoValidate() bool
	SetNoValidate(bool)
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
	Reset(args ...interface{})
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	Submit(args ...interface{})
	GetTagName() string
	GetTarget() string
	SetTarget(string)
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLFormElement struct {
	Value
}

func JSValueToHTMLFormElement(val js.Value) HTMLFormElement {
	return HTMLFormElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLFormElement() HTMLFormElement { return HTMLFormElement{Value: v} }
func NewHTMLFormElement(args ...interface{}) HTMLFormElement {
	return HTMLFormElement{Value: JSValueToValue(js.Global().Get("HTMLFormElement").New(args...))}
}
func (h HTMLFormElement) GetAcceptCharset() string {
	val := h.Get("acceptCharset")
	return val.String()
}
func (h HTMLFormElement) SetAcceptCharset(val string) {
	h.Set("acceptCharset", val)
}
func (h HTMLFormElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLFormElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLFormElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLFormElement) GetAction() string {
	val := h.Get("action")
	return val.String()
}
func (h HTMLFormElement) SetAction(val string) {
	h.Set("action", val)
}
func (h HTMLFormElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLFormElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLFormElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLFormElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLFormElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLFormElement) GetAutocomplete() string {
	val := h.Get("autocomplete")
	return val.String()
}
func (h HTMLFormElement) SetAutocomplete(val string) {
	h.Set("autocomplete", val)
}
func (h HTMLFormElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLFormElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLFormElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLFormElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLFormElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLFormElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLFormElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLFormElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLFormElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLFormElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLFormElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLFormElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLFormElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLFormElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLFormElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLFormElement) GetElements() HTMLFormControlsCollection {
	val := h.Get("elements")
	return JSValueToHTMLFormControlsCollection(val.JSValue())
}
func (h HTMLFormElement) GetEncoding() string {
	val := h.Get("encoding")
	return val.String()
}
func (h HTMLFormElement) SetEncoding(val string) {
	h.Set("encoding", val)
}
func (h HTMLFormElement) GetEnctype() string {
	val := h.Get("enctype")
	return val.String()
}
func (h HTMLFormElement) SetEnctype(val string) {
	h.Set("enctype", val)
}
func (h HTMLFormElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLFormElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLFormElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLFormElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFormElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFormElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFormElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFormElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFormElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLFormElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLFormElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLFormElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLFormElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLFormElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLFormElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLFormElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLFormElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLFormElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLFormElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLFormElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLFormElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLFormElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLFormElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLFormElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLFormElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLFormElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLFormElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) GetLength() int {
	val := h.Get("length")
	return val.Int()
}
func (h HTMLFormElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLFormElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLFormElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLFormElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLFormElement) GetMethod() string {
	val := h.Get("method")
	return val.String()
}
func (h HTMLFormElement) SetMethod(val string) {
	h.Set("method", val)
}
func (h HTMLFormElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLFormElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLFormElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLFormElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) GetNoValidate() bool {
	val := h.Get("noValidate")
	return val.Bool()
}
func (h HTMLFormElement) SetNoValidate(val bool) {
	h.Set("noValidate", val)
}
func (h HTMLFormElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLFormElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLFormElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLFormElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLFormElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLFormElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLFormElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLFormElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLFormElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLFormElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLFormElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFormElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLFormElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFormElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLFormElement) Reset(args ...interface{}) {
	h.Call("reset", args...)
}
func (h HTMLFormElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLFormElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLFormElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFormElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFormElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLFormElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLFormElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLFormElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLFormElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLFormElement) Submit(args ...interface{}) {
	h.Call("submit", args...)
}
func (h HTMLFormElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLFormElement) GetTarget() string {
	val := h.Get("target")
	return val.String()
}
func (h HTMLFormElement) SetTarget(val string) {
	h.Set("target", val)
}
func (h HTMLFormElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLFormElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLFormElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLFormElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLFormElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLFormElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLFormElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLFormElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
