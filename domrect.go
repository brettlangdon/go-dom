// Code generated DO NOT EDIT
// domrect.go
package dom

import "syscall/js"

type DOMRectIFace interface {
	GetBottom() float64
	FromRect(args ...interface{}) DOMRect
	GetHeight() float64
	SetHeight(float64)
	GetLeft() float64
	GetRight() float64
	ToJSON(args ...interface{}) Value
	GetTop() float64
	GetWidth() float64
	SetWidth(float64)
	GetX() float64
	SetX(float64)
	GetY() float64
	SetY(float64)
}
type DOMRect struct {
	Value
	DOMRectReadOnly
}

func jsValueToDOMRect(val js.Value) DOMRect { return DOMRect{Value: Value{Value: val}} }
func (v Value) AsDOMRect() DOMRect          { return DOMRect{Value: v} }
func (d DOMRect) FromRect(args ...interface{}) DOMRect {
	val := d.Call("fromRect", args...)
	return jsValueToDOMRect(val.JSValue())
}
func (d DOMRect) GetHeight() float64 {
	val := d.Get("height")
	return val.Float()
}
func (d DOMRect) SetHeight(val float64) {
	d.Set("height", val)
}
func (d DOMRect) GetWidth() float64 {
	val := d.Get("width")
	return val.Float()
}
func (d DOMRect) SetWidth(val float64) {
	d.Set("width", val)
}
func (d DOMRect) GetX() float64 {
	val := d.Get("x")
	return val.Float()
}
func (d DOMRect) SetX(val float64) {
	d.Set("x", val)
}
func (d DOMRect) GetY() float64 {
	val := d.Get("y")
	return val.Float()
}
func (d DOMRect) SetY(val float64) {
	d.Set("y", val)
}
