// DO NOT EDIT - generated file
package dom

import "syscall/js"

type ShadowRoot struct {
	Value
}

func NewShadowRoot(v js.Value) *ShadowRoot {
	val := Value{Value: v}
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return val.ToShadowRoot()
}
func (v Value) ToShadowRoot() *ShadowRoot { return &ShadowRoot{Value: v} }
func (s *ShadowRoot) GetMode() string {
	val := Value{Value: s.Get("mode")}
	return val.String()
}
func (s *ShadowRoot) GetHost() *Element {
	val := Value{Value: s.Get("host")}
	return NewElement(val.JSValue())
}
func (s *ShadowRoot) GetInnerHTML() string {
	val := Value{Value: s.Get("innerHTML")}
	return val.String()
}
func (s *ShadowRoot) SetInnerHTML(v string) {
	s.Set("innerHTML", v)
}
func (s *ShadowRoot) AddEventListener(t string, listener *Callback) Value {
	val := Value{Value: s.Call("addEventListener", ToJSValue(t), ToJSValue(listener))}
	return val
}
func (s *ShadowRoot) AppendChild(aChild *Element) *Element {
	val := Value{Value: s.Call("appendChild", ToJSValue(aChild))}
	return NewElement(val.JSValue())
}
func (s *ShadowRoot) GetBaseURI() string {
	val := Value{Value: s.Get("baseURI")}
	return val.String()
}
func (s *ShadowRoot) GetFirstChild() *Element {
	val := Value{Value: s.Get("firstChild")}
	return NewElement(val.JSValue())
}
func (s *ShadowRoot) GetLastChild() *Element {
	val := Value{Value: s.Get("lastChild")}
	return NewElement(val.JSValue())
}
func (s *ShadowRoot) GetNextSibling() *Element {
	val := Value{Value: s.Get("nextSibling")}
	return NewElement(val.JSValue())
}
func (s *ShadowRoot) GetPreviousSibling() *Element {
	val := Value{Value: s.Get("previousSibling")}
	return NewElement(val.JSValue())
}
func (s *ShadowRoot) GetParentElement() *Element {
	val := Value{Value: s.Get("parentElement")}
	return NewElement(val.JSValue())
}
func (s *ShadowRoot) GetRootElement() *Element {
	val := Value{Value: s.Get("rootElement")}
	return NewElement(val.JSValue())
}
func (s *ShadowRoot) GetPrefix() string {
	val := Value{Value: s.Get("prefix")}
	return val.String()
}
func (s *ShadowRoot) GetNodeName() string {
	val := Value{Value: s.Get("nodeName")}
	return val.String()
}
func (s *ShadowRoot) GetTextContent() string {
	val := Value{Value: s.Get("textContent")}
	return val.String()
}
func (s *ShadowRoot) SetTextContent(v string) {
	s.Set("textContent", v)
}
