// Code generated DO NOT EDIT
// offscreencanvasrenderingcontext2d.go
package dom

import "syscall/js"

type OffscreenCanvasRenderingContext2DIFace interface {
	Arc(args ...interface{})
	ArcTo(args ...interface{})
	BeginPath(args ...interface{})
	BezierCurveTo(args ...interface{})
	GetCanvas() OffscreenCanvas
	ClearRect(args ...interface{})
	Clip(args ...interface{})
	ClipWithArgs(args ...interface{})
	ClosePath(args ...interface{})
	Commit(args ...interface{})
	CreateImageData(args ...interface{}) ImageData
	CreateImageDataWithArgs(args ...interface{}) ImageData
	CreateLinearGradient(args ...interface{}) CanvasGradient
	CreatePattern(args ...interface{}) CanvasPattern
	CreateRadialGradient(args ...interface{}) CanvasGradient
	GetDirection() CanvasDirection
	SetDirection(CanvasDirection)
	DrawImage(args ...interface{})
	Ellipse(args ...interface{})
	Fill(args ...interface{})
	FillRect(args ...interface{})
	GetFillStyle()
	SetFillStyle()
	FillText(args ...interface{})
	FillWithArgs(args ...interface{})
	GetFilter() string
	SetFilter(string)
	GetFont() string
	SetFont(string)
	GetImageData(args ...interface{}) ImageData
	GetLineDash(args ...interface{})
	GetTransform(args ...interface{}) DOMMatrix
	GetGlobalAlpha() float64
	SetGlobalAlpha(float64)
	GetGlobalCompositeOperation() string
	SetGlobalCompositeOperation(string)
	GetImageSmoothingEnabled() bool
	SetImageSmoothingEnabled(bool)
	GetImageSmoothingQuality() ImageSmoothingQuality
	SetImageSmoothingQuality(ImageSmoothingQuality)
	IsPointInPath(args ...interface{}) bool
	IsPointInPathWithArgs(args ...interface{}) bool
	IsPointInStroke(args ...interface{}) bool
	IsPointInStrokeWithArgs(args ...interface{}) bool
	GetLineCap() CanvasLineCap
	SetLineCap(CanvasLineCap)
	GetLineDashOffset() float64
	SetLineDashOffset(float64)
	GetLineJoin() CanvasLineJoin
	SetLineJoin(CanvasLineJoin)
	LineTo(args ...interface{})
	GetLineWidth() float64
	SetLineWidth(float64)
	MeasureText(args ...interface{}) TextMetrics
	GetMiterLimit() float64
	SetMiterLimit(float64)
	MoveTo(args ...interface{})
	PutImageData(args ...interface{})
	PutImageDataWithArgs(args ...interface{})
	QuadraticCurveTo(args ...interface{})
	Rect(args ...interface{})
	ResetTransform(args ...interface{})
	Restore(args ...interface{})
	Rotate(args ...interface{})
	Save(args ...interface{})
	Scale(args ...interface{})
	SetLineDash(args ...interface{})
	SetTransform(args ...interface{})
	SetTransformWithArgs(args ...interface{})
	GetShadowBlur() float64
	SetShadowBlur(float64)
	GetShadowColor() string
	SetShadowColor(string)
	GetShadowOffsetX() float64
	SetShadowOffsetX(float64)
	GetShadowOffsetY() float64
	SetShadowOffsetY(float64)
	Stroke(args ...interface{})
	StrokeRect(args ...interface{})
	GetStrokeStyle()
	SetStrokeStyle()
	StrokeText(args ...interface{})
	StrokeWithArgs(args ...interface{})
	GetTextAlign() CanvasTextAlign
	SetTextAlign(CanvasTextAlign)
	GetTextBaseline() CanvasTextBaseline
	SetTextBaseline(CanvasTextBaseline)
	Transform(args ...interface{})
	Translate(args ...interface{})
}
type OffscreenCanvasRenderingContext2D struct {
	Value
}

