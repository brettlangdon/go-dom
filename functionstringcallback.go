// Code generated DO NOT EDIT
// functionstringcallback.go
package dom

import "syscall/js"

type FunctionStringCallbackCallback func(data string)
type FunctionStringCallback struct {
	Callback
}

func JSValueToFunctionStringCallback(val js.Value) FunctionStringCallback {
	return FunctionStringCallback{Callback: JSValueToCallback(val)}
}
func NewFunctionStringCallback(c FunctionStringCallbackCallback) FunctionStringCallback {
	callback := js.NewCallback(func(args []js.Value) {
		data := args[0].String()
		c(data)
	})
	return FunctionStringCallback{Callback: Callback{Callback: callback}}
}
