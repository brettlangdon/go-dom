// Code generated DO NOT EDIT
// dommatrix.go
package dom

import "syscall/js"

type DOMMatrixIFace interface {
	GetA() float64
	SetA(float64)
	GetB() float64
	SetB(float64)
	GetC() float64
	SetC(float64)
	GetD() float64
	SetD(float64)
	GetE() float64
	SetE(float64)
	GetF() float64
	SetF(float64)
	FlipX(args ...interface{}) DOMMatrix
	FlipY(args ...interface{}) DOMMatrix
	FromFloat32Array(args ...interface{}) DOMMatrix
	FromFloat64Array(args ...interface{}) DOMMatrix
	FromMatrix(args ...interface{}) DOMMatrix
	Inverse(args ...interface{}) DOMMatrix
	InvertSelf(args ...interface{}) DOMMatrix
	GetIs2D() bool
	GetIsIdentity() bool
	GetM11() float64
	SetM11(float64)
	GetM12() float64
	SetM12(float64)
	GetM13() float64
	SetM13(float64)
	GetM14() float64
	SetM14(float64)
	GetM21() float64
	SetM21(float64)
	GetM22() float64
	SetM22(float64)
	GetM23() float64
	SetM23(float64)
	GetM24() float64
	SetM24(float64)
	GetM31() float64
	SetM31(float64)
	GetM32() float64
	SetM32(float64)
	GetM33() float64
	SetM33(float64)
	GetM34() float64
	SetM34(float64)
	GetM41() float64
	SetM41(float64)
	GetM42() float64
	SetM42(float64)
	GetM43() float64
	SetM43(float64)
	GetM44() float64
	SetM44(float64)
	Multiply(args ...interface{}) DOMMatrix
	MultiplySelf(args ...interface{}) DOMMatrix
	PreMultiplySelf(args ...interface{}) DOMMatrix
	Rotate(args ...interface{}) DOMMatrix
	RotateAxisAngle(args ...interface{}) DOMMatrix
	RotateAxisAngleSelf(args ...interface{}) DOMMatrix
	RotateFromVector(args ...interface{}) DOMMatrix
	RotateFromVectorSelf(args ...interface{}) DOMMatrix
	RotateSelf(args ...interface{}) DOMMatrix
	Scale(args ...interface{}) DOMMatrix
	Scale3d(args ...interface{}) DOMMatrix
	Scale3dSelf(args ...interface{}) DOMMatrix
	ScaleSelf(args ...interface{}) DOMMatrix
	SetMatrixValue(args ...interface{}) DOMMatrix
	SkewX(args ...interface{}) DOMMatrix
	SkewXSelf(args ...interface{}) DOMMatrix
	SkewY(args ...interface{}) DOMMatrix
	SkewYSelf(args ...interface{}) DOMMatrix
	ToFloat32Array(args ...interface{}) Float32Array
	ToFloat64Array(args ...interface{}) Float64Array
	ToJSON(args ...interface{}) Value
	TransformPoint(args ...interface{}) DOMPoint
	Translate(args ...interface{}) DOMMatrix
	TranslateSelf(args ...interface{}) DOMMatrix
}
type DOMMatrix struct {
	Value
	DOMMatrixReadOnly
}

