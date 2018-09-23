// Code generated DO NOT EDIT
// eventsource.go
package dom

import "syscall/js"

type EventSourceIFace interface {
	AddEventListener(args ...interface{})
	Close(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetOnerror() EventHandler
	SetOnerror(EventHandler)
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnopen() EventHandler
	SetOnopen(EventHandler)
	GetReadyState() int
	RemoveEventListener(args ...interface{})
	GetUrl() string
	GetWithCredentials() bool
}
type EventSource struct {
	Value
	EventTarget
}

func JSValueToEventSource(val js.Value) EventSource { return EventSource{Value: Value{Value: val}} }
func (v Value) AsEventSource() EventSource          { return EventSource{Value: v} }
func (e EventSource) Close(args ...interface{}) {
	e.Call("close", args...)
}
func (e EventSource) GetOnerror() EventHandler {
	val := e.Get("onerror")
	return JSValueToEventHandler(val.JSValue())
}
func (e EventSource) SetOnerror(val EventHandler) {
	e.Set("onerror", val)
}
func (e EventSource) GetOnmessage() EventHandler {
	val := e.Get("onmessage")
	return JSValueToEventHandler(val.JSValue())
}
func (e EventSource) SetOnmessage(val EventHandler) {
	e.Set("onmessage", val)
}
func (e EventSource) GetOnopen() EventHandler {
	val := e.Get("onopen")
	return JSValueToEventHandler(val.JSValue())
}
func (e EventSource) SetOnopen(val EventHandler) {
	e.Set("onopen", val)
}
func (e EventSource) GetReadyState() int {
	val := e.Get("readyState")
	return val.Int()
}
func (e EventSource) GetUrl() string {
	val := e.Get("url")
	return val.String()
}
func (e EventSource) GetWithCredentials() bool {
	val := e.Get("withCredentials")
	return val.Bool()
}
