// Code generated DO NOT EDIT
// htmlimageelement.go
package dom

import "syscall/js"

type HTMLImageElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	GetAlign() string
	SetAlign(string)
	GetAlt() string
	SetAlt(string)
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetBaseURI() string
	GetBorder() string
	SetBorder(string)
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	GetComplete() bool
	Contains(args ...interface{}) bool
	GetCrossOrigin() string
	SetCrossOrigin(string)
	GetCurrentSrc() string
	Decode(args ...interface{})
	GetDecoding() string
	SetDecoding(string)
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
	GetHeight() int
	SetHeight(int)
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
	GetIsMap() bool
	SetIsMap(bool)
	IsSameNode(args ...interface{}) bool
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	GetLocalName() string
	GetLongDesc() string
	SetLongDesc(string)
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetLowsrc() string
	SetLowsrc(string)
	Matches(args ...interface{}) bool
	GetName() string
	SetName(string)
	GetNamespaceURI() string
	GetNaturalHeight() int
	GetNaturalWidth() int
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
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSizes() string
	SetSizes(string)
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetSrc() string
	SetSrc(string)
	GetSrcset() string
	SetSrcset(string)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetUseMap() string
	SetUseMap(string)
	GetVspace() int
	SetVspace(int)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() int
	SetWidth(int)
}
type HTMLImageElement struct {
	Value
}

func JSValueToHTMLImageElement(val js.Value) HTMLImageElement {
	return HTMLImageElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLImageElement() HTMLImageElement { return HTMLImageElement{Value: v} }
func NewHTMLImageElement(args ...interface{}) HTMLImageElement {
	return HTMLImageElement{Value: JSValueToValue(js.Global().Get("HTMLImageElement").New(args...))}
}
func (h HTMLImageElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLImageElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLImageElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLImageElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLImageElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLImageElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLImageElement) GetAlt() string {
	val := h.Get("alt")
	return val.String()
}
func (h HTMLImageElement) SetAlt(val string) {
	h.Set("alt", val)
}
func (h HTMLImageElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLImageElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLImageElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLImageElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLImageElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLImageElement) GetBorder() string {
	val := h.Get("border")
	return val.String()
}
func (h HTMLImageElement) SetBorder(val string) {
	h.Set("border", val)
}
func (h HTMLImageElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLImageElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLImageElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLImageElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLImageElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLImageElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLImageElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLImageElement) GetComplete() bool {
	val := h.Get("complete")
	return val.Bool()
}
func (h HTMLImageElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLImageElement) GetCrossOrigin() string {
	val := h.Get("crossOrigin")
	return val.String()
}
func (h HTMLImageElement) SetCrossOrigin(val string) {
	h.Set("crossOrigin", val)
}
func (h HTMLImageElement) GetCurrentSrc() string {
	val := h.Get("currentSrc")
	return val.String()
}
func (h HTMLImageElement) Decode(args ...interface{}) {
	h.Call("decode", args...)
}
func (h HTMLImageElement) GetDecoding() string {
	val := h.Get("decoding")
	return val.String()
}
func (h HTMLImageElement) SetDecoding(val string) {
	h.Set("decoding", val)
}
func (h HTMLImageElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLImageElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLImageElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLImageElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLImageElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLImageElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLImageElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLImageElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLImageElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLImageElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLImageElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLImageElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLImageElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLImageElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLImageElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLImageElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLImageElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLImageElement) GetHeight() int {
	val := h.Get("height")
	return val.Int()
}
func (h HTMLImageElement) SetHeight(val int) {
	h.Set("height", val)
}
func (h HTMLImageElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLImageElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLImageElement) GetHspace() int {
	val := h.Get("hspace")
	return val.Int()
}
func (h HTMLImageElement) SetHspace(val int) {
	h.Set("hspace", val)
}
func (h HTMLImageElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLImageElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLImageElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLImageElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLImageElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLImageElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLImageElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLImageElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLImageElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLImageElement) GetIsMap() bool {
	val := h.Get("isMap")
	return val.Bool()
}
func (h HTMLImageElement) SetIsMap(val bool) {
	h.Set("isMap", val)
}
func (h HTMLImageElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLImageElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLImageElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLImageElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLImageElement) GetLongDesc() string {
	val := h.Get("longDesc")
	return val.String()
}
func (h HTMLImageElement) SetLongDesc(val string) {
	h.Set("longDesc", val)
}
func (h HTMLImageElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLImageElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLImageElement) GetLowsrc() string {
	val := h.Get("lowsrc")
	return val.String()
}
func (h HTMLImageElement) SetLowsrc(val string) {
	h.Set("lowsrc", val)
}
func (h HTMLImageElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLImageElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLImageElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLImageElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLImageElement) GetNaturalHeight() int {
	val := h.Get("naturalHeight")
	return val.Int()
}
func (h HTMLImageElement) GetNaturalWidth() int {
	val := h.Get("naturalWidth")
	return val.Int()
}
func (h HTMLImageElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLImageElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLImageElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLImageElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLImageElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLImageElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLImageElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLImageElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLImageElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLImageElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLImageElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLImageElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLImageElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLImageElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLImageElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLImageElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLImageElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLImageElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLImageElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLImageElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLImageElement) GetSizes() string {
	val := h.Get("sizes")
	return val.String()
}
func (h HTMLImageElement) SetSizes(val string) {
	h.Set("sizes", val)
}
func (h HTMLImageElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLImageElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLImageElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLImageElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLImageElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLImageElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLImageElement) GetSrcset() string {
	val := h.Get("srcset")
	return val.String()
}
func (h HTMLImageElement) SetSrcset(val string) {
	h.Set("srcset", val)
}
func (h HTMLImageElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLImageElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLImageElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLImageElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLImageElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLImageElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLImageElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLImageElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLImageElement) GetUseMap() string {
	val := h.Get("useMap")
	return val.String()
}
func (h HTMLImageElement) SetUseMap(val string) {
	h.Set("useMap", val)
}
func (h HTMLImageElement) GetVspace() int {
	val := h.Get("vspace")
	return val.Int()
}
func (h HTMLImageElement) SetVspace(val int) {
	h.Set("vspace", val)
}
func (h HTMLImageElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLImageElement) GetWidth() int {
	val := h.Get("width")
	return val.Int()
}
func (h HTMLImageElement) SetWidth(val int) {
	h.Set("width", val)
}
