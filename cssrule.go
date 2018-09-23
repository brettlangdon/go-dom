// Code generated DO NOT EDIT
// cssrule.go
package dom

import "syscall/js"

type CSSRuleIFace interface {
	GetCssText() string
	SetCssText(string)
	GetParentRule() CSSRule
	GetParentStyleSheet() CSSStyleSheet
	GetType() int
}
type CSSRule struct {
	Value
}

func JSValueToCSSRule(val js.Value) CSSRule { return CSSRule{Value: Value{Value: val}} }
func (v Value) AsCSSRule() CSSRule          { return CSSRule{Value: v} }
func (c CSSRule) GetCssText() string {
	val := c.Get("cssText")
	return val.String()
}
func (c CSSRule) SetCssText(val string) {
	c.Set("cssText", val)
}
func (c CSSRule) GetParentRule() CSSRule {
	val := c.Get("parentRule")
	return JSValueToCSSRule(val.JSValue())
}
func (c CSSRule) GetParentStyleSheet() CSSStyleSheet {
	val := c.Get("parentStyleSheet")
	return JSValueToCSSStyleSheet(val.JSValue())
}
func (c CSSRule) GetType() int {
	val := c.Get("type")
	return val.Int()
}
