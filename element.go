// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Element struct {
	js.Value
}

func (e *Element) JSValue() js.Value { return e.Value }
func (e *Element) AddEventListener(t string, listener js.Callback) {
	e.Call("addEventListener", ToJSValue(t), ToJSValue(listener))
}
func (e *Element) AppendChild(aChild *Element) *Element {
	val := e.Call("appendChild", ToJSValue(aChild))
	return &Element{Value: val}
}
func (e *Element) GetBaseURI() string {
	val := e.Get("baseURI")
	return val.String()
}
func (e *Element) GetFirstChild() *Element {
	val := e.Get("firstChild")
	return &Element{Value: val}
}
func (e *Element) GetLastChild() *Element {
	val := e.Get("lastChild")
	return &Element{Value: val}
}
func (e *Element) GetNextSibling() *Element {
	val := e.Get("nextSibling")
	return &Element{Value: val}
}
func (e *Element) GetPreviousSibling() *Element {
	val := e.Get("previousSibling")
	return &Element{Value: val}
}
func (e *Element) GetParentElement() *Element {
	val := e.Get("parentElement")
	return &Element{Value: val}
}
func (e *Element) GetRootElement() *Element {
	val := e.Get("rootElement")
	return &Element{Value: val}
}
func (e *Element) GetPrefix() string {
	val := e.Get("prefix")
	return val.String()
}
func (e *Element) GetNodeName() string {
	val := e.Get("nodeName")
	return val.String()
}
func (e *Element) GetTextContent() string {
	val := e.Get("textContent")
	return val.String()
}
func (e *Element) SetTextContent(v string) {
	e.Set("textContent", v)
}
func (e *Element) QuerySelector(selector string) *Element {
	val := e.Call("querySelector", ToJSValue(selector))
	return &Element{Value: val}
}
func (e *Element) QuerySelectorAll(selector string) []*Element {
	val := e.Call("querySelectorAll", ToJSValue(selector))
	elms := make([]*Element, 0)
	for i := 0; i < val.Length(); i += 1 {
		elms = append(elms, &Element{Value: val.Index(i)})
	}
	return elms
}
func (e *Element) GetClassName() string {
	val := e.Get("className")
	return val.String()
}
func (e *Element) SetClassName(v string) {
	e.Set("className", v)
}
func (e *Element) GetId() string {
	val := e.Get("id")
	return val.String()
}
func (e *Element) SetId(v string) {
	e.Set("id", v)
}
