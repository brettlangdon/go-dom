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

func JSValueToDragEvent(val js.Value) DragEvent { return DragEvent{Value: JSValueToValue(val)} }
func (v Value) AsDragEvent() DragEvent          { return DragEvent{Value: v} }
func NewDragEvent(args ...interface{}) DragEvent {
	return DragEvent{Value: JSValueToValue(js.Global().Get("DragEvent").New(args...))}
}
func (d DragEvent) GetDataTransfer() DataTransfer {
	val := d.Get("dataTransfer")
	return JSValueToDataTransfer(val.JSValue())
}
