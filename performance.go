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
	EventTarget
}

func jsValueToPerformance(val js.Value) Performance { return Performance{Value: Value{Value: val}} }
func (v Value) AsPerformance() Performance          { return Performance{Value: v} }
func (p Performance) Now(args ...interface{}) DOMHighResTimeStamp {
	val := p.Call("now", args...)
	return jsValueToDOMHighResTimeStamp(val.JSValue())
}
func (p Performance) GetTimeOrigin() DOMHighResTimeStamp {
	val := p.Get("timeOrigin")
	return jsValueToDOMHighResTimeStamp(val.JSValue())
}
func (p Performance) ToJSON(args ...interface{}) Value {
	val := p.Call("toJSON", args...)
	return val
}
