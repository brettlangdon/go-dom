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
}

func JSValueToCSSGroupingRule(val js.Value) CSSGroupingRule {
	return CSSGroupingRule{Value: JSValueToValue(val)}
}
func (v Value) AsCSSGroupingRule() CSSGroupingRule { return CSSGroupingRule{Value: v} }
func NewCSSGroupingRule(args ...interface{}) CSSGroupingRule {
	return CSSGroupingRule{Value: JSValueToValue(js.Global().Get("CSSGroupingRule").New(args...))}
}
func (c CSSGroupingRule) GetCssRules() CSSRuleList {
	val := c.Get("cssRules")
	return JSValueToCSSRuleList(val.JSValue())
}
func (c CSSGroupingRule) GetCssText() string {
	val := c.Get("cssText")
	return val.String()
}
func (c CSSGroupingRule) SetCssText(val string) {
	c.Set("cssText", val)
}
func (c CSSGroupingRule) DeleteRule(args ...interface{}) {
	c.Call("deleteRule", args...)
}
func (c CSSGroupingRule) InsertRule(args ...interface{}) int {
	val := c.Call("insertRule", args...)
	return val.Int()
}
func (c CSSGroupingRule) GetParentRule() CSSRule {
	val := c.Get("parentRule")
	return JSValueToCSSRule(val.JSValue())
}
func (c CSSGroupingRule) GetParentStyleSheet() CSSStyleSheet {
	val := c.Get("parentStyleSheet")
	return JSValueToCSSStyleSheet(val.JSValue())
}
func (c CSSGroupingRule) GetType() int {
	val := c.Get("type")
	return val.Int()
}
