// Code generated DO NOT EDIT
// cssstylerule.go
package dom

import "syscall/js"

type CSSStyleRuleIFace interface {
	GetCssText() string
	SetCssText(string)
	GetParentRule() CSSRule
	GetParentStyleSheet() CSSStyleSheet
	GetSelectorText() string
	SetSelectorText(string)
	GetStyle() CSSStyleDeclaration
	GetType() int
}
type CSSStyleRule struct {
	Value
	CSSRule
}

func jsValueToCSSStyleRule(val js.Value) CSSStyleRule { return CSSStyleRule{Value: Value{Value: val}} }
func (v Value) AsCSSStyleRule() CSSStyleRule          { return CSSStyleRule{Value: v} }
func (c CSSStyleRule) GetSelectorText() string {
	val := c.Get("selectorText")
	return val.String()
}
func (c CSSStyleRule) SetSelectorText(val string) {
	c.Set("selectorText", val)
}
func (c CSSStyleRule) GetStyle() CSSStyleDeclaration {
	val := c.Get("style")
	return jsValueToCSSStyleDeclaration(val.JSValue())
}
