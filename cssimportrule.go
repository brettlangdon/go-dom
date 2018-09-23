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

func jsValueToCSSImportRule(val js.Value) CSSImportRule {
	return CSSImportRule{Value: Value{Value: val}}
}
func (v Value) AsCSSImportRule() CSSImportRule { return CSSImportRule{Value: v} }
func (c CSSImportRule) GetHref() string {
	val := c.Get("href")
	return val.String()
}
func (c CSSImportRule) GetMedia() MediaList {
	val := c.Get("media")
	return jsValueToMediaList(val.JSValue())
}
func (c CSSImportRule) GetStyleSheet() CSSStyleSheet {
	val := c.Get("styleSheet")
	return jsValueToCSSStyleSheet(val.JSValue())
}
