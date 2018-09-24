// Code generated DO NOT EDIT
// example.go
package dom

import "syscall/js"

type ExampleIFace interface {
}
type Example struct {
	Value
}

func JSValueToExample(val js.Value) Example { return Example{Value: JSValueToValue(val)} }
func (v Value) AsExample() Example          { return Example{Value: v} }
func NewExample(args ...interface{}) Example {
	return Example{Value: JSValueToValue(js.Global().Get("Example").New(args...))}
}