func JSValueToDOMMatrix(val js.Value) DOMMatrix { return DOMMatrix{Value: Value{Value: val}} }
func (v Value) AsDOMMatrix() DOMMatrix          { return DOMMatrix{Value: v} }
func (d DOMMatrix) GetA() float64 {
	val := d.Get("a")
	return val.Float()
}
func (d DOMMatrix) SetA(val float64) {
	d.Set("a", val)
}
func (d DOMMatrix) GetB() float64 {
	val := d.Get("b")
	return val.Float()
}
func (d DOMMatrix) SetB(val float64) {
	d.Set("b", val)
}
func (d DOMMatrix) GetC() float64 {
	val := d.Get("c")
	return val.Float()
}
func (d DOMMatrix) SetC(val float64) {
	d.Set("c", val)
}
func (d DOMMatrix) GetD() float64 {
	val := d.Get("d")
	return val.Float()
}
func (d DOMMatrix) SetD(val float64) {
	d.Set("d", val)
}
func (d DOMMatrix) GetE() float64 {
	val := d.Get("e")
	return val.Float()
}
func (d DOMMatrix) SetE(val float64) {
	d.Set("e", val)
}
func (d DOMMatrix) GetF() float64 {
	val := d.Get("f")
	return val.Float()
}
func (d DOMMatrix) SetF(val float64) {
	d.Set("f", val)
}
func (d DOMMatrix) FromFloat32Array(args ...interface{}) DOMMatrix {
	val := d.Call("fromFloat32Array", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) FromFloat64Array(args ...interface{}) DOMMatrix {
	val := d.Call("fromFloat64Array", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) FromMatrix(args ...interface{}) DOMMatrix {
	val := d.Call("fromMatrix", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) InvertSelf(args ...interface{}) DOMMatrix {
	val := d.Call("invertSelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) GetM11() float64 {
	val := d.Get("m11")
	return val.Float()
}
func (d DOMMatrix) SetM11(val float64) {
	d.Set("m11", val)
}
func (d DOMMatrix) GetM12() float64 {
	val := d.Get("m12")
	return val.Float()
}
func (d DOMMatrix) SetM12(val float64) {
	d.Set("m12", val)
}
func (d DOMMatrix) GetM13() float64 {
	val := d.Get("m13")
	return val.Float()
}
func (d DOMMatrix) SetM13(val float64) {
	d.Set("m13", val)
}
func (d DOMMatrix) GetM14() float64 {
	val := d.Get("m14")
	return val.Float()
}
func (d DOMMatrix) SetM14(val float64) {
	d.Set("m14", val)
}
func (d DOMMatrix) GetM21() float64 {
	val := d.Get("m21")
	return val.Float()
}
func (d DOMMatrix) SetM21(val float64) {
	d.Set("m21", val)
}
func (d DOMMatrix) GetM22() float64 {
	val := d.Get("m22")
	return val.Float()
}
func (d DOMMatrix) SetM22(val float64) {
	d.Set("m22", val)
}
func (d DOMMatrix) GetM23() float64 {
	val := d.Get("m23")
	return val.Float()
}
func (d DOMMatrix) SetM23(val float64) {
	d.Set("m23", val)
}
func (d DOMMatrix) GetM24() float64 {
	val := d.Get("m24")
	return val.Float()
}
func (d DOMMatrix) SetM24(val float64) {
	d.Set("m24", val)
}
func (d DOMMatrix) GetM31() float64 {
	val := d.Get("m31")
	return val.Float()
}
func (d DOMMatrix) SetM31(val float64) {
	d.Set("m31", val)
}
func (d DOMMatrix) GetM32() float64 {
	val := d.Get("m32")
	return val.Float()
}
func (d DOMMatrix) SetM32(val float64) {
	d.Set("m32", val)
}
func (d DOMMatrix) GetM33() float64 {
	val := d.Get("m33")
	return val.Float()
}
func (d DOMMatrix) SetM33(val float64) {
	d.Set("m33", val)
}
func (d DOMMatrix) GetM34() float64 {
	val := d.Get("m34")
	return val.Float()
}
func (d DOMMatrix) SetM34(val float64) {
	d.Set("m34", val)
}
func (d DOMMatrix) GetM41() float64 {
	val := d.Get("m41")
	return val.Float()
}
func (d DOMMatrix) SetM41(val float64) {
	d.Set("m41", val)
}
func (d DOMMatrix) GetM42() float64 {
	val := d.Get("m42")
	return val.Float()
}
func (d DOMMatrix) SetM42(val float64) {
	d.Set("m42", val)
}
func (d DOMMatrix) GetM43() float64 {
	val := d.Get("m43")
	return val.Float()
}
func (d DOMMatrix) SetM43(val float64) {
	d.Set("m43", val)
}
func (d DOMMatrix) GetM44() float64 {
	val := d.Get("m44")
	return val.Float()
}
func (d DOMMatrix) SetM44(val float64) {
	d.Set("m44", val)
}
func (d DOMMatrix) MultiplySelf(args ...interface{}) DOMMatrix {
	val := d.Call("multiplySelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) PreMultiplySelf(args ...interface{}) DOMMatrix {
	val := d.Call("preMultiplySelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) RotateAxisAngleSelf(args ...interface{}) DOMMatrix {
	val := d.Call("rotateAxisAngleSelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) RotateFromVectorSelf(args ...interface{}) DOMMatrix {
	val := d.Call("rotateFromVectorSelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) RotateSelf(args ...interface{}) DOMMatrix {
	val := d.Call("rotateSelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) Scale3dSelf(args ...interface{}) DOMMatrix {
	val := d.Call("scale3dSelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) ScaleSelf(args ...interface{}) DOMMatrix {
	val := d.Call("scaleSelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) SetMatrixValue(args ...interface{}) DOMMatrix {
	val := d.Call("setMatrixValue", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) SkewXSelf(args ...interface{}) DOMMatrix {
	val := d.Call("skewXSelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) SkewYSelf(args ...interface{}) DOMMatrix {
	val := d.Call("skewYSelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (d DOMMatrix) TranslateSelf(args ...interface{}) DOMMatrix {
	val := d.Call("translateSelf", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
