// Code generated DO NOT EDIT
// file.go
package dom

import "syscall/js"

type FileIFace interface {
	GetLastModified() float64
	GetName() string
	GetSize() float64
	Slice(args ...interface{}) Blob
	GetType() string
}
type File struct {
	Value
	Blob
}

func jsValueToFile(val js.Value) File { return File{Value: Value{Value: val}} }
func (v Value) AsFile() File          { return File{Value: v} }
func (f File) GetLastModified() float64 {
	val := f.Get("lastModified")
	return val.Float()
}
func (f File) GetName() string {
	val := f.Get("name")
	return val.String()
}
