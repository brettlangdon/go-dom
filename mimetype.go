// Code generated DO NOT EDIT
// mimetype.go
package dom

import "syscall/js"

type MimeTypeIFace interface {
	GetDescription() string
	GetEnabledPlugin() Plugin
	GetSuffixes() string
	GetType() string
}
type MimeType struct {
	Value
}

func JSValueToMimeType(val js.Value) MimeType { return MimeType{Value: Value{Value: val}} }
func (v Value) AsMimeType() MimeType          { return MimeType{Value: v} }
func (m MimeType) GetDescription() string {
	val := m.Get("description")
	return val.String()
}
func (m MimeType) GetEnabledPlugin() Plugin {
	val := m.Get("enabledPlugin")
	return JSValueToPlugin(val.JSValue())
}
func (m MimeType) GetSuffixes() string {
	val := m.Get("suffixes")
	return val.String()
}
func (m MimeType) GetType() string {
	val := m.Get("type")
	return val.String()
}
