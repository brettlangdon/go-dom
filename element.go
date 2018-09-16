// DO NOT EDIT - generated file
package dom

import "syscall/js"

type Element struct {
	js.Value
}

func (e *Element) GetClassName() string {
	val := e.Get("className")
	return val.String()
}
func (e *Element) GetId() string {
	val := e.Get("id")
	return val.String()
}
func (e *Element) AddEventListener(t string, listener js.Callback) {
	e.Call("addEventListener", ToValue(t), ToValue(listener))
}
func (e *Element) AppendChild(aChild *Element) *Element {
	val := e.Call("appendChild", ToValue(aChild))
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
