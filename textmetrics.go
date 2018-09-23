// Code generated DO NOT EDIT
// textmetrics.go
package dom

import "syscall/js"

type TextMetricsIFace interface {
	GetActualBoundingBoxAscent() float64
	GetActualBoundingBoxDescent() float64
	GetActualBoundingBoxLeft() float64
	GetActualBoundingBoxRight() float64
	GetAlphabeticBaseline() float64
	GetEmHeightAscent() float64
	GetEmHeightDescent() float64
	GetFontBoundingBoxAscent() float64
	GetFontBoundingBoxDescent() float64
	GetHangingBaseline() float64
	GetIdeographicBaseline() float64
	GetWidth() float64
}
type TextMetrics struct {
	Value
}

func JSValueToTextMetrics(val js.Value) TextMetrics { return TextMetrics{Value: Value{Value: val}} }
func (v Value) AsTextMetrics() TextMetrics          { return TextMetrics{Value: v} }
func (t TextMetrics) GetActualBoundingBoxAscent() float64 {
	val := t.Get("actualBoundingBoxAscent")
	return val.Float()
}
func (t TextMetrics) GetActualBoundingBoxDescent() float64 {
	val := t.Get("actualBoundingBoxDescent")
	return val.Float()
}
func (t TextMetrics) GetActualBoundingBoxLeft() float64 {
	val := t.Get("actualBoundingBoxLeft")
	return val.Float()
}
func (t TextMetrics) GetActualBoundingBoxRight() float64 {
	val := t.Get("actualBoundingBoxRight")
	return val.Float()
}
func (t TextMetrics) GetAlphabeticBaseline() float64 {
	val := t.Get("alphabeticBaseline")
	return val.Float()
}
func (t TextMetrics) GetEmHeightAscent() float64 {
	val := t.Get("emHeightAscent")
	return val.Float()
}
func (t TextMetrics) GetEmHeightDescent() float64 {
	val := t.Get("emHeightDescent")
	return val.Float()
}
func (t TextMetrics) GetFontBoundingBoxAscent() float64 {
	val := t.Get("fontBoundingBoxAscent")
	return val.Float()
}
func (t TextMetrics) GetFontBoundingBoxDescent() float64 {
	val := t.Get("fontBoundingBoxDescent")
	return val.Float()
}
func (t TextMetrics) GetHangingBaseline() float64 {
	val := t.Get("hangingBaseline")
	return val.Float()
}
func (t TextMetrics) GetIdeographicBaseline() float64 {
	val := t.Get("ideographicBaseline")
	return val.Float()
}
func (t TextMetrics) GetWidth() float64 {
	val := t.Get("width")
	return val.Float()
}
