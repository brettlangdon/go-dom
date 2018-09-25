// Code generated DO NOT EDIT
// abortsignal.go
package dom

import "syscall/js"

type AbortSignalIFace interface {
	GetAborted() bool
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetOnabort() EventHandler
	SetOnabort(EventHandler)
	RemoveEventListener(args ...interface{})
}
type AbortSignal struct {
	Value
}

func JSValueToAbortSignal(val js.Value) AbortSignal { return AbortSignal{Value: JSValueToValue(val)} }
func (v Value) AsAbortSignal() AbortSignal          { return AbortSignal{Value: v} }
func NewAbortSignal(args ...interface{}) AbortSignal {
	return AbortSignal{Value: JSValueToValue(js.Global().Get("AbortSignal").New(args...))}
}
func (a AbortSignal) GetAborted() bool {
	val := a.Get("aborted")
	return val.Bool()
}
func (a AbortSignal) AddEventListener(args ...interface{}) {
	a.Call("addEventListener", args...)
}
func (a AbortSignal) DispatchEvent(args ...interface{}) bool {
	val := a.Call("dispatchEvent", args...)
	return val.Bool()
}
func (a AbortSignal) GetOnabort() EventHandler {
	val := a.Get("onabort")
	return JSValueToEventHandler(val.JSValue())
}
func (a AbortSignal) SetOnabort(val EventHandler) {
	a.Set("onabort", val)
}
func (a AbortSignal) RemoveEventListener(args ...interface{}) {
	a.Call("removeEventListener", args...)
}
