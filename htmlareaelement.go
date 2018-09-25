// Code generated DO NOT EDIT
// htmlareaelement.go
package dom

import "syscall/js"

type HTMLAreaElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	GetAlt() string
	SetAlt(string)
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
	GetCoords() string
	SetCoords(string)
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDownload() string
	SetDownload(string)
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
	GetHash() string
	SetHash(string)
	GetHidden() bool
	SetHidden(bool)
	GetHost() string
	SetHost(string)
	GetHostname() string
	SetHostname(string)
	GetHref() string
	SetHref(string)
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
	GetNoHref() bool
	SetNoHref(bool)
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOrigin() string
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPassword() string
	SetPassword(string)
	GetPathname() string
	SetPathname(string)
	GetPing() string
	SetPing(string)
	GetPort() string
	SetPort(string)
	GetPrefix() string
	GetPreviousSibling() Node
	GetProtocol() string
	SetProtocol(string)
	GetReferrerPolicy() string
	SetReferrerPolicy(string)
	GetRel() string
	SetRel(string)
	GetRelList() DOMTokenList
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetSearch() string
	SetSearch(string)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetShape() string
	SetShape(string)
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetTagName() string
	GetTarget() string
	SetTarget(string)
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetUsername() string
	SetUsername(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLAreaElement struct {
	Value
}

func JSValueToHTMLAreaElement(val js.Value) HTMLAreaElement {
	return HTMLAreaElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLAreaElement() HTMLAreaElement { return HTMLAreaElement{Value: v} }
func NewHTMLAreaElement(args ...interface{}) HTMLAreaElement {
	return HTMLAreaElement{Value: JSValueToValue(js.Global().Get("HTMLAreaElement").New(args...))}
}
func (h HTMLAreaElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLAreaElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLAreaElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLAreaElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLAreaElement) GetAlt() string {
	val := h.Get("alt")
	return val.String()
}
func (h HTMLAreaElement) SetAlt(val string) {
	h.Set("alt", val)
}
func (h HTMLAreaElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLAreaElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLAreaElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLAreaElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLAreaElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLAreaElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLAreaElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLAreaElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLAreaElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLAreaElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLAreaElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLAreaElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLAreaElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLAreaElement) GetCoords() string {
	val := h.Get("coords")
	return val.String()
}
func (h HTMLAreaElement) SetCoords(val string) {
	h.Set("coords", val)
}
func (h HTMLAreaElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLAreaElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLAreaElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLAreaElement) GetDownload() string {
	val := h.Get("download")
	return val.String()
}
func (h HTMLAreaElement) SetDownload(val string) {
	h.Set("download", val)
}
func (h HTMLAreaElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLAreaElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLAreaElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLAreaElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLAreaElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLAreaElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAreaElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAreaElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLAreaElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLAreaElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLAreaElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLAreaElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLAreaElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLAreaElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLAreaElement) GetHash() string {
	val := h.Get("hash")
	return val.String()
}
func (h HTMLAreaElement) SetHash(val string) {
	h.Set("hash", val)
}
func (h HTMLAreaElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLAreaElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLAreaElement) GetHost() string {
	val := h.Get("host")
	return val.String()
}
func (h HTMLAreaElement) SetHost(val string) {
	h.Set("host", val)
}
func (h HTMLAreaElement) GetHostname() string {
	val := h.Get("hostname")
	return val.String()
}
func (h HTMLAreaElement) SetHostname(val string) {
	h.Set("hostname", val)
}
func (h HTMLAreaElement) GetHref() string {
	val := h.Get("href")
	return val.String()
}
func (h HTMLAreaElement) SetHref(val string) {
	h.Set("href", val)
}
func (h HTMLAreaElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLAreaElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLAreaElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLAreaElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLAreaElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLAreaElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLAreaElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLAreaElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLAreaElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLAreaElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLAreaElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLAreaElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLAreaElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLAreaElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLAreaElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLAreaElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLAreaElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLAreaElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) GetNoHref() bool {
	val := h.Get("noHref")
	return val.Bool()
}
func (h HTMLAreaElement) SetNoHref(val bool) {
	h.Set("noHref", val)
}
func (h HTMLAreaElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLAreaElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLAreaElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLAreaElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLAreaElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLAreaElement) GetOrigin() string {
	val := h.Get("origin")
	return val.String()
}
func (h HTMLAreaElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLAreaElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLAreaElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) GetPassword() string {
	val := h.Get("password")
	return val.String()
}
func (h HTMLAreaElement) SetPassword(val string) {
	h.Set("password", val)
}
func (h HTMLAreaElement) GetPathname() string {
	val := h.Get("pathname")
	return val.String()
}
func (h HTMLAreaElement) SetPathname(val string) {
	h.Set("pathname", val)
}
func (h HTMLAreaElement) GetPing() string {
	val := h.Get("ping")
	return val.String()
}
func (h HTMLAreaElement) SetPing(val string) {
	h.Set("ping", val)
}
func (h HTMLAreaElement) GetPort() string {
	val := h.Get("port")
	return val.String()
}
func (h HTMLAreaElement) SetPort(val string) {
	h.Set("port", val)
}
func (h HTMLAreaElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLAreaElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) GetProtocol() string {
	val := h.Get("protocol")
	return val.String()
}
func (h HTMLAreaElement) SetProtocol(val string) {
	h.Set("protocol", val)
}
func (h HTMLAreaElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLAreaElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLAreaElement) GetRel() string {
	val := h.Get("rel")
	return val.String()
}
func (h HTMLAreaElement) SetRel(val string) {
	h.Set("rel", val)
}
func (h HTMLAreaElement) GetRelList() DOMTokenList {
	val := h.Get("relList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLAreaElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLAreaElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLAreaElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAreaElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLAreaElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAreaElement) GetSearch() string {
	val := h.Get("search")
	return val.String()
}
func (h HTMLAreaElement) SetSearch(val string) {
	h.Set("search", val)
}
func (h HTMLAreaElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLAreaElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLAreaElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAreaElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAreaElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLAreaElement) GetShape() string {
	val := h.Get("shape")
	return val.String()
}
func (h HTMLAreaElement) SetShape(val string) {
	h.Set("shape", val)
}
func (h HTMLAreaElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLAreaElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLAreaElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLAreaElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLAreaElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLAreaElement) GetTarget() string {
	val := h.Get("target")
	return val.String()
}
func (h HTMLAreaElement) SetTarget(val string) {
	h.Set("target", val)
}
func (h HTMLAreaElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLAreaElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLAreaElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLAreaElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLAreaElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLAreaElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLAreaElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLAreaElement) GetUsername() string {
	val := h.Get("username")
	return val.String()
}
func (h HTMLAreaElement) SetUsername(val string) {
	h.Set("username", val)
}
func (h HTMLAreaElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
