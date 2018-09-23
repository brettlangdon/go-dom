// Code generated DO NOT EDIT
// history.go
package dom

import "syscall/js"

type HistoryIFace interface {
	Back(args ...interface{})
	Forward(args ...interface{})
	Go(args ...interface{})
	GetLength() float64
	PushState(args ...interface{})
	ReplaceState(args ...interface{})
	GetScrollRestoration() ScrollRestoration
	SetScrollRestoration(ScrollRestoration)
	GetState() Value
}
type History struct {
	Value
}

func JSValueToHistory(val js.Value) History { return History{Value: Value{Value: val}} }
func (v Value) AsHistory() History          { return History{Value: v} }
func (h History) Back(args ...interface{}) {
	h.Call("back", args...)
}
func (h History) Forward(args ...interface{}) {
	h.Call("forward", args...)
}
func (h History) Go(args ...interface{}) {
	h.Call("go", args...)
}
func (h History) GetLength() float64 {
	val := h.Get("length")
	return val.Float()
}
func (h History) PushState(args ...interface{}) {
	h.Call("pushState", args...)
}
func (h History) ReplaceState(args ...interface{}) {
	h.Call("replaceState", args...)
}
func (h History) GetScrollRestoration() ScrollRestoration {
	val := h.Get("scrollRestoration")
	return JSValueToScrollRestoration(val.JSValue())
}
func (h History) SetScrollRestoration(val ScrollRestoration) {
	h.Set("scrollRestoration", val)
}
func (h History) GetState() Value {
	val := h.Get("state")
	return val
}
