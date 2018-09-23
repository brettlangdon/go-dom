// Code generated DO NOT EDIT
// htmloptionscollection.go
package dom

import "syscall/js"

type HTMLOptionsCollectionIFace interface {
	Add(args ...interface{})
	Item(args ...interface{}) Element
	GetLength() float64
	SetLength(float64)
	NamedItem(args ...interface{}) Element
	Remove(args ...interface{})
	GetSelectedIndex() float64
	SetSelectedIndex(float64)
}
type HTMLOptionsCollection struct {
	Value
	HTMLCollection
}

func jsValueToHTMLOptionsCollection(val js.Value) HTMLOptionsCollection {
	return HTMLOptionsCollection{Value: Value{Value: val}}
}
func (v Value) AsHTMLOptionsCollection() HTMLOptionsCollection { return HTMLOptionsCollection{Value: v} }
func (h HTMLOptionsCollection) Add(args ...interface{}) {
	h.Call("add", args...)
}
func (h HTMLOptionsCollection) GetLength() float64 {
	val := h.Get("length")
	return val.Float()
}
func (h HTMLOptionsCollection) SetLength(val float64) {
	h.Set("length", val)
}
func (h HTMLOptionsCollection) Remove(args ...interface{}) {
	h.Call("remove", args...)
}
func (h HTMLOptionsCollection) GetSelectedIndex() float64 {
	val := h.Get("selectedIndex")
	return val.Float()
}
func (h HTMLOptionsCollection) SetSelectedIndex(val float64) {
	h.Set("selectedIndex", val)
}
