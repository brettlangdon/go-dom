// Code generated DO NOT EDIT
// htmlbodyelement.go
package dom

import "syscall/js"

type HTMLBodyElementIFace interface {
	GetALink() string
	SetALink(string)
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetBackground() string
	SetBackground(string)
	GetBaseURI() string
	GetBgColor() string
	SetBgColor(string)
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
	GetLink() string
	SetLink(string)
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
	GetVLink() string
	SetVLink(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLBodyElement struct {
	Value
}

func JSValueToHTMLBodyElement(val js.Value) HTMLBodyElement {
	return HTMLBodyElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLBodyElement() HTMLBodyElement { return HTMLBodyElement{Value: v} }
func NewHTMLBodyElement(args ...interface{}) HTMLBodyElement {
	return HTMLBodyElement{Value: JSValueToValue(js.Global().Get("HTMLBodyElement").New(args...))}
}
func (h HTMLBodyElement) GetALink() string {
	val := h.Get("aLink")
	return val.String()
}
func (h HTMLBodyElement) SetALink(val string) {
	h.Set("aLink", val)
}
func (h HTMLBodyElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLBodyElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLBodyElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLBodyElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLBodyElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLBodyElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLBodyElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLBodyElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLBodyElement) GetBackground() string {
	val := h.Get("background")
	return val.String()
}
func (h HTMLBodyElement) SetBackground(val string) {
	h.Set("background", val)
}
func (h HTMLBodyElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLBodyElement) GetBgColor() string {
	val := h.Get("bgColor")
	return val.String()
}
func (h HTMLBodyElement) SetBgColor(val string) {
	h.Set("bgColor", val)
}
func (h HTMLBodyElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLBodyElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLBodyElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLBodyElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLBodyElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLBodyElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLBodyElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLBodyElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLBodyElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLBodyElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLBodyElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLBodyElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLBodyElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLBodyElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLBodyElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLBodyElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLBodyElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLBodyElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLBodyElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLBodyElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLBodyElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLBodyElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLBodyElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLBodyElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLBodyElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLBodyElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLBodyElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLBodyElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLBodyElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLBodyElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLBodyElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLBodyElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLBodyElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLBodyElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLBodyElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLBodyElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLBodyElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLBodyElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLBodyElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLBodyElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) GetLink() string {
	val := h.Get("link")
	return val.String()
}
func (h HTMLBodyElement) SetLink(val string) {
	h.Set("link", val)
}
func (h HTMLBodyElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLBodyElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLBodyElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLBodyElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLBodyElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLBodyElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLBodyElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLBodyElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLBodyElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLBodyElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLBodyElement) GetOnafterprint() EventHandler {
	val := h.Get("onafterprint")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnafterprint(val EventHandler) {
	h.Set("onafterprint", val)
}
func (h HTMLBodyElement) GetOnbeforeprint() EventHandler {
	val := h.Get("onbeforeprint")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnbeforeprint(val EventHandler) {
	h.Set("onbeforeprint", val)
}
func (h HTMLBodyElement) GetOnbeforeunload() OnBeforeUnloadEventHandler {
	val := h.Get("onbeforeunload")
	return JSValueToOnBeforeUnloadEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnbeforeunload(val OnBeforeUnloadEventHandler) {
	h.Set("onbeforeunload", val)
}
func (h HTMLBodyElement) GetOnhashchange() EventHandler {
	val := h.Get("onhashchange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnhashchange(val EventHandler) {
	h.Set("onhashchange", val)
}
func (h HTMLBodyElement) GetOnlanguagechange() EventHandler {
	val := h.Get("onlanguagechange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnlanguagechange(val EventHandler) {
	h.Set("onlanguagechange", val)
}
func (h HTMLBodyElement) GetOnmessage() EventHandler {
	val := h.Get("onmessage")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnmessage(val EventHandler) {
	h.Set("onmessage", val)
}
func (h HTMLBodyElement) GetOnmessageerror() EventHandler {
	val := h.Get("onmessageerror")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnmessageerror(val EventHandler) {
	h.Set("onmessageerror", val)
}
func (h HTMLBodyElement) GetOnoffline() EventHandler {
	val := h.Get("onoffline")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnoffline(val EventHandler) {
	h.Set("onoffline", val)
}
func (h HTMLBodyElement) GetOnonline() EventHandler {
	val := h.Get("ononline")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnonline(val EventHandler) {
	h.Set("ononline", val)
}
func (h HTMLBodyElement) GetOnpagehide() EventHandler {
	val := h.Get("onpagehide")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnpagehide(val EventHandler) {
	h.Set("onpagehide", val)
}
func (h HTMLBodyElement) GetOnpageshow() EventHandler {
	val := h.Get("onpageshow")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnpageshow(val EventHandler) {
	h.Set("onpageshow", val)
}
func (h HTMLBodyElement) GetOnpopstate() EventHandler {
	val := h.Get("onpopstate")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnpopstate(val EventHandler) {
	h.Set("onpopstate", val)
}
func (h HTMLBodyElement) GetOnrejectionhandled() EventHandler {
	val := h.Get("onrejectionhandled")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnrejectionhandled(val EventHandler) {
	h.Set("onrejectionhandled", val)
}
func (h HTMLBodyElement) GetOnstorage() EventHandler {
	val := h.Get("onstorage")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnstorage(val EventHandler) {
	h.Set("onstorage", val)
}
func (h HTMLBodyElement) GetOnunhandledrejection() EventHandler {
	val := h.Get("onunhandledrejection")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnunhandledrejection(val EventHandler) {
	h.Set("onunhandledrejection", val)
}
func (h HTMLBodyElement) GetOnunload() EventHandler {
	val := h.Get("onunload")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnunload(val EventHandler) {
	h.Set("onunload", val)
}
func (h HTMLBodyElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLBodyElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLBodyElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLBodyElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLBodyElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLBodyElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLBodyElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLBodyElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLBodyElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLBodyElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLBodyElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLBodyElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLBodyElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLBodyElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLBodyElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLBodyElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLBodyElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLBodyElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLBodyElement) GetText() string {
	val := h.Get("text")
	return val.String()
}
func (h HTMLBodyElement) SetText(val string) {
	h.Set("text", val)
}
func (h HTMLBodyElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLBodyElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLBodyElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLBodyElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLBodyElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLBodyElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLBodyElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLBodyElement) GetVLink() string {
	val := h.Get("vLink")
	return val.String()
}
func (h HTMLBodyElement) SetVLink(val string) {
	h.Set("vLink", val)
}
func (h HTMLBodyElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
