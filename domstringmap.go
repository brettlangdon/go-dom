// Code generated DO NOT EDIT
// domstringmap.go
package dom

import "syscall/js"

type DOMStringMapIFace interface {
}
type DOMStringMap struct {
	Value
}

func JSValueToDOMStringMap(val js.Value) DOMStringMap { return DOMStringMap{Value: JSValueToValue(val)} }
func (v Value) AsDOMStringMap() DOMStringMap          { return DOMStringMap{Value: v} }
func NewDOMStringMap(args ...interface{}) DOMStringMap {
	return DOMStringMap{Value: JSValueToValue(js.Global().Get("DOMStringMap").New(args...))}
}
