// Code generated DO NOT EDIT
// messageport.go
package dom

import "syscall/js"

type MessagePortIFace interface {
	AddEventListener(args ...interface{})
	Close(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnmessageerror() EventHandler
	SetOnmessageerror(EventHandler)
	PostMessage(args ...interface{})
	RemoveEventListener(args ...interface{})
	Start(args ...interface{})
}
type MessagePort struct {
	Value
}

func JSValueToMessagePort(val js.Value) MessagePort { return MessagePort{Value: JSValueToValue(val)} }
func (v Value) AsMessagePort() MessagePort          { return MessagePort{Value: v} }
func NewMessagePort(args ...interface{}) MessagePort {
	return MessagePort{Value: JSValueToValue(js.Global().Get("MessagePort").New(args...))}
}
func (m MessagePort) AddEventListener(args ...interface{}) {
	m.Call("addEventListener", args...)
}
func (m MessagePort) Close(args ...interface{}) {
	m.Call("close", args...)
}
func (m MessagePort) DispatchEvent(args ...interface{}) bool {
	val := m.Call("dispatchEvent", args...)
	return val.Bool()
}
func (m MessagePort) GetOnmessage() EventHandler {
	val := m.Get("onmessage")
	return JSValueToEventHandler(val.JSValue())
}
func (m MessagePort) SetOnmessage(val EventHandler) {
	m.Set("onmessage", val)
}
func (m MessagePort) GetOnmessageerror() EventHandler {
	val := m.Get("onmessageerror")
	return JSValueToEventHandler(val.JSValue())
}
func (m MessagePort) SetOnmessageerror(val EventHandler) {
	m.Set("onmessageerror", val)
}
func (m MessagePort) PostMessage(args ...interface{}) {
	m.Call("postMessage", args...)
}
func (m MessagePort) RemoveEventListener(args ...interface{}) {
	m.Call("removeEventListener", args...)
}
func (m MessagePort) Start(args ...interface{}) {
	m.Call("start", args...)
}
