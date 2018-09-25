// Code generated DO NOT EDIT
// promiserejectionevent.go
package dom

import "syscall/js"

type PromiseRejectionEventIFace interface {
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
	GetPromise()
	GetReason() Value
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetType() string
}
type PromiseRejectionEvent struct {
	Value
}

func JSValueToPromiseRejectionEvent(val js.Value) PromiseRejectionEvent {
	return PromiseRejectionEvent{Value: JSValueToValue(val)}
}
func (v Value) AsPromiseRejectionEvent() PromiseRejectionEvent { return PromiseRejectionEvent{Value: v} }
func NewPromiseRejectionEvent(args ...interface{}) PromiseRejectionEvent {
	return PromiseRejectionEvent{Value: JSValueToValue(js.Global().Get("PromiseRejectionEvent").New(args...))}
}
func (p PromiseRejectionEvent) GetBubbles() bool {
	val := p.Get("bubbles")
	return val.Bool()
}
func (p PromiseRejectionEvent) GetCancelBubble() bool {
	val := p.Get("cancelBubble")
	return val.Bool()
}
func (p PromiseRejectionEvent) SetCancelBubble(val bool) {
	p.Set("cancelBubble", val)
}
func (p PromiseRejectionEvent) GetCancelable() bool {
	val := p.Get("cancelable")
	return val.Bool()
}
func (p PromiseRejectionEvent) GetComposed() bool {
	val := p.Get("composed")
	return val.Bool()
}
func (p PromiseRejectionEvent) ComposedPath(args ...interface{}) {
	p.Call("composedPath", args...)
}
func (p PromiseRejectionEvent) GetCurrentTarget() EventTarget {
	val := p.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (p PromiseRejectionEvent) GetDefaultPrevented() bool {
	val := p.Get("defaultPrevented")
	return val.Bool()
}
func (p PromiseRejectionEvent) GetEventPhase() int {
	val := p.Get("eventPhase")
	return val.Int()
}
func (p PromiseRejectionEvent) InitEvent(args ...interface{}) {
	p.Call("initEvent", args...)
}
func (p PromiseRejectionEvent) GetIsTrusted() bool {
	val := p.Get("isTrusted")
	return val.Bool()
}
func (p PromiseRejectionEvent) PreventDefault(args ...interface{}) {
	p.Call("preventDefault", args...)
}
func (p PromiseRejectionEvent) GetPromise() Value {
	val := p.Get("promise")
	return val
}
func (p PromiseRejectionEvent) GetReason() Value {
	val := p.Get("reason")
	return val
}
func (p PromiseRejectionEvent) GetReturnValue() bool {
	val := p.Get("returnValue")
	return val.Bool()
}
func (p PromiseRejectionEvent) SetReturnValue(val bool) {
	p.Set("returnValue", val)
}
func (p PromiseRejectionEvent) GetSrcElement() EventTarget {
	val := p.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (p PromiseRejectionEvent) StopImmediatePropagation(args ...interface{}) {
	p.Call("stopImmediatePropagation", args...)
}
func (p PromiseRejectionEvent) StopPropagation(args ...interface{}) {
	p.Call("stopPropagation", args...)
}
func (p PromiseRejectionEvent) GetTarget() EventTarget {
	val := p.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (p PromiseRejectionEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := p.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (p PromiseRejectionEvent) GetType() string {
	val := p.Get("type")
	return val.String()
}
