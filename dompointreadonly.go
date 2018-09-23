// Code generated DO NOT EDIT
// dompointreadonly.go
package dom

import "syscall/js"

type DOMPointReadOnlyIFace interface {
	FromPoint(args ...interface{}) DOMPointReadOnly
	MatrixTransform(args ...interface{}) DOMPoint
	ToJSON(args ...interface{}) Value
	GetW() float64
	GetX() float64
	GetY() float64
	GetZ() float64
}
type DOMPointReadOnly struct {
	Value
}

func jsValueToDOMPointReadOnly(val js.Value) DOMPointReadOnly {
	return DOMPointReadOnly{Value: Value{Value: val}}
}
func (v Value) AsDOMPointReadOnly() DOMPointReadOnly { return DOMPointReadOnly{Value: v} }
func (d DOMPointReadOnly) FromPoint(args ...interface{}) DOMPointReadOnly {
	val := d.Call("fromPoint", args...)
	return jsValueToDOMPointReadOnly(val.JSValue())
}
func (d DOMPointReadOnly) MatrixTransform(args ...interface{}) DOMPoint {
	val := d.Call("matrixTransform", args...)
	return jsValueToDOMPoint(val.JSValue())
}
func (d DOMPointReadOnly) ToJSON(args ...interface{}) Value {
	val := d.Call("toJSON", args...)
	return val
}
func (d DOMPointReadOnly) GetW() float64 {
	val := d.Get("w")
	return val.Float()
}
func (d DOMPointReadOnly) GetX() float64 {
	val := d.Get("x")
	return val.Float()
}
func (d DOMPointReadOnly) GetY() float64 {
	val := d.Get("y")
	return val.Float()
}
func (d DOMPointReadOnly) GetZ() float64 {
	val := d.Get("z")
	return val.Float()
}
