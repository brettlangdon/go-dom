package dom

import "syscall/js"

type Uint8Array struct {
	js.TypedArray
}

func jsValueToUint8Array(val js.Value) Uint8Array {
	return Uint8Array{TypedArray: js.TypedArray{Value: val}}
}
func (u Uint8Array) JSValue() js.Value { return u.TypedArray.Value }
