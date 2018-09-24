// Code generated DO NOT EDIT
// texttrackcue.go
package dom

import "syscall/js"

type TextTrackCueIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetEndTime() float64
	SetEndTime(float64)
	GetId() string
	SetId(string)
	GetOnenter() EventHandler
	SetOnenter(EventHandler)
	GetOnexit() EventHandler
	SetOnexit(EventHandler)
	GetPauseOnExit() bool
	SetPauseOnExit(bool)
	RemoveEventListener(args ...interface{})
	GetStartTime() float64
	SetStartTime(float64)
	GetTrack() TextTrack
}
type TextTrackCue struct {
	Value
	EventTarget
}

func JSValueToTextTrackCue(val js.Value) TextTrackCue { return TextTrackCue{Value: JSValueToValue(val)} }
func (v Value) AsTextTrackCue() TextTrackCue          { return TextTrackCue{Value: v} }
func NewTextTrackCue(args ...interface{}) TextTrackCue {
	return TextTrackCue{Value: JSValueToValue(js.Global().Get("TextTrackCue").New(args...))}
}
func (t TextTrackCue) GetEndTime() float64 {
	val := t.Get("endTime")
	return val.Float()
}
func (t TextTrackCue) SetEndTime(val float64) {
	t.Set("endTime", val)
}
func (t TextTrackCue) GetId() string {
	val := t.Get("id")
	return val.String()
}
func (t TextTrackCue) SetId(val string) {
	t.Set("id", val)
}
func (t TextTrackCue) GetOnenter() EventHandler {
	val := t.Get("onenter")
	return JSValueToEventHandler(val.JSValue())
}
func (t TextTrackCue) SetOnenter(val EventHandler) {
	t.Set("onenter", val)
}
func (t TextTrackCue) GetOnexit() EventHandler {
	val := t.Get("onexit")
	return JSValueToEventHandler(val.JSValue())
}
func (t TextTrackCue) SetOnexit(val EventHandler) {
	t.Set("onexit", val)
}
func (t TextTrackCue) GetPauseOnExit() bool {
	val := t.Get("pauseOnExit")
	return val.Bool()
}
func (t TextTrackCue) SetPauseOnExit(val bool) {
	t.Set("pauseOnExit", val)
}
func (t TextTrackCue) GetStartTime() float64 {
	val := t.Get("startTime")
	return val.Float()
}
func (t TextTrackCue) SetStartTime(val float64) {
	t.Set("startTime", val)
}
func (t TextTrackCue) GetTrack() TextTrack {
	val := t.Get("track")
	return JSValueToTextTrack(val.JSValue())
}
