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
	Event
}

func JSValueToPageTransitionEvent(val js.Value) PageTransitionEvent {
	return PageTransitionEvent{Value: JSValueToValue(val)}
}
func (v Value) AsPageTransitionEvent() PageTransitionEvent { return PageTransitionEvent{Value: v} }
func NewPageTransitionEvent(args ...interface{}) PageTransitionEvent {
	return PageTransitionEvent{Value: JSValueToValue(js.Global().Get("PageTransitionEvent").New(args...))}
}
func (p PageTransitionEvent) GetPersisted() bool {
	val := p.Get("persisted")
	return val.Bool()
}
