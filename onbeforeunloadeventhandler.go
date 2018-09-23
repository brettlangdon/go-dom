// Code generated DO NOT EDIT
// onbeforeunloadeventhandler.go
package dom

import "syscall/js"

type OnBeforeUnloadEventHandler OnBeforeUnloadEventHandlerNonNull

func JSValueToOnBeforeUnloadEventHandler(val js.Value) OnBeforeUnloadEventHandler {
	return OnBeforeUnloadEventHandler(JSValueToOnBeforeUnloadEventHandlerNonNull(val))
}
func NewOnBeforeUnloadEventHandler(c OnBeforeUnloadEventHandlerNonNullCallback) OnBeforeUnloadEventHandler {
	return OnBeforeUnloadEventHandler(NewOnBeforeUnloadEventHandlerNonNull(c))
}
