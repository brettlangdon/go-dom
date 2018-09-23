// Code generated DO NOT EDIT
// imagesmoothingquality.go
package dom

import "syscall/js"

type ImageSmoothingQuality string

const (
	ImageSmoothingQualityLow    ImageSmoothingQuality = "low"
	ImageSmoothingQualityMedium ImageSmoothingQuality = "medium"
	ImageSmoothingQualityHigh   ImageSmoothingQuality = "high"
)

func jsValueToImageSmoothingQuality(val js.Value) ImageSmoothingQuality {
	return ImageSmoothingQuality(val.String())
}
