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

func JSValueToBarProp(val js.Value) BarProp { return BarProp{Value: Value{Value: val}} }
func (v Value) AsBarProp() BarProp          { return BarProp{Value: v} }
func (b BarProp) GetVisible() bool {
	val := b.Get("visible")
	return val.Bool()
}
