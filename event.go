// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Event struct {
	Value
}

func NewEvent(v js.Value) *Event {
	val := Value{Value: v}
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return val.ToEvent()
}
func (v Value) ToEvent() *Event { return &Event{Value: v} }
func (e *Event) PreventDefault() Value {
	val := Value{Value: e.Call("preventDefault")}
	return val
}
func (e *Event) StopPropagation() Value {
	val := Value{Value: e.Call("stopPropagation")}
	return val
}
func (e *Event) StopImmediatePropagation() Value {
	val := Value{Value: e.Call("stopImmediatePropagation")}
	return val
}
func (e *Event) GetCurrentTarget() Value {
	val := Value{Value: e.Get("currentTarget")}
	return val
}
func (e *Event) GetTarget() Value {
	val := Value{Value: e.Get("target")}
	return val
}
func (e *Event) GetType() string {
	val := Value{Value: e.Get("type")}
	return val.String()
}
func (e *Event) GetSrcElement() Value {
	val := Value{Value: e.Get("srcElement")}
	return val
}
