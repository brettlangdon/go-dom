package dom

import "syscall/js"

type Callback struct {
	js.Callback
}

func jsValueToCallback(val js.Value) Callback { return Callback{Callback: js.Callback{Value: val}} }
func (c Callback) JSValue() js.Value          { return c.Callback.Value }
