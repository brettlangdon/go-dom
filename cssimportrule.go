// Code generated DO NOT EDIT
// cssimportrule.go
package dom

import "syscall/js"

type CSSImportRuleIFace interface {
	GetCssText() string
	SetCssText(string)
	GetHref() string
	GetMedia() MediaList
	GetParentRule() CSSRule
	GetParentStyleSheet() CSSStyleSheet
	GetStyleSheet() CSSStyleSheet
	GetType() int
}
type CSSImportRule struct {
	Value
	CSSRule
}

func JSValueToCSSImportRule(val js.Value) CSSImportRule {
	return CSSImportRule{Value: JSValueToValue(val)}
}
func (v Value) AsCSSImportRule() CSSImportRule { return CSSImportRule{Value: v} }
func NewCSSImportRule(args ...interface{}) CSSImportRule {
	return CSSImportRule{Value: JSValueToValue(js.Global().Get("CSSImportRule").New(args...))}
}
func (c CSSImportRule) GetHref() string {
	val := c.Get("href")
	return val.String()
}
func (c CSSImportRule) GetMedia() MediaList {
	val := c.Get("media")
	return JSValueToMediaList(val.JSValue())
}
func (c CSSImportRule) GetStyleSheet() CSSStyleSheet {
	val := c.Get("styleSheet")
	return JSValueToCSSStyleSheet(val.JSValue())
}
