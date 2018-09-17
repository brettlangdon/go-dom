// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Event struct {
	js.Value
}

func (e *Event) JSValue() js.Value { return e.Value }
func (e *Event) PreventDefault() {
	e.Call("preventDefault")
}
func (e *Event) StopPropagation() {
	e.Call("stopPropagation")
}
func (e *Event) StopImmediatePropagation() {
	e.Call("stopImmediatePropagation")
}
func (e *Event) GetCurrentTarget() js.Value {
	val := e.Get("currentTarget")
	return val
}
func (e *Event) GetTarget() js.Value {
	val := e.Get("target")
	return val
}
func (e *Event) GetType() string {
	val := e.Get("type")
	return val.String()
}
func (e *Event) GetSrcElement() js.Value {
	val := e.Get("srcElement")
	return val
}
