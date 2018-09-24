// Code generated DO NOT EDIT
// event.go
package dom

import "syscall/js"

type EventIFace interface {
	GetBubbles() bool
	GetCancelBubble() bool
	SetCancelBubble(bool)
	GetCancelable() bool
	GetComposed() bool
	ComposedPath(args ...interface{})
	GetCurrentTarget() EventTarget
	GetDefaultPrevented() bool
	GetEventPhase() int
	InitEvent(args ...interface{})
	GetIsTrusted() bool
	PreventDefault(args ...interface{})
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetType() string
}
type Event struct {
	Value
}

func JSValueToEvent(val js.Value) Event { return Event{Value: JSValueToValue(val)} }
func (v Value) AsEvent() Event          { return Event{Value: v} }
func NewEvent(args ...interface{}) Event {
	return Event{Value: JSValueToValue(js.Global().Get("Event").New(args...))}
}
func (e Event) GetBubbles() bool {
	val := e.Get("bubbles")
	return val.Bool()
}
func (e Event) GetCancelBubble() bool {
	val := e.Get("cancelBubble")
	return val.Bool()
}
func (e Event) SetCancelBubble(val bool) {
	e.Set("cancelBubble", val)
}
func (e Event) GetCancelable() bool {
	val := e.Get("cancelable")
	return val.Bool()
}
func (e Event) GetComposed() bool {
	val := e.Get("composed")
	return val.Bool()
}
func (e Event) ComposedPath(args ...interface{}) {
	e.Call("composedPath", args...)
}
func (e Event) GetCurrentTarget() EventTarget {
	val := e.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (e Event) GetDefaultPrevented() bool {
	val := e.Get("defaultPrevented")
	return val.Bool()
}
func (e Event) GetEventPhase() int {
	val := e.Get("eventPhase")
	return val.Int()
}
func (e Event) InitEvent(args ...interface{}) {
	e.Call("initEvent", args...)
}
func (e Event) GetIsTrusted() bool {
	val := e.Get("isTrusted")
	return val.Bool()
}
func (e Event) PreventDefault(args ...interface{}) {
	e.Call("preventDefault", args...)
}
func (e Event) GetReturnValue() bool {
	val := e.Get("returnValue")
	return val.Bool()
}
func (e Event) SetReturnValue(val bool) {
	e.Set("returnValue", val)
}
func (e Event) GetSrcElement() EventTarget {
	val := e.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (e Event) StopImmediatePropagation(args ...interface{}) {
	e.Call("stopImmediatePropagation", args...)
}
func (e Event) StopPropagation(args ...interface{}) {
	e.Call("stopPropagation", args...)
}
func (e Event) GetTarget() EventTarget {
	val := e.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (e Event) GetTimeStamp() DOMHighResTimeStamp {
	val := e.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (e Event) GetType() string {
	val := e.Get("type")
	return val.String()
}
