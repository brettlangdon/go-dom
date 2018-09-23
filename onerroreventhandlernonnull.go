// Code generated DO NOT EDIT
// onerroreventhandlernonnull.go
package dom

import "syscall/js"

type OnErrorEventHandlerNonNullCallback func(event Value, source string, lineno float64, colno float64, error Value) Value
type OnErrorEventHandlerNonNull struct {
	Callback
}

func jsValueToOnErrorEventHandlerNonNull(val js.Value) OnErrorEventHandlerNonNull {
	return OnErrorEventHandlerNonNull{Callback: jsValueToCallback(val)}
}
func NewOnErrorEventHandlerNonNull(c OnErrorEventHandlerNonNullCallback) OnErrorEventHandlerNonNull {
	callback := js.NewCallback(func(args []js.Value) {
		event := jsValueToValue(args[0])
		source := args[1].String()
		lineno := args[2].Float()
		colno := args[3].Float()
		error := jsValueToValue(args[4])
		c(event, source, lineno, colno, error)
	})
	return OnErrorEventHandlerNonNull{Callback: Callback{Callback: callback}}
}
