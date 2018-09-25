// Code generated DO NOT EDIT
// progressevent.go
package dom

import "syscall/js"

type ProgressEventIFace interface {
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
	GetLengthComputable() bool
	GetLoaded() int
	PreventDefault(args ...interface{})
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetTotal() int
	GetType() string
}
type ProgressEvent struct {
	Value
}

func JSValueToProgressEvent(val js.Value) ProgressEvent {
	return ProgressEvent{Value: JSValueToValue(val)}
}
func (v Value) AsProgressEvent() ProgressEvent { return ProgressEvent{Value: v} }
func NewProgressEvent(args ...interface{}) ProgressEvent {
	return ProgressEvent{Value: JSValueToValue(js.Global().Get("ProgressEvent").New(args...))}
}
func (p ProgressEvent) GetBubbles() bool {
	val := p.Get("bubbles")
	return val.Bool()
}
func (p ProgressEvent) GetCancelBubble() bool {
	val := p.Get("cancelBubble")
	return val.Bool()
}
func (p ProgressEvent) SetCancelBubble(val bool) {
	p.Set("cancelBubble", val)
}
func (p ProgressEvent) GetCancelable() bool {
	val := p.Get("cancelable")
	return val.Bool()
}
func (p ProgressEvent) GetComposed() bool {
	val := p.Get("composed")
	return val.Bool()
}
func (p ProgressEvent) ComposedPath(args ...interface{}) {
	p.Call("composedPath", args...)
}
func (p ProgressEvent) GetCurrentTarget() EventTarget {
	val := p.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (p ProgressEvent) GetDefaultPrevented() bool {
	val := p.Get("defaultPrevented")
	return val.Bool()
}
func (p ProgressEvent) GetEventPhase() int {
	val := p.Get("eventPhase")
	return val.Int()
}
func (p ProgressEvent) InitEvent(args ...interface{}) {
	p.Call("initEvent", args...)
}
func (p ProgressEvent) GetIsTrusted() bool {
	val := p.Get("isTrusted")
	return val.Bool()
}
func (p ProgressEvent) GetLengthComputable() bool {
	val := p.Get("lengthComputable")
	return val.Bool()
}
func (p ProgressEvent) GetLoaded() int {
	val := p.Get("loaded")
	return val.Int()
}
func (p ProgressEvent) PreventDefault(args ...interface{}) {
	p.Call("preventDefault", args...)
}
func (p ProgressEvent) GetReturnValue() bool {
	val := p.Get("returnValue")
	return val.Bool()
}
func (p ProgressEvent) SetReturnValue(val bool) {
	p.Set("returnValue", val)
}
func (p ProgressEvent) GetSrcElement() EventTarget {
	val := p.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (p ProgressEvent) StopImmediatePropagation(args ...interface{}) {
	p.Call("stopImmediatePropagation", args...)
}
func (p ProgressEvent) StopPropagation(args ...interface{}) {
	p.Call("stopPropagation", args...)
}
func (p ProgressEvent) GetTarget() EventTarget {
	val := p.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (p ProgressEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := p.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (p ProgressEvent) GetTotal() int {
	val := p.Get("total")
	return val.Int()
}
func (p ProgressEvent) GetType() string {
	val := p.Get("type")
	return val.String()
}
