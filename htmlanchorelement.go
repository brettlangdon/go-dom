// Code generated DO NOT EDIT
// htmlanchorelement.go
package dom

import "syscall/js"

type HTMLAnchorElementIFace interface {
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
	GetHreflang() string
	SetHreflang(string)
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
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetTagName() string
	GetTarget() string
	SetTarget(string)
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
	GetUsername() string
	SetUsername(string)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLAnchorElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func jsValueToHTMLAnchorElement(val js.Value) HTMLAnchorElement {
	return HTMLAnchorElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLAnchorElement() HTMLAnchorElement { return HTMLAnchorElement{Value: v} }
func (h HTMLAnchorElement) GetDownload() string {
	val := h.Get("download")
	return val.String()
}
func (h HTMLAnchorElement) SetDownload(val string) {
	h.Set("download", val)
}
func (h HTMLAnchorElement) GetHash() string {
	val := h.Get("hash")
	return val.String()
}
func (h HTMLAnchorElement) SetHash(val string) {
	h.Set("hash", val)
}
func (h HTMLAnchorElement) GetHost() string {
	val := h.Get("host")
	return val.String()
}
func (h HTMLAnchorElement) SetHost(val string) {
	h.Set("host", val)
}
func (h HTMLAnchorElement) GetHostname() string {
	val := h.Get("hostname")
	return val.String()
}
func (h HTMLAnchorElement) SetHostname(val string) {
	h.Set("hostname", val)
}
func (h HTMLAnchorElement) GetHref() string {
	val := h.Get("href")
	return val.String()
}
func (h HTMLAnchorElement) SetHref(val string) {
	h.Set("href", val)
}
func (h HTMLAnchorElement) GetHreflang() string {
	val := h.Get("hreflang")
	return val.String()
}
func (h HTMLAnchorElement) SetHreflang(val string) {
	h.Set("hreflang", val)
}
func (h HTMLAnchorElement) GetOrigin() string {
	val := h.Get("origin")
	return val.String()
}
func (h HTMLAnchorElement) GetPassword() string {
	val := h.Get("password")
	return val.String()
}
func (h HTMLAnchorElement) SetPassword(val string) {
	h.Set("password", val)
}
func (h HTMLAnchorElement) GetPathname() string {
	val := h.Get("pathname")
	return val.String()
}
func (h HTMLAnchorElement) SetPathname(val string) {
	h.Set("pathname", val)
}
func (h HTMLAnchorElement) GetPing() string {
	val := h.Get("ping")
	return val.String()
}
func (h HTMLAnchorElement) SetPing(val string) {
	h.Set("ping", val)
}
func (h HTMLAnchorElement) GetPort() string {
	val := h.Get("port")
	return val.String()
}
func (h HTMLAnchorElement) SetPort(val string) {
	h.Set("port", val)
}
func (h HTMLAnchorElement) GetProtocol() string {
	val := h.Get("protocol")
	return val.String()
}
func (h HTMLAnchorElement) SetProtocol(val string) {
	h.Set("protocol", val)
}
func (h HTMLAnchorElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLAnchorElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLAnchorElement) GetRel() string {
	val := h.Get("rel")
	return val.String()
}
func (h HTMLAnchorElement) SetRel(val string) {
	h.Set("rel", val)
}
func (h HTMLAnchorElement) GetRelList() DOMTokenList {
	val := h.Get("relList")
	return jsValueToDOMTokenList(val.JSValue())
}
func (h HTMLAnchorElement) GetSearch() string {
	val := h.Get("search")
	return val.String()
}
func (h HTMLAnchorElement) SetSearch(val string) {
	h.Set("search", val)
}
func (h HTMLAnchorElement) GetTarget() string {
	val := h.Get("target")
	return val.String()
}
func (h HTMLAnchorElement) SetTarget(val string) {
	h.Set("target", val)
}
func (h HTMLAnchorElement) GetText() string {
	val := h.Get("text")
	return val.String()
}
func (h HTMLAnchorElement) SetText(val string) {
	h.Set("text", val)
}
func (h HTMLAnchorElement) GetType() string {
	val := h.Get("type")
	return val.String()
}
func (h HTMLAnchorElement) SetType(val string) {
	h.Set("type", val)
}
func (h HTMLAnchorElement) GetUsername() string {
	val := h.Get("username")
	return val.String()
}
func (h HTMLAnchorElement) SetUsername(val string) {
	h.Set("username", val)
}
