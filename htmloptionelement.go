// Code generated DO NOT EDIT
// htmloptionelement.go
package dom

import "syscall/js"

type HTMLOptionElementIFace interface {
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
	GetDefaultSelected() bool
	SetDefaultSelected(bool)
	GetDir() string
	SetDir(string)
	GetDisabled() bool
	SetDisabled(bool)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
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
	GetIndex() int
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
	GetSelected() bool
	SetSelected(bool)
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
	GetText() string
	SetText(string)
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetValue() string
	SetValue(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLOptionElement struct {
	Value
}

func JSValueToHTMLOptionElement(val js.Value) HTMLOptionElement {
	return HTMLOptionElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLOptionElement() HTMLOptionElement { return HTMLOptionElement{Value: v} }
func NewHTMLOptionElement(args ...interface{}) HTMLOptionElement {
	return HTMLOptionElement{Value: JSValueToValue(js.Global().Get("HTMLOptionElement").New(args...))}
}
func (h HTMLOptionElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLOptionElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLOptionElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLOptionElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLOptionElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLOptionElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLOptionElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLOptionElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLOptionElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLOptionElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLOptionElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLOptionElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLOptionElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLOptionElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLOptionElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLOptionElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLOptionElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLOptionElement) GetDefaultSelected() bool {
	val := h.Get("defaultSelected")
	return val.Bool()
}
func (h HTMLOptionElement) SetDefaultSelected(val bool) {
	h.Set("defaultSelected", val)
}
func (h HTMLOptionElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLOptionElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLOptionElement) GetDisabled() bool {
	val := h.Get("disabled")
	return val.Bool()
}
func (h HTMLOptionElement) SetDisabled(val bool) {
	h.Set("disabled", val)
}
func (h HTMLOptionElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLOptionElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLOptionElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLOptionElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLOptionElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLOptionElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLOptionElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLOptionElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOptionElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOptionElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLOptionElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLOptionElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLOptionElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLOptionElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLOptionElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLOptionElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLOptionElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLOptionElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLOptionElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLOptionElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLOptionElement) GetIndex() int {
	val := h.Get("index")
	return val.Int()
}
func (h HTMLOptionElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLOptionElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLOptionElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLOptionElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLOptionElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLOptionElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLOptionElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLOptionElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLOptionElement) GetLabel() string {
	val := h.Get("label")
	return val.String()
}
func (h HTMLOptionElement) SetLabel(val string) {
	h.Set("label", val)
}
func (h HTMLOptionElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLOptionElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLOptionElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLOptionElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLOptionElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLOptionElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLOptionElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLOptionElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLOptionElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLOptionElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLOptionElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLOptionElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLOptionElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLOptionElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLOptionElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLOptionElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLOptionElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLOptionElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOptionElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLOptionElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLOptionElement) GetSelected() bool {
	val := h.Get("selected")
	return val.Bool()
}
func (h HTMLOptionElement) SetSelected(val bool) {
	h.Set("selected", val)
}
func (h HTMLOptionElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLOptionElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLOptionElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOptionElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLOptionElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLOptionElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLOptionElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLOptionElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLOptionElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLOptionElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLOptionElement) GetText() string {
	val := h.Get("text")
	return val.String()
}
func (h HTMLOptionElement) SetText(val string) {
	h.Set("text", val)
}
func (h HTMLOptionElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLOptionElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLOptionElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLOptionElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLOptionElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLOptionElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLOptionElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLOptionElement) GetValue() string {
	val := h.Get("value")
	return val.String()
}
func (h HTMLOptionElement) SetValue(val string) {
	h.Set("value", val)
}
func (h HTMLOptionElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
