// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Document struct {
	Value
}

func NewDocument(v js.Value) *Document {
	val := Value{Value: v}
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return val.ToDocument()
}
func (v Value) ToDocument() *Document { return &Document{Value: v} }
func (d *Document) GetBody() *Element {
	val := Value{Value: d.Get("body")}
	return NewElement(val.JSValue())
}
func (d *Document) CreateElement(tagName string) *Element {
	val := Value{Value: d.Call("createElement", ToJSValue(tagName))}
	return NewElement(val.JSValue())
}
func (d *Document) GetElementById(id string) *Element {
	val := Value{Value: d.Call("getElementById", ToJSValue(id))}
	return NewElement(val.JSValue())
}
func (d *Document) GetElementsByName(name string) []*Element {
	val := Value{Value: d.Call("getElementsByName", ToJSValue(name))}
	elms := make([]*Element, 0)
	for i := 0; i < val.Length(); i += 1 {
		elms = append(elms, NewElement(val.Index(i)))
	}
	return elms
}
func (d *Document) Write(markup string) Value {
	val := Value{Value: d.Call("write", ToJSValue(markup))}
	return val
}
func (d *Document) AddEventListener(t string, listener *Callback) Value {
	val := Value{Value: d.Call("addEventListener", ToJSValue(t), ToJSValue(listener))}
	return val
}
func (d *Document) AppendChild(aChild *Element) *Element {
	val := Value{Value: d.Call("appendChild", ToJSValue(aChild))}
	return NewElement(val.JSValue())
}
func (d *Document) GetBaseURI() string {
	val := Value{Value: d.Get("baseURI")}
	return val.String()
}
func (d *Document) GetFirstChild() *Element {
	val := Value{Value: d.Get("firstChild")}
	return NewElement(val.JSValue())
}
func (d *Document) GetLastChild() *Element {
	val := Value{Value: d.Get("lastChild")}
	return NewElement(val.JSValue())
}
func (d *Document) GetNextSibling() *Element {
	val := Value{Value: d.Get("nextSibling")}
	return NewElement(val.JSValue())
}
func (d *Document) GetPreviousSibling() *Element {
	val := Value{Value: d.Get("previousSibling")}
	return NewElement(val.JSValue())
}
func (d *Document) GetParentElement() *Element {
	val := Value{Value: d.Get("parentElement")}
	return NewElement(val.JSValue())
}
func (d *Document) GetRootElement() *Element {
	val := Value{Value: d.Get("rootElement")}
	return NewElement(val.JSValue())
}
func (d *Document) GetPrefix() string {
	val := Value{Value: d.Get("prefix")}
	return val.String()
}
func (d *Document) GetNodeName() string {
	val := Value{Value: d.Get("nodeName")}
	return val.String()
}
func (d *Document) GetTextContent() string {
	val := Value{Value: d.Get("textContent")}
	return val.String()
}
func (d *Document) SetTextContent(v string) {
	d.Set("textContent", v)
}
func (d *Document) QuerySelector(selector string) *Element {
	val := Value{Value: d.Call("querySelector", ToJSValue(selector))}
	return NewElement(val.JSValue())
}
func (d *Document) QuerySelectorAll(selector string) []*Element {
	val := Value{Value: d.Call("querySelectorAll", ToJSValue(selector))}
	elms := make([]*Element, 0)
	for i := 0; i < val.Length(); i += 1 {
		elms = append(elms, NewElement(val.Index(i)))
	}
	return elms
}
func (d *Document) AttachShadow(shadowRootInit ShadowRootInit) *ShadowRoot {
	val := Value{Value: d.Call("attachShadow", ToJSValue(shadowRootInit))}
	return NewShadowRoot(val.JSValue())
}
func (d *Document) GetClassName() string {
	val := Value{Value: d.Get("className")}
	return val.String()
}
func (d *Document) SetClassName(v string) {
	d.Set("className", v)
}
func (d *Document) GetId() string {
	val := Value{Value: d.Get("id")}
	return val.String()
}
func (d *Document) SetId(v string) {
	d.Set("id", v)
}
