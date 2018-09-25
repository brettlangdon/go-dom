// Code generated DO NOT EDIT
// trackevent.go
package dom

import "syscall/js"

type TrackEventIFace interface {
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
	GetTrack()
	GetType() string
}
type TrackEvent struct {
	Value
}

func JSValueToTrackEvent(val js.Value) TrackEvent { return TrackEvent{Value: JSValueToValue(val)} }
func (v Value) AsTrackEvent() TrackEvent          { return TrackEvent{Value: v} }
func NewTrackEvent(args ...interface{}) TrackEvent {
	return TrackEvent{Value: JSValueToValue(js.Global().Get("TrackEvent").New(args...))}
}
func (t TrackEvent) GetBubbles() bool {
	val := t.Get("bubbles")
	return val.Bool()
}
func (t TrackEvent) GetCancelBubble() bool {
	val := t.Get("cancelBubble")
	return val.Bool()
}
func (t TrackEvent) SetCancelBubble(val bool) {
	t.Set("cancelBubble", val)
}
func (t TrackEvent) GetCancelable() bool {
	val := t.Get("cancelable")
	return val.Bool()
}
func (t TrackEvent) GetComposed() bool {
	val := t.Get("composed")
	return val.Bool()
}
func (t TrackEvent) ComposedPath(args ...interface{}) {
	t.Call("composedPath", args...)
}
func (t TrackEvent) GetCurrentTarget() EventTarget {
	val := t.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (t TrackEvent) GetDefaultPrevented() bool {
	val := t.Get("defaultPrevented")
	return val.Bool()
}
func (t TrackEvent) GetEventPhase() int {
	val := t.Get("eventPhase")
	return val.Int()
}
func (t TrackEvent) InitEvent(args ...interface{}) {
	t.Call("initEvent", args...)
}
func (t TrackEvent) GetIsTrusted() bool {
	val := t.Get("isTrusted")
	return val.Bool()
}
func (t TrackEvent) PreventDefault(args ...interface{}) {
	t.Call("preventDefault", args...)
}
func (t TrackEvent) GetReturnValue() bool {
	val := t.Get("returnValue")
	return val.Bool()
}
func (t TrackEvent) SetReturnValue(val bool) {
	t.Set("returnValue", val)
}
func (t TrackEvent) GetSrcElement() EventTarget {
	val := t.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (t TrackEvent) StopImmediatePropagation(args ...interface{}) {
	t.Call("stopImmediatePropagation", args...)
}
func (t TrackEvent) StopPropagation(args ...interface{}) {
	t.Call("stopPropagation", args...)
}
func (t TrackEvent) GetTarget() EventTarget {
	val := t.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (t TrackEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := t.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (t TrackEvent) GetTrack() Value {
	val := t.Get("track")
	return val
}
func (t TrackEvent) GetType() string {
	val := t.Get("type")
	return val.String()
}
