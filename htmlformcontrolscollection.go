// Code generated DO NOT EDIT
// htmlformcontrolscollection.go
package dom

import "syscall/js"

type HTMLFormControlsCollectionIFace interface {
	Item(args ...interface{}) Element
	GetLength() int
	NamedItem(args ...interface{})
}
type HTMLFormControlsCollection struct {
	Value
	HTMLCollection
}

func JSValueToHTMLFormControlsCollection(val js.Value) HTMLFormControlsCollection {
	return HTMLFormControlsCollection{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLFormControlsCollection() HTMLFormControlsCollection {
	return HTMLFormControlsCollection{Value: v}
}
func NewHTMLFormControlsCollection(args ...interface{}) HTMLFormControlsCollection {
	return HTMLFormControlsCollection{Value: JSValueToValue(js.Global().Get("HTMLFormControlsCollection").New(args...))}
}
func (h HTMLFormControlsCollection) NamedItem(args ...interface{}) {
	h.Call("namedItem", args...)
}
