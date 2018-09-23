package dom

import "syscall/js"

type Float32Array struct {
	js.TypedArray
}

func JSValueToFloat32Array(val js.Value) Float32Array {
	return Float32Array{TypedArray: js.TypedArray{Value: val}}
}
func (u Float32Array) JSValue() js.Value { return u.TypedArray.Value }
