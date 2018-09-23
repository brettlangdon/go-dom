// Code generated DO NOT EDIT
// htmlframeelement.go
package dom

import "syscall/js"

type HTMLFrameElementIFace interface {
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
	GetContentDocument() Document
	GetContentWindow() WindowProxy
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetFirstChild() Node
	GetFrameBorder() string
	SetFrameBorder(string)
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
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	GetLocalName() string
	GetLongDesc() string
	SetLongDesc(string)
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetMarginHeight() string
	SetMarginHeight(string)
	GetMarginWidth() string
	SetMarginWidth(string)
	Matches(args ...interface{}) bool
	GetName() string
	SetName(string)
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNoResize() bool
	SetNoResize(bool)
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
	GetScrolling() string
	SetScrolling(string)
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
type HTMLFrameElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLFrameElement(val js.Value) HTMLFrameElement {
	return HTMLFrameElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLFrameElement() HTMLFrameElement { return HTMLFrameElement{Value: v} }
func (h HTMLFrameElement) GetContentDocument() Document {
	val := h.Get("contentDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLFrameElement) GetContentWindow() WindowProxy {
	val := h.Get("contentWindow")
	return JSValueToWindowProxy(val.JSValue())
}
func (h HTMLFrameElement) GetFrameBorder() string {
	val := h.Get("frameBorder")
	return val.String()
}
func (h HTMLFrameElement) SetFrameBorder(val string) {
	h.Set("frameBorder", val)
}
func (h HTMLFrameElement) GetLongDesc() string {
	val := h.Get("longDesc")
	return val.String()
}
func (h HTMLFrameElement) SetLongDesc(val string) {
	h.Set("longDesc", val)
}
func (h HTMLFrameElement) GetMarginHeight() string {
	val := h.Get("marginHeight")
	return val.String()
}
func (h HTMLFrameElement) SetMarginHeight(val string) {
	h.Set("marginHeight", val)
}
func (h HTMLFrameElement) GetMarginWidth() string {
	val := h.Get("marginWidth")
	return val.String()
}
func (h HTMLFrameElement) SetMarginWidth(val string) {
	h.Set("marginWidth", val)
}
func (h HTMLFrameElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLFrameElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLFrameElement) GetNoResize() bool {
	val := h.Get("noResize")
	return val.Bool()
}
func (h HTMLFrameElement) SetNoResize(val bool) {
	h.Set("noResize", val)
}
func (h HTMLFrameElement) GetScrolling() string {
	val := h.Get("scrolling")
	return val.String()
}
func (h HTMLFrameElement) SetScrolling(val string) {
	h.Set("scrolling", val)
}
func (h HTMLFrameElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLFrameElement) SetSrc(val string) {
	h.Set("src", val)
}
