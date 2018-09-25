// Code generated DO NOT EDIT
// popstateevent.go
package dom

import "syscall/js"

type PopStateEventIFace interface {
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
	GetState() Value
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetType() string
}
type PopStateEvent struct {
	Value
}

func JSValueToPopStateEvent(val js.Value) PopStateEvent {
	return PopStateEvent{Value: JSValueToValue(val)}
}
func (v Value) AsPopStateEvent() PopStateEvent { return PopStateEvent{Value: v} }
func NewPopStateEvent(args ...interface{}) PopStateEvent {
	return PopStateEvent{Value: JSValueToValue(js.Global().Get("PopStateEvent").New(args...))}
}
func (p PopStateEvent) GetBubbles() bool {
	val := p.Get("bubbles")
	return val.Bool()
}
func (p PopStateEvent) GetCancelBubble() bool {
	val := p.Get("cancelBubble")
	return val.Bool()
}
func (p PopStateEvent) SetCancelBubble(val bool) {
	p.Set("cancelBubble", val)
}
func (p PopStateEvent) GetCancelable() bool {
	val := p.Get("cancelable")
	return val.Bool()
}
func (p PopStateEvent) GetComposed() bool {
	val := p.Get("composed")
	return val.Bool()
}
func (p PopStateEvent) ComposedPath(args ...interface{}) {
	p.Call("composedPath", args...)
}
func (p PopStateEvent) GetCurrentTarget() EventTarget {
	val := p.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (p PopStateEvent) GetDefaultPrevented() bool {
	val := p.Get("defaultPrevented")
	return val.Bool()
}
func (p PopStateEvent) GetEventPhase() int {
	val := p.Get("eventPhase")
	return val.Int()
}
func (p PopStateEvent) InitEvent(args ...interface{}) {
	p.Call("initEvent", args...)
}
func (p PopStateEvent) GetIsTrusted() bool {
	val := p.Get("isTrusted")
	return val.Bool()
}
func (p PopStateEvent) PreventDefault(args ...interface{}) {
	p.Call("preventDefault", args...)
}
func (p PopStateEvent) GetReturnValue() bool {
	val := p.Get("returnValue")
	return val.Bool()
}
func (p PopStateEvent) SetReturnValue(val bool) {
	p.Set("returnValue", val)
}
func (p PopStateEvent) GetSrcElement() EventTarget {
	val := p.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (p PopStateEvent) GetState() Value {
	val := p.Get("state")
	return val
}
func (p PopStateEvent) StopImmediatePropagation(args ...interface{}) {
	p.Call("stopImmediatePropagation", args...)
}
func (p PopStateEvent) StopPropagation(args ...interface{}) {
	p.Call("stopPropagation", args...)
}
func (p PopStateEvent) GetTarget() EventTarget {
	val := p.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (p PopStateEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := p.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (p PopStateEvent) GetType() string {
	val := p.Get("type")
	return val.String()
}
