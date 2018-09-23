// Code generated DO NOT EDIT
// stylesheet.go
package dom

import "syscall/js"

type StyleSheetIFace interface {
	GetDisabled() bool
	SetDisabled(bool)
	GetHref() string
	GetMedia() MediaList
	GetOwnerNode()
	GetParentStyleSheet() StyleSheet
	GetTitle() string
	GetType() string
}
type StyleSheet struct {
	Value
}

func jsValueToStyleSheet(val js.Value) StyleSheet { return StyleSheet{Value: Value{Value: val}} }
func (v Value) AsStyleSheet() StyleSheet          { return StyleSheet{Value: v} }
func (s StyleSheet) GetDisabled() bool {
	val := s.Get("disabled")
	return val.Bool()
}
func (s StyleSheet) SetDisabled(val bool) {
	s.Set("disabled", val)
}
func (s StyleSheet) GetHref() string {
	val := s.Get("href")
	return val.String()
}
func (s StyleSheet) GetMedia() MediaList {
	val := s.Get("media")
	return jsValueToMediaList(val.JSValue())
}
func (s StyleSheet) GetOwnerNode() Value {
	val := s.Get("ownerNode")
	return val
}
func (s StyleSheet) GetParentStyleSheet() StyleSheet {
	val := s.Get("parentStyleSheet")
	return jsValueToStyleSheet(val.JSValue())
}
func (s StyleSheet) GetTitle() string {
	val := s.Get("title")
	return val.String()
}
func (s StyleSheet) GetType() string {
	val := s.Get("type")
	return val.String()
}
