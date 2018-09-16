// DO NOT EDIT - generated file
package dom

import "syscall/js"

type document struct {
	js.Value
}

func (d *document) GetBody() *Element {
	val := d.Get("body")
	return &Element{Value: val}
}
func (d *document) CreateElement(tagName string) *Element {
	val := d.Call("createElement", ToValue(tagName))
	return &Element{Value: val}
}
func (d *document) GetElementById(id string) *Element {
	val := d.Call("getElementById", ToValue(id))
	return &Element{Value: val}
}
func (d *document) QuerySelector(selector string) *Element {
	val := d.Call("querySelector", ToValue(selector))
	return &Element{Value: val}
}
func (d *document) QuerySelectorAll(selector string) []*Element {
	val := d.Call("querySelectorAll", ToValue(selector))
	elms := make([]*Element, 0)
	for i := 0; i < val.Length(); i += 1 {
		elms = append(elms, &Element{Value: val.Index(i)})
	}
	return elms
}
func (d *document) GetElementsByName(name string) []*Element {
	val := d.Call("getElementsByName", ToValue(name))
	elms := make([]*Element, 0)
	for i := 0; i < val.Length(); i += 1 {
		elms = append(elms, &Element{Value: val.Index(i)})
	}
	return elms
}
func (d *document) Write(markup string) {
	d.Call("write", ToValue(markup))
}
func (d *document) AddEventListener(t string, listener js.Callback) {
	d.Call("addEventListener", ToValue(t), ToValue(listener))
}
func (d *document) AppendChild(aChild *Element) *Element {
	val := d.Call("appendChild", ToValue(aChild))
	return &Element{Value: val}
}
func (d *document) GetBaseURI() string {
	val := d.Get("baseURI")
	return val.String()
}
func (d *document) GetFirstChild() *Element {
	val := d.Get("firstChild")
	return &Element{Value: val}
}
func (d *document) GetLastChild() *Element {
	val := d.Get("lastChild")
	return &Element{Value: val}
}
func (d *document) GetNextSibling() *Element {
	val := d.Get("nextSibling")
	return &Element{Value: val}
}
func (d *document) GetPreviousSibling() *Element {
	val := d.Get("previousSibling")
	return &Element{Value: val}
}
func (d *document) GetParentElement() *Element {
	val := d.Get("parentElement")
	return &Element{Value: val}
}
func (d *document) GetRootElement() *Element {
	val := d.Get("rootElement")
	return &Element{Value: val}
}
func (d *document) GetPrefix() string {
	val := d.Get("prefix")
	return val.String()
}
func (d *document) GetNodeName() string {
	val := d.Get("nodeName")
	return val.String()
}
func (d *document) GetTextContent() string {
	val := d.Get("textContent")
	return val.String()
}
func (d *document) SetTextContent(v string) {
	d.Set("textContent", v)
}
