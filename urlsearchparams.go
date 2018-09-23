// Code generated DO NOT EDIT
// urlsearchparams.go
package dom

import "syscall/js"

type URLSearchParamsIFace interface {
	Append(args ...interface{})
	Delete(args ...interface{})
	Get(args ...interface{}) string
	GetAll(args ...interface{})
	Has(args ...interface{}) bool
	Set(args ...interface{})
	Sort(args ...interface{})
}
type URLSearchParams struct {
	Value
}

func jsValueToURLSearchParams(val js.Value) URLSearchParams {
	return URLSearchParams{Value: Value{Value: val}}
}
func (v Value) AsURLSearchParams() URLSearchParams { return URLSearchParams{Value: v} }
func (u URLSearchParams) Append(args ...interface{}) {
	u.Call("append", args...)
}
func (u URLSearchParams) Delete(args ...interface{}) {
	u.Call("delete", args...)
}
func (u URLSearchParams) Get(args ...interface{}) string {
	val := u.Call("get", args...)
	return val.String()
}
func (u URLSearchParams) GetAll(args ...interface{}) {
	u.Call("getAll", args...)
}
func (u URLSearchParams) Has(args ...interface{}) bool {
	val := u.Call("has", args...)
	return val.Bool()
}
func (u URLSearchParams) Set(args ...interface{}) {
	u.Call("set", args...)
}
func (u URLSearchParams) Sort(args ...interface{}) {
	u.Call("sort", args...)
}
