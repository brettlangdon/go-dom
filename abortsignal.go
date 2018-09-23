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
	EventTarget
}

func jsValueToAbortSignal(val js.Value) AbortSignal { return AbortSignal{Value: Value{Value: val}} }
func (v Value) AsAbortSignal() AbortSignal          { return AbortSignal{Value: v} }
func (a AbortSignal) GetAborted() bool {
	val := a.Get("aborted")
	return val.Bool()
}
func (a AbortSignal) GetOnabort() EventHandler {
	val := a.Get("onabort")
	return jsValueToEventHandler(val.JSValue())
}
func (a AbortSignal) SetOnabort(val EventHandler) {
	a.Set("onabort", val)
}
