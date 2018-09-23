// Code generated DO NOT EDIT
// onbeforeunloadeventhandlernonnull.go
package dom

import "syscall/js"

type OnBeforeUnloadEventHandlerNonNullCallback func(event Event) string
type OnBeforeUnloadEventHandlerNonNull struct {
	Callback
}

func jsValueToOnBeforeUnloadEventHandlerNonNull(val js.Value) OnBeforeUnloadEventHandlerNonNull {
	return OnBeforeUnloadEventHandlerNonNull{Callback: jsValueToCallback(val)}
}
func NewOnBeforeUnloadEventHandlerNonNull(c OnBeforeUnloadEventHandlerNonNullCallback) OnBeforeUnloadEventHandlerNonNull {
	callback := js.NewCallback(func(args []js.Value) {
		event := jsValueToEvent(args[0])
		c(event)
	})
	return OnBeforeUnloadEventHandlerNonNull{Callback: Callback{Callback: callback}}
}
