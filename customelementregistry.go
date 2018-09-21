// DO NOT EDIT - generated file
package dom

import "syscall/js"

type CustomElementRegistry struct {
	Value
}

func NewCustomElementRegistry(v js.Value) *CustomElementRegistry {
	val := Value{Value: v}
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return val.ToCustomElementRegistry()
}
func (v Value) ToCustomElementRegistry() *CustomElementRegistry {
	return &CustomElementRegistry{Value: v}
}
func (c *CustomElementRegistry) Define(name string, constructor interface{}) Value {
	val := Value{Value: c.Call("define", ToJSValue(name), ToJSValue(constructor))}
	return val
}
func (c *CustomElementRegistry) Get(name string) Value {
	val := Value{Value: c.Call("get", ToJSValue(name))}
	return val
}
func (c *CustomElementRegistry) WhenDefined(name string) *Promise {
	val := Value{Value: c.Call("whenDefined", ToJSValue(name))}
	return NewPromise(val.JSValue())
}
