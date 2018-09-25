// Code generated DO NOT EDIT
// htmlmarqueeelement.go
package dom

import "syscall/js"

type HTMLMarqueeElementIFace interface {
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
	GetBehavior() string
	SetBehavior(string)
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
	GetDirection() string
	SetDirection(string)
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
	GetHeight() string
	SetHeight(string)
	GetHidden() bool
	SetHidden(bool)
	GetHspace() int
	SetHspace(int)
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
	GetLoop() int
	SetLoop(int)
	Matches(args ...interface{}) bool
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOnbounce() EventHandler
	SetOnbounce(EventHandler)
	GetOnfinish() EventHandler
	SetOnfinish(EventHandler)
	GetOnstart() EventHandler
	SetOnstart(EventHandler)
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
	GetScrollAmount() int
	SetScrollAmount(int)
	GetScrollDelay() int
	SetScrollDelay(int)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	Start(args ...interface{})
	Stop(args ...interface{})
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetTrueSpeed() bool
	SetTrueSpeed(bool)
	GetVspace() int
	SetVspace(int)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() string
	SetWidth(string)
}
type HTMLMarqueeElement struct {
	Value
}

func JSValueToHTMLMarqueeElement(val js.Value) HTMLMarqueeElement {
	return HTMLMarqueeElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLMarqueeElement() HTMLMarqueeElement { return HTMLMarqueeElement{Value: v} }
func NewHTMLMarqueeElement(args ...interface{}) HTMLMarqueeElement {
	return HTMLMarqueeElement{Value: JSValueToValue(js.Global().Get("HTMLMarqueeElement").New(args...))}
}
func (h HTMLMarqueeElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLMarqueeElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLMarqueeElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLMarqueeElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLMarqueeElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLMarqueeElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLMarqueeElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLMarqueeElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLMarqueeElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLMarqueeElement) GetBehavior() string {
	val := h.Get("behavior")
	return val.String()
}
func (h HTMLMarqueeElement) SetBehavior(val string) {
	h.Set("behavior", val)
}
func (h HTMLMarqueeElement) GetBgColor() string {
	val := h.Get("bgColor")
	return val.String()
}
func (h HTMLMarqueeElement) SetBgColor(val string) {
	h.Set("bgColor", val)
}
func (h HTMLMarqueeElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLMarqueeElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLMarqueeElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLMarqueeElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLMarqueeElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLMarqueeElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLMarqueeElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLMarqueeElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLMarqueeElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLMarqueeElement) GetDirection() string {
	val := h.Get("direction")
	return val.String()
}
func (h HTMLMarqueeElement) SetDirection(val string) {
	h.Set("direction", val)
}
func (h HTMLMarqueeElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLMarqueeElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLMarqueeElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLMarqueeElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLMarqueeElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLMarqueeElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMarqueeElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMarqueeElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLMarqueeElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLMarqueeElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLMarqueeElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) GetHeight() string {
	val := h.Get("height")
	return val.String()
}
func (h HTMLMarqueeElement) SetHeight(val string) {
	h.Set("height", val)
}
func (h HTMLMarqueeElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLMarqueeElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLMarqueeElement) GetHspace() int {
	val := h.Get("hspace")
	return val.Int()
}
func (h HTMLMarqueeElement) SetHspace(val int) {
	h.Set("hspace", val)
}
func (h HTMLMarqueeElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLMarqueeElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLMarqueeElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLMarqueeElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLMarqueeElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLMarqueeElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLMarqueeElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLMarqueeElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLMarqueeElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLMarqueeElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLMarqueeElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLMarqueeElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLMarqueeElement) GetLoop() int {
	val := h.Get("loop")
	return val.Int()
}
func (h HTMLMarqueeElement) SetLoop(val int) {
	h.Set("loop", val)
}
func (h HTMLMarqueeElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLMarqueeElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLMarqueeElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLMarqueeElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLMarqueeElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLMarqueeElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLMarqueeElement) GetOnbounce() EventHandler {
	val := h.Get("onbounce")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLMarqueeElement) SetOnbounce(val EventHandler) {
	h.Set("onbounce", val)
}
func (h HTMLMarqueeElement) GetOnfinish() EventHandler {
	val := h.Get("onfinish")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLMarqueeElement) SetOnfinish(val EventHandler) {
	h.Set("onfinish", val)
}
func (h HTMLMarqueeElement) GetOnstart() EventHandler {
	val := h.Get("onstart")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLMarqueeElement) SetOnstart(val EventHandler) {
	h.Set("onstart", val)
}
func (h HTMLMarqueeElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLMarqueeElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLMarqueeElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLMarqueeElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLMarqueeElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLMarqueeElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMarqueeElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLMarqueeElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMarqueeElement) GetScrollAmount() int {
	val := h.Get("scrollAmount")
	return val.Int()
}
func (h HTMLMarqueeElement) SetScrollAmount(val int) {
	h.Set("scrollAmount", val)
}
func (h HTMLMarqueeElement) GetScrollDelay() int {
	val := h.Get("scrollDelay")
	return val.Int()
}
func (h HTMLMarqueeElement) SetScrollDelay(val int) {
	h.Set("scrollDelay", val)
}
func (h HTMLMarqueeElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLMarqueeElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLMarqueeElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMarqueeElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMarqueeElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLMarqueeElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLMarqueeElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLMarqueeElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLMarqueeElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLMarqueeElement) Start(args ...interface{}) {
	h.Call("start", args...)
}
func (h HTMLMarqueeElement) Stop(args ...interface{}) {
	h.Call("stop", args...)
}
func (h HTMLMarqueeElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLMarqueeElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLMarqueeElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLMarqueeElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLMarqueeElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLMarqueeElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLMarqueeElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLMarqueeElement) GetTrueSpeed() bool {
	val := h.Get("trueSpeed")
	return val.Bool()
}
func (h HTMLMarqueeElement) SetTrueSpeed(val bool) {
	h.Set("trueSpeed", val)
}
func (h HTMLMarqueeElement) GetVspace() int {
	val := h.Get("vspace")
	return val.Int()
}
func (h HTMLMarqueeElement) SetVspace(val int) {
	h.Set("vspace", val)
}
func (h HTMLMarqueeElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLMarqueeElement) GetWidth() string {
	val := h.Get("width")
	return val.String()
}
func (h HTMLMarqueeElement) SetWidth(val string) {
	h.Set("width", val)
}
