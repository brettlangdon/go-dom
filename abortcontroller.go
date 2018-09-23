// Code generated DO NOT EDIT
// abortcontroller.go
package dom

import "syscall/js"

type AbortControllerIFace interface {
	Abort(args ...interface{})
	GetSignal() AbortSignal
}
type AbortController struct {
	Value
}

func jsValueToAbortController(val js.Value) AbortController {
	return AbortController{Value: Value{Value: val}}
}
func (v Value) AsAbortController() AbortController { return AbortController{Value: v} }
func (a AbortController) Abort(args ...interface{}) {
	a.Call("abort", args...)
}
func (a AbortController) GetSignal() AbortSignal {
	val := a.Get("signal")
	return jsValueToAbortSignal(val.JSValue())
}
