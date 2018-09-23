// Code generated DO NOT EDIT
// errorevent.go
package dom

import "syscall/js"

type ErrorEventIFace interface {
	GetBubbles() bool
	GetCancelBubble() bool
	SetCancelBubble(bool)
	GetCancelable() bool
	GetColno() float64
	GetComposed() bool
	ComposedPath(args ...interface{})
	GetCurrentTarget() EventTarget
	GetDefaultPrevented() bool
	GetError() Value
	GetEventPhase() int
	GetFilename() string
	InitEvent(args ...interface{})
	GetIsTrusted() bool
	GetLineno() float64
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

func jsValueToErrorEvent(val js.Value) ErrorEvent { return ErrorEvent{Value: Value{Value: val}} }
func (v Value) AsErrorEvent() ErrorEvent          { return ErrorEvent{Value: v} }
func (e ErrorEvent) GetColno() float64 {
	val := e.Get("colno")
	return val.Float()
}
func (e ErrorEvent) GetError() Value {
	val := e.Get("error")
	return val
}
func (e ErrorEvent) GetFilename() string {
	val := e.Get("filename")
	return val.String()
}
func (e ErrorEvent) GetLineno() float64 {
	val := e.Get("lineno")
	return val.Float()
}
func (e ErrorEvent) GetMessage() string {
	val := e.Get("message")
	return val.String()
}
