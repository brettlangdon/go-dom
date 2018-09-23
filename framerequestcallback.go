// Code generated DO NOT EDIT
// framerequestcallback.go
package dom

import "syscall/js"

type FrameRequestCallbackCallback func(time DOMHighResTimeStamp)
type FrameRequestCallback struct {
	Callback
}

func jsValueToFrameRequestCallback(val js.Value) FrameRequestCallback {
	return FrameRequestCallback{Callback: jsValueToCallback(val)}
}
func NewFrameRequestCallback(c FrameRequestCallbackCallback) FrameRequestCallback {
	callback := js.NewCallback(func(args []js.Value) {
		time := jsValueToDOMHighResTimeStamp(args[0])
		c(time)
	})
	return FrameRequestCallback{Callback: Callback{Callback: callback}}
}
