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
	Blob
}

func JSValueToFile(val js.Value) File { return File{Value: Value{Value: val}} }
func (v Value) AsFile() File          { return File{Value: v} }
func (f File) GetLastModified() int {
	val := f.Get("lastModified")
	return val.Int()
}
func (f File) GetName() string {
	val := f.Get("name")
	return val.String()
}
