// Code generated DO NOT EDIT
// videotracklist.go
package dom

import "syscall/js"

type VideoTrackListIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetTrackById(args ...interface{}) VideoTrack
	GetLength() int
	GetOnaddtrack() EventHandler
	SetOnaddtrack(EventHandler)
	GetOnchange() EventHandler
	SetOnchange(EventHandler)
	GetOnremovetrack() EventHandler
	SetOnremovetrack(EventHandler)
	RemoveEventListener(args ...interface{})
	GetSelectedIndex() int
}
type VideoTrackList struct {
	Value
	EventTarget
}

func JSValueToVideoTrackList(val js.Value) VideoTrackList {
	return VideoTrackList{Value: JSValueToValue(val)}
}
func (v Value) AsVideoTrackList() VideoTrackList { return VideoTrackList{Value: v} }
func NewVideoTrackList(args ...interface{}) VideoTrackList {
	return VideoTrackList{Value: JSValueToValue(js.Global().Get("VideoTrackList").New(args...))}
}
func (v VideoTrackList) GetTrackById(args ...interface{}) VideoTrack {
	val := v.Call("getTrackById", args...)
	return JSValueToVideoTrack(val.JSValue())
}
func (v VideoTrackList) GetLength() int {
	val := v.Get("length")
	return val.Int()
}
func (v VideoTrackList) GetOnaddtrack() EventHandler {
	val := v.Get("onaddtrack")
	return JSValueToEventHandler(val.JSValue())
}
func (v VideoTrackList) SetOnaddtrack(val EventHandler) {
	v.Set("onaddtrack", val)
}
func (v VideoTrackList) GetOnchange() EventHandler {
	val := v.Get("onchange")
	return JSValueToEventHandler(val.JSValue())
}
func (v VideoTrackList) SetOnchange(val EventHandler) {
	v.Set("onchange", val)
}
func (v VideoTrackList) GetOnremovetrack() EventHandler {
	val := v.Get("onremovetrack")
	return JSValueToEventHandler(val.JSValue())
}
func (v VideoTrackList) SetOnremovetrack(val EventHandler) {
	v.Set("onremovetrack", val)
}
func (v VideoTrackList) GetSelectedIndex() int {
	val := v.Get("selectedIndex")
	return val.Int()
}
