// Code generated DO NOT EDIT
// dedicatedworkerglobalscope.go
package dom

import "syscall/js"

type DedicatedWorkerGlobalScopeIFace interface {
	AddEventListener(args ...interface{})
	CancelAnimationFrame(args ...interface{})
	Close(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	ImportScripts(args ...interface{})
	GetLocation() WorkerLocation
	GetName() string
	GetNavigator() WorkerNavigator
	GetOnerror() OnErrorEventHandler
	SetOnerror(OnErrorEventHandler)
	GetOnlanguagechange() EventHandler
	SetOnlanguagechange(EventHandler)
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnmessageerror() EventHandler
	SetOnmessageerror(EventHandler)
	GetOnoffline() EventHandler
	SetOnoffline(EventHandler)
	GetOnonline() EventHandler
	SetOnonline(EventHandler)
	GetOnrejectionhandled() EventHandler
	SetOnrejectionhandled(EventHandler)
	GetOnunhandledrejection() EventHandler
	SetOnunhandledrejection(EventHandler)
	PostMessage(args ...interface{})
	RemoveEventListener(args ...interface{})
	RequestAnimationFrame(args ...interface{}) int
	GetSelf() WorkerGlobalScope
}
type DedicatedWorkerGlobalScope struct {
	Value
}

func JSValueToDedicatedWorkerGlobalScope(val js.Value) DedicatedWorkerGlobalScope {
	return DedicatedWorkerGlobalScope{Value: JSValueToValue(val)}
}
func (v Value) AsDedicatedWorkerGlobalScope() DedicatedWorkerGlobalScope {
	return DedicatedWorkerGlobalScope{Value: v}
}
func NewDedicatedWorkerGlobalScope(args ...interface{}) DedicatedWorkerGlobalScope {
	return DedicatedWorkerGlobalScope{Value: JSValueToValue(js.Global().Get("DedicatedWorkerGlobalScope").New(args...))}
}
func (d DedicatedWorkerGlobalScope) AddEventListener(args ...interface{}) {
	d.Call("addEventListener", args...)
}
func (d DedicatedWorkerGlobalScope) CancelAnimationFrame(args ...interface{}) {
	d.Call("cancelAnimationFrame", args...)
}
func (d DedicatedWorkerGlobalScope) Close(args ...interface{}) {
	d.Call("close", args...)
}
func (d DedicatedWorkerGlobalScope) DispatchEvent(args ...interface{}) bool {
	val := d.Call("dispatchEvent", args...)
	return val.Bool()
}
func (d DedicatedWorkerGlobalScope) ImportScripts(args ...interface{}) {
	d.Call("importScripts", args...)
}
func (d DedicatedWorkerGlobalScope) GetLocation() WorkerLocation {
	val := d.Get("location")
	return JSValueToWorkerLocation(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) GetName() string {
	val := d.Get("name")
	return val.String()
}
func (d DedicatedWorkerGlobalScope) GetNavigator() WorkerNavigator {
	val := d.Get("navigator")
	return JSValueToWorkerNavigator(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) GetOnerror() OnErrorEventHandler {
	val := d.Get("onerror")
	return JSValueToOnErrorEventHandler(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) SetOnerror(val OnErrorEventHandler) {
	d.Set("onerror", val)
}
func (d DedicatedWorkerGlobalScope) GetOnlanguagechange() EventHandler {
	val := d.Get("onlanguagechange")
	return JSValueToEventHandler(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) SetOnlanguagechange(val EventHandler) {
	d.Set("onlanguagechange", val)
}
func (d DedicatedWorkerGlobalScope) GetOnmessage() EventHandler {
	val := d.Get("onmessage")
	return JSValueToEventHandler(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) SetOnmessage(val EventHandler) {
	d.Set("onmessage", val)
}
func (d DedicatedWorkerGlobalScope) GetOnmessageerror() EventHandler {
	val := d.Get("onmessageerror")
	return JSValueToEventHandler(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) SetOnmessageerror(val EventHandler) {
	d.Set("onmessageerror", val)
}
func (d DedicatedWorkerGlobalScope) GetOnoffline() EventHandler {
	val := d.Get("onoffline")
	return JSValueToEventHandler(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) SetOnoffline(val EventHandler) {
	d.Set("onoffline", val)
}
func (d DedicatedWorkerGlobalScope) GetOnonline() EventHandler {
	val := d.Get("ononline")
	return JSValueToEventHandler(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) SetOnonline(val EventHandler) {
	d.Set("ononline", val)
}
func (d DedicatedWorkerGlobalScope) GetOnrejectionhandled() EventHandler {
	val := d.Get("onrejectionhandled")
	return JSValueToEventHandler(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) SetOnrejectionhandled(val EventHandler) {
	d.Set("onrejectionhandled", val)
}
func (d DedicatedWorkerGlobalScope) GetOnunhandledrejection() EventHandler {
	val := d.Get("onunhandledrejection")
	return JSValueToEventHandler(val.JSValue())
}
func (d DedicatedWorkerGlobalScope) SetOnunhandledrejection(val EventHandler) {
	d.Set("onunhandledrejection", val)
}
func (d DedicatedWorkerGlobalScope) PostMessage(args ...interface{}) {
	d.Call("postMessage", args...)
}
func (d DedicatedWorkerGlobalScope) RemoveEventListener(args ...interface{}) {
	d.Call("removeEventListener", args...)
}
func (d DedicatedWorkerGlobalScope) RequestAnimationFrame(args ...interface{}) int {
	val := d.Call("requestAnimationFrame", args...)
	return val.Int()
}
func (d DedicatedWorkerGlobalScope) GetSelf() WorkerGlobalScope {
	val := d.Get("self")
	return JSValueToWorkerGlobalScope(val.JSValue())
}
