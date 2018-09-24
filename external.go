// Code generated DO NOT EDIT
// external.go
package dom

import "syscall/js"

type ExternalIFace interface {
	AddSearchProvider(args ...interface{})
	IsSearchProviderInstalled(args ...interface{})
}
type External struct {
	Value
}

func JSValueToExternal(val js.Value) External { return External{Value: JSValueToValue(val)} }
func (v Value) AsExternal() External          { return External{Value: v} }
func NewExternal(args ...interface{}) External {
	return External{Value: JSValueToValue(js.Global().Get("External").New(args...))}
}
func (e External) AddSearchProvider(args ...interface{}) {
	e.Call("AddSearchProvider", args...)
}
func (e External) IsSearchProviderInstalled(args ...interface{}) {
	e.Call("IsSearchProviderInstalled", args...)
}
