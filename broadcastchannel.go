// Code generated DO NOT EDIT
// broadcastchannel.go
package dom

import "syscall/js"

type BroadcastChannelIFace interface {
	AddEventListener(args ...interface{})
	Close(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetName() string
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnmessageerror() EventHandler
	SetOnmessageerror(EventHandler)
	PostMessage(args ...interface{})
	RemoveEventListener(args ...interface{})
}
type BroadcastChannel struct {
	Value
	EventTarget
}

func jsValueToBroadcastChannel(val js.Value) BroadcastChannel {
	return BroadcastChannel{Value: Value{Value: val}}
}
func (v Value) AsBroadcastChannel() BroadcastChannel { return BroadcastChannel{Value: v} }
func (b BroadcastChannel) Close(args ...interface{}) {
	b.Call("close", args...)
}
func (b BroadcastChannel) GetName() string {
	val := b.Get("name")
	return val.String()
}
func (b BroadcastChannel) GetOnmessage() EventHandler {
	val := b.Get("onmessage")
	return jsValueToEventHandler(val.JSValue())
}
func (b BroadcastChannel) SetOnmessage(val EventHandler) {
	b.Set("onmessage", val)
}
func (b BroadcastChannel) GetOnmessageerror() EventHandler {
	val := b.Get("onmessageerror")
	return jsValueToEventHandler(val.JSValue())
}
func (b BroadcastChannel) SetOnmessageerror(val EventHandler) {
	b.Set("onmessageerror", val)
}
func (b BroadcastChannel) PostMessage(args ...interface{}) {
	b.Call("postMessage", args...)
}
