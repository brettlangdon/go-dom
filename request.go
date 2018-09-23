// Code generated DO NOT EDIT
// request.go
package dom

import "syscall/js"

type RequestIFace interface {
	ArrayBuffer(args ...interface{})
	Blob(args ...interface{})
	GetBody() Value
	GetBodyUsed() bool
	GetCache() RequestCache
	Clone(args ...interface{}) Request
	GetCredentials() RequestCredentials
	GetDestination() RequestDestination
	FormData(args ...interface{})
	GetHeaders() Headers
	GetIntegrity() string
	GetIsHistoryNavigation() bool
	GetIsReloadNavigation() bool
	Json(args ...interface{})
	GetKeepalive() bool
	GetMethod() []byte
	GetMode() RequestMode
	GetRedirect() RequestRedirect
	GetReferrer() string
	GetReferrerPolicy() ReferrerPolicy
	GetSignal() AbortSignal
	Text(args ...interface{})
	GetUrl() string
}
type Request struct {
	Value
}

func JSValueToRequest(val js.Value) Request { return Request{Value: Value{Value: val}} }
func (v Value) AsRequest() Request          { return Request{Value: v} }
func (r Request) ArrayBuffer(args ...interface{}) {
	r.Call("arrayBuffer", args...)
}
func (r Request) Blob(args ...interface{}) {
	r.Call("blob", args...)
}
func (r Request) GetBody() Value {
	val := r.Get("body")
	return val
}
func (r Request) GetBodyUsed() bool {
	val := r.Get("bodyUsed")
	return val.Bool()
}
func (r Request) GetCache() RequestCache {
	val := r.Get("cache")
	return JSValueToRequestCache(val.JSValue())
}
func (r Request) Clone(args ...interface{}) Request {
	val := r.Call("clone", args...)
	return JSValueToRequest(val.JSValue())
}
func (r Request) GetCredentials() RequestCredentials {
	val := r.Get("credentials")
	return JSValueToRequestCredentials(val.JSValue())
}
func (r Request) GetDestination() RequestDestination {
	val := r.Get("destination")
	return JSValueToRequestDestination(val.JSValue())
}
func (r Request) FormData(args ...interface{}) {
	r.Call("formData", args...)
}
func (r Request) GetHeaders() Headers {
	val := r.Get("headers")
	return JSValueToHeaders(val.JSValue())
}
func (r Request) GetIntegrity() string {
	val := r.Get("integrity")
	return val.String()
}
func (r Request) GetIsHistoryNavigation() bool {
	val := r.Get("isHistoryNavigation")
	return val.Bool()
}
func (r Request) GetIsReloadNavigation() bool {
	val := r.Get("isReloadNavigation")
	return val.Bool()
}
func (r Request) Json(args ...interface{}) {
	r.Call("json", args...)
}
func (r Request) GetKeepalive() bool {
	val := r.Get("keepalive")
	return val.Bool()
}
func (r Request) GetMethod() []byte {
	val := r.Get("method")
	return []byte(val.String())
}
func (r Request) GetMode() RequestMode {
	val := r.Get("mode")
	return JSValueToRequestMode(val.JSValue())
}
func (r Request) GetRedirect() RequestRedirect {
	val := r.Get("redirect")
	return JSValueToRequestRedirect(val.JSValue())
}
func (r Request) GetReferrer() string {
	val := r.Get("referrer")
	return val.String()
}
func (r Request) GetReferrerPolicy() ReferrerPolicy {
	val := r.Get("referrerPolicy")
	return JSValueToReferrerPolicy(val.JSValue())
}
func (r Request) GetSignal() AbortSignal {
	val := r.Get("signal")
	return JSValueToAbortSignal(val.JSValue())
}
func (r Request) Text(args ...interface{}) {
	r.Call("text", args...)
}
func (r Request) GetUrl() string {
	val := r.Get("url")
	return val.String()
}
