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
	Event
}

func JSValueToStorageEvent(val js.Value) StorageEvent { return StorageEvent{Value: JSValueToValue(val)} }
func (v Value) AsStorageEvent() StorageEvent          { return StorageEvent{Value: v} }
func NewStorageEvent(args ...interface{}) StorageEvent {
	return StorageEvent{Value: JSValueToValue(js.Global().Get("StorageEvent").New(args...))}
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
func (s StorageEvent) GetStorageArea() Storage {
	val := s.Get("storageArea")
	return JSValueToStorage(val.JSValue())
}
func (s StorageEvent) GetUrl() string {
	val := s.Get("url")
	return val.String()
}
