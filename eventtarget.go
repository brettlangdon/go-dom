// Code generated DO NOT EDIT
// eventtarget.go
package dom

import "syscall/js"

type EventTargetIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	RemoveEventListener(args ...interface{})
}
type EventTarget struct {
	Value
}

func JSValueToEventTarget(val js.Value) EventTarget { return EventTarget{Value: Value{Value: val}} }
func (v Value) AsEventTarget() EventTarget          { return EventTarget{Value: v} }
func (e EventTarget) AddEventListener(args ...interface{}) {
	e.Call("addEventListener", args...)
}
func (e EventTarget) DispatchEvent(args ...interface{}) bool {
	val := e.Call("dispatchEvent", args...)
	return val.Bool()
}
func (e EventTarget) RemoveEventListener(args ...interface{}) {
	e.Call("removeEventListener", args...)
}
