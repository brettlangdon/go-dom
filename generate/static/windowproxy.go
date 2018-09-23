package dom

import "syscall/js"

type WindowProxy Window

func JSValueToWindowProxy(val js.Value) WindowProxy { return WindowProxy{Value: Value{Value: val}} }
func (v Value) AsWindowProxy() WindowProxy          { return WindowProxy{Value: v} }
