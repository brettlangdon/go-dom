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

func jsValueToPageTransitionEvent(val js.Value) PageTransitionEvent {
	return PageTransitionEvent{Value: Value{Value: val}}
}
func (v Value) AsPageTransitionEvent() PageTransitionEvent { return PageTransitionEvent{Value: v} }
func (p PageTransitionEvent) GetPersisted() bool {
	val := p.Get("persisted")
	return val.Bool()
}
