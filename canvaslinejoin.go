// Code generated DO NOT EDIT
// canvaslinejoin.go
package dom

import "syscall/js"

type CanvasLineJoin string

const (
	CanvasLineJoinRound CanvasLineJoin = "round"
	CanvasLineJoinBevel CanvasLineJoin = "bevel"
	CanvasLineJoinMiter CanvasLineJoin = "miter"
)

func jsValueToCanvasLineJoin(val js.Value) CanvasLineJoin {
	return CanvasLineJoin(val.String())
}
