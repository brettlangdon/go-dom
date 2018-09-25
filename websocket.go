// Code generated DO NOT EDIT
// websocket.go
package dom

import "syscall/js"

type WebSocketIFace interface {
	AddEventListener(args ...interface{})
	GetBinaryType() BinaryType
	SetBinaryType(BinaryType)
	GetBufferedAmount() int
	Close(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetExtensions() string
	GetOnclose() EventHandler
	SetOnclose(EventHandler)
	GetOnerror() EventHandler
	SetOnerror(EventHandler)
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnopen() EventHandler
	SetOnopen(EventHandler)
	GetProtocol() string
	GetReadyState() int
	RemoveEventListener(args ...interface{})
	SendWithArrayBuffer(args ...interface{})
	SendWithArrayBufferView(args ...interface{})
	SendWithBlob(args ...interface{})
	SendWithUSVString(args ...interface{})
	GetUrl() string
}
type WebSocket struct {
	Value
}

func JSValueToWebSocket(val js.Value) WebSocket { return WebSocket{Value: JSValueToValue(val)} }
func (v Value) AsWebSocket() WebSocket          { return WebSocket{Value: v} }
func NewWebSocket(args ...interface{}) WebSocket {
	return WebSocket{Value: JSValueToValue(js.Global().Get("WebSocket").New(args...))}
}
func (w WebSocket) AddEventListener(args ...interface{}) {
	w.Call("addEventListener", args...)
}
func (w WebSocket) GetBinaryType() BinaryType {
	val := w.Get("binaryType")
	return JSValueToBinaryType(val.JSValue())
}
func (w WebSocket) SetBinaryType(val BinaryType) {
	w.Set("binaryType", val)
}
func (w WebSocket) GetBufferedAmount() int {
	val := w.Get("bufferedAmount")
	return val.Int()
}
func (w WebSocket) Close(args ...interface{}) {
	w.Call("close", args...)
}
func (w WebSocket) DispatchEvent(args ...interface{}) bool {
	val := w.Call("dispatchEvent", args...)
	return val.Bool()
}
func (w WebSocket) GetExtensions() string {
	val := w.Get("extensions")
	return val.String()
}
func (w WebSocket) GetOnclose() EventHandler {
	val := w.Get("onclose")
	return JSValueToEventHandler(val.JSValue())
}
func (w WebSocket) SetOnclose(val EventHandler) {
	w.Set("onclose", val)
}
func (w WebSocket) GetOnerror() EventHandler {
	val := w.Get("onerror")
	return JSValueToEventHandler(val.JSValue())
}
func (w WebSocket) SetOnerror(val EventHandler) {
	w.Set("onerror", val)
}
func (w WebSocket) GetOnmessage() EventHandler {
	val := w.Get("onmessage")
	return JSValueToEventHandler(val.JSValue())
}
func (w WebSocket) SetOnmessage(val EventHandler) {
	w.Set("onmessage", val)
}
func (w WebSocket) GetOnopen() EventHandler {
	val := w.Get("onopen")
	return JSValueToEventHandler(val.JSValue())
}
func (w WebSocket) SetOnopen(val EventHandler) {
	w.Set("onopen", val)
}
func (w WebSocket) GetProtocol() string {
	val := w.Get("protocol")
	return val.String()
}
func (w WebSocket) GetReadyState() int {
	val := w.Get("readyState")
	return val.Int()
}
func (w WebSocket) RemoveEventListener(args ...interface{}) {
	w.Call("removeEventListener", args...)
}
func (w WebSocket) SendWithArrayBuffer(args ...interface{}) {
	w.Call("sendWithArrayBuffer", args...)
}
func (w WebSocket) SendWithArrayBufferView(args ...interface{}) {
	w.Call("sendWithArrayBufferView", args...)
}
func (w WebSocket) SendWithBlob(args ...interface{}) {
	w.Call("sendWithBlob", args...)
}
func (w WebSocket) SendWithUSVString(args ...interface{}) {
	w.Call("sendWithUSVString", args...)
}
func (w WebSocket) GetUrl() string {
	val := w.Get("url")
	return val.String()
}
