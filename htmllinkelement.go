// Code generated DO NOT EDIT
// htmllinkelement.go
package dom

import "syscall/js"

type HTMLLinkElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetAs() string
	SetAs(string)
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
	GetCrossOrigin() string
	SetCrossOrigin(string)
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
	GetHref() string
	SetHref(string)
	GetHreflang() string
	SetHreflang(string)
	GetId() string
	SetId(string)
	GetInnerText() string
	SetInnerText(string)
	InsertAdjacentElement(args ...interface{}) Element
	InsertAdjacentText(args ...interface{})
	InsertBefore(args ...interface{}) Node
	GetIntegrity() string
	SetIntegrity(string)
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
	GetMedia() string
	SetMedia(string)
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
	GetReferrerPolicy() string
	SetReferrerPolicy(string)
	GetRel() string
	SetRel(string)
	GetRelList() DOMTokenList
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
	GetSheet() StyleSheet
	GetSizes() DOMTokenList
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
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLLinkElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLLinkElement(val js.Value) HTMLLinkElement {
	return HTMLLinkElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLLinkElement() HTMLLinkElement { return HTMLLinkElement{Value: v} }
func (h HTMLLinkElement) GetAs() string {
	val := h.Get("as")
	return val.String()
}
func (h HTMLLinkElement) SetAs(val string) {
	h.Set("as", val)
}
func (h HTMLLinkElement) GetCrossOrigin() string {
	val := h.Get("crossOrigin")
	return val.String()
}
func (h HTMLLinkElement) SetCrossOrigin(val string) {
	h.Set("crossOrigin", val)
}
func (h HTMLLinkElement) GetHref() string {
	val := h.Get("href")
	return val.String()
}
func (h HTMLLinkElement) SetHref(val string) {
	h.Set("href", val)
}
func (h HTMLLinkElement) GetHreflang() string {
	val := h.Get("hreflang")
	return val.String()
}
func (h HTMLLinkElement) SetHreflang(val string) {
	h.Set("hreflang", val)
}
func (h HTMLLinkElement) GetIntegrity() string {
	val := h.Get("integrity")
	return val.String()
}
func (h HTMLLinkElement) SetIntegrity(val string) {
	h.Set("integrity", val)
}
func (h HTMLLinkElement) GetMedia() string {
	val := h.Get("media")
	return val.String()
}
func (h HTMLLinkElement) SetMedia(val string) {
	h.Set("media", val)
}
func (h HTMLLinkElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLLinkElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLLinkElement) GetRel() string {
	val := h.Get("rel")
	return val.String()
}
func (h HTMLLinkElement) SetRel(val string) {
	h.Set("rel", val)
}
func (h HTMLLinkElement) GetRelList() DOMTokenList {
	val := h.Get("relList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLLinkElement) GetSheet() StyleSheet {
	val := h.Get("sheet")
	return JSValueToStyleSheet(val.JSValue())
}
func (h HTMLLinkElement) GetSizes() DOMTokenList {
	val := h.Get("sizes")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLLinkElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLLinkElement) SetType(val string) {
	h.Set("type", val)
}
