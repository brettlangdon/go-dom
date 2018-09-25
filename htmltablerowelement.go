// Code generated DO NOT EDIT
// htmltablerowelement.go
package dom

import "syscall/js"

type HTMLTableRowElementIFace interface {
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
	GetBgColor() string
	SetBgColor(string)
	GetCells() HTMLCollection
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
	DeleteCell(args ...interface{})
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
	InsertCell(args ...interface{}) HTMLTableCellElement
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
	GetRowIndex() int
	GetSectionRowIndex() int
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
type HTMLTableRowElement struct {
	Value
}

func JSValueToHTMLTableRowElement(val js.Value) HTMLTableRowElement {
	return HTMLTableRowElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLTableRowElement() HTMLTableRowElement { return HTMLTableRowElement{Value: v} }
func NewHTMLTableRowElement(args ...interface{}) HTMLTableRowElement {
	return HTMLTableRowElement{Value: JSValueToValue(js.Global().Get("HTMLTableRowElement").New(args...))}
}
func (h HTMLTableRowElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLTableRowElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLTableRowElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLTableRowElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLTableRowElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLTableRowElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLTableRowElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTableRowElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLTableRowElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLTableRowElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLTableRowElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLTableRowElement) GetBgColor() string {
	val := h.Get("bgColor")
	return val.String()
}
func (h HTMLTableRowElement) SetBgColor(val string) {
	h.Set("bgColor", val)
}
func (h HTMLTableRowElement) GetCells() HTMLCollection {
	val := h.Get("cells")
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableRowElement) GetCh() string {
	val := h.Get("ch")
	return val.String()
}
func (h HTMLTableRowElement) SetCh(val string) {
	h.Set("ch", val)
}
func (h HTMLTableRowElement) GetChOff() string {
	val := h.Get("chOff")
	return val.String()
}
func (h HTMLTableRowElement) SetChOff(val string) {
	h.Set("chOff", val)
}
func (h HTMLTableRowElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLTableRowElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLTableRowElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLTableRowElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLTableRowElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLTableRowElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableRowElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLTableRowElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) DeleteCell(args ...interface{}) {
	h.Call("deleteCell", args...)
}
func (h HTMLTableRowElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLTableRowElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLTableRowElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLTableRowElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLTableRowElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLTableRowElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLTableRowElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLTableRowElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableRowElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableRowElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableRowElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableRowElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableRowElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLTableRowElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLTableRowElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLTableRowElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLTableRowElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLTableRowElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLTableRowElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableRowElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLTableRowElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) InsertCell(args ...interface{}) HTMLTableCellElement {
	val := h.Call("insertCell", args...)
	return JSValueToHTMLTableCellElement(val.JSValue())
}
func (h HTMLTableRowElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLTableRowElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLTableRowElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLTableRowElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLTableRowElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLTableRowElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLTableRowElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLTableRowElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLTableRowElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLTableRowElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLTableRowElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLTableRowElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLTableRowElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLTableRowElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableRowElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLTableRowElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLTableRowElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLTableRowElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableRowElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLTableRowElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableRowElement) GetRowIndex() int {
	val := h.Get("rowIndex")
	return val.Int()
}
func (h HTMLTableRowElement) GetSectionRowIndex() int {
	val := h.Get("sectionRowIndex")
	return val.Int()
}
func (h HTMLTableRowElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLTableRowElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLTableRowElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableRowElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableRowElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTableRowElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLTableRowElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLTableRowElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLTableRowElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLTableRowElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLTableRowElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLTableRowElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLTableRowElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLTableRowElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLTableRowElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLTableRowElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLTableRowElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLTableRowElement) GetVAlign() string {
	val := h.Get("vAlign")
	return val.String()
}
func (h HTMLTableRowElement) SetVAlign(val string) {
	h.Set("vAlign", val)
}
func (h HTMLTableRowElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
