// Code generated DO NOT EDIT
// xmlhttprequestupload.go
package dom

import "syscall/js"

type XMLHttpRequestUploadIFace interface {
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
type XMLHttpRequestUpload struct {
	Value
}

func JSValueToXMLHttpRequestUpload(val js.Value) XMLHttpRequestUpload {
	return XMLHttpRequestUpload{Value: JSValueToValue(val)}
}
func (v Value) AsXMLHttpRequestUpload() XMLHttpRequestUpload { return XMLHttpRequestUpload{Value: v} }
func NewXMLHttpRequestUpload(args ...interface{}) XMLHttpRequestUpload {
	return XMLHttpRequestUpload{Value: JSValueToValue(js.Global().Get("XMLHttpRequestUpload").New(args...))}
}
func (x XMLHttpRequestUpload) AddEventListener(args ...interface{}) {
	x.Call("addEventListener", args...)
}
func (x XMLHttpRequestUpload) DispatchEvent(args ...interface{}) bool {
	val := x.Call("dispatchEvent", args...)
	return val.Bool()
}
func (x XMLHttpRequestUpload) GetOnabort() EventHandler {
	val := x.Get("onabort")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestUpload) SetOnabort(val EventHandler) {
	x.Set("onabort", val)
}
func (x XMLHttpRequestUpload) GetOnerror() EventHandler {
	val := x.Get("onerror")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestUpload) SetOnerror(val EventHandler) {
	x.Set("onerror", val)
}
func (x XMLHttpRequestUpload) GetOnload() EventHandler {
	val := x.Get("onload")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestUpload) SetOnload(val EventHandler) {
	x.Set("onload", val)
}
func (x XMLHttpRequestUpload) GetOnloadend() EventHandler {
	val := x.Get("onloadend")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestUpload) SetOnloadend(val EventHandler) {
	x.Set("onloadend", val)
}
func (x XMLHttpRequestUpload) GetOnloadstart() EventHandler {
	val := x.Get("onloadstart")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestUpload) SetOnloadstart(val EventHandler) {
	x.Set("onloadstart", val)
}
func (x XMLHttpRequestUpload) GetOnprogress() EventHandler {
	val := x.Get("onprogress")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestUpload) SetOnprogress(val EventHandler) {
	x.Set("onprogress", val)
}
func (x XMLHttpRequestUpload) GetOntimeout() EventHandler {
	val := x.Get("ontimeout")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequestUpload) SetOntimeout(val EventHandler) {
	x.Set("ontimeout", val)
}
func (x XMLHttpRequestUpload) RemoveEventListener(args ...interface{}) {
	x.Call("removeEventListener", args...)
}
