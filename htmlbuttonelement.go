// Code generated DO NOT EDIT
// htmlbuttonelement.go
package dom

import "syscall/js"

type HTMLButtonElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetAutofocus() bool
	SetAutofocus(bool)
	GetBaseURI() string
	CheckValidity(args ...interface{}) bool
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
	GetDisabled() bool
	SetDisabled(bool)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetFirstChild() Node
	GetForm() HTMLFormElement
	GetFormAction() string
	SetFormAction(string)
	GetFormEnctype() string
	SetFormEnctype(string)
	GetFormMethod() string
	SetFormMethod(string)
	GetFormNoValidate() bool
	SetFormNoValidate(bool)
	GetFormTarget() string
	SetFormTarget(string)
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
	GetLabels() NodeList
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
	GetValidationMessage() string
	GetValidity() ValidityState
	GetValue() string
	SetValue(string)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWillValidate() bool
}
type HTMLButtonElement struct {
	Value
}

func JSValueToHTMLButtonElement(val js.Value) HTMLButtonElement {
	return HTMLButtonElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLButtonElement() HTMLButtonElement { return HTMLButtonElement{Value: v} }
func NewHTMLButtonElement(args ...interface{}) HTMLButtonElement {
	return HTMLButtonElement{Value: JSValueToValue(js.Global().Get("HTMLButtonElement").New(args...))}
}
func (h HTMLButtonElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLButtonElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLButtonElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLButtonElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLButtonElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLButtonElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLButtonElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLButtonElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLButtonElement) GetAutofocus() bool {
	val := h.Get("autofocus")
	return val.Bool()
}
func (h HTMLButtonElement) SetAutofocus(val bool) {
	h.Set("autofocus", val)
}
func (h HTMLButtonElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLButtonElement) CheckValidity(args ...interface{}) bool {
	val := h.Call("checkValidity", args...)
	return val.Bool()
}
func (h HTMLButtonElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLButtonElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLButtonElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLButtonElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLButtonElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLButtonElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLButtonElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLButtonElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLButtonElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLButtonElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLButtonElement) GetDisabled() bool {
	val := h.Get("disabled")
	return val.Bool()
}
func (h HTMLButtonElement) SetDisabled(val bool) {
	h.Set("disabled", val)
}
func (h HTMLButtonElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLButtonElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLButtonElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLButtonElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLButtonElement) GetFormAction() string {
	val := h.Get("formAction")
	return val.String()
}
func (h HTMLButtonElement) SetFormAction(val string) {
	h.Set("formAction", val)
}
func (h HTMLButtonElement) GetFormEnctype() string {
	val := h.Get("formEnctype")
	return val.String()
}
func (h HTMLButtonElement) SetFormEnctype(val string) {
	h.Set("formEnctype", val)
}
func (h HTMLButtonElement) GetFormMethod() string {
	val := h.Get("formMethod")
	return val.String()
}
func (h HTMLButtonElement) SetFormMethod(val string) {
	h.Set("formMethod", val)
}
func (h HTMLButtonElement) GetFormNoValidate() bool {
	val := h.Get("formNoValidate")
	return val.Bool()
}
func (h HTMLButtonElement) SetFormNoValidate(val bool) {
	h.Set("formNoValidate", val)
}
func (h HTMLButtonElement) GetFormTarget() string {
	val := h.Get("formTarget")
	return val.String()
}
func (h HTMLButtonElement) SetFormTarget(val string) {
	h.Set("formTarget", val)
}
func (h HTMLButtonElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLButtonElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLButtonElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLButtonElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLButtonElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLButtonElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLButtonElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLButtonElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLButtonElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLButtonElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLButtonElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLButtonElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLButtonElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLButtonElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLButtonElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLButtonElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLButtonElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLButtonElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLButtonElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLButtonElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLButtonElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLButtonElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLButtonElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLButtonElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLButtonElement) GetLabels() NodeList {
	val := h.Get("labels")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLButtonElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLButtonElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLButtonElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLButtonElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLButtonElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLButtonElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLButtonElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLButtonElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLButtonElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLButtonElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLButtonElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLButtonElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLButtonElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLButtonElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLButtonElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLButtonElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLButtonElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLButtonElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLButtonElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLButtonElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLButtonElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLButtonElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLButtonElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLButtonElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLButtonElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLButtonElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLButtonElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLButtonElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLButtonElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLButtonElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLButtonElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLButtonElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLButtonElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLButtonElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLButtonElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLButtonElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLButtonElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLButtonElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLButtonElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLButtonElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLButtonElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLButtonElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLButtonElement) SetType(val string) {
	h.Set("type", val)
}
func (h HTMLButtonElement) GetValidationMessage() string {
	val := h.Get("validationMessage")
	return val.String()
}
func (h HTMLButtonElement) GetValidity() ValidityState {
	val := h.Get("validity")
	return JSValueToValidityState(val.JSValue())
}
func (h HTMLButtonElement) GetValue() string {
	val := h.Get("value")
	return val.String()
}
func (h HTMLButtonElement) SetValue(val string) {
	h.Set("value", val)
}
func (h HTMLButtonElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLButtonElement) GetWillValidate() bool {
	val := h.Get("willValidate")
	return val.Bool()
}
