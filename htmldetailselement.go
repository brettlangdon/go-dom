// Code generated DO NOT EDIT
// htmldetailselement.go
package dom

import "syscall/js"

type HTMLDetailsElementIFace interface {
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
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOpen() bool
	SetOpen(bool)
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
type HTMLDetailsElement struct {
	Value
}

func JSValueToHTMLDetailsElement(val js.Value) HTMLDetailsElement {
	return HTMLDetailsElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLDetailsElement() HTMLDetailsElement { return HTMLDetailsElement{Value: v} }
func NewHTMLDetailsElement(args ...interface{}) HTMLDetailsElement {
	return HTMLDetailsElement{Value: JSValueToValue(js.Global().Get("HTMLDetailsElement").New(args...))}
}
func (h HTMLDetailsElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLDetailsElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLDetailsElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLDetailsElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLDetailsElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLDetailsElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLDetailsElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLDetailsElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLDetailsElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLDetailsElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLDetailsElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLDetailsElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLDetailsElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLDetailsElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLDetailsElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLDetailsElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLDetailsElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLDetailsElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLDetailsElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLDetailsElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLDetailsElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLDetailsElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLDetailsElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLDetailsElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLDetailsElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLDetailsElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLDetailsElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLDetailsElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLDetailsElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLDetailsElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLDetailsElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLDetailsElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLDetailsElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLDetailsElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLDetailsElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLDetailsElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLDetailsElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLDetailsElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLDetailsElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLDetailsElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLDetailsElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLDetailsElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLDetailsElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLDetailsElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLDetailsElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLDetailsElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLDetailsElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLDetailsElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLDetailsElement) GetOpen() bool {
	val := h.Get("open")
	return val.Bool()
}
func (h HTMLDetailsElement) SetOpen(val bool) {
	h.Set("open", val)
}
func (h HTMLDetailsElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLDetailsElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLDetailsElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLDetailsElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLDetailsElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLDetailsElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLDetailsElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLDetailsElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLDetailsElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLDetailsElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLDetailsElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLDetailsElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLDetailsElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLDetailsElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLDetailsElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLDetailsElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLDetailsElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLDetailsElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLDetailsElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLDetailsElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLDetailsElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLDetailsElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLDetailsElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLDetailsElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLDetailsElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLDetailsElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
