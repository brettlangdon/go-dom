// Code generated DO NOT EDIT
// headers.go
package dom

import "syscall/js"

type HeadersIFace interface {
	Append(args ...interface{})
	Delete(args ...interface{})
	Get(args ...interface{}) []byte
	Has(args ...interface{}) bool
	Set(args ...interface{})
}
type Headers struct {
	Value
}

func JSValueToHeaders(val js.Value) Headers { return Headers{Value: JSValueToValue(val)} }
func (v Value) AsHeaders() Headers          { return Headers{Value: v} }
func NewHeaders(args ...interface{}) Headers {
	return Headers{Value: JSValueToValue(js.Global().Get("Headers").New(args...))}
}
func (h Headers) Append(args ...interface{}) {
	h.Call("append", args...)
}
func (h Headers) Delete(args ...interface{}) {
	h.Call("delete", args...)
}
func (h Headers) Get(args ...interface{}) []byte {
	val := h.Call("get", args...)
	return []byte(val.String())
}
func (h Headers) Has(args ...interface{}) bool {
	val := h.Call("has", args...)
	return val.Bool()
}
func (h Headers) Set(args ...interface{}) {
	h.Call("set", args...)
}
