// Code generated DO NOT EDIT
// domrectlist.go
package dom

import "syscall/js"

type DOMRectListIFace interface {
	Item(args ...interface{}) DOMRect
	GetLength() int
}
type DOMRectList struct {
	Value
}

func JSValueToDOMRectList(val js.Value) DOMRectList { return DOMRectList{Value: JSValueToValue(val)} }
func (v Value) AsDOMRectList() DOMRectList          { return DOMRectList{Value: v} }
func NewDOMRectList(args ...interface{}) DOMRectList {
	return DOMRectList{Value: JSValueToValue(js.Global().Get("DOMRectList").New(args...))}
}
func (d DOMRectList) Item(args ...interface{}) DOMRect {
	val := d.Call("item", args...)
	return JSValueToDOMRect(val.JSValue())
}
func (d DOMRectList) GetLength() int {
	val := d.Get("length")
	return val.Int()
}
