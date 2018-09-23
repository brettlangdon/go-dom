// Code generated DO NOT EDIT
// datatransferitemlist.go
package dom

import "syscall/js"

type DataTransferItemListIFace interface {
	Add(args ...interface{}) DataTransferItem
	AddWithArgs(args ...interface{}) DataTransferItem
	Clear(args ...interface{})
	GetLength() float64
	Remove(args ...interface{})
}
type DataTransferItemList struct {
	Value
}

func JSValueToDataTransferItemList(val js.Value) DataTransferItemList {
	return DataTransferItemList{Value: Value{Value: val}}
}
func (v Value) AsDataTransferItemList() DataTransferItemList { return DataTransferItemList{Value: v} }
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
func (d DataTransferItemList) GetLength() float64 {
	val := d.Get("length")
	return val.Float()
}
func (d DataTransferItemList) Remove(args ...interface{}) {
	d.Call("remove", args...)
}
