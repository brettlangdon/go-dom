// Code generated DO NOT EDIT
// eventhandler.go
package dom

import "syscall/js"

type EventHandler EventHandlerNonNull

func JSValueToEventHandler(val js.Value) EventHandler {
	return EventHandler(JSValueToEventHandlerNonNull(val))
}
func NewEventHandler(c EventHandlerNonNullCallback) EventHandler {
	return EventHandler(NewEventHandlerNonNull(c))
}
