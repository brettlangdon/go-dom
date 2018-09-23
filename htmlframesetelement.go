// Code generated DO NOT EDIT
// htmlframesetelement.go
package dom

import "syscall/js"

type HTMLFrameSetElementIFace interface {
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
	GetCols() string
	SetCols(string)
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
	GetRows() string
	SetRows(string)
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
type HTMLFrameSetElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func jsValueToHTMLFrameSetElement(val js.Value) HTMLFrameSetElement {
	return HTMLFrameSetElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLFrameSetElement() HTMLFrameSetElement { return HTMLFrameSetElement{Value: v} }
func (h HTMLFrameSetElement) GetCols() string {
	val := h.Get("cols")
	return val.String()
}
func (h HTMLFrameSetElement) SetCols(val string) {
	h.Set("cols", val)
}
func (h HTMLFrameSetElement) GetOnafterprint() EventHandler {
	val := h.Get("onafterprint")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnafterprint(val EventHandler) {
	h.Set("onafterprint", val)
}
func (h HTMLFrameSetElement) GetOnbeforeprint() EventHandler {
	val := h.Get("onbeforeprint")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnbeforeprint(val EventHandler) {
	h.Set("onbeforeprint", val)
}
func (h HTMLFrameSetElement) GetOnbeforeunload() OnBeforeUnloadEventHandler {
	val := h.Get("onbeforeunload")
	return jsValueToOnBeforeUnloadEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnbeforeunload(val OnBeforeUnloadEventHandler) {
	h.Set("onbeforeunload", val)
}
func (h HTMLFrameSetElement) GetOnhashchange() EventHandler {
	val := h.Get("onhashchange")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnhashchange(val EventHandler) {
	h.Set("onhashchange", val)
}
func (h HTMLFrameSetElement) GetOnlanguagechange() EventHandler {
	val := h.Get("onlanguagechange")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnlanguagechange(val EventHandler) {
	h.Set("onlanguagechange", val)
}
func (h HTMLFrameSetElement) GetOnmessage() EventHandler {
	val := h.Get("onmessage")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnmessage(val EventHandler) {
	h.Set("onmessage", val)
}
func (h HTMLFrameSetElement) GetOnmessageerror() EventHandler {
	val := h.Get("onmessageerror")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnmessageerror(val EventHandler) {
	h.Set("onmessageerror", val)
}
func (h HTMLFrameSetElement) GetOnoffline() EventHandler {
	val := h.Get("onoffline")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnoffline(val EventHandler) {
	h.Set("onoffline", val)
}
func (h HTMLFrameSetElement) GetOnonline() EventHandler {
	val := h.Get("ononline")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnonline(val EventHandler) {
	h.Set("ononline", val)
}
func (h HTMLFrameSetElement) GetOnpagehide() EventHandler {
	val := h.Get("onpagehide")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnpagehide(val EventHandler) {
	h.Set("onpagehide", val)
}
func (h HTMLFrameSetElement) GetOnpageshow() EventHandler {
	val := h.Get("onpageshow")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnpageshow(val EventHandler) {
	h.Set("onpageshow", val)
}
func (h HTMLFrameSetElement) GetOnpopstate() EventHandler {
	val := h.Get("onpopstate")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnpopstate(val EventHandler) {
	h.Set("onpopstate", val)
}
func (h HTMLFrameSetElement) GetOnrejectionhandled() EventHandler {
	val := h.Get("onrejectionhandled")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnrejectionhandled(val EventHandler) {
	h.Set("onrejectionhandled", val)
}
func (h HTMLFrameSetElement) GetOnstorage() EventHandler {
	val := h.Get("onstorage")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnstorage(val EventHandler) {
	h.Set("onstorage", val)
}
func (h HTMLFrameSetElement) GetOnunhandledrejection() EventHandler {
	val := h.Get("onunhandledrejection")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnunhandledrejection(val EventHandler) {
	h.Set("onunhandledrejection", val)
}
func (h HTMLFrameSetElement) GetOnunload() EventHandler {
	val := h.Get("onunload")
	return jsValueToEventHandler(val.JSValue())
}
func (h HTMLFrameSetElement) SetOnunload(val EventHandler) {
	h.Set("onunload", val)
}
func (h HTMLFrameSetElement) GetRows() string {
	val := h.Get("rows")
	return val.String()
}
func (h HTMLFrameSetElement) SetRows(val string) {
	h.Set("rows", val)
}
