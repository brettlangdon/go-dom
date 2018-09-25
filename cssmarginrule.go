// Code generated DO NOT EDIT
// cssmarginrule.go
package dom

import "syscall/js"

type CSSMarginRuleIFace interface {
	GetCssText() string
	SetCssText(string)
	GetName() string
	GetParentRule() CSSRule
	GetParentStyleSheet() CSSStyleSheet
	GetStyle() CSSStyleDeclaration
	GetType() int
}
type CSSMarginRule struct {
	Value
}

func JSValueToCSSMarginRule(val js.Value) CSSMarginRule {
	return CSSMarginRule{Value: JSValueToValue(val)}
}
func (v Value) AsCSSMarginRule() CSSMarginRule { return CSSMarginRule{Value: v} }
func NewCSSMarginRule(args ...interface{}) CSSMarginRule {
	return CSSMarginRule{Value: JSValueToValue(js.Global().Get("CSSMarginRule").New(args...))}
}
func (c CSSMarginRule) GetCssText() string {
	val := c.Get("cssText")
	return val.String()
}
func (c CSSMarginRule) SetCssText(val string) {
	c.Set("cssText", val)
}
func (c CSSMarginRule) GetName() string {
	val := c.Get("name")
	return val.String()
}
func (c CSSMarginRule) GetParentRule() CSSRule {
	val := c.Get("parentRule")
	return JSValueToCSSRule(val.JSValue())
}
func (c CSSMarginRule) GetParentStyleSheet() CSSStyleSheet {
	val := c.Get("parentStyleSheet")
	return JSValueToCSSStyleSheet(val.JSValue())
}
func (c CSSMarginRule) GetStyle() CSSStyleDeclaration {
	val := c.Get("style")
	return JSValueToCSSStyleDeclaration(val.JSValue())
}
func (c CSSMarginRule) GetType() int {
	val := c.Get("type")
	return val.Int()
}
