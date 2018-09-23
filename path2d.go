// Code generated DO NOT EDIT
// path2d.go
package dom

import "syscall/js"

type Path2DIFace interface {
	AddPath(args ...interface{})
	Arc(args ...interface{})
	ArcTo(args ...interface{})
	BezierCurveTo(args ...interface{})
	ClosePath(args ...interface{})
	Ellipse(args ...interface{})
	LineTo(args ...interface{})
	MoveTo(args ...interface{})
	QuadraticCurveTo(args ...interface{})
	Rect(args ...interface{})
}
type Path2D struct {
	Value
}

func jsValueToPath2D(val js.Value) Path2D { return Path2D{Value: Value{Value: val}} }
func (v Value) AsPath2D() Path2D          { return Path2D{Value: v} }
func (p Path2D) AddPath(args ...interface{}) {
	p.Call("addPath", args...)
}
func (p Path2D) Arc(args ...interface{}) {
	p.Call("arc", args...)
}
func (p Path2D) ArcTo(args ...interface{}) {
	p.Call("arcTo", args...)
}
func (p Path2D) BezierCurveTo(args ...interface{}) {
	p.Call("bezierCurveTo", args...)
}
func (p Path2D) ClosePath(args ...interface{}) {
	p.Call("closePath", args...)
}
func (p Path2D) Ellipse(args ...interface{}) {
	p.Call("ellipse", args...)
}
func (p Path2D) LineTo(args ...interface{}) {
	p.Call("lineTo", args...)
}
func (p Path2D) MoveTo(args ...interface{}) {
	p.Call("moveTo", args...)
}
func (p Path2D) QuadraticCurveTo(args ...interface{}) {
	p.Call("quadraticCurveTo", args...)
}
func (p Path2D) Rect(args ...interface{}) {
	p.Call("rect", args...)
}
