// Code generated DO NOT EDIT
// canvastextbaseline.go
package dom

import "syscall/js"

type CanvasTextBaseline string

const (
	CanvasTextBaselineTop         CanvasTextBaseline = "top"
	CanvasTextBaselineHanging     CanvasTextBaseline = "hanging"
	CanvasTextBaselineMiddle      CanvasTextBaseline = "middle"
	CanvasTextBaselineAlphabetic  CanvasTextBaseline = "alphabetic"
	CanvasTextBaselineIdeographic CanvasTextBaseline = "ideographic"
	CanvasTextBaselineBottom      CanvasTextBaseline = "bottom"
)

func JSValueToCanvasTextBaseline(val js.Value) CanvasTextBaseline {
	return CanvasTextBaseline(val.String())
}
