// Code generated DO NOT EDIT
// storage.go
package dom

import "syscall/js"

type StorageIFace interface {
	Clear(args ...interface{})
	GetItem(args ...interface{}) string
	Key(args ...interface{}) string
	GetLength() float64
	RemoveItem(args ...interface{})
	SetItem(args ...interface{})
}
type Storage struct {
	Value
}

func jsValueToStorage(val js.Value) Storage { return Storage{Value: Value{Value: val}} }
func (v Value) AsStorage() Storage          { return Storage{Value: v} }
func (s Storage) Clear(args ...interface{}) {
	s.Call("clear", args...)
}
func (s Storage) GetItem(args ...interface{}) string {
	val := s.Call("getItem", args...)
	return val.String()
}
func (s Storage) Key(args ...interface{}) string {
	val := s.Call("key", args...)
	return val.String()
}
func (s Storage) GetLength() float64 {
	val := s.Get("length")
	return val.Float()
}
func (s Storage) RemoveItem(args ...interface{}) {
	s.Call("removeItem", args...)
}
func (s Storage) SetItem(args ...interface{}) {
	s.Call("setItem", args...)
}
