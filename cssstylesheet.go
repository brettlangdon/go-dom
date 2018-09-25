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
}

func JSValueToCSSStyleSheet(val js.Value) CSSStyleSheet {
	return CSSStyleSheet{Value: JSValueToValue(val)}
}
func (v Value) AsCSSStyleSheet() CSSStyleSheet { return CSSStyleSheet{Value: v} }
func NewCSSStyleSheet(args ...interface{}) CSSStyleSheet {
	return CSSStyleSheet{Value: JSValueToValue(js.Global().Get("CSSStyleSheet").New(args...))}
}
func (c CSSStyleSheet) GetCssRules() CSSRuleList {
	val := c.Get("cssRules")
	return JSValueToCSSRuleList(val.JSValue())
}
func (c CSSStyleSheet) DeleteRule(args ...interface{}) {
	c.Call("deleteRule", args...)
}
func (c CSSStyleSheet) GetDisabled() bool {
	val := c.Get("disabled")
	return val.Bool()
}
func (c CSSStyleSheet) SetDisabled(val bool) {
	c.Set("disabled", val)
}
func (c CSSStyleSheet) GetHref() string {
	val := c.Get("href")
	return val.String()
}
func (c CSSStyleSheet) InsertRule(args ...interface{}) int {
	val := c.Call("insertRule", args...)
	return val.Int()
}
func (c CSSStyleSheet) GetMedia() MediaList {
	val := c.Get("media")
	return JSValueToMediaList(val.JSValue())
}
func (c CSSStyleSheet) GetOwnerNode() Value {
	val := c.Get("ownerNode")
	return val
}
func (c CSSStyleSheet) GetOwnerRule() CSSRule {
	val := c.Get("ownerRule")
	return JSValueToCSSRule(val.JSValue())
}
func (c CSSStyleSheet) GetParentStyleSheet() StyleSheet {
	val := c.Get("parentStyleSheet")
	return JSValueToStyleSheet(val.JSValue())
}
func (c CSSStyleSheet) GetTitle() string {
	val := c.Get("title")
	return val.String()
}
func (c CSSStyleSheet) GetType() string {
	val := c.Get("type")
	return val.String()
}
