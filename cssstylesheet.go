// Code generated DO NOT EDIT
// cssstylesheet.go
package dom

import "syscall/js"

type CSSStyleSheetIFace interface {
	GetCssRules() CSSRuleList
	DeleteRule(args ...interface{})
	GetDisabled() bool
	SetDisabled(bool)
	GetHref() string
	InsertRule(args ...interface{}) int
	GetMedia() MediaList
	GetOwnerNode()
	GetOwnerRule() CSSRule
	GetParentStyleSheet() StyleSheet
	GetTitle() string
	GetType() string
}
type CSSStyleSheet struct {
	Value
	StyleSheet
}

func JSValueToCSSStyleSheet(val js.Value) CSSStyleSheet {
	return CSSStyleSheet{Value: Value{Value: val}}
}
func (v Value) AsCSSStyleSheet() CSSStyleSheet { return CSSStyleSheet{Value: v} }
func (c CSSStyleSheet) GetCssRules() CSSRuleList {
	val := c.Get("cssRules")
	return JSValueToCSSRuleList(val.JSValue())
}
func (c CSSStyleSheet) DeleteRule(args ...interface{}) {
	c.Call("deleteRule", args...)
}
func (c CSSStyleSheet) InsertRule(args ...interface{}) int {
	val := c.Call("insertRule", args...)
	return val.Int()
}
func (c CSSStyleSheet) GetOwnerRule() CSSRule {
	val := c.Get("ownerRule")
	return JSValueToCSSRule(val.JSValue())
}
