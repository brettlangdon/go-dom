// Code generated DO NOT EDIT
// requestcredentials.go
package dom

import "syscall/js"

type RequestCredentials string

const (
	RequestCredentialsOmit       RequestCredentials = "omit"
	RequestCredentialsSameOrigin RequestCredentials = "same-origin"
	RequestCredentialsInclude    RequestCredentials = "include"
)

func jsValueToRequestCredentials(val js.Value) RequestCredentials {
	return RequestCredentials(val.String())
}
