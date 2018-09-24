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
	Event
}

func JSValueToPopStateEvent(val js.Value) PopStateEvent {
	return PopStateEvent{Value: JSValueToValue(val)}
}
func (v Value) AsPopStateEvent() PopStateEvent { return PopStateEvent{Value: v} }
func NewPopStateEvent(args ...interface{}) PopStateEvent {
	return PopStateEvent{Value: JSValueToValue(js.Global().Get("PopStateEvent").New(args...))}
}
func (p PopStateEvent) GetState() Value {
	val := p.Get("state")
	return val
}
