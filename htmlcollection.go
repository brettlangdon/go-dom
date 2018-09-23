// Code generated DO NOT EDIT
// htmlcollection.go
package dom

import "syscall/js"

type HTMLCollectionIFace interface {
	Item(args ...interface{}) Element
	GetLength() float64
	NamedItem(args ...interface{}) Element
}
type HTMLCollection struct {
	Value
}

func jsValueToHTMLCollection(val js.Value) HTMLCollection {
	return HTMLCollection{Value: Value{Value: val}}
}
func (v Value) AsHTMLCollection() HTMLCollection { return HTMLCollection{Value: v} }
func (h HTMLCollection) Item(args ...interface{}) Element {
	val := h.Call("item", args...)
	return jsValueToElement(val.JSValue())
}
func (h HTMLCollection) GetLength() float64 {
	val := h.Get("length")
	return val.Float()
}
func (h HTMLCollection) NamedItem(args ...interface{}) Element {
	val := h.Call("namedItem", args...)
	return jsValueToElement(val.JSValue())
}
