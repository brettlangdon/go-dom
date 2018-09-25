// Code generated DO NOT EDIT
// htmltablesectionelement.go
package dom

import "syscall/js"

type HTMLTableSectionElementIFace interface {
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
	GetCh() string
	SetCh(string)
	GetChOff() string
	SetChOff(string)
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	DeleteRow(args ...interface{})
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
	InsertRow(args ...interface{}) HTMLTableRowElement
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
	GetRows() HTMLCollection
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
	GetVAlign() string
	SetVAlign(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLTableSectionElement struct {
	Value
}

func JSValueToHTMLTableSectionElement(val js.Value) HTMLTableSectionElement {
	return HTMLTableSectionElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLTableSectionElement() HTMLTableSectionElement {
	return HTMLTableSectionElement{Value: v}
}
func NewHTMLTableSectionElement(args ...interface{}) HTMLTableSectionElement {
	return HTMLTableSectionElement{Value: JSValueToValue(js.Global().Get("HTMLTableSectionElement").New(args...))}
}
func (h HTMLTableSectionElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLTableSectionElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLTableSectionElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLTableSectionElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLTableSectionElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLTableSectionElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLTableSectionElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTableSectionElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLTableSectionElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLTableSectionElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLTableSectionElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLTableSectionElement) GetCh() string {
	val := h.Get("ch")
	return val.String()
}
func (h HTMLTableSectionElement) SetCh(val string) {
	h.Set("ch", val)
}
func (h HTMLTableSectionElement) GetChOff() string {
	val := h.Get("chOff")
	return val.String()
}
func (h HTMLTableSectionElement) SetChOff(val string) {
	h.Set("chOff", val)
}
func (h HTMLTableSectionElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLTableSectionElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLTableSectionElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLTableSectionElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLTableSectionElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLTableSectionElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableSectionElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLTableSectionElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) DeleteRow(args ...interface{}) {
	h.Call("deleteRow", args...)
}
func (h HTMLTableSectionElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLTableSectionElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLTableSectionElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLTableSectionElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLTableSectionElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLTableSectionElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLTableSectionElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLTableSectionElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableSectionElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableSectionElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableSectionElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableSectionElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableSectionElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLTableSectionElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLTableSectionElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLTableSectionElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLTableSectionElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLTableSectionElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLTableSectionElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableSectionElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLTableSectionElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) InsertRow(args ...interface{}) HTMLTableRowElement {
	val := h.Call("insertRow", args...)
	return JSValueToHTMLTableRowElement(val.JSValue())
}
func (h HTMLTableSectionElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLTableSectionElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLTableSectionElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLTableSectionElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLTableSectionElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLTableSectionElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLTableSectionElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLTableSectionElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLTableSectionElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLTableSectionElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLTableSectionElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLTableSectionElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLTableSectionElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLTableSectionElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableSectionElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLTableSectionElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLTableSectionElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLTableSectionElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableSectionElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLTableSectionElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableSectionElement) GetRows() HTMLCollection {
	val := h.Get("rows")
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableSectionElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLTableSectionElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLTableSectionElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableSectionElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableSectionElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTableSectionElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLTableSectionElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLTableSectionElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLTableSectionElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLTableSectionElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLTableSectionElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLTableSectionElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLTableSectionElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLTableSectionElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLTableSectionElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLTableSectionElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLTableSectionElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLTableSectionElement) GetVAlign() string {
	val := h.Get("vAlign")
	return val.String()
}
func (h HTMLTableSectionElement) SetVAlign(val string) {
	h.Set("vAlign", val)
}
func (h HTMLTableSectionElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
