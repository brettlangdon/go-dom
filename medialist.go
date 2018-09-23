// Code generated DO NOT EDIT
// medialist.go
package dom

import "syscall/js"

type MediaListIFace interface {
	AppendMedium(args ...interface{})
	DeleteMedium(args ...interface{})
	Item(args ...interface{}) string
	GetLength() float64
	GetMediaText() string
	SetMediaText(string)
}
type MediaList struct {
	Value
}

func JSValueToMediaList(val js.Value) MediaList { return MediaList{Value: Value{Value: val}} }
func (v Value) AsMediaList() MediaList          { return MediaList{Value: v} }
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
func (m MediaList) GetLength() float64 {
	val := m.Get("length")
	return val.Float()
}
func (m MediaList) GetMediaText() string {
	val := m.Get("mediaText")
	return val.String()
}
func (m MediaList) SetMediaText(val string) {
	m.Set("mediaText", val)
}
