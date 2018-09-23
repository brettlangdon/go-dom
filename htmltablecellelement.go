// Code generated DO NOT EDIT
// htmltablecellelement.go
package dom

import "syscall/js"

type HTMLTableCellElementIFace interface {
	GetAbbr() string
	SetAbbr(string)
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
	GetCellIndex() int
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	GetColSpan() int
	SetColSpan(int)
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
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
	GetHeaders() string
	SetHeaders(string)
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
	GetRowSpan() int
	SetRowSpan(int)
	GetScope() string
	SetScope(string)
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
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLTableCellElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLTableCellElement(val js.Value) HTMLTableCellElement {
	return HTMLTableCellElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLTableCellElement() HTMLTableCellElement { return HTMLTableCellElement{Value: v} }
func (h HTMLTableCellElement) GetAbbr() string {
	val := h.Get("abbr")
	return val.String()
}
func (h HTMLTableCellElement) SetAbbr(val string) {
	h.Set("abbr", val)
}
func (h HTMLTableCellElement) GetCellIndex() int {
	val := h.Get("cellIndex")
	return val.Int()
}
func (h HTMLTableCellElement) GetColSpan() int {
	val := h.Get("colSpan")
	return val.Int()
}
func (h HTMLTableCellElement) SetColSpan(val int) {
	h.Set("colSpan", val)
}
func (h HTMLTableCellElement) GetHeaders() string {
	val := h.Get("headers")
	return val.String()
}
func (h HTMLTableCellElement) SetHeaders(val string) {
	h.Set("headers", val)
}
func (h HTMLTableCellElement) GetRowSpan() int {
	val := h.Get("rowSpan")
	return val.Int()
}
func (h HTMLTableCellElement) SetRowSpan(val int) {
	h.Set("rowSpan", val)
}
func (h HTMLTableCellElement) GetScope() string {
	val := h.Get("scope")
	return val.String()
}
func (h HTMLTableCellElement) SetScope(val string) {
	h.Set("scope", val)
}
