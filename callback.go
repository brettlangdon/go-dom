package dom

import "syscall/js"

type Callback struct {
	js.Callback
}

func (c *Callback) JSValue() js.Value { return c.Callback.Value }

type EventCallbackFunction func(*Event)

func NewEventCallback(c EventCallbackFunction) *Callback {
	return &Callback{
		Callback: js.NewEventCallback(0, func(evt js.Value) {
			c(NewEvent(evt))
		}),
	}
}
