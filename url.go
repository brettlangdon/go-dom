// Code generated DO NOT EDIT
// url.go
package dom

import "syscall/js"

type URLIFace interface {
	CreateObjectURL(args ...interface{}) string
	GetHash() string
	SetHash(string)
	GetHost() string
	SetHost(string)
	GetHostname() string
	SetHostname(string)
	GetHref() string
	SetHref(string)
	GetOrigin() string
	GetPassword() string
	SetPassword(string)
	GetPathname() string
	SetPathname(string)
	GetPort() string
	SetPort(string)
	GetProtocol() string
	SetProtocol(string)
	RevokeObjectURL(args ...interface{})
	GetSearch() string
	SetSearch(string)
	GetSearchParams() URLSearchParams
	ToJSON(args ...interface{}) string
	GetUsername() string
	SetUsername(string)
}
type URL struct {
	Value
}

func JSValueToURL(val js.Value) URL { return URL{Value: JSValueToValue(val)} }
func (v Value) AsURL() URL          { return URL{Value: v} }
func NewURL(args ...interface{}) URL {
	return URL{Value: JSValueToValue(js.Global().Get("URL").New(args...))}
}
func (u URL) CreateObjectURL(args ...interface{}) string {
	val := u.Call("createObjectURL", args...)
	return val.String()
}
func (u URL) GetHash() string {
	val := u.Get("hash")
	return val.String()
}
func (u URL) SetHash(val string) {
	u.Set("hash", val)
}
func (u URL) GetHost() string {
	val := u.Get("host")
	return val.String()
}
func (u URL) SetHost(val string) {
	u.Set("host", val)
}
func (u URL) GetHostname() string {
	val := u.Get("hostname")
	return val.String()
}
func (u URL) SetHostname(val string) {
	u.Set("hostname", val)
}
func (u URL) GetHref() string {
	val := u.Get("href")
	return val.String()
}
func (u URL) SetHref(val string) {
	u.Set("href", val)
}
func (u URL) GetOrigin() string {
	val := u.Get("origin")
	return val.String()
}
func (u URL) GetPassword() string {
	val := u.Get("password")
	return val.String()
}
func (u URL) SetPassword(val string) {
	u.Set("password", val)
}
func (u URL) GetPathname() string {
	val := u.Get("pathname")
	return val.String()
}
func (u URL) SetPathname(val string) {
	u.Set("pathname", val)
}
func (u URL) GetPort() string {
	val := u.Get("port")
	return val.String()
}
func (u URL) SetPort(val string) {
	u.Set("port", val)
}
func (u URL) GetProtocol() string {
	val := u.Get("protocol")
	return val.String()
}
func (u URL) SetProtocol(val string) {
	u.Set("protocol", val)
}
func (u URL) RevokeObjectURL(args ...interface{}) {
	u.Call("revokeObjectURL", args...)
}
func (u URL) GetSearch() string {
	val := u.Get("search")
	return val.String()
}
func (u URL) SetSearch(val string) {
	u.Set("search", val)
}
func (u URL) GetSearchParams() URLSearchParams {
	val := u.Get("searchParams")
	return JSValueToURLSearchParams(val.JSValue())
}
func (u URL) ToJSON(args ...interface{}) string {
	val := u.Call("toJSON", args...)
	return val.String()
}
func (u URL) GetUsername() string {
	val := u.Get("username")
	return val.String()
}
func (u URL) SetUsername(val string) {
	u.Set("username", val)
}
