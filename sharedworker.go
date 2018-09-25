// Code generated DO NOT EDIT
// sharedworker.go
package dom

import "syscall/js"

type SharedWorkerIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetOnerror() EventHandler
	SetOnerror(EventHandler)
	GetPort() MessagePort
	RemoveEventListener(args ...interface{})
}
type SharedWorker struct {
	Value
}

func JSValueToSharedWorker(val js.Value) SharedWorker { return SharedWorker{Value: JSValueToValue(val)} }
func (v Value) AsSharedWorker() SharedWorker          { return SharedWorker{Value: v} }
func NewSharedWorker(args ...interface{}) SharedWorker {
	return SharedWorker{Value: JSValueToValue(js.Global().Get("SharedWorker").New(args...))}
}
func (s SharedWorker) AddEventListener(args ...interface{}) {
	s.Call("addEventListener", args...)
}
func (s SharedWorker) DispatchEvent(args ...interface{}) bool {
	val := s.Call("dispatchEvent", args...)
	return val.Bool()
}
func (s SharedWorker) GetOnerror() EventHandler {
	val := s.Get("onerror")
	return JSValueToEventHandler(val.JSValue())
}
func (s SharedWorker) SetOnerror(val EventHandler) {
	s.Set("onerror", val)
}
func (s SharedWorker) GetPort() MessagePort {
	val := s.Get("port")
	return JSValueToMessagePort(val.JSValue())
}
func (s SharedWorker) RemoveEventListener(args ...interface{}) {
	s.Call("removeEventListener", args...)
}
