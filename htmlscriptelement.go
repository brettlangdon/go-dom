// Code generated DO NOT EDIT
// htmlscriptelement.go
package dom

import "syscall/js"

type HTMLScriptElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetAsync() bool
	SetAsync(bool)
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
	GetDefer() bool
	SetDefer(bool)
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
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNoModule() bool
	SetNoModule(bool)
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
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetSrc() string
	SetSrc(string)
	GetTagName() string
	GetText() string
	SetText(string)
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
type HTMLScriptElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func jsValueToHTMLScriptElement(val js.Value) HTMLScriptElement {
	return HTMLScriptElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLScriptElement() HTMLScriptElement { return HTMLScriptElement{Value: v} }
func (h HTMLScriptElement) GetAsync() bool {
	val := h.Get("async")
	return val.Bool()
}
func (h HTMLScriptElement) SetAsync(val bool) {
	h.Set("async", val)
}
func (h HTMLScriptElement) GetCrossOrigin() string {
	val := h.Get("crossOrigin")
	return val.String()
}
func (h HTMLScriptElement) SetCrossOrigin(val string) {
	h.Set("crossOrigin", val)
}
func (h HTMLScriptElement) GetDefer() bool {
	val := h.Get("defer")
	return val.Bool()
}
func (h HTMLScriptElement) SetDefer(val bool) {
	h.Set("defer", val)
}
func (h HTMLScriptElement) GetIntegrity() string {
	val := h.Get("integrity")
	return val.String()
}
func (h HTMLScriptElement) SetIntegrity(val string) {
	h.Set("integrity", val)
}
func (h HTMLScriptElement) GetNoModule() bool {
	val := h.Get("noModule")
	return val.Bool()
}
func (h HTMLScriptElement) SetNoModule(val bool) {
	h.Set("noModule", val)
}
func (h HTMLScriptElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLScriptElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLScriptElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLScriptElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLScriptElement) GetText() string {
	val := h.Get("text")
	return val.String()
}
func (h HTMLScriptElement) SetText(val string) {
	h.Set("text", val)
}
func (h HTMLScriptElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLScriptElement) SetType(val string) {
	h.Set("type", val)
}
