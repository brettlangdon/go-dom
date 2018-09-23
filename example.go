// Code generated DO NOT EDIT
// example.go
package dom

import "syscall/js"

type ExampleIFace interface {
}
type Example struct {
	Value
}

func JSValueToExample(val js.Value) Example { return Example{Value: Value{Value: val}} }
func (v Value) AsExample() Example          { return Example{Value: v} }
