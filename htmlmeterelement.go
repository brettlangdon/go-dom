// Code generated DO NOT EDIT
// htmlmeterelement.go
package dom

import "syscall/js"

type HTMLMeterElementIFace interface {
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
	GetHigh() float64
	SetHigh(float64)
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
	GetLabels() NodeList
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetLow() float64
	SetLow(float64)
	Matches(args ...interface{}) bool
	GetMax() float64
	SetMax(float64)
	GetMin() float64
	SetMin(float64)
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOptimum() float64
	SetOptimum(float64)
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
	GetValue() float64
	SetValue(float64)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLMeterElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLMeterElement(val js.Value) HTMLMeterElement {
	return HTMLMeterElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLMeterElement() HTMLMeterElement { return HTMLMeterElement{Value: v} }
func (h HTMLMeterElement) GetHigh() float64 {
	val := h.Get("high")
	return val.Float()
}
func (h HTMLMeterElement) SetHigh(val float64) {
	h.Set("high", val)
}
func (h HTMLMeterElement) GetLabels() NodeList {
	val := h.Get("labels")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLMeterElement) GetLow() float64 {
	val := h.Get("low")
	return val.Float()
}
func (h HTMLMeterElement) SetLow(val float64) {
	h.Set("low", val)
}
func (h HTMLMeterElement) GetMax() float64 {
	val := h.Get("max")
	return val.Float()
}
func (h HTMLMeterElement) SetMax(val float64) {
	h.Set("max", val)
}
func (h HTMLMeterElement) GetMin() float64 {
	val := h.Get("min")
	return val.Float()
}
func (h HTMLMeterElement) SetMin(val float64) {
	h.Set("min", val)
}
func (h HTMLMeterElement) GetOptimum() float64 {
	val := h.Get("optimum")
	return val.Float()
}
func (h HTMLMeterElement) SetOptimum(val float64) {
	h.Set("optimum", val)
}
func (h HTMLMeterElement) GetValue() float64 {
	val := h.Get("value")
	return val.Float()
}
func (h HTMLMeterElement) SetValue(val float64) {
	h.Set("value", val)
}
