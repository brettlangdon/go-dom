// Code generated DO NOT EDIT
// plugin.go
package dom

import "syscall/js"

type PluginIFace interface {
	GetDescription() string
	GetFilename() string
	Item(args ...interface{}) MimeType
	GetLength() int
	GetName() string
	NamedItem(args ...interface{}) MimeType
}
type Plugin struct {
	Value
}

func JSValueToPlugin(val js.Value) Plugin { return Plugin{Value: JSValueToValue(val)} }
func (v Value) AsPlugin() Plugin          { return Plugin{Value: v} }
func NewPlugin(args ...interface{}) Plugin {
	return Plugin{Value: JSValueToValue(js.Global().Get("Plugin").New(args...))}
}
func (p Plugin) GetDescription() string {
	val := p.Get("description")
	return val.String()
}
func (p Plugin) GetFilename() string {
	val := p.Get("filename")
	return val.String()
}
func (p Plugin) Item(args ...interface{}) MimeType {
	val := p.Call("item", args...)
	return JSValueToMimeType(val.JSValue())
}
func (p Plugin) GetLength() int {
	val := p.Get("length")
	return val.Int()
}
func (p Plugin) GetName() string {
	val := p.Get("name")
	return val.String()
}
func (p Plugin) NamedItem(args ...interface{}) MimeType {
	val := p.Call("namedItem", args...)
	return JSValueToMimeType(val.JSValue())
}
