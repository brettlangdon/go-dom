// Code generated DO NOT EDIT
// htmltrackelement.go
package dom

import "syscall/js"

type HTMLTrackElementIFace interface {
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
	GetDefault() bool
	SetDefault(bool)
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
	GetKind() string
	SetKind(string)
	GetLabel() string
	SetLabel(string)
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
	GetReadyState() int
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
	GetSrc() string
	SetSrc(string)
	GetSrclang() string
	SetSrclang(string)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTrack() TextTrack
	GetTranslate() bool
	SetTranslate(bool)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLTrackElement struct {
	Value
}

func JSValueToHTMLTrackElement(val js.Value) HTMLTrackElement {
	return HTMLTrackElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLTrackElement() HTMLTrackElement { return HTMLTrackElement{Value: v} }
func NewHTMLTrackElement(args ...interface{}) HTMLTrackElement {
	return HTMLTrackElement{Value: JSValueToValue(js.Global().Get("HTMLTrackElement").New(args...))}
}
func (h HTMLTrackElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLTrackElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLTrackElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLTrackElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLTrackElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTrackElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLTrackElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLTrackElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLTrackElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLTrackElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLTrackElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLTrackElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLTrackElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLTrackElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLTrackElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTrackElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLTrackElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLTrackElement) GetDefault() bool {
	val := h.Get("default")
	return val.Bool()
}
func (h HTMLTrackElement) SetDefault(val bool) {
	h.Set("default", val)
}
func (h HTMLTrackElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLTrackElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLTrackElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLTrackElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLTrackElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLTrackElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLTrackElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLTrackElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLTrackElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTrackElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTrackElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTrackElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTrackElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTrackElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLTrackElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLTrackElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLTrackElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLTrackElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLTrackElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLTrackElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLTrackElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLTrackElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLTrackElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLTrackElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLTrackElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLTrackElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLTrackElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLTrackElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLTrackElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLTrackElement) GetKind() string {
	val := h.Get("kind")
	return val.String()
}
func (h HTMLTrackElement) SetKind(val string) {
	h.Set("kind", val)
}
func (h HTMLTrackElement) GetLabel() string {
	val := h.Get("label")
	return val.String()
}
func (h HTMLTrackElement) SetLabel(val string) {
	h.Set("label", val)
}
func (h HTMLTrackElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLTrackElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLTrackElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLTrackElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLTrackElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLTrackElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLTrackElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLTrackElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLTrackElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLTrackElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLTrackElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLTrackElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLTrackElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLTrackElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLTrackElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLTrackElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) GetReadyState() int {
	val := h.Get("readyState")
	return val.Int()
}
func (h HTMLTrackElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLTrackElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLTrackElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTrackElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLTrackElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLTrackElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLTrackElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLTrackElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTrackElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLTrackElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLTrackElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLTrackElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLTrackElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLTrackElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLTrackElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLTrackElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLTrackElement) GetSrclang() string {
	val := h.Get("srclang")
	return val.String()
}
func (h HTMLTrackElement) SetSrclang(val string) {
	h.Set("srclang", val)
}
func (h HTMLTrackElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLTrackElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLTrackElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLTrackElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLTrackElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLTrackElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLTrackElement) GetTrack() TextTrack {
	val := h.Get("track")
	return JSValueToTextTrack(val.JSValue())
}
func (h HTMLTrackElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLTrackElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLTrackElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
