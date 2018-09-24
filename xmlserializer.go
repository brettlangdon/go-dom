// Code generated DO NOT EDIT
// xmlserializer.go
package dom

import "syscall/js"

type XMLSerializerIFace interface {
	SerializeToString(args ...interface{}) string
}
type XMLSerializer struct {
	Value
}

func JSValueToXMLSerializer(val js.Value) XMLSerializer {
	return XMLSerializer{Value: JSValueToValue(val)}
}
func (v Value) AsXMLSerializer() XMLSerializer { return XMLSerializer{Value: v} }
func NewXMLSerializer(args ...interface{}) XMLSerializer {
	return XMLSerializer{Value: JSValueToValue(js.Global().Get("XMLSerializer").New(args...))}
}
func (x XMLSerializer) SerializeToString(args ...interface{}) string {
	val := x.Call("serializeToString", args...)
	return val.String()
}
