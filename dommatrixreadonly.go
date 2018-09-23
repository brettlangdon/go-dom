// Code generated DO NOT EDIT
// dommatrixreadonly.go
package dom

import "syscall/js"

type DOMMatrixReadOnlyIFace interface {
	GetA() float64
	GetB() float64
	GetC() float64
	GetD() float64
	GetE() float64
	GetF() float64
	FlipX(args ...interface{}) DOMMatrix
	FlipY(args ...interface{}) DOMMatrix
	FromFloat32Array(args ...interface{}) DOMMatrixReadOnly
	FromFloat64Array(args ...interface{}) DOMMatrixReadOnly
	FromMatrix(args ...interface{}) DOMMatrixReadOnly
	Inverse(args ...interface{}) DOMMatrix
	GetIs2D() bool
	GetIsIdentity() bool
	GetM11() float64
	GetM12() float64
	GetM13() float64
	GetM14() float64
	GetM21() float64
	GetM22() float64
	GetM23() float64
	GetM24() float64
	GetM31() float64
	GetM32() float64
	GetM33() float64
	GetM34() float64
	GetM41() float64
	GetM42() float64
	GetM43() float64
	GetM44() float64
	Multiply(args ...interface{}) DOMMatrix
	Rotate(args ...interface{}) DOMMatrix
	RotateAxisAngle(args ...interface{}) DOMMatrix
	RotateFromVector(args ...interface{}) DOMMatrix
	Scale(args ...interface{}) DOMMatrix
	Scale3d(args ...interface{}) DOMMatrix
	SkewX(args ...interface{}) DOMMatrix
	SkewY(args ...interface{}) DOMMatrix
	ToFloat32Array(args ...interface{}) Float32Array
	ToFloat64Array(args ...interface{}) Float64Array
	ToJSON(args ...interface{}) Value
	TransformPoint(args ...interface{}) DOMPoint
	Translate(args ...interface{}) DOMMatrix
}
type DOMMatrixReadOnly struct {
	Value
}

func jsValueToDOMMatrixReadOnly(val js.Value) DOMMatrixReadOnly {
	return DOMMatrixReadOnly{Value: Value{Value: val}}
}
func (v Value) AsDOMMatrixReadOnly() DOMMatrixReadOnly { return DOMMatrixReadOnly{Value: v} }
func (d DOMMatrixReadOnly) GetA() float64 {
	val := d.Get("a")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetB() float64 {
	val := d.Get("b")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetC() float64 {
	val := d.Get("c")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetD() float64 {
	val := d.Get("d")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetE() float64 {
	val := d.Get("e")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetF() float64 {
	val := d.Get("f")
	return val.Float()
}
func (d DOMMatrixReadOnly) FlipX(args ...interface{}) DOMMatrix {
	val := d.Call("flipX", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) FlipY(args ...interface{}) DOMMatrix {
	val := d.Call("flipY", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) FromFloat32Array(args ...interface{}) DOMMatrixReadOnly {
	val := d.Call("fromFloat32Array", args...)
	return jsValueToDOMMatrixReadOnly(val.JSValue())
}
func (d DOMMatrixReadOnly) FromFloat64Array(args ...interface{}) DOMMatrixReadOnly {
	val := d.Call("fromFloat64Array", args...)
	return jsValueToDOMMatrixReadOnly(val.JSValue())
}
func (d DOMMatrixReadOnly) FromMatrix(args ...interface{}) DOMMatrixReadOnly {
	val := d.Call("fromMatrix", args...)
	return jsValueToDOMMatrixReadOnly(val.JSValue())
}
func (d DOMMatrixReadOnly) Inverse(args ...interface{}) DOMMatrix {
	val := d.Call("inverse", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) GetIs2D() bool {
	val := d.Get("is2D")
	return val.Bool()
}
func (d DOMMatrixReadOnly) GetIsIdentity() bool {
	val := d.Get("isIdentity")
	return val.Bool()
}
func (d DOMMatrixReadOnly) GetM11() float64 {
	val := d.Get("m11")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM12() float64 {
	val := d.Get("m12")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM13() float64 {
	val := d.Get("m13")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM14() float64 {
	val := d.Get("m14")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM21() float64 {
	val := d.Get("m21")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM22() float64 {
	val := d.Get("m22")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM23() float64 {
	val := d.Get("m23")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM24() float64 {
	val := d.Get("m24")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM31() float64 {
	val := d.Get("m31")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM32() float64 {
	val := d.Get("m32")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM33() float64 {
	val := d.Get("m33")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM34() float64 {
	val := d.Get("m34")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM41() float64 {
	val := d.Get("m41")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM42() float64 {
	val := d.Get("m42")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM43() float64 {
	val := d.Get("m43")
	return val.Float()
}
func (d DOMMatrixReadOnly) GetM44() float64 {
	val := d.Get("m44")
	return val.Float()
}
func (d DOMMatrixReadOnly) Multiply(args ...interface{}) DOMMatrix {
	val := d.Call("multiply", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) Rotate(args ...interface{}) DOMMatrix {
	val := d.Call("rotate", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) RotateAxisAngle(args ...interface{}) DOMMatrix {
	val := d.Call("rotateAxisAngle", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) RotateFromVector(args ...interface{}) DOMMatrix {
	val := d.Call("rotateFromVector", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) Scale(args ...interface{}) DOMMatrix {
	val := d.Call("scale", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) Scale3d(args ...interface{}) DOMMatrix {
	val := d.Call("scale3d", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) SkewX(args ...interface{}) DOMMatrix {
	val := d.Call("skewX", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) SkewY(args ...interface{}) DOMMatrix {
	val := d.Call("skewY", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrixReadOnly) ToFloat32Array(args ...interface{}) Float32Array {
	val := d.Call("toFloat32Array", args...)
	return jsValueToFloat32Array(val.JSValue())
}
func (d DOMMatrixReadOnly) ToFloat64Array(args ...interface{}) Float64Array {
	val := d.Call("toFloat64Array", args...)
	return jsValueToFloat64Array(val.JSValue())
}
func (d DOMMatrixReadOnly) ToJSON(args ...interface{}) Value {
	val := d.Call("toJSON", args...)
	return val
}
func (d DOMMatrixReadOnly) TransformPoint(args ...interface{}) DOMPoint {
	val := d.Call("transformPoint", args...)
	return jsValueToDOMPoint(val.JSValue())
}
func (d DOMMatrixReadOnly) Translate(args ...interface{}) DOMMatrix {
	val := d.Call("translate", args...)
	return jsValueToDOMMatrix(val.JSValue())
}
