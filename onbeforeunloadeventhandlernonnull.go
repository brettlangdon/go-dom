// Code generated DO NOT EDIT
// onbeforeunloadeventhandlernonnull.go
package dom

import "syscall/js"

type OnBeforeUnloadEventHandlerNonNullCallback func(event Event)
type OnBeforeUnloadEventHandlerNonNull struct {
	Callback
}

func JSValueToOnBeforeUnloadEventHandlerNonNull(val js.Value) OnBeforeUnloadEventHandlerNonNull {
	return OnBeforeUnloadEventHandlerNonNull{Callback: JSValueToCallback(val)}
}
func NewOnBeforeUnloadEventHandlerNonNull(c OnBeforeUnloadEventHandlerNonNullCallback) OnBeforeUnloadEventHandlerNonNull {
	callback := js.NewCallback(func(args []js.Value) {
		event := JSValueToEvent(args[0])
		c(event)
	})
	return OnBeforeUnloadEventHandlerNonNull{Callback: Callback{Callback: callback}}
}
