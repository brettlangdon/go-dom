// Code generated DO NOT EDIT
// cssnamespacerule.go
package dom

import "syscall/js"

type CSSNamespaceRuleIFace interface {
	GetCssText() string
	SetCssText(string)
	GetNamespaceURI() string
	GetParentRule() CSSRule
	GetParentStyleSheet() CSSStyleSheet
	GetPrefix() string
	GetType() int
}
type CSSNamespaceRule struct {
	Value
}

func JSValueToCSSNamespaceRule(val js.Value) CSSNamespaceRule {
	return CSSNamespaceRule{Value: JSValueToValue(val)}
}
func (v Value) AsCSSNamespaceRule() CSSNamespaceRule { return CSSNamespaceRule{Value: v} }
func NewCSSNamespaceRule(args ...interface{}) CSSNamespaceRule {
	return CSSNamespaceRule{Value: JSValueToValue(js.Global().Get("CSSNamespaceRule").New(args...))}
}
func (c CSSNamespaceRule) GetCssText() string {
	val := c.Get("cssText")
	return val.String()
}
func (c CSSNamespaceRule) SetCssText(val string) {
	c.Set("cssText", val)
}
func (c CSSNamespaceRule) GetNamespaceURI() string {
	val := c.Get("namespaceURI")
	return val.String()
}
func (c CSSNamespaceRule) GetParentRule() CSSRule {
	val := c.Get("parentRule")
	return JSValueToCSSRule(val.JSValue())
}
func (c CSSNamespaceRule) GetParentStyleSheet() CSSStyleSheet {
	val := c.Get("parentStyleSheet")
	return JSValueToCSSStyleSheet(val.JSValue())
}
func (c CSSNamespaceRule) GetPrefix() string {
	val := c.Get("prefix")
	return val.String()
}
func (c CSSNamespaceRule) GetType() int {
	val := c.Get("type")
	return val.Int()
}
