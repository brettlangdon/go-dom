// Code generated DO NOT EDIT
// requestmode.go
package dom

import "syscall/js"

type RequestMode string

const (
	RequestModeNavigate   RequestMode = "navigate"
	RequestModeSameOrigin RequestMode = "same-origin"
	RequestModeNoCors     RequestMode = "no-cors"
	RequestModeCors       RequestMode = "cors"
)

func jsValueToRequestMode(val js.Value) RequestMode {
	return RequestMode(val.String())
}
