// Code generated DO NOT EDIT
// eventhandler.go
package dom

import "syscall/js"

type EventHandler EventHandlerNonNull

func jsValueToEventHandler(val js.Value) EventHandler {
	return EventHandler(jsValueToEventHandlerNonNull(val))
}
