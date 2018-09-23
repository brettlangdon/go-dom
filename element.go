// Code generated DO NOT EDIT
// element.go
package dom

import "syscall/js"

type ElementIFace interface {
	AddEventListener(args ...interface{})
	After(args ...interface{})
	Append(args ...interface{})
	AppendChild(args ...interface{}) Node
	GetAssignedSlot() HTMLSlotElement
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetBaseURI() string
	Before(args ...interface{})
	GetChildElementCount() int
	GetChildNodes() NodeList
	GetChildren() HTMLCollection
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	DispatchEvent(args ...interface{}) bool
	GetFirstChild() Node
	GetFirstElementChild() Element
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
	GetId() string
	SetId(string)
	InsertAdjacentElement(args ...interface{}) Element
	InsertAdjacentText(args ...interface{})
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLastChild() Node
	GetLastElementChild() Element
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
	GetNamespaceURI() string
	GetNextElementSibling() Element
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
	Prepend(args ...interface{})
	GetPreviousElementSibling() Element
	GetPreviousSibling() Node
	QuerySelector(args ...interface{}) Element
	QuerySelectorAll(args ...interface{}) NodeList
	Remove(args ...interface{})
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	ReplaceWith(args ...interface{})
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	ToggleAttribute(args ...interface{}) bool
	WebkitMatchesSelector(args ...interface{}) bool
}
type Element struct {
	Value
	Node
	EventTarget
}

func JSValueToElement(val js.Value) Element { return Element{Value: Value{Value: val}} }
func (v Value) AsElement() Element          { return Element{Value: v} }
func (e Element) After(args ...interface{}) {
	e.Call("after", args...)
}
func (e Element) Append(args ...interface{}) {
	e.Call("append", args...)
}
func (e Element) GetAssignedSlot() HTMLSlotElement {
	val := e.Get("assignedSlot")
	return JSValueToHTMLSlotElement(val.JSValue())
}
func (e Element) AttachShadow(args ...interface{}) ShadowRoot {
	val := e.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (e Element) GetAttributes() NamedNodeMap {
	val := e.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (e Element) Before(args ...interface{}) {
	e.Call("before", args...)
}
func (e Element) GetChildElementCount() int {
	val := e.Get("childElementCount")
	return val.Int()
}
func (e Element) GetChildren() HTMLCollection {
	val := e.Get("children")
	return JSValueToHTMLCollection(val.JSValue())
}
func (e Element) GetClassList() DOMTokenList {
	val := e.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (e Element) GetClassName() string {
	val := e.Get("className")
	return val.String()
}
func (e Element) SetClassName(val string) {
	e.Set("className", val)
}
func (e Element) Closest(args ...interface{}) Element {
	val := e.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (e Element) GetFirstElementChild() Element {
	val := e.Get("firstElementChild")
	return JSValueToElement(val.JSValue())
}
func (e Element) GetAttribute(args ...interface{}) string {
	val := e.Call("getAttribute", args...)
	return val.String()
}
func (e Element) GetAttributeNS(args ...interface{}) string {
	val := e.Call("getAttributeNS", args...)
	return val.String()
}
func (e Element) GetAttributeNames(args ...interface{}) {
	e.Call("getAttributeNames", args...)
}
func (e Element) GetAttributeNode(args ...interface{}) Attr {
	val := e.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (e Element) GetAttributeNodeNS(args ...interface{}) Attr {
	val := e.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (e Element) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := e.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (e Element) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := e.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (e Element) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := e.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (e Element) HasAttribute(args ...interface{}) bool {
	val := e.Call("hasAttribute", args...)
	return val.Bool()
}
func (e Element) HasAttributeNS(args ...interface{}) bool {
	val := e.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (e Element) HasAttributes(args ...interface{}) bool {
	val := e.Call("hasAttributes", args...)
	return val.Bool()
}
func (e Element) GetId() string {
	val := e.Get("id")
	return val.String()
}
func (e Element) SetId(val string) {
	e.Set("id", val)
}
func (e Element) InsertAdjacentElement(args ...interface{}) Element {
	val := e.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (e Element) InsertAdjacentText(args ...interface{}) {
	e.Call("insertAdjacentText", args...)
}
func (e Element) GetLastElementChild() Element {
	val := e.Get("lastElementChild")
	return JSValueToElement(val.JSValue())
}
func (e Element) GetLocalName() string {
	val := e.Get("localName")
	return val.String()
}
func (e Element) Matches(args ...interface{}) bool {
	val := e.Call("matches", args...)
	return val.Bool()
}
func (e Element) GetNamespaceURI() string {
	val := e.Get("namespaceURI")
	return val.String()
}
func (e Element) GetNextElementSibling() Element {
	val := e.Get("nextElementSibling")
	return JSValueToElement(val.JSValue())
}
func (e Element) GetPrefix() string {
	val := e.Get("prefix")
	return val.String()
}
func (e Element) Prepend(args ...interface{}) {
	e.Call("prepend", args...)
}
func (e Element) GetPreviousElementSibling() Element {
	val := e.Get("previousElementSibling")
	return JSValueToElement(val.JSValue())
}
func (e Element) QuerySelector(args ...interface{}) Element {
	val := e.Call("querySelector", args...)
	return JSValueToElement(val.JSValue())
}
func (e Element) QuerySelectorAll(args ...interface{}) NodeList {
	val := e.Call("querySelectorAll", args...)
	return JSValueToNodeList(val.JSValue())
}
func (e Element) Remove(args ...interface{}) {
	e.Call("remove", args...)
}
func (e Element) RemoveAttribute(args ...interface{}) {
	e.Call("removeAttribute", args...)
}
func (e Element) RemoveAttributeNS(args ...interface{}) {
	e.Call("removeAttributeNS", args...)
}
func (e Element) RemoveAttributeNode(args ...interface{}) Attr {
	val := e.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (e Element) ReplaceWith(args ...interface{}) {
	e.Call("replaceWith", args...)
}
func (e Element) SetAttribute(args ...interface{}) {
	e.Call("setAttribute", args...)
}
func (e Element) SetAttributeNS(args ...interface{}) {
	e.Call("setAttributeNS", args...)
}
func (e Element) SetAttributeNode(args ...interface{}) Attr {
	val := e.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (e Element) SetAttributeNodeNS(args ...interface{}) Attr {
	val := e.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (e Element) GetShadowRoot() ShadowRoot {
	val := e.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (e Element) GetSlot() string {
	val := e.Get("slot")
	return val.String()
}
func (e Element) SetSlot(val string) {
	e.Set("slot", val)
}
func (e Element) GetTagName() string {
	val := e.Get("tagName")
	return val.String()
}
func (e Element) ToggleAttribute(args ...interface{}) bool {
	val := e.Call("toggleAttribute", args...)
	return val.Bool()
}
func (e Element) WebkitMatchesSelector(args ...interface{}) bool {
	val := e.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
