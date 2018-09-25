// Code generated DO NOT EDIT
// htmlframeelement.go
package dom

import "syscall/js"

type HTMLFrameElementIFace interface {
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
	GetContentDocument() Document
	GetContentWindow() WindowProxy
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetFirstChild() Node
	GetFrameBorder() string
	SetFrameBorder(string)
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
	GetLongDesc() string
	SetLongDesc(string)
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetMarginHeight() string
	SetMarginHeight(string)
	GetMarginWidth() string
	SetMarginWidth(string)
	Matches(args ...interface{}) bool
	GetName() string
	SetName(string)
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNoResize() bool
	SetNoResize(bool)
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
	GetScrolling() string
	SetScrolling(string)
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
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLFrameElement struct {
	Value
}

func JSValueToHTMLFrameElement(val js.Value) HTMLFrameElement {
	return HTMLFrameElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLFrameElement() HTMLFrameElement { return HTMLFrameElement{Value: v} }
func NewHTMLFrameElement(args ...interface{}) HTMLFrameElement {
	return HTMLFrameElement{Value: JSValueToValue(js.Global().Get("HTMLFrameElement").New(args...))}
}
func (h HTMLFrameElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLFrameElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLFrameElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLFrameElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLFrameElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLFrameElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLFrameElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLFrameElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLFrameElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLFrameElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLFrameElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLFrameElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLFrameElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLFrameElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLFrameElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLFrameElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLFrameElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLFrameElement) GetContentDocument() Document {
	val := h.Get("contentDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLFrameElement) GetContentWindow() WindowProxy {
	val := h.Get("contentWindow")
	return JSValueToWindowProxy(val.JSValue())
}
func (h HTMLFrameElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLFrameElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLFrameElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLFrameElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLFrameElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLFrameElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) GetFrameBorder() string {
	val := h.Get("frameBorder")
	return val.String()
}
func (h HTMLFrameElement) SetFrameBorder(val string) {
	h.Set("frameBorder", val)
}
func (h HTMLFrameElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLFrameElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLFrameElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLFrameElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFrameElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFrameElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFrameElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFrameElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFrameElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLFrameElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLFrameElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLFrameElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLFrameElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLFrameElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLFrameElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLFrameElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLFrameElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLFrameElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLFrameElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLFrameElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLFrameElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLFrameElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLFrameElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLFrameElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLFrameElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLFrameElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLFrameElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLFrameElement) GetLongDesc() string {
	val := h.Get("longDesc")
	return val.String()
}
func (h HTMLFrameElement) SetLongDesc(val string) {
	h.Set("longDesc", val)
}
func (h HTMLFrameElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLFrameElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLFrameElement) GetMarginHeight() string {
	val := h.Get("marginHeight")
	return val.String()
}
func (h HTMLFrameElement) SetMarginHeight(val string) {
	h.Set("marginHeight", val)
}
func (h HTMLFrameElement) GetMarginWidth() string {
	val := h.Get("marginWidth")
	return val.String()
}
func (h HTMLFrameElement) SetMarginWidth(val string) {
	h.Set("marginWidth", val)
}
func (h HTMLFrameElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLFrameElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLFrameElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLFrameElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLFrameElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) GetNoResize() bool {
	val := h.Get("noResize")
	return val.Bool()
}
func (h HTMLFrameElement) SetNoResize(val bool) {
	h.Set("noResize", val)
}
func (h HTMLFrameElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLFrameElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLFrameElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLFrameElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLFrameElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLFrameElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLFrameElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLFrameElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLFrameElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLFrameElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLFrameElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFrameElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLFrameElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameElement) GetScrolling() string {
	val := h.Get("scrolling")
	return val.String()
}
func (h HTMLFrameElement) SetScrolling(val string) {
	h.Set("scrolling", val)
}
func (h HTMLFrameElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLFrameElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLFrameElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFrameElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFrameElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLFrameElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLFrameElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLFrameElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLFrameElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLFrameElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLFrameElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLFrameElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLFrameElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLFrameElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLFrameElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLFrameElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLFrameElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLFrameElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLFrameElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLFrameElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
