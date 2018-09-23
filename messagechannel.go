// Code generated DO NOT EDIT
// messagechannel.go
package dom

import "syscall/js"

type MessageChannelIFace interface {
	GetPort1() MessagePort
	GetPort2() MessagePort
}
type MessageChannel struct {
	Value
}

func jsValueToMessageChannel(val js.Value) MessageChannel {
	return MessageChannel{Value: Value{Value: val}}
}
func (v Value) AsMessageChannel() MessageChannel { return MessageChannel{Value: v} }
func (m MessageChannel) GetPort1() MessagePort {
	val := m.Get("port1")
	return jsValueToMessagePort(val.JSValue())
}
func (m MessageChannel) GetPort2() MessagePort {
	val := m.Get("port2")
	return jsValueToMessagePort(val.JSValue())
}
