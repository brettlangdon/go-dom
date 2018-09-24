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

func JSValueToAbortController(val js.Value) AbortController {
	return AbortController{Value: JSValueToValue(val)}
}
func (v Value) AsAbortController() AbortController { return AbortController{Value: v} }
func NewAbortController(args ...interface{}) AbortController {
	return AbortController{Value: JSValueToValue(js.Global().Get("AbortController").New(args...))}
}
func (a AbortController) Abort(args ...interface{}) {
	a.Call("abort", args...)
}
func (a AbortController) GetSignal() AbortSignal {
	val := a.Get("signal")
	return JSValueToAbortSignal(val.JSValue())
}
