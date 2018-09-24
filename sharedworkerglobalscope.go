// Code generated DO NOT EDIT
// sharedworkerglobalscope.go
package dom

import "syscall/js"

type SharedWorkerGlobalScopeIFace interface {
	AddEventListener(args ...interface{})
	Close(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	ImportScripts(args ...interface{})
	GetLocation() WorkerLocation
	GetName() string
	GetNavigator() WorkerNavigator
	GetOnconnect() EventHandler
	SetOnconnect(EventHandler)
	GetOnerror() OnErrorEventHandler
	SetOnerror(OnErrorEventHandler)
	GetOnlanguagechange() EventHandler
	SetOnlanguagechange(EventHandler)
	GetOnoffline() EventHandler
	SetOnoffline(EventHandler)
	GetOnonline() EventHandler
	SetOnonline(EventHandler)
	GetOnrejectionhandled() EventHandler
	SetOnrejectionhandled(EventHandler)
	GetOnunhandledrejection() EventHandler
	SetOnunhandledrejection(EventHandler)
	RemoveEventListener(args ...interface{})
	GetSelf() WorkerGlobalScope
}
type SharedWorkerGlobalScope struct {
	Value
	WorkerGlobalScope
	EventTarget
}

func JSValueToSharedWorkerGlobalScope(val js.Value) SharedWorkerGlobalScope {
	return SharedWorkerGlobalScope{Value: JSValueToValue(val)}
}
func (v Value) AsSharedWorkerGlobalScope() SharedWorkerGlobalScope {
	return SharedWorkerGlobalScope{Value: v}
}
func NewSharedWorkerGlobalScope(args ...interface{}) SharedWorkerGlobalScope {
	return SharedWorkerGlobalScope{Value: JSValueToValue(js.Global().Get("SharedWorkerGlobalScope").New(args...))}
}
func (s SharedWorkerGlobalScope) Close(args ...interface{}) {
	s.Call("close", args...)
}
func (s SharedWorkerGlobalScope) GetName() string {
	val := s.Get("name")
	return val.String()
}
func (s SharedWorkerGlobalScope) GetOnconnect() EventHandler {
	val := s.Get("onconnect")
	return JSValueToEventHandler(val.JSValue())
}
func (s SharedWorkerGlobalScope) SetOnconnect(val EventHandler) {
	s.Set("onconnect", val)
}
