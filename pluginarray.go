// Code generated DO NOT EDIT
// pluginarray.go
package dom

import "syscall/js"

type PluginArrayIFace interface {
	Item(args ...interface{}) Plugin
	GetLength() int
	NamedItem(args ...interface{}) Plugin
	Refresh(args ...interface{})
}
type PluginArray struct {
	Value
}

func JSValueToPluginArray(val js.Value) PluginArray { return PluginArray{Value: JSValueToValue(val)} }
func (v Value) AsPluginArray() PluginArray          { return PluginArray{Value: v} }
func NewPluginArray(args ...interface{}) PluginArray {
	return PluginArray{Value: JSValueToValue(js.Global().Get("PluginArray").New(args...))}
}
func (p PluginArray) Item(args ...interface{}) Plugin {
	val := p.Call("item", args...)
	return JSValueToPlugin(val.JSValue())
}
func (p PluginArray) GetLength() int {
	val := p.Get("length")
	return val.Int()
}
func (p PluginArray) NamedItem(args ...interface{}) Plugin {
	val := p.Call("namedItem", args...)
	return JSValueToPlugin(val.JSValue())
}
func (p PluginArray) Refresh(args ...interface{}) {
	p.Call("refresh", args...)
}
