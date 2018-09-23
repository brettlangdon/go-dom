// Code generated DO NOT EDIT
// htmlallcollection.go
package dom

import "syscall/js"

type HTMLAllCollectionIFace interface {
	Item(args ...interface{})
	GetLength() float64
	NamedItem(args ...interface{})
}
type HTMLAllCollection struct {
	Value
}

func JSValueToHTMLAllCollection(val js.Value) HTMLAllCollection {
	return HTMLAllCollection{Value: Value{Value: val}}
}
func (v Value) AsHTMLAllCollection() HTMLAllCollection { return HTMLAllCollection{Value: v} }
func (h HTMLAllCollection) Item(args ...interface{}) {
	h.Call("item", args...)
}
func (h HTMLAllCollection) GetLength() float64 {
	val := h.Get("length")
	return val.Float()
}
func (h HTMLAllCollection) NamedItem(args ...interface{}) {
	h.Call("namedItem", args...)
}
