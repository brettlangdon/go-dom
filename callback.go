package dom

import "syscall/js"

type Callback struct {
	js.Callback
}

func JSValueToCallback(val js.Value) Callback { return Callback{Callback: js.Callback{Value: val}} }
func (c Callback) JSValue() js.Value          { return c.Callback.Value }
