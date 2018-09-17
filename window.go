// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Window struct {
	Value
}

func NewWindow(v js.Value) *Window {
	val := Value{Value: v}
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return val.ToWindow()
}
func (v Value) ToWindow() *Window { return &Window{Value: v} }
