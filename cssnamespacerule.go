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
	CSSRule
}

func JSValueToCSSNamespaceRule(val js.Value) CSSNamespaceRule {
	return CSSNamespaceRule{Value: JSValueToValue(val)}
}
func (v Value) AsCSSNamespaceRule() CSSNamespaceRule { return CSSNamespaceRule{Value: v} }
func NewCSSNamespaceRule(args ...interface{}) CSSNamespaceRule {
	return CSSNamespaceRule{Value: JSValueToValue(js.Global().Get("CSSNamespaceRule").New(args...))}
}
func (c CSSNamespaceRule) GetNamespaceURI() string {
	val := c.Get("namespaceURI")
	return val.String()
}
func (c CSSNamespaceRule) GetPrefix() string {
	val := c.Get("prefix")
	return val.String()
}
