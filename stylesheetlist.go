// Code generated DO NOT EDIT
// stylesheetlist.go
package dom

import "syscall/js"

type StyleSheetListIFace interface {
	Item(args ...interface{}) StyleSheet
	GetLength() float64
}
type StyleSheetList struct {
	Value
}

func jsValueToStyleSheetList(val js.Value) StyleSheetList {
	return StyleSheetList{Value: Value{Value: val}}
}
func (v Value) AsStyleSheetList() StyleSheetList { return StyleSheetList{Value: v} }
func (s StyleSheetList) Item(args ...interface{}) StyleSheet {
	val := s.Call("item", args...)
	return jsValueToStyleSheet(val.JSValue())
}
func (s StyleSheetList) GetLength() float64 {
	val := s.Get("length")
	return val.Float()
}
