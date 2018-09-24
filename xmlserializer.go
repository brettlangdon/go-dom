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
	return XMLSerializer{Value: Value{Value: val}}
}
func (v Value) AsXMLSerializer() XMLSerializer { return XMLSerializer{Value: v} }
func (x XMLSerializer) SerializeToString(args ...interface{}) string {
	val := x.Call("serializeToString", args...)
	return val.String()
}
