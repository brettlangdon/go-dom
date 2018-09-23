// Code generated DO NOT EDIT
// scrollrestoration.go
package dom

import "syscall/js"

type ScrollRestoration string

const (
	ScrollRestorationAuto   ScrollRestoration = "auto"
	ScrollRestorationManual ScrollRestoration = "manual"
)

func JSValueToScrollRestoration(val js.Value) ScrollRestoration {
	return ScrollRestoration(val.String())
}
