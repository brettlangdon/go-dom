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
	GetLoaded() float64
	PreventDefault(args ...interface{})
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetTotal() float64
	GetType() string
}
type ProgressEvent struct {
	Value
	Event
}

func jsValueToProgressEvent(val js.Value) ProgressEvent {
	return ProgressEvent{Value: Value{Value: val}}
}
func (v Value) AsProgressEvent() ProgressEvent { return ProgressEvent{Value: v} }
func (p ProgressEvent) GetLengthComputable() bool {
	val := p.Get("lengthComputable")
	return val.Bool()
}
func (p ProgressEvent) GetLoaded() float64 {
	val := p.Get("loaded")
	return val.Float()
}
func (p ProgressEvent) GetTotal() float64 {
	val := p.Get("total")
	return val.Float()
}
