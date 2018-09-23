// Code generated DO NOT EDIT
// eventhandlernonnull.go
package dom

import "syscall/js"

type EventHandlerNonNullCallback func(event Event) Value
type EventHandlerNonNull struct {
	Callback
}

func JSValueToEventHandlerNonNull(val js.Value) EventHandlerNonNull {
	return EventHandlerNonNull{Callback: JSValueToCallback(val)}
}
func NewEventHandlerNonNull(c EventHandlerNonNullCallback) EventHandlerNonNull {
	callback := js.NewCallback(func(args []js.Value) {
		event := JSValueToEvent(args[0])
		c(event)
	})
	return EventHandlerNonNull{Callback: Callback{Callback: callback}}
}
