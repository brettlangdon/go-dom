// Code generated DO NOT EDIT
// csspagerule.go
package dom

import "syscall/js"

type CSSPageRuleIFace interface {
	GetCssRules() CSSRuleList
	GetCssText() string
	SetCssText(string)
	DeleteRule(args ...interface{})
	InsertRule(args ...interface{}) float64
	GetParentRule() CSSRule
	GetParentStyleSheet() CSSStyleSheet
	GetSelectorText() string
	SetSelectorText(string)
	GetStyle() CSSStyleDeclaration
	GetType() int
}
type CSSPageRule struct {
	Value
	CSSGroupingRule
	CSSRule
}

func JSValueToCSSPageRule(val js.Value) CSSPageRule { return CSSPageRule{Value: Value{Value: val}} }
func (v Value) AsCSSPageRule() CSSPageRule          { return CSSPageRule{Value: v} }
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
