// Code generated DO NOT EDIT
// filelist.go
package dom

import "syscall/js"

type FileListIFace interface {
	Item(args ...interface{}) File
	GetLength() float64
}
type FileList struct {
	Value
}

func JSValueToFileList(val js.Value) FileList { return FileList{Value: Value{Value: val}} }
func (v Value) AsFileList() FileList          { return FileList{Value: v} }
func (f FileList) Item(args ...interface{}) File {
	val := f.Call("item", args...)
	return JSValueToFile(val.JSValue())
}
func (f FileList) GetLength() float64 {
	val := f.Get("length")
	return val.Float()
}
