// Code generated DO NOT EDIT
// pluginarray.go
package dom

import "syscall/js"

type PluginArrayIFace interface {
	Item(args ...interface{}) Plugin
	GetLength() float64
	NamedItem(args ...interface{}) Plugin
	Refresh(args ...interface{})
}
type PluginArray struct {
	Value
}

func jsValueToPluginArray(val js.Value) PluginArray { return PluginArray{Value: Value{Value: val}} }
func (v Value) AsPluginArray() PluginArray          { return PluginArray{Value: v} }
func (p PluginArray) Item(args ...interface{}) Plugin {
	val := p.Call("item", args...)
	return jsValueToPlugin(val.JSValue())
}
func (p PluginArray) GetLength() float64 {
	val := p.Get("length")
	return val.Float()
}
func (p PluginArray) NamedItem(args ...interface{}) Plugin {
	val := p.Call("namedItem", args...)
	return jsValueToPlugin(val.JSValue())
}
func (p PluginArray) Refresh(args ...interface{}) {
	p.Call("refresh", args...)
}
