// Code generated DO NOT EDIT
// workerglobalscope.go
package dom

import "syscall/js"

type WorkerGlobalScopeIFace interface {
	AddEventListener(args ...interface{})
	Atob(args ...interface{}) []byte
	Btoa(args ...interface{}) string
	ClearInterval(args ...interface{})
	ClearTimeout(args ...interface{})
	CreateImageBitmap(args ...interface{})
	CreateImageBitmapWithArgs(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	ImportScripts(args ...interface{})
	GetLocation() WorkerLocation
	GetNavigator() WorkerNavigator
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
	GetOrigin() string
	QueueMicrotask(args ...interface{})
	RemoveEventListener(args ...interface{})
	GetSelf() WorkerGlobalScope
	SetInterval(args ...interface{}) int
	SetTimeout(args ...interface{}) int
}
type WorkerGlobalScope struct {
	Value
	EventTarget
}

func JSValueToWorkerGlobalScope(val js.Value) WorkerGlobalScope {
	return WorkerGlobalScope{Value: Value{Value: val}}
}
func (v Value) AsWorkerGlobalScope() WorkerGlobalScope { return WorkerGlobalScope{Value: v} }
func (w WorkerGlobalScope) Atob(args ...interface{}) []byte {
	val := w.Call("atob", args...)
	return []byte(val.String())
}
func (w WorkerGlobalScope) Btoa(args ...interface{}) string {
	val := w.Call("btoa", args...)
	return val.String()
}
func (w WorkerGlobalScope) ClearInterval(args ...interface{}) {
	w.Call("clearInterval", args...)
}
func (w WorkerGlobalScope) ClearTimeout(args ...interface{}) {
	w.Call("clearTimeout", args...)
}
func (w WorkerGlobalScope) CreateImageBitmap(args ...interface{}) {
	w.Call("createImageBitmap", args...)
}
func (w WorkerGlobalScope) CreateImageBitmapWithArgs(args ...interface{}) {
	w.Call("createImageBitmapWithArgs", args...)
}
func (w WorkerGlobalScope) ImportScripts(args ...interface{}) {
	w.Call("importScripts", args...)
}
func (w WorkerGlobalScope) GetLocation() WorkerLocation {
	val := w.Get("location")
	return JSValueToWorkerLocation(val.JSValue())
}
func (w WorkerGlobalScope) GetNavigator() WorkerNavigator {
	val := w.Get("navigator")
	return JSValueToWorkerNavigator(val.JSValue())
}
func (w WorkerGlobalScope) GetOnerror() OnErrorEventHandler {
	val := w.Get("onerror")
	return JSValueToOnErrorEventHandler(val.JSValue())
}
func (w WorkerGlobalScope) SetOnerror(val OnErrorEventHandler) {
	w.Set("onerror", val)
}
func (w WorkerGlobalScope) GetOnlanguagechange() EventHandler {
	val := w.Get("onlanguagechange")
	return JSValueToEventHandler(val.JSValue())
}
func (w WorkerGlobalScope) SetOnlanguagechange(val EventHandler) {
	w.Set("onlanguagechange", val)
}
func (w WorkerGlobalScope) GetOnoffline() EventHandler {
	val := w.Get("onoffline")
	return JSValueToEventHandler(val.JSValue())
}
func (w WorkerGlobalScope) SetOnoffline(val EventHandler) {
	w.Set("onoffline", val)
}
func (w WorkerGlobalScope) GetOnonline() EventHandler {
	val := w.Get("ononline")
	return JSValueToEventHandler(val.JSValue())
}
func (w WorkerGlobalScope) SetOnonline(val EventHandler) {
	w.Set("ononline", val)
}
func (w WorkerGlobalScope) GetOnrejectionhandled() EventHandler {
	val := w.Get("onrejectionhandled")
	return JSValueToEventHandler(val.JSValue())
}
func (w WorkerGlobalScope) SetOnrejectionhandled(val EventHandler) {
	w.Set("onrejectionhandled", val)
}
func (w WorkerGlobalScope) GetOnunhandledrejection() EventHandler {
	val := w.Get("onunhandledrejection")
	return JSValueToEventHandler(val.JSValue())
}
func (w WorkerGlobalScope) SetOnunhandledrejection(val EventHandler) {
	w.Set("onunhandledrejection", val)
}
func (w WorkerGlobalScope) GetOrigin() string {
	val := w.Get("origin")
	return val.String()
}
func (w WorkerGlobalScope) QueueMicrotask(args ...interface{}) {
	w.Call("queueMicrotask", args...)
}
func (w WorkerGlobalScope) GetSelf() WorkerGlobalScope {
	val := w.Get("self")
	return JSValueToWorkerGlobalScope(val.JSValue())
}
func (w WorkerGlobalScope) SetInterval(args ...interface{}) int {
	val := w.Call("setInterval", args...)
	return val.Int()
}
func (w WorkerGlobalScope) SetTimeout(args ...interface{}) int {
	val := w.Call("setTimeout", args...)
	return val.Int()
}
