// Code generated DO NOT EDIT
// datatransferitemlist.go
package dom

import "syscall/js"

type DataTransferItemListIFace interface {
	Add(args ...interface{}) DataTransferItem
	AddWithArgs(args ...interface{}) DataTransferItem
	Clear(args ...interface{})
	GetLength() int
	Remove(args ...interface{})
}
type DataTransferItemList struct {
	Value
}

func JSValueToDataTransferItemList(val js.Value) DataTransferItemList {
	return DataTransferItemList{Value: JSValueToValue(val)}
}
func (v Value) AsDataTransferItemList() DataTransferItemList { return DataTransferItemList{Value: v} }
func NewDataTransferItemList(args ...interface{}) DataTransferItemList {
	return DataTransferItemList{Value: JSValueToValue(js.Global().Get("DataTransferItemList").New(args...))}
}
func (d DataTransferItemList) Add(args ...interface{}) DataTransferItem {
	val := d.Call("add", args...)
	return JSValueToDataTransferItem(val.JSValue())
}
func (d DataTransferItemList) AddWithArgs(args ...interface{}) DataTransferItem {
	val := d.Call("addWithArgs", args...)
	return JSValueToDataTransferItem(val.JSValue())
}
func (d DataTransferItemList) Clear(args ...interface{}) {
	d.Call("clear", args...)
}
func (d DataTransferItemList) GetLength() int {
	val := d.Get("length")
	return val.Int()
}
func (d DataTransferItemList) Remove(args ...interface{}) {
	d.Call("remove", args...)
}
