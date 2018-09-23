// Code generated DO NOT EDIT
// cssrulelist.go
package dom

import "syscall/js"

type CSSRuleListIFace interface {
	Item(args ...interface{}) CSSRule
	GetLength() int
}
type CSSRuleList struct {
	Value
}

func JSValueToCSSRuleList(val js.Value) CSSRuleList { return CSSRuleList{Value: Value{Value: val}} }
func (v Value) AsCSSRuleList() CSSRuleList          { return CSSRuleList{Value: v} }
func (c CSSRuleList) Item(args ...interface{}) CSSRule {
	val := c.Call("item", args...)
	return JSValueToCSSRule(val.JSValue())
}
func (c CSSRuleList) GetLength() int {
	val := c.Get("length")
	return val.Int()
}
