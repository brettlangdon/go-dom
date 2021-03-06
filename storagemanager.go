// Code generated DO NOT EDIT
// storagemanager.go
package dom

import "syscall/js"

type StorageManagerIFace interface {
	Estimate(args ...interface{})
	Persist(args ...interface{})
	Persisted(args ...interface{})
}
type StorageManager struct {
	Value
}

func JSValueToStorageManager(val js.Value) StorageManager {
	return StorageManager{Value: JSValueToValue(val)}
}
func (v Value) AsStorageManager() StorageManager { return StorageManager{Value: v} }
func NewStorageManager(args ...interface{}) StorageManager {
	return StorageManager{Value: JSValueToValue(js.Global().Get("StorageManager").New(args...))}
}
func (s StorageManager) Estimate(args ...interface{}) {
	s.Call("estimate", args...)
}
func (s StorageManager) Persist(args ...interface{}) {
	s.Call("persist", args...)
}
func (s StorageManager) Persisted(args ...interface{}) {
	s.Call("persisted", args...)
}
