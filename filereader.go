// Code generated DO NOT EDIT
// filereader.go
package dom

import "syscall/js"

type FileReaderIFace interface {
	Abort(args ...interface{})
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetError() DOMException
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
	ReadAsArrayBuffer(args ...interface{})
	ReadAsBinaryString(args ...interface{})
	ReadAsDataURL(args ...interface{})
	ReadAsText(args ...interface{})
	GetReadyState() int
	RemoveEventListener(args ...interface{})
	GetResult()
}
type FileReader struct {
	Value
}

func JSValueToFileReader(val js.Value) FileReader { return FileReader{Value: JSValueToValue(val)} }
func (v Value) AsFileReader() FileReader          { return FileReader{Value: v} }
func NewFileReader(args ...interface{}) FileReader {
	return FileReader{Value: JSValueToValue(js.Global().Get("FileReader").New(args...))}
}
func (f FileReader) Abort(args ...interface{}) {
	f.Call("abort", args...)
}
func (f FileReader) AddEventListener(args ...interface{}) {
	f.Call("addEventListener", args...)
}
func (f FileReader) DispatchEvent(args ...interface{}) bool {
	val := f.Call("dispatchEvent", args...)
	return val.Bool()
}
func (f FileReader) GetError() DOMException {
	val := f.Get("error")
	return JSValueToDOMException(val.JSValue())
}
func (f FileReader) GetOnabort() EventHandler {
	val := f.Get("onabort")
	return JSValueToEventHandler(val.JSValue())
}
func (f FileReader) SetOnabort(val EventHandler) {
	f.Set("onabort", val)
}
func (f FileReader) GetOnerror() EventHandler {
	val := f.Get("onerror")
	return JSValueToEventHandler(val.JSValue())
}
func (f FileReader) SetOnerror(val EventHandler) {
	f.Set("onerror", val)
}
func (f FileReader) GetOnload() EventHandler {
	val := f.Get("onload")
	return JSValueToEventHandler(val.JSValue())
}
func (f FileReader) SetOnload(val EventHandler) {
	f.Set("onload", val)
}
func (f FileReader) GetOnloadend() EventHandler {
	val := f.Get("onloadend")
	return JSValueToEventHandler(val.JSValue())
}
func (f FileReader) SetOnloadend(val EventHandler) {
	f.Set("onloadend", val)
}
func (f FileReader) GetOnloadstart() EventHandler {
	val := f.Get("onloadstart")
	return JSValueToEventHandler(val.JSValue())
}
func (f FileReader) SetOnloadstart(val EventHandler) {
	f.Set("onloadstart", val)
}
func (f FileReader) GetOnprogress() EventHandler {
	val := f.Get("onprogress")
	return JSValueToEventHandler(val.JSValue())
}
func (f FileReader) SetOnprogress(val EventHandler) {
	f.Set("onprogress", val)
}
func (f FileReader) ReadAsArrayBuffer(args ...interface{}) {
	f.Call("readAsArrayBuffer", args...)
}
func (f FileReader) ReadAsBinaryString(args ...interface{}) {
	f.Call("readAsBinaryString", args...)
}
func (f FileReader) ReadAsDataURL(args ...interface{}) {
	f.Call("readAsDataURL", args...)
}
func (f FileReader) ReadAsText(args ...interface{}) {
	f.Call("readAsText", args...)
}
func (f FileReader) GetReadyState() int {
	val := f.Get("readyState")
	return val.Int()
}
func (f FileReader) RemoveEventListener(args ...interface{}) {
	f.Call("removeEventListener", args...)
}
func (f FileReader) GetResult() Value {
	val := f.Get("result")
	return val
}
