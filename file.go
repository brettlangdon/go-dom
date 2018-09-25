// Code generated DO NOT EDIT
// file.go
package dom

import "syscall/js"

type FileIFace interface {
	GetLastModified() int
	GetName() string
	GetSize() int
	Slice(args ...interface{}) Blob
	GetType() string
}
type File struct {
	Value
}

func JSValueToFile(val js.Value) File { return File{Value: JSValueToValue(val)} }
func (v Value) AsFile() File          { return File{Value: v} }
func NewFile(args ...interface{}) File {
	return File{Value: JSValueToValue(js.Global().Get("File").New(args...))}
}
func (f File) GetLastModified() int {
	val := f.Get("lastModified")
	return val.Int()
}
func (f File) GetName() string {
	val := f.Get("name")
	return val.String()
}
func (f File) GetSize() int {
	val := f.Get("size")
	return val.Int()
}
func (f File) Slice(args ...interface{}) Blob {
	val := f.Call("slice", args...)
	return JSValueToBlob(val.JSValue())
}
func (f File) GetType() string {
	val := f.Get("type")
	return val.String()
}
