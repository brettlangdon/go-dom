// Code generated DO NOT EDIT
// mutationobserver.go
package dom

import "syscall/js"

type MutationObserverIFace interface {
	Disconnect(args ...interface{})
	Observe(args ...interface{})
	TakeRecords(args ...interface{})
}
type MutationObserver struct {
	Value
}

func jsValueToMutationObserver(val js.Value) MutationObserver {
	return MutationObserver{Value: Value{Value: val}}
}
func (v Value) AsMutationObserver() MutationObserver { return MutationObserver{Value: v} }
func (m MutationObserver) Disconnect(args ...interface{}) {
	m.Call("disconnect", args...)
}
func (m MutationObserver) Observe(args ...interface{}) {
	m.Call("observe", args...)
}
func (m MutationObserver) TakeRecords(args ...interface{}) {
	m.Call("takeRecords", args...)
}
