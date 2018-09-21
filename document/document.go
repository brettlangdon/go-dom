// DO NOT EDIT - generated file
package document

import "syscall/js"
import dom "github.com/brettlangdon/go-dom/v1"

var d *dom.Document

func init() {
	d = dom.NewDocument(js.Global().Get("document"))
}
func GetBody() *dom.Element {
	return d.GetBody()
}
func CreateElement(tagName string) *dom.Element {
	return d.CreateElement(tagName)
}
func GetElementById(id string) *dom.Element {
	return d.GetElementById(id)
}
func GetElementsByName(name string) []*dom.Element {
	return d.GetElementsByName(name)
}
func Write(markup string) dom.Value {
	return d.Write(markup)
}
func AddEventListener(t string, listener *dom.Callback) dom.Value {
	return d.AddEventListener(t, listener)
}
func AppendChild(aChild *dom.Element) *dom.Element {
	return d.AppendChild(aChild)
}
func GetBaseURI() string {
	return d.GetBaseURI()
}
func GetFirstChild() *dom.Element {
	return d.GetFirstChild()
}
func GetLastChild() *dom.Element {
	return d.GetLastChild()
}
func GetNextSibling() *dom.Element {
	return d.GetNextSibling()
}
func GetPreviousSibling() *dom.Element {
	return d.GetPreviousSibling()
}
func GetParentElement() *dom.Element {
	return d.GetParentElement()
}
func GetRootElement() *dom.Element {
	return d.GetRootElement()
}
func GetPrefix() string {
	return d.GetPrefix()
}
func GetNodeName() string {
	return d.GetNodeName()
}
func GetTextContent() string {
	return d.GetTextContent()
}
func SetTextContent(v string) {
	d.SetTextContent(v)
}
func QuerySelector(selector string) *dom.Element {
	return d.QuerySelector(selector)
}
func QuerySelectorAll(selector string) []*dom.Element {
	return d.QuerySelectorAll(selector)
}
func AttachShadow(shadowRootInit dom.ShadowRootInit) *dom.ShadowRoot {
	return d.AttachShadow(shadowRootInit)
}
func GetClassName() string {
	return d.GetClassName()
}
func SetClassName(v string) {
	d.SetClassName(v)
}
func GetId() string {
	return d.GetId()
}
func SetId(v string) {
	d.SetId(v)
}
