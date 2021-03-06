// Code generated DO NOT EDIT
// shadowrootmode.go
package dom

import "syscall/js"

type ShadowRootMode string

const (
	ShadowRootModeOpen   ShadowRootMode = "open"
	ShadowRootModeClosed ShadowRootMode = "closed"
)

func JSValueToShadowRootMode(val js.Value) ShadowRootMode {
	return ShadowRootMode(val.String())
}
