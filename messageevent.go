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
	Event
}

func JSValueToMessageEvent(val js.Value) MessageEvent { return MessageEvent{Value: JSValueToValue(val)} }
func (v Value) AsMessageEvent() MessageEvent          { return MessageEvent{Value: v} }
func NewMessageEvent(args ...interface{}) MessageEvent {
	return MessageEvent{Value: JSValueToValue(js.Global().Get("MessageEvent").New(args...))}
}
func (m MessageEvent) GetData() Value {
	val := m.Get("data")
	return val
}
func (m MessageEvent) InitMessageEvent(args ...interface{}) {
	m.Call("initMessageEvent", args...)
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
func (m MessageEvent) GetSource() MessageEventSource {
	val := m.Get("source")
	return JSValueToMessageEventSource(val.JSValue())
}
