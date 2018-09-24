// Code generated DO NOT EDIT
// domparser.go
package dom

import "syscall/js"

type DOMParserIFace interface {
	ParseFromString(args ...interface{}) Document
}
type DOMParser struct {
	Value
}

func JSValueToDOMParser(val js.Value) DOMParser { return DOMParser{Value: Value{Value: val}} }
func (v Value) AsDOMParser() DOMParser          { return DOMParser{Value: v} }
func (d DOMParser) ParseFromString(args ...interface{}) Document {
	val := d.Call("parseFromString", args...)
	return JSValueToDocument(val.JSValue())
}
