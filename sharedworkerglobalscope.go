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
func (s SharedWorkerGlobalScope) AddEventListener(args ...interface{}) {
	s.Call("addEventListener", args...)
}
func (s SharedWorkerGlobalScope) Close(args ...interface{}) {
	s.Call("close", args...)
}
func (s SharedWorkerGlobalScope) DispatchEvent(args ...interface{}) bool {
	val := s.Call("dispatchEvent", args...)
	return val.Bool()
}
func (s SharedWorkerGlobalScope) ImportScripts(args ...interface{}) {
	s.Call("importScripts", args...)
}
func (s SharedWorkerGlobalScope) GetLocation() WorkerLocation {
	val := s.Get("location")
	return JSValueToWorkerLocation(val.JSValue())
}
func (s SharedWorkerGlobalScope) GetName() string {
	val := s.Get("name")
	return val.String()
}
func (s SharedWorkerGlobalScope) GetNavigator() WorkerNavigator {
	val := s.Get("navigator")
	return JSValueToWorkerNavigator(val.JSValue())
}
func (s SharedWorkerGlobalScope) GetOnconnect() EventHandler {
	val := s.Get("onconnect")
	return JSValueToEventHandler(val.JSValue())
}
func (s SharedWorkerGlobalScope) SetOnconnect(val EventHandler) {
	s.Set("onconnect", val)
}
func (s SharedWorkerGlobalScope) GetOnerror() OnErrorEventHandler {
	val := s.Get("onerror")
	return JSValueToOnErrorEventHandler(val.JSValue())
}
func (s SharedWorkerGlobalScope) SetOnerror(val OnErrorEventHandler) {
	s.Set("onerror", val)
}
func (s SharedWorkerGlobalScope) GetOnlanguagechange() EventHandler {
	val := s.Get("onlanguagechange")
	return JSValueToEventHandler(val.JSValue())
}
func (s SharedWorkerGlobalScope) SetOnlanguagechange(val EventHandler) {
	s.Set("onlanguagechange", val)
}
func (s SharedWorkerGlobalScope) GetOnoffline() EventHandler {
	val := s.Get("onoffline")
	return JSValueToEventHandler(val.JSValue())
}
func (s SharedWorkerGlobalScope) SetOnoffline(val EventHandler) {
	s.Set("onoffline", val)
}
func (s SharedWorkerGlobalScope) GetOnonline() EventHandler {
	val := s.Get("ononline")
	return JSValueToEventHandler(val.JSValue())
}
func (s SharedWorkerGlobalScope) SetOnonline(val EventHandler) {
	s.Set("ononline", val)
}
func (s SharedWorkerGlobalScope) GetOnrejectionhandled() EventHandler {
	val := s.Get("onrejectionhandled")
	return JSValueToEventHandler(val.JSValue())
}
func (s SharedWorkerGlobalScope) SetOnrejectionhandled(val EventHandler) {
	s.Set("onrejectionhandled", val)
}
func (s SharedWorkerGlobalScope) GetOnunhandledrejection() EventHandler {
	val := s.Get("onunhandledrejection")
	return JSValueToEventHandler(val.JSValue())
}
func (s SharedWorkerGlobalScope) SetOnunhandledrejection(val EventHandler) {
	s.Set("onunhandledrejection", val)
}
func (s SharedWorkerGlobalScope) RemoveEventListener(args ...interface{}) {
	s.Call("removeEventListener", args...)
}
func (s SharedWorkerGlobalScope) GetSelf() WorkerGlobalScope {
	val := s.Get("self")
	return JSValueToWorkerGlobalScope(val.JSValue())
}
