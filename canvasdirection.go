// Code generated DO NOT EDIT
// canvasdirection.go
package dom

import "syscall/js"

type CanvasDirection string

const (
	CanvasDirectionLtr     CanvasDirection = "ltr"
	CanvasDirectionRtl     CanvasDirection = "rtl"
	CanvasDirectionInherit CanvasDirection = "inherit"
)

func jsValueToCanvasDirection(val js.Value) CanvasDirection {
	return CanvasDirection(val.String())
}
