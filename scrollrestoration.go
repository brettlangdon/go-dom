// Code generated DO NOT EDIT
// scrollrestoration.go
package dom

import "syscall/js"

type ScrollRestoration string

const (
	ScrollRestorationAuto   ScrollRestoration = "auto"
	ScrollRestorationManual ScrollRestoration = "manual"
)

func jsValueToScrollRestoration(val js.Value) ScrollRestoration {
	return ScrollRestoration(val.String())
}
