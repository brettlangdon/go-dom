// Code generated DO NOT EDIT
// navigator.go
package dom

import "syscall/js"

type NavigatorIFace interface {
	GetAppCodeName() string
	GetAppName() string
	GetAppVersion() string
	GetCookieEnabled() bool
	GetHardwareConcurrency() int
	JavaEnabled(args ...interface{}) bool
	GetLanguage() string
	GetLanguages()
	GetMimeTypes() MimeTypeArray
	GetOnLine() bool
	GetPlatform() string
	GetPlugins() PluginArray
	GetProduct() string
	GetProductSub() string
	RegisterProtocolHandler(args ...interface{})
	GetStorage() StorageManager
	UnregisterProtocolHandler(args ...interface{})
	GetUserAgent() string
	GetVendor() string
	GetVendorSub() string
}
type Navigator struct {
	Value
}

func JSValueToNavigator(val js.Value) Navigator { return Navigator{Value: Value{Value: val}} }
func (v Value) AsNavigator() Navigator          { return Navigator{Value: v} }
func (n Navigator) GetAppCodeName() string {
	val := n.Get("appCodeName")
	return val.String()
}
func (n Navigator) GetAppName() string {
	val := n.Get("appName")
	return val.String()
}
func (n Navigator) GetAppVersion() string {
	val := n.Get("appVersion")
	return val.String()
}
func (n Navigator) GetCookieEnabled() bool {
	val := n.Get("cookieEnabled")
	return val.Bool()
}
func (n Navigator) GetHardwareConcurrency() int {
	val := n.Get("hardwareConcurrency")
	return val.Int()
}
func (n Navigator) JavaEnabled(args ...interface{}) bool {
	val := n.Call("javaEnabled", args...)
	return val.Bool()
}
func (n Navigator) GetLanguage() string {
	val := n.Get("language")
	return val.String()
}
func (n Navigator) GetLanguages() Value {
	val := n.Get("languages")
	return val
}
func (n Navigator) GetMimeTypes() MimeTypeArray {
	val := n.Get("mimeTypes")
	return JSValueToMimeTypeArray(val.JSValue())
}
func (n Navigator) GetOnLine() bool {
	val := n.Get("onLine")
	return val.Bool()
}
func (n Navigator) GetPlatform() string {
	val := n.Get("platform")
	return val.String()
}
func (n Navigator) GetPlugins() PluginArray {
	val := n.Get("plugins")
	return JSValueToPluginArray(val.JSValue())
}
func (n Navigator) GetProduct() string {
	val := n.Get("product")
	return val.String()
}
func (n Navigator) GetProductSub() string {
	val := n.Get("productSub")
	return val.String()
}
func (n Navigator) RegisterProtocolHandler(args ...interface{}) {
	n.Call("registerProtocolHandler", args...)
}
func (n Navigator) GetStorage() StorageManager {
	val := n.Get("storage")
	return JSValueToStorageManager(val.JSValue())
}
func (n Navigator) UnregisterProtocolHandler(args ...interface{}) {
	n.Call("unregisterProtocolHandler", args...)
}
func (n Navigator) GetUserAgent() string {
	val := n.Get("userAgent")
	return val.String()
}
func (n Navigator) GetVendor() string {
	val := n.Get("vendor")
	return val.String()
}
func (n Navigator) GetVendorSub() string {
	val := n.Get("vendorSub")
	return val.String()
}
