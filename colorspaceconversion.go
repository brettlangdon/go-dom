// Code generated DO NOT EDIT
// colorspaceconversion.go
package dom

import "syscall/js"

type ColorSpaceConversion string

const (
	ColorSpaceConversionNone    ColorSpaceConversion = "none"
	ColorSpaceConversionDefault ColorSpaceConversion = "default"
)

func JSValueToColorSpaceConversion(val js.Value) ColorSpaceConversion {
	return ColorSpaceConversion(val.String())
}
