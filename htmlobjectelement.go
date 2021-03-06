// Code generated DO NOT EDIT
// htmlobjectelement.go
package dom

import "syscall/js"

type HTMLObjectElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	GetAlign() string
	SetAlign(string)
	AppendChild(args ...interface{}) Node
	GetArchive() string
	SetArchive(string)
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetBaseURI() string
	GetBorder() string
	SetBorder(string)
	CheckValidity(args ...interface{}) bool
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	GetCode() string
	SetCode(string)
	GetCodeBase() string
	SetCodeBase(string)
	GetCodeType() string
	SetCodeType(string)
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	GetContentDocument() Document
	GetContentWindow() WindowProxy
	GetData() string
	SetData(string)
	GetDeclare() bool
	SetDeclare(bool)
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetFirstChild() Node
	GetForm() HTMLFormElement
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
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	ReportValidity(args ...interface{}) bool
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	SetCustomValidity(args ...interface{})
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetStandby() string
	SetStandby(string)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetType() string
	SetType(string)
	GetTypeMustMatch() bool
	SetTypeMustMatch(bool)
	GetUseMap() string
	SetUseMap(string)
	GetValidationMessage() string
	GetValidity() ValidityState
	GetVspace() int
	SetVspace(int)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() string
	SetWidth(string)
	GetWillValidate() bool
}
type HTMLObjectElement struct {
	Value
}

func JSValueToHTMLObjectElement(val js.Value) HTMLObjectElement {
	return HTMLObjectElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLObjectElement() HTMLObjectElement { return HTMLObjectElement{Value: v} }
func NewHTMLObjectElement(args ...interface{}) HTMLObjectElement {
	return HTMLObjectElement{Value: JSValueToValue(js.Global().Get("HTMLObjectElement").New(args...))}
}
func (h HTMLObjectElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLObjectElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLObjectElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLObjectElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLObjectElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLObjectElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLObjectElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) GetArchive() string {
	val := h.Get("archive")
	return val.String()
}
func (h HTMLObjectElement) SetArchive(val string) {
	h.Set("archive", val)
}
func (h HTMLObjectElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLObjectElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLObjectElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLObjectElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLObjectElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLObjectElement) GetBorder() string {
	val := h.Get("border")
	return val.String()
}
func (h HTMLObjectElement) SetBorder(val string) {
	h.Set("border", val)
}
func (h HTMLObjectElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLObjectElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLObjectElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLObjectElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLObjectElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLObjectElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLObjectElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLObjectElement) GetCode() string {
	val := h.Get("code")
	return val.String()
}
func (h HTMLObjectElement) SetCode(val string) {
	h.Set("code", val)
}
func (h HTMLObjectElement) GetCodeBase() string {
	val := h.Get("codeBase")
	return val.String()
}
func (h HTMLObjectElement) SetCodeBase(val string) {
	h.Set("codeBase", val)
}
func (h HTMLObjectElement) GetCodeType() string {
	val := h.Get("codeType")
	return val.String()
}
func (h HTMLObjectElement) SetCodeType(val string) {
	h.Set("codeType", val)
}
func (h HTMLObjectElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLObjectElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLObjectElement) GetContentDocument() Document {
	val := h.Get("contentDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLObjectElement) GetContentWindow() WindowProxy {
	val := h.Get("contentWindow")
	return JSValueToWindowProxy(val.JSValue())
}
func (h HTMLObjectElement) GetData() string {
	val := h.Get("data")
	return val.String()
}
func (h HTMLObjectElement) SetData(val string) {
	h.Set("data", val)
}
func (h HTMLObjectElement) GetDeclare() bool {
	val := h.Get("declare")
	return val.Bool()
}
func (h HTMLObjectElement) SetDeclare(val bool) {
	h.Set("declare", val)
}
func (h HTMLObjectElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLObjectElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLObjectElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLObjectElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLObjectElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLObjectElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLObjectElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLObjectElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLObjectElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLObjectElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLObjectElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLObjectElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLObjectElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLObjectElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLObjectElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) GetSVGDocument(args ...interface{}) Document {
	val := h.Call("getSVGDocument", args...)
	return JSValueToDocument(val.JSValue())
}
func (h HTMLObjectElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLObjectElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLObjectElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLObjectElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLObjectElement) GetHeight() string {
	val := h.Get("height")
	return val.String()
}
func (h HTMLObjectElement) SetHeight(val string) {
	h.Set("height", val)
}
func (h HTMLObjectElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLObjectElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLObjectElement) GetHspace() int {
	val := h.Get("hspace")
	return val.Int()
}
func (h HTMLObjectElement) SetHspace(val int) {
	h.Set("hspace", val)
}
func (h HTMLObjectElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLObjectElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLObjectElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLObjectElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLObjectElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLObjectElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLObjectElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLObjectElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLObjectElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLObjectElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLObjectElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLObjectElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLObjectElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLObjectElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLObjectElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLObjectElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLObjectElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLObjectElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLObjectElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLObjectElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLObjectElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLObjectElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLObjectElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLObjectElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLObjectElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLObjectElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLObjectElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLObjectElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLObjectElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLObjectElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLObjectElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLObjectElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLObjectElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLObjectElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLObjectElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLObjectElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLObjectElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLObjectElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLObjectElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLObjectElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLObjectElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLObjectElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLObjectElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLObjectElement) GetStandby() string {
	val := h.Get("standby")
	return val.String()
}
func (h HTMLObjectElement) SetStandby(val string) {
	h.Set("standby", val)
}
func (h HTMLObjectElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLObjectElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLObjectElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLObjectElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLObjectElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLObjectElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLObjectElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLObjectElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLObjectElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLObjectElement) SetType(val string) {
	h.Set("type", val)
}
func (h HTMLObjectElement) GetTypeMustMatch() bool {
	val := h.Get("typeMustMatch")
	return val.Bool()
}
func (h HTMLObjectElement) SetTypeMustMatch(val bool) {
	h.Set("typeMustMatch", val)
}
func (h HTMLObjectElement) GetUseMap() string {
	val := h.Get("useMap")
	return val.String()
}
func (h HTMLObjectElement) SetUseMap(val string) {
	h.Set("useMap", val)
}
func (h HTMLObjectElement) GetValidationMessage() string {
	val := h.Get("validationMessage")
	return val.String()
}
func (h HTMLObjectElement) GetValidity() ValidityState {
	val := h.Get("validity")
	return JSValueToValidityState(val.JSValue())
}
func (h HTMLObjectElement) GetVspace() int {
	val := h.Get("vspace")
	return val.Int()
}
func (h HTMLObjectElement) SetVspace(val int) {
	h.Set("vspace", val)
}
func (h HTMLObjectElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLObjectElement) GetWidth() string {
	val := h.Get("width")
	return val.String()
}
func (h HTMLObjectElement) SetWidth(val string) {
	h.Set("width", val)
}
func (h HTMLObjectElement) GetWillValidate() bool {
	val := h.Get("willValidate")
	return val.Bool()
}
