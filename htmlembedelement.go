// Code generated DO NOT EDIT
// htmlembedelement.go
package dom

import "syscall/js"

type HTMLEmbedElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	GetAlign() string
	SetAlign(string)
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
	GetSVGDocument(args ...interface{}) Document
	HasAttribute(args ...interface{}) bool
	HasAttributeNS(args ...interface{}) bool
	HasAttributes(args ...interface{}) bool
	HasChildNodes(args ...interface{}) bool
	GetHeight() string
	SetHeight(string)
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
	GetType() string
	SetType(string)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() string
	SetWidth(string)
}
type HTMLEmbedElement struct {
	Value
}

func JSValueToHTMLEmbedElement(val js.Value) HTMLEmbedElement {
	return HTMLEmbedElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLEmbedElement() HTMLEmbedElement { return HTMLEmbedElement{Value: v} }
func NewHTMLEmbedElement(args ...interface{}) HTMLEmbedElement {
	return HTMLEmbedElement{Value: JSValueToValue(js.Global().Get("HTMLEmbedElement").New(args...))}
}
func (h HTMLEmbedElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLEmbedElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLEmbedElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLEmbedElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLEmbedElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLEmbedElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLEmbedElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLEmbedElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLEmbedElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLEmbedElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLEmbedElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLEmbedElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLEmbedElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLEmbedElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLEmbedElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLEmbedElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLEmbedElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLEmbedElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLEmbedElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLEmbedElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLEmbedElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLEmbedElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLEmbedElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLEmbedElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLEmbedElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLEmbedElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLEmbedElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLEmbedElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLEmbedElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLEmbedElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLEmbedElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) GetSVGDocument(args ...interface{}) Document {
	val := h.Call("getSVGDocument", args...)
	return JSValueToDocument(val.JSValue())
}
func (h HTMLEmbedElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) GetHeight() string {
	val := h.Get("height")
	return val.String()
}
func (h HTMLEmbedElement) SetHeight(val string) {
	h.Set("height", val)
}
func (h HTMLEmbedElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLEmbedElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLEmbedElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLEmbedElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLEmbedElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLEmbedElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLEmbedElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLEmbedElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLEmbedElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLEmbedElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLEmbedElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLEmbedElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLEmbedElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLEmbedElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLEmbedElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLEmbedElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLEmbedElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLEmbedElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLEmbedElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLEmbedElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLEmbedElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLEmbedElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLEmbedElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLEmbedElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLEmbedElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLEmbedElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLEmbedElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLEmbedElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLEmbedElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLEmbedElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLEmbedElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLEmbedElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLEmbedElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLEmbedElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLEmbedElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLEmbedElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLEmbedElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLEmbedElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLEmbedElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLEmbedElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLEmbedElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLEmbedElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLEmbedElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLEmbedElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLEmbedElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLEmbedElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLEmbedElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLEmbedElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLEmbedElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLEmbedElement) SetType(val string) {
	h.Set("type", val)
}
func (h HTMLEmbedElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLEmbedElement) GetWidth() string {
	val := h.Get("width")
	return val.String()
}
func (h HTMLEmbedElement) SetWidth(val string) {
	h.Set("width", val)
}
