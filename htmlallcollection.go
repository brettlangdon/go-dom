// Code generated DO NOT EDIT
// htmlallcollection.go
package dom

import "syscall/js"

type HTMLAllCollectionIFace interface {
	Item(args ...interface{})
	GetLength() int
	NamedItem(args ...interface{})
}
type HTMLAllCollection struct {
	Value
}

func JSValueToHTMLAllCollection(val js.Value) HTMLAllCollection {
	return HTMLAllCollection{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLAllCollection() HTMLAllCollection { return HTMLAllCollection{Value: v} }
func NewHTMLAllCollection(args ...interface{}) HTMLAllCollection {
	return HTMLAllCollection{Value: JSValueToValue(js.Global().Get("HTMLAllCollection").New(args...))}
}
func (h HTMLAllCollection) Item(args ...interface{}) {
	h.Call("item", args...)
}
func (h HTMLAllCollection) GetLength() int {
	val := h.Get("length")
	return val.Int()
}
func (h HTMLAllCollection) NamedItem(args ...interface{}) {
	h.Call("namedItem", args...)
}
