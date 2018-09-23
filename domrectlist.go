// Code generated DO NOT EDIT
// domrectlist.go
package dom

import "syscall/js"

type DOMRectListIFace interface {
	Item(args ...interface{}) DOMRect
	GetLength() float64
}
type DOMRectList struct {
	Value
}

func JSValueToDOMRectList(val js.Value) DOMRectList { return DOMRectList{Value: Value{Value: val}} }
func (v Value) AsDOMRectList() DOMRectList          { return DOMRectList{Value: v} }
func (d DOMRectList) Item(args ...interface{}) DOMRect {
	val := d.Call("item", args...)
	return JSValueToDOMRect(val.JSValue())
}
func (d DOMRectList) GetLength() float64 {
	val := d.Get("length")
	return val.Float()
}
