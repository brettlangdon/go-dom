// Code generated DO NOT EDIT
// imageorientation.go
package dom

import "syscall/js"

type ImageOrientation string

const (
	ImageOrientationNone  ImageOrientation = "none"
	ImageOrientationFlipY ImageOrientation = "flipY"
)

func jsValueToImageOrientation(val js.Value) ImageOrientation {
	return ImageOrientation(val.String())
}
