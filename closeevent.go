// Code generated DO NOT EDIT
// closeevent.go
package dom

import "syscall/js"

type CloseEventIFace interface {
	GetBubbles() bool
	GetCancelBubble() bool
	SetCancelBubble(bool)
	GetCancelable() bool
	GetCode() int
	GetComposed() bool
	ComposedPath(args ...interface{})
	GetCurrentTarget() EventTarget
	GetDefaultPrevented() bool
	GetEventPhase() int
	InitEvent(args ...interface{})
	GetIsTrusted() bool
	PreventDefault(args ...interface{})
	GetReason() string
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetType() string
	GetWasClean() bool
}
type CloseEvent struct {
	Value
}

func JSValueToCloseEvent(val js.Value) CloseEvent { return CloseEvent{Value: JSValueToValue(val)} }
func (v Value) AsCloseEvent() CloseEvent          { return CloseEvent{Value: v} }
func NewCloseEvent(args ...interface{}) CloseEvent {
	return CloseEvent{Value: JSValueToValue(js.Global().Get("CloseEvent").New(args...))}
}
func (c CloseEvent) GetBubbles() bool {
	val := c.Get("bubbles")
	return val.Bool()
}
func (c CloseEvent) GetCancelBubble() bool {
	val := c.Get("cancelBubble")
	return val.Bool()
}
func (c CloseEvent) SetCancelBubble(val bool) {
	c.Set("cancelBubble", val)
}
func (c CloseEvent) GetCancelable() bool {
	val := c.Get("cancelable")
	return val.Bool()
}
func (c CloseEvent) GetCode() int {
	val := c.Get("code")
	return val.Int()
}
func (c CloseEvent) GetComposed() bool {
	val := c.Get("composed")
	return val.Bool()
}
func (c CloseEvent) ComposedPath(args ...interface{}) {
	c.Call("composedPath", args...)
}
func (c CloseEvent) GetCurrentTarget() EventTarget {
	val := c.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (c CloseEvent) GetDefaultPrevented() bool {
	val := c.Get("defaultPrevented")
	return val.Bool()
}
func (c CloseEvent) GetEventPhase() int {
	val := c.Get("eventPhase")
	return val.Int()
}
func (c CloseEvent) InitEvent(args ...interface{}) {
	c.Call("initEvent", args...)
}
func (c CloseEvent) GetIsTrusted() bool {
	val := c.Get("isTrusted")
	return val.Bool()
}
func (c CloseEvent) PreventDefault(args ...interface{}) {
	c.Call("preventDefault", args...)
}
func (c CloseEvent) GetReason() string {
	val := c.Get("reason")
	return val.String()
}
func (c CloseEvent) GetReturnValue() bool {
	val := c.Get("returnValue")
	return val.Bool()
}
func (c CloseEvent) SetReturnValue(val bool) {
	c.Set("returnValue", val)
}
func (c CloseEvent) GetSrcElement() EventTarget {
	val := c.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (c CloseEvent) StopImmediatePropagation(args ...interface{}) {
	c.Call("stopImmediatePropagation", args...)
}
func (c CloseEvent) StopPropagation(args ...interface{}) {
	c.Call("stopPropagation", args...)
}
func (c CloseEvent) GetTarget() EventTarget {
	val := c.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (c CloseEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := c.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (c CloseEvent) GetType() string {
	val := c.Get("type")
	return val.String()
}
func (c CloseEvent) GetWasClean() bool {
	val := c.Get("wasClean")
	return val.Bool()
}
