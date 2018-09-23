// Code generated DO NOT EDIT
// datatransferitem.go
package dom

import "syscall/js"

type DataTransferItemIFace interface {
	GetAsFile(args ...interface{}) File
	GetAsString(args ...interface{})
	GetKind() string
	GetType() string
}
type DataTransferItem struct {
	Value
}

func JSValueToDataTransferItem(val js.Value) DataTransferItem {
	return DataTransferItem{Value: Value{Value: val}}
}
func (v Value) AsDataTransferItem() DataTransferItem { return DataTransferItem{Value: v} }
func (d DataTransferItem) GetAsFile(args ...interface{}) File {
	val := d.Call("getAsFile", args...)
	return JSValueToFile(val.JSValue())
}
func (d DataTransferItem) GetAsString(args ...interface{}) {
	d.Call("getAsString", args...)
}
func (d DataTransferItem) GetKind() string {
	val := d.Get("kind")
	return val.String()
}
func (d DataTransferItem) GetType() string {
	val := d.Get("type")
	return val.String()
}
