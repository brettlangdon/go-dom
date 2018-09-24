// Code generated DO NOT EDIT
// stylesheetlist.go
package dom

import "syscall/js"

type StyleSheetListIFace interface {
	Item(args ...interface{}) StyleSheet
	GetLength() int
}
type StyleSheetList struct {
	Value
}

func JSValueToStyleSheetList(val js.Value) StyleSheetList {
	return StyleSheetList{Value: JSValueToValue(val)}
}
func (v Value) AsStyleSheetList() StyleSheetList { return StyleSheetList{Value: v} }
func NewStyleSheetList(args ...interface{}) StyleSheetList {
	return StyleSheetList{Value: JSValueToValue(js.Global().Get("StyleSheetList").New(args...))}
}
func (s StyleSheetList) Item(args ...interface{}) StyleSheet {
	val := s.Call("item", args...)
	return JSValueToStyleSheet(val.JSValue())
}
func (s StyleSheetList) GetLength() int {
	val := s.Get("length")
	return val.Int()
}
