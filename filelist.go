// Code generated DO NOT EDIT
// filelist.go
package dom

import "syscall/js"

type FileListIFace interface {
	Item(args ...interface{}) File
	GetLength() int
}
type FileList struct {
	Value
}

func JSValueToFileList(val js.Value) FileList { return FileList{Value: JSValueToValue(val)} }
func (v Value) AsFileList() FileList          { return FileList{Value: v} }
func NewFileList(args ...interface{}) FileList {
	return FileList{Value: JSValueToValue(js.Global().Get("FileList").New(args...))}
}
func (f FileList) Item(args ...interface{}) File {
	val := f.Call("item", args...)
	return JSValueToFile(val.JSValue())
}
func (f FileList) GetLength() int {
	val := f.Get("length")
	return val.Int()
}
