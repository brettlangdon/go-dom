// Code generated DO NOT EDIT
// texttrackmode.go
package dom

import "syscall/js"

type TextTrackMode string

const (
	TextTrackModeDisabled TextTrackMode = "disabled"
	TextTrackModeHidden   TextTrackMode = "hidden"
	TextTrackModeShowing  TextTrackMode = "showing"
)

func jsValueToTextTrackMode(val js.Value) TextTrackMode {
	return TextTrackMode(val.String())
}
