// Code generated DO NOT EDIT
// xmlhttprequest.go
package dom

import "syscall/js"

type XMLHttpRequestIFace interface {
	Abort(args ...interface{})
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetAllResponseHeaders(args ...interface{}) []byte
	GetResponseHeader(args ...interface{}) []byte
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
	GetOnreadystatechange() EventHandler
	SetOnreadystatechange(EventHandler)
	GetOntimeout() EventHandler
	SetOntimeout(EventHandler)
	Open(args ...interface{})
	OpenWithArgs(args ...interface{})
	OverrideMimeType(args ...interface{})
	GetReadyState() int
	RemoveEventListener(args ...interface{})
	GetResponse() Value
	GetResponseText() string
	GetResponseType() XMLHttpRequestResponseType
	SetResponseType(XMLHttpRequestResponseType)
	GetResponseURL() string
	GetResponseXML() Document
	Send(args ...interface{})
	SetRequestHeader(args ...interface{})
	GetStatus() int
	GetStatusText() []byte
	GetTimeout() int
	SetTimeout(int)
	GetUpload() XMLHttpRequestUpload
	GetWithCredentials() bool
	SetWithCredentials(bool)
}
type XMLHttpRequest struct {
	Value
}

func JSValueToXMLHttpRequest(val js.Value) XMLHttpRequest {
	return XMLHttpRequest{Value: JSValueToValue(val)}
}
func (v Value) AsXMLHttpRequest() XMLHttpRequest { return XMLHttpRequest{Value: v} }
func NewXMLHttpRequest(args ...interface{}) XMLHttpRequest {
	return XMLHttpRequest{Value: JSValueToValue(js.Global().Get("XMLHttpRequest").New(args...))}
}
func (x XMLHttpRequest) Abort(args ...interface{}) {
	x.Call("abort", args...)
}
func (x XMLHttpRequest) AddEventListener(args ...interface{}) {
	x.Call("addEventListener", args...)
}
func (x XMLHttpRequest) DispatchEvent(args ...interface{}) bool {
	val := x.Call("dispatchEvent", args...)
	return val.Bool()
}
func (x XMLHttpRequest) GetAllResponseHeaders(args ...interface{}) []byte {
	val := x.Call("getAllResponseHeaders", args...)
	return []byte(val.String())
}
func (x XMLHttpRequest) GetResponseHeader(args ...interface{}) []byte {
	val := x.Call("getResponseHeader", args...)
	return []byte(val.String())
}
func (x XMLHttpRequest) GetOnabort() EventHandler {
	val := x.Get("onabort")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequest) SetOnabort(val EventHandler) {
	x.Set("onabort", val)
}
func (x XMLHttpRequest) GetOnerror() EventHandler {
	val := x.Get("onerror")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequest) SetOnerror(val EventHandler) {
	x.Set("onerror", val)
}
func (x XMLHttpRequest) GetOnload() EventHandler {
	val := x.Get("onload")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequest) SetOnload(val EventHandler) {
	x.Set("onload", val)
}
func (x XMLHttpRequest) GetOnloadend() EventHandler {
	val := x.Get("onloadend")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequest) SetOnloadend(val EventHandler) {
	x.Set("onloadend", val)
}
func (x XMLHttpRequest) GetOnloadstart() EventHandler {
	val := x.Get("onloadstart")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequest) SetOnloadstart(val EventHandler) {
	x.Set("onloadstart", val)
}
func (x XMLHttpRequest) GetOnprogress() EventHandler {
	val := x.Get("onprogress")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequest) SetOnprogress(val EventHandler) {
	x.Set("onprogress", val)
}
func (x XMLHttpRequest) GetOnreadystatechange() EventHandler {
	val := x.Get("onreadystatechange")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequest) SetOnreadystatechange(val EventHandler) {
	x.Set("onreadystatechange", val)
}
func (x XMLHttpRequest) GetOntimeout() EventHandler {
	val := x.Get("ontimeout")
	return JSValueToEventHandler(val.JSValue())
}
func (x XMLHttpRequest) SetOntimeout(val EventHandler) {
	x.Set("ontimeout", val)
}
func (x XMLHttpRequest) Open(args ...interface{}) {
	x.Call("open", args...)
}
func (x XMLHttpRequest) OpenWithArgs(args ...interface{}) {
	x.Call("openWithArgs", args...)
}
func (x XMLHttpRequest) OverrideMimeType(args ...interface{}) {
	x.Call("overrideMimeType", args...)
}
func (x XMLHttpRequest) GetReadyState() int {
	val := x.Get("readyState")
	return val.Int()
}
func (x XMLHttpRequest) RemoveEventListener(args ...interface{}) {
	x.Call("removeEventListener", args...)
}
func (x XMLHttpRequest) GetResponse() Value {
	val := x.Get("response")
	return val
}
func (x XMLHttpRequest) GetResponseText() string {
	val := x.Get("responseText")
	return val.String()
}
func (x XMLHttpRequest) GetResponseType() XMLHttpRequestResponseType {
	val := x.Get("responseType")
	return JSValueToXMLHttpRequestResponseType(val.JSValue())
}
func (x XMLHttpRequest) SetResponseType(val XMLHttpRequestResponseType) {
	x.Set("responseType", val)
}
func (x XMLHttpRequest) GetResponseURL() string {
	val := x.Get("responseURL")
	return val.String()
}
func (x XMLHttpRequest) GetResponseXML() Document {
	val := x.Get("responseXML")
	return JSValueToDocument(val.JSValue())
}
func (x XMLHttpRequest) Send(args ...interface{}) {
	x.Call("send", args...)
}
func (x XMLHttpRequest) SetRequestHeader(args ...interface{}) {
	x.Call("setRequestHeader", args...)
}
func (x XMLHttpRequest) GetStatus() int {
	val := x.Get("status")
	return val.Int()
}
func (x XMLHttpRequest) GetStatusText() []byte {
	val := x.Get("statusText")
	return []byte(val.String())
}
func (x XMLHttpRequest) GetTimeout() int {
	val := x.Get("timeout")
	return val.Int()
}
func (x XMLHttpRequest) SetTimeout(val int) {
	x.Set("timeout", val)
}
func (x XMLHttpRequest) GetUpload() XMLHttpRequestUpload {
	val := x.Get("upload")
	return JSValueToXMLHttpRequestUpload(val.JSValue())
}
func (x XMLHttpRequest) GetWithCredentials() bool {
	val := x.Get("withCredentials")
	return val.Bool()
}
func (x XMLHttpRequest) SetWithCredentials(val bool) {
	x.Set("withCredentials", val)
}
