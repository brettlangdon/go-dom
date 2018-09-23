// Code generated DO NOT EDIT
// htmlbodyelement.go
package dom

import "syscall/js"

type HTMLBodyElementIFace interface {
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
	GetOnafterprint() EventHandler
	SetOnafterprint(EventHandler)
	GetOnbeforeprint() EventHandler
	SetOnbeforeprint(EventHandler)
	GetOnbeforeunload() OnBeforeUnloadEventHandler
	SetOnbeforeunload(OnBeforeUnloadEventHandler)
	GetOnhashchange() EventHandler
	SetOnhashchange(EventHandler)
	GetOnlanguagechange() EventHandler
	SetOnlanguagechange(EventHandler)
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnmessageerror() EventHandler
	SetOnmessageerror(EventHandler)
	GetOnoffline() EventHandler
	SetOnoffline(EventHandler)
	GetOnonline() EventHandler
	SetOnonline(EventHandler)
	GetOnpagehide() EventHandler
	SetOnpagehide(EventHandler)
	GetOnpageshow() EventHandler
	SetOnpageshow(EventHandler)
	GetOnpopstate() EventHandler
	SetOnpopstate(EventHandler)
	GetOnrejectionhandled() EventHandler
	SetOnrejectionhandled(EventHandler)
	GetOnstorage() EventHandler
	SetOnstorage(EventHandler)
	GetOnunhandledrejection() EventHandler
	SetOnunhandledrejection(EventHandler)
	GetOnunload() EventHandler
	SetOnunload(EventHandler)
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
type HTMLBodyElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func jsValueToHTMLBodyElement(val js.Value) HTMLBodyElement {
	return HTMLBodyElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLBodyElement() HTMLBodyElement { return HTMLBodyElement{Value: v} }
func (h HTMLBodyElement) GetOnafterprint() EventHandler {
	val := h.Get("onafterprint")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnafterprint(val EventHandler) {
	h.Set("onafterprint", val)
}
func (h HTMLBodyElement) GetOnbeforeprint() EventHandler {
	val := h.Get("onbeforeprint")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnbeforeprint(val EventHandler) {
	h.Set("onbeforeprint", val)
}
func (h HTMLBodyElement) GetOnbeforeunload() OnBeforeUnloadEventHandler {
	val := h.Get("onbeforeunload")
	return jsValueToOnBeforeUnloadEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnbeforeunload(val OnBeforeUnloadEventHandler) {
	h.Set("onbeforeunload", val)
}
func (h HTMLBodyElement) GetOnhashchange() EventHandler {
	val := h.Get("onhashchange")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnhashchange(val EventHandler) {
	h.Set("onhashchange", val)
}
func (h HTMLBodyElement) GetOnlanguagechange() EventHandler {
	val := h.Get("onlanguagechange")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnlanguagechange(val EventHandler) {
	h.Set("onlanguagechange", val)
}
func (h HTMLBodyElement) GetOnmessage() EventHandler {
	val := h.Get("onmessage")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnmessage(val EventHandler) {
	h.Set("onmessage", val)
}
func (h HTMLBodyElement) GetOnmessageerror() EventHandler {
	val := h.Get("onmessageerror")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnmessageerror(val EventHandler) {
	h.Set("onmessageerror", val)
}
func (h HTMLBodyElement) GetOnoffline() EventHandler {
	val := h.Get("onoffline")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnoffline(val EventHandler) {
	h.Set("onoffline", val)
}
func (h HTMLBodyElement) GetOnonline() EventHandler {
	val := h.Get("ononline")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnonline(val EventHandler) {
	h.Set("ononline", val)
}
func (h HTMLBodyElement) GetOnpagehide() EventHandler {
	val := h.Get("onpagehide")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnpagehide(val EventHandler) {
	h.Set("onpagehide", val)
}
func (h HTMLBodyElement) GetOnpageshow() EventHandler {
	val := h.Get("onpageshow")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnpageshow(val EventHandler) {
	h.Set("onpageshow", val)
}
func (h HTMLBodyElement) GetOnpopstate() EventHandler {
	val := h.Get("onpopstate")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnpopstate(val EventHandler) {
	h.Set("onpopstate", val)
}
func (h HTMLBodyElement) GetOnrejectionhandled() EventHandler {
	val := h.Get("onrejectionhandled")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnrejectionhandled(val EventHandler) {
	h.Set("onrejectionhandled", val)
}
func (h HTMLBodyElement) GetOnstorage() EventHandler {
	val := h.Get("onstorage")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnstorage(val EventHandler) {
	h.Set("onstorage", val)
}
func (h HTMLBodyElement) GetOnunhandledrejection() EventHandler {
	val := h.Get("onunhandledrejection")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnunhandledrejection(val EventHandler) {
	h.Set("onunhandledrejection", val)
}
func (h HTMLBodyElement) GetOnunload() EventHandler {
	val := h.Get("onunload")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLBodyElement) SetOnunload(val EventHandler) {
	h.Set("onunload", val)
}
