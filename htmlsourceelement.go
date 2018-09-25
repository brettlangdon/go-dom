// Code generated DO NOT EDIT
// htmlsourceelement.go
package dom

import "syscall/js"

type HTMLSourceElementIFace interface {
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
	GetSizes() string
	SetSizes(string)
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetSrc() string
	SetSrc(string)
	GetSrcset() string
	SetSrcset(string)
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
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLSourceElement struct {
	Value
}

func JSValueToHTMLSourceElement(val js.Value) HTMLSourceElement {
	return HTMLSourceElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLSourceElement() HTMLSourceElement { return HTMLSourceElement{Value: v} }
func NewHTMLSourceElement(args ...interface{}) HTMLSourceElement {
	return HTMLSourceElement{Value: JSValueToValue(js.Global().Get("HTMLSourceElement").New(args...))}
}
func (h HTMLSourceElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLSourceElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLSourceElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLSourceElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLSourceElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLSourceElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLSourceElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLSourceElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLSourceElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLSourceElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLSourceElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLSourceElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLSourceElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLSourceElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLSourceElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLSourceElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLSourceElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLSourceElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLSourceElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLSourceElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLSourceElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLSourceElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLSourceElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLSourceElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLSourceElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLSourceElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLSourceElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLSourceElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLSourceElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLSourceElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLSourceElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLSourceElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLSourceElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLSourceElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLSourceElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLSourceElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLSourceElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLSourceElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLSourceElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLSourceElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLSourceElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLSourceElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLSourceElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLSourceElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLSourceElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLSourceElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLSourceElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLSourceElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLSourceElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLSourceElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLSourceElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLSourceElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLSourceElement) GetMedia() string {
	val := h.Get("media")
	return val.String()
}
func (h HTMLSourceElement) SetMedia(val string) {
	h.Set("media", val)
}
func (h HTMLSourceElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLSourceElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLSourceElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLSourceElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLSourceElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLSourceElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLSourceElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLSourceElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLSourceElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLSourceElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLSourceElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLSourceElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLSourceElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLSourceElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLSourceElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLSourceElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLSourceElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLSourceElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLSourceElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLSourceElement) GetSizes() string {
	val := h.Get("sizes")
	return val.String()
}
func (h HTMLSourceElement) SetSizes(val string) {
	h.Set("sizes", val)
}
func (h HTMLSourceElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLSourceElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLSourceElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLSourceElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLSourceElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLSourceElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLSourceElement) GetSrcset() string {
	val := h.Get("srcset")
	return val.String()
}
func (h HTMLSourceElement) SetSrcset(val string) {
	h.Set("srcset", val)
}
func (h HTMLSourceElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLSourceElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLSourceElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLSourceElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLSourceElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLSourceElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLSourceElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLSourceElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLSourceElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLSourceElement) SetType(val string) {
	h.Set("type", val)
}
func (h HTMLSourceElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
