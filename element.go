// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Element struct {
	Value
}

func NewElement(v js.Value) *Element {
	val := Value{Value: v}
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return val.ToElement()
}
func (v Value) ToElement() *Element { return &Element{Value: v} }
func (e *Element) AddEventListener(t string, listener *Callback) Value {
	val := Value{Value: e.Call("addEventListener", ToJSValue(t), ToJSValue(listener))}
	return val
}
func (e *Element) AppendChild(aChild *Element) *Element {
	val := Value{Value: e.Call("appendChild", ToJSValue(aChild))}
	return NewElement(val.JSValue())
}
func (e *Element) GetBaseURI() string {
	val := Value{Value: e.Get("baseURI")}
	return val.String()
}
func (e *Element) GetFirstChild() *Element {
	val := Value{Value: e.Get("firstChild")}
	return NewElement(val.JSValue())
}
func (e *Element) GetLastChild() *Element {
	val := Value{Value: e.Get("lastChild")}
	return NewElement(val.JSValue())
}
func (e *Element) GetNextSibling() *Element {
	val := Value{Value: e.Get("nextSibling")}
	return NewElement(val.JSValue())
}
func (e *Element) GetPreviousSibling() *Element {
	val := Value{Value: e.Get("previousSibling")}
	return NewElement(val.JSValue())
}
func (e *Element) GetParentElement() *Element {
	val := Value{Value: e.Get("parentElement")}
	return NewElement(val.JSValue())
}
func (e *Element) GetRootElement() *Element {
	val := Value{Value: e.Get("rootElement")}
	return NewElement(val.JSValue())
}
func (e *Element) GetPrefix() string {
	val := Value{Value: e.Get("prefix")}
	return val.String()
}
func (e *Element) GetNodeName() string {
	val := Value{Value: e.Get("nodeName")}
	return val.String()
}
func (e *Element) GetTextContent() string {
	val := Value{Value: e.Get("textContent")}
	return val.String()
}
func (e *Element) SetTextContent(v string) {
	e.Set("textContent", v)
}
func (e *Element) QuerySelector(selector string) *Element {
	val := Value{Value: e.Call("querySelector", ToJSValue(selector))}
	return NewElement(val.JSValue())
}
func (e *Element) QuerySelectorAll(selector string) []*Element {
	val := Value{Value: e.Call("querySelectorAll", ToJSValue(selector))}
	elms := make([]*Element, 0)
	for i := 0; i < val.Length(); i += 1 {
		elms = append(elms, NewElement(val.Index(i)))
	}
	return elms
}
func (e *Element) AttachShadow(shadowRootInit ShadowRootInit) *ShadowRoot {
	val := Value{Value: e.Call("attachShadow", ToJSValue(shadowRootInit))}
	return NewShadowRoot(val.JSValue())
}
func (e *Element) GetClassName() string {
	val := Value{Value: e.Get("className")}
	return val.String()
}
func (e *Element) SetClassName(v string) {
	e.Set("className", v)
}
func (e *Element) GetId() string {
	val := Value{Value: e.Get("id")}
	return val.String()
}
func (e *Element) SetId(v string) {
	e.Set("id", v)
}
