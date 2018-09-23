// Code generated DO NOT EDIT
// htmlcanvaselement.go
package dom

import "syscall/js"

type HTMLCanvasElementIFace interface {
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
	GetContext(args ...interface{}) RenderingContext
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
	ToBlob(args ...interface{})
	ToDataURL(args ...interface{}) string
	ToggleAttribute(args ...interface{}) bool
	TransferControlToOffscreen(args ...interface{}) OffscreenCanvas
	GetTranslate() bool
	SetTranslate(bool)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() int
	SetWidth(int)
}
type HTMLCanvasElement struct {
	Value
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLCanvasElement(val js.Value) HTMLCanvasElement {
	return HTMLCanvasElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLCanvasElement() HTMLCanvasElement { return HTMLCanvasElement{Value: v} }
func (h HTMLCanvasElement) GetContext(args ...interface{}) RenderingContext {
	val := h.Call("getContext", args...)
	return JSValueToRenderingContext(val.JSValue())
}
func (h HTMLCanvasElement) GetHeight() int {
	val := h.Get("height")
	return val.Int()
}
func (h HTMLCanvasElement) SetHeight(val int) {
	h.Set("height", val)
}
func (h HTMLCanvasElement) ToBlob(args ...interface{}) {
	h.Call("toBlob", args...)
}
func (h HTMLCanvasElement) ToDataURL(args ...interface{}) string {
	val := h.Call("toDataURL", args...)
	return val.String()
}
func (h HTMLCanvasElement) TransferControlToOffscreen(args ...interface{}) OffscreenCanvas {
	val := h.Call("transferControlToOffscreen", args...)
	return JSValueToOffscreenCanvas(val.JSValue())
}
func (h HTMLCanvasElement) GetWidth() int {
	val := h.Get("width")
	return val.Int()
}
func (h HTMLCanvasElement) SetWidth(val int) {
	h.Set("width", val)
}
