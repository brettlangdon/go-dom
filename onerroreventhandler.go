// Code generated DO NOT EDIT
// onerroreventhandler.go
package dom

import "syscall/js"

type OnErrorEventHandler OnErrorEventHandlerNonNull

func JSValueToOnErrorEventHandler(val js.Value) OnErrorEventHandler {
	return OnErrorEventHandler(JSValueToOnErrorEventHandlerNonNull(val))
}
func NewOnErrorEventHandler(c OnErrorEventHandlerNonNullCallback) OnErrorEventHandler {
	return OnErrorEventHandler(NewOnErrorEventHandlerNonNull(c))
}
