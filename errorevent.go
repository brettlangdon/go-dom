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
}

func JSValueToErrorEvent(val js.Value) ErrorEvent { return ErrorEvent{Value: JSValueToValue(val)} }
func (v Value) AsErrorEvent() ErrorEvent          { return ErrorEvent{Value: v} }
func NewErrorEvent(args ...interface{}) ErrorEvent {
	return ErrorEvent{Value: JSValueToValue(js.Global().Get("ErrorEvent").New(args...))}
}
func (e ErrorEvent) GetBubbles() bool {
	val := e.Get("bubbles")
	return val.Bool()
}
func (e ErrorEvent) GetCancelBubble() bool {
	val := e.Get("cancelBubble")
	return val.Bool()
}
func (e ErrorEvent) SetCancelBubble(val bool) {
	e.Set("cancelBubble", val)
}
func (e ErrorEvent) GetCancelable() bool {
	val := e.Get("cancelable")
	return val.Bool()
}
func (e ErrorEvent) GetColno() int {
	val := e.Get("colno")
	return val.Int()
}
func (e ErrorEvent) GetComposed() bool {
	val := e.Get("composed")
	return val.Bool()
}
func (e ErrorEvent) ComposedPath(args ...interface{}) {
	e.Call("composedPath", args...)
}
func (e ErrorEvent) GetCurrentTarget() EventTarget {
	val := e.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (e ErrorEvent) GetDefaultPrevented() bool {
	val := e.Get("defaultPrevented")
	return val.Bool()
}
func (e ErrorEvent) GetError() Value {
	val := e.Get("error")
	return val
}
func (e ErrorEvent) GetEventPhase() int {
	val := e.Get("eventPhase")
	return val.Int()
}
func (e ErrorEvent) GetFilename() string {
	val := e.Get("filename")
	return val.String()
}
func (e ErrorEvent) InitEvent(args ...interface{}) {
	e.Call("initEvent", args...)
}
func (e ErrorEvent) GetIsTrusted() bool {
	val := e.Get("isTrusted")
	return val.Bool()
}
func (e ErrorEvent) GetLineno() int {
	val := e.Get("lineno")
	return val.Int()
}
func (e ErrorEvent) GetMessage() string {
	val := e.Get("message")
	return val.String()
}
func (e ErrorEvent) PreventDefault(args ...interface{}) {
	e.Call("preventDefault", args...)
}
func (e ErrorEvent) GetReturnValue() bool {
	val := e.Get("returnValue")
	return val.Bool()
}
func (e ErrorEvent) SetReturnValue(val bool) {
	e.Set("returnValue", val)
}
func (e ErrorEvent) GetSrcElement() EventTarget {
	val := e.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (e ErrorEvent) StopImmediatePropagation(args ...interface{}) {
	e.Call("stopImmediatePropagation", args...)
}
func (e ErrorEvent) StopPropagation(args ...interface{}) {
	e.Call("stopPropagation", args...)
}
func (e ErrorEvent) GetTarget() EventTarget {
	val := e.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (e ErrorEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := e.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (e ErrorEvent) GetType() string {
	val := e.Get("type")
	return val.String()
}
