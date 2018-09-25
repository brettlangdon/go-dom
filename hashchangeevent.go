// Code generated DO NOT EDIT
// hashchangeevent.go
package dom

import "syscall/js"

type HashChangeEventIFace interface {
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
	GetNewURL() string
	GetOldURL() string
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
type HashChangeEvent struct {
	Value
}

func JSValueToHashChangeEvent(val js.Value) HashChangeEvent {
	return HashChangeEvent{Value: JSValueToValue(val)}
}
func (v Value) AsHashChangeEvent() HashChangeEvent { return HashChangeEvent{Value: v} }
func NewHashChangeEvent(args ...interface{}) HashChangeEvent {
	return HashChangeEvent{Value: JSValueToValue(js.Global().Get("HashChangeEvent").New(args...))}
}
func (h HashChangeEvent) GetBubbles() bool {
	val := h.Get("bubbles")
	return val.Bool()
}
func (h HashChangeEvent) GetCancelBubble() bool {
	val := h.Get("cancelBubble")
	return val.Bool()
}
func (h HashChangeEvent) SetCancelBubble(val bool) {
	h.Set("cancelBubble", val)
}
func (h HashChangeEvent) GetCancelable() bool {
	val := h.Get("cancelable")
	return val.Bool()
}
func (h HashChangeEvent) GetComposed() bool {
	val := h.Get("composed")
	return val.Bool()
}
func (h HashChangeEvent) ComposedPath(args ...interface{}) {
	h.Call("composedPath", args...)
}
func (h HashChangeEvent) GetCurrentTarget() EventTarget {
	val := h.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (h HashChangeEvent) GetDefaultPrevented() bool {
	val := h.Get("defaultPrevented")
	return val.Bool()
}
func (h HashChangeEvent) GetEventPhase() int {
	val := h.Get("eventPhase")
	return val.Int()
}
func (h HashChangeEvent) InitEvent(args ...interface{}) {
	h.Call("initEvent", args...)
}
func (h HashChangeEvent) GetIsTrusted() bool {
	val := h.Get("isTrusted")
	return val.Bool()
}
func (h HashChangeEvent) GetNewURL() string {
	val := h.Get("newURL")
	return val.String()
}
func (h HashChangeEvent) GetOldURL() string {
	val := h.Get("oldURL")
	return val.String()
}
func (h HashChangeEvent) PreventDefault(args ...interface{}) {
	h.Call("preventDefault", args...)
}
func (h HashChangeEvent) GetReturnValue() bool {
	val := h.Get("returnValue")
	return val.Bool()
}
func (h HashChangeEvent) SetReturnValue(val bool) {
	h.Set("returnValue", val)
}
func (h HashChangeEvent) GetSrcElement() EventTarget {
	val := h.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (h HashChangeEvent) StopImmediatePropagation(args ...interface{}) {
	h.Call("stopImmediatePropagation", args...)
}
func (h HashChangeEvent) StopPropagation(args ...interface{}) {
	h.Call("stopPropagation", args...)
}
func (h HashChangeEvent) GetTarget() EventTarget {
	val := h.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (h HashChangeEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := h.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (h HashChangeEvent) GetType() string {
	val := h.Get("type")
	return val.String()
}
