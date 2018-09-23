// Code generated DO NOT EDIT
// texttrackkind.go
package dom

import "syscall/js"

type TextTrackKind string

const (
	TextTrackKindSubtitles    TextTrackKind = "subtitles"
	TextTrackKindCaptions     TextTrackKind = "captions"
	TextTrackKindDescriptions TextTrackKind = "descriptions"
	TextTrackKindChapters     TextTrackKind = "chapters"
	TextTrackKindMetadata     TextTrackKind = "metadata"
)

func JSValueToTextTrackKind(val js.Value) TextTrackKind {
	return TextTrackKind(val.String())
}
