// Code generated DO NOT EDIT
// domimplementation.go
package dom

import "syscall/js"

type DOMImplementationIFace interface {
	CreateDocument(args ...interface{}) XMLDocument
	CreateDocumentType(args ...interface{}) DocumentType
	CreateHTMLDocument(args ...interface{}) Document
	HasFeature(args ...interface{}) bool
}
type DOMImplementation struct {
	Value
}

func JSValueToDOMImplementation(val js.Value) DOMImplementation {
	return DOMImplementation{Value: JSValueToValue(val)}
}
func (v Value) AsDOMImplementation() DOMImplementation { return DOMImplementation{Value: v} }
func NewDOMImplementation(args ...interface{}) DOMImplementation {
	return DOMImplementation{Value: JSValueToValue(js.Global().Get("DOMImplementation").New(args...))}
}
func (d DOMImplementation) CreateDocument(args ...interface{}) XMLDocument {
	val := d.Call("createDocument", args...)
	return JSValueToXMLDocument(val.JSValue())
}
func (d DOMImplementation) CreateDocumentType(args ...interface{}) DocumentType {
	val := d.Call("createDocumentType", args...)
	return JSValueToDocumentType(val.JSValue())
}
func (d DOMImplementation) CreateHTMLDocument(args ...interface{}) Document {
	val := d.Call("createHTMLDocument", args...)
	return JSValueToDocument(val.JSValue())
}
func (d DOMImplementation) HasFeature(args ...interface{}) bool {
	val := d.Call("hasFeature", args...)
	return val.Bool()
}
