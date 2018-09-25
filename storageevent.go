// Code generated DO NOT EDIT
// storageevent.go
package dom

import "syscall/js"

type StorageEventIFace interface {
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
	GetKey() string
	GetNewValue() string
	GetOldValue() string
	PreventDefault(args ...interface{})
	GetReturnValue() bool
	SetReturnValue(bool)
	GetSrcElement() EventTarget
	StopImmediatePropagation(args ...interface{})
	StopPropagation(args ...interface{})
	GetStorageArea() Storage
	GetTarget() EventTarget
	GetTimeStamp() DOMHighResTimeStamp
	GetType() string
	GetUrl() string
}
type StorageEvent struct {
	Value
}

func JSValueToStorageEvent(val js.Value) StorageEvent { return StorageEvent{Value: JSValueToValue(val)} }
func (v Value) AsStorageEvent() StorageEvent          { return StorageEvent{Value: v} }
func NewStorageEvent(args ...interface{}) StorageEvent {
	return StorageEvent{Value: JSValueToValue(js.Global().Get("StorageEvent").New(args...))}
}
func (s StorageEvent) GetBubbles() bool {
	val := s.Get("bubbles")
	return val.Bool()
}
func (s StorageEvent) GetCancelBubble() bool {
	val := s.Get("cancelBubble")
	return val.Bool()
}
func (s StorageEvent) SetCancelBubble(val bool) {
	s.Set("cancelBubble", val)
}
func (s StorageEvent) GetCancelable() bool {
	val := s.Get("cancelable")
	return val.Bool()
}
func (s StorageEvent) GetComposed() bool {
	val := s.Get("composed")
	return val.Bool()
}
func (s StorageEvent) ComposedPath(args ...interface{}) {
	s.Call("composedPath", args...)
}
func (s StorageEvent) GetCurrentTarget() EventTarget {
	val := s.Get("currentTarget")
	return JSValueToEventTarget(val.JSValue())
}
func (s StorageEvent) GetDefaultPrevented() bool {
	val := s.Get("defaultPrevented")
	return val.Bool()
}
func (s StorageEvent) GetEventPhase() int {
	val := s.Get("eventPhase")
	return val.Int()
}
func (s StorageEvent) InitEvent(args ...interface{}) {
	s.Call("initEvent", args...)
}
func (s StorageEvent) GetIsTrusted() bool {
	val := s.Get("isTrusted")
	return val.Bool()
}
func (s StorageEvent) GetKey() string {
	val := s.Get("key")
	return val.String()
}
func (s StorageEvent) GetNewValue() string {
	val := s.Get("newValue")
	return val.String()
}
func (s StorageEvent) GetOldValue() string {
	val := s.Get("oldValue")
	return val.String()
}
func (s StorageEvent) PreventDefault(args ...interface{}) {
	s.Call("preventDefault", args...)
}
func (s StorageEvent) GetReturnValue() bool {
	val := s.Get("returnValue")
	return val.Bool()
}
func (s StorageEvent) SetReturnValue(val bool) {
	s.Set("returnValue", val)
}
func (s StorageEvent) GetSrcElement() EventTarget {
	val := s.Get("srcElement")
	return JSValueToEventTarget(val.JSValue())
}
func (s StorageEvent) StopImmediatePropagation(args ...interface{}) {
	s.Call("stopImmediatePropagation", args...)
}
func (s StorageEvent) StopPropagation(args ...interface{}) {
	s.Call("stopPropagation", args...)
}
func (s StorageEvent) GetStorageArea() Storage {
	val := s.Get("storageArea")
	return JSValueToStorage(val.JSValue())
}
func (s StorageEvent) GetTarget() EventTarget {
	val := s.Get("target")
	return JSValueToEventTarget(val.JSValue())
}
func (s StorageEvent) GetTimeStamp() DOMHighResTimeStamp {
	val := s.Get("timeStamp")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (s StorageEvent) GetType() string {
	val := s.Get("type")
	return val.String()
}
func (s StorageEvent) GetUrl() string {
	val := s.Get("url")
	return val.String()
}
