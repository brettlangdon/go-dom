// Code generated DO NOT EDIT
// dragevent.go
package dom

import "syscall/js"

type DragEventIFace interface {
	GetDataTransfer() DataTransfer
}
type DragEvent struct {
	Value
}

func jsValueToDragEvent(val js.Value) DragEvent { return DragEvent{Value: Value{Value: val}} }
func (v Value) AsDragEvent() DragEvent          { return DragEvent{Value: v} }
func (d DragEvent) GetDataTransfer() DataTransfer {
	val := d.Get("dataTransfer")
	return jsValueToDataTransfer(val.JSValue())
}
