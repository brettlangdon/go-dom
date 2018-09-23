// Code generated DO NOT EDIT
// response.go
package dom

import "syscall/js"

type ResponseIFace interface {
	ArrayBuffer(args ...interface{})
	Blob(args ...interface{})
	GetBody() Value
	GetBodyUsed() bool
	Clone(args ...interface{}) Response
	Error(args ...interface{}) Response
	FormData(args ...interface{})
	GetHeaders() Headers
	Json(args ...interface{})
	GetOk() bool
	Redirect(args ...interface{}) Response
	GetRedirected() bool
	GetStatus() int
	GetStatusText() []byte
	Text(args ...interface{})
	GetTrailer()
	GetType() ResponseType
	GetUrl() string
}
type Response struct {
	Value
}

func JSValueToResponse(val js.Value) Response { return Response{Value: Value{Value: val}} }
func (v Value) AsResponse() Response          { return Response{Value: v} }
func (r Response) ArrayBuffer(args ...interface{}) {
	r.Call("arrayBuffer", args...)
}
func (r Response) Blob(args ...interface{}) {
	r.Call("blob", args...)
}
func (r Response) GetBody() Value {
	val := r.Get("body")
	return val
}
func (r Response) GetBodyUsed() bool {
	val := r.Get("bodyUsed")
	return val.Bool()
}
func (r Response) Clone(args ...interface{}) Response {
	val := r.Call("clone", args...)
	return JSValueToResponse(val.JSValue())
}
func (r Response) Error(args ...interface{}) Response {
	val := r.Call("error", args...)
	return JSValueToResponse(val.JSValue())
}
func (r Response) FormData(args ...interface{}) {
	r.Call("formData", args...)
}
func (r Response) GetHeaders() Headers {
	val := r.Get("headers")
	return JSValueToHeaders(val.JSValue())
}
func (r Response) Json(args ...interface{}) {
	r.Call("json", args...)
}
func (r Response) GetOk() bool {
	val := r.Get("ok")
	return val.Bool()
}
func (r Response) Redirect(args ...interface{}) Response {
	val := r.Call("redirect", args...)
	return JSValueToResponse(val.JSValue())
}
func (r Response) GetRedirected() bool {
	val := r.Get("redirected")
	return val.Bool()
}
func (r Response) GetStatus() int {
	val := r.Get("status")
	return val.Int()
}
func (r Response) GetStatusText() []byte {
	val := r.Get("statusText")
	return []byte(val.String())
}
func (r Response) Text(args ...interface{}) {
	r.Call("text", args...)
}
func (r Response) GetTrailer() Value {
	val := r.Get("trailer")
	return val
}
func (r Response) GetType() ResponseType {
	val := r.Get("type")
	return JSValueToResponseType(val.JSValue())
}
func (r Response) GetUrl() string {
	val := r.Get("url")
	return val.String()
}
