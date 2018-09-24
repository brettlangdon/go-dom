// Code generated DO NOT EDIT
// htmltableelement.go
package dom

import "syscall/js"

type HTMLTableElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	GetAlign() string
	SetAlign(string)
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetBaseURI() string
	GetBgColor() string
	SetBgColor(string)
	GetBorder() string
	SetBorder(string)
	GetCaption() HTMLTableCaptionElement
	SetCaption(HTMLTableCaptionElement)
	GetCellPadding() string
	SetCellPadding(string)
	GetCellSpacing() string
	SetCellSpacing(string)
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	CreateCaption(args ...interface{}) HTMLTableCaptionElement
	CreateTBody(args ...interface{}) HTMLTableSectionElement
	CreateTFoot(args ...interface{}) HTMLTableSectionElement
	CreateTHead(args ...interface{}) HTMLTableSectionElement
	DeleteCaption(args ...interface{})
	DeleteRow(args ...interface{})
	DeleteTFoot(args ...interface{})
	DeleteTHead(args ...interface{})
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetFirstChild() Node
	GetFrame() string
	SetFrame(string)
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
	InsertRow(args ...interface{}) HTMLTableRowElement
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
	GetRows() HTMLCollection
	GetRules() string
	SetRules(string)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetSummary() string
	SetSummary(string)
	GetTBodies() HTMLCollection
	GetTFoot() HTMLTableSectionElement
	SetTFoot(HTMLTableSectionElement)
	GetTHead() HTMLTableSectionElement
	SetTHead(HTMLTableSectionElement)
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
type HTMLTableElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLTableElement(val js.Value) HTMLTableElement {
	return HTMLTableElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLTableElement() HTMLTableElement { return HTMLTableElement{Value: v} }
func NewHTMLTableElement(args ...interface{}) HTMLTableElement {
	return HTMLTableElement{Value: JSValueToValue(js.Global().Get("HTMLTableElement").New(args...))}
}
func (h HTMLTableElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLTableElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLTableElement) GetBgColor() string {
	val := h.Get("bgColor")
	return val.String()
}
func (h HTMLTableElement) SetBgColor(val string) {
	h.Set("bgColor", val)
}
func (h HTMLTableElement) GetBorder() string {
	val := h.Get("border")
	return val.String()
}
func (h HTMLTableElement) SetBorder(val string) {
	h.Set("border", val)
}
func (h HTMLTableElement) GetCaption() HTMLTableCaptionElement {
	val := h.Get("caption")
	return JSValueToHTMLTableCaptionElement(val.JSValue())
}
func (h HTMLTableElement) SetCaption(val HTMLTableCaptionElement) {
	h.Set("caption", val)
}
func (h HTMLTableElement) GetCellPadding() string {
	val := h.Get("cellPadding")
	return val.String()
}
func (h HTMLTableElement) SetCellPadding(val string) {
	h.Set("cellPadding", val)
}
func (h HTMLTableElement) GetCellSpacing() string {
	val := h.Get("cellSpacing")
	return val.String()
}
func (h HTMLTableElement) SetCellSpacing(val string) {
	h.Set("cellSpacing", val)
}
func (h HTMLTableElement) CreateCaption(args ...interface{}) HTMLTableCaptionElement {
	val := h.Call("createCaption", args...)
	return JSValueToHTMLTableCaptionElement(val.JSValue())
}
func (h HTMLTableElement) CreateTBody(args ...interface{}) HTMLTableSectionElement {
	val := h.Call("createTBody", args...)
	return JSValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) CreateTFoot(args ...interface{}) HTMLTableSectionElement {
	val := h.Call("createTFoot", args...)
	return JSValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) CreateTHead(args ...interface{}) HTMLTableSectionElement {
	val := h.Call("createTHead", args...)
	return JSValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) DeleteCaption(args ...interface{}) {
	h.Call("deleteCaption", args...)
}
func (h HTMLTableElement) DeleteRow(args ...interface{}) {
	h.Call("deleteRow", args...)
}
func (h HTMLTableElement) DeleteTFoot(args ...interface{}) {
	h.Call("deleteTFoot", args...)
}
func (h HTMLTableElement) DeleteTHead(args ...interface{}) {
	h.Call("deleteTHead", args...)
}
func (h HTMLTableElement) GetFrame() string {
	val := h.Get("frame")
	return val.String()
}
func (h HTMLTableElement) SetFrame(val string) {
	h.Set("frame", val)
}
func (h HTMLTableElement) InsertRow(args ...interface{}) HTMLTableRowElement {
	val := h.Call("insertRow", args...)
	return JSValueToHTMLTableRowElement(val.JSValue())
}
func (h HTMLTableElement) GetRows() HTMLCollection {
	val := h.Get("rows")
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableElement) GetRules() string {
	val := h.Get("rules")
	return val.String()
}
func (h HTMLTableElement) SetRules(val string) {
	h.Set("rules", val)
}
func (h HTMLTableElement) GetSummary() string {
	val := h.Get("summary")
	return val.String()
}
func (h HTMLTableElement) SetSummary(val string) {
	h.Set("summary", val)
}
func (h HTMLTableElement) GetTBodies() HTMLCollection {
	val := h.Get("tBodies")
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableElement) GetTFoot() HTMLTableSectionElement {
	val := h.Get("tFoot")
	return JSValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) SetTFoot(val HTMLTableSectionElement) {
	h.Set("tFoot", val)
}
func (h HTMLTableElement) GetTHead() HTMLTableSectionElement {
	val := h.Get("tHead")
	return JSValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) SetTHead(val HTMLTableSectionElement) {
	h.Set("tHead", val)
}
func (h HTMLTableElement) GetWidth() string {
	val := h.Get("width")
	return val.String()
}
func (h HTMLTableElement) SetWidth(val string) {
	h.Set("width", val)
}
