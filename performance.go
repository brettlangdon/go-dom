// Code generated DO NOT EDIT
// performance.go
package dom

import "syscall/js"

type PerformanceIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	Now(args ...interface{}) DOMHighResTimeStamp
	RemoveEventListener(args ...interface{})
	GetTimeOrigin() DOMHighResTimeStamp
	ToJSON(args ...interface{}) Value
}
type Performance struct {
	Value
}

func JSValueToPerformance(val js.Value) Performance { return Performance{Value: JSValueToValue(val)} }
func (v Value) AsPerformance() Performance          { return Performance{Value: v} }
func NewPerformance(args ...interface{}) Performance {
	return Performance{Value: JSValueToValue(js.Global().Get("Performance").New(args...))}
}
func (p Performance) AddEventListener(args ...interface{}) {
	p.Call("addEventListener", args...)
}
func (p Performance) DispatchEvent(args ...interface{}) bool {
	val := p.Call("dispatchEvent", args...)
	return val.Bool()
}
func (p Performance) Now(args ...interface{}) DOMHighResTimeStamp {
	val := p.Call("now", args...)
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (p Performance) RemoveEventListener(args ...interface{}) {
	p.Call("removeEventListener", args...)
}
func (p Performance) GetTimeOrigin() DOMHighResTimeStamp {
	val := p.Get("timeOrigin")
	return JSValueToDOMHighResTimeStamp(val.JSValue())
}
func (p Performance) ToJSON(args ...interface{}) Value {
	val := p.Call("toJSON", args...)
	return val
}
