// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Console struct {
	Value
}

func NewConsole(v js.Value) *Console {
	val := Value{Value: v}
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return val.ToConsole()
}
func (v Value) ToConsole() *Console { return &Console{Value: v} }
func (c *Console) Log(v ...interface{}) Value {
	vVaridic := make([]interface{}, 0)
	for _, a := range v {
		vVaridic = append(vVaridic, ToJSValue(a))
	}
	val := Value{Value: c.Call("log", vVaridic...)}
	return val
}
func (c *Console) Error(v ...interface{}) Value {
	vVaridic := make([]interface{}, 0)
	for _, a := range v {
		vVaridic = append(vVaridic, ToJSValue(a))
	}
	val := Value{Value: c.Call("error", vVaridic...)}
	return val
}
func (c *Console) Dir(v JSValue) Value {
	val := Value{Value: c.Call("dir", ToJSValue(v))}
	return val
}
