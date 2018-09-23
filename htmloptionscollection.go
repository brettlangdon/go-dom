// Code generated DO NOT EDIT
// htmloptionscollection.go
package dom

import "syscall/js"

type HTMLOptionsCollectionIFace interface {
	Add(args ...interface{})
	Item(args ...interface{}) Element
	GetLength() int
	SetLength(int)
	NamedItem(args ...interface{}) Element
	Remove(args ...interface{})
	GetSelectedIndex() int
	SetSelectedIndex(int)
}
type HTMLOptionsCollection struct {
	Value
	HTMLCollection
}

func JSValueToHTMLOptionsCollection(val js.Value) HTMLOptionsCollection {
	return HTMLOptionsCollection{Value: Value{Value: val}}
}
func (v Value) AsHTMLOptionsCollection() HTMLOptionsCollection { return HTMLOptionsCollection{Value: v} }
func (h HTMLOptionsCollection) Add(args ...interface{}) {
	h.Call("add", args...)
}
func (h HTMLOptionsCollection) GetLength() int {
	val := h.Get("length")
	return val.Int()
}
func (h HTMLOptionsCollection) SetLength(val int) {
	h.Set("length", val)
}
func (h HTMLOptionsCollection) Remove(args ...interface{}) {
	h.Call("remove", args...)
}
func (h HTMLOptionsCollection) GetSelectedIndex() int {
	val := h.Get("selectedIndex")
	return val.Int()
}
func (h HTMLOptionsCollection) SetSelectedIndex(val int) {
	h.Set("selectedIndex", val)
}
