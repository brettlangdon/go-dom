// Code generated DO NOT EDIT
// resizequality.go
package dom

import "syscall/js"

type ResizeQuality string

const (
	ResizeQualityPixelated ResizeQuality = "pixelated"
	ResizeQualityLow       ResizeQuality = "low"
	ResizeQualityMedium    ResizeQuality = "medium"
	ResizeQualityHigh      ResizeQuality = "high"
)

func jsValueToResizeQuality(val js.Value) ResizeQuality {
	return ResizeQuality(val.String())
}
