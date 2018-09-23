// Code generated DO NOT EDIT
// closeevent.go
package dom

import "syscall/js"

type CloseEventIFace interface {
	GetBubbles() bool
	GetCancelBubble() bool
	SetCancelBubble(bool)
	GetCancelable() bool
	GetCode() int
	GetComposed() bool
	ComposedPath(args ...interface{})
	GetCurrentTarget() EventTarget
	GetDefaultPrevented() bool
	GetEventPhase() int
	InitEvent(args ...interface{})
	GetIsTrusted() bool
	PreventDefault(args ...interface{})
	GetReason() string
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetType() string
	GetWasClean() bool
}
type CloseEvent struct {
	Value
	Event
}

func jsValueToCloseEvent(val js.Value) CloseEvent { return CloseEvent{Value: Value{Value: val}} }
func (v Value) AsCloseEvent() CloseEvent          { return CloseEvent{Value: v} }
func (c CloseEvent) GetCode() int {
	val := c.Get("code")
	return val.Int()
}
func (c CloseEvent) GetReason() string {
	val := c.Get("reason")
	return val.String()
}
func (c CloseEvent) GetWasClean() bool {
	val := c.Get("wasClean")
	return val.Bool()
}
