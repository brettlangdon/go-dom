// Code generated DO NOT EDIT
// htmlareaelement.go
package dom

import "syscall/js"

type HTMLAreaElementIFace interface {
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
	Contains(args ...interface{}) bool
	GetCoords() string
	SetCoords(string)
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDownload() string
	SetDownload(string)
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
	GetHash() string
	SetHash(string)
	GetHidden() bool
	SetHidden(bool)
	GetHost() string
	SetHost(string)
	GetHostname() string
	SetHostname(string)
	GetHref() string
	SetHref(string)
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
	GetOrigin() string
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPassword() string
	SetPassword(string)
	GetPathname() string
	SetPathname(string)
	GetPing() string
	SetPing(string)
	GetPort() string
	SetPort(string)
	GetPrefix() string
	GetPreviousSibling() Node
	GetProtocol() string
	SetProtocol(string)
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
	GetSearch() string
	SetSearch(string)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetShape() string
	SetShape(string)
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetTagName() string
	GetTarget() string
	SetTarget(string)
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetUsername() string
	SetUsername(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLAreaElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLAreaElement(val js.Value) HTMLAreaElement {
	return HTMLAreaElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLAreaElement() HTMLAreaElement { return HTMLAreaElement{Value: v} }
func (h HTMLAreaElement) GetAlt() string {
	val := h.Get("alt")
	return val.String()
}
func (h HTMLAreaElement) SetAlt(val string) {
	h.Set("alt", val)
}
func (h HTMLAreaElement) GetCoords() string {
	val := h.Get("coords")
	return val.String()
}
func (h HTMLAreaElement) SetCoords(val string) {
	h.Set("coords", val)
}
func (h HTMLAreaElement) GetDownload() string {
	val := h.Get("download")
	return val.String()
}
func (h HTMLAreaElement) SetDownload(val string) {
	h.Set("download", val)
}
func (h HTMLAreaElement) GetHash() string {
	val := h.Get("hash")
	return val.String()
}
func (h HTMLAreaElement) SetHash(val string) {
	h.Set("hash", val)
}
func (h HTMLAreaElement) GetHost() string {
	val := h.Get("host")
	return val.String()
}
func (h HTMLAreaElement) SetHost(val string) {
	h.Set("host", val)
}
func (h HTMLAreaElement) GetHostname() string {
	val := h.Get("hostname")
	return val.String()
}
func (h HTMLAreaElement) SetHostname(val string) {
	h.Set("hostname", val)
}
func (h HTMLAreaElement) GetHref() string {
	val := h.Get("href")
	return val.String()
}
func (h HTMLAreaElement) SetHref(val string) {
	h.Set("href", val)
}
func (h HTMLAreaElement) GetOrigin() string {
	val := h.Get("origin")
	return val.String()
}
func (h HTMLAreaElement) GetPassword() string {
	val := h.Get("password")
	return val.String()
}
func (h HTMLAreaElement) SetPassword(val string) {
	h.Set("password", val)
}
func (h HTMLAreaElement) GetPathname() string {
	val := h.Get("pathname")
	return val.String()
}
func (h HTMLAreaElement) SetPathname(val string) {
	h.Set("pathname", val)
}
func (h HTMLAreaElement) GetPing() string {
	val := h.Get("ping")
	return val.String()
}
func (h HTMLAreaElement) SetPing(val string) {
	h.Set("ping", val)
}
func (h HTMLAreaElement) GetPort() string {
	val := h.Get("port")
	return val.String()
}
func (h HTMLAreaElement) SetPort(val string) {
	h.Set("port", val)
}
func (h HTMLAreaElement) GetProtocol() string {
	val := h.Get("protocol")
	return val.String()
}
func (h HTMLAreaElement) SetProtocol(val string) {
	h.Set("protocol", val)
}
func (h HTMLAreaElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLAreaElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLAreaElement) GetRel() string {
	val := h.Get("rel")
	return val.String()
}
func (h HTMLAreaElement) SetRel(val string) {
	h.Set("rel", val)
}
func (h HTMLAreaElement) GetRelList() DOMTokenList {
	val := h.Get("relList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLAreaElement) GetSearch() string {
	val := h.Get("search")
	return val.String()
}
func (h HTMLAreaElement) SetSearch(val string) {
	h.Set("search", val)
}
func (h HTMLAreaElement) GetShape() string {
	val := h.Get("shape")
	return val.String()
}
func (h HTMLAreaElement) SetShape(val string) {
	h.Set("shape", val)
}
func (h HTMLAreaElement) GetTarget() string {
	val := h.Get("target")
	return val.String()
}
func (h HTMLAreaElement) SetTarget(val string) {
	h.Set("target", val)
}
func (h HTMLAreaElement) GetUsername() string {
	val := h.Get("username")
	return val.String()
}
func (h HTMLAreaElement) SetUsername(val string) {
	h.Set("username", val)
}
