// Code generated DO NOT EDIT
// htmliframeelement.go
package dom

import "syscall/js"

type HTMLIFrameElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	GetAlign() string
	SetAlign(string)
	GetAllow() string
	SetAllow(string)
	GetAllowFullscreen() bool
	SetAllowFullscreen(bool)
	GetAllowPaymentRequest() bool
	SetAllowPaymentRequest(bool)
	GetAllowUserMedia() bool
	SetAllowUserMedia(bool)
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
	GetContentDocument() Document
	GetContentWindow() WindowProxy
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetFirstChild() Node
	GetFrameBorder() string
	SetFrameBorder(string)
	GetAttribute(args ...interface{}) string
	GetAttributeNS(args ...interface{}) string
	GetAttributeNames(args ...interface{})
	GetAttributeNode(args ...interface{}) Attr
	GetAttributeNodeNS(args ...interface{}) Attr
	GetElementsByClassName(args ...interface{}) HTMLCollection
	GetElementsByTagName(args ...interface{}) HTMLCollection
	GetElementsByTagNameNS(args ...interface{}) HTMLCollection
	GetRootNode(args ...interface{}) Node
	GetSVGDocument(args ...interface{}) Document
	HasAttribute(args ...interface{}) bool
	HasAttributeNS(args ...interface{}) bool
	HasAttributes(args ...interface{}) bool
	HasChildNodes(args ...interface{}) bool
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
	GetLongDesc() string
	SetLongDesc(string)
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetMarginHeight() string
	SetMarginHeight(string)
	GetMarginWidth() string
	SetMarginWidth(string)
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
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPrefix() string
	GetPreviousSibling() Node
	GetReferrerPolicy() string
	SetReferrerPolicy(string)
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetSandbox() DOMTokenList
	GetScrolling() string
	SetScrolling(string)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetSrc() string
	SetSrc(string)
	GetSrcdoc() string
	SetSrcdoc(string)
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
type HTMLIFrameElement struct {
	Value
}

func JSValueToHTMLIFrameElement(val js.Value) HTMLIFrameElement {
	return HTMLIFrameElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLIFrameElement() HTMLIFrameElement { return HTMLIFrameElement{Value: v} }
func NewHTMLIFrameElement(args ...interface{}) HTMLIFrameElement {
	return HTMLIFrameElement{Value: JSValueToValue(js.Global().Get("HTMLIFrameElement").New(args...))}
}
func (h HTMLIFrameElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLIFrameElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLIFrameElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLIFrameElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLIFrameElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLIFrameElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLIFrameElement) GetAllow() string {
	val := h.Get("allow")
	return val.String()
}
func (h HTMLIFrameElement) SetAllow(val string) {
	h.Set("allow", val)
}
func (h HTMLIFrameElement) GetAllowFullscreen() bool {
	val := h.Get("allowFullscreen")
	return val.Bool()
}
func (h HTMLIFrameElement) SetAllowFullscreen(val bool) {
	h.Set("allowFullscreen", val)
}
func (h HTMLIFrameElement) GetAllowPaymentRequest() bool {
	val := h.Get("allowPaymentRequest")
	return val.Bool()
}
func (h HTMLIFrameElement) SetAllowPaymentRequest(val bool) {
	h.Set("allowPaymentRequest", val)
}
func (h HTMLIFrameElement) GetAllowUserMedia() bool {
	val := h.Get("allowUserMedia")
	return val.Bool()
}
func (h HTMLIFrameElement) SetAllowUserMedia(val bool) {
	h.Set("allowUserMedia", val)
}
func (h HTMLIFrameElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLIFrameElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLIFrameElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLIFrameElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLIFrameElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLIFrameElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLIFrameElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLIFrameElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLIFrameElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLIFrameElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLIFrameElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLIFrameElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLIFrameElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) GetContentDocument() Document {
	val := h.Get("contentDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLIFrameElement) GetContentWindow() WindowProxy {
	val := h.Get("contentWindow")
	return JSValueToWindowProxy(val.JSValue())
}
func (h HTMLIFrameElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLIFrameElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLIFrameElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLIFrameElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLIFrameElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) GetFrameBorder() string {
	val := h.Get("frameBorder")
	return val.String()
}
func (h HTMLIFrameElement) SetFrameBorder(val string) {
	h.Set("frameBorder", val)
}
func (h HTMLIFrameElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLIFrameElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLIFrameElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLIFrameElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLIFrameElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLIFrameElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLIFrameElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLIFrameElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLIFrameElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) GetSVGDocument(args ...interface{}) Document {
	val := h.Call("getSVGDocument", args...)
	return JSValueToDocument(val.JSValue())
}
func (h HTMLIFrameElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) GetHeight() string {
	val := h.Get("height")
	return val.String()
}
func (h HTMLIFrameElement) SetHeight(val string) {
	h.Set("height", val)
}
func (h HTMLIFrameElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLIFrameElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLIFrameElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLIFrameElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLIFrameElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLIFrameElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLIFrameElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLIFrameElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLIFrameElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLIFrameElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLIFrameElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLIFrameElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLIFrameElement) GetLongDesc() string {
	val := h.Get("longDesc")
	return val.String()
}
func (h HTMLIFrameElement) SetLongDesc(val string) {
	h.Set("longDesc", val)
}
func (h HTMLIFrameElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLIFrameElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLIFrameElement) GetMarginHeight() string {
	val := h.Get("marginHeight")
	return val.String()
}
func (h HTMLIFrameElement) SetMarginHeight(val string) {
	h.Set("marginHeight", val)
}
func (h HTMLIFrameElement) GetMarginWidth() string {
	val := h.Get("marginWidth")
	return val.String()
}
func (h HTMLIFrameElement) SetMarginWidth(val string) {
	h.Set("marginWidth", val)
}
func (h HTMLIFrameElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLIFrameElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLIFrameElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLIFrameElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLIFrameElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLIFrameElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLIFrameElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLIFrameElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLIFrameElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLIFrameElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLIFrameElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLIFrameElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLIFrameElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLIFrameElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLIFrameElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLIFrameElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLIFrameElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLIFrameElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLIFrameElement) GetSandbox() DOMTokenList {
	val := h.Get("sandbox")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLIFrameElement) GetScrolling() string {
	val := h.Get("scrolling")
	return val.String()
}
func (h HTMLIFrameElement) SetScrolling(val string) {
	h.Set("scrolling", val)
}
func (h HTMLIFrameElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLIFrameElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLIFrameElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLIFrameElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLIFrameElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLIFrameElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLIFrameElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLIFrameElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLIFrameElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLIFrameElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLIFrameElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLIFrameElement) GetSrcdoc() string {
	val := h.Get("srcdoc")
	return val.String()
}
func (h HTMLIFrameElement) SetSrcdoc(val string) {
	h.Set("srcdoc", val)
}
func (h HTMLIFrameElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLIFrameElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLIFrameElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLIFrameElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLIFrameElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLIFrameElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLIFrameElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLIFrameElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLIFrameElement) GetWidth() string {
	val := h.Get("width")
	return val.String()
}
func (h HTMLIFrameElement) SetWidth(val string) {
	h.Set("width", val)
}
