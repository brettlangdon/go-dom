// Code generated DO NOT EDIT
// beforeunloadevent.go
package dom

import "syscall/js"

type BeforeUnloadEventIFace interface {
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
	GetReturnValue() string
	SetReturnValue(string)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetType() string
}
type BeforeUnloadEvent struct {
	Value
}

func JSValueToBeforeUnloadEvent(val js.Value) BeforeUnloadEvent {
	return BeforeUnloadEvent{Value: JSValueToValue(val)}
}
func (v Value) AsBeforeUnloadEvent() BeforeUnloadEvent { return BeforeUnloadEvent{Value: v} }
func NewBeforeUnloadEvent(args ...interface{}) BeforeUnloadEvent {
	return BeforeUnloadEvent{Value: JSValueToValue(js.Global().Get("BeforeUnloadEvent").New(args...))}
}
func (b BeforeUnloadEvent) GetBubbles() bool {
	val := b.Get("bubbles")
	return val.Bool()
}
func (b BeforeUnloadEvent) GetCancelBubble() bool {
	val := b.Get("cancelBubble")
	return val.Bool()
}
func (b BeforeUnloadEvent) SetCancelBubble(val bool) {
	b.Set("cancelBubble", val)
}
func (b BeforeUnloadEvent) GetCancelable() bool {
	val := b.Get("cancelable")
	return val.Bool()
}
func (b BeforeUnloadEvent) GetComposed() bool {
	val := b.Get("composed")
	return val.Bool()
}
func (b BeforeUnloadEvent) ComposedPath(args ...interface{}) {
	b.Call("composedPath", args...)
}
func (b BeforeUnloadEvent) GetCurrentTarget() EventTarget {
	val := b.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (b BeforeUnloadEvent) GetDefaultPrevented() bool {
	val := b.Get("defaultPrevented")
	return val.Bool()
}
func (b BeforeUnloadEvent) GetEventPhase() int {
	val := b.Get("eventPhase")
	return val.Int()
}
func (b BeforeUnloadEvent) InitEvent(args ...interface{}) {
	b.Call("initEvent", args...)
}
func (b BeforeUnloadEvent) GetIsTrusted() bool {
	val := b.Get("isTrusted")
	return val.Bool()
}
func (b BeforeUnloadEvent) PreventDefault(args ...interface{}) {
	b.Call("preventDefault", args...)
}
func (b BeforeUnloadEvent) GetReturnValue() string {
	val := b.Get("returnValue")
	return val.String()
}
func (b BeforeUnloadEvent) SetReturnValue(val string) {
	b.Set("returnValue", val)
}
func (b BeforeUnloadEvent) GetSrcElement() EventTarget {
	val := b.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (b BeforeUnloadEvent) StopImmediatePropagation(args ...interface{}) {
	b.Call("stopImmediatePropagation", args...)
}
func (b BeforeUnloadEvent) StopPropagation(args ...interface{}) {
	b.Call("stopPropagation", args...)
}
func (b BeforeUnloadEvent) GetTarget() EventTarget {
	val := b.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (b BeforeUnloadEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := b.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (b BeforeUnloadEvent) GetType() string {
	val := b.Get("type")
	return val.String()
}
