// Code generated DO NOT EDIT
// formdata.go
package dom

import "syscall/js"

type FormDataIFace interface {
	Append(args ...interface{})
	AppendWithArgs(args ...interface{})
	Delete(args ...interface{})
	Get(args ...interface{}) FormDataEntryValue
	GetAll(args ...interface{})
	Has(args ...interface{}) bool
	Set(args ...interface{})
	SetWithArgs(args ...interface{})
}
type FormData struct {
	Value
}

func jsValueToFormData(val js.Value) FormData { return FormData{Value: Value{Value: val}} }
func (v Value) AsFormData() FormData          { return FormData{Value: v} }
func (f FormData) Append(args ...interface{}) {
	f.Call("append", args...)
}
func (f FormData) AppendWithArgs(args ...interface{}) {
	f.Call("appendWithArgs", args...)
}
func (f FormData) Delete(args ...interface{}) {
	f.Call("delete", args...)
}
func (f FormData) Get(args ...interface{}) FormDataEntryValue {
	val := f.Call("get", args...)
	return jsValueToFormDataEntryValue(val.JSValue())
}
func (f FormData) GetAll(args ...interface{}) {
	f.Call("getAll", args...)
}
func (f FormData) Has(args ...interface{}) bool {
	val := f.Call("has", args...)
	return val.Bool()
}
func (f FormData) Set(args ...interface{}) {
	f.Call("set", args...)
}
func (f FormData) SetWithArgs(args ...interface{}) {
	f.Call("setWithArgs", args...)
}
