// Code generated DO NOT EDIT
// location.go
package dom

import "syscall/js"

type LocationIFace interface {
	GetAncestorOrigins() DOMStringList
	Assign(args ...interface{})
	GetHash() string
	SetHash(string)
	GetHost() string
	SetHost(string)
	GetHostname() string
	SetHostname(string)
	GetHref() string
	SetHref(string)
	GetOrigin() string
	GetPathname() string
	SetPathname(string)
	GetPort() string
	SetPort(string)
	GetProtocol() string
	SetProtocol(string)
	Reload(args ...interface{})
	Replace(args ...interface{})
	GetSearch() string
	SetSearch(string)
}
type Location struct {
	Value
}

func JSValueToLocation(val js.Value) Location { return Location{Value: JSValueToValue(val)} }
func (v Value) AsLocation() Location          { return Location{Value: v} }
func NewLocation(args ...interface{}) Location {
	return Location{Value: JSValueToValue(js.Global().Get("Location").New(args...))}
}
func (l Location) GetAncestorOrigins() DOMStringList {
	val := l.Get("ancestorOrigins")
	return JSValueToDOMStringList(val.JSValue())
}
func (l Location) Assign(args ...interface{}) {
	l.Call("assign", args...)
}
func (l Location) GetHash() string {
	val := l.Get("hash")
	return val.String()
}
func (l Location) SetHash(val string) {
	l.Set("hash", val)
}
func (l Location) GetHost() string {
	val := l.Get("host")
	return val.String()
}
func (l Location) SetHost(val string) {
	l.Set("host", val)
}
func (l Location) GetHostname() string {
	val := l.Get("hostname")
	return val.String()
}
func (l Location) SetHostname(val string) {
	l.Set("hostname", val)
}
func (l Location) GetHref() string {
	val := l.Get("href")
	return val.String()
}
func (l Location) SetHref(val string) {
	l.Set("href", val)
}
func (l Location) GetOrigin() string {
	val := l.Get("origin")
	return val.String()
}
func (l Location) GetPathname() string {
	val := l.Get("pathname")
	return val.String()
}
func (l Location) SetPathname(val string) {
	l.Set("pathname", val)
}
func (l Location) GetPort() string {
	val := l.Get("port")
	return val.String()
}
func (l Location) SetPort(val string) {
	l.Set("port", val)
}
func (l Location) GetProtocol() string {
	val := l.Get("protocol")
	return val.String()
}
func (l Location) SetProtocol(val string) {
	l.Set("protocol", val)
}
func (l Location) Reload(args ...interface{}) {
	l.Call("reload", args...)
}
func (l Location) Replace(args ...interface{}) {
	l.Call("replace", args...)
}
func (l Location) GetSearch() string {
	val := l.Get("search")
	return val.String()
}
func (l Location) SetSearch(val string) {
	l.Set("search", val)
}
