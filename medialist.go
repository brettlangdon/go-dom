// Code generated DO NOT EDIT
// medialist.go
package dom

import "syscall/js"

type MediaListIFace interface {
	AppendMedium(args ...interface{})
	DeleteMedium(args ...interface{})
	Item(args ...interface{}) string
	GetLength() int
	GetMediaText() string
	SetMediaText(string)
}
type MediaList struct {
	Value
}

func JSValueToMediaList(val js.Value) MediaList { return MediaList{Value: JSValueToValue(val)} }
func (v Value) AsMediaList() MediaList          { return MediaList{Value: v} }
func NewMediaList(args ...interface{}) MediaList {
	return MediaList{Value: JSValueToValue(js.Global().Get("MediaList").New(args...))}
}
func (m MediaList) AppendMedium(args ...interface{}) {
	m.Call("appendMedium", args...)
}
func (m MediaList) DeleteMedium(args ...interface{}) {
	m.Call("deleteMedium", args...)
}
func (m MediaList) Item(args ...interface{}) string {
	val := m.Call("item", args...)
	return val.String()
}
func (m MediaList) GetLength() int {
	val := m.Get("length")
	return val.Int()
}
func (m MediaList) GetMediaText() string {
	val := m.Get("mediaText")
	return val.String()
}
func (m MediaList) SetMediaText(val string) {
	m.Set("mediaText", val)
}
