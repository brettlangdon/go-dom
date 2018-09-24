// Code generated DO NOT EDIT
// canvasrenderingcontext2d.go
package dom

import "syscall/js"

type CanvasRenderingContext2DIFace interface {
	Arc(args ...interface{})
	ArcTo(args ...interface{})
	BeginPath(args ...interface{})
	BezierCurveTo(args ...interface{})
	GetCanvas() HTMLCanvasElement
	ClearRect(args ...interface{})
	Clip(args ...interface{})
	ClipWithArgs(args ...interface{})
	ClosePath(args ...interface{})
	CreateImageData(args ...interface{}) ImageData
	CreateImageDataWithArgs(args ...interface{}) ImageData
	CreateLinearGradient(args ...interface{}) CanvasGradient
	CreatePattern(args ...interface{}) CanvasPattern
	CreateRadialGradient(args ...interface{}) CanvasGradient
	GetDirection() CanvasDirection
	SetDirection(CanvasDirection)
	DrawFocusIfNeeded(args ...interface{})
	DrawFocusIfNeededWithArgs(args ...interface{})
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
	ScrollPathIntoView(args ...interface{})
	ScrollPathIntoViewWithArgs(args ...interface{})
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
type CanvasRenderingContext2D struct {
	Value
}

func JSValueToCanvasRenderingContext2D(val js.Value) CanvasRenderingContext2D {
	return CanvasRenderingContext2D{Value: JSValueToValue(val)}
}
func (v Value) AsCanvasRenderingContext2D() CanvasRenderingContext2D {
	return CanvasRenderingContext2D{Value: v}
}
func NewCanvasRenderingContext2D(args ...interface{}) CanvasRenderingContext2D {
	return CanvasRenderingContext2D{Value: JSValueToValue(js.Global().Get("CanvasRenderingContext2D").New(args...))}
}
func (c CanvasRenderingContext2D) Arc(args ...interface{}) {
	c.Call("arc", args...)
}
func (c CanvasRenderingContext2D) ArcTo(args ...interface{}) {
	c.Call("arcTo", args...)
}
func (c CanvasRenderingContext2D) BeginPath(args ...interface{}) {
	c.Call("beginPath", args...)
}
func (c CanvasRenderingContext2D) BezierCurveTo(args ...interface{}) {
	c.Call("bezierCurveTo", args...)
}
func (c CanvasRenderingContext2D) GetCanvas() HTMLCanvasElement {
	val := c.Get("canvas")
	return JSValueToHTMLCanvasElement(val.JSValue())
}
func (c CanvasRenderingContext2D) ClearRect(args ...interface{}) {
	c.Call("clearRect", args...)
}
func (c CanvasRenderingContext2D) Clip(args ...interface{}) {
	c.Call("clip", args...)
}
func (c CanvasRenderingContext2D) ClipWithArgs(args ...interface{}) {
	c.Call("clipWithArgs", args...)
}
func (c CanvasRenderingContext2D) ClosePath(args ...interface{}) {
	c.Call("closePath", args...)
}
func (c CanvasRenderingContext2D) CreateImageData(args ...interface{}) ImageData {
	val := c.Call("createImageData", args...)
	return JSValueToImageData(val.JSValue())
}
func (c CanvasRenderingContext2D) CreateImageDataWithArgs(args ...interface{}) ImageData {
	val := c.Call("createImageDataWithArgs", args...)
	return JSValueToImageData(val.JSValue())
}
func (c CanvasRenderingContext2D) CreateLinearGradient(args ...interface{}) CanvasGradient {
	val := c.Call("createLinearGradient", args...)
	return JSValueToCanvasGradient(val.JSValue())
}
func (c CanvasRenderingContext2D) CreatePattern(args ...interface{}) CanvasPattern {
	val := c.Call("createPattern", args...)
	return JSValueToCanvasPattern(val.JSValue())
}
func (c CanvasRenderingContext2D) CreateRadialGradient(args ...interface{}) CanvasGradient {
	val := c.Call("createRadialGradient", args...)
	return JSValueToCanvasGradient(val.JSValue())
}
func (c CanvasRenderingContext2D) GetDirection() CanvasDirection {
	val := c.Get("direction")
	return JSValueToCanvasDirection(val.JSValue())
}
func (c CanvasRenderingContext2D) SetDirection(val CanvasDirection) {
	c.Set("direction", val)
}
func (c CanvasRenderingContext2D) DrawFocusIfNeeded(args ...interface{}) {
	c.Call("drawFocusIfNeeded", args...)
}
func (c CanvasRenderingContext2D) DrawFocusIfNeededWithArgs(args ...interface{}) {
	c.Call("drawFocusIfNeededWithArgs", args...)
}
func (c CanvasRenderingContext2D) DrawImage(args ...interface{}) {
	c.Call("drawImage", args...)
}
func (c CanvasRenderingContext2D) Ellipse(args ...interface{}) {
	c.Call("ellipse", args...)
}
func (c CanvasRenderingContext2D) Fill(args ...interface{}) {
	c.Call("fill", args...)
}
func (c CanvasRenderingContext2D) FillRect(args ...interface{}) {
	c.Call("fillRect", args...)
}
func (c CanvasRenderingContext2D) GetFillStyle() Value {
	val := c.Get("fillStyle")
	return val
}
func (c CanvasRenderingContext2D) SetFillStyle(val Value) {
	c.Set("fillStyle", val)
}
func (c CanvasRenderingContext2D) FillText(args ...interface{}) {
	c.Call("fillText", args...)
}
func (c CanvasRenderingContext2D) FillWithArgs(args ...interface{}) {
	c.Call("fillWithArgs", args...)
}
func (c CanvasRenderingContext2D) GetFilter() string {
	val := c.Get("filter")
	return val.String()
}
func (c CanvasRenderingContext2D) SetFilter(val string) {
	c.Set("filter", val)
}
func (c CanvasRenderingContext2D) GetFont() string {
	val := c.Get("font")
	return val.String()
}
func (c CanvasRenderingContext2D) SetFont(val string) {
	c.Set("font", val)
}
func (c CanvasRenderingContext2D) GetImageData(args ...interface{}) ImageData {
	val := c.Call("getImageData", args...)
	return JSValueToImageData(val.JSValue())
}
func (c CanvasRenderingContext2D) GetLineDash(args ...interface{}) {
	c.Call("getLineDash", args...)
}
func (c CanvasRenderingContext2D) GetTransform(args ...interface{}) DOMMatrix {
	val := c.Call("getTransform", args...)
	return JSValueToDOMMatrix(val.JSValue())
}
func (c CanvasRenderingContext2D) GetGlobalAlpha() float64 {
	val := c.Get("globalAlpha")
	return val.Float()
}
func (c CanvasRenderingContext2D) SetGlobalAlpha(val float64) {
	c.Set("globalAlpha", val)
}
func (c CanvasRenderingContext2D) GetGlobalCompositeOperation() string {
	val := c.Get("globalCompositeOperation")
	return val.String()
}
func (c CanvasRenderingContext2D) SetGlobalCompositeOperation(val string) {
	c.Set("globalCompositeOperation", val)
}
func (c CanvasRenderingContext2D) GetImageSmoothingEnabled() bool {
	val := c.Get("imageSmoothingEnabled")
	return val.Bool()
}
func (c CanvasRenderingContext2D) SetImageSmoothingEnabled(val bool) {
	c.Set("imageSmoothingEnabled", val)
}
func (c CanvasRenderingContext2D) GetImageSmoothingQuality() ImageSmoothingQuality {
	val := c.Get("imageSmoothingQuality")
	return JSValueToImageSmoothingQuality(val.JSValue())
}
func (c CanvasRenderingContext2D) SetImageSmoothingQuality(val ImageSmoothingQuality) {
	c.Set("imageSmoothingQuality", val)
}
func (c CanvasRenderingContext2D) IsPointInPath(args ...interface{}) bool {
	val := c.Call("isPointInPath", args...)
	return val.Bool()
}
func (c CanvasRenderingContext2D) IsPointInPathWithArgs(args ...interface{}) bool {
	val := c.Call("isPointInPathWithArgs", args...)
	return val.Bool()
}
func (c CanvasRenderingContext2D) IsPointInStroke(args ...interface{}) bool {
	val := c.Call("isPointInStroke", args...)
	return val.Bool()
}
func (c CanvasRenderingContext2D) IsPointInStrokeWithArgs(args ...interface{}) bool {
	val := c.Call("isPointInStrokeWithArgs", args...)
	return val.Bool()
}
func (c CanvasRenderingContext2D) GetLineCap() CanvasLineCap {
	val := c.Get("lineCap")
	return JSValueToCanvasLineCap(val.JSValue())
}
func (c CanvasRenderingContext2D) SetLineCap(val CanvasLineCap) {
	c.Set("lineCap", val)
}
func (c CanvasRenderingContext2D) GetLineDashOffset() float64 {
	val := c.Get("lineDashOffset")
	return val.Float()
}
func (c CanvasRenderingContext2D) SetLineDashOffset(val float64) {
	c.Set("lineDashOffset", val)
}
func (c CanvasRenderingContext2D) GetLineJoin() CanvasLineJoin {
	val := c.Get("lineJoin")
	return JSValueToCanvasLineJoin(val.JSValue())
}
func (c CanvasRenderingContext2D) SetLineJoin(val CanvasLineJoin) {
	c.Set("lineJoin", val)
}
func (c CanvasRenderingContext2D) LineTo(args ...interface{}) {
	c.Call("lineTo", args...)
}
func (c CanvasRenderingContext2D) GetLineWidth() float64 {
	val := c.Get("lineWidth")
	return val.Float()
}
func (c CanvasRenderingContext2D) SetLineWidth(val float64) {
	c.Set("lineWidth", val)
}
func (c CanvasRenderingContext2D) MeasureText(args ...interface{}) TextMetrics {
	val := c.Call("measureText", args...)
	return JSValueToTextMetrics(val.JSValue())
}
func (c CanvasRenderingContext2D) GetMiterLimit() float64 {
	val := c.Get("miterLimit")
	return val.Float()
}
func (c CanvasRenderingContext2D) SetMiterLimit(val float64) {
	c.Set("miterLimit", val)
}
func (c CanvasRenderingContext2D) MoveTo(args ...interface{}) {
	c.Call("moveTo", args...)
}
func (c CanvasRenderingContext2D) PutImageData(args ...interface{}) {
	c.Call("putImageData", args...)
}
func (c CanvasRenderingContext2D) PutImageDataWithArgs(args ...interface{}) {
	c.Call("putImageDataWithArgs", args...)
}
func (c CanvasRenderingContext2D) QuadraticCurveTo(args ...interface{}) {
	c.Call("quadraticCurveTo", args...)
}
func (c CanvasRenderingContext2D) Rect(args ...interface{}) {
	c.Call("rect", args...)
}
func (c CanvasRenderingContext2D) ResetTransform(args ...interface{}) {
	c.Call("resetTransform", args...)
}
func (c CanvasRenderingContext2D) Restore(args ...interface{}) {
	c.Call("restore", args...)
}
func (c CanvasRenderingContext2D) Rotate(args ...interface{}) {
	c.Call("rotate", args...)
}
func (c CanvasRenderingContext2D) Save(args ...interface{}) {
	c.Call("save", args...)
}
func (c CanvasRenderingContext2D) Scale(args ...interface{}) {
	c.Call("scale", args...)
}
func (c CanvasRenderingContext2D) ScrollPathIntoView(args ...interface{}) {
	c.Call("scrollPathIntoView", args...)
}
func (c CanvasRenderingContext2D) ScrollPathIntoViewWithArgs(args ...interface{}) {
	c.Call("scrollPathIntoViewWithArgs", args...)
}
func (c CanvasRenderingContext2D) SetLineDash(args ...interface{}) {
	c.Call("setLineDash", args...)
}
func (c CanvasRenderingContext2D) SetTransform(args ...interface{}) {
	c.Call("setTransform", args...)
}
func (c CanvasRenderingContext2D) SetTransformWithArgs(args ...interface{}) {
	c.Call("setTransformWithArgs", args...)
}
func (c CanvasRenderingContext2D) GetShadowBlur() float64 {
	val := c.Get("shadowBlur")
	return val.Float()
}
func (c CanvasRenderingContext2D) SetShadowBlur(val float64) {
	c.Set("shadowBlur", val)
}
func (c CanvasRenderingContext2D) GetShadowColor() string {
	val := c.Get("shadowColor")
	return val.String()
}
func (c CanvasRenderingContext2D) SetShadowColor(val string) {
	c.Set("shadowColor", val)
}
func (c CanvasRenderingContext2D) GetShadowOffsetX() float64 {
	val := c.Get("shadowOffsetX")
	return val.Float()
}
func (c CanvasRenderingContext2D) SetShadowOffsetX(val float64) {
	c.Set("shadowOffsetX", val)
}
func (c CanvasRenderingContext2D) GetShadowOffsetY() float64 {
	val := c.Get("shadowOffsetY")
	return val.Float()
}
func (c CanvasRenderingContext2D) SetShadowOffsetY(val float64) {
	c.Set("shadowOffsetY", val)
}
func (c CanvasRenderingContext2D) Stroke(args ...interface{}) {
	c.Call("stroke", args...)
}
func (c CanvasRenderingContext2D) StrokeRect(args ...interface{}) {
	c.Call("strokeRect", args...)
}
func (c CanvasRenderingContext2D) GetStrokeStyle() Value {
	val := c.Get("strokeStyle")
	return val
}
func (c CanvasRenderingContext2D) SetStrokeStyle(val Value) {
	c.Set("strokeStyle", val)
}
func (c CanvasRenderingContext2D) StrokeText(args ...interface{}) {
	c.Call("strokeText", args...)
}
func (c CanvasRenderingContext2D) StrokeWithArgs(args ...interface{}) {
	c.Call("strokeWithArgs", args...)
}
func (c CanvasRenderingContext2D) GetTextAlign() CanvasTextAlign {
	val := c.Get("textAlign")
	return JSValueToCanvasTextAlign(val.JSValue())
}
func (c CanvasRenderingContext2D) SetTextAlign(val CanvasTextAlign) {
	c.Set("textAlign", val)
}
func (c CanvasRenderingContext2D) GetTextBaseline() CanvasTextBaseline {
	val := c.Get("textBaseline")
	return JSValueToCanvasTextBaseline(val.JSValue())
}
func (c CanvasRenderingContext2D) SetTextBaseline(val CanvasTextBaseline) {
	c.Set("textBaseline", val)
}
func (c CanvasRenderingContext2D) Transform(args ...interface{}) {
	c.Call("transform", args...)
}
func (c CanvasRenderingContext2D) Translate(args ...interface{}) {
	c.Call("translate", args...)
}
