// Code generated DO NOT EDIT
// videotrack.go
package dom

import "syscall/js"

type VideoTrackIFace interface {
	GetId() string
	GetKind() string
	GetLabel() string
	GetLanguage() string
	GetSelected() bool
	SetSelected(bool)
}
type VideoTrack struct {
	Value
}

func jsValueToVideoTrack(val js.Value) VideoTrack { return VideoTrack{Value: Value{Value: val}} }
func (v Value) AsVideoTrack() VideoTrack          { return VideoTrack{Value: v} }
func (v VideoTrack) GetId() string {
	val := v.Get("id")
	return val.String()
}
func (v VideoTrack) GetKind() string {
	val := v.Get("kind")
	return val.String()
}
func (v VideoTrack) GetLabel() string {
	val := v.Get("label")
	return val.String()
}
func (v VideoTrack) GetLanguage() string {
	val := v.Get("language")
	return val.String()
}
func (v VideoTrack) GetSelected() bool {
	val := v.Get("selected")
	return val.Bool()
}
func (v VideoTrack) SetSelected(val bool) {
	v.Set("selected", val)
}
