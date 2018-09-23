// Code generated DO NOT EDIT
// dompoint.go
package dom

import "syscall/js"

type DOMPointIFace interface {
	FromPoint(args ...interface{}) DOMPoint
	MatrixTransform(args ...interface{}) DOMPoint
	ToJSON(args ...interface{}) Value
	GetW() float64
	SetW(float64)
	GetX() float64
	SetX(float64)
	GetY() float64
	SetY(float64)
	GetZ() float64
	SetZ(float64)
}
type DOMPoint struct {
	Value
	DOMPointReadOnly
}

func jsValueToDOMPoint(val js.Value) DOMPoint { return DOMPoint{Value: Value{Value: val}} }
func (v Value) AsDOMPoint() DOMPoint          { return DOMPoint{Value: v} }
func (d DOMPoint) FromPoint(args ...interface{}) DOMPoint {
	val := d.Call("fromPoint", args...)
	return jsValueToDOMPoint(val.JSValue())
}
func (d DOMPoint) GetW() float64 {
	val := d.Get("w")
	return val.Float()
}
func (d DOMPoint) SetW(val float64) {
	d.Set("w", val)
}
func (d DOMPoint) GetX() float64 {
	val := d.Get("x")
	return val.Float()
}
func (d DOMPoint) SetX(val float64) {
	d.Set("x", val)
}
func (d DOMPoint) GetY() float64 {
	val := d.Get("y")
	return val.Float()
}
func (d DOMPoint) SetY(val float64) {
	d.Set("y", val)
}
func (d DOMPoint) GetZ() float64 {
	val := d.Get("z")
	return val.Float()
}
func (d DOMPoint) SetZ(val float64) {
	d.Set("z", val)
}
