package dom

import "syscall/js"

type ArrayBuffer struct {
	js.TypedArray
}

func jsValueToArrayBuffer(val js.Value) ArrayBuffer {
	return ArrayBuffer{TypedArray: js.TypedArray{Value: val}}
}
func (u ArrayBuffer) JSValue() js.Value { return u.TypedArray.Value }
