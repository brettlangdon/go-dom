// Code generated DO NOT EDIT
// cssgroupingrule.go
package dom

import "syscall/js"

type CSSGroupingRuleIFace interface {
	GetCssRules() CSSRuleList
	GetCssText() string
	SetCssText(string)
	DeleteRule(args ...interface{})
	InsertRule(args ...interface{}) int
	GetParentRule() CSSRule
	GetParentStyleSheet() CSSStyleSheet
	GetType() int
}
type CSSGroupingRule struct {
	Value
	CSSRule
}

func JSValueToCSSGroupingRule(val js.Value) CSSGroupingRule {
	return CSSGroupingRule{Value: Value{Value: val}}
}
func (v Value) AsCSSGroupingRule() CSSGroupingRule { return CSSGroupingRule{Value: v} }
func (c CSSGroupingRule) GetCssRules() CSSRuleList {
	val := c.Get("cssRules")
	return JSValueToCSSRuleList(val.JSValue())
}
func (c CSSGroupingRule) DeleteRule(args ...interface{}) {
	c.Call("deleteRule", args...)
}
func (c CSSGroupingRule) InsertRule(args ...interface{}) int {
	val := c.Call("insertRule", args...)
	return val.Int()
}
