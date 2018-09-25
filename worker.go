// Code generated DO NOT EDIT
// worker.go
package dom

import "syscall/js"

type WorkerIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetOnerror() EventHandler
	SetOnerror(EventHandler)
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnmessageerror() EventHandler
	SetOnmessageerror(EventHandler)
	PostMessage(args ...interface{})
	RemoveEventListener(args ...interface{})
	Terminate(args ...interface{})
}
type Worker struct {
	Value
}

func JSValueToWorker(val js.Value) Worker { return Worker{Value: JSValueToValue(val)} }
func (v Value) AsWorker() Worker          { return Worker{Value: v} }
func NewWorker(args ...interface{}) Worker {
	return Worker{Value: JSValueToValue(js.Global().Get("Worker").New(args...))}
}
func (w Worker) AddEventListener(args ...interface{}) {
	w.Call("addEventListener", args...)
}
func (w Worker) DispatchEvent(args ...interface{}) bool {
	val := w.Call("dispatchEvent", args...)
	return val.Bool()
}
func (w Worker) GetOnerror() EventHandler {
	val := w.Get("onerror")
	return JSValueToEventHandler(val.JSValue())
}
func (w Worker) SetOnerror(val EventHandler) {
	w.Set("onerror", val)
}
func (w Worker) GetOnmessage() EventHandler {
	val := w.Get("onmessage")
	return JSValueToEventHandler(val.JSValue())
}
func (w Worker) SetOnmessage(val EventHandler) {
	w.Set("onmessage", val)
}
func (w Worker) GetOnmessageerror() EventHandler {
	val := w.Get("onmessageerror")
	return JSValueToEventHandler(val.JSValue())
}
func (w Worker) SetOnmessageerror(val EventHandler) {
	w.Set("onmessageerror", val)
}
func (w Worker) PostMessage(args ...interface{}) {
	w.Call("postMessage", args...)
}
func (w Worker) RemoveEventListener(args ...interface{}) {
	w.Call("removeEventListener", args...)
}
func (w Worker) Terminate(args ...interface{}) {
	w.Call("terminate", args...)
}
