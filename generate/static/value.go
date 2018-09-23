package dom

import "syscall/js"

type JSValue interface {
	JSValue() js.Value
}

type Value struct {
	js.Value
}

func JSValueToValue(val js.Value) Value { return Value{Value: val} }

func (v Value) JSValue() js.Value  { return v.Value }
func (v Value) Get(p string) Value { return JSValueToValue(v.Value.Get(p)) }
func (v Value) Index(i int) Value  { return JSValueToValue(v.Value.Index(i)) }
func (v Value) Call(m string, args ...interface{}) Value {
	return JSValueToValue(v.Value.Call(m, args...))
}
func (v Value) Invoke(args ...interface{}) Value {
	return JSValueToValue(v.Value.Invoke(args...))
}
func (v Value) New(args ...interface{}) Value {
	return JSValueToValue(v.Value.New(args...))
}
