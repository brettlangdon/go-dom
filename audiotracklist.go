// Code generated DO NOT EDIT
// audiotracklist.go
package dom

import "syscall/js"

type AudioTrackListIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetTrackById(args ...interface{}) AudioTrack
	GetLength() float64
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
	EventTarget
}

func jsValueToAudioTrackList(val js.Value) AudioTrackList {
	return AudioTrackList{Value: Value{Value: val}}
}
func (v Value) AsAudioTrackList() AudioTrackList { return AudioTrackList{Value: v} }
func (a AudioTrackList) GetTrackById(args ...interface{}) AudioTrack {
	val := a.Call("getTrackById", args...)
	return jsValueToAudioTrack(val.JSValue())
}
func (a AudioTrackList) GetLength() float64 {
	val := a.Get("length")
	return val.Float()
}
func (a AudioTrackList) GetOnaddtrack() EventHandler {
	val := a.Get("onaddtrack")
	return jsValueToEventHandler(val.JSValue())
}
func (a AudioTrackList) SetOnaddtrack(val EventHandler) {
	a.Set("onaddtrack", val)
}
func (a AudioTrackList) GetOnchange() EventHandler {
	val := a.Get("onchange")
	return jsValueToEventHandler(val.JSValue())
}
func (a AudioTrackList) SetOnchange(val EventHandler) {
	a.Set("onchange", val)
}
func (a AudioTrackList) GetOnremovetrack() EventHandler {
	val := a.Get("onremovetrack")
	return jsValueToEventHandler(val.JSValue())
}
func (a AudioTrackList) SetOnremovetrack(val EventHandler) {
	a.Set("onremovetrack", val)
}