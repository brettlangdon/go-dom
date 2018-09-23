// Code generated DO NOT EDIT
// datatransfer.go
package dom

import "syscall/js"

type DataTransferIFace interface {
	ClearData(args ...interface{})
	GetDropEffect() string
	SetDropEffect(string)
	GetEffectAllowed() string
	SetEffectAllowed(string)
	GetFiles() FileList
	GetData(args ...interface{}) string
	GetItems() DataTransferItemList
	SetData(args ...interface{})
	SetDragImage(args ...interface{})
	GetTypes()
}
type DataTransfer struct {
	Value
}

func JSValueToDataTransfer(val js.Value) DataTransfer { return DataTransfer{Value: Value{Value: val}} }
func (v Value) AsDataTransfer() DataTransfer          { return DataTransfer{Value: v} }
func (d DataTransfer) ClearData(args ...interface{}) {
	d.Call("clearData", args...)
}
func (d DataTransfer) GetDropEffect() string {
	val := d.Get("dropEffect")
	return val.String()
}
func (d DataTransfer) SetDropEffect(val string) {
	d.Set("dropEffect", val)
}
func (d DataTransfer) GetEffectAllowed() string {
	val := d.Get("effectAllowed")
	return val.String()
}
func (d DataTransfer) SetEffectAllowed(val string) {
	d.Set("effectAllowed", val)
}
func (d DataTransfer) GetFiles() FileList {
	val := d.Get("files")
	return JSValueToFileList(val.JSValue())
}
func (d DataTransfer) GetData(args ...interface{}) string {
	val := d.Call("getData", args...)
	return val.String()
}
func (d DataTransfer) GetItems() DataTransferItemList {
	val := d.Get("items")
	return JSValueToDataTransferItemList(val.JSValue())
}
func (d DataTransfer) SetData(args ...interface{}) {
	d.Call("setData", args...)
}
func (d DataTransfer) SetDragImage(args ...interface{}) {
	d.Call("setDragImage", args...)
}
func (d DataTransfer) GetTypes() Value {
	val := d.Get("types")
	return val
}
