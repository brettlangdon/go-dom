// Code generated DO NOT EDIT
// htmltableelement.go
package dom

import "syscall/js"

type HTMLTableElementIFace interface {
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
	GetCaption() HTMLTableCaptionElement
	SetCaption(HTMLTableCaptionElement)
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
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
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
}
type HTMLTableElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func jsValueToHTMLTableElement(val js.Value) HTMLTableElement {
	return HTMLTableElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLTableElement() HTMLTableElement { return HTMLTableElement{Value: v} }
func (h HTMLTableElement) GetCaption() HTMLTableCaptionElement {
	val := h.Get("caption")
	return jsValueToHTMLTableCaptionElement(val.JSValue())
}
func (h HTMLTableElement) SetCaption(val HTMLTableCaptionElement) {
	h.Set("caption", val)
}
func (h HTMLTableElement) CreateCaption(args ...interface{}) HTMLTableCaptionElement {
	val := h.Call("createCaption", args...)
	return jsValueToHTMLTableCaptionElement(val.JSValue())
}
func (h HTMLTableElement) CreateTBody(args ...interface{}) HTMLTableSectionElement {
	val := h.Call("createTBody", args...)
	return jsValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) CreateTFoot(args ...interface{}) HTMLTableSectionElement {
	val := h.Call("createTFoot", args...)
	return jsValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) CreateTHead(args ...interface{}) HTMLTableSectionElement {
	val := h.Call("createTHead", args...)
	return jsValueToHTMLTableSectionElement(val.JSValue())
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
func (h HTMLTableElement) InsertRow(args ...interface{}) HTMLTableRowElement {
	val := h.Call("insertRow", args...)
	return jsValueToHTMLTableRowElement(val.JSValue())
}
func (h HTMLTableElement) GetRows() HTMLCollection {
	val := h.Get("rows")
	return jsValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableElement) GetTBodies() HTMLCollection {
	val := h.Get("tBodies")
	return jsValueToHTMLCollection(val.JSValue())
}
func (h HTMLTableElement) GetTFoot() HTMLTableSectionElement {
	val := h.Get("tFoot")
	return jsValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) SetTFoot(val HTMLTableSectionElement) {
	h.Set("tFoot", val)
}
func (h HTMLTableElement) GetTHead() HTMLTableSectionElement {
	val := h.Get("tHead")
	return jsValueToHTMLTableSectionElement(val.JSValue())
}
func (h HTMLTableElement) SetTHead(val HTMLTableSectionElement) {
	h.Set("tHead", val)
}
