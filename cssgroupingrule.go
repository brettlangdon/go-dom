// Code generated DO NOT EDIT
// cssgroupingrule.go
package dom

import "syscall/js"

type CSSGroupingRuleIFace interface {
	GetCssRules() CSSRuleList
	GetCssText() string
	SetCssText(string)
	DeleteRule(args ...interface{})
	InsertRule(args ...interface{}) float64
	GetParentRule() CSSRule
	GetParentStyleSheet() CSSStyleSheet
	GetType() int
}
type CSSGroupingRule struct {
	Value
	CSSRule
}

func jsValueToCSSGroupingRule(val js.Value) CSSGroupingRule {
	return CSSGroupingRule{Value: Value{Value: val}}
}
func (v Value) AsCSSGroupingRule() CSSGroupingRule { return CSSGroupingRule{Value: v} }
func (c CSSGroupingRule) GetCssRules() CSSRuleList {
	val := c.Get("cssRules")
	return jsValueToCSSRuleList(val.JSValue())
}
func (c CSSGroupingRule) DeleteRule(args ...interface{}) {
	c.Call("deleteRule", args...)
}
func (c CSSGroupingRule) InsertRule(args ...interface{}) float64 {
	val := c.Call("insertRule", args...)
	return val.Float()
}
