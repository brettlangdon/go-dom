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
	Event
}

func JSValueToBeforeUnloadEvent(val js.Value) BeforeUnloadEvent {
	return BeforeUnloadEvent{Value: Value{Value: val}}
}
func (v Value) AsBeforeUnloadEvent() BeforeUnloadEvent { return BeforeUnloadEvent{Value: v} }
func (b BeforeUnloadEvent) GetReturnValue() string {
	val := b.Get("returnValue")
	return val.String()
}
func (b BeforeUnloadEvent) SetReturnValue(val string) {
	b.Set("returnValue", val)
}
