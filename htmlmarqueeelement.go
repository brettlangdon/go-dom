// Code generated DO NOT EDIT
// htmlmarqueeelement.go
package dom

import "syscall/js"

type HTMLMarqueeElementIFace interface {
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
	GetBehavior() string
	SetBehavior(string)
	GetBgColor() string
	SetBgColor(string)
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
	GetDirection() string
	SetDirection(string)
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
	GetHeight() string
	SetHeight(string)
	GetHidden() bool
	SetHidden(bool)
	GetHspace() int
	SetHspace(int)
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
	GetLoop() int
	SetLoop(int)
	Matches(args ...interface{}) bool
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOnbounce() EventHandler
	SetOnbounce(EventHandler)
	GetOnfinish() EventHandler
	SetOnfinish(EventHandler)
	GetOnstart() EventHandler
	SetOnstart(EventHandler)
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
	GetScrollAmount() int
	SetScrollAmount(int)
	GetScrollDelay() int
	SetScrollDelay(int)
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	Start(args ...interface{})
	Stop(args ...interface{})
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetTrueSpeed() bool
	SetTrueSpeed(bool)
	GetVspace() int
	SetVspace(int)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() string
	SetWidth(string)
}
type HTMLMarqueeElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLMarqueeElement(val js.Value) HTMLMarqueeElement {
	return HTMLMarqueeElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLMarqueeElement() HTMLMarqueeElement { return HTMLMarqueeElement{Value: v} }
func (h HTMLMarqueeElement) GetBehavior() string {
	val := h.Get("behavior")
	return val.String()
}
func (h HTMLMarqueeElement) SetBehavior(val string) {
	h.Set("behavior", val)
}
func (h HTMLMarqueeElement) GetBgColor() string {
	val := h.Get("bgColor")
	return val.String()
}
func (h HTMLMarqueeElement) SetBgColor(val string) {
	h.Set("bgColor", val)
}
func (h HTMLMarqueeElement) GetDirection() string {
	val := h.Get("direction")
	return val.String()
}
func (h HTMLMarqueeElement) SetDirection(val string) {
	h.Set("direction", val)
}
func (h HTMLMarqueeElement) GetHeight() string {
	val := h.Get("height")
	return val.String()
}
func (h HTMLMarqueeElement) SetHeight(val string) {
	h.Set("height", val)
}
func (h HTMLMarqueeElement) GetHspace() int {
	val := h.Get("hspace")
	return val.Int()
}
func (h HTMLMarqueeElement) SetHspace(val int) {
	h.Set("hspace", val)
}
func (h HTMLMarqueeElement) GetLoop() int {
	val := h.Get("loop")
	return val.Int()
}
func (h HTMLMarqueeElement) SetLoop(val int) {
	h.Set("loop", val)
}
func (h HTMLMarqueeElement) GetOnbounce() EventHandler {
	val := h.Get("onbounce")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLMarqueeElement) SetOnbounce(val EventHandler) {
	h.Set("onbounce", val)
}
func (h HTMLMarqueeElement) GetOnfinish() EventHandler {
	val := h.Get("onfinish")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLMarqueeElement) SetOnfinish(val EventHandler) {
	h.Set("onfinish", val)
}
func (h HTMLMarqueeElement) GetOnstart() EventHandler {
	val := h.Get("onstart")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLMarqueeElement) SetOnstart(val EventHandler) {
	h.Set("onstart", val)
}
func (h HTMLMarqueeElement) GetScrollAmount() int {
	val := h.Get("scrollAmount")
	return val.Int()
}
func (h HTMLMarqueeElement) SetScrollAmount(val int) {
	h.Set("scrollAmount", val)
}
func (h HTMLMarqueeElement) GetScrollDelay() int {
	val := h.Get("scrollDelay")
	return val.Int()
}
func (h HTMLMarqueeElement) SetScrollDelay(val int) {
	h.Set("scrollDelay", val)
}
func (h HTMLMarqueeElement) Start(args ...interface{}) {
	h.Call("start", args...)
}
func (h HTMLMarqueeElement) Stop(args ...interface{}) {
	h.Call("stop", args...)
}
func (h HTMLMarqueeElement) GetTrueSpeed() bool {
	val := h.Get("trueSpeed")
	return val.Bool()
}
func (h HTMLMarqueeElement) SetTrueSpeed(val bool) {
	h.Set("trueSpeed", val)
}
func (h HTMLMarqueeElement) GetVspace() int {
	val := h.Get("vspace")
	return val.Int()
}
func (h HTMLMarqueeElement) SetVspace(val int) {
	h.Set("vspace", val)
}
func (h HTMLMarqueeElement) GetWidth() string {
	val := h.Get("width")
	return val.String()
}
func (h HTMLMarqueeElement) SetWidth(val string) {
	h.Set("width", val)
}
