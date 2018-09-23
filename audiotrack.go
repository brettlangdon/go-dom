// Code generated DO NOT EDIT
// audiotrack.go
package dom

import "syscall/js"

type AudioTrackIFace interface {
	GetEnabled() bool
	SetEnabled(bool)
	GetId() string
	GetKind() string
	GetLabel() string
	GetLanguage() string
}
type AudioTrack struct {
	Value
}

func JSValueToAudioTrack(val js.Value) AudioTrack { return AudioTrack{Value: Value{Value: val}} }
func (v Value) AsAudioTrack() AudioTrack          { return AudioTrack{Value: v} }
func (a AudioTrack) GetEnabled() bool {
	val := a.Get("enabled")
	return val.Bool()
}
func (a AudioTrack) SetEnabled(val bool) {
	a.Set("enabled", val)
}
func (a AudioTrack) GetId() string {
	val := a.Get("id")
	return val.String()
}
func (a AudioTrack) GetKind() string {
	val := a.Get("kind")
	return val.String()
}
func (a AudioTrack) GetLabel() string {
	val := a.Get("label")
	return val.String()
}
func (a AudioTrack) GetLanguage() string {
	val := a.Get("language")
	return val.String()
}
