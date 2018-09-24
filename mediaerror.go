// Code generated DO NOT EDIT
// mediaerror.go
package dom

import "syscall/js"

type MediaErrorIFace interface {
	GetCode() int
	GetMessage() string
}
type MediaError struct {
	Value
}

func JSValueToMediaError(val js.Value) MediaError { return MediaError{Value: JSValueToValue(val)} }
func (v Value) AsMediaError() MediaError          { return MediaError{Value: v} }
func NewMediaError(args ...interface{}) MediaError {
	return MediaError{Value: JSValueToValue(js.Global().Get("MediaError").New(args...))}
}
func (m MediaError) GetCode() int {
	val := m.Get("code")
	return val.Int()
}
func (m MediaError) GetMessage() string {
	val := m.Get("message")
	return val.String()
}
