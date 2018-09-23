// Code generated DO NOT EDIT
// domrectreadonly.go
package dom

import "syscall/js"

type DOMRectReadOnlyIFace interface {
	GetBottom() float64
	FromRect(args ...interface{}) DOMRectReadOnly
	GetHeight() float64
	GetLeft() float64
	GetRight() float64
	ToJSON(args ...interface{}) Value
	GetTop() float64
	GetWidth() float64
	GetX() float64
	GetY() float64
}
type DOMRectReadOnly struct {
	Value
}

func jsValueToDOMRectReadOnly(val js.Value) DOMRectReadOnly {
	return DOMRectReadOnly{Value: Value{Value: val}}
}
func (v Value) AsDOMRectReadOnly() DOMRectReadOnly { return DOMRectReadOnly{Value: v} }
func (d DOMRectReadOnly) GetBottom() float64 {
	val := d.Get("bottom")
	return val.Float()
}
func (d DOMRectReadOnly) FromRect(args ...interface{}) DOMRectReadOnly {
	val := d.Call("fromRect", args...)
	return jsValueToDOMRectReadOnly(val.JSValue())
}
func (d DOMRectReadOnly) GetHeight() float64 {
	val := d.Get("height")
	return val.Float()
}
func (d DOMRectReadOnly) GetLeft() float64 {
	val := d.Get("left")
	return val.Float()
}
func (d DOMRectReadOnly) GetRight() float64 {
	val := d.Get("right")
	return val.Float()
}
func (d DOMRectReadOnly) ToJSON(args ...interface{}) Value {
	val := d.Call("toJSON", args...)
	return val
}
func (d DOMRectReadOnly) GetTop() float64 {
	val := d.Get("top")
	return val.Float()
}
func (d DOMRectReadOnly) GetWidth() float64 {
	val := d.Get("width")
	return val.Float()
}
func (d DOMRectReadOnly) GetX() float64 {
	val := d.Get("x")
	return val.Float()
}
func (d DOMRectReadOnly) GetY() float64 {
	val := d.Get("y")
	return val.Float()
}
