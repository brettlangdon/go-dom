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
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLObjectElement(val js.Value) HTMLObjectElement {
	return HTMLObjectElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLObjectElement() HTMLObjectElement { return HTMLObjectElement{Value: v} }
func (h HTMLObjectElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLObjectElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLObjectElement) GetArchive() string {
	val := h.Get("archive")
	return val.String()
}
func (h HTMLObjectElement) SetArchive(val string) {
	h.Set("archive", val)
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
func (h HTMLObjectElement) GetForm() HTMLFormElement {
	val := h.Get("form")
	return JSValueToHTMLFormElement(val.JSValue())
}
func (h HTMLObjectElement) GetSVGDocument(args ...interface{}) Document {
	val := h.Call("getSVGDocument", args...)
	return JSValueToDocument(val.JSValue())
}
func (h HTMLObjectElement) GetHeight() string {
	val := h.Get("height")
	return val.String()
}
func (h HTMLObjectElement) SetHeight(val string) {
	h.Set("height", val)
}
func (h HTMLObjectElement) GetHspace() int {
	val := h.Get("hspace")
	return val.Int()
}
func (h HTMLObjectElement) SetHspace(val int) {
	h.Set("hspace", val)
}
func (h HTMLObjectElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLObjectElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLObjectElement) ReportValidity(args ...interface{}) bool {
	val := h.Call("reportValidity", args...)
	return val.Bool()
}
func (h HTMLObjectElement) SetCustomValidity(args ...interface{}) {
	h.Call("setCustomValidity", args...)
}
func (h HTMLObjectElement) GetStandby() string {
	val := h.Get("standby")
	return val.String()
}
func (h HTMLObjectElement) SetStandby(val string) {
	h.Set("standby", val)
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
