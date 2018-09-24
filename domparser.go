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

func JSValueToDOMParser(val js.Value) DOMParser { return DOMParser{Value: JSValueToValue(val)} }
func (v Value) AsDOMParser() DOMParser          { return DOMParser{Value: v} }
func NewDOMParser(args ...interface{}) DOMParser {
	return DOMParser{Value: JSValueToValue(js.Global().Get("DOMParser").New(args...))}
}
func (d DOMParser) ParseFromString(args ...interface{}) Document {
	val := d.Call("parseFromString", args...)
	return JSValueToDocument(val.JSValue())
}
