// Code generated DO NOT EDIT
// workernavigator.go
package dom

import "syscall/js"

type WorkerNavigatorIFace interface {
	GetAppCodeName() string
	GetAppName() string
	GetAppVersion() string
	GetHardwareConcurrency() int
	GetLanguage() string
	GetLanguages()
	GetOnLine() bool
	GetPlatform() string
	GetProduct() string
	GetProductSub() string
	GetStorage() StorageManager
	GetUserAgent() string
	GetVendor() string
	GetVendorSub() string
}
type WorkerNavigator struct {
	Value
}

func JSValueToWorkerNavigator(val js.Value) WorkerNavigator {
	return WorkerNavigator{Value: Value{Value: val}}
}
func (v Value) AsWorkerNavigator() WorkerNavigator { return WorkerNavigator{Value: v} }
func (w WorkerNavigator) GetAppCodeName() string {
	val := w.Get("appCodeName")
	return val.String()
}
func (w WorkerNavigator) GetAppName() string {
	val := w.Get("appName")
	return val.String()
}
func (w WorkerNavigator) GetAppVersion() string {
	val := w.Get("appVersion")
	return val.String()
}
func (w WorkerNavigator) GetHardwareConcurrency() int {
	val := w.Get("hardwareConcurrency")
	return val.Int()
}
func (w WorkerNavigator) GetLanguage() string {
	val := w.Get("language")
	return val.String()
}
func (w WorkerNavigator) GetLanguages() Value {
	val := w.Get("languages")
	return val
}
func (w WorkerNavigator) GetOnLine() bool {
	val := w.Get("onLine")
	return val.Bool()
}
func (w WorkerNavigator) GetPlatform() string {
	val := w.Get("platform")
	return val.String()
}
func (w WorkerNavigator) GetProduct() string {
	val := w.Get("product")
	return val.String()
}
func (w WorkerNavigator) GetProductSub() string {
	val := w.Get("productSub")
	return val.String()
}
func (w WorkerNavigator) GetStorage() StorageManager {
	val := w.Get("storage")
	return JSValueToStorageManager(val.JSValue())
}
func (w WorkerNavigator) GetUserAgent() string {
	val := w.Get("userAgent")
	return val.String()
}
func (w WorkerNavigator) GetVendor() string {
	val := w.Get("vendor")
	return val.String()
}
func (w WorkerNavigator) GetVendorSub() string {
	val := w.Get("vendorSub")
	return val.String()
}
