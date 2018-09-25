// Code generated DO NOT EDIT
// csspagerule.go
package dom

import "syscall/js"

type CSSPageRuleIFace interface {
	GetCssRules() CSSRuleList
	GetCssText() string
	SetCssText(string)
	DeleteRule(args ...interface{})
	InsertRule(args ...interface{}) int
	GetParentRule() CSSRule
	GetParentStyleSheet() CSSStyleSheet
	GetSelectorText() string
	SetSelectorText(string)
	GetStyle() CSSStyleDeclaration
	GetType() int
}
type CSSPageRule struct {
	Value
}

func JSValueToCSSPageRule(val js.Value) CSSPageRule { return CSSPageRule{Value: JSValueToValue(val)} }
func (v Value) AsCSSPageRule() CSSPageRule          { return CSSPageRule{Value: v} }
func NewCSSPageRule(args ...interface{}) CSSPageRule {
	return CSSPageRule{Value: JSValueToValue(js.Global().Get("CSSPageRule").New(args...))}
}
func (c CSSPageRule) GetCssRules() CSSRuleList {
	val := c.Get("cssRules")
	return JSValueToCSSRuleList(val.JSValue())
}
func (c CSSPageRule) GetCssText() string {
	val := c.Get("cssText")
	return val.String()
}
func (c CSSPageRule) SetCssText(val string) {
	c.Set("cssText", val)
}
func (c CSSPageRule) DeleteRule(args ...interface{}) {
	c.Call("deleteRule", args...)
}
func (c CSSPageRule) InsertRule(args ...interface{}) int {
	val := c.Call("insertRule", args...)
	return val.Int()
}
func (c CSSPageRule) GetParentRule() CSSRule {
	val := c.Get("parentRule")
	return JSValueToCSSRule(val.JSValue())
}
func (c CSSPageRule) GetParentStyleSheet() CSSStyleSheet {
	val := c.Get("parentStyleSheet")
	return JSValueToCSSStyleSheet(val.JSValue())
}
func (c CSSPageRule) GetSelectorText() string {
	val := c.Get("selectorText")
	return val.String()
}
func (c CSSPageRule) SetSelectorText(val string) {
	c.Set("selectorText", val)
}
func (c CSSPageRule) GetStyle() CSSStyleDeclaration {
	val := c.Get("style")
	return JSValueToCSSStyleDeclaration(val.JSValue())
}
func (c CSSPageRule) GetType() int {
	val := c.Get("type")
	return val.Int()
}
