// Code generated DO NOT EDIT
// htmlimageelement.go
package dom

import "syscall/js"

type HTMLImageElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	GetAlt() string
	SetAlt(string)
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
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
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
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() int
	SetWidth(int)
}
type HTMLImageElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLImageElement(val js.Value) HTMLImageElement {
	return HTMLImageElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLImageElement() HTMLImageElement { return HTMLImageElement{Value: v} }
func (h HTMLImageElement) GetAlt() string {
	val := h.Get("alt")
	return val.String()
}
func (h HTMLImageElement) SetAlt(val string) {
	h.Set("alt", val)
}
func (h HTMLImageElement) GetComplete() bool {
	val := h.Get("complete")
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
func (h HTMLImageElement) GetHeight() int {
	val := h.Get("height")
	return val.Int()
}
func (h HTMLImageElement) SetHeight(val int) {
	h.Set("height", val)
}
func (h HTMLImageElement) GetIsMap() bool {
	val := h.Get("isMap")
	return val.Bool()
}
func (h HTMLImageElement) SetIsMap(val bool) {
	h.Set("isMap", val)
}
func (h HTMLImageElement) GetNaturalHeight() int {
	val := h.Get("naturalHeight")
	return val.Int()
}
func (h HTMLImageElement) GetNaturalWidth() int {
	val := h.Get("naturalWidth")
	return val.Int()
}
func (h HTMLImageElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLImageElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLImageElement) GetSizes() string {
	val := h.Get("sizes")
	return val.String()
}
func (h HTMLImageElement) SetSizes(val string) {
	h.Set("sizes", val)
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
func (h HTMLImageElement) GetUseMap() string {
	val := h.Get("useMap")
	return val.String()
}
func (h HTMLImageElement) SetUseMap(val string) {
	h.Set("useMap", val)
}
func (h HTMLImageElement) GetWidth() int {
	val := h.Get("width")
	return val.Int()
}
func (h HTMLImageElement) SetWidth(val int) {
	h.Set("width", val)
}
