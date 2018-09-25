// Code generated DO NOT EDIT
// htmltableelement.go
package dom

import "syscall/js"

type HTMLTableElementIFace interface {
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
	GetBorder() string
	SetBorder(string)
	GetCaption() HTMLTableCaptionElement
	SetCaption(HTMLTableCaptionElement)
	GetCellPadding() string
	SetCellPadding(string)
	GetCellSpacing() string
	SetCellSpacing(string)
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	CreateCaption(args ...interface{}) HTMLTableCaptionElement
	CreateTBody(args ...interface{}) HTMLTableSectionElement
	CreateTFoot(args ...interface{}) HTMLTableSectionElement
	CreateTHead(args ...interface{}) HTMLTableSectionElement
	DeleteCaption(args ...interface{})
	DeleteRow(args ...interface{})
	DeleteTFoot(args ...interface{})
	DeleteTHead(args ...interface{})
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetFirstChild() Node
	GetFrame() string
	SetFrame(string)
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
	GetRules() string
	SetRules(string)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetSummary() string
	SetSummary(string)
	GetTBodies() HTMLCollection
	GetTFoot() HTMLTableSectionElement
	SetTFoot(HTMLTableSectionElement)
	GetTHead() HTMLTableSectionElement
	SetTHead(HTMLTableSectionElement)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() string
	SetWidth(string)
}
type HTMLTableElement struct {
	Value
}

func JSValueToHTMLTableElement(val js.Value) HTMLTableElement {
	return HTMLTableElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLTableElement() HTMLTableElement { return HTMLTableElement{Value: v} }
func NewHTMLTableElement(args ...interface{}) HTMLTableElement {
	return HTMLTableElement{Value: JSValueToValue(js.Global().Get("HTMLTableElement").New(args...))}
}
func (h HTMLTableElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLTableElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLTableElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLTableElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLTableElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLTableElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLTableElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTableElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLTableElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLTableElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLTableElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLTableElement) GetBgColor() string {
	val := h.Get("bgColor")
	return val.String()
}
func (h HTMLTableElement) SetBgColor(val string) {
	h.Set("bgColor", val)
}
func (h HTMLTableElement) GetBorder() string {
	val := h.Get("border")
	return val.String()
}
func (h HTMLTableElement) SetBorder(val string) {
	h.Set("border", val)
}
func (h HTMLTableElement) GetCaption() HTMLTableCaptionElement {
	val := h.Get("caption")
	return JSValueToHTMLTableCaptionElement(val.JSValue())
}
func (h HTMLTableElement) SetCaption(val HTMLTableCaptionElement) {
	h.Set("caption", val)
}
func (h HTMLTableElement) GetCellPadding() string {
	val := h.Get("cellPadding")
	return val.String()
}
func (h HTMLTableElement) SetCellPadding(val string) {
	h.Set("cellPadding", val)
}
func (h HTMLTableElement) GetCellSpacing() string {
	val := h.Get("cellSpacing")
	return val.String()
}
func (h HTMLTableElement) SetCellSpacing(val string) {
	h.Set("cellSpacing", val)
}
func (h HTMLTableElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLTableElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLTableElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLTableElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLTableElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLTableElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLTableElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLTableElement) CreateCaption(args ...interface{}) HTMLTableCaptionElement {
	val := h.Call("createCaption", args...)
	return JSValueToHTMLTableCaptionElement(val.JSValue())
}
func (h HTMLTableElement) CreateTBody(args ...interface{}) HTMLTableSectionElement {
	val := h.Call("createTBody", args...)
	return JSValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) CreateTFoot(args ...interface{}) HTMLTableSectionElement {
	val := h.Call("createTFoot", args...)
	return JSValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) CreateTHead(args ...interface{}) HTMLTableSectionElement {
	val := h.Call("createTHead", args...)
	return JSValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) DeleteCaption(args ...interface{}) {
	h.Call("deleteCaption", args...)
}
func (h HTMLTableElement) DeleteRow(args ...interface{}) {
	h.Call("deleteRow", args...)
}
func (h HTMLTableElement) DeleteTFoot(args ...interface{}) {
	h.Call("deleteTFoot", args...)
}
func (h HTMLTableElement) DeleteTHead(args ...interface{}) {
	h.Call("deleteTHead", args...)
}
func (h HTMLTableElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLTableElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLTableElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLTableElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLTableElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLTableElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) GetFrame() string {
	val := h.Get("frame")
	return val.String()
}
func (h HTMLTableElement) SetFrame(val string) {
	h.Set("frame", val)
}
func (h HTMLTableElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLTableElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLTableElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLTableElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLTableElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLTableElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLTableElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLTableElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLTableElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLTableElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLTableElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLTableElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLTableElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLTableElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLTableElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) InsertRow(args ...interface{}) HTMLTableRowElement {
	val := h.Call("insertRow", args...)
	return JSValueToHTMLTableRowElement(val.JSValue())
}
func (h HTMLTableElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLTableElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLTableElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLTableElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLTableElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLTableElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLTableElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLTableElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLTableElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLTableElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLTableElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLTableElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLTableElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLTableElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLTableElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLTableElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLTableElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLTableElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLTableElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLTableElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLTableElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLTableElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLTableElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTableElement) GetRows() HTMLCollection {
	val := h.Get("rows")
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableElement) GetRules() string {
	val := h.Get("rules")
	return val.String()
}
func (h HTMLTableElement) SetRules(val string) {
	h.Set("rules", val)
}
func (h HTMLTableElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLTableElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLTableElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTableElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTableElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLTableElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLTableElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLTableElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLTableElement) GetSummary() string {
	val := h.Get("summary")
	return val.String()
}
func (h HTMLTableElement) SetSummary(val string) {
	h.Set("summary", val)
}
func (h HTMLTableElement) GetTBodies() HTMLCollection {
	val := h.Get("tBodies")
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableElement) GetTFoot() HTMLTableSectionElement {
	val := h.Get("tFoot")
	return JSValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) SetTFoot(val HTMLTableSectionElement) {
	h.Set("tFoot", val)
}
func (h HTMLTableElement) GetTHead() HTMLTableSectionElement {
	val := h.Get("tHead")
	return JSValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) SetTHead(val HTMLTableSectionElement) {
	h.Set("tHead", val)
}
func (h HTMLTableElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLTableElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLTableElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLTableElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLTableElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLTableElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLTableElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLTableElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLTableElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLTableElement) GetWidth() string {
	val := h.Get("width")
	return val.String()
}
func (h HTMLTableElement) SetWidth(val string) {
	h.Set("width", val)
}
