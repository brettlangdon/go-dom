// Code generated DO NOT EDIT
// texttrack.go
package dom

import "syscall/js"

type TextTrackIFace interface {
	GetActiveCues() TextTrackCueList
	AddCue(args ...interface{})
	AddEventListener(args ...interface{})
	GetCues() TextTrackCueList
	DispatchEvent(args ...interface{}) bool
	GetId() string
	GetInBandMetadataTrackDispatchType() string
	GetKind() TextTrackKind
	GetLabel() string
	GetLanguage() string
	GetMode() TextTrackMode
	SetMode(TextTrackMode)
	GetOncuechange() EventHandler
	SetOncuechange(EventHandler)
	RemoveCue(args ...interface{})
	RemoveEventListener(args ...interface{})
}
type TextTrack struct {
	Value
}

func JSValueToTextTrack(val js.Value) TextTrack { return TextTrack{Value: JSValueToValue(val)} }
func (v Value) AsTextTrack() TextTrack          { return TextTrack{Value: v} }
func NewTextTrack(args ...interface{}) TextTrack {
	return TextTrack{Value: JSValueToValue(js.Global().Get("TextTrack").New(args...))}
}
func (t TextTrack) GetActiveCues() TextTrackCueList {
	val := t.Get("activeCues")
	return JSValueToTextTrackCueList(val.JSValue())
}
func (t TextTrack) AddCue(args ...interface{}) {
	t.Call("addCue", args...)
}
func (t TextTrack) AddEventListener(args ...interface{}) {
	t.Call("addEventListener", args...)
}
func (t TextTrack) GetCues() TextTrackCueList {
	val := t.Get("cues")
	return JSValueToTextTrackCueList(val.JSValue())
}
func (t TextTrack) DispatchEvent(args ...interface{}) bool {
	val := t.Call("dispatchEvent", args...)
	return val.Bool()
}
func (t TextTrack) GetId() string {
	val := t.Get("id")
	return val.String()
}
func (t TextTrack) GetInBandMetadataTrackDispatchType() string {
	val := t.Get("inBandMetadataTrackDispatchType")
	return val.String()
}
func (t TextTrack) GetKind() TextTrackKind {
	val := t.Get("kind")
	return JSValueToTextTrackKind(val.JSValue())
}
func (t TextTrack) GetLabel() string {
	val := t.Get("label")
	return val.String()
}
func (t TextTrack) GetLanguage() string {
	val := t.Get("language")
	return val.String()
}
func (t TextTrack) GetMode() TextTrackMode {
	val := t.Get("mode")
	return JSValueToTextTrackMode(val.JSValue())
}
func (t TextTrack) SetMode(val TextTrackMode) {
	t.Set("mode", val)
}
func (t TextTrack) GetOncuechange() EventHandler {
	val := t.Get("oncuechange")
	return JSValueToEventHandler(val.JSValue())
}
func (t TextTrack) SetOncuechange(val EventHandler) {
	t.Set("oncuechange", val)
}
func (t TextTrack) RemoveCue(args ...interface{}) {
	t.Call("removeCue", args...)
}
func (t TextTrack) RemoveEventListener(args ...interface{}) {
	t.Call("removeEventListener", args...)
}
