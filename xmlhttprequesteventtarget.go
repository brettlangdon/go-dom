// Code generated DO NOT EDIT
// xmlhttprequesteventtarget.go
package dom

import "syscall/js"

type XMLHttpRequestEventTargetIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetOnabort() EventHandler
	SetOnabort(EventHandler)
	GetOnerror() EventHandler
	SetOnerror(EventHandler)
	GetOnload() EventHandler
	SetOnload(EventHandler)
	GetOnloadend() EventHandler
	SetOnloadend(EventHandler)
	GetOnloadstart() EventHandler
	SetOnloadstart(EventHandler)
	GetOnprogress() EventHandler
	SetOnprogress(EventHandler)
	GetOntimeout() EventHandler
	SetOntimeout(EventHandler)
	RemoveEventListener(args ...interface{})
}
type XMLHttpRequestEventTarget struct {
	Value
}

func JSValueToXMLHttpRequestEventTarget(val js.Value) XMLHttpRequestEventTarget {
	return XMLHttpRequestEventTarget{Value: JSValueToValue(val)}
}
func (v Value) AsXMLHttpRequestEventTarget() XMLHttpRequestEventTarget {
	return XMLHttpRequestEventTarget{Value: v}
}
func NewXMLHttpRequestEventTarget(args ...interface{}) XMLHttpRequestEventTarget {
	return XMLHttpRequestEventTarget{Value: JSValueToValue(js.Global().Get("XMLHttpRequestEventTarget").New(args...))}
}
func (x XMLHttpRequestEventTarget) AddEventListener(args ...interface{}) {
	x.Call("addEventListener", args...)
}
func (x XMLHttpRequestEventTarget) DispatchEvent(args ...interface{}) bool {
	val := x.Call("dispatchEvent", args...)
	return val.Bool()
}
func (x XMLHttpRequestEventTarget) GetOnabort() EventHandler {
	val := x.Get("onabort")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestEventTarget) SetOnabort(val EventHandler) {
	x.Set("onabort", val)
}
func (x XMLHttpRequestEventTarget) GetOnerror() EventHandler {
	val := x.Get("onerror")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestEventTarget) SetOnerror(val EventHandler) {
	x.Set("onerror", val)
}
func (x XMLHttpRequestEventTarget) GetOnload() EventHandler {
	val := x.Get("onload")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestEventTarget) SetOnload(val EventHandler) {
	x.Set("onload", val)
}
func (x XMLHttpRequestEventTarget) GetOnloadend() EventHandler {
	val := x.Get("onloadend")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestEventTarget) SetOnloadend(val EventHandler) {
	x.Set("onloadend", val)
}
func (x XMLHttpRequestEventTarget) GetOnloadstart() EventHandler {
	val := x.Get("onloadstart")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestEventTarget) SetOnloadstart(val EventHandler) {
	x.Set("onloadstart", val)
}
func (x XMLHttpRequestEventTarget) GetOnprogress() EventHandler {
	val := x.Get("onprogress")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestEventTarget) SetOnprogress(val EventHandler) {
	x.Set("onprogress", val)
}
func (x XMLHttpRequestEventTarget) GetOntimeout() EventHandler {
	val := x.Get("ontimeout")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestEventTarget) SetOntimeout(val EventHandler) {
	x.Set("ontimeout", val)
}
func (x XMLHttpRequestEventTarget) RemoveEventListener(args ...interface{}) {
	x.Call("removeEventListener", args...)
}
