// Code generated DO NOT EDIT
// domstringmap.go
package dom

import "syscall/js"

type DOMStringMapIFace interface {
}
type DOMStringMap struct {
	Value
}

func JSValueToDOMStringMap(val js.Value) DOMStringMap { return DOMStringMap{Value: Value{Value: val}} }
func (v Value) AsDOMStringMap() DOMStringMap          { return DOMStringMap{Value: v} }
