// Code generated DO NOT EDIT
// xmlhttprequestupload.go
package dom

import "syscall/js"

type XMLHttpRequestUploadIFace interface {
	AddEventListener(args ...interface{})
	DispatchEvent(args ...interface{}) bool
	GetOnabort() EventHandler
	SetOnabort(EventHandler)
	GetOnerror() EventHandler
	SetOnerror(EventHandler)
	GetOnload() EventHandler
	SetOnload(EventHandler)
	GetOnloadend() EventHandler
	SetOnloadend(EventHandler)
	GetOnloadstart() EventHandler
	SetOnloadstart(EventHandler)
	GetOnprogress() EventHandler
	SetOnprogress(EventHandler)
	GetOntimeout() EventHandler
	SetOntimeout(EventHandler)
	RemoveEventListener(args ...interface{})
}
type XMLHttpRequestUpload struct {
	Value
	XMLHttpRequestEventTarget
	EventTarget
}

func JSValueToXMLHttpRequestUpload(val js.Value) XMLHttpRequestUpload {
	return XMLHttpRequestUpload{Value: Value{Value: val}}
}
func (v Value) AsXMLHttpRequestUpload() XMLHttpRequestUpload { return XMLHttpRequestUpload{Value: v} }
