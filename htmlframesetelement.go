// Code generated DO NOT EDIT
// htmlframesetelement.go
package dom

import "syscall/js"

type HTMLFrameSetElementIFace interface {
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
	GetCols() string
	SetCols(string)
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
	GetOnafterprint() EventHandler
	SetOnafterprint(EventHandler)
	GetOnbeforeprint() EventHandler
	SetOnbeforeprint(EventHandler)
	GetOnbeforeunload() OnBeforeUnloadEventHandler
	SetOnbeforeunload(OnBeforeUnloadEventHandler)
	GetOnhashchange() EventHandler
	SetOnhashchange(EventHandler)
	GetOnlanguagechange() EventHandler
	SetOnlanguagechange(EventHandler)
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnmessageerror() EventHandler
	SetOnmessageerror(EventHandler)
	GetOnoffline() EventHandler
	SetOnoffline(EventHandler)
	GetOnonline() EventHandler
	SetOnonline(EventHandler)
	GetOnpagehide() EventHandler
	SetOnpagehide(EventHandler)
	GetOnpageshow() EventHandler
	SetOnpageshow(EventHandler)
	GetOnpopstate() EventHandler
	SetOnpopstate(EventHandler)
	GetOnrejectionhandled() EventHandler
	SetOnrejectionhandled(EventHandler)
	GetOnstorage() EventHandler
	SetOnstorage(EventHandler)
	GetOnunhandledrejection() EventHandler
	SetOnunhandledrejection(EventHandler)
	GetOnunload() EventHandler
	SetOnunload(EventHandler)
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
	GetRows() string
	SetRows(string)
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
type HTMLFrameSetElement struct {
	Value
}

func JSValueToHTMLFrameSetElement(val js.Value) HTMLFrameSetElement {
	return HTMLFrameSetElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLFrameSetElement() HTMLFrameSetElement { return HTMLFrameSetElement{Value: v} }
func NewHTMLFrameSetElement(args ...interface{}) HTMLFrameSetElement {
	return HTMLFrameSetElement{Value: JSValueToValue(js.Global().Get("HTMLFrameSetElement").New(args...))}
}
func (h HTMLFrameSetElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLFrameSetElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLFrameSetElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLFrameSetElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLFrameSetElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLFrameSetElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLFrameSetElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLFrameSetElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLFrameSetElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLFrameSetElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLFrameSetElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLFrameSetElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLFrameSetElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLFrameSetElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLFrameSetElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLFrameSetElement) GetCols() string {
	val := h.Get("cols")
	return val.String()
}
func (h HTMLFrameSetElement) SetCols(val string) {
	h.Set("cols", val)
}
func (h HTMLFrameSetElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLFrameSetElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLFrameSetElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLFrameSetElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLFrameSetElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLFrameSetElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLFrameSetElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLFrameSetElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLFrameSetElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFrameSetElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFrameSetElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFrameSetElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFrameSetElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLFrameSetElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLFrameSetElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLFrameSetElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLFrameSetElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLFrameSetElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLFrameSetElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLFrameSetElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLFrameSetElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLFrameSetElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLFrameSetElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLFrameSetElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLFrameSetElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLFrameSetElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLFrameSetElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLFrameSetElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLFrameSetElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLFrameSetElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLFrameSetElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLFrameSetElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLFrameSetElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLFrameSetElement) GetOnafterprint() EventHandler {
	val := h.Get("onafterprint")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnafterprint(val EventHandler) {
	h.Set("onafterprint", val)
}
func (h HTMLFrameSetElement) GetOnbeforeprint() EventHandler {
	val := h.Get("onbeforeprint")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnbeforeprint(val EventHandler) {
	h.Set("onbeforeprint", val)
}
func (h HTMLFrameSetElement) GetOnbeforeunload() OnBeforeUnloadEventHandler {
	val := h.Get("onbeforeunload")
	return JSValueToOnBeforeUnloadEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnbeforeunload(val OnBeforeUnloadEventHandler) {
	h.Set("onbeforeunload", val)
}
func (h HTMLFrameSetElement) GetOnhashchange() EventHandler {
	val := h.Get("onhashchange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnhashchange(val EventHandler) {
	h.Set("onhashchange", val)
}
func (h HTMLFrameSetElement) GetOnlanguagechange() EventHandler {
	val := h.Get("onlanguagechange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnlanguagechange(val EventHandler) {
	h.Set("onlanguagechange", val)
}
func (h HTMLFrameSetElement) GetOnmessage() EventHandler {
	val := h.Get("onmessage")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnmessage(val EventHandler) {
	h.Set("onmessage", val)
}
func (h HTMLFrameSetElement) GetOnmessageerror() EventHandler {
	val := h.Get("onmessageerror")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnmessageerror(val EventHandler) {
	h.Set("onmessageerror", val)
}
func (h HTMLFrameSetElement) GetOnoffline() EventHandler {
	val := h.Get("onoffline")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnoffline(val EventHandler) {
	h.Set("onoffline", val)
}
func (h HTMLFrameSetElement) GetOnonline() EventHandler {
	val := h.Get("ononline")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnonline(val EventHandler) {
	h.Set("ononline", val)
}
func (h HTMLFrameSetElement) GetOnpagehide() EventHandler {
	val := h.Get("onpagehide")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnpagehide(val EventHandler) {
	h.Set("onpagehide", val)
}
func (h HTMLFrameSetElement) GetOnpageshow() EventHandler {
	val := h.Get("onpageshow")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnpageshow(val EventHandler) {
	h.Set("onpageshow", val)
}
func (h HTMLFrameSetElement) GetOnpopstate() EventHandler {
	val := h.Get("onpopstate")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnpopstate(val EventHandler) {
	h.Set("onpopstate", val)
}
func (h HTMLFrameSetElement) GetOnrejectionhandled() EventHandler {
	val := h.Get("onrejectionhandled")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnrejectionhandled(val EventHandler) {
	h.Set("onrejectionhandled", val)
}
func (h HTMLFrameSetElement) GetOnstorage() EventHandler {
	val := h.Get("onstorage")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnstorage(val EventHandler) {
	h.Set("onstorage", val)
}
func (h HTMLFrameSetElement) GetOnunhandledrejection() EventHandler {
	val := h.Get("onunhandledrejection")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnunhandledrejection(val EventHandler) {
	h.Set("onunhandledrejection", val)
}
func (h HTMLFrameSetElement) GetOnunload() EventHandler {
	val := h.Get("onunload")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnunload(val EventHandler) {
	h.Set("onunload", val)
}
func (h HTMLFrameSetElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLFrameSetElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLFrameSetElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLFrameSetElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLFrameSetElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLFrameSetElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFrameSetElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLFrameSetElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLFrameSetElement) GetRows() string {
	val := h.Get("rows")
	return val.String()
}
func (h HTMLFrameSetElement) SetRows(val string) {
	h.Set("rows", val)
}
func (h HTMLFrameSetElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLFrameSetElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLFrameSetElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFrameSetElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLFrameSetElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLFrameSetElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLFrameSetElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLFrameSetElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLFrameSetElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLFrameSetElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLFrameSetElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLFrameSetElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLFrameSetElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLFrameSetElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLFrameSetElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLFrameSetElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLFrameSetElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLFrameSetElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
