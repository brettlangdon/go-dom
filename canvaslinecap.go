// Code generated DO NOT EDIT
// canvaslinecap.go
package dom

import "syscall/js"

type CanvasLineCap string

const (
	CanvasLineCapButt   CanvasLineCap = "butt"
	CanvasLineCapRound  CanvasLineCap = "round"
	CanvasLineCapSquare CanvasLineCap = "square"
)

func jsValueToCanvasLineCap(val js.Value) CanvasLineCap {
	return CanvasLineCap(val.String())
}
