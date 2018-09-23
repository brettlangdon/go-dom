// Code generated DO NOT EDIT
// framerequestcallback.go
package dom

import "syscall/js"

type FrameRequestCallbackCallback func(time DOMHighResTimeStamp)
type FrameRequestCallback struct {
	Callback
}

func JSValueToFrameRequestCallback(val js.Value) FrameRequestCallback {
	return FrameRequestCallback{Callback: JSValueToCallback(val)}
}
func NewFrameRequestCallback(c FrameRequestCallbackCallback) FrameRequestCallback {
	callback := js.NewCallback(func(args []js.Value) {
		time := JSValueToDOMHighResTimeStamp(args[0])
		c(time)
	})
	return FrameRequestCallback{Callback: Callback{Callback: callback}}
}
