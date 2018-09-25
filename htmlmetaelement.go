// Code generated DO NOT EDIT
// htmlmetaelement.go
package dom

import "syscall/js"

type HTMLMetaElementIFace interface {
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
	GetContent() string
	SetContent(string)
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
	GetHttpEquiv() string
	SetHttpEquiv(string)
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
	GetScheme() string
	SetScheme(string)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
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
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLMetaElement struct {
	Value
}

func JSValueToHTMLMetaElement(val js.Value) HTMLMetaElement {
	return HTMLMetaElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLMetaElement() HTMLMetaElement { return HTMLMetaElement{Value: v} }
func NewHTMLMetaElement(args ...interface{}) HTMLMetaElement {
	return HTMLMetaElement{Value: JSValueToValue(js.Global().Get("HTMLMetaElement").New(args...))}
}
func (h HTMLMetaElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLMetaElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLMetaElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLMetaElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLMetaElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLMetaElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLMetaElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLMetaElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLMetaElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLMetaElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLMetaElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLMetaElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLMetaElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLMetaElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLMetaElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLMetaElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLMetaElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLMetaElement) GetContent() string {
	val := h.Get("content")
	return val.String()
}
func (h HTMLMetaElement) SetContent(val string) {
	h.Set("content", val)
}
func (h HTMLMetaElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLMetaElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLMetaElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLMetaElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLMetaElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLMetaElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLMetaElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLMetaElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLMetaElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMetaElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMetaElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLMetaElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLMetaElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLMetaElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLMetaElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLMetaElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLMetaElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLMetaElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLMetaElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLMetaElement) GetHttpEquiv() string {
	val := h.Get("httpEquiv")
	return val.String()
}
func (h HTMLMetaElement) SetHttpEquiv(val string) {
	h.Set("httpEquiv", val)
}
func (h HTMLMetaElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLMetaElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLMetaElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLMetaElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLMetaElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLMetaElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLMetaElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLMetaElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLMetaElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLMetaElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLMetaElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLMetaElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLMetaElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLMetaElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLMetaElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLMetaElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLMetaElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLMetaElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLMetaElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLMetaElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLMetaElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLMetaElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLMetaElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLMetaElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLMetaElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLMetaElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLMetaElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLMetaElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLMetaElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLMetaElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMetaElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLMetaElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMetaElement) GetScheme() string {
	val := h.Get("scheme")
	return val.String()
}
func (h HTMLMetaElement) SetScheme(val string) {
	h.Set("scheme", val)
}
func (h HTMLMetaElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLMetaElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLMetaElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMetaElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMetaElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLMetaElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLMetaElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLMetaElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLMetaElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLMetaElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLMetaElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLMetaElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLMetaElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLMetaElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLMetaElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLMetaElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLMetaElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLMetaElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
