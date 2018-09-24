// Code generated DO NOT EDIT
// cssstyledeclaration.go
package dom

import "syscall/js"

type CSSStyleDeclarationIFace interface {
	GetCamel_cased_attribute() string
	SetCamel_cased_attribute(string)
	GetCssFloat() string
	SetCssFloat(string)
	GetCssText() string
	SetCssText(string)
	GetDashed_attribute() string
	SetDashed_attribute(string)
	GetPropertyPriority(args ...interface{}) string
	GetPropertyValue(args ...interface{}) string
	Item(args ...interface{}) string
	GetLength() int
	GetParentRule() CSSRule
	RemoveProperty(args ...interface{}) string
	SetProperty(args ...interface{})
	GetWebkit_cased_attribute() string
	SetWebkit_cased_attribute(string)
}
type CSSStyleDeclaration struct {
	Value
}

func JSValueToCSSStyleDeclaration(val js.Value) CSSStyleDeclaration {
	return CSSStyleDeclaration{Value: JSValueToValue(val)}
}
func (v Value) AsCSSStyleDeclaration() CSSStyleDeclaration { return CSSStyleDeclaration{Value: v} }
func NewCSSStyleDeclaration(args ...interface{}) CSSStyleDeclaration {
	return CSSStyleDeclaration{Value: JSValueToValue(js.Global().Get("CSSStyleDeclaration").New(args...))}
}
func (c CSSStyleDeclaration) GetCamel_cased_attribute() string {
	val := c.Get("camel_cased_attribute")
	return val.String()
}
func (c CSSStyleDeclaration) SetCamel_cased_attribute(val string) {
	c.Set("camel_cased_attribute", val)
}
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
func (c CSSStyleDeclaration) GetDashed_attribute() string {
	val := c.Get("dashed_attribute")
	return val.String()
}
func (c CSSStyleDeclaration) SetDashed_attribute(val string) {
	c.Set("dashed_attribute", val)
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
func (c CSSStyleDeclaration) GetLength() int {
	val := c.Get("length")
	return val.Int()
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
func (c CSSStyleDeclaration) GetWebkit_cased_attribute() string {
	val := c.Get("webkit_cased_attribute")
	return val.String()
}
func (c CSSStyleDeclaration) SetWebkit_cased_attribute(val string) {
	c.Set("webkit_cased_attribute", val)
}
