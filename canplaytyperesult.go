// Code generated DO NOT EDIT
// canplaytyperesult.go
package dom

import "syscall/js"

type CanPlayTypeResult string

const (
	CanPlayTypeResultEmpty    CanPlayTypeResult = "Empty"
	CanPlayTypeResultMaybe    CanPlayTypeResult = "maybe"
	CanPlayTypeResultProbably CanPlayTypeResult = "probably"
)

func jsValueToCanPlayTypeResult(val js.Value) CanPlayTypeResult {
	return CanPlayTypeResult(val.String())
}
