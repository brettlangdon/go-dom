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
}

func JSValueToDOMRect(val js.Value) DOMRect { return DOMRect{Value: JSValueToValue(val)} }
func (v Value) AsDOMRect() DOMRect          { return DOMRect{Value: v} }
func NewDOMRect(args ...interface{}) DOMRect {
	return DOMRect{Value: JSValueToValue(js.Global().Get("DOMRect").New(args...))}
}
func (d DOMRect) GetBottom() float64 {
	val := d.Get("bottom")
	return val.Float()
}
func (d DOMRect) FromRect(args ...interface{}) DOMRect {
	val := d.Call("fromRect", args...)
	return JSValueToDOMRect(val.JSValue())
}
func (d DOMRect) GetHeight() float64 {
	val := d.Get("height")
	return val.Float()
}
func (d DOMRect) SetHeight(val float64) {
	d.Set("height", val)
}
func (d DOMRect) GetLeft() float64 {
	val := d.Get("left")
	return val.Float()
}
func (d DOMRect) GetRight() float64 {
	val := d.Get("right")
	return val.Float()
}
func (d DOMRect) ToJSON(args ...interface{}) Value {
	val := d.Call("toJSON", args...)
	return val
}
func (d DOMRect) GetTop() float64 {
	val := d.Get("top")
	return val.Float()
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
