// Code generated DO NOT EDIT
// domquad.go
package dom

import "syscall/js"

type DOMQuadIFace interface {
	FromQuad(args ...interface{}) DOMQuad
	FromRect(args ...interface{}) DOMQuad
	GetBounds(args ...interface{}) DOMRect
	GetP1() DOMPoint
	GetP2() DOMPoint
	GetP3() DOMPoint
	GetP4() DOMPoint
	ToJSON(args ...interface{}) Value
}
type DOMQuad struct {
	Value
}

func JSValueToDOMQuad(val js.Value) DOMQuad { return DOMQuad{Value: Value{Value: val}} }
func (v Value) AsDOMQuad() DOMQuad          { return DOMQuad{Value: v} }
func (d DOMQuad) FromQuad(args ...interface{}) DOMQuad {
	val := d.Call("fromQuad", args...)
	return JSValueToDOMQuad(val.JSValue())
}
func (d DOMQuad) FromRect(args ...interface{}) DOMQuad {
	val := d.Call("fromRect", args...)
	return JSValueToDOMQuad(val.JSValue())
}
func (d DOMQuad) GetBounds(args ...interface{}) DOMRect {
	val := d.Call("getBounds", args...)
	return JSValueToDOMRect(val.JSValue())
}
func (d DOMQuad) GetP1() DOMPoint {
	val := d.Get("p1")
	return JSValueToDOMPoint(val.JSValue())
}
func (d DOMQuad) GetP2() DOMPoint {
	val := d.Get("p2")
	return JSValueToDOMPoint(val.JSValue())
}
func (d DOMQuad) GetP3() DOMPoint {
	val := d.Get("p3")
	return JSValueToDOMPoint(val.JSValue())
}
func (d DOMQuad) GetP4() DOMPoint {
	val := d.Get("p4")
	return JSValueToDOMPoint(val.JSValue())
}
func (d DOMQuad) ToJSON(args ...interface{}) Value {
	val := d.Call("toJSON", args...)
	return val
}
