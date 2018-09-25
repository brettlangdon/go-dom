// Code generated DO NOT EDIT
// htmltablecellelement.go
package dom

import "syscall/js"

type HTMLTableCellElementIFace interface {
	GetAbbr() string
	SetAbbr(string)
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
	GetAxis() string
	SetAxis(string)
	GetBaseURI() string
	GetBgColor() string
	SetBgColor(string)
	GetCellIndex() int
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
	GetColSpan() int
	SetColSpan(int)
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
	GetHeaders() string
	SetHeaders(string)
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
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNoWrap() bool
	SetNoWrap(bool)
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
	GetRowSpan() int
	SetRowSpan(int)
	GetScope() string
	SetScope(string)
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
	GetWidth() string
	SetWidth(string)
}
type HTMLTableCellElement struct {
	Value
}

func JSValueToHTMLTableCellElement(val js.Value) HTMLTableCellElement {
	return HTMLTableCellElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLTableCellElement() HTMLTableCellElement { return HTMLTableCellElement{Value: v} }
func NewHTMLTableCellElement(args ...interface{}) HTMLTableCellElement {
	return HTMLTableCellElement{Value: JSValueToValue(js.Global().Get("HTMLTableCellElement").New(args...))}
}
func (h HTMLTableCellElement) GetAbbr() string {
	val := h.Get("abbr")
	return val.String()
}
func (h HTMLTableCellElement) SetAbbr(val string) {
	h.Set("abbr", val)
}
func (h HTMLTableCellElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLTableCellElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLTableCellElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLTableCellElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLTableCellElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLTableCellElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLTableCellElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTableCellElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLTableCellElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLTableCellElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLTableCellElement) GetAxis() string {
	val := h.Get("axis")
	return val.String()
}
func (h HTMLTableCellElement) SetAxis(val string) {
	h.Set("axis", val)
}
func (h HTMLTableCellElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLTableCellElement) GetBgColor() string {
	val := h.Get("bgColor")
	return val.String()
}
func (h HTMLTableCellElement) SetBgColor(val string) {
	h.Set("bgColor", val)
}
func (h HTMLTableCellElement) GetCellIndex() int {
	val := h.Get("cellIndex")
	return val.Int()
}
func (h HTMLTableCellElement) GetCh() string {
	val := h.Get("ch")
	return val.String()
}
func (h HTMLTableCellElement) SetCh(val string) {
	h.Set("ch", val)
}
func (h HTMLTableCellElement) GetChOff() string {
	val := h.Get("chOff")
	return val.String()
}
func (h HTMLTableCellElement) SetChOff(val string) {
	h.Set("chOff", val)
}
func (h HTMLTableCellElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLTableCellElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLTableCellElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLTableCellElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLTableCellElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLTableCellElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableCellElement) GetColSpan() int {
	val := h.Get("colSpan")
	return val.Int()
}
func (h HTMLTableCellElement) SetColSpan(val int) {
	h.Set("colSpan", val)
}
func (h HTMLTableCellElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLTableCellElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLTableCellElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLTableCellElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLTableCellElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLTableCellElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLTableCellElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLTableCellElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLTableCellElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableCellElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableCellElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableCellElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableCellElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableCellElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) GetHeaders() string {
	val := h.Get("headers")
	return val.String()
}
func (h HTMLTableCellElement) SetHeaders(val string) {
	h.Set("headers", val)
}
func (h HTMLTableCellElement) GetHeight() string {
	val := h.Get("height")
	return val.String()
}
func (h HTMLTableCellElement) SetHeight(val string) {
	h.Set("height", val)
}
func (h HTMLTableCellElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLTableCellElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLTableCellElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLTableCellElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLTableCellElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLTableCellElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLTableCellElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableCellElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLTableCellElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLTableCellElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLTableCellElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLTableCellElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLTableCellElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLTableCellElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLTableCellElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLTableCellElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) GetNoWrap() bool {
	val := h.Get("noWrap")
	return val.Bool()
}
func (h HTMLTableCellElement) SetNoWrap(val bool) {
	h.Set("noWrap", val)
}
func (h HTMLTableCellElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLTableCellElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLTableCellElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLTableCellElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLTableCellElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLTableCellElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLTableCellElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableCellElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLTableCellElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLTableCellElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLTableCellElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableCellElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLTableCellElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableCellElement) GetRowSpan() int {
	val := h.Get("rowSpan")
	return val.Int()
}
func (h HTMLTableCellElement) SetRowSpan(val int) {
	h.Set("rowSpan", val)
}
func (h HTMLTableCellElement) GetScope() string {
	val := h.Get("scope")
	return val.String()
}
func (h HTMLTableCellElement) SetScope(val string) {
	h.Set("scope", val)
}
func (h HTMLTableCellElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLTableCellElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLTableCellElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableCellElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableCellElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTableCellElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLTableCellElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLTableCellElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLTableCellElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLTableCellElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLTableCellElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLTableCellElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLTableCellElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLTableCellElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLTableCellElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLTableCellElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLTableCellElement) GetVAlign() string {
	val := h.Get("vAlign")
	return val.String()
}
func (h HTMLTableCellElement) SetVAlign(val string) {
	h.Set("vAlign", val)
}
func (h HTMLTableCellElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLTableCellElement) GetWidth() string {
	val := h.Get("width")
	return val.String()
}
func (h HTMLTableCellElement) SetWidth(val string) {
	h.Set("width", val)
}
