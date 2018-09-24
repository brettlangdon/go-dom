// Code generated DO NOT EDIT
// radionodelist.go
package dom

import "syscall/js"

type RadioNodeListIFace interface {
	Item(args ...interface{}) Node
	GetLength() int
	GetValue() string
	SetValue(string)
}
type RadioNodeList struct {
	Value
	NodeList
}

func JSValueToRadioNodeList(val js.Value) RadioNodeList {
	return RadioNodeList{Value: JSValueToValue(val)}
}
func (v Value) AsRadioNodeList() RadioNodeList { return RadioNodeList{Value: v} }
func NewRadioNodeList(args ...interface{}) RadioNodeList {
	return RadioNodeList{Value: JSValueToValue(js.Global().Get("RadioNodeList").New(args...))}
}
func (r RadioNodeList) GetValue() string {
	val := r.Get("value")
	return val.String()
}
func (r RadioNodeList) SetValue(val string) {
	r.Set("value", val)
}
