package dom

import "syscall/js"

type Uint8ClampedArray struct {
	js.TypedArray
}

func JSValueToUint8ClampedArray(val js.Value) Uint8ClampedArray {
	return Uint8ClampedArray{TypedArray: js.TypedArray{Value: val}}
}
func (u Uint8ClampedArray) JSValue() js.Value { return u.TypedArray.Value }
