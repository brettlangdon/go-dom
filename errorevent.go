// Code generated DO NOT EDIT
// errorevent.go
package dom

import "syscall/js"

type ErrorEventIFace interface {
	GetBubbles() bool
	GetCancelBubble() bool
	SetCancelBubble(bool)
	GetCancelable() bool
	GetColno() int
	GetComposed() bool
	ComposedPath(args ...interface{})
	GetCurrentTarget() EventTarget
	GetDefaultPrevented() bool
	GetError() Value
	GetEventPhase() int
	GetFilename() string
	InitEvent(args ...interface{})
	GetIsTrusted() bool
	GetLineno() int
	GetMessage() string
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
type ErrorEvent struct {
	Value
	Event
}

func JSValueToErrorEvent(val js.Value) ErrorEvent { return ErrorEvent{Value: Value{Value: val}} }
func (v Value) AsErrorEvent() ErrorEvent          { return ErrorEvent{Value: v} }
func (e ErrorEvent) GetColno() int {
	val := e.Get("colno")
	return val.Int()
}
func (e ErrorEvent) GetError() Value {
	val := e.Get("error")
	return val
}
func (e ErrorEvent) GetFilename() string {
	val := e.Get("filename")
	return val.String()
}
func (e ErrorEvent) GetLineno() int {
	val := e.Get("lineno")
	return val.Int()
}
func (e ErrorEvent) GetMessage() string {
	val := e.Get("message")
	return val.String()
}
