// Code generated DO NOT EDIT
// requestredirect.go
package dom

import "syscall/js"

type RequestRedirect string

const (
	RequestRedirectFollow RequestRedirect = "follow"
	RequestRedirectError  RequestRedirect = "error"
	RequestRedirectManual RequestRedirect = "manual"
)

func jsValueToRequestRedirect(val js.Value) RequestRedirect {
	return RequestRedirect(val.String())
}
