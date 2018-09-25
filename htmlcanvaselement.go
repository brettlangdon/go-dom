// Code generated DO NOT EDIT
// htmlcanvaselement.go
package dom

import "syscall/js"

type HTMLCanvasElementIFace interface {
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
	GetContext(args ...interface{}) RenderingContext
	GetElementsByClassName(args ...interface{}) HTMLCollection
	GetElementsByTagName(args ...interface{}) HTMLCollection
	GetElementsByTagNameNS(args ...interface{}) HTMLCollection
	GetRootNode(args ...interface{}) Node
	HasAttribute(args ...interface{}) bool
	HasAttributeNS(args ...interface{}) bool
	HasAttributes(args ...interface{}) bool
	HasChildNodes(args ...interface{}) bool
	GetHeight() int
	SetHeight(int)
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
	ToBlob(args ...interface{})
	ToDataURL(args ...interface{}) string
	ToggleAttribute(args ...interface{}) bool
	TransferControlToOffscreen(args ...interface{}) OffscreenCanvas
	GetTranslate() bool
	SetTranslate(bool)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() int
	SetWidth(int)
}
type HTMLCanvasElement struct {
	Value
}

func JSValueToHTMLCanvasElement(val js.Value) HTMLCanvasElement {
	return HTMLCanvasElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLCanvasElement() HTMLCanvasElement { return HTMLCanvasElement{Value: v} }
func NewHTMLCanvasElement(args ...interface{}) HTMLCanvasElement {
	return HTMLCanvasElement{Value: JSValueToValue(js.Global().Get("HTMLCanvasElement").New(args...))}
}
func (h HTMLCanvasElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLCanvasElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLCanvasElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLCanvasElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLCanvasElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLCanvasElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLCanvasElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLCanvasElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLCanvasElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLCanvasElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLCanvasElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLCanvasElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLCanvasElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLCanvasElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLCanvasElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLCanvasElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLCanvasElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLCanvasElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLCanvasElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLCanvasElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLCanvasElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLCanvasElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLCanvasElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLCanvasElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLCanvasElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLCanvasElement) GetContext(args ...interface{}) RenderingContext {
	val := h.Call("getContext", args...)
	return JSValueToRenderingContext(val.JSValue())
}
func (h HTMLCanvasElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLCanvasElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLCanvasElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLCanvasElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) GetHeight() int {
	val := h.Get("height")
	return val.Int()
}
func (h HTMLCanvasElement) SetHeight(val int) {
	h.Set("height", val)
}
func (h HTMLCanvasElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLCanvasElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLCanvasElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLCanvasElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLCanvasElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLCanvasElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLCanvasElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLCanvasElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLCanvasElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLCanvasElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLCanvasElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLCanvasElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLCanvasElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLCanvasElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLCanvasElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLCanvasElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLCanvasElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLCanvasElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLCanvasElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLCanvasElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLCanvasElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLCanvasElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLCanvasElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLCanvasElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLCanvasElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLCanvasElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLCanvasElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLCanvasElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLCanvasElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLCanvasElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLCanvasElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLCanvasElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLCanvasElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLCanvasElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLCanvasElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLCanvasElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLCanvasElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLCanvasElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLCanvasElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLCanvasElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLCanvasElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLCanvasElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLCanvasElement) ToBlob(args ...interface{}) {
	h.Call("toBlob", args...)
}
func (h HTMLCanvasElement) ToDataURL(args ...interface{}) string {
	val := h.Call("toDataURL", args...)
	return val.String()
}
func (h HTMLCanvasElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) TransferControlToOffscreen(args ...interface{}) OffscreenCanvas {
	val := h.Call("transferControlToOffscreen", args...)
	return JSValueToOffscreenCanvas(val.JSValue())
}
func (h HTMLCanvasElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLCanvasElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLCanvasElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLCanvasElement) GetWidth() int {
	val := h.Get("width")
	return val.Int()
}
func (h HTMLCanvasElement) SetWidth(val int) {
	h.Set("width", val)
}
