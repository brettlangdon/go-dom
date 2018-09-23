// Code generated DO NOT EDIT
// domstringlist.go
package dom

import "syscall/js"

type DOMStringListIFace interface {
	Contains(args ...interface{}) bool
	Item(args ...interface{}) string
	GetLength() float64
}
type DOMStringList struct {
	Value
}

func jsValueToDOMStringList(val js.Value) DOMStringList {
	return DOMStringList{Value: Value{Value: val}}
}
func (v Value) AsDOMStringList() DOMStringList { return DOMStringList{Value: v} }
func (d DOMStringList) Contains(args ...interface{}) bool {
	val := d.Call("contains", args...)
	return val.Bool()
}
func (d DOMStringList) Item(args ...interface{}) string {
	val := d.Call("item", args...)
	return val.String()
}
func (d DOMStringList) GetLength() float64 {
	val := d.Get("length")
	return val.Float()
}
