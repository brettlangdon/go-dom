package dom

import "syscall/js"

type ShadowRootInit struct {
	Value
	Mode string
}

func (s ShadowRootInit) JSValue() js.Value {
	v := js.Global().Get("Object").New()
	v.Set("mode", s.Mode)
	return v
}
