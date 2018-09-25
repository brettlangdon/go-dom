// Code generated DO NOT EDIT
// texttracklist.go
package dom

import "syscall/js"

type TextTrackListIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetTrackById(args ...interface{}) TextTrack
	GetLength() int
	GetOnaddtrack() EventHandler
	SetOnaddtrack(EventHandler)
	GetOnchange() EventHandler
	SetOnchange(EventHandler)
	GetOnremovetrack() EventHandler
	SetOnremovetrack(EventHandler)
	RemoveEventListener(args ...interface{})
}
type TextTrackList struct {
	Value
}

func JSValueToTextTrackList(val js.Value) TextTrackList {
	return TextTrackList{Value: JSValueToValue(val)}
}
func (v Value) AsTextTrackList() TextTrackList { return TextTrackList{Value: v} }
func NewTextTrackList(args ...interface{}) TextTrackList {
	return TextTrackList{Value: JSValueToValue(js.Global().Get("TextTrackList").New(args...))}
}
func (t TextTrackList) AddEventListener(args ...interface{}) {
	t.Call("addEventListener", args...)
}
func (t TextTrackList) DispatchEvent(args ...interface{}) bool {
	val := t.Call("dispatchEvent", args...)
	return val.Bool()
}
func (t TextTrackList) GetTrackById(args ...interface{}) TextTrack {
	val := t.Call("getTrackById", args...)
	return JSValueToTextTrack(val.JSValue())
}
func (t TextTrackList) GetLength() int {
	val := t.Get("length")
	return val.Int()
}
func (t TextTrackList) GetOnaddtrack() EventHandler {
	val := t.Get("onaddtrack")
	return JSValueToEventHandler(val.JSValue())
}
func (t TextTrackList) SetOnaddtrack(val EventHandler) {
	t.Set("onaddtrack", val)
}
func (t TextTrackList) GetOnchange() EventHandler {
	val := t.Get("onchange")
	return JSValueToEventHandler(val.JSValue())
}
func (t TextTrackList) SetOnchange(val EventHandler) {
	t.Set("onchange", val)
}
func (t TextTrackList) GetOnremovetrack() EventHandler {
	val := t.Get("onremovetrack")
	return JSValueToEventHandler(val.JSValue())
}
func (t TextTrackList) SetOnremovetrack(val EventHandler) {
	t.Set("onremovetrack", val)
}
func (t TextTrackList) RemoveEventListener(args ...interface{}) {
	t.Call("removeEventListener", args...)
}
