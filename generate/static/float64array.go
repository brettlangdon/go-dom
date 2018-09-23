package dom

import "syscall/js"

type Float64Array struct {
	js.TypedArray
}

func JSValueToFloat64Array(val js.Value) Float64Array {
	return Float64Array{TypedArray: js.TypedArray{Value: val}}
}
func (u Float64Array) JSValue() js.Value { return u.TypedArray.Value }
