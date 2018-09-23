// Code generated DO NOT EDIT
// endingtype.go
package dom

import "syscall/js"

type EndingType string

const (
	EndingTypeTransparent EndingType = "transparent"
	EndingTypeNative      EndingType = "native"
)

func JSValueToEndingType(val js.Value) EndingType {
	return EndingType(val.String())
}
