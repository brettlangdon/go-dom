// Code generated DO NOT EDIT
// mutationcallback.go
package dom

import "syscall/js"

type MutationCallbackCallback func(mutations Value, observer MutationObserver)
type MutationCallback struct {
	Callback
}

func JSValueToMutationCallback(val js.Value) MutationCallback {
	return MutationCallback{Callback: JSValueToCallback(val)}
}
func NewMutationCallback(c MutationCallbackCallback) MutationCallback {
	callback := js.NewCallback(func(args []js.Value) {
		mutations := JSValueToValue(args[0])
		observer := JSValueToMutationObserver(args[1])
		c(mutations, observer)
	})
	return MutationCallback{Callback: Callback{Callback: callback}}
}
