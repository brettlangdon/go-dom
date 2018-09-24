// Code generated DO NOT EDIT
// texttrackcuelist.go
package dom

import "syscall/js"

type TextTrackCueListIFace interface {
	GetCueById(args ...interface{}) TextTrackCue
	GetLength() int
}
type TextTrackCueList struct {
	Value
}

func JSValueToTextTrackCueList(val js.Value) TextTrackCueList {
	return TextTrackCueList{Value: JSValueToValue(val)}
}
func (v Value) AsTextTrackCueList() TextTrackCueList { return TextTrackCueList{Value: v} }
func NewTextTrackCueList(args ...interface{}) TextTrackCueList {
	return TextTrackCueList{Value: JSValueToValue(js.Global().Get("TextTrackCueList").New(args...))}
}
func (t TextTrackCueList) GetCueById(args ...interface{}) TextTrackCue {
	val := t.Call("getCueById", args...)
	return JSValueToTextTrackCue(val.JSValue())
}
func (t TextTrackCueList) GetLength() int {
	val := t.Get("length")
	return val.Int()
}
