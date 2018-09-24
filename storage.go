// Code generated DO NOT EDIT
// storage.go
package dom

import "syscall/js"

type StorageIFace interface {
	Clear(args ...interface{})
	GetItem(args ...interface{}) string
	Key(args ...interface{}) string
	GetLength() int
	RemoveItem(args ...interface{})
	SetItem(args ...interface{})
}
type Storage struct {
	Value
}

func JSValueToStorage(val js.Value) Storage { return Storage{Value: JSValueToValue(val)} }
func (v Value) AsStorage() Storage          { return Storage{Value: v} }
func NewStorage(args ...interface{}) Storage {
	return Storage{Value: JSValueToValue(js.Global().Get("Storage").New(args...))}
}
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
func (s Storage) GetLength() int {
	val := s.Get("length")
	return val.Int()
}
func (s Storage) RemoveItem(args ...interface{}) {
	s.Call("removeItem", args...)
}
func (s Storage) SetItem(args ...interface{}) {
	s.Call("setItem", args...)
}
