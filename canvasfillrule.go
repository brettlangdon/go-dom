// Code generated DO NOT EDIT
// canvasfillrule.go
package dom

import "syscall/js"

type CanvasFillRule string

const (
	CanvasFillRuleNonzero CanvasFillRule = "nonzero"
	CanvasFillRuleEvenodd CanvasFillRule = "evenodd"
)

func JSValueToCanvasFillRule(val js.Value) CanvasFillRule {
	return CanvasFillRule(val.String())
}