func JSValueToOffscreenCanvasRenderingContext2D(val js.Value) OffscreenCanvasRenderingContext2D {
	return OffscreenCanvasRenderingContext2D{Value: Value{Value: val}}
}
func (v Value) AsOffscreenCanvasRenderingContext2D() OffscreenCanvasRenderingContext2D {
	return OffscreenCanvasRenderingContext2D{Value: v}
}
func (o OffscreenCanvasRenderingContext2D) Arc(args ...interface{}) {
	o.Call("arc", args...)
}
func (o OffscreenCanvasRenderingContext2D) ArcTo(args ...interface{}) {
	o.Call("arcTo", args...)
}
func (o OffscreenCanvasRenderingContext2D) BeginPath(args ...interface{}) {
	o.Call("beginPath", args...)
}
func (o OffscreenCanvasRenderingContext2D) BezierCurveTo(args ...interface{}) {
	o.Call("bezierCurveTo", args...)
}
func (o OffscreenCanvasRenderingContext2D) GetCanvas() OffscreenCanvas {
	val := o.Get("canvas")
	return JSValueToOffscreenCanvas(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) ClearRect(args ...interface{}) {
	o.Call("clearRect", args...)
}
func (o OffscreenCanvasRenderingContext2D) Clip(args ...interface{}) {
	o.Call("clip", args...)
}
func (o OffscreenCanvasRenderingContext2D) ClipWithArgs(args ...interface{}) {
	o.Call("clipWithArgs", args...)
}
func (o OffscreenCanvasRenderingContext2D) ClosePath(args ...interface{}) {
	o.Call("closePath", args...)
}
func (o OffscreenCanvasRenderingContext2D) Commit(args ...interface{}) {
	o.Call("commit", args...)
}
func (o OffscreenCanvasRenderingContext2D) CreateImageData(args ...interface{}) ImageData {
	val := o.Call("createImageData", args...)
	return JSValueToImageData(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) CreateImageDataWithArgs(args ...interface{}) ImageData {
	val := o.Call("createImageDataWithArgs", args...)
	return JSValueToImageData(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) CreateLinearGradient(args ...interface{}) CanvasGradient {
	val := o.Call("createLinearGradient", args...)
	return JSValueToCanvasGradient(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) CreatePattern(args ...interface{}) CanvasPattern {
	val := o.Call("createPattern", args...)
	return JSValueToCanvasPattern(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) CreateRadialGradient(args ...interface{}) CanvasGradient {
	val := o.Call("createRadialGradient", args...)
	return JSValueToCanvasGradient(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) GetDirection() CanvasDirection {
	val := o.Get("direction")
	return JSValueToCanvasDirection(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) SetDirection(val CanvasDirection) {
	o.Set("direction", val)
}
func (o OffscreenCanvasRenderingContext2D) DrawImage(args ...interface{}) {
	o.Call("drawImage", args...)
}
func (o OffscreenCanvasRenderingContext2D) Ellipse(args ...interface{}) {
	o.Call("ellipse", args...)
}
func (o OffscreenCanvasRenderingContext2D) Fill(args ...interface{}) {
	o.Call("fill", args...)
}
func (o OffscreenCanvasRenderingContext2D) FillRect(args ...interface{}) {
	o.Call("fillRect", args...)
}
func (o OffscreenCanvasRenderingContext2D) GetFillStyle() Value {
	val := o.Get("fillStyle")
	return val
}
func (o OffscreenCanvasRenderingContext2D) SetFillStyle(val Value) {
	o.Set("fillStyle", val)
}
func (o OffscreenCanvasRenderingContext2D) FillText(args ...interface{}) {
	o.Call("fillText", args...)
}
func (o OffscreenCanvasRenderingContext2D) FillWithArgs(args ...interface{}) {
	o.Call("fillWithArgs", args...)
}
func (o OffscreenCanvasRenderingContext2D) GetFilter() string {
	val := o.Get("filter")
	return val.String()
}
func (o OffscreenCanvasRenderingContext2D) SetFilter(val string) {
	o.Set("filter", val)
}
func (o OffscreenCanvasRenderingContext2D) GetFont() string {
	val := o.Get("font")
	return val.String()
}
func (o OffscreenCanvasRenderingContext2D) SetFont(val string) {
	o.Set("font", val)
}
func (o OffscreenCanvasRenderingContext2D) GetImageData(args ...interface{}) ImageData {
	val := o.Call("getImageData", args...)
	return JSValueToImageData(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) GetLineDash(args ...interface{}) {
	o.Call("getLineDash", args...)
}
func (o OffscreenCanvasRenderingContext2D) GetTransform(args ...interface{}) DOMMatrix {
	val := o.Call("getTransform", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) GetGlobalAlpha() float64 {
	val := o.Get("globalAlpha")
	return val.Float()
}
func (o OffscreenCanvasRenderingContext2D) SetGlobalAlpha(val float64) {
	o.Set("globalAlpha", val)
}
func (o OffscreenCanvasRenderingContext2D) GetGlobalCompositeOperation() string {
	val := o.Get("globalCompositeOperation")
	return val.String()
}
func (o OffscreenCanvasRenderingContext2D) SetGlobalCompositeOperation(val string) {
	o.Set("globalCompositeOperation", val)
}
func (o OffscreenCanvasRenderingContext2D) GetImageSmoothingEnabled() bool {
	val := o.Get("imageSmoothingEnabled")
	return val.Bool()
}
func (o OffscreenCanvasRenderingContext2D) SetImageSmoothingEnabled(val bool) {
	o.Set("imageSmoothingEnabled", val)
}
func (o OffscreenCanvasRenderingContext2D) GetImageSmoothingQuality() ImageSmoothingQuality {
	val := o.Get("imageSmoothingQuality")
	return JSValueToImageSmoothingQuality(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) SetImageSmoothingQuality(val ImageSmoothingQuality) {
	o.Set("imageSmoothingQuality", val)
}
func (o OffscreenCanvasRenderingContext2D) IsPointInPath(args ...interface{}) bool {
	val := o.Call("isPointInPath", args...)
	return val.Bool()
}
func (o OffscreenCanvasRenderingContext2D) IsPointInPathWithArgs(args ...interface{}) bool {
	val := o.Call("isPointInPathWithArgs", args...)
	return val.Bool()
}
func (o OffscreenCanvasRenderingContext2D) IsPointInStroke(args ...interface{}) bool {
	val := o.Call("isPointInStroke", args...)
	return val.Bool()
}
func (o OffscreenCanvasRenderingContext2D) IsPointInStrokeWithArgs(args ...interface{}) bool {
	val := o.Call("isPointInStrokeWithArgs", args...)
	return val.Bool()
}
func (o OffscreenCanvasRenderingContext2D) GetLineCap() CanvasLineCap {
	val := o.Get("lineCap")
	return JSValueToCanvasLineCap(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) SetLineCap(val CanvasLineCap) {
	o.Set("lineCap", val)
}
func (o OffscreenCanvasRenderingContext2D) GetLineDashOffset() float64 {
	val := o.Get("lineDashOffset")
	return val.Float()
}
func (o OffscreenCanvasRenderingContext2D) SetLineDashOffset(val float64) {
	o.Set("lineDashOffset", val)
}
func (o OffscreenCanvasRenderingContext2D) GetLineJoin() CanvasLineJoin {
	val := o.Get("lineJoin")
	return JSValueToCanvasLineJoin(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) SetLineJoin(val CanvasLineJoin) {
	o.Set("lineJoin", val)
}
func (o OffscreenCanvasRenderingContext2D) LineTo(args ...interface{}) {
	o.Call("lineTo", args...)
}
func (o OffscreenCanvasRenderingContext2D) GetLineWidth() float64 {
	val := o.Get("lineWidth")
	return val.Float()
}
func (o OffscreenCanvasRenderingContext2D) SetLineWidth(val float64) {
	o.Set("lineWidth", val)
}
func (o OffscreenCanvasRenderingContext2D) MeasureText(args ...interface{}) TextMetrics {
	val := o.Call("measureText", args...)
	return JSValueToTextMetrics(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) GetMiterLimit() float64 {
	val := o.Get("miterLimit")
	return val.Float()
}
func (o OffscreenCanvasRenderingContext2D) SetMiterLimit(val float64) {
	o.Set("miterLimit", val)
}
func (o OffscreenCanvasRenderingContext2D) MoveTo(args ...interface{}) {
	o.Call("moveTo", args...)
}
func (o OffscreenCanvasRenderingContext2D) PutImageData(args ...interface{}) {
	o.Call("putImageData", args...)
}
func (o OffscreenCanvasRenderingContext2D) PutImageDataWithArgs(args ...interface{}) {
	o.Call("putImageDataWithArgs", args...)
}
func (o OffscreenCanvasRenderingContext2D) QuadraticCurveTo(args ...interface{}) {
	o.Call("quadraticCurveTo", args...)
}
func (o OffscreenCanvasRenderingContext2D) Rect(args ...interface{}) {
	o.Call("rect", args...)
}
func (o OffscreenCanvasRenderingContext2D) ResetTransform(args ...interface{}) {
	o.Call("resetTransform", args...)
}
func (o OffscreenCanvasRenderingContext2D) Restore(args ...interface{}) {
	o.Call("restore", args...)
}
func (o OffscreenCanvasRenderingContext2D) Rotate(args ...interface{}) {
	o.Call("rotate", args...)
}
func (o OffscreenCanvasRenderingContext2D) Save(args ...interface{}) {
	o.Call("save", args...)
}
func (o OffscreenCanvasRenderingContext2D) Scale(args ...interface{}) {
	o.Call("scale", args...)
}
func (o OffscreenCanvasRenderingContext2D) SetLineDash(args ...interface{}) {
	o.Call("setLineDash", args...)
}
func (o OffscreenCanvasRenderingContext2D) SetTransform(args ...interface{}) {
	o.Call("setTransform", args...)
}
func (o OffscreenCanvasRenderingContext2D) SetTransformWithArgs(args ...interface{}) {
	o.Call("setTransformWithArgs", args...)
}
func (o OffscreenCanvasRenderingContext2D) GetShadowBlur() float64 {
	val := o.Get("shadowBlur")
	return val.Float()
}
func (o OffscreenCanvasRenderingContext2D) SetShadowBlur(val float64) {
	o.Set("shadowBlur", val)
}
func (o OffscreenCanvasRenderingContext2D) GetShadowColor() string {
	val := o.Get("shadowColor")
	return val.String()
}
func (o OffscreenCanvasRenderingContext2D) SetShadowColor(val string) {
	o.Set("shadowColor", val)
}
func (o OffscreenCanvasRenderingContext2D) GetShadowOffsetX() float64 {
	val := o.Get("shadowOffsetX")
	return val.Float()
}
func (o OffscreenCanvasRenderingContext2D) SetShadowOffsetX(val float64) {
	o.Set("shadowOffsetX", val)
}
func (o OffscreenCanvasRenderingContext2D) GetShadowOffsetY() float64 {
	val := o.Get("shadowOffsetY")
	return val.Float()
}
func (o OffscreenCanvasRenderingContext2D) SetShadowOffsetY(val float64) {
	o.Set("shadowOffsetY", val)
}
func (o OffscreenCanvasRenderingContext2D) Stroke(args ...interface{}) {
	o.Call("stroke", args...)
}
func (o OffscreenCanvasRenderingContext2D) StrokeRect(args ...interface{}) {
	o.Call("strokeRect", args...)
}
func (o OffscreenCanvasRenderingContext2D) GetStrokeStyle() Value {
	val := o.Get("strokeStyle")
	return val
}
func (o OffscreenCanvasRenderingContext2D) SetStrokeStyle(val Value) {
	o.Set("strokeStyle", val)
}
func (o OffscreenCanvasRenderingContext2D) StrokeText(args ...interface{}) {
	o.Call("strokeText", args...)
}
func (o OffscreenCanvasRenderingContext2D) StrokeWithArgs(args ...interface{}) {
	o.Call("strokeWithArgs", args...)
}
func (o OffscreenCanvasRenderingContext2D) GetTextAlign() CanvasTextAlign {
	val := o.Get("textAlign")
	return JSValueToCanvasTextAlign(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) SetTextAlign(val CanvasTextAlign) {
	o.Set("textAlign", val)
}
func (o OffscreenCanvasRenderingContext2D) GetTextBaseline() CanvasTextBaseline {
	val := o.Get("textBaseline")
	return JSValueToCanvasTextBaseline(val.JSValue())
}
func (o OffscreenCanvasRenderingContext2D) SetTextBaseline(val CanvasTextBaseline) {
	o.Set("textBaseline", val)
}
func (o OffscreenCanvasRenderingContext2D) Transform(args ...interface{}) {
	o.Call("transform", args...)
}
func (o OffscreenCanvasRenderingContext2D) Translate(args ...interface{}) {
	o.Call("translate", args...)
}
