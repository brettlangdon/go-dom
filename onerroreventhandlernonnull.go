// Code generated DO NOT EDIT
// onerroreventhandlernonnull.go
package dom

import "syscall/js"

type OnErrorEventHandlerNonNullCallback func(event Value, source string, lineno int, colno int, error Value)
type OnErrorEventHandlerNonNull struct {
	Callback
}

func JSValueToOnErrorEventHandlerNonNull(val js.Value) OnErrorEventHandlerNonNull {
	return OnErrorEventHandlerNonNull{Callback: JSValueToCallback(val)}
}
func NewOnErrorEventHandlerNonNull(c OnErrorEventHandlerNonNullCallback) OnErrorEventHandlerNonNull {
	callback := js.NewCallback(func(args []js.Value) {
		event := JSValueToValue(args[0])
		source := args[1].String()
		lineno := args[2].Int()
		colno := args[3].Int()
		error := JSValueToValue(args[4])
		c(event, source, lineno, colno, error)
	})
	return OnErrorEventHandlerNonNull{Callback: Callback{Callback: callback}}
}
