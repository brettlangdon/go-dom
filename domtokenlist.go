// Code generated DO NOT EDIT
// domtokenlist.go
package dom

import "syscall/js"

type DOMTokenListIFace interface {
	Add(args ...interface{})
	Contains(args ...interface{}) bool
	Item(args ...interface{}) string
	GetLength() float64
	Remove(args ...interface{})
	Replace(args ...interface{}) bool
	Supports(args ...interface{}) bool
	Toggle(args ...interface{}) bool
	GetValue() string
	SetValue(string)
}
type DOMTokenList struct {
	Value
}

func jsValueToDOMTokenList(val js.Value) DOMTokenList { return DOMTokenList{Value: Value{Value: val}} }
func (v Value) AsDOMTokenList() DOMTokenList          { return DOMTokenList{Value: v} }
func (d DOMTokenList) Add(args ...interface{}) {
	d.Call("add", args...)
}
func (d DOMTokenList) Contains(args ...interface{}) bool {
	val := d.Call("contains", args...)
	return val.Bool()
}
func (d DOMTokenList) Item(args ...interface{}) string {
	val := d.Call("item", args...)
	return val.String()
}
func (d DOMTokenList) GetLength() float64 {
	val := d.Get("length")
	return val.Float()
}
func (d DOMTokenList) Remove(args ...interface{}) {
	d.Call("remove", args...)
}
func (d DOMTokenList) Replace(args ...interface{}) bool {
	val := d.Call("replace", args...)
	return val.Bool()
}
func (d DOMTokenList) Supports(args ...interface{}) bool {
	val := d.Call("supports", args...)
	return val.Bool()
}
func (d DOMTokenList) Toggle(args ...interface{}) bool {
	val := d.Call("toggle", args...)
	return val.Bool()
}
func (d DOMTokenList) GetValue() string {
	val := d.Get("value")
	return val.String()
}
func (d DOMTokenList) SetValue(val string) {
	d.Set("value", val)
}
