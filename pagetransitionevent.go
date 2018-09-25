// Code generated DO NOT EDIT
// pagetransitionevent.go
package dom

import "syscall/js"

type PageTransitionEventIFace interface {
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
	GetPersisted() bool
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
type PageTransitionEvent struct {
	Value
}

func JSValueToPageTransitionEvent(val js.Value) PageTransitionEvent {
	return PageTransitionEvent{Value: JSValueToValue(val)}
}
func (v Value) AsPageTransitionEvent() PageTransitionEvent { return PageTransitionEvent{Value: v} }
func NewPageTransitionEvent(args ...interface{}) PageTransitionEvent {
	return PageTransitionEvent{Value: JSValueToValue(js.Global().Get("PageTransitionEvent").New(args...))}
}
func (p PageTransitionEvent) GetBubbles() bool {
	val := p.Get("bubbles")
	return val.Bool()
}
func (p PageTransitionEvent) GetCancelBubble() bool {
	val := p.Get("cancelBubble")
	return val.Bool()
}
func (p PageTransitionEvent) SetCancelBubble(val bool) {
	p.Set("cancelBubble", val)
}
func (p PageTransitionEvent) GetCancelable() bool {
	val := p.Get("cancelable")
	return val.Bool()
}
func (p PageTransitionEvent) GetComposed() bool {
	val := p.Get("composed")
	return val.Bool()
}
func (p PageTransitionEvent) ComposedPath(args ...interface{}) {
	p.Call("composedPath", args...)
}
func (p PageTransitionEvent) GetCurrentTarget() EventTarget {
	val := p.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (p PageTransitionEvent) GetDefaultPrevented() bool {
	val := p.Get("defaultPrevented")
	return val.Bool()
}
func (p PageTransitionEvent) GetEventPhase() int {
	val := p.Get("eventPhase")
	return val.Int()
}
func (p PageTransitionEvent) InitEvent(args ...interface{}) {
	p.Call("initEvent", args...)
}
func (p PageTransitionEvent) GetIsTrusted() bool {
	val := p.Get("isTrusted")
	return val.Bool()
}
func (p PageTransitionEvent) GetPersisted() bool {
	val := p.Get("persisted")
	return val.Bool()
}
func (p PageTransitionEvent) PreventDefault(args ...interface{}) {
	p.Call("preventDefault", args...)
}
func (p PageTransitionEvent) GetReturnValue() bool {
	val := p.Get("returnValue")
	return val.Bool()
}
func (p PageTransitionEvent) SetReturnValue(val bool) {
	p.Set("returnValue", val)
}
func (p PageTransitionEvent) GetSrcElement() EventTarget {
	val := p.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (p PageTransitionEvent) StopImmediatePropagation(args ...interface{}) {
	p.Call("stopImmediatePropagation", args...)
}
func (p PageTransitionEvent) StopPropagation(args ...interface{}) {
	p.Call("stopPropagation", args...)
}
func (p PageTransitionEvent) GetTarget() EventTarget {
	val := p.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (p PageTransitionEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := p.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (p PageTransitionEvent) GetType() string {
	val := p.Get("type")
	return val.String()
}
