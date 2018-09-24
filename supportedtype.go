// Code generated DO NOT EDIT
// supportedtype.go
package dom

import "syscall/js"

type SupportedType string

const (
	SupportedTypeTextHtml            SupportedType = "text/html"
	SupportedTypeTextXml             SupportedType = "text/xml"
	SupportedTypeApplicationXml      SupportedType = "application/xml"
	SupportedTypeApplicationXhtmlXml SupportedType = "application/xhtml+xml"
	SupportedTypeImageSvgXml         SupportedType = "image/svg+xml"
)

func JSValueToSupportedType(val js.Value) SupportedType {
	return SupportedType(val.String())
}
