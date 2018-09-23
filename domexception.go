// Code generated DO NOT EDIT
// domexception.go
package dom

import "syscall/js"

type DOMExceptionIFace interface {
	GetCode() int
	GetMessage() string
	GetName() string
}
type DOMException struct {
	Value
}

func jsValueToDOMException(val js.Value) DOMException { return DOMException{Value: Value{Value: val}} }
func (v Value) AsDOMException() DOMException          { return DOMException{Value: v} }
func (d DOMException) GetCode() int {
	val := d.Get("code")
	return val.Int()
}
func (d DOMException) GetMessage() string {
	val := d.Get("message")
	return val.String()
}
func (d DOMException) GetName() string {
	val := d.Get("name")
	return val.String()
}
