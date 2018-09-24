// Code generated DO NOT EDIT
// namednodemap.go
package dom

import "syscall/js"

type NamedNodeMapIFace interface {
	GetNamedItem(args ...interface{}) Attr
	GetNamedItemNS(args ...interface{}) Attr
	Item(args ...interface{}) Attr
	GetLength() int
	RemoveNamedItem(args ...interface{}) Attr
	RemoveNamedItemNS(args ...interface{}) Attr
	SetNamedItem(args ...interface{}) Attr
	SetNamedItemNS(args ...interface{}) Attr
}
type NamedNodeMap struct {
	Value
}

func JSValueToNamedNodeMap(val js.Value) NamedNodeMap { return NamedNodeMap{Value: JSValueToValue(val)} }
func (v Value) AsNamedNodeMap() NamedNodeMap          { return NamedNodeMap{Value: v} }
func NewNamedNodeMap(args ...interface{}) NamedNodeMap {
	return NamedNodeMap{Value: JSValueToValue(js.Global().Get("NamedNodeMap").New(args...))}
}
func (n NamedNodeMap) GetNamedItem(args ...interface{}) Attr {
	val := n.Call("getNamedItem", args...)
	return JSValueToAttr(val.JSValue())
}
func (n NamedNodeMap) GetNamedItemNS(args ...interface{}) Attr {
	val := n.Call("getNamedItemNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (n NamedNodeMap) Item(args ...interface{}) Attr {
	val := n.Call("item", args...)
	return JSValueToAttr(val.JSValue())
}
func (n NamedNodeMap) GetLength() int {
	val := n.Get("length")
	return val.Int()
}
func (n NamedNodeMap) RemoveNamedItem(args ...interface{}) Attr {
	val := n.Call("removeNamedItem", args...)
	return JSValueToAttr(val.JSValue())
}
func (n NamedNodeMap) RemoveNamedItemNS(args ...interface{}) Attr {
	val := n.Call("removeNamedItemNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (n NamedNodeMap) SetNamedItem(args ...interface{}) Attr {
	val := n.Call("setNamedItem", args...)
	return JSValueToAttr(val.JSValue())
}
func (n NamedNodeMap) SetNamedItemNS(args ...interface{}) Attr {
	val := n.Call("setNamedItemNS", args...)
	return JSValueToAttr(val.JSValue())
}
