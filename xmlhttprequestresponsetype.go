// Code generated DO NOT EDIT
// xmlhttprequestresponsetype.go
package dom

import "syscall/js"

type XMLHttpRequestResponseType string

const (
	XMLHttpRequestResponseTypeEmpty       XMLHttpRequestResponseType = "Empty"
	XMLHttpRequestResponseTypeArraybuffer XMLHttpRequestResponseType = "arraybuffer"
	XMLHttpRequestResponseTypeBlob        XMLHttpRequestResponseType = "blob"
	XMLHttpRequestResponseTypeDocument    XMLHttpRequestResponseType = "document"
	XMLHttpRequestResponseTypeJson        XMLHttpRequestResponseType = "json"
	XMLHttpRequestResponseTypeText        XMLHttpRequestResponseType = "text"
)

func JSValueToXMLHttpRequestResponseType(val js.Value) XMLHttpRequestResponseType {
	return XMLHttpRequestResponseType(val.String())
}
