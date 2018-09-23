// Code generated DO NOT EDIT
// texttrackcuelist.go
package dom

import "syscall/js"

type TextTrackCueListIFace interface {
	GetCueById(args ...interface{}) TextTrackCue
	GetLength() float64
}
type TextTrackCueList struct {
	Value
}

func jsValueToTextTrackCueList(val js.Value) TextTrackCueList {
	return TextTrackCueList{Value: Value{Value: val}}
}
func (v Value) AsTextTrackCueList() TextTrackCueList { return TextTrackCueList{Value: v} }
func (t TextTrackCueList) GetCueById(args ...interface{}) TextTrackCue {
	val := t.Call("getCueById", args...)
	return jsValueToTextTrackCue(val.JSValue())
}
func (t TextTrackCueList) GetLength() float64 {
	val := t.Get("length")
	return val.Float()
}
