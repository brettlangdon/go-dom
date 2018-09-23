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
	EventTarget
}

func JSValueToSharedWorker(val js.Value) SharedWorker { return SharedWorker{Value: Value{Value: val}} }
func (v Value) AsSharedWorker() SharedWorker          { return SharedWorker{Value: v} }
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
