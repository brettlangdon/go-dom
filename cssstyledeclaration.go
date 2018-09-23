// Code generated DO NOT EDIT
// cssstyledeclaration.go
package dom

import "syscall/js"

type CSSStyleDeclarationIFace interface {
	GetCssFloat() string
	SetCssFloat(string)
	GetCssText() string
	SetCssText(string)
	GetPropertyPriority(args ...interface{}) string
	GetPropertyValue(args ...interface{}) string
	Item(args ...interface{}) string
	GetLength() float64
	GetParentRule() CSSRule
	RemoveProperty(args ...interface{}) string
	SetProperty(args ...interface{})
}
type CSSStyleDeclaration struct {
	Value
}

func JSValueToCSSStyleDeclaration(val js.Value) CSSStyleDeclaration {
	return CSSStyleDeclaration{Value: Value{Value: val}}
}
func (v Value) AsCSSStyleDeclaration() CSSStyleDeclaration { return CSSStyleDeclaration{Value: v} }
func (c CSSStyleDeclaration) GetCssFloat() string {
	val := c.Get("cssFloat")
	return val.String()
}
func (c CSSStyleDeclaration) SetCssFloat(val string) {
	c.Set("cssFloat", val)
}
func (c CSSStyleDeclaration) GetCssText() string {
	val := c.Get("cssText")
	return val.String()
}
func (c CSSStyleDeclaration) SetCssText(val string) {
	c.Set("cssText", val)
}
func (c CSSStyleDeclaration) GetPropertyPriority(args ...interface{}) string {
	val := c.Call("getPropertyPriority", args...)
	return val.String()
}
func (c CSSStyleDeclaration) GetPropertyValue(args ...interface{}) string {
	val := c.Call("getPropertyValue", args...)
	return val.String()
}
func (c CSSStyleDeclaration) Item(args ...interface{}) string {
	val := c.Call("item", args...)
	return val.String()
}
func (c CSSStyleDeclaration) GetLength() float64 {
	val := c.Get("length")
	return val.Float()
}
func (c CSSStyleDeclaration) GetParentRule() CSSRule {
	val := c.Get("parentRule")
	return JSValueToCSSRule(val.JSValue())
}
func (c CSSStyleDeclaration) RemoveProperty(args ...interface{}) string {
	val := c.Call("removeProperty", args...)
	return val.String()
}
func (c CSSStyleDeclaration) SetProperty(args ...interface{}) {
	c.Call("setProperty", args...)
}
