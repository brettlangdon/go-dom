// Code generated DO NOT EDIT
// customevent.go
package dom

import "syscall/js"

type CustomEventIFace interface {
	GetBubbles() bool
	GetCancelBubble() bool
	SetCancelBubble(bool)
	GetCancelable() bool
	GetComposed() bool
	ComposedPath(args ...interface{})
	GetCurrentTarget() EventTarget
	GetDefaultPrevented() bool
	GetDetail() Value
	GetEventPhase() int
	InitCustomEvent(args ...interface{})
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
type CustomEvent struct {
	Value
}

func JSValueToCustomEvent(val js.Value) CustomEvent { return CustomEvent{Value: JSValueToValue(val)} }
func (v Value) AsCustomEvent() CustomEvent          { return CustomEvent{Value: v} }
func NewCustomEvent(args ...interface{}) CustomEvent {
	return CustomEvent{Value: JSValueToValue(js.Global().Get("CustomEvent").New(args...))}
}
func (c CustomEvent) GetBubbles() bool {
	val := c.Get("bubbles")
	return val.Bool()
}
func (c CustomEvent) GetCancelBubble() bool {
	val := c.Get("cancelBubble")
	return val.Bool()
}
func (c CustomEvent) SetCancelBubble(val bool) {
	c.Set("cancelBubble", val)
}
func (c CustomEvent) GetCancelable() bool {
	val := c.Get("cancelable")
	return val.Bool()
}
func (c CustomEvent) GetComposed() bool {
	val := c.Get("composed")
	return val.Bool()
}
func (c CustomEvent) ComposedPath(args ...interface{}) {
	c.Call("composedPath", args...)
}
func (c CustomEvent) GetCurrentTarget() EventTarget {
	val := c.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (c CustomEvent) GetDefaultPrevented() bool {
	val := c.Get("defaultPrevented")
	return val.Bool()
}
func (c CustomEvent) GetDetail() Value {
	val := c.Get("detail")
	return val
}
func (c CustomEvent) GetEventPhase() int {
	val := c.Get("eventPhase")
	return val.Int()
}
func (c CustomEvent) InitCustomEvent(args ...interface{}) {
	c.Call("initCustomEvent", args...)
}
func (c CustomEvent) InitEvent(args ...interface{}) {
	c.Call("initEvent", args...)
}
func (c CustomEvent) GetIsTrusted() bool {
	val := c.Get("isTrusted")
	return val.Bool()
}
func (c CustomEvent) PreventDefault(args ...interface{}) {
	c.Call("preventDefault", args...)
}
func (c CustomEvent) GetReturnValue() bool {
	val := c.Get("returnValue")
	return val.Bool()
}
func (c CustomEvent) SetReturnValue(val bool) {
	c.Set("returnValue", val)
}
func (c CustomEvent) GetSrcElement() EventTarget {
	val := c.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (c CustomEvent) StopImmediatePropagation(args ...interface{}) {
	c.Call("stopImmediatePropagation", args...)
}
func (c CustomEvent) StopPropagation(args ...interface{}) {
	c.Call("stopPropagation", args...)
}
func (c CustomEvent) GetTarget() EventTarget {
	val := c.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (c CustomEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := c.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (c CustomEvent) GetType() string {
	val := c.Get("type")
	return val.String()
}
