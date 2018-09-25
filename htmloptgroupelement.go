// Code generated DO NOT EDIT
// htmloptgroupelement.go
package dom

import "syscall/js"

type HTMLOptGroupElementIFace interface {
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
	GetDisabled() bool
	SetDisabled(bool)
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
	GetLabel() string
	SetLabel(string)
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
type HTMLOptGroupElement struct {
	Value
}

func JSValueToHTMLOptGroupElement(val js.Value) HTMLOptGroupElement {
	return HTMLOptGroupElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLOptGroupElement() HTMLOptGroupElement { return HTMLOptGroupElement{Value: v} }
func NewHTMLOptGroupElement(args ...interface{}) HTMLOptGroupElement {
	return HTMLOptGroupElement{Value: JSValueToValue(js.Global().Get("HTMLOptGroupElement").New(args...))}
}
func (h HTMLOptGroupElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLOptGroupElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLOptGroupElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLOptGroupElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLOptGroupElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLOptGroupElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLOptGroupElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLOptGroupElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLOptGroupElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLOptGroupElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLOptGroupElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLOptGroupElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLOptGroupElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLOptGroupElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLOptGroupElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLOptGroupElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLOptGroupElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLOptGroupElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLOptGroupElement) GetDisabled() bool {
	val := h.Get("disabled")
	return val.Bool()
}
func (h HTMLOptGroupElement) SetDisabled(val bool) {
	h.Set("disabled", val)
}
func (h HTMLOptGroupElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLOptGroupElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLOptGroupElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLOptGroupElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLOptGroupElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLOptGroupElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOptGroupElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOptGroupElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLOptGroupElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLOptGroupElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLOptGroupElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLOptGroupElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLOptGroupElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLOptGroupElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLOptGroupElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLOptGroupElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLOptGroupElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLOptGroupElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLOptGroupElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLOptGroupElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) GetLabel() string {
	val := h.Get("label")
	return val.String()
}
func (h HTMLOptGroupElement) SetLabel(val string) {
	h.Set("label", val)
}
func (h HTMLOptGroupElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLOptGroupElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLOptGroupElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLOptGroupElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLOptGroupElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLOptGroupElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLOptGroupElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLOptGroupElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLOptGroupElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLOptGroupElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLOptGroupElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLOptGroupElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLOptGroupElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLOptGroupElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLOptGroupElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLOptGroupElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLOptGroupElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOptGroupElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLOptGroupElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptGroupElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLOptGroupElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLOptGroupElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOptGroupElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOptGroupElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLOptGroupElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLOptGroupElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLOptGroupElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLOptGroupElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLOptGroupElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLOptGroupElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLOptGroupElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLOptGroupElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLOptGroupElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLOptGroupElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLOptGroupElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLOptGroupElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLOptGroupElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
