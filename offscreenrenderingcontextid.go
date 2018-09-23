// Code generated DO NOT EDIT
// offscreenrenderingcontextid.go
package dom

import "syscall/js"

type OffscreenRenderingContextId string

const (
	OffscreenRenderingContextId2d     OffscreenRenderingContextId = "2d"
	OffscreenRenderingContextIdWebgl  OffscreenRenderingContextId = "webgl"
	OffscreenRenderingContextIdWebgl2 OffscreenRenderingContextId = "webgl2"
)

func jsValueToOffscreenRenderingContextId(val js.Value) OffscreenRenderingContextId {
	return OffscreenRenderingContextId(val.String())
}
