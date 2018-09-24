// Code generated DO NOT EDIT
// htmlcollection.go
package dom

import "syscall/js"

type HTMLCollectionIFace interface {
	Item(args ...interface{}) Element
	GetLength() int
	NamedItem(args ...interface{}) Element
}
type HTMLCollection struct {
	Value
}

func JSValueToHTMLCollection(val js.Value) HTMLCollection {
	return HTMLCollection{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLCollection() HTMLCollection { return HTMLCollection{Value: v} }
func NewHTMLCollection(args ...interface{}) HTMLCollection {
	return HTMLCollection{Value: JSValueToValue(js.Global().Get("HTMLCollection").New(args...))}
}
func (h HTMLCollection) Item(args ...interface{}) Element {
	val := h.Call("item", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLCollection) GetLength() int {
	val := h.Get("length")
	return val.Int()
}
func (h HTMLCollection) NamedItem(args ...interface{}) Element {
	val := h.Call("namedItem", args...)
	return JSValueToElement(val.JSValue())
}
