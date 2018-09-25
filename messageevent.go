// Code generated DO NOT EDIT
// messageevent.go
package dom

import "syscall/js"

type MessageEventIFace interface {
	GetBubbles() bool
	GetCancelBubble() bool
	SetCancelBubble(bool)
	GetCancelable() bool
	GetComposed() bool
	ComposedPath(args ...interface{})
	GetCurrentTarget() EventTarget
	GetData() Value
	GetDefaultPrevented() bool
	GetEventPhase() int
	InitEvent(args ...interface{})
	InitMessageEvent(args ...interface{})
	GetIsTrusted() bool
	GetLastEventId() string
	GetOrigin() string
	GetPorts()
	PreventDefault(args ...interface{})
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSource() MessageEventSource
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetType() string
}
type MessageEvent struct {
	Value
}

func JSValueToMessageEvent(val js.Value) MessageEvent { return MessageEvent{Value: JSValueToValue(val)} }
func (v Value) AsMessageEvent() MessageEvent          { return MessageEvent{Value: v} }
func NewMessageEvent(args ...interface{}) MessageEvent {
	return MessageEvent{Value: JSValueToValue(js.Global().Get("MessageEvent").New(args...))}
}
func (m MessageEvent) GetBubbles() bool {
	val := m.Get("bubbles")
	return val.Bool()
}
func (m MessageEvent) GetCancelBubble() bool {
	val := m.Get("cancelBubble")
	return val.Bool()
}
func (m MessageEvent) SetCancelBubble(val bool) {
	m.Set("cancelBubble", val)
}
func (m MessageEvent) GetCancelable() bool {
	val := m.Get("cancelable")
	return val.Bool()
}
func (m MessageEvent) GetComposed() bool {
	val := m.Get("composed")
	return val.Bool()
}
func (m MessageEvent) ComposedPath(args ...interface{}) {
	m.Call("composedPath", args...)
}
func (m MessageEvent) GetCurrentTarget() EventTarget {
	val := m.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (m MessageEvent) GetData() Value {
	val := m.Get("data")
	return val
}
func (m MessageEvent) GetDefaultPrevented() bool {
	val := m.Get("defaultPrevented")
	return val.Bool()
}
func (m MessageEvent) GetEventPhase() int {
	val := m.Get("eventPhase")
	return val.Int()
}
func (m MessageEvent) InitEvent(args ...interface{}) {
	m.Call("initEvent", args...)
}
func (m MessageEvent) InitMessageEvent(args ...interface{}) {
	m.Call("initMessageEvent", args...)
}
func (m MessageEvent) GetIsTrusted() bool {
	val := m.Get("isTrusted")
	return val.Bool()
}
func (m MessageEvent) GetLastEventId() string {
	val := m.Get("lastEventId")
	return val.String()
}
func (m MessageEvent) GetOrigin() string {
	val := m.Get("origin")
	return val.String()
}
func (m MessageEvent) GetPorts() Value {
	val := m.Get("ports")
	return val
}
func (m MessageEvent) PreventDefault(args ...interface{}) {
	m.Call("preventDefault", args...)
}
func (m MessageEvent) GetReturnValue() bool {
	val := m.Get("returnValue")
	return val.Bool()
}
func (m MessageEvent) SetReturnValue(val bool) {
	m.Set("returnValue", val)
}
func (m MessageEvent) GetSource() MessageEventSource {
	val := m.Get("source")
	return JSValueToMessageEventSource(val.JSValue())
}
func (m MessageEvent) GetSrcElement() EventTarget {
	val := m.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (m MessageEvent) StopImmediatePropagation(args ...interface{}) {
	m.Call("stopImmediatePropagation", args...)
}
func (m MessageEvent) StopPropagation(args ...interface{}) {
	m.Call("stopPropagation", args...)
}
func (m MessageEvent) GetTarget() EventTarget {
	val := m.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (m MessageEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := m.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (m MessageEvent) GetType() string {
	val := m.Get("type")
	return val.String()
}
