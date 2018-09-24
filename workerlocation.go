// Code generated DO NOT EDIT
// workerlocation.go
package dom

import "syscall/js"

type WorkerLocationIFace interface {
	GetHash() string
	GetHost() string
	GetHostname() string
	GetHref() string
	GetOrigin() string
	GetPathname() string
	GetPort() string
	GetProtocol() string
	GetSearch() string
}
type WorkerLocation struct {
	Value
}

func JSValueToWorkerLocation(val js.Value) WorkerLocation {
	return WorkerLocation{Value: JSValueToValue(val)}
}
func (v Value) AsWorkerLocation() WorkerLocation { return WorkerLocation{Value: v} }
func NewWorkerLocation(args ...interface{}) WorkerLocation {
	return WorkerLocation{Value: JSValueToValue(js.Global().Get("WorkerLocation").New(args...))}
}
func (w WorkerLocation) GetHash() string {
	val := w.Get("hash")
	return val.String()
}
func (w WorkerLocation) GetHost() string {
	val := w.Get("host")
	return val.String()
}
func (w WorkerLocation) GetHostname() string {
	val := w.Get("hostname")
	return val.String()
}
func (w WorkerLocation) GetHref() string {
	val := w.Get("href")
	return val.String()
}
func (w WorkerLocation) GetOrigin() string {
	val := w.Get("origin")
	return val.String()
}
func (w WorkerLocation) GetPathname() string {
	val := w.Get("pathname")
	return val.String()
}
func (w WorkerLocation) GetPort() string {
	val := w.Get("port")
	return val.String()
}
func (w WorkerLocation) GetProtocol() string {
	val := w.Get("protocol")
	return val.String()
}
func (w WorkerLocation) GetSearch() string {
	val := w.Get("search")
	return val.String()
}
