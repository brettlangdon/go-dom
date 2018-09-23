// Code generated DO NOT EDIT
// applicationcache.go
package dom

import "syscall/js"

type ApplicationCacheIFace interface {
	Abort(args ...interface{})
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetOncached() EventHandler
	SetOncached(EventHandler)
	GetOnchecking() EventHandler
	SetOnchecking(EventHandler)
	GetOndownloading() EventHandler
	SetOndownloading(EventHandler)
	GetOnerror() EventHandler
	SetOnerror(EventHandler)
	GetOnnoupdate() EventHandler
	SetOnnoupdate(EventHandler)
	GetOnobsolete() EventHandler
	SetOnobsolete(EventHandler)
	GetOnprogress() EventHandler
	SetOnprogress(EventHandler)
	GetOnupdateready() EventHandler
	SetOnupdateready(EventHandler)
	RemoveEventListener(args ...interface{})
	GetStatus() int
	SwapCache(args ...interface{})
	Update(args ...interface{})
}
type ApplicationCache struct {
	Value
	EventTarget
}

func JSValueToApplicationCache(val js.Value) ApplicationCache {
	return ApplicationCache{Value: Value{Value: val}}
}
func (v Value) AsApplicationCache() ApplicationCache { return ApplicationCache{Value: v} }
func (a ApplicationCache) Abort(args ...interface{}) {
	a.Call("abort", args...)
}
func (a ApplicationCache) GetOncached() EventHandler {
	val := a.Get("oncached")
	return JSValueToEventHandler(val.JSValue())
}
func (a ApplicationCache) SetOncached(val EventHandler) {
	a.Set("oncached", val)
}
func (a ApplicationCache) GetOnchecking() EventHandler {
	val := a.Get("onchecking")
	return JSValueToEventHandler(val.JSValue())
}
func (a ApplicationCache) SetOnchecking(val EventHandler) {
	a.Set("onchecking", val)
}
func (a ApplicationCache) GetOndownloading() EventHandler {
	val := a.Get("ondownloading")
	return JSValueToEventHandler(val.JSValue())
}
func (a ApplicationCache) SetOndownloading(val EventHandler) {
	a.Set("ondownloading", val)
}
func (a ApplicationCache) GetOnerror() EventHandler {
	val := a.Get("onerror")
	return JSValueToEventHandler(val.JSValue())
}
func (a ApplicationCache) SetOnerror(val EventHandler) {
	a.Set("onerror", val)
}
func (a ApplicationCache) GetOnnoupdate() EventHandler {
	val := a.Get("onnoupdate")
	return JSValueToEventHandler(val.JSValue())
}
func (a ApplicationCache) SetOnnoupdate(val EventHandler) {
	a.Set("onnoupdate", val)
}
func (a ApplicationCache) GetOnobsolete() EventHandler {
	val := a.Get("onobsolete")
	return JSValueToEventHandler(val.JSValue())
}
func (a ApplicationCache) SetOnobsolete(val EventHandler) {
	a.Set("onobsolete", val)
}
func (a ApplicationCache) GetOnprogress() EventHandler {
	val := a.Get("onprogress")
	return JSValueToEventHandler(val.JSValue())
}
func (a ApplicationCache) SetOnprogress(val EventHandler) {
	a.Set("onprogress", val)
}
func (a ApplicationCache) GetOnupdateready() EventHandler {
	val := a.Get("onupdateready")
	return JSValueToEventHandler(val.JSValue())
}
func (a ApplicationCache) SetOnupdateready(val EventHandler) {
	a.Set("onupdateready", val)
}
func (a ApplicationCache) GetStatus() int {
	val := a.Get("status")
	return val.Int()
}
func (a ApplicationCache) SwapCache(args ...interface{}) {
	a.Call("swapCache", args...)
}
func (a ApplicationCache) Update(args ...interface{}) {
	a.Call("update", args...)
}
