// Code generated DO NOT EDIT
// htmllinkelement.go
package dom

import "syscall/js"

type HTMLLinkElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetAs() string
	SetAs(string)
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
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
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
	GetHref() string
	SetHref(string)
	GetHreflang() string
	SetHreflang(string)
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
	GetMedia() string
	SetMedia(string)
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
	GetReferrerPolicy() string
	SetReferrerPolicy(string)
	GetRel() string
	SetRel(string)
	GetRelList() DOMTokenList
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetRev() string
	SetRev(string)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSheet() StyleSheet
	GetSizes() DOMTokenList
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
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
	GetType() string
	SetType(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLLinkElement struct {
	Value
}

func JSValueToHTMLLinkElement(val js.Value) HTMLLinkElement {
	return HTMLLinkElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLLinkElement() HTMLLinkElement { return HTMLLinkElement{Value: v} }
func NewHTMLLinkElement(args ...interface{}) HTMLLinkElement {
	return HTMLLinkElement{Value: JSValueToValue(js.Global().Get("HTMLLinkElement").New(args...))}
}
func (h HTMLLinkElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLLinkElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLLinkElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLLinkElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLLinkElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) GetAs() string {
	val := h.Get("as")
	return val.String()
}
func (h HTMLLinkElement) SetAs(val string) {
	h.Set("as", val)
}
func (h HTMLLinkElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLLinkElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLLinkElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLLinkElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLLinkElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLLinkElement) GetCharset() string {
	val := h.Get("charset")
	return val.String()
}
func (h HTMLLinkElement) SetCharset(val string) {
	h.Set("charset", val)
}
func (h HTMLLinkElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLLinkElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLLinkElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLLinkElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLLinkElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLLinkElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLLinkElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLLinkElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLLinkElement) GetCrossOrigin() string {
	val := h.Get("crossOrigin")
	return val.String()
}
func (h HTMLLinkElement) SetCrossOrigin(val string) {
	h.Set("crossOrigin", val)
}
func (h HTMLLinkElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLLinkElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLLinkElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLLinkElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLLinkElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLLinkElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLLinkElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLLinkElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLLinkElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLLinkElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLLinkElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLLinkElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLLinkElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLLinkElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLLinkElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLLinkElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLLinkElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLLinkElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLLinkElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLLinkElement) GetHref() string {
	val := h.Get("href")
	return val.String()
}
func (h HTMLLinkElement) SetHref(val string) {
	h.Set("href", val)
}
func (h HTMLLinkElement) GetHreflang() string {
	val := h.Get("hreflang")
	return val.String()
}
func (h HTMLLinkElement) SetHreflang(val string) {
	h.Set("hreflang", val)
}
func (h HTMLLinkElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLLinkElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLLinkElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLLinkElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLLinkElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLLinkElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLLinkElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) GetIntegrity() string {
	val := h.Get("integrity")
	return val.String()
}
func (h HTMLLinkElement) SetIntegrity(val string) {
	h.Set("integrity", val)
}
func (h HTMLLinkElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLLinkElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLLinkElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLLinkElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLLinkElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLLinkElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLLinkElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLLinkElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLLinkElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLLinkElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLLinkElement) GetMedia() string {
	val := h.Get("media")
	return val.String()
}
func (h HTMLLinkElement) SetMedia(val string) {
	h.Set("media", val)
}
func (h HTMLLinkElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLLinkElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLLinkElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLLinkElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLLinkElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLLinkElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLLinkElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLLinkElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLLinkElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLLinkElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLLinkElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLLinkElement) GetRel() string {
	val := h.Get("rel")
	return val.String()
}
func (h HTMLLinkElement) SetRel(val string) {
	h.Set("rel", val)
}
func (h HTMLLinkElement) GetRelList() DOMTokenList {
	val := h.Get("relList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLLinkElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLLinkElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLLinkElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLLinkElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLLinkElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLLinkElement) GetRev() string {
	val := h.Get("rev")
	return val.String()
}
func (h HTMLLinkElement) SetRev(val string) {
	h.Set("rev", val)
}
func (h HTMLLinkElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLLinkElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLLinkElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLLinkElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLLinkElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLLinkElement) GetSheet() StyleSheet {
	val := h.Get("sheet")
	return JSValueToStyleSheet(val.JSValue())
}
func (h HTMLLinkElement) GetSizes() DOMTokenList {
	val := h.Get("sizes")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLLinkElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLLinkElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLLinkElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLLinkElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLLinkElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLLinkElement) GetTarget() string {
	val := h.Get("target")
	return val.String()
}
func (h HTMLLinkElement) SetTarget(val string) {
	h.Set("target", val)
}
func (h HTMLLinkElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLLinkElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLLinkElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLLinkElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLLinkElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLLinkElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLLinkElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLLinkElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLLinkElement) SetType(val string) {
	h.Set("type", val)
}
func (h HTMLLinkElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
