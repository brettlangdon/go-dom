// Code generated DO NOT EDIT
// customelementconstructor.go
package dom

import "syscall/js"

type CustomElementConstructorCallback func() Value
type CustomElementConstructor struct {
	Callback
}

func jsValueToCustomElementConstructor(val js.Value) CustomElementConstructor {
	return CustomElementConstructor{Callback: jsValueToCallback(val)}
}
func NewCustomElementConstructor(c CustomElementConstructorCallback) CustomElementConstructor {
	callback := js.NewCallback(func(args []js.Value) {
		c()
	})
	return CustomElementConstructor{Callback: Callback{Callback: callback}}
}