// Code generated DO NOT EDIT
// barprop.go
package dom

import "syscall/js"

type BarPropIFace interface {
	GetVisible() bool
}
type BarProp struct {
	Value
}

func JSValueToBarProp(val js.Value) BarProp { return BarProp{Value: JSValueToValue(val)} }
func (v Value) AsBarProp() BarProp          { return BarProp{Value: v} }
func NewBarProp(args ...interface{}) BarProp {
	return BarProp{Value: JSValueToValue(js.Global().Get("BarProp").New(args...))}
}
func (b BarProp) GetVisible() bool {
	val := b.Get("visible")
	return val.Bool()
}
