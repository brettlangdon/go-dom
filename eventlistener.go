// Code generated DO NOT EDIT
// eventlistener.go
package dom

import "syscall/js"

type EventListenerHandleEventCallback func(event Event)
type EventListenerHandleEvent struct {
	Callback
}

func JSValueToEventListenerHandleEvent(val js.Value) EventListenerHandleEvent {
	return EventListenerHandleEvent{Callback: JSValueToCallback(val)}
}
func NewEventListenerHandleEvent(c EventListenerHandleEventCallback) EventListenerHandleEvent {
	callback := js.NewCallback(func(args []js.Value) {
		event := JSValueToEvent(args[0])
		c(event)
	})
	return EventListenerHandleEvent{Callback: Callback{Callback: callback}}
}

type EventListener struct {
	Value
	HandleEvent EventListenerHandleEventCallback
}

func JSValueToEventListener(val js.Value) EventListener {
	return EventListener{Value: Value{Value: val}}
}
func (v Value) AsEventListener() EventListener { return EventListener{Value: v} }
