// Code generated DO NOT EDIT
// audiotracklist.go
package dom

import "syscall/js"

type AudioTrackListIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetTrackById(args ...interface{}) AudioTrack
	GetLength() int
	GetOnaddtrack() EventHandler
	SetOnaddtrack(EventHandler)
	GetOnchange() EventHandler
	SetOnchange(EventHandler)
	GetOnremovetrack() EventHandler
	SetOnremovetrack(EventHandler)
	RemoveEventListener(args ...interface{})
}
type AudioTrackList struct {
	Value
}

func JSValueToAudioTrackList(val js.Value) AudioTrackList {
	return AudioTrackList{Value: JSValueToValue(val)}
}
func (v Value) AsAudioTrackList() AudioTrackList { return AudioTrackList{Value: v} }
func NewAudioTrackList(args ...interface{}) AudioTrackList {
	return AudioTrackList{Value: JSValueToValue(js.Global().Get("AudioTrackList").New(args...))}
}
func (a AudioTrackList) AddEventListener(args ...interface{}) {
	a.Call("addEventListener", args...)
}
func (a AudioTrackList) DispatchEvent(args ...interface{}) bool {
	val := a.Call("dispatchEvent", args...)
	return val.Bool()
}
func (a AudioTrackList) GetTrackById(args ...interface{}) AudioTrack {
	val := a.Call("getTrackById", args...)
	return JSValueToAudioTrack(val.JSValue())
}
func (a AudioTrackList) GetLength() int {
	val := a.Get("length")
	return val.Int()
}
func (a AudioTrackList) GetOnaddtrack() EventHandler {
	val := a.Get("onaddtrack")
	return JSValueToEventHandler(val.JSValue())
}
func (a AudioTrackList) SetOnaddtrack(val EventHandler) {
	a.Set("onaddtrack", val)
}
func (a AudioTrackList) GetOnchange() EventHandler {
	val := a.Get("onchange")
	return JSValueToEventHandler(val.JSValue())
}
func (a AudioTrackList) SetOnchange(val EventHandler) {
	a.Set("onchange", val)
}
func (a AudioTrackList) GetOnremovetrack() EventHandler {
	val := a.Get("onremovetrack")
	return JSValueToEventHandler(val.JSValue())
}
func (a AudioTrackList) SetOnremovetrack(val EventHandler) {
	a.Set("onremovetrack", val)
}
func (a AudioTrackList) RemoveEventListener(args ...interface{}) {
	a.Call("removeEventListener", args...)
}
