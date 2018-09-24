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
	Event
}

func JSValueToHashChangeEvent(val js.Value) HashChangeEvent {
	return HashChangeEvent{Value: JSValueToValue(val)}
}
func (v Value) AsHashChangeEvent() HashChangeEvent { return HashChangeEvent{Value: v} }
func NewHashChangeEvent(args ...interface{}) HashChangeEvent {
	return HashChangeEvent{Value: JSValueToValue(js.Global().Get("HashChangeEvent").New(args...))}
}
func (h HashChangeEvent) GetNewURL() string {
	val := h.Get("newURL")
	return val.String()
}
func (h HashChangeEvent) GetOldURL() string {
	val := h.Get("oldURL")
	return val.String()
}
