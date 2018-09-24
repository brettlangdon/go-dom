// Code generated DO NOT EDIT
// htmliframeelement.go
package dom

import "syscall/js"

type HTMLIFrameElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	GetAlign() string
	SetAlign(string)
	GetAllow() string
	SetAllow(string)
	GetAllowFullscreen() bool
	SetAllowFullscreen(bool)
	GetAllowPaymentRequest() bool
	SetAllowPaymentRequest(bool)
	GetAllowUserMedia() bool
	SetAllowUserMedia(bool)
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
	GetSVGDocument(args ...interface{}) Document
	HasAttribute(args ...interface{}) bool
	HasAttributeNS(args ...interface{}) bool
	HasAttributes(args ...interface{}) bool
	HasChildNodes(args ...interface{}) bool
	GetHeight() string
	SetHeight(string)
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
	GetSandbox() DOMTokenList
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
	GetSrcdoc() string
	SetSrcdoc(string)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() string
	SetWidth(string)
}
type HTMLIFrameElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLIFrameElement(val js.Value) HTMLIFrameElement {
	return HTMLIFrameElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLIFrameElement() HTMLIFrameElement { return HTMLIFrameElement{Value: v} }
func (h HTMLIFrameElement) GetAlign() string {
	val := h.Get("align")
	return val.String()
}
func (h HTMLIFrameElement) SetAlign(val string) {
	h.Set("align", val)
}
func (h HTMLIFrameElement) GetAllow() string {
	val := h.Get("allow")
	return val.String()
}
func (h HTMLIFrameElement) SetAllow(val string) {
	h.Set("allow", val)
}
func (h HTMLIFrameElement) GetAllowFullscreen() bool {
	val := h.Get("allowFullscreen")
	return val.Bool()
}
func (h HTMLIFrameElement) SetAllowFullscreen(val bool) {
	h.Set("allowFullscreen", val)
}
func (h HTMLIFrameElement) GetAllowPaymentRequest() bool {
	val := h.Get("allowPaymentRequest")
	return val.Bool()
}
func (h HTMLIFrameElement) SetAllowPaymentRequest(val bool) {
	h.Set("allowPaymentRequest", val)
}
func (h HTMLIFrameElement) GetAllowUserMedia() bool {
	val := h.Get("allowUserMedia")
	return val.Bool()
}
func (h HTMLIFrameElement) SetAllowUserMedia(val bool) {
	h.Set("allowUserMedia", val)
}
func (h HTMLIFrameElement) GetContentDocument() Document {
	val := h.Get("contentDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLIFrameElement) GetContentWindow() WindowProxy {
	val := h.Get("contentWindow")
	return JSValueToWindowProxy(val.JSValue())
}
func (h HTMLIFrameElement) GetFrameBorder() string {
	val := h.Get("frameBorder")
	return val.String()
}
func (h HTMLIFrameElement) SetFrameBorder(val string) {
	h.Set("frameBorder", val)
}
func (h HTMLIFrameElement) GetSVGDocument(args ...interface{}) Document {
	val := h.Call("getSVGDocument", args...)
	return JSValueToDocument(val.JSValue())
}
func (h HTMLIFrameElement) GetHeight() string {
	val := h.Get("height")
	return val.String()
}
func (h HTMLIFrameElement) SetHeight(val string) {
	h.Set("height", val)
}
func (h HTMLIFrameElement) GetLongDesc() string {
	val := h.Get("longDesc")
	return val.String()
}
func (h HTMLIFrameElement) SetLongDesc(val string) {
	h.Set("longDesc", val)
}
func (h HTMLIFrameElement) GetMarginHeight() string {
	val := h.Get("marginHeight")
	return val.String()
}
func (h HTMLIFrameElement) SetMarginHeight(val string) {
	h.Set("marginHeight", val)
}
func (h HTMLIFrameElement) GetMarginWidth() string {
	val := h.Get("marginWidth")
	return val.String()
}
func (h HTMLIFrameElement) SetMarginWidth(val string) {
	h.Set("marginWidth", val)
}
func (h HTMLIFrameElement) GetName() string {
	val := h.Get("name")
	return val.String()
}
func (h HTMLIFrameElement) SetName(val string) {
	h.Set("name", val)
}
func (h HTMLIFrameElement) GetReferrerPolicy() string {
	val := h.Get("referrerPolicy")
	return val.String()
}
func (h HTMLIFrameElement) SetReferrerPolicy(val string) {
	h.Set("referrerPolicy", val)
}
func (h HTMLIFrameElement) GetSandbox() DOMTokenList {
	val := h.Get("sandbox")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLIFrameElement) GetScrolling() string {
	val := h.Get("scrolling")
	return val.String()
}
func (h HTMLIFrameElement) SetScrolling(val string) {
	h.Set("scrolling", val)
}
func (h HTMLIFrameElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLIFrameElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLIFrameElement) GetSrcdoc() string {
	val := h.Get("srcdoc")
	return val.String()
}
func (h HTMLIFrameElement) SetSrcdoc(val string) {
	h.Set("srcdoc", val)
}
func (h HTMLIFrameElement) GetWidth() string {
	val := h.Get("width")
	return val.String()
}
func (h HTMLIFrameElement) SetWidth(val string) {
	h.Set("width", val)
}
