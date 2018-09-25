// Code generated DO NOT EDIT
// htmlanchorelement.go
package dom

import "syscall/js"

type HTMLAnchorElementIFace interface {
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
	GetCharset() string
	SetCharset(string)
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
	GetHreflang() string
	SetHreflang(string)
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
	GetName() string
	SetName(string)
	GetNamespaceURI() string
	GetNextSibling() Node
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
	GetRev() string
	SetRev(string)
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
	GetText() string
	SetText(string)
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetType() string
	SetType(string)
	GetUsername() string
	SetUsername(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLAnchorElement struct {
	Value
}

func JSValueToHTMLAnchorElement(val js.Value) HTMLAnchorElement {
	return HTMLAnchorElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLAnchorElement() HTMLAnchorElement { return HTMLAnchorElement{Value: v} }
func NewHTMLAnchorElement(args ...interface{}) HTMLAnchorElement {
	return HTMLAnchorElement{Value: JSValueToValue(js.Global().Get("HTMLAnchorElement").New(args...))}
}
func (h HTMLAnchorElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLAnchorElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLAnchorElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLAnchorElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLAnchorElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLAnchorElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLAnchorElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLAnchorElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLAnchorElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLAnchorElement) GetCharset() string {
	val := h.Get("charset")
	return val.String()
}
func (h HTMLAnchorElement) SetCharset(val string) {
	h.Set("charset", val)
}
func (h HTMLAnchorElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLAnchorElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLAnchorElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLAnchorElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLAnchorElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLAnchorElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLAnchorElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLAnchorElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) GetCoords() string {
	val := h.Get("coords")
	return val.String()
}
func (h HTMLAnchorElement) SetCoords(val string) {
	h.Set("coords", val)
}
func (h HTMLAnchorElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLAnchorElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLAnchorElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) GetDownload() string {
	val := h.Get("download")
	return val.String()
}
func (h HTMLAnchorElement) SetDownload(val string) {
	h.Set("download", val)
}
func (h HTMLAnchorElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLAnchorElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLAnchorElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLAnchorElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLAnchorElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLAnchorElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAnchorElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAnchorElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLAnchorElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLAnchorElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLAnchorElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) GetHash() string {
	val := h.Get("hash")
	return val.String()
}
func (h HTMLAnchorElement) SetHash(val string) {
	h.Set("hash", val)
}
func (h HTMLAnchorElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLAnchorElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLAnchorElement) GetHost() string {
	val := h.Get("host")
	return val.String()
}
func (h HTMLAnchorElement) SetHost(val string) {
	h.Set("host", val)
}
func (h HTMLAnchorElement) GetHostname() string {
	val := h.Get("hostname")
	return val.String()
}
func (h HTMLAnchorElement) SetHostname(val string) {
	h.Set("hostname", val)
}
func (h HTMLAnchorElement) GetHref() string {
	val := h.Get("href")
	return val.String()
}
func (h HTMLAnchorElement) SetHref(val string) {
	h.Set("href", val)
}
func (h HTMLAnchorElement) GetHreflang() string {
	val := h.Get("hreflang")
	return val.String()
}
func (h HTMLAnchorElement) SetHreflang(val string) {
	h.Set("hreflang", val)
}
func (h HTMLAnchorElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLAnchorElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLAnchorElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLAnchorElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLAnchorElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLAnchorElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLAnchorElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLAnchorElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLAnchorElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLAnchorElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLAnchorElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLAnchorElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLAnchorElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLAnchorElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLAnchorElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLAnchorElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLAnchorElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLAnchorElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLAnchorElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLAnchorElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLAnchorElement) GetOrigin() string {
	val := h.Get("origin")
	return val.String()
}
func (h HTMLAnchorElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLAnchorElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLAnchorElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) GetPassword() string {
	val := h.Get("password")
	return val.String()
}
func (h HTMLAnchorElement) SetPassword(val string) {
	h.Set("password", val)
}
func (h HTMLAnchorElement) GetPathname() string {
	val := h.Get("pathname")
	return val.String()
}
func (h HTMLAnchorElement) SetPathname(val string) {
	h.Set("pathname", val)
}
func (h HTMLAnchorElement) GetPing() string {
	val := h.Get("ping")
	return val.String()
}
func (h HTMLAnchorElement) SetPing(val string) {
	h.Set("ping", val)
}
func (h HTMLAnchorElement) GetPort() string {
	val := h.Get("port")
	return val.String()
}
func (h HTMLAnchorElement) SetPort(val string) {
	h.Set("port", val)
}
func (h HTMLAnchorElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLAnchorElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) GetProtocol() string {
	val := h.Get("protocol")
	return val.String()
}
func (h HTMLAnchorElement) SetProtocol(val string) {
	h.Set("protocol", val)
}
func (h HTMLAnchorElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLAnchorElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLAnchorElement) GetRel() string {
	val := h.Get("rel")
	return val.String()
}
func (h HTMLAnchorElement) SetRel(val string) {
	h.Set("rel", val)
}
func (h HTMLAnchorElement) GetRelList() DOMTokenList {
	val := h.Get("relList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLAnchorElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLAnchorElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLAnchorElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAnchorElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLAnchorElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAnchorElement) GetRev() string {
	val := h.Get("rev")
	return val.String()
}
func (h HTMLAnchorElement) SetRev(val string) {
	h.Set("rev", val)
}
func (h HTMLAnchorElement) GetSearch() string {
	val := h.Get("search")
	return val.String()
}
func (h HTMLAnchorElement) SetSearch(val string) {
	h.Set("search", val)
}
func (h HTMLAnchorElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLAnchorElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLAnchorElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAnchorElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAnchorElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLAnchorElement) GetShape() string {
	val := h.Get("shape")
	return val.String()
}
func (h HTMLAnchorElement) SetShape(val string) {
	h.Set("shape", val)
}
func (h HTMLAnchorElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLAnchorElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLAnchorElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLAnchorElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLAnchorElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLAnchorElement) GetTarget() string {
	val := h.Get("target")
	return val.String()
}
func (h HTMLAnchorElement) SetTarget(val string) {
	h.Set("target", val)
}
func (h HTMLAnchorElement) GetText() string {
	val := h.Get("text")
	return val.String()
}
func (h HTMLAnchorElement) SetText(val string) {
	h.Set("text", val)
}
func (h HTMLAnchorElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLAnchorElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLAnchorElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLAnchorElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLAnchorElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLAnchorElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLAnchorElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLAnchorElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLAnchorElement) SetType(val string) {
	h.Set("type", val)
}
func (h HTMLAnchorElement) GetUsername() string {
	val := h.Get("username")
	return val.String()
}
func (h HTMLAnchorElement) SetUsername(val string) {
	h.Set("username", val)
}
func (h HTMLAnchorElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
