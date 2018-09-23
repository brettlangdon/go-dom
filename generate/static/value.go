package dom

import "syscall/js"

type JSValue interface {
	JSValue() js.Value
}

type Value struct {
	js.Value
}

func jsValueToValue(val js.Value) Value { return Value{Value: val} }

func (v Value) JSValue() js.Value  { return v.Value }
func (v Value) Get(p string) Value { return jsValueToValue(v.Value.Get(p)) }
func (v Value) Index(i int) Value  { return jsValueToValue(v.Value.Index(i)) }
func (v Value) Call(m string, args ...interface{}) Value {
	return jsValueToValue(v.Value.Call(m, args...))
}
func (v Value) Invoke(args ...interface{}) Value {
	return jsValueToValue(v.Value.Invoke(args...))
}
func (v Value) New(args ...interface{}) Value {
	return jsValueToValue(v.Value.New(args...))
}
