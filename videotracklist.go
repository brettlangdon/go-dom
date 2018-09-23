// Code generated DO NOT EDIT
// videotracklist.go
package dom

import "syscall/js"

type VideoTrackListIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetTrackById(args ...interface{}) VideoTrack
	GetLength() float64
	GetOnaddtrack() EventHandler
	SetOnaddtrack(EventHandler)
	GetOnchange() EventHandler
	SetOnchange(EventHandler)
	GetOnremovetrack() EventHandler
	SetOnremovetrack(EventHandler)
	RemoveEventListener(args ...interface{})
	GetSelectedIndex() float64
}
type VideoTrackList struct {
	Value
	EventTarget
}

func jsValueToVideoTrackList(val js.Value) VideoTrackList {
	return VideoTrackList{Value: Value{Value: val}}
}
func (v Value) AsVideoTrackList() VideoTrackList { return VideoTrackList{Value: v} }
func (v VideoTrackList) GetTrackById(args ...interface{}) VideoTrack {
	val := v.Call("getTrackById", args...)
	return jsValueToVideoTrack(val.JSValue())
}
func (v VideoTrackList) GetLength() float64 {
	val := v.Get("length")
	return val.Float()
}
func (v VideoTrackList) GetOnaddtrack() EventHandler {
	val := v.Get("onaddtrack")
	return jsValueToEventHandler(val.JSValue())
}
func (v VideoTrackList) SetOnaddtrack(val EventHandler) {
	v.Set("onaddtrack", val)
}
func (v VideoTrackList) GetOnchange() EventHandler {
	val := v.Get("onchange")
	return jsValueToEventHandler(val.JSValue())
}
func (v VideoTrackList) SetOnchange(val EventHandler) {
	v.Set("onchange", val)
}
func (v VideoTrackList) GetOnremovetrack() EventHandler {
	val := v.Get("onremovetrack")
	return jsValueToEventHandler(val.JSValue())
}
func (v VideoTrackList) SetOnremovetrack(val EventHandler) {
	v.Set("onremovetrack", val)
}
func (v VideoTrackList) GetSelectedIndex() float64 {
	val := v.Get("selectedIndex")
	return val.Float()
}
