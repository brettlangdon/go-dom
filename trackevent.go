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
	Event
}

func JSValueToTrackEvent(val js.Value) TrackEvent { return TrackEvent{Value: JSValueToValue(val)} }
func (v Value) AsTrackEvent() TrackEvent          { return TrackEvent{Value: v} }
func NewTrackEvent(args ...interface{}) TrackEvent {
	return TrackEvent{Value: JSValueToValue(js.Global().Get("TrackEvent").New(args...))}
}
func (t TrackEvent) GetTrack() Value {
	val := t.Get("track")
	return val
}
