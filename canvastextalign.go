// Code generated DO NOT EDIT
// canvastextalign.go
package dom

import "syscall/js"

type CanvasTextAlign string

const (
	CanvasTextAlignStart  CanvasTextAlign = "start"
	CanvasTextAlignEnd    CanvasTextAlign = "end"
	CanvasTextAlignLeft   CanvasTextAlign = "left"
	CanvasTextAlignRight  CanvasTextAlign = "right"
	CanvasTextAlignCenter CanvasTextAlign = "center"
)

func jsValueToCanvasTextAlign(val js.Value) CanvasTextAlign {
	return CanvasTextAlign(val.String())
}
