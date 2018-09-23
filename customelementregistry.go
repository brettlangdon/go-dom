// Code generated DO NOT EDIT
// customelementregistry.go
package dom

import "syscall/js"

type CustomElementRegistryIFace interface {
	Define(args ...interface{})
	Get(args ...interface{}) Value
	Upgrade(args ...interface{})
	WhenDefined(args ...interface{})
}
type CustomElementRegistry struct {
	Value
}

func JSValueToCustomElementRegistry(val js.Value) CustomElementRegistry {
	return CustomElementRegistry{Value: Value{Value: val}}
}
func (v Value) AsCustomElementRegistry() CustomElementRegistry { return CustomElementRegistry{Value: v} }
func (c CustomElementRegistry) Define(args ...interface{}) {
	c.Call("define", args...)
}
func (c CustomElementRegistry) Get(args ...interface{}) Value {
	val := c.Call("get", args...)
	return val
}
func (c CustomElementRegistry) Upgrade(args ...interface{}) {
	c.Call("upgrade", args...)
}
func (c CustomElementRegistry) WhenDefined(args ...interface{}) {
	c.Call("whenDefined", args...)
}
