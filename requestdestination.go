// Code generated DO NOT EDIT
// requestdestination.go
package dom

import "syscall/js"

type RequestDestination string

const (
	RequestDestinationEmpty        RequestDestination = "Empty"
	RequestDestinationAudio        RequestDestination = "audio"
	RequestDestinationAudioworklet RequestDestination = "audioworklet"
	RequestDestinationDocument     RequestDestination = "document"
	RequestDestinationEmbed        RequestDestination = "embed"
	RequestDestinationFont         RequestDestination = "font"
	RequestDestinationImage        RequestDestination = "image"
	RequestDestinationManifest     RequestDestination = "manifest"
	RequestDestinationObject       RequestDestination = "object"
	RequestDestinationPaintworklet RequestDestination = "paintworklet"
	RequestDestinationReport       RequestDestination = "report"
	RequestDestinationScript       RequestDestination = "script"
	RequestDestinationSharedworker RequestDestination = "sharedworker"
	RequestDestinationStyle        RequestDestination = "style"
	RequestDestinationTrack        RequestDestination = "track"
	RequestDestinationVideo        RequestDestination = "video"
	RequestDestinationWorker       RequestDestination = "worker"
	RequestDestinationXslt         RequestDestination = "xslt"
)

func JSValueToRequestDestination(val js.Value) RequestDestination {
	return RequestDestination(val.String())
}
