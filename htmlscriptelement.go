// Code generated DO NOT EDIT
// htmlscriptelement.go
package dom

import "syscall/js"

type HTMLScriptElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetAsync() bool
	SetAsync(bool)
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetBaseURI() string
	GetCharset() string
	SetCharset(string)
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	GetCrossOrigin() string
	SetCrossOrigin(string)
	GetDefer() bool
	SetDefer(bool)
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetEvent() string
	SetEvent(string)
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
	GetHtmlFor() string
	SetHtmlFor(string)
	GetId() string
	SetId(string)
	GetInnerText() string
	SetInnerText(string)
	InsertAdjacentElement(args ...interface{}) Element
	InsertAdjacentText(args ...interface{})
	InsertBefore(args ...interface{}) Node
	GetIntegrity() string
	SetIntegrity(string)
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
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNoModule() bool
	SetNoModule(bool)
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
	GetReferrerPolicy() string
	SetReferrerPolicy(string)
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetSrc() string
	SetSrc(string)
	GetTagName() string
	GetText() string
	SetText(string)
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetType() string
	SetType(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLScriptElement struct {
	Value
}

func JSValueToHTMLScriptElement(val js.Value) HTMLScriptElement {
	return HTMLScriptElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLScriptElement() HTMLScriptElement { return HTMLScriptElement{Value: v} }
func NewHTMLScriptElement(args ...interface{}) HTMLScriptElement {
	return HTMLScriptElement{Value: JSValueToValue(js.Global().Get("HTMLScriptElement").New(args...))}
}
func (h HTMLScriptElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLScriptElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLScriptElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLScriptElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLScriptElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) GetAsync() bool {
	val := h.Get("async")
	return val.Bool()
}
func (h HTMLScriptElement) SetAsync(val bool) {
	h.Set("async", val)
}
func (h HTMLScriptElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLScriptElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLScriptElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLScriptElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLScriptElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLScriptElement) GetCharset() string {
	val := h.Get("charset")
	return val.String()
}
func (h HTMLScriptElement) SetCharset(val string) {
	h.Set("charset", val)
}
func (h HTMLScriptElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLScriptElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLScriptElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLScriptElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLScriptElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLScriptElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLScriptElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLScriptElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLScriptElement) GetCrossOrigin() string {
	val := h.Get("crossOrigin")
	return val.String()
}
func (h HTMLScriptElement) SetCrossOrigin(val string) {
	h.Set("crossOrigin", val)
}
func (h HTMLScriptElement) GetDefer() bool {
	val := h.Get("defer")
	return val.Bool()
}
func (h HTMLScriptElement) SetDefer(val bool) {
	h.Set("defer", val)
}
func (h HTMLScriptElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLScriptElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLScriptElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLScriptElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLScriptElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLScriptElement) GetEvent() string {
	val := h.Get("event")
	return val.String()
}
func (h HTMLScriptElement) SetEvent(val string) {
	h.Set("event", val)
}
func (h HTMLScriptElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLScriptElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLScriptElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLScriptElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLScriptElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLScriptElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLScriptElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLScriptElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLScriptElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLScriptElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLScriptElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLScriptElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLScriptElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLScriptElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLScriptElement) GetHtmlFor() string {
	val := h.Get("htmlFor")
	return val.String()
}
func (h HTMLScriptElement) SetHtmlFor(val string) {
	h.Set("htmlFor", val)
}
func (h HTMLScriptElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLScriptElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLScriptElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLScriptElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLScriptElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLScriptElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLScriptElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) GetIntegrity() string {
	val := h.Get("integrity")
	return val.String()
}
func (h HTMLScriptElement) SetIntegrity(val string) {
	h.Set("integrity", val)
}
func (h HTMLScriptElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLScriptElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLScriptElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLScriptElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLScriptElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLScriptElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLScriptElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLScriptElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLScriptElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLScriptElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLScriptElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLScriptElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) GetNoModule() bool {
	val := h.Get("noModule")
	return val.Bool()
}
func (h HTMLScriptElement) SetNoModule(val bool) {
	h.Set("noModule", val)
}
func (h HTMLScriptElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLScriptElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLScriptElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLScriptElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLScriptElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLScriptElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLScriptElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLScriptElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLScriptElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLScriptElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLScriptElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLScriptElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLScriptElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLScriptElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLScriptElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLScriptElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLScriptElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLScriptElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLScriptElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLScriptElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLScriptElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLScriptElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLScriptElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLScriptElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLScriptElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLScriptElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLScriptElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLScriptElement) GetText() string {
	val := h.Get("text")
	return val.String()
}
func (h HTMLScriptElement) SetText(val string) {
	h.Set("text", val)
}
func (h HTMLScriptElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLScriptElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLScriptElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLScriptElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLScriptElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLScriptElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLScriptElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLScriptElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLScriptElement) SetType(val string) {
	h.Set("type", val)
}
func (h HTMLScriptElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
