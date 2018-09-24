// Code generated DO NOT EDIT
// domstringlist.go
package dom

import "syscall/js"

type DOMStringListIFace interface {
	Contains(args ...interface{}) bool
	Item(args ...interface{}) string
	GetLength() int
}
type DOMStringList struct {
	Value
}

func JSValueToDOMStringList(val js.Value) DOMStringList {
	return DOMStringList{Value: JSValueToValue(val)}
}
func (v Value) AsDOMStringList() DOMStringList { return DOMStringList{Value: v} }
func NewDOMStringList(args ...interface{}) DOMStringList {
	return DOMStringList{Value: JSValueToValue(js.Global().Get("DOMStringList").New(args...))}
}
func (d DOMStringList) Contains(args ...interface{}) bool {
	val := d.Call("contains", args...)
	return val.Bool()
}
func (d DOMStringList) Item(args ...interface{}) string {
	val := d.Call("item", args...)
	return val.String()
}
func (d DOMStringList) GetLength() int {
	val := d.Get("length")
	return val.Int()
}
