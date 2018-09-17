package dom

import "syscall/js"

type JSValue interface {
	JSValue() js.Value
}

func ToJSValue(x interface{}) js.Value {
	if v, ok := x.(JSValue); ok {
		return v.JSValue()
	}
	return js.ValueOf(x)
}
