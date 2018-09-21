// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Promise struct {
	Value
}

func NewPromise(v js.Value) *Promise {
	val := Value{Value: v}
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return val.ToPromise()
}
func (v Value) ToPromise() *Promise { return &Promise{Value: v} }
