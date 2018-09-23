// Code generated DO NOT EDIT
// premultiplyalpha.go
package dom

import "syscall/js"

type PremultiplyAlpha string

const (
	PremultiplyAlphaNone        PremultiplyAlpha = "none"
	PremultiplyAlphaPremultiply PremultiplyAlpha = "premultiply"
	PremultiplyAlphaDefault     PremultiplyAlpha = "default"
)

func jsValueToPremultiplyAlpha(val js.Value) PremultiplyAlpha {
	return PremultiplyAlpha(val.String())
}
