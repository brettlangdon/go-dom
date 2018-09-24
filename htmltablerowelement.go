// Code generated DO NOT EDIT
// htmltablerowelement.go
package dom

import "syscall/js"

type HTMLTableRowElementIFace interface {
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
	GetCells() HTMLCollection
	GetCh() string
	SetCh(string)
	GetChOff() string
	SetChOff(string)
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	DeleteCell(args ...interface{})
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
	InsertCell(args ...interface{}) HTMLTableCellElement
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
	GetRowIndex() int
	GetSectionRowIndex() int
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
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
	GetVAlign() string
	SetVAlign(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLTableRowElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLTableRowElement(val js.Value) HTMLTableRowElement {
	return HTMLTableRowElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLTableRowElement() HTMLTableRowElement { return HTMLTableRowElement{Value: v} }
func (h HTMLTableRowElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLTableRowElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLTableRowElement) GetBgColor() string {
	val := h.Get("bgColor")
	return val.String()
}
func (h HTMLTableRowElement) SetBgColor(val string) {
	h.Set("bgColor", val)
}
func (h HTMLTableRowElement) GetCells() HTMLCollection {
	val := h.Get("cells")
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableRowElement) GetCh() string {
	val := h.Get("ch")
	return val.String()
}
func (h HTMLTableRowElement) SetCh(val string) {
	h.Set("ch", val)
}
func (h HTMLTableRowElement) GetChOff() string {
	val := h.Get("chOff")
	return val.String()
}
func (h HTMLTableRowElement) SetChOff(val string) {
	h.Set("chOff", val)
}
func (h HTMLTableRowElement) DeleteCell(args ...interface{}) {
	h.Call("deleteCell", args...)
}
func (h HTMLTableRowElement) InsertCell(args ...interface{}) HTMLTableCellElement {
	val := h.Call("insertCell", args...)
	return JSValueToHTMLTableCellElement(val.JSValue())
}
func (h HTMLTableRowElement) GetRowIndex() int {
	val := h.Get("rowIndex")
	return val.Int()
}
func (h HTMLTableRowElement) GetSectionRowIndex() int {
	val := h.Get("sectionRowIndex")
	return val.Int()
}
func (h HTMLTableRowElement) GetVAlign() string {
	val := h.Get("vAlign")
	return val.String()
}
func (h HTMLTableRowElement) SetVAlign(val string) {
	h.Set("vAlign", val)
}
